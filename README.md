# Kreditbureau - API driven Credit Reporting.

[Kreditbureau](https://github.com/wondenge/kreditbureau) adheres to Version 4 of the Data Specification Template Validation that was released by the [Central Bank of Kenya](https://www.centralbank.go.ke/), following consultations among commercial banks, the [Kenya Deposit Insurance Corporation (KDIC)](https://kdic.go.ke), Microfinance Banks and the Credit Reference Bureaus; for the implementation of full-file credit information sharing.

# 1. Monthly Batch APIs

## 1.1 NonIndividualAccount API

This API is mandatory for credit providers who have advanced a credit facility to non-individual consumers.

## 1.2 IndividualAccount API

This API is mandatory for credit providers who have advanced credit facilities to individual consumers.

## 1.3 Stakeholders API

This API is mandatory if non-individual accounts exist.

## 1.4 Guarantors API

This API ptovides guarantor information for a loan. This includes group guarantors.

## 1.5 BouncedCheques API

This API provides information about bounced cheques.

## 1.6 Collaterals API

This API updates a collaterals (items, funds and/or property pledged to secure a loan or a debt.)

## 1.7 Fraud API

This API submits fraudlent information that has been proven in a court of law.

## 1.8 CreditApplications API

This API submits new loan applications.

## 1.9 GroupGuarantee API

This API lists members of a group in a group-lending methodology.

# 2. Daily Batch APIs

## 2.1 DailyPayment API

This API reports any payment(s) made into a credit facility on a given date.

## 2.2 MobileFacility API

This API reports new mobile loans granted by a lender.

## 2.3 DelinkID API

This API delinks an ID from a record that was previously submitted to the Bureau.

## 2.4 LinkDelinkIDs API

This API links IDs e.g. ID to a Passport or Service ID or Alien IDs.

## 2.5 AccountsMerger API

The API is to merge accounts.This could be as a result of system change.

## 2.6 DeletionRelist API

This API deletes default history or relist deleted historical data of a customer.

## 2.7 HistoricalCredit API

This API updates the credit information of a facility which was erroneously submitted.

## 2.8 ContactUpload API

This API updates contact information on a given borrowing client.
