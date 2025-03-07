# Cmd args: <number of days> <number of ATM> <max number of transactions per ATM>
#
# ATM data directory: atm-transactions
# ATM data filename: {id}_YYYYMMDD.{ext}
# ATM data file ext: .csv/json/yaml/xml
#
# Transaction data:
# - transactionId (hex string)
# - transactionDate (datetime)
# - transactionType (integer)
# - amount (integer)
# - cardNumber (16-digit number)
# - destinationCardNumber (16-digit number)

import sys
import random
import time
import hashlib
import os
import csv
import json
from datetime import datetime, timedelta

CSV = "csv"
JSON = "json"
YAML = "yaml"
XML = "xml"

FILENAME_FORMAT = '%Y%m%d'
EXT = [CSV, JSON, YAML, XML]

MAX_DAYS = 7
MAX_ATM = 5
MAX_TX = 10

TYPE_WITHDRAW = 0
TYPE_DEPOSIT = 1
TYPE_TRANSFER = 2

TYPES = [TYPE_WITHDRAW, TYPE_DEPOSIT, TYPE_TRANSFER]

KEY_ID = 'transactionId'
KEY_DATE = 'transactionDate'
KEY_TYPE = 'transactionType'
KEY_AMOUNT = 'amount'
KEY_SRC = 'cardNumber'
KEY_DEST = 'destinationCardNumber'

DATE_FORMAT = "%H:%M:%S"
CSV_HEADER = [KEY_ID, KEY_DATE, KEY_TYPE, KEY_AMOUNT, KEY_SRC, KEY_DEST]

class Args:
    def __init__(self, argv):
        self.days = int(argv[1])
        self.atm = int(argv[2])
        self.max_tx = int(argv[3])

class Tx:
    def __init__(self):
        self.id = ''
        self.date = None
        self.type = 0
        self.amount = 0
        self.src = ''
        self.dest = ''
        self.unix = 0

    def gen(self, date):
        self.date = randtime(date)
        self.unix = time.mktime(self.date.timetuple())
        self.id = hashlib.md5(str(self.unix).encode()).digest().hex()[:12]

        self.type = TYPES[random.randrange(0, len(TYPES))]
        self.amount = randdigits(4)
        self.src = randcardnum()

        if self.type == TYPE_TRANSFER:
            self.dest = randcardnum()

    def list(self):
        return [self.id, Tx._strftime(self.date) , self.type, self.amount, self.src, self.dest]
    
    def dict(self):
        return {KEY_ID: self.id, KEY_DATE: Tx._strftime(self.date), KEY_TYPE: self.type, KEY_AMOUNT: self.amount, KEY_SRC: self.src, KEY_DEST: self.dest}

    def _strftime(date):
        return date.strftime(DATE_FORMAT)

def randtime(date):
    hour = random.randrange(0, 24)
    min = random.randrange(0, 60)
    sec = random.randrange(0, 60)

    return date + timedelta(hours = hour, minutes = min, seconds = sec)

def randcardnum():
    digits = 8
    return str(randdigits(digits)).zfill(digits) + str(randdigits(digits)).zfill(digits)

def randdigits(digits):
    end = 1

    while digits:
        end *= 10
        digits -= 1

    return random.randrange(1, end)

def generate(cnt, date):
    gen = []

    while cnt:
        tx = Tx()
        tx.gen(date)
        gen.append(tx)
        cnt -= 1

    gen.sort(key = lambda tx : tx.unix)
    return gen

def parseargs():
    if not len(sys.argv) == 4:
        return None, err('invalid number of cmd args, expect 3')
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

def cleardir(dir):
    try:
        files = os.listdir(dir)

        for file in files:
            path = os.path.join(dir, file) 
            os.remove(path)

    except OSError:
        return False
    
    return True

def write_csv(gen, f):
    writer = csv.writer(f)
    writer.writerow(CSV_HEADER)

    for row in gen:
        writer.writerow(row.list())

def write_json(gen, f):
    kv = []

    for tx in gen:
        kv.append(tx.dict())
    
    json.dump(kv, f)

def err(msg):
    return 'err: {}'.format(msg)

def main():
    args, msg = parseargs()

    if not args:
        print(msg)
        return
    
    today = datetime.today().replace(hour = 0, minute = 0, second = 0)
    dir = '../atm-transactions'

    if not cleardir(dir):
        print(err('failed to delete files in {}'.format(dir)))
        return

    for d in range(1, args.days+1):
        date = today - timedelta(days = d)

        for a in range(0, args.atm):
            id = chr(ord('A') + a)
            ext = EXT[a % len(EXT)]
            name = '{}_{}.{}'.format(id, date.strftime(FILENAME_FORMAT), ext)

            num_tx = random.randint(0, args.max_tx)
            gen = generate(num_tx, date)
            
            with open(os.path.join(dir, name), 'w') as f:
                if ext == CSV:
                    write_csv(gen, f)
                if ext == JSON:
                    write_json(gen, f)
                
main()
