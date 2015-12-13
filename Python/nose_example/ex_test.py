#! /usr/bin/env python

import os
import sys

from nose import config
from nose import core
from nose.tools import with_setup

conf = config.Config(stream=sys.stdout,
                     env=os.environ,
                     verbosity=int(os.environ.get('NOSE_VERSOSE', 3)),
                     includeExe=True,
                     traverseNamespace=True,
                     plugins=core.DefaultPluginManager())

var = 0

def setup_env():
    print 'Setting up environment'

def teardown_env():
    print 'Teardown enviroment'

@with_setup(setup_env, teardown_env)
def test():
    print 'Hello World'

def b_test():
    global var
    var += 1
    assert var == 1

def a_test():
    global var
    var += 1
    print var
    assert var == 3

def c_test():
    global var
    var += 1
    assert var == 3
