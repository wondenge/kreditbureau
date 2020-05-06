package design

import (
	. "goa.design/goa/v3/dsl"
)

// The account cannot be updated until the dispute is resolved.
// The account Status cannot be Closed, Fully Settled with days
// in arrears greater than 0

/* Company Account Status */

// A - Closed – No more admin processes running such as instalment demands
// or interest charges to account, and no further facilities can be offered
// on this account.
var ClosedCompanyAccountStatus = Type("ClosedCompanyAccountStatus", func() {
	Description("Closed: No further facilities can be offered on this account")

	// Cannot be in the future
	// Not greater than snapshot date
	Attribute("AccountStatusDate", String, func() {
		Description("Date of Status Change")
		MaxLength(8)
		Meta("rpc:tag", "49")
	})

	// If Option “A” (closed) on Field 4.1.48 (Account Status) Must be filled
	Attribute("AccountClosureReason", String, func() {
		Description("Reason for Account Closure")
		MinLength(1)
		MaxLength(50)
		Meta("rpc:tag", "50")
	})

	Required("AccountStatusDate", "AccountClosureReason")

})

// C - Write-Off – For facilities that don’t form part of the outstanding portfolio
// in the Balance Sheet, but are still outstanding in the books of accounts.
var WriteOffCompanyAccountStatus = Type("WriteOffCompanyAccountStatus", func() {
	Description("WriteOff: Facilities that don’t form part of the outstanding portfolio")

	// Cannot be in the future
	// Not greater than snapshot date
	Attribute("AccountStatusDate", String, func() {
		Description("Date of Status Change")
		MaxLength(8)
		Meta("rpc:tag", "49")
	})

	Required("AccountStatusDate")
})

// D - Legal -with legal officer in court
var LegalCompanyAccountStatus = Type("LegalCompanyAccountStatus", func() {
	Description("Legal: With legal officer in court ")

	// Cannot be in the future
	// Not greater than snapshot date
	Attribute("AccountStatusDate", String, func() {
		Description("Date of Status Change")
		MaxLength(8)
		Meta("rpc:tag", "49")
	})

	Required("AccountStatusDate")
})

// E - Collection- with collection bureau
var CollectionCompanyAccountStatus = Type("CollectionCompanyAccountStatus", func() {
	Description("Collection: With collection bureau")

	// Cannot be in the future
	// Not greater than snapshot date
	Attribute("AccountStatusDate", String, func() {
		Description("Date of Status Change")
		MaxLength(8)
		Meta("rpc:tag", "49")
	})

	Required("AccountStatusDate")
})

// F – Active - For facilities that form part of the outstanding portfolio, and are
// reported in the Balance Sheet.
var ActiveCompanyAccountStatus = Type("ActiveCompanyAccountStatus", func() {
	Description("Active : Facilities that form part of the outstanding portfolio")

	// Cannot be in the future
	// Not greater than snapshot date
	Attribute("AccountStatusDate", String, func() {
		Description("Date of Status Change")
		MaxLength(8)
		Meta("rpc:tag", "49")
	})

	Required("AccountStatusDate")
})

// G – Facility Rescheduled – For Rescheduled/Restructured Facilities
var RescheduledCompanyAccountStatus = Type("RescheduledCompanyAccountStatus", func() {
	Description("Facility Rescheduled: Rescheduled/Restructured Facilities")

	// Cannot be in the future
	// Not greater than snapshot date
	Attribute("AccountStatusDate", String, func() {
		Description("Date of Status Change")
		MaxLength(8)
		Meta("rpc:tag", "49")
	})

	Required("AccountStatusDate")
})

// H –Settled – The facility has been cleared. This status can only be used for
// revolving facilities
var SettledCompanyAccountStatus = Type("SettledCompanyAccountStatus", func() {
	Description("Settled: Facility has been cleared")

	// Cannot be in the future
	// Not greater than snapshot date
	Attribute("AccountStatusDate", String, func() {
		Description("Date of Status Change")
		MaxLength(8)
		Meta("rpc:tag", "49")
	})

	Required("AccountStatusDate")
})

// J – Called Up : The facility has been called up. Once the client has paid up,
// the status should be updated to Option A, // H; or otherwise Option L
var CalledUpCompanyAccountStatus = Type("CalledUpCompanyAccountStatus", func() {
	Description("Called Up: Facility has been called up ")

	// Cannot be in the future
	// Not greater than snapshot date
	Attribute("AccountStatusDate", String, func() {
		Description("Date of Status Change")
		MaxLength(8)
		Meta("rpc:tag", "49")
	})

	Required("AccountStatusDate")
})

// K – Suspended – The facility has been put on hold for an indefinite period of time
var SuspendedCompanyAccountStatus = Type("SuspendedCompanyAccountStatus", func() {
	Description("Suspended: Facility put on hold for an indefinite period of time")

	// Cannot be in the future
	// Not greater than snapshot date
	Attribute("AccountStatusDate", String, func() {
		Description("Date of Status Change")
		MaxLength(8)
		Meta("rpc:tag", "49")
	})

	Required("AccountStatusDate")
})

// L– Client Deceased
var ClientDeceasedCompanyAccountStatus = Type("ClientDeceasedCompanyAccountStatus", func() {
	Description("Client Deceased")

	// Cannot be in the future
	// Not greater than snapshot date
	Attribute("AccountStatusDate", String, func() {
		Description("Date of Status Change")
		MaxLength(8)
		Meta("rpc:tag", "49")
	})

	Required("AccountStatusDate")

})

// M – Deferred – This refers to facilities whose payments have been put on hold
// for a definite period or in moratorium (Grace Period)
var DeferredCompanyAccountStatus = Type("DeferredCompanyAccountStatus", func() {
	Description("Deferred: Facilities whose payments put on hold for a definite period ")

	// Cannot be in the future
	// Not greater than snapshot date
	Attribute("AccountStatusDate", String, func() {
		Description("Date of Status Change")
		MaxLength(8)
		Meta("rpc:tag", "49")
	})

	// Mandatory if Account Status = Deferred.
	// Deferred Payment Date must be in the future.
	// If account status =M "Deferred" we expect Deferred Payment date
	Attribute("DeferredPaymentDate", String, func() {
		Description("Deferred Payment Date")
		Format(FormatDate)
		Meta("rpc:tag", "52")
	})

	// Conditional to field Account Status, Deferred Payment Date.
	// Conditional Fields to 4.1.49, 4.1.52
	Attribute("DeferredPaymentAmount", Int64, func() {
		Description("Deferred Payment Amount")
		MinLength(1)
		MaxLength(16)
		Meta("rpc:tag", "53")
	})

	Required("AccountStatusDate", "DeferredPaymentDate", "DeferredPaymentAmount")

})

// N – Not Updated – This status is reserved for CRBs (if last record status is not CLOSED)
var NotUpdatedCompanyAccountStatus = Type("NotUpdatedCompanyAccountStatus", func() {
	Description("Not Updated: reserved for CRBs (if last record status is not CLOSED)")

	// Cannot be in the future
	// Not greater than snapshot date
	Attribute("AccountStatusDate", String, func() {
		Description("Date of Status Change")
		MaxLength(8)
		Meta("rpc:tag", "49")
	})

	Required("AccountStatusDate")
})

// P – Disputed – Refers to a Record that the Client has disputed at the CRB.
var DisputedCompanyAccountStatus = Type("DisputedCompanyAccountStatus", func() {
	Description("Disputed: Client has disputed at the CRB")

	// Cannot be in the future
	// Not greater than snapshot date
	Attribute("AccountStatusDate", String, func() {
		Description("Date of Status Change")
		MaxLength(8)
		Meta("rpc:tag", "49")
	})

	Required("AccountStatusDate")
})
