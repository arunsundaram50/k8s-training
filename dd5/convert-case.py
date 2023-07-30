#!/usr/bin/env python3

import sys, os

type_of_conversion = os.environ['TYPE_OF_CONVERSION']

if type_of_conversion=="upper":
  s = sys.argv[1].upper()
else:
  s = sys.argv[1].lower()

print(s)

