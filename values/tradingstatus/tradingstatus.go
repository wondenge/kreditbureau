package tradingstatus

type (
	TradingStatus int
)

const (
	Dormant TradingStatus = iota
	ActivelyTrading
	UnderReceivershipLiquidation
	Liquidated
	UnderManagement
	Dissolved
	Bankrupt
)
