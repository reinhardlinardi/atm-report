# ATM report service dataset remover

import os

def main():
    dir = '../dataset'

    try:
        files = os.listdir(dir)

        for file in files:
            path = os.path.join(dir, file) 
            os.remove(path)

    except OSError:
        print('remove files failed')

main()
