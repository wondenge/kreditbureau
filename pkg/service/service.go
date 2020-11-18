package service

import (
	"context"
	//"github.com/go-pg/pg/v10"
	//"github.com/go-redis/redis/v8"
	//"github.com/hashicorp/vault/api"
	. "github.com/wondenge/kreditbureau/internal"
	//"go.mongodb.org/mongo-driver/mongo"
	//"go.mongodb.org/mongo-driver/mongo/options"
)

// KreditbureauService describes the service.
type KreditbureauService interface {

	// Daily Payment Information File
	PostRepayment(ctx context.Context, p *Repayments) error
	GetRepayment(ctx context.Context, id string) (*Repayments, error)
	PutRepayment(ctx context.Context, id string, p *Repayments) error
	PatchRepayment(ctx context.Context, id string, p *Repayments) error
	DeleteRepayment(ctx context.Context, id string) error

	// Mobile Facilities File
	PostMobileFacility(ctx context.Context, p *MobileFacility) error
	GetMobileFacility(ctx context.Context, id string) (*MobileFacility, error)
	PutMobileFacility(ctx context.Context, id string, p *MobileFacility) error
	PatchMobileFacility(ctx context.Context, id string, p *MobileFacility) error
	DeleteMobileFacility(ctx context.Context, id string) error

	// Historical Credit Information Update File
	PostCreditHistory(ctx context.Context, p *CreditHistory) error
	GetCreditHistory(ctx context.Context, id string) (*CreditHistory, error)
	PutCreditHistory(ctx context.Context, id string, p *CreditHistory) error
	PatchCreditHistory(ctx context.Context, id string, p *CreditHistory) error
	DeleteCreditHistory(ctx context.Context, id string) error

	// Contact Upload File
	PostContact(ctx context.Context, p *Contact) error
	GetContact(ctx context.Context, id string) (*Contact, error)
	PutContact(ctx context.Context, id string, p *Contact) error
	PatchContact(ctx context.Context, id string, p *Contact) error
	DeleteContact(ctx context.Context, id string) error

	// Delink IDs from an Account File:  ID to Link from Account
	PostIDToLink(ctx context.Context, p *IDToLink) error
	GetIDToLink(ctx context.Context, id string) (*IDToLink, error)
	PutIDToLink(ctx context.Context, id string, p *IDToLink) error
	PatchIDToLink(ctx context.Context, id string, p *IDToLink) error
	DeleteIDToLink(ctx context.Context, id string) error

	// Delink IDs from an Account File: ID to Delink from Account
	PostIDToDelink(ctx context.Context, p *IDToDelink) error
	GetIDToDelink(ctx context.Context, id string) (*IDToDelink, error)
	PutIDToDelink(ctx context.Context, id string, p *IDToDelink) error
	PatchIDToDelink(ctx context.Context, id string, p *IDToDelink) error
	DeleteIDToDelink(ctx context.Context, id string) error

	// Link-Delink IDs File
	PostLinkID(ctx context.Context, p *LinkDelinkID) error
	GetLinkID(ctx context.Context, id string) (*LinkDelinkID, error)
	PutLinkID(ctx context.Context, id string, p *LinkDelinkID) error
	PatchLinkID(ctx context.Context, id string, p *LinkDelinkID) error
	DeleteLinkID(ctx context.Context, id string) error

	// Accounts Merger File
	PostMerger(ctx context.Context, p *Merger) error
	GetMerger(ctx context.Context, id string) (*Merger, error)
	PutMerger(ctx context.Context, id string, p *Merger) error
	PatchMerger(ctx context.Context, id string, p *Merger) error
	DeleteMerger(ctx context.Context, id string) error

	// Deletion/Relist File: Delete History
	PostHistoryPurge(ctx context.Context, p *HistoryPurge) error
	GetHistoryPurge(ctx context.Context, id string) (*HistoryPurge, error)
	PutHistoryPurge(ctx context.Context, id string, p *HistoryPurge) error
	PatchHistoryPurge(ctx context.Context, id string, p *HistoryPurge) error
	DeleteHistoryPurge(ctx context.Context, id string) error

	// Deletion/Relist File: Relist History
	PostHistoryRelist(ctx context.Context, p *HistoryRelist) error
	GetHistoryRelist(ctx context.Context, id string) (*HistoryRelist, error)
	PutHistoryRelist(ctx context.Context, id string, p *HistoryRelist) error
	PatchHistoryRelist(ctx context.Context, id string, p *HistoryRelist) error
	DeleteHistoryRelist(ctx context.Context, id string) error

	// Non-Individual Consumer and Account File
	PostCompany(ctx context.Context, p *Company) error
	GetCompany(ctx context.Context, id string) (*Company, error)
	PutCompany(ctx context.Context, id string, p *Company) error
	PatchCompany(ctx context.Context, id string, p *Company) error
	DeleteCompany(ctx context.Context, id string) error

	// Individual Consumer, Employer & Account File
	PostIndividual(ctx context.Context, p *Individual) error
	GetIndividual(ctx context.Context, id string) (*Individual, error)
	PutIndividual(ctx context.Context, id string, p *Individual) error
	PatchIndividual(ctx context.Context, id string, p *Individual) error
	DeleteIndividual(ctx context.Context, id string) error

	// Stakeholder File
	PostStakeholder(ctx context.Context, p *Stakeholder) error
	GetStakeholder(ctx context.Context, id string) (*Stakeholder, error)
	PutStakeholder(ctx context.Context, id string, p *Stakeholder) error
	PatchStakeholder(ctx context.Context, id string, p *Stakeholder) error
	DeleteStakeholder(ctx context.Context, id string) error

	// Guarantor File
	PostGuarantor(ctx context.Context, p *NewGuarantor) error
	GetGuarantor(ctx context.Context, id string) (*NewGuarantor, error)
	PutGuarantor(ctx context.Context, id string, p *NewGuarantor) error
	PatchGuarantor(ctx context.Context, id string, p *NewGuarantor) error
	DeleteGuarantor(ctx context.Context, id string) error

	// Bounced Cheque File
	PostCheque(ctx context.Context, p *Cheque) error
	GetCheque(ctx context.Context, id string) (*Cheque, error)
	PutCheque(ctx context.Context, id string, p *Cheque) error
	PatchCheque(ctx context.Context, id string, p *Cheque) error
	DeleteCheque(ctx context.Context, id string) error

	// Collateral Register File
	PostRegister(ctx context.Context, p *Register) error
	GetRegister(ctx context.Context, id string) (*Register, error)
	PutRegister(ctx context.Context, id string, p *Register) error
	PatchRegister(ctx context.Context, id string, p *Register) error
	DeleteRegister(ctx context.Context, id string) error

	// Fraudulent Activities File
	PostFraud(ctx context.Context, p *Fraud) error
	GetFraud(ctx context.Context, id string) (*Fraud, error)
	PutFraud(ctx context.Context, id string, p *Fraud) error
	PatchFraud(ctx context.Context, id string, p *Fraud) error
	DeleteFraud(ctx context.Context, id string) error

	// Credit Application File
	PostFacility(ctx context.Context, p *Facility) error
	GetFacility(ctx context.Context, id string) (*Facility, error)
	PutFacility(ctx context.Context, id string, p *Facility) error
	PatchFacility(ctx context.Context, id string, p *Facility) error
	DeleteFacility(ctx context.Context, id string) error

	// Group Guarantee File
	PostGuarantee(ctx context.Context, p *Guarantee) error
	GetGuarantee(ctx context.Context, id string) (*Guarantee, error)
	PutGuarantee(ctx context.Context, id string, p *Guarantee) error
	PatchGuarantee(ctx context.Context, id string, p *Guarantee) error
	DeleteGuarantee(ctx context.Context, id string) error
}
