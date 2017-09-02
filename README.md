# HOMEWORK

Reimplement the following programs:

* cat
* grep
* cut
* ls
* head
* tail

in the following languages:

* Python
* C
* C++
* Go
* OCaml
* Haskell
* Rust
* Elixer
* Ruby
* Node.js
* Erlang
* Swift
* Lisp

and solve the following challenges:

* How many files are on my computer?
* What are the top 10 posts on reddit?

# General Requirements

* Do not use any third party libraries
* Package an installer for each program in the language's native package manager (e.g. npm, pypi, etc) // SHOUT OUT [@schep_](https://twitter.com/schep_)

# Language Requirements

## Python

* Must support python2 and python3

# Program Requirements

## cat

* At minimum: Read a file to stdout
* At minimum: Read multiple files to stdout
* At minimum: Output errors to stderr

## head

* At minimum: Read the first 10 lines of a file
* Extra credit: Read the first 10 lines of stdin
* Extra credit: Implement an option to define how many lines to read
* Extra credit: Implement an option to read bytes instead of lines
* Extra credit: Don't read the entire file

## tail

* At minimum: Read the last 10 lines of a file
* Extra credit: Read the last 10 lines of stdin
* Extra credit: Implement an option to define how many lines to read
* Extra credit: Implement an option to read bytes instead of lines
* Extra credit: Don't read the entire file

## ls

* At minimum: Handle multiple paths or files
* Extra credit: Include an option to show hidden files
* Extra credit: Include an option to show exact file size
* Extra credit: Include an option to show file size in human
* Extra credit: Implement separate terminal colors for files and directories

## grep

* At minimum: Read lines from a file and return lines that include a keyword
* Extra credit: Implement case-insensitivity flag
* Extra credit: Implement -v (ignore) flag
* Extra credit: Optionally apply keyword to stdin as well as a file
