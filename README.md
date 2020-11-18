# Kreditbureau - Credit Reporting for Financial Inclusivity.

[Kreditbureau](https://github.com/wondenge/kreditbureau) is a financial tool for CRBSs to help in managing of credit reports from across the Financial Sector in Kenya. Kreditbureau adheres to v4 of the DST Validation that was released by the [Central Bank of Kenya](https://www.centralbank.go.ke/), following consultations among commercial banks, the [KDIC](https://kdic.go.ke), Microfinance Banks and the Credit Reference Bureaus; for the implementation of full-file credit information sharing.

# 1. Manthle - Monthly Batch APIs

## Non-Individual Consumer Account API

- API is mandatory for credit providers who have advanced a credit facility to non-individual consumers. These are registered institutions such as companies, partnerships, societies, church, school, etc.
- An account is considered non-individual, if the facility or loan contract is given to the registered entity and not its owner.
- In the case of the sole proprietor or the micro-entrepreneur, the account is considered non-individual if the loan contract is written and granted to the business name (as opposed to the owner’s name).
- Also, as a rule of thumb, credit providers lending for a business are encouraged to record the loan under the business name (not the owner) if the business is a registered entity.
- This API also includes the account information for the entity. This is the core of credit API as it contains information on the key behaviour aspects of the borrower, specifically data on their payment patterns and level of exposure. It also provides information on the longevity of the account.
- If a record is listed in this API, it is expected that there will be an associated entry in the stakeholder API to detail the significant stakeholders for the loan.

## Individual Consumer, Employment, and Account API

- API is mandatory for credit providers who have advanced credit facilities to individual consumers.

- Data under the rubric of Individual Consumer information is the central source of data for identity for all loans made to individuals.

Also included, is the Consumer’s corresponding account information is included and shows the actual performance of their credit facility

## Stakeholders API

- API is mandatory if non-individual accounts exits. Records in this file pertain to individuals who;

  1. Hold at least 10% shareholding in a non-individual consumer
  2. Are key stakeholders of the non-individual consumer( Directors,Partners, Trustees or Officials )

- Stakeholder information can assist risk assessment for non-individual borrowers, but is more like to have greater value in establishing identity. The data can assist in due diligence check for corporate borrowers.

## Guarantor API

- Guarantor information is mandatory if it exists for a loan. This includes group guarantors.

- A guarantor is defined as a person or entity that agrees to be responsible for another’s debt or performance under a contract, if the borrower fails to pay or perform. The records submitted in this file refer to persons or entities that meet the definition.

- Guarantor information assists in providing information on the extent to which a borrower is exposed in their guarantees on the loans of others. This data also assists in assessing the risk of loans of an individual as guarantors mitigate
  the risk of default.

- As a rule of thumb, individuals or entities who signed as a guarantor on a loan contract are considered guarantors and should be included in this data file. However Credit Providers can use their own criteria to decide subject to the condition that the entities listed as guarantors are actually liable to pay for a facility that falls in arrears.

## Bounced Cheque API

- API is non-mandatory and is only if the lender has this information available.

- Information in this API should be submitted by the payee bank that bounced the cheque. The data submitted refers to cheques bounced due to insufficient funds or linked to fraud. As per the CIS Regulations, cheques to be submitted are
  only those meant for settlement of credits in favour of institutions.

- Bounced cheque data is often a valuable indicator of credit risk and could be useful as part of the assessment of the Consumer.

## Collateral Register API

- Records are mandatory in this API, if in the Individual and/or Non Individual APIs have the field “Type of Security” updated as "Fully Secured”

- Collaterals are items, funds and/or property pledged to secure a loan or a debt.

- Knowing whether a loan is secured by collaterals, and the nature of those collaterals may be valuable information during the course of loan analysis. For the purpose of this submission, only items that are pledged as collaterals on signed loan agreements are considered collaterals.

## Fraudulent Activities API

- Non-mandatory API but only shared if fraud is proven by the courts.

- Information is submitted in this API when a fraud occurs and has been proven in a court of law.

- Data on fraudulent activities is very useful in the issuance of credit, as fraud is a special category of risk. As per the Regulations, this API is to be submitted only where fraud has been proven in a court of law.

## Credit Application API

- New loan applications made should be submitted in this API.

- Credit application data is important as indicators whether a borrower is suddenly or excessively shopping for credit is an indicator of risk and impending exposure. In line with international practice, this information will in the near future be provided real time in order to be of greater value to lenders who receive applications from serial applicants.

## Group Guarantee API

- API lists members of a group in a group-lending methodology.

- All members of a group are unified by the Group ID. A member of the group who obtains a facility under the group-lending methodology will have a record in the Individual.

- Some group methodologies have sub-groups, where members of a sub-group directly guarantee a loan of any of the members in that sub-group. In this case, the lender must assign a sub-group ID to each sub-group within a group.

- If a sub-group exists within a group (thus a Sub-Group ID has been filled), only members listed with the same sub-group will be considered as guarantors of the loan. If no sub-group exists, all members of the group are considered as direct guarantors of the loan.

# 2. Dale - Daily Batch APIs

## Daily Payment API

- API is used for accounts which have received a payment on a given day.

- The API reports any payment(s) made into a credit facility on a given date. If more than one payment is received in a loan on a given day, the lender will sum the payments.

- If a payment to an account was done erroneously (and subsequently updated at the Bureau), this API can be used to update the Bureau, with a negative payment value.

- This API is the standard template to be used when a lender wishes to make an adhoc update to the Bureau (e.g. when a client request their current loan status to be updated to the Bureau).
- Lenders who wish to update the Bureau on a daily basis will be required to use this API to update payments (or reversal of payments) into accounts on a given day.

- The updated status should be reflected in the subsequent monthly batch submission.

## Mobile Facilities API

- API is used by lenders with Mobile lending facilities.This helps to report new mobile loans granted by the lender. The API has been designed to be brief to ease the process of a lender to update the information at the Bureau; and thus increase their ability to increase the frequency of updating information at the Bureau.

## Delink ID from Account API

- API is used to Delink an ID from an account. This is helpful for lenders to delink an ID from a record that was previously submitted to the Bureau. This could be as a result of a dispute or shared IDs.

## Link-Delink IDs API

API is used to Link and Delink IDs from an account:

- Links IDs e.g. ID to a Passport or Service ID or Alien IDs.
- Delinks a wrong ID and link the right ID to an account. This could be as a result of a dispute or if the lender realises that they submitted erroneous information submitted.

## Accounts Merger API

- API is used to Merge Accounts. This is helpful for lenders to merge accounts as a result of system change.

## Deletion/Relist API

- API used whenever the default history is to be deleted. This API is used by lenders to delete Default History/ Relist Deleted historical data of a customer.
- The API will be used in case there is a court order or an agreement of the 2 parties i.e. Customer and Lender to delete default history.

## Historical Credit Information Update API

- API used by a lender wishing to correct an erroneous state of an account previously submitted to the Bureau. This updates the credit information of a facility which was erroneously submitted to the Bureau.
- This could also be used to resolve a dispute, if the lender realizes it submitted the wrong credit information to the Bureau.

## Contact Upload API

- API used by a lender who wishes to update a Bureau with contact information on a given borrowing client.
