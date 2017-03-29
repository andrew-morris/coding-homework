#!/usr/bin/env python

import sys

#TODO: Grep stdin
#TODO: Grep from a file

def main():
    
    arg = sys.argv[1]
    
    stdin = sys.stdin.read()
    tokenizedInput = stdin.split('\n')
    for line in tokenizedInput:
        if arg in line:
            sys.stdout.write(line + '\n')
        else:
            pass

    sys.exit(0)

if __name__ == "__main__":
    main()
