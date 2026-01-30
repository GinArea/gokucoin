package kucoinv1

type WsToken struct {
	Token           string
	InstanceServers []InstanceServers
}

type InstanceServers struct {
	Endpoint     string
	Encrypt      bool
	Protocol     string
	PingInterval int
	PingTimeout  int
}

func (c *Client) GetPrivateWsToken() Response[WsToken] {
	return Post(c, "bullet-private", nil, forward[WsToken])
}

func (c *Client) GetPublicWsToken() Response[WsToken] {
	return PostPub(c, "bullet-public", nil, forward[WsToken])
}
