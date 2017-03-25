#!/usr/bin/env python

import sys

#TODO: Need to go back and write this in such a way that I
# don't read in the entire file to memory before printing it
def readFile(filename):
    try:
        f = open(filename, 'rb')
        data = f.read()
        sys.stdout.write(data)
        f.close()
        return
    
    except Exception as e:
        sys.stderr.write('%s\n' % str(e))
        pass

def main():
    for filename in sys.argv[1:]:
        readFile(filename)
    sys.exit(0)

if __name__ == '__main__':
    main()
