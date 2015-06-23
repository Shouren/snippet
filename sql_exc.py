#! /usr/bin/env python
# from cls import g_var
# from cls import test
from cls import p_ext

# class b(test):
#     _resource_name_ = 'network'
#     __tablename__ = 'networks_tbl'

def func_with_kwargs(param, **kwargs):
    print param
    print kwargs['data']['data']

if __name__ == "__main__":
    # a = g_var()
    # a.print_cls()
    # a.print_obj()
    # c = b()
    # print type(a)
    # print type(b)
    # print type(c)

    #res = p_ext()
    #print res.version
    #print res._length
    #print res.subtype
    #print res
    data = {'data': {'data': 1}}
    func_with_kwargs('hello', **data)
