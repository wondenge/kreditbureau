package kreditbureau

type ChequeType string
type BouncedChequeReason string

const (
	CorporateCheque ChequeType = "A - Corporate"
	PersonalCheque             = "B - Personal"
	DividendCheque             = "C - Dividend"
)

const (
	ChequeUnpaidBecauseOfSuspectedCriminalActivity BouncedChequeReason = "30 – Cheque Unpaid Because of Suspected Criminal Activity"
	DateExpiredOrChequeStale                                           = "31 – Date Expired/Cheque Stale"
	ChequePostDated                                                    = "32 – Cheque Post-dated"
	DateIrregular                                                      = "33 – Date Irregular"
	PayeeNameIncomplete                                                = "37 – Payee Name Incomplete"
	PayeeNameIrregular                                                 = "38 – Payee Name Irregular"
	AmountsInWordsAndFiguresDiffer                                     = "40 – Amounts in Words and Figures Differ"
	AmountInWordsRequired                                              = "41 – Amount in Words required"
	AmountInFiguresRequired                                            = "42 – Amount in Figures required"
	AmountInFiguresIrregularOrIncomplete                               = "43 – Amount in Figures Irregular/incomplete"
	SignatureDiffers                                                   = "53 – Signature Differs"
	DrawerSignatureRequired                                            = "54 – Drawer Signature Required"
	NotSignedInAccordanceWithMandate                                   = "55 – Not signed in accordance with Mandate"
	EndorsementIrregular                                               = "56 – Endorsement Irregular"
	AlterationRequiresDrawersSignature                                 = "57 – Alteration, requires drawers signature"
	PayeeNameRequired                                                  = "58 – Payee Name Required"
	EffectsNotCleared                                                  = "62 – Effects not Cleared"
	InsufficientFundsReferToDrawer                                     = "63 – Insufficient Funds – Refer to Drawer"
	AccountDormantReferToDrawer                                        = "64 – Account Dormant – Refer to Drawer"
	ChequeRepresentedMoreThanOnce                                      = "66 – Cheque Re-presented more than Once"
	InvalidAccountNumber                                               = "69 – Invalid Account Number"
	TitleOfAccountIrregular                                            = "70 – Title of Account Irregular"
	OriginatorReferenceNotQuoted                                       = "71 – Originator Reference Not Quoted"
	AccountTransferred                                                 = "72 – Account Transferred"
	CustomerOrDrawerDeceased                                           = "73 – Customer/Drawer Deceased"
	AccountClosed                                                      = "74 – Account Closed"
	TitleOfAccountRequired                                             = "75 – Title of Account Required"
	DebitInExcessOfDirectDebitAuthority                                = "76 – Debit in Excess of Direct Debit Authority"
	FrozenAccount                                                      = "77 – Frozen Account"
	PaymentStoppedByDrawer                                             = "79 – Payment Stopped by Drawer"
	ConfirmationAwaited                                                = "80 – Confirmation Awaited"
)

type BouncedCheque struct {

	// The Type of Drawer.
	ClientType   ClientType
	CustomerName CustomerName
	CompanyName  CompanyName
	PrimaryID    PrimaryID

	// Bank and Branch code on Bounced Cheque
	BranchCodeOnCheque string

	// Client Reference Code Linking Drawer to Banking System
	ClientNumber ClientNumber

	// Bank Account from which the Bounced Cheque was drawn
	AccountNumber AccountNumber

	// Type of Cheque.
	ChequeAccountType ChequeType

	// Amount on cheque
	ChequeAmount float64

	// Cheque Number
	ChequeNumber string

	// ISO Currency Code
	ChequeCurrency CurrencyCode

	// Date on Bounced Cheque
	ChequeDate string

	// The Date the Cheque was unpaid by the Drawee Bank
	ChequeBounceDate string

	// Reason Code for Bouncing cheque
	ChequeBounceReasonCodes BouncedChequeReason

	// The Loan Account to which the Cheque was deposited to.
	LoanAccount string

	// The amount on the Cheque in Kenya Shillings.
	// If the Cheque is drawn in foreign currency, use the Exchange Rate on the Reporting Date.
	// If the Cheque is drawn in Kenya Shillings, report the actual amount on the cheque.
	ChequeAmountInKES float64
}
