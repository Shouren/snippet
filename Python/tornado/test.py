from tornado import web
from tornado import ioloop

class test(web.RequestHandler):
    def get(self):
        err = web.HTTPError(406, 'test')
        raise err


if __name__ == '__main__':
    app = web.Application([(r'/test', test)])
    app.listen(65534)
    ioloop.IOLoop.current().start()
