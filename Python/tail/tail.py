#!/usr/bin/env python

import sys

#TODO: Add arg for amount of lines
#TODO: Add arg for amount of bytes
#TODO: Support reading from a file
#TODO: Support reading from stdin
#TODO: Fix bug that adds extra newline at the end

def main():
    data = sys.stdin.read()
    lines = data.split('\n')
    for line in lines[-10:]:
        sys.stdout.write(line + '\n')

if __name__ == '__main__':
    main()
