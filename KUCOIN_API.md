# KuCoin Classic REST API Reference

## Overview

This document covers the **Classic REST API** (production-ready). KuCoin also has a "Pro REST API" which is in beta and not recommended for production use.

### Base URLs

| Environment | Spot/Margin | Futures |
|-------------|-------------|---------|
| Production | `https://api.kucoin.com` | `https://api-futures.kucoin.com` |

### API Versions

- Spot: `/api/v1/`, `/api/v2/`, `/api/v3/` (varies by endpoint)
- Futures: `/api/v1/`

---

## Authentication

KuCoin uses HMAC-SHA256 signature with base64 encoding (similar to OKX).

### Required Headers

| Header | Description |
|--------|-------------|
| KC-API-KEY | API Key |
| KC-API-SIGN | Signature (base64) |
| KC-API-TIMESTAMP | Request timestamp (ms) |
| KC-API-PASSPHRASE | Encrypted passphrase (base64) |
| KC-API-KEY-VERSION | Set to `2` for v2 signature |

### Signature Generation

```
Pre-sign string = timestamp + method + requestPath + body
Signature = base64(HMAC-SHA256(pre-sign, secretKey))
Passphrase = base64(HMAC-SHA256(passphrase, secretKey))
```

**Example:**
```
timestamp: 1659523885000
method: POST
requestPath: /api/v1/orders
body: {"symbol":"BTC-USDT","side":"buy","type":"limit","price":"20000","size":"0.001"}

pre-sign = "1659523885000POST/api/v1/orders{\"symbol\":\"BTC-USDT\"...}"
```

---

## Rate Limits

- **Window:** 30-second rolling
- **VIP0:** 2,000-4,000 requests per 30s (weight-based)
- **Response:** HTTP 429 when exceeded
- **Headers:** `X-Rate-Limit-Limit`, `X-Rate-Limit-Remaining`, `X-Rate-Limit-Reset`

---

## Account & Funding

### User Info
Docs: https://www.kucoin.com/docs-new/rest/account-info/account-funding/get-account-summary-info
```
GET /api/v2/user-info
```
Returns user account info including level, sub-account count.

### Accounts List
Docs: https://www.kucoin.com/docs-new/rest/account-info/account-funding/get-account-list-spot
```
GET /api/v1/accounts
Query: type (main/trade/margin/isolated), currency
```
Returns all accounts with balance info.

### Account Balance (Single)
Docs: https://www.kucoin.com/docs-new/rest/account-info/account-funding/get-account-detail-spot
```
GET /api/v1/accounts/{accountId}
```

### Transferable Balance
Docs: https://www.kucoin.com/docs-new/rest/account-info/account-funding/get-transferable
```
GET /api/v1/accounts/transferable
Query: currency, type (MAIN/TRADE/MARGIN)
```

### Inner Transfer
Docs: https://www.kucoin.com/docs-new/rest/account-info/account-funding/inner-transfer
```
POST /api/v2/accounts/inner-transfer
Body: clientOid, currency, from, to, amount
```

### Deposit Address
Docs: https://www.kucoin.com/docs-new/rest/account-info/deposit/get-deposit-address
```
GET /api/v1/deposit-addresses
Query: currency, chain (optional)
```

### Create Deposit Address
Docs: https://www.kucoin.com/docs-new/rest/account-info/deposit/create-deposit-address
```
POST /api/v1/deposit-addresses
Body: currency, chain (optional)
```

### Deposits History
Docs: https://www.kucoin.com/docs-new/rest/account-info/deposit/get-deposit-list
```
GET /api/v1/deposits
Query: currency, status, startAt, endAt
```

### Withdrawals History
Docs: https://www.kucoin.com/docs-new/rest/account-info/withdrawal/get-withdrawal-list
```
GET /api/v1/withdrawals
Query: currency, status, startAt, endAt
```

### Apply Withdrawal
Docs: https://www.kucoin.com/docs-new/rest/account-info/withdrawal/apply-withdraw
```
POST /api/v1/withdrawals
Body: currency, address, amount, chain, memo (optional)
```

---

## Spot Trading

### Market Data (Public)

#### Symbols List
Docs: https://www.kucoin.com/docs-new/rest/spot-trading/market-data/get-all-symbols
```
GET /api/v2/symbols
Query: market (optional, e.g., "USDS")
```
Returns trading pairs with specs (baseCurrency, quoteCurrency, baseMinSize, priceIncrement, etc.)

#### All Tickers
Docs: https://www.kucoin.com/docs-new/rest/spot-trading/market-data/get-all-tickers
```
GET /api/v1/market/allTickers
```
Returns price summary for all symbols.

#### Single Ticker
Docs: https://www.kucoin.com/docs-new/rest/spot-trading/market-data/get-24hr-stats
```
GET /api/v1/market/stats
Query: symbol
```

#### Order Book (Level 2)
Docs (Partial): https://www.kucoin.com/docs-new/rest/spot-trading/market-data/get-part-orderbook
Docs (Full): https://www.kucoin.com/docs-new/rest/spot-trading/market-data/get-full-orderbook
```
GET /api/v1/market/orderbook/level2_20
GET /api/v1/market/orderbook/level2_100
GET /api/v3/market/orderbook/level2
Query: symbol
```
Level2_20/100 are public, full Level2 requires auth.

#### Candles (Klines)
Docs: https://www.kucoin.com/docs-new/rest/spot-trading/market-data/get-klines
```
GET /api/v1/market/candles
Query: symbol, type (1min/5min/15min/30min/1hour/2hour/4hour/6hour/8hour/12hour/1day/1week), startAt, endAt
```

#### Trade History
Docs: https://www.kucoin.com/docs-new/rest/spot-trading/market-data/get-trade-histories
```
GET /api/v1/market/histories
Query: symbol
```

### Orders (Private)

#### Place Order (HF)
Docs: https://www.kucoin.com/docs-new/rest/spot-trading/orders/add-order
```
POST /api/v1/hf/orders
Body: {
  clientOid: string (optional, unique ID, UUID recommended),
  side: "buy" | "sell",
  symbol: string,
  type: "limit" | "market",
  price: string (for limit),
  size: string (base currency qty),
  funds: string (quote currency qty, market only),
  timeInForce: "GTC" | "GTT" | "IOC" | "FOK",
  cancelAfter: long (seconds, requires GTT),
  postOnly: boolean,
  hidden: boolean,
  iceberg: boolean,
  visibleSize: string,
  stp: "CN" | "CO" | "CB" | "DC",
  tags: string (max 20 ASCII),
  remark: string (max 20 ASCII)
}
```
Returns: `{ orderId: string, clientOid: string }`

#### Place Multiple Orders
Docs: https://www.kucoin.com/docs-new/rest/spot-trading/orders/place-multiple-orders
```
POST /api/v1/orders/multi
Body: { symbol, orderList: [order...] }
```
Max 5 orders per request.

#### Cancel Order
Docs: https://www.kucoin.com/docs-new/rest/spot-trading/orders/cancel-order-by-orderid
```
DELETE /api/v1/orders/{orderId}
```

#### Cancel by Client OID
Docs: https://www.kucoin.com/docs-new/rest/spot-trading/orders/cancel-order-by-clientoid
```
DELETE /api/v1/order/client-order/{clientOid}
```

#### Cancel All Orders
Docs: https://www.kucoin.com/docs-new/rest/spot-trading/orders/cancel-all-orders
```
DELETE /api/v1/orders
Query: symbol, tradeType (TRADE/MARGIN_TRADE/MARGIN_ISOLATED_TRADE)
```

#### Get Order by ClientOid (HF)
Docs: https://www.kucoin.com/docs-new/rest/spot-trading/orders/get-order-by-clientoid
```
GET /api/v1/hf/orders/client-order/{clientOid}
Query: symbol (required)
```
Note: Data available for 3x24 hours only for inactive orders.

#### Get by Client OID
Docs: https://www.kucoin.com/docs-new/rest/spot-trading/orders/get-order-by-clientoid
```
GET /api/v1/order/client-order/{clientOid}
```

#### List Orders
Docs: https://www.kucoin.com/docs-new/rest/spot-trading/orders/get-order-list
```
GET /api/v1/orders
Query: status (active/done), symbol, side, type, tradeType, startAt, endAt
```

#### Recent Orders (24h)
Docs: https://www.kucoin.com/docs-new/rest/spot-trading/orders/get-recent-orders-list
```
GET /api/v1/limit/orders
```

### Fills (Private)

#### List Fills
Docs: https://www.kucoin.com/docs-new/rest/spot-trading/fills/get-filled-list
```
GET /api/v1/fills
Query: orderId, symbol, side, type, tradeType, startAt, endAt
```

#### Recent Fills (24h)
Docs: https://www.kucoin.com/docs-new/rest/spot-trading/fills/get-recent-filled-list
```
GET /api/v1/limit/fills
```

### Stop Orders

#### Place Stop Order
Docs: https://www.kucoin.com/docs-new/rest/spot-trading/stop-order/place-stop-order
```
POST /api/v1/stop-order
Body: {
  clientOid, side, symbol, type,
  stopPrice: string,
  price: string (limit),
  size: string,
  stop: "loss" | "entry"
}
```

#### Cancel Stop Order
Docs: https://www.kucoin.com/docs-new/rest/spot-trading/stop-order/cancel-stop-order
```
DELETE /api/v1/stop-order/{orderId}
```

#### Cancel All Stop Orders
Docs: https://www.kucoin.com/docs-new/rest/spot-trading/stop-order/cancel-stop-orders
```
DELETE /api/v1/stop-order/cancel
Query: symbol, tradeType, orderIds
```

#### Get Stop Order
Docs: https://www.kucoin.com/docs-new/rest/spot-trading/stop-order/get-stop-order
```
GET /api/v1/stop-order/{orderId}
```

#### List Stop Orders
Docs: https://www.kucoin.com/docs-new/rest/spot-trading/stop-order/get-stop-order-list
```
GET /api/v1/stop-order
Query: symbol, side, type, tradeType, startAt, endAt
```

---

## Margin Trading

### Cross Margin

#### Borrow
Docs: https://www.kucoin.com/docs-new/rest/margin-trading/credit/margin-borrow
```
POST /api/v3/margin/borrow
Body: currency, size, timeInForce (IOC/FOK)
```

#### Repay
Docs: https://www.kucoin.com/docs-new/rest/margin-trading/credit/margin-repay
```
POST /api/v3/margin/repay
Body: currency, size
```

#### Borrow History
Docs: https://www.kucoin.com/docs-new/rest/margin-trading/credit/get-margin-borrow-history
```
GET /api/v3/margin/borrow
Query: currency, status, startAt, endAt
```

#### Repay History
Docs: https://www.kucoin.com/docs-new/rest/margin-trading/credit/get-margin-repay-history
```
GET /api/v3/margin/repay
Query: currency, startAt, endAt
```

### Isolated Margin

#### Borrow
Docs: https://www.kucoin.com/docs-new/rest/margin-trading/isolated-margin/isolated-margin-borrow
```
POST /api/v3/isolated/borrow
Body: symbol, currency, size, borrowStrategy (FOK/IOC)
```

#### Repay
Docs: https://www.kucoin.com/docs-new/rest/margin-trading/isolated-margin/isolated-margin-repay
```
POST /api/v3/isolated/repay
Body: symbol, currency, size
```

#### Isolated Accounts
Docs: https://www.kucoin.com/docs-new/rest/margin-trading/isolated-margin/get-isolated-margin-account
```
GET /api/v3/isolated/accounts
Query: symbol, quoteCurrency, queryType
```

---

## Futures Trading

Base URL: `https://api-futures.kucoin.com`

### Market Data (Public)

#### Active Contracts
Docs: https://www.kucoin.com/docs-new/rest/futures-trading/market-data/get-all-symbols
```
GET /api/v1/contracts/active
```
Returns all active futures contracts with specs.

#### Contract Detail
Docs: https://www.kucoin.com/docs-new/rest/futures-trading/market-data/get-symbol
```
GET /api/v1/contracts/{symbol}
```

#### Ticker
Docs: https://www.kucoin.com/docs-new/rest/futures-trading/market-data/get-ticker
```
GET /api/v1/ticker
Query: symbol
```

#### Order Book (Level 2)
Docs: https://www.kucoin.com/docs-new/rest/futures-trading/market-data/get-part-orderbook
```
GET /api/v1/level2/snapshot
Query: symbol
```

#### Klines
Docs: https://www.kucoin.com/docs-new/rest/futures-trading/market-data/get-klines
```
GET /api/v1/kline/query
Query: symbol, granularity (1/5/15/30/60/120/240/480/720/1440/10080), from, to
```
Granularity in minutes.

#### Trade History
Docs: https://www.kucoin.com/docs-new/rest/futures-trading/market-data/get-trade-history
```
GET /api/v1/trade/history
Query: symbol
```

#### Index Price
Docs: https://www.kucoin.com/docs-new/rest/futures-trading/market-data/get-index-list
```
GET /api/v1/index/query
Query: symbol, startAt, endAt
```

#### Mark Price
Docs: https://www.kucoin.com/docs-new/rest/futures-trading/market-data/get-current-mark-price
```
GET /api/v1/mark-price/{symbol}/current
```

#### Funding Rate (Current)
Docs: https://www.kucoin.com/docs-new/rest/futures-trading/market-data/get-current-funding-rate
```
GET /api/v1/funding-rate/{symbol}/current
```

#### Funding Rate (History)
Docs: https://www.kucoin.com/docs-new/rest/futures-trading/market-data/get-public-funding-history
```
GET /api/v1/contract/funding-rates
Query: symbol, from, to
```

### Account (Private)

#### Account Overview
Docs: https://www.kucoin.com/docs-new/rest/futures-trading/account/get-account-overview
```
GET /api/v1/account-overview
Query: currency (optional, default XBT)
```

#### Transaction History
Docs: https://www.kucoin.com/docs-new/rest/futures-trading/account/get-transaction-history
```
GET /api/v1/transaction-history
Query: type, currency, startAt, endAt
```

### Orders (Private)

#### Place Order
Docs: https://www.kucoin.com/docs-new/rest/futures-trading/orders/place-order
```
POST /api/v1/orders
Body: {
  clientOid: string (required),
  side: "buy" | "sell",
  symbol: string,
  type: "limit" | "market",
  leverage: string,
  price: string (limit),
  size: integer (contracts),
  timeInForce: "GTC" | "IOC" | "FOK",
  postOnly: boolean,
  hidden: boolean,
  reduceOnly: boolean,
  closeOrder: boolean
}
```

#### Cancel Order
Docs: https://www.kucoin.com/docs-new/rest/futures-trading/orders/cancel-order-by-id
```
DELETE /api/v1/orders/{orderId}
```

#### Cancel All Orders
Docs: https://www.kucoin.com/docs-new/rest/futures-trading/orders/cancel-multiple-orders
```
DELETE /api/v3/orders
Query: symbol
```
Note: `DELETE /api/v1/orders` deprecated since December 2024.

#### Get Order
Docs: https://www.kucoin.com/docs-new/rest/futures-trading/orders/get-order-by-orderid
```
GET /api/v1/orders/{orderId}
```

#### Get by Client OID
Docs: https://www.kucoin.com/docs-new/rest/futures-trading/orders/get-order-by-clientoid
```
GET /api/v1/orders/byClientOid
Query: clientOid
```

#### List Orders
Docs: https://www.kucoin.com/docs-new/rest/futures-trading/orders/get-order-list
```
GET /api/v1/orders
Query: status (active/done), symbol, side, type, startAt, endAt
```

#### Recent Done Orders (24h)
Docs: https://www.kucoin.com/docs-new/rest/futures-trading/orders/get-recent-closed-orders
```
GET /api/v1/recentDoneOrders
```

### Stop Orders (Private)

#### Place Stop Order
Docs: https://www.kucoin.com/docs-new/rest/futures-trading/orders/place-order
```
POST /api/v1/orders
Body: {
  ...order fields...,
  stop: "up" | "down",
  stopPriceType: "TP" | "IP" | "MP",
  stopPrice: string
}
```

#### Get Untriggered Stop Orders
Docs: https://www.kucoin.com/docs-new/rest/futures-trading/orders/get-stop-order-list
```
GET /api/v1/stopOrders
Query: symbol, side, type, startAt, endAt
```

### Positions (Private)

#### Get Position
Docs: https://www.kucoin.com/docs-new/rest/futures-trading/positions/get-position-details
```
GET /api/v1/position
Query: symbol
```

#### Get All Positions
Docs: https://www.kucoin.com/docs-new/rest/futures-trading/positions/get-position-list
```
GET /api/v1/positions
```

#### Auto-Deposit Status
Docs: https://www.kucoin.com/docs-new/rest/futures-trading/positions/get-auto-deposit-status
```
GET /api/v1/position/autoDepositStatus
Query: symbol
```

#### Enable Auto-Deposit
Docs: https://www.kucoin.com/docs-new/rest/futures-trading/positions/enable-auto-deposit-margin
```
POST /api/v1/position/margin/auto-deposit-status
Body: symbol, status (true/false)
```

#### Add Margin
Docs: https://www.kucoin.com/docs-new/rest/futures-trading/positions/add-margin-manually
```
POST /api/v1/position/margin/deposit-margin
Body: symbol, margin, bizNo
```

### Risk Limit

#### Get Risk Limit
Docs: https://www.kucoin.com/docs-new/rest/futures-trading/risk-limit/get-futures-risk-limit
```
GET /api/v1/contracts/risk-limit/{symbol}
```

#### Adjust Risk Limit
Docs: https://www.kucoin.com/docs-new/rest/futures-trading/risk-limit/modify-futures-risk-limit
```
POST /api/v1/position/risk-limit-level/change
Body: symbol, level
```

### Fills (Private)

#### List Fills
Docs: https://www.kucoin.com/docs-new/rest/futures-trading/fills/get-filled-list
```
GET /api/v1/fills
Query: orderId, symbol, side, type, startAt, endAt
```

#### Recent Fills (24h)
Docs: https://www.kucoin.com/docs-new/rest/futures-trading/fills/get-recent-filled-list
```
GET /api/v1/recentFills
```

---

## WebSocket API

Docs (Spot): https://www.kucoin.com/docs-new/websocket/spot-trading
Docs (Futures): https://www.kucoin.com/docs-new/websocket/futures-trading

### Connection

KuCoin WebSocket requires a token obtained from REST API:

#### Spot Token
```
POST /api/v1/bullet-public  (public)
POST /api/v1/bullet-private (private, requires auth)
```

Returns:
```json
{
  "token": "xxx",
  "instanceServers": [
    {"endpoint": "wss://...", "pingInterval": 18000, "pingTimeout": 10000}
  ]
}
```

#### Futures Token
```
POST /api/v1/bullet-public  (on api-futures.kucoin.com)
POST /api/v1/bullet-private
```

### WebSocket URLs

Connect with token as query param:
```
wss://{endpoint}?token={token}&connectId={unique_id}
```

### Ping/Pong

Send ping every `pingInterval` ms:
```json
{"id": "123", "type": "ping"}
```

### Subscription Format

```json
{
  "id": "unique_id",
  "type": "subscribe",
  "topic": "/market/ticker:BTC-USDT",
  "privateChannel": false,
  "response": true
}
```

### Public Channels (Spot)

#### Ticker
```
/market/ticker:{symbol}
/market/ticker:all
```

#### Order Book (Level 2)
```
/market/level2:{symbol}
```

#### Candles
```
/market/candles:{symbol}_{interval}
```
Intervals: 1min, 5min, 15min, 30min, 1hour, 2hour, 4hour, 6hour, 8hour, 12hour, 1day, 1week

#### Match (Trades)
```
/market/match:{symbol}
```

#### Snapshot (Level 2 Depth 5)
```
/spotMarket/level2Depth5:{symbol}
/spotMarket/level2Depth50:{symbol}
```

### Private Channels (Spot)

#### Orders
```
/spotMarket/tradeOrders
/spotMarket/tradeOrdersV2
```

#### Balance
```
/account/balance
```

#### Stop Order
```
/spotMarket/advancedOrders
```

### Public Channels (Futures)

#### Ticker
```
/contractMarket/tickerV2:{symbol}
```

#### Order Book (Level 2)
```
/contractMarket/level2:{symbol}
```

#### Level 2 Depth
```
/contractMarket/level2Depth5:{symbol}
/contractMarket/level2Depth50:{symbol}
```

#### Execution
```
/contractMarket/execution:{symbol}
```

#### Mark Price
```
/contract/instrument:{symbol}
```

#### Funding Rate
```
/contract/funding:{symbol}
```

### Private Channels (Futures)

#### Orders
```
/contractMarket/tradeOrders
/contractMarket/tradeOrdersV2
```

#### Positions
```
/contract/position:{symbol}
/contract/positionAll
```

#### Balance
```
/contractAccount/wallet
```

#### Stop Orders
```
/contractMarket/advancedOrders
```

---

## Error Codes

### Common Codes

| Code | Description |
|------|-------------|
| 200000 | Success |
| 400001 | Invalid parameter |
| 400002 | Invalid request |
| 400003 | Invalid timestamp |
| 400004 | Invalid signature |
| 400005 | Invalid passphrase |
| 400006 | Invalid API key |
| 400007 | Invalid IP |
| 411100 | User frozen |
| 415000 | Unsupported media type |
| 429000 | Too many requests |
| 500000 | Internal error |

### Trading Codes

| Code | Description |
|------|-------------|
| 200004 | Insufficient balance |
| 300000 | Order not found |
| 400100 | Order already done |
| 400200 | Order amount too small |
| 400500 | Order quantity invalid |
| 400600 | Symbol not traded |
| 400700 | Price out of range |

---

## Response Format

### Success
```json
{
  "code": "200000",
  "data": { ... }
}
```

### Error
```json
{
  "code": "400001",
  "msg": "Invalid parameter"
}
```

### Paginated
```json
{
  "code": "200000",
  "data": {
    "currentPage": 1,
    "pageSize": 50,
    "totalNum": 100,
    "totalPage": 2,
    "items": [ ... ]
  }
}
```

---

## Comparison with Other Exchanges

| Aspect | KuCoin | Bybit | OKX |
|--------|--------|-------|-----|
| Signature | HMAC-SHA256 base64 | HMAC-SHA256 hex | HMAC-SHA256 base64 |
| Passphrase | Yes (encrypted) | No | Yes (encrypted in v2) |
| Timestamp | Unix ms | Unix ms | ISO 8601 |
| WS Auth | Token via REST | Signature in connect | Signature in subscribe |
| Futures URL | Separate domain | Same domain | Same domain |
