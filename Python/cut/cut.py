#!/usr/bin/env python

import sys

#TODO: Make the argument order flexible
#TODO: Add error handling
#TODO: Add support for multiple delimiters
#TODO: Add support for first and last fields

def outputUsage():
    print('Usage: something | python cut.py -d deliminator -f field_number')
    print('Example: something | python cut.py -d \':\' -f 3')
    sys.exit(1)

def main():
    
    # count amount of args, if not the expected amount, print usage
    if len(sys.argv) != 5:
        outputUsage()

    delimitor = sys.argv[2]
    field = int(sys.argv[4])

    data = sys.stdin.read()
    for line in data.split('\n'):
        print(line.split(delimitor)[field - 1])

if __name__ == '__main__':
    main()
