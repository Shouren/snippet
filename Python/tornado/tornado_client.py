#! /usr/bin/env python

import os
import sys
import socket
import functools
from tornado.ioloop import IOLoop
from tornado.iostream import IOStream

class file_client(object):

    def __init__(self, file_path):
        self.file_path = file_path

    def add_handler(self, sock, handler, event):
        io_loop = IOLoop.instance()
        callback = functools.partial(handler, sock)
        try:
            io_loop.remove_handler(sock.fileno())
        except:
            pass
        io_loop.add_handler(sock.fileno(), callback, event)

    def add_callback(self, func, *argc):
        io_loop = IOLoop.instance()
        callback = functools.partial(func, *argc)
        io_loop.add_callback(callback)

    def connect(self, host, port):
        #connect to server
        s = socket.socket(socket.AF_INET, socket.SOCK_STREAM, 0)
        s.setsockopt(socket.SOL_SOCKET, socket.SO_REUSEADDR, 1)
        s.connect((host, port))
        s.setblocking(0)
        self.add_handler(s, self.request_file, IOLoop.WRITE)

    def request_file(self, sock, fd, events):
        sock.send(self.file_path + '\n')
        self.add_handler(sock, self.handle_recv_data, IOLoop.READ)

    def get_store_path(self):
        self.store_dir = os.getcwd() + '/client'
        return self.store_dir + '/' + self.file_path.split('/')[-1]

    def save(self, data):
        if data != '':
            self.fd.write(data)
        else:
            self.fd.close()
            print "File is saved to " + self.store_path
            IOLoop.instance().stop()

    def handle_recv_file(self, sock, fd, events):
        try:
            data = sock.recv(4096)
        except socket.error as err:
            print err
            pass

        self.add_callback(self.save, data)

    def handle_recv_data(self, sock, fd, events):
        data = sock.recv(15)
        if data == 'File Not Found!':
            print "Server Return: File Not Found"
            IOLoop.instance().stop()
            return

        self.store_path = self.get_store_path()
        self.fd = open(self.store_path, 'w')
        self.fd.write(data)
        self.add_handler(sock, self.handle_recv_file, IOLoop.READ)

if __name__ == '__main__':
    file_name = sys.argv[1]
    Host, Port = 'localhost', 6634

    client = file_client(file_name)
    client.connect(Host, Port)
    IOLoop.instance().start()
