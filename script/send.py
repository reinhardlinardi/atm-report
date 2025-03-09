# ATM report service dataset producer (send dataset file to consumer)

import os
import shutil

SRC = '../gen'
DEST = '../dataset'

def listdir(src):
    arr = []

    try:
        files = os.listdir(src)
        
        for file in files:
            arr.append(file)
        
        arr.sort()
        
        for i in range(0, len(arr)):
            print('{}. {}'.format(i+1, arr[i]))
    
    except OSError:
        return None, False
    
    return True, arr

def copyfile(src, dest, name):
    try:
        shutil.copy(os.path.join(src, name), dest)

    except shutil.Error:
        return False
    
    return True

def prompt():
    print()
    print('Send file number: ', end = '')

def sent():
    print('File sent.')
    print()

def invalid():
    print('Invalid. Please try again.')
    print()

def main():
    print('Type number and press Enter to send file.')
    print('To exit, press Ctrl+C.')
    print()

    while True:
        try:
            ok, arr = listdir(SRC)
            if not ok:
                print('list error')
                break

            prompt()
            s = input().strip()

            if not s.isdigit():
                invalid()
                continue

            num = int(s)
            
            if num == 0 or num > len(arr):
                invalid()
                continue

            ok = copyfile(SRC, DEST, arr[num-1])
            if not ok:
                print('copy error')
                break

            sent()

        except KeyboardInterrupt:
            break

main()
