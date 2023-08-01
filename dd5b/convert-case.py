#!/usr/bin/env python3

import sys, os

op = os.environ.get('op', 'lower')

if op == "upper":
  s = sys.argv[1].upper()
else:
  s = sys.argv[1].lower()

print(s)

