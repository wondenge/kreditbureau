package kreditbureau

import (
	"github.com/wondenge/kreditbureau/values/accountindicator"
	"github.com/wondenge/kreditbureau/values/accountproductype"
	"github.com/wondenge/kreditbureau/values/accountstatus"
	"github.com/wondenge/kreditbureau/values/applicationdecline"
	"github.com/wondenge/kreditbureau/values/applicationstatus"
	"github.com/wondenge/kreditbureau/values/assetclassification"
	"github.com/wondenge/kreditbureau/values/bouncedcheques"
	"github.com/wondenge/kreditbureau/values/chequetype"
	"github.com/wondenge/kreditbureau/values/clientype"
	"github.com/wondenge/kreditbureau/values/collateralspart"
	"github.com/wondenge/kreditbureau/values/collateraltype"
	"github.com/wondenge/kreditbureau/values/consumerid"
	"github.com/wondenge/kreditbureau/values/functiontype"
	"github.com/wondenge/kreditbureau/values/employmentype"
	"github.com/wondenge/kreditbureau/values/gender13"
	"github.com/wondenge/kreditbureau/values/gender13b"
	"github.com/wondenge/kreditbureau/values/guaranteetype"
	"github.com/wondenge/kreditbureau/values/guarantorelationship"
	"github.com/wondenge/kreditbureau/values/income"
	"github.com/wondenge/kreditbureau/values/industrycode"
	"github.com/wondenge/kreditbureau/values/linkdelink"
	"github.com/wondenge/kreditbureau/values/maritalstatus"
	"github.com/wondenge/kreditbureau/values/memberstatus"
	"github.com/wondenge/kreditbureau/values/orgclass"
	"github.com/wondenge/kreditbureau/values/orgtype"
	"github.com/wondenge/kreditbureau/values/paymentfrequency"
	"github.com/wondenge/kreditbureau/values/residencytype"
	"github.com/wondenge/kreditbureau/values/securitytype"
	"github.com/wondenge/kreditbureau/values/sharedcollateral"
	"github.com/wondenge/kreditbureau/values/stakeholderid"
	"github.com/wondenge/kreditbureau/values/stakeholdertype"
	"github.com/wondenge/kreditbureau/values/statusupdate"
	"github.com/wondenge/kreditbureau/values/tradingstatus"
	"github.com/wondenge/kreditbureau/values/updatefunction"
)

type (
	LOV1 struct {
		OrganizationType orgtype.OrganizationType
	}
	LOV2 struct {
		IndustryCode industrycode.IndustryCode
	}
	LOV3 struct {
	}
	LOV4 struct {
		TradingStatus tradingstatus.TradingStatus
	}
	LOV5 struct {
		AccountIndicator accountindicator.AccountIndicator
	}
	LOV6 struct {
		OrganizationClass orgclass.OrganizationClass
	}
	LOV7 struct {
		AccountProductType accountproductype.AccountProductType
	}
	LOV8 struct {
		AssetClassification assetclassification.AssetClassification
	}
	LOV9 struct {
		AccountStatus accountstatus.AccountStatus
	}
	LOV10 struct {
		PaymentFrequency paymentfrequency.PaymentFrequency
	}
	LOV11 struct {
		SecurityType securitytype.SecurityType
	}
	LOV12 struct {
		ResidencyType residencytype.ResidencyType
	}
	LOV13 struct {
		Gender gender13.Gender
	}
	LOV13B struct {
		Gender gender13b.Gender
	}
	LOV14 struct {
		MaritalStatus maritalstatus.MaritalStatus
	}
	LOV15M struct {
		IDType consumerid.IDType
	}
	LOV15S struct {
		IDType stakeholderid.IDType
	}
	LOV16 struct {
		EmploymentType employmentype.EmploymentType
	}
	LOV17 struct {
		IncomeBandOptions income.IncomeBandOptions
	}
	LOV18 struct {
		StakeholderType stakeholdertype.StakeholderType
	}
	LOV19 struct {
		GuaranteeType guaranteetype.GuaranteeType
	}
	LOV20 struct {
		GuarantorRelationship guarantorelationship.GuarantorRelationship
	}
	LOV21 struct {
		ChequeType chequetype.ChequeType
	}
	LOV22 struct {
		CollateralType collateraltype.CollateralType
	}
	LOV23 struct {
		ClientType clientype.ClientType
	}
	LOV24 struct {
		ApplicationStatus applicationstatus.ApplicationStatus
	}
	LOV25 struct {
		ApplicationDecline applicationdecline.ApplicationDecline
	}
	LOV26 struct {
		StatusUpdate statusupdate.StatusUpdate
	}
	LOV27 struct {
		UpdateFunction updatefunction.UpdateFunction
	}
	LOV28 struct {
		BouncedChequeReason bouncedcheques.BouncedChequeReason
	}
	LOV29 struct {
		SharedCollateral sharedcollateral.SharedCollateral
	}
	LOV30 struct {
		CollateralsPart collateralspart.CollateralsPart
	}
	LOV31 struct {
		MemberStatus memberstatus.MemberStatus
	}
	LOV32 struct {
		DRFunctionType functiontype.DRFunctionType
	}
	LOV33 struct {
		LinkDelink linkdelink.LinkDelink
	}
)
