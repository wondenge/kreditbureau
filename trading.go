package kreditbureau

type TradingStatus string

const (
	DormantTradingStatus           = "001=Dormant"
	ActivelyTrading                = "002=Actively Trading"
	UnderReceivershipOrLiquidation = "003=Under Receivership/Liquidation"
	Liquidated                     = "004=Liquidated"
	UnderManagement                = "005=Under Management"
	Dissolved                      = "006=Dissolved"
	Bankrupt                       = "007=Bankrupt"
)
