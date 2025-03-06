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

class Args:
    def __init__(self, argv):
        self.days = int(argv[1])
        self.atm = int(argv[2])
        self.max_tx = int(argv[3])

    def __str__(self):
        return str(dict(days = self.days, atm = self.atm, max_tx = self.max_tx))

def parse_args():
    if not len(sys.argv) >= 4:
        return None, err('invalid number of cmd args')
    if not sys.argv[1].isdigit():
        return None, err('invalid number of days')
    if not sys.argv[2].isdigit():
        return None, err('invalid number of ATM')
    if not sys.argv[3].isdigit():
        return None, err('invalid max number of transactions per ATM')
    
    args = Args(sys.argv)

    if args.days > MAX_DAYS:
        return None, err('max number of days is {}'.format(MAX_DAYS))
    if args.atm > MAX_ATM:
        return None, err('max number of ATM is {}'.format(MAX_ATM))
    if args.max_tx > MAX_TX:
        return None, err('max number of transactions per ATM is {}'.format(MAX_TX))
    
    if args.days == 0:
        args.days = random.randint(1, MAX_DAYS)
    if args.atm == 0:
        args.atm = random.randint(1, MAX_ATM)
    if args.max_tx == 0:
        args.max_tx = random.randint(1, MAX_TX)

    return args, ''

def err(msg):
    return 'err: {}'.format(msg)

def main():
    args, msg = parse_args()

    if not args:
        print(msg)
        return
    
    print(args)

main()
