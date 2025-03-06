# ATM Report Service

## Description
ATM Report Service is designed to process transaction data from different ATMs (A, B, C, D, .. etc), which generate transactions in different formats (XML, CSV, JSON, YAML) on daily basis. The service fetches those files & generates reports based on the transaction data and exposes them as REST APIs.

Features:  

Supports multiple transaction file formats: XML, CSV, JSON, YAML.  
All files have the same format:

- transactionId
- transactionDate
- cardNumber
- transactionType (deposit/withdraw/transfer)
- amount
- destinationCardNumber (in case of transfer)

Expected to continuously fetching files of those ATMs & load into DB.  
Reports to expose via APIs:

- Number of transactions per day
- Number of transactions per transaction type per day
- Number of transactions per day and per transaction type per day
- ATM with max withdraw per day
