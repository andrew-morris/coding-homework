#!/usr/bin/env python

import sys

#TODO: Grep stdin
#TODO: Grep from a file

def main():
    
    # Figure out if we're grepping a file or stdin
    if sys.stdin.read():
        print('stdin detected')
    else:
        print('stdin not detected')

    sys.exit()
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
