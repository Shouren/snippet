rb_stack = []

def _roll_back(rb_func):
    def _fn(func):
        def func_wrapper(*args, **kwargs):
            ret = func(*args, **kwargs)
            print ret
            if ret is False:
                rb_func()
            rb_stack.append(1)
            return ret
        return func_wrapper
    return _fn

def test_rb():
    print 'hello'

@_roll_back(test_rb)
def test(a):
    return False

test(1)

test(2)

print rb_stack
