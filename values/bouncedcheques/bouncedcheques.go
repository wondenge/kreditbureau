package bouncedcheques

type (
	BouncedChequeReason int
)

const (
	ChequeUnpaid                        BouncedChequeReason = 30
	DateExpiredChequeStale                                  = 31
	ChequePostDated                                         = 32
	DateIrregular                                           = 33
	PayeeNameIncomplete                                     = 37
	PayeeNameIrregular                                      = 38
	AmountsInWordsAndFiguresDiffer                          = 40
	AmountInWordsRequired                                   = 41
	AmountInFiguresRequired                                 = 42
	AmountInFiguresIrregularIncomplete                      = 43
	SignatureDiffers                                        = 53
	DrawerSignatureRequired                                 = 54
	NotSignedInAccordanceWithMandate                        = 55
	EndorsementIrregular                                    = 56
	AlterationRequiresDrawersSignature                      = 57
	PayeeNameRequired                                       = 58
	EffectsNotCleared                                       = 62
	InsufficientFunds                                       = 63
	AccountDormant                                          = 64
	ChequeRepresentedMoreThanOnce                           = 66
	InvalidAccountNumber                                    = 69
	TitleOfAccountIrregular                                 = 70
	OriginatorReferenceNotQuoted                            = 71
	AccountTransferred                                      = 72
	CustomerDrawerDeceased                                  = 73
	AccountClosed                                           = 74
	TitleOfAccountRequired                                  = 75
	DebitInExcessOfDirectDebitAuthority                     = 76
	FrozenAccount                                           = 77
	PaymentStoppedByDrawer                                  = 79
	ConfirmationAwaited                                     = 80
)
