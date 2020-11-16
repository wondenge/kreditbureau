package kreditbureau

type CollateralType string
type SharedCollateral string
type CollateralPart string

const (
	Land                      CollateralType = "A – Land"
	PropertyAndBuildings                     = "B – Property and Buildings"
	Cash                                     = "C – Cash"
	TreasuryBillsAndBonds                    = "D – Treasury Bills and Bonds"
	Shares                                   = "E – Shares"
	Debentures                               = "F – Debentures"
	ChattelsOverMovableAssets                = "G – Chattels/Charges over movable assets" // excluding vehicles
	ChattelsOverVehicles                     = "H – Chattels over Vehicles"
	InsurancePolicy                          = "I – Insurance Policy"
)

const (
	CollateralIsShared  SharedCollateral = "Y= Collateral is shared"
	CollateralNotShared                  = "N= Collateral is not shared"
)

const (
	PartOfSetCollaterals CollateralPart = "Y= Part of a Set of Collaterals"
	SoleCollateral                      = "N= Sole Collateral"
)
