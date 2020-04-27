package kreditbureau

import (
	"github.com/wondenge/kreditbureau/values/accountindicator"
	"github.com/wondenge/kreditbureau/values/accountproductype"
	"github.com/wondenge/kreditbureau/values/accountstatus"
	"github.com/wondenge/kreditbureau/values/assetclassification"
	"github.com/wondenge/kreditbureau/values/consumerid"
	"github.com/wondenge/kreditbureau/values/employmentype"
	"github.com/wondenge/kreditbureau/values/gender13"
	"github.com/wondenge/kreditbureau/values/gender13b"
	"github.com/wondenge/kreditbureau/values/guaranteetype"
	"github.com/wondenge/kreditbureau/values/guarantorelationship"
	"github.com/wondenge/kreditbureau/values/income"
	"github.com/wondenge/kreditbureau/values/industrycode"
	"github.com/wondenge/kreditbureau/values/maritalstatus"
	"github.com/wondenge/kreditbureau/values/orgclass"
	"github.com/wondenge/kreditbureau/values/orgtype"
	"github.com/wondenge/kreditbureau/values/paymentfrequency"
	"github.com/wondenge/kreditbureau/values/residencytype"
	"github.com/wondenge/kreditbureau/values/securitytype"
	"github.com/wondenge/kreditbureau/values/stakeholderid"
	"github.com/wondenge/kreditbureau/values/stakeholdertype"
	"github.com/wondenge/kreditbureau/values/tradingstatus"
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
	}
	LOV22 struct {
	}
	LOV23 struct {
	}
	LOV24 struct {
	}
	LOV25 struct {
	}
	LOV26 struct {
	}
	LOV27 struct {
	}
	LOV28 struct {
	}
	LOV29 struct {
	}
	LOV30 struct {
	}
	LOV31 struct {
	}
	LOV32 struct {
	}
	LOV33 struct {
	}
)
