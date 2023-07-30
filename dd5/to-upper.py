#!/usr/bin/env python3

import sys, os

p = os.environ['PASSWORD'] # environment
print(p) # stdout (fp==1)

s = sys.argv[1].upper() # argument
print(s, file=sys.stderr) # (fp==2)

in_string = input() # read from stdin (fp==0)
while in_string != 'quit':
  print(f"You entered {in_string}") # write to stdout
  in_string = input() 

