# ATM report service dataset generator
#
# cmd args: <# of days> <# of ATM> <max files per ATM> <max transactions per ATM>
#
# ATM data directory: gen
# ATM data filename: {id}_YYYYMMDD_{seq}.{ext}
# ATM data file ext: .csv/json/yaml/xml
#
# Transaction data:
# - transactionId (hex string)
# - transactionDate (YYYY-MM-DD)
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
import yaml
import xml.etree.ElementTree as ET
from datetime import datetime, timedelta

CSV = "csv"
JSON = "json"
YAML = "yaml"
XML = "xml"

FILENAME_FORMAT = '%Y%m%d'
EXT = [CSV, JSON, YAML, XML]

MAX_DAYS = 7
MAX_ATM = 8
MAX_FILES = 5
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

DATE_FORMAT = "%Y-%m-%d"
XML_ROOT = 'transactions'
XML_TAG = 'transaction'
CSV_HEADER = [KEY_ID, KEY_DATE, KEY_TYPE, KEY_AMOUNT, KEY_SRC, KEY_DEST]

class Args:
    def __init__(self, argv):
        self.days = int(argv[1])
        self.atm = int(argv[2])
        self.max_files = int(argv[3])
        self.max_tx = int(argv[4])

class Tx:
    def __init__(self):
        self.id = ''
        self.date = None
        self.type = 0
        self.amount = 0
        self.src = ''
        self.dest = ''

    def gen(self, date):
        self.date = date
        
        t = time.mktime(randtime(self.date).timetuple())
        self.id = hashlib.md5(str(t).encode()).digest().hex()[:12]

        self.type = TYPES[random.randrange(0, len(TYPES))]
        self.amount = randdigits(4)
        self.src = randcardnum()

        if self.type == TYPE_TRANSFER:
            self.dest = randcardnum()

    def list(self):
        return [self.id, self.date.strftime(DATE_FORMAT), self.type, self.amount, self.src, self.dest]
    
    def dict(self):
        return {KEY_ID: self.id, KEY_DATE: self.date.strftime(DATE_FORMAT), KEY_TYPE: self.type, KEY_AMOUNT: self.amount, KEY_SRC: self.src, KEY_DEST: self.dest}

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

    return gen

def parseargs():
    num_args = 4

    if not len(sys.argv) == num_args+1:
        return None, err('invalid number of cmd args, expect {}'.format(num_args))
    if not sys.argv[1].isdigit():
        return None, err('invalid number of days')
    if not sys.argv[2].isdigit():
        return None, err('invalid number of ATM')
    if not sys.argv[3].isdigit():
        return None, err('invalid max files per ATM')
    if not sys.argv[4].isdigit():
        return None, err('invalid max transactions per ATM')
    
    args = Args(sys.argv)

    if args.days > MAX_DAYS:
        return None, err('max number of days is {}'.format(MAX_DAYS))
    if args.atm > MAX_ATM:
        return None, err('max number of ATM is {}'.format(MAX_ATM))
    if args.atm > MAX_FILES:
        return None, err('max files per ATM is {}'.format(MAX_FILES))
    if args.max_tx > MAX_TX:
        return None, err('max transactions per ATM is {}'.format(MAX_TX))
    
    if args.days == 0:
        args.days = random.randint(1, MAX_DAYS)
    if args.atm == 0:
        args.atm = random.randint(1, MAX_ATM)
    if args.max_files == 0:
        args.max_files = random.randint(1, MAX_FILES)
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

def write_csv(data, f):
    writer = csv.writer(f)
    writer.writerow(CSV_HEADER)

    for row in data:
        writer.writerow(row.list())

def write_json(data, f):
    d = []

    for tx in data:
        d.append(tx.dict())
    
    json.dump(d, f)

def write_yaml(data, f):
    d = []

    for tx in data:
        d.append(tx.dict())

    yaml.dump(d, f)

def write_xml(data, f):
    root = ET.Element(XML_ROOT)

    for tx in data:
        tag = ET.SubElement(root, XML_TAG)
        tag.set(KEY_ID, tx.id)

        date = ET.SubElement(tag, KEY_DATE)
        type = ET.SubElement(tag, KEY_TYPE)
        amount = ET.SubElement(tag, KEY_AMOUNT)
        src = ET.SubElement(tag, KEY_SRC)
        dest = ET.SubElement(tag, KEY_DEST)
        
        date.text = tx.date.strftime(DATE_FORMAT)
        type.text = str(tx.type)
        amount.text = str(tx.amount)
        src.text = tx.src
        dest.text = tx.dest

    s = ET.tostring(root)
    f.write(s)

def err(msg):
    return 'err: {}'.format(msg)

def main():
    args, msg = parseargs()

    if not args:
        print(msg)
        return
    
    today = datetime.today().replace(hour = 0, minute = 0, second = 0)
    dir = '../gen'

    if not cleardir(dir):
        print(err('failed to delete files in {}'.format(dir)))
        return

    for d in range(1, args.days+1):
        date = today - timedelta(days = d)

        for a in range(0, args.atm):
            for seq in range(1, args.max_files):
                id = chr(ord('A') + a)
            
                ext = EXT[a % len(EXT)]
                name = '{}_{}_{}.{}'.format(id, date.strftime(FILENAME_FORMAT), seq, ext)
                path = os.path.join(dir, name)
            
                num_tx = random.randint(1, args.max_tx)
                gen = generate(num_tx, date)

                if ext == XML:
                    with open(path, 'wb') as f:
                        write_xml(gen, f)
                
                else:
                    with open(path, 'w') as f:
                        if ext == CSV:
                            write_csv(gen, f)
                        if ext == JSON:
                            write_json(gen, f)
                        if ext == YAML:
                            write_yaml(gen, f)
                
main()
