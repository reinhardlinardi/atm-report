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

## Stack

- Golang 1.23
- Python 3.10
- MySQL 8.0

## Setup

### Dataset

1. Open terminal, navigate to project root
2. Create dataset folder and navigate to script directory
```bash
mkdir dataset && cd script
```
3. Run dataset generator
```bash
# cmd args: <# of days> <# of ATM> <max # of transactions per ATM>
python3 gen.py 2 4 0
```

### DB

1. Create new DB schema
```sql
CREATE SCHEMA `atm-report`;
```
2. Install [golang-migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate#installation)

3. Migrate DB tables
```bash
./migrate -path /path/to/migrations -database 'mysql://user:pass@tcp(localhost:3306)/atm-report' up
```
