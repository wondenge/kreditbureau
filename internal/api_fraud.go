package internal

type Fraud struct {
	Lenders       Lenders
	ClientNumber  ClientNumber
	AccountNumber AccountNumber

	// Brief Description of fraud type.
	FraudType string

	// Date on which Fraud took place.
	FraudIncidentDate string

	// Date on which Fraud was reported.
	FraudReportDate string

	// Amount involved in Fraud.
	Amount float64

	// Actual loss incurred in the fraud as at reporting date.
	LossAmount float64

	// Currency code of the Amount involved in the Fraud.
	// Default is Kenya Shillings.
	CurrencyCode CurrencyCode

	// A brief on the fraudulent incident.
	IncidentDetails string

	// A brief on the Forensic evidence.
	ForensicInformation string
	CustomerName        CustomerName

	// Primary Identification Document of the person Involved in the Fraud.
	PrimaryID PrimaryID

	// The Equivalent Amount involved in the fraud in KES.
	// If the Amounts involved in the Fraud (Field 4.7.10) are in a foreign
	// currency, use the exchange rate as at the Reporting Period
	// If the Amounts involved in the Fraud in Kenya Shillings, then report
	// the same value as in Field 4.7.10
	AmountInKES float64
}
