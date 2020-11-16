# Kreditbureau - API driven Credit Reporting.

[Kreditbureau](https://github.com/wondenge/kreditbureau) adheres to Version 4 of the Data Specification Template Validation that was released by the [Central Bank of Kenya](https://www.centralbank.go.ke/), following consultations among commercial banks, the [Kenya Deposit Insurance Corporation (KDIC)](https://kdic.go.ke), Microfinance Banks and the Credit Reference Bureaus; for the implementation of full-file credit information sharing.

# 1. Monthly Batch APIs

## 1.1 Non-Individual Consumer Account API

This file is mandatory for credit providers who have advanced a credit facility to non-individual consumers

A non-individual consumer refers to registered institutions such as companies, partnerships, societies, church, school, etc. An account is considered non-individual, if the facility or loan contract is given to the registered entity and not its owner.

In the case of the sole proprietor or the micro-entrepreneur, the account is considered non-individual if the loan contract is written and granted to the business name (as opposed to the owner’s name). Also, as a rule of thumb, credit providers lending for a business are encouraged to record the loan under the business name (not the owner) if the business is a registered entity.

This file is central for establishing the identity of non-individual borrowers. The fields cover:

- registered name,
- trading name,
- registration number,
- company type,
- industry code,
- PIN number,
- VAT number
- and other fields serve to establish an institutional borrower’s identity.

This file also includes the account information for the entity. This is the core of credit file as it contains information on the key behaviour aspects of the borrower, specifically data on their payment patterns and level of exposure. It also provides information on the longevity of the account.

If a record is listed in this file, it is expected that there will be an associated entry in the stakeholder file to detail the significant stakeholders for the loan.

## 1.2 Individual Consumer, Employment, and Account API

This is mandatory for credit providers who have advanced credit facilities to individual consumers.

Data under the rubric of Individual Consumer information is the central source of data for identity for all loans made to individuals. The fields cover name, identifying numbers, addresses, and other contact information.

Where applicable, this file also contains the individual’s employment data. Employment data is valuable both for matching (or establishing identity) and for risk assessment especially as it can speak to capacity. Type of employment and gross monthly income assist in establishing the ability of a borrower to afford a loan. Employer data helps match a client to an account, though again mismatches on employment data will not necessarily indicate different individuals if Consumer change jobs regularly.

In this same file, the Consumer’s corresponding account information is included and shows the actual performance of their credit facility

## 1.3 Stakeholders API

This file is mandatory if non-individual accounts exits.

Records in this file pertain to individuals who;

1. Hold at least 10% shareholding in a non-individual consumer
2. Are key stakeholders of the non-individual consumer

- Directors
- Partners
- Trustees or Officials

Stakeholder information can assist risk assessment for non-individual borrowers, but is more like to have greater value in establishing identity. The data can assist in due diligence check for corporate borrowers.

## 1.4 Guarantor API

Guarantor information is mandatory if it exists for a loan. This includes group guarantors.

A guarantor is defined as a person or entity that agrees to be responsible for another’s debt or performance under a contract, if the borrower fails to pay or perform. The records submitted in this file refer to persons or entities that meet the definition.

Guarantor information assists in providing information on the extent to which a borrower is exposed in their guarantees on the loans of others. This data also assists in assessing the risk of loans of an individual as guarantors mitigate
the risk of default.

As a rule of thumb, individuals or entities who signed as a guarantor on a loan contract are considered guarantors and should be included in this data file. However Credit Providers can use their own criteria to decide subject to the condition that the entities listed as guarantors are actually liable to pay for a facility that falls in arrears.

## 1.5 Bounced Cheque API

This file is non-mandatory and is only if the lender has this information available.

Information in this file should be submitted by the payee bank that bounced the cheque. The data submitted refers to cheques bounced due to insufficient funds or linked to fraud. As per the CIS Regulations, cheques to be submitted are
only those meant for settlement of credits in favour of institutions.

Bounced cheque data is often a valuable indicator of credit risk and could be useful as part of the assessment of the Consumer.

## 1.6 Collateral Register API

Records are mandatory in this file, if in the Individual and/or Non Individual files have the field “Type of Security” updated as "Fully Secured”

Collaterals are items, funds and/or property pledged to secure a loan or a debt.

Knowing whether a loan is secured by collaterals, and the nature of those collaterals may be valuable information during the course of loan analysis. For the purpose of this submission, only items that are pledged as collaterals on signed
loan agreements are considered collaterals.

## 1.7 Fraudulent Activities API

Non-mandatory file but only shared if fraud is proven by the courts

Information is submitted in this file when a fraud occurs and has been proven in a court of law.

Data on fraudulent activities is very useful in the issuance of credit, as fraud is a special category of risk. As per the Regulations, this file is to be submitted only where fraud has been proven in a court of law.

## 1.8 Credit Application API

New loan applications made should be submitted in this file.

Credit application data is important as indicators whether a borrower is suddenly or excessively shopping for credit is an indicator of risk and impending exposure. In line with international practice, this information will in the near future be provided real time in order to be of greater value to lenders who receive applications from serial applicants.

## 1.9 Group Guarantee API

This file lists members of a group in a group-lending methodology.

All members of a group are unified by the Group ID. A member of the group who obtains a facility under the group-lending methodology will have a record in the Individual
File with Field 4.2.60 filled in with the Group ID of the group he/she belongs to. The borrower should also have an entry in the Group Guarantee File.

Some group methodologies have sub-groups, where members of a sub-group directly guarantee a loan of any of the members in that sub-group. In this case, the lender must assign a sub-group ID to each sub-group within a group.

If a sub-group exists within a group (thus a Sub-Group ID has been filled), only members listed
with the same sub-group will be considered as guarantors of the loan. If no sub-group exists, all
members of the group are considered as direct guarantors of the loan.

# 2. Daily Batch APIs

## 2.1 Daily Payment API

This file should be filled for accounts which have received a payment on a given day.

This file will report any payment(s) made into a credit facility on a given date. If more than one payment is received in a loan on a given day, the lender will sum the payments as the value of field 5.1.11. Other fields should be populated with the new status of the account after the payments. If a payment to an account was done erroneously (and subsequently updated at the Bureau), this file can be used to update the Bureau, with a negative payment value.

This file is the standard template to be used when a lender wishes to make an adhoc update to the Bureau (e.g. when a client request their current loan status to be updated to the Bureau). Lenders who wish to update the Bureau on a daily basis will be required to use this file to update payments (or reversal of payments) into accounts on a given day.

The updated status should be reflected in the subsequent monthly batch submission.

## 2.2 Mobile Facilities API

This file should be filled in by lenders with Mobile lending facilities.

This file will report new mobile loans granted by a lender. The file has been designed to be brief to ease the process of a lender to update the information at the Bureau; and thus increase their ability to increase the frequency of updating information at the Bureau.

```bash
goa gen github.com/wondenge/kreditbureau/mobile-facilities/design
goa example github.com/wondenge/kreditbureau/mobile-facilities/design
go build ./cmd/mobilefacilities && go build ./cmd/mobilefacilities-cli
```

## 2.3 Delink ID from Account API

This File should be used to Delink an ID from an account.

This File will be used by lenders to delink an ID from a record that was previously submitted to the Bureau. This could be as a result of a dispute or shared IDs.

## 2.4 Link-Delink IDs API

The file should be used to Link and Delink IDs from an account.

This file will be used to Link IDs e.g. ID to a Passport or Service ID or Alien IDs. The file will also be used to Delink a wrong ID and link the right ID to an account. This could be as a result of a dispute or if the lender realises that they submitted erroneous information submitted.

## 2.5 Accounts Merger API

This file will be used to Merge Accounts.

The File will be used by lenders to merge accounts. This is could be as a result of system change.

## 2.6 Deletion/Relist API

This file should be used whenever the default history is to be deleted.

This file will be used by lenders to delete Default History/ Relist Deleted historical data of a customer. The file will be used in case there is a court order or an agreement of the 2 parties i.e. Customer and Lender to delete default history.

## 2.7 Historical Credit Information Update API

This file should be submitted by a lender wishing to correct an erroneous state of an account previously submitted to the Bureau.

This File will be used to update the credit information of a facility which was erroneously submitted to the Bureau. For example, if a lender submitted the wrong current balance amount (either due to a typographical or other human error; and thus does not require payment information to update the current balance amount) and later realize the error, the lender will
use this file to update the information. This could also be used to resolve a dispute, if the lender realizes it submitted the wrong credit information to the Bureau.

## 2.8 Contact Upload API

This file should be used by a lender wishing to add contact information for a given profile.

This file will be used by a lender who wishes to update a Bureau with contact information on a given borrowing client. The Bureau will add the contact information to existing contacts of a profile.
