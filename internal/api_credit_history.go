package internal

type CreditHistory struct {
	SnapshotDate      SnapshotDate
	CustomerName      CustomerName
	Registrar         Registrar
	PrimaryID         PrimaryID
	PersonalPINNumber PersonalPINNumber
	AccountNumber     AccountNumber

	// The currency code for the facility's currency
	FacilityCurrency CurrencyCode

	// For Loans, the approved loan amount.
	// For Overdrafts, Credit Cards and Trade Finance products; the approved limits
	// For overdrafts with no limits, input the maximum amount overdrawn
	OriginalAmount float64

	// The current balance in the account.
	CurrentBalance float64

	// Arrears amount in a facility.
	// If the account is not in arrears, return 0.
	OverdueBalance float64

	// The number of days the account has been in arrears
	// (The Difference between the reporting date and the overdue date, calculated in days).
	// If the account is not in arrears, return 0.
	DaysInArrears int

	// The Number of missed Instalments in the Facility.
	InstalmentsInArrears int

	// The Account status as at the time of reporting.
	AccountStatus AccountStatus

	// Classification of the Account
	PrudentialAssetClassification AssetClassification

	// The date when the facility last received a payment.
	// If no amount has been paid into the account, leave the field blank.
	LastPaymentDate string

	// The Account Product Type
	AccountProductType AccountProductType

	// Instalment amount for Loans.
	InstalmentAmount float64

	// Any additional information on the facility
	AdditionalInformation string
}
