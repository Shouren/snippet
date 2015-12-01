from sqlalchemy.ext.declarative.api import declarative_base

vxlan_id = -1
vm_id = 0

BASEOBJ = declarative_base()

class base(BASEOBJ):
    """Base class for Nova and Glance Models"""
    __abstract__ = True
    __table_args__ = {'mysql_engine': 'InnoDB'}
    __table_initialized__ = False

class test(base):
    """ Base class for resource type object """
    __abstract__ = True

    _resource_name_ = None
    _alias_name_ = None

    __protected_attributes__ = set([
        "created_at", "updated_at", "deleted_at", "deleted"])

class g_var(object):
    def __init__(self):
        global vxlan_id
        pass

    @classmethod
    def print_cls(cls):
        print cls.__class__

    def print_obj(self):
        print self.__class__

class p_base(object):
    _length = 8

    def __init__(self):
        self.version = 0

class p_ext(p_base):
    _length = 16

    def __init__(self):
        super(p_ext, self).__init__()
        self.version = 1
        self.subtype = 0

    def __str__(self):
        return "%s %d" % (self.version, self.subtype)

class Configurable(object):
    __impl_class = None
    __impl_kwargs = None

    def __new__(cls, **kwargs):
        base = cls.configurable_base()
        print "cls: " + cls.__name__
        print "base: " + base.__name__
        if base is cls:
            impl = cls.configured_class()
        else:
            impl = cls
        print "impl: " + impl.__name__
        instance = super(Configurable, cls).__new__(impl)
        instance.initialize(**kwargs)
        return instance

    @classmethod
    def configured_class(cls):
        base = cls.configurable_base()
        if cls.__impl_class is None:
            base.__impl_class = cls.configurable_default()
        return base.__impl_class

class IOLoop(Configurable):

    @classmethod
    def configurable_base(cls):
        return IOLoop

    @classmethod
    def configurable_default(cls):
        return KQueueIOLoop


class KQueueIOLoop(IOLoop):
    def initialize(self, **kwargs):
        print "Hello World"

class property_test(object):
    def __init__(self):
        self.name = 'test'

    @property
    def is_true(self):
        return True if self.name == 'test' else False

if __name__ == "__main__":
    print 'Class testing'
    obj = property_test()
    if obj.is_true:
        print 'True'
    else:
        print False

