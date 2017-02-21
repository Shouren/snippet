#! /usr/bin/python3

import time
import threading

from sqlalchemy import create_engine
from sqlalchemy import Column, Integer, String, DateTime

from sqlalchemy.orm import sessionmaker, scoped_session

from sqlalchemy.ext.declarative import declarative_base


Base = declarative_base()

engine = create_engine('postgresql:///test_1.db', echo=True)
Session = scoped_session(sessionmaker(bind=engine))


class IPs(Base):
    __tablename__ = 'ips'

    id = Column(Integer, primary_key=True)
    address = Column(String)
    task_id = Column(Integer, nullable=False)

    def __init__(self, address, task_id):
        self.address = address
        self.task_id = task_id

    def __repr__(self):
        return "<IPs({0}, {1}, {2})>".format(self.id, self.address, self.task_id)


class Tasks(Base):
    __tablename__ = 'tasks'

    id = Column(Integer, primary_key=True)
    complete_time = Column(DateTime, nullable=False)

    def __init__(self, complete_time):
        self.complete_time = complete_time

    def __repr__(self):
        return "<Task({0}, {1})>".format(self.id, self.complete_time)


class Whitelist(Base):
    __tablename__ = 'whitelist'

    ip = Column(String, primary_key=True)

    def __init__(self, ip):
        self.ip = ip

    def __repr__(self):
        return "<Whitelist({0})>".format(self.ip)


class test(threading.Thread):
    def __init__(self):
        threading.Thread.__init__(self)

    def run(self):
        time.sleep(1)
        session = Session()
        ret = session.query(Whitelist).all()
        print(ret)
        Session.remove()

if __name__ == '__main__':
    Base.metadata.create_all(engine)
    t = test()
    t.start()
    print(2)
    session = Session()
    ret = session.query(IPs).all()
    print(ret)
    session.add(Whitelist('103.37.149.31'))
    session.commit()
