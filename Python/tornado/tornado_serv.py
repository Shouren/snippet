#! /usr/bin/env python

import os
from tornado import process
from tornado.ioloop import IOLoop
from tornado.tcpserver import TCPServer
from tornado.netutil import bind_sockets

class FileServer(TCPServer):

    def handle_stream(self, stream, addr):
        self.stream = stream
        print "====== Handle new stream from ======"
        print process.task_id(), addr
        self.handle_request(stream)

    def handle_request(self, stream):
        f_handler = FileHandler(stream)
        stream.read_until('\n', f_handler.get_file_data)

class FileHandler(object):

    def __init__(self, stream):
        self.stream = stream

    def file_exist(self, path):
        # Make sure that file exist on the server side
        return os.path.exists(path)

    def get_file_data(self, file_path):
        print "====== Handle file request ======"
        print "====== " + file_path[0:-1] + " ======"

        file_path = file_path[0:-1]
        if self.file_exist(file_path):
            self.fd = open(file_path, 'r')
            self.send_file()
        else:
            self.stream.write('File Not Found')

    def send_file(self):
        content = self.fd.read(4096)
        if content != '':
            self.stream.write(content, self.send_file)
        else:
            self.fd.close()
            self.stream.close()

if __name__ == "__main__":
    print "====== Starting File Server ======"
    sockets = bind_sockets(6634)
    process.fork_processes(2)

    print os.getcwd()

    server = FileServer()
    server.add_sockets(sockets)
    IOLoop.instance().start()
