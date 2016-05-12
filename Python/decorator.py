class OP(object):
    self.op = []

def tags(func_b):
    def tags_decorator(func_a):
        def func_wrapper(name):
            return "<{0}>{1}</{0}>".format('test', func_b(name))
        return func_wrapper
    return tags_decorator

def get_text_b(name):
    return 'hello ' + name

@tags(get_text_b)
def get_text_a(name):
    return "Hello " + name

print get_text_a("John")

# Outputs <p>Hello John</p>
