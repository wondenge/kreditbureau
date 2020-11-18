package internal

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

type CollateralRegister struct {
	Lenders           Lenders
	ClientNumber      ClientNumber
	AccountNumber     AccountNumber
	UserID            UserID
	PersonalPINNumber PersonalPINNumber
	Registrar         Registrar
	Collateral        Collateral
}

type Collateral struct {

	// Type of Collateral
	CollateralType CollateralType

	// The Registration number or Reference Number to uniquely
	// identify the collateral.
	CollateralReferenceNumber string

	// Value of the collateral as at Last valuation
	CollateralLastValuationAmount float64

	// Currency Code of the Collateral currency.
	CollateralCurrency CurrencyCode

	// The Forced Sale value as per professional valuation.
	CollateralForcedSaleValue float64

	// Date when collateral is scheduled for valuation
	NextValuationDate string

	// Required if the Collateral has an expiry Date
	CollateralExpiryDate string

	// Brief description of instruments used to recover
	InstrumentOfClaimOrRecoveryType string

	// The Date the Collateral was last Valued
	LastValuationDate string

	// If the Collateral secures multiple loans
	SharedCollateral SharedCollateral

	// The value should be expressed as a percentage,
	// but the “%” sign should be excluded
	PortionOfCollateralShared int

	// If the collateral is part of a set of Collaterals
	MultipleCollateral CollateralPart
}
