import os
import random  # noqa: F401
import this  # noqa
import sys
from json import *

x = 1
if x == 1:
    # indented four spaces
    print("x is 1.")

    
def some_function():  # pylint: disable=missing-docstring
    print(
        "this should not raise a documentation related issue because it "
        "is silenced."
    )


def some_other_function():
    print("this should raise a documentation related issue.")

    # skipcq
    raise NotImplemented

    
def dummy_function():
    print('foo')

def new_dummy_function():
    print('foobar')
    raise NotImplemented  # skipcq: pyl-e0702
