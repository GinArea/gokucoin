package v1

// contractStatuc (https://docs.kucoin.com/futures/#get-open-contract-list)
type ContractStatus string

const (
	Init         ContractStatus = "Init"
	Open         ContractStatus = "Open"
	BeingSettled ContractStatus = "BeingSettled"
	Settled      ContractStatus = "Settled"
	Paused       ContractStatus = "Paused"
	Closed       ContractStatus = "Closed"
	CancelOnly   ContractStatus = "CancelOnly"
)
