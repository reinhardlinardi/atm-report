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

1. Open terminal
2. Navigate to script directory
3. Run dataset generator, e.g. `python3 gen.py 2 4 3 5`

### DB

1. Install [golang-migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate#installation)
2. Create new DB schema
```sql
CREATE SCHEMA `atm_report`;
```
3. Migrate tables
```bash
# Replace path, user, and pass in this command
./migrate -path migrations -database 'mysql://user:pass@tcp(localhost:3306)/atm_report' up
```

### App
Run `make build`

## Run

### App
Run `./atm-report`

### Producer
1. Open terminal
2. Navigate to script directory
3. Run `python3 send.py`
4. Follow instructions to send file to cron consumer
