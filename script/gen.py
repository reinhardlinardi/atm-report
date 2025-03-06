import sys
import random

# Cmd args: <number of days> <number of ATM> <max number of transactions per ATM>
#
# ATM data folder: atm-transactions/{id}
#
# ATM data filename: YYYYMMDD.{ext}
# ATM data file ext: .xml/csv/json/yaml
#
# Transaction data:
# - transactionId (hex alnum)
# - transactionDate (unix timestamp)
# - transactionType (deposit/withdraw/transfer) (integer)
# - amount (integer)
# - cardNumber (16-digit numeric)
# - destinationCardNumber (in case of transfer) (16-digit numeric)

MAX_DAYS = 7
MAX_ATM = 5
MAX_TX = 10

def valid_args():
    if not len(sys.argv) == 4:
        return False, 'Invalid number of cmd args'
    
    INVALID = 'Invalid args: '
    
    if not sys.argv[1].isdigit():
        return False, INVALID + 'number of days'
    
    if not sys.argv[2].isdigit():
        return False, INVALID + 'number of ATM'

    if not sys.argv[3].isdigit():
        return False, INVALID + 'max number of transactions per ATM'
    
    days = int(sys.argv[1])
    atm = int(sys.argv[2])
    max_tx = int(sys.argv[3])

    if days > MAX_DAYS:
        return False, INVALID + 'max ' + str(MAX_DAYS) + ' days'

    if atm > MAX_ATM:
        return False, INVALID + 'max ' + str(MAX_ATM) + ' ATM'
    
    if max_tx > MAX_TX:
        return False, INVALID + 'max ' + str(MAX_TX) + ' transactions per ATM'

    return True, ''

def get_args():
    days = int(sys.argv[1])
    atm = int(sys.argv[2])
    max_tx = int(sys.argv[3])

    if days == 0:
        days = random.randint(1, MAX_DAYS)

    if atm == 0:
        atm = random.randint(1, MAX_ATM)

    if max_tx == 0:
        max_tx = random.randint(1, MAX_TX)

    return {'days': days, 'atm': atm, 'max_tx': max_tx}

def main():
    valid, msg = valid_args()

    if not valid:
        print('FAIL:', msg)
        return
    
    print('OK')
    print(get_args())

main()
