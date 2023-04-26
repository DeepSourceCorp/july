import os
import random
import this

x = 1
if x == 1:
    # indented four spaces
    print("x is 1.")

    
def some_function():
    print("this should raise a documentation related issue.")


def some_other_function():
    print("this should raise a documentation related issue.")

    # skipcq
    # some comment here
    # another comment
    raise NotImplemented


def return_none():
    return None


def dummy_function():
    print('foo')
    val = return_none()  # skipcq: pyl-e1128


def new_dummy_function():
    print('foobar')
    # This is a comment, we test whether the next line gets highlighted
    raise NotImplementedError  # raise an issue here
