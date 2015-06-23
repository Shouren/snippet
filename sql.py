#! /usr/bin/env python

from sqlalchemy import create_engine
#from sqlalchemy import and_
from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy import Column, Integer, String#, func
from sqlalchemy.orm import sessionmaker

Base = declarative_base()
engine = create_engine('sqlite:///:memory:', echo=True)
Session = sessionmaker(bind=engine)

class User(Base):
    __tablename__ = 'users'

    id = Column(Integer, primary_key=True)
    name = Column(String)
    fullname = Column(String, nullable=True)
    password = Column(String)
    department = Column(String)

    def __init__(self, name, fullname, password, department):
        self.name = name
        self.fullname = fullname
        self.password = password
        self.department = department

    def __repr__(self):
        return "<User('%s', '%s', '%s', '%s')>" % (self.name, self.fullname, self.password, self.department)

    @classmethod
    def get_name(cls):
        return cls.name

class Department(Base):
    __tablename__ = 'department'

    id = Column(Integer, primary_key=True)
    name = Column(String)
    location = Column(String)

    def __init__(self, name, location):
        self.name = name
        self.location = location

    def __repr__(self):
        return "<Department('%s')>" % self.name

if __name__ == "__main__":
    Base.metadata.create_all(engine)
    session = Session()

    user = User('Sichao', None, 'xiangyoukanqi', 'Design')
    session.add(user)
    user = User('miaoge', 'Zhaoyuan', 'miaoge', 'Design')
    session.add(user)
    user = User('shouren', 'yangshouren', '0192837465', 'tech')
    session.add(user)
    user = User('kezi', 'yukezi', 'ykz', 'research')
    session.add(user)
    dep = Department("Design", "BJ")
    session.add(dep)
    dep = Department("research", "NYK")
    session.add(dep)
    session.commit()
    #q = session.query(User).outerjoin(Department, and_(User.department==Department.name)).filter(Department.location=="NYK")
    #for item in q.all():
    #    print item
    q = session.query(User).filter(User.name=='Sichao')
    for item in q.all():
        print item
        if item.fullname == None:
            print 'Equal'
