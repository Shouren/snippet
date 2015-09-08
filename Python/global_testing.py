#! /usr/bin/env python

print locals()

import global_test

a = 3

print locals()

global_test.testing()

print global_test.test_res
print locals()
