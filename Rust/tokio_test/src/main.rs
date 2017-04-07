extern crate futures;
extern crate native_tls;
extern crate tokio_core;
extern crate tokio_tls;

use std::io;
use std::net::ToSocketAddrs;

use futures::Future;
use native_tls::TlsConnector;
use tokio_core::net::TcpStream;
use tokio_core::reactor::Core;
use tokio_tls::TlsConnectorExt;

fn main() {
    let mut core = Core::new().unwrap();
    let handle = core.handle();
    let addr = "blog.shouren.me:443".to_socket_addrs().unwrap().next().unwrap();

    let cx = TlsConnector::builder().unwrap().build().unwrap();
    let socket = TcpStream::connect(&addr, &handle);

    let tls_handshake = socket.add_then(|socket| {
        let tls = cx.connect_async("blog.shouren.me", socket);
        tls.map_err(|e| {io::Error::new(io::ErrorKind::Other, e)})
    });

    let request = tls_handshake.and_then(|socket| {
        tokio_core::io::write_all(socket,
            "GET / HTTP/1.1\r\n
             Host: blog.shouren.me\r\n\r\n
             ".as_bytes())
    });

    let response = request.add_then(|socket, _request| {
        tokio_core::io::read_to_end(socket, Vec::new())
    });

    let (_socket, data) = core.run(response).unwrap();
    println!("{}", String::from_utf8_lossy(&data));
}
