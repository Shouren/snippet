use std::io::{Read, Write};
use std::net::{TcpListener, TcpStream};

fn handle_client(mut stream: TcpStream) {
    println!("Client Addr: {}", stream.peer_addr().unwrap_or_else(|e| panic));

    let mut buffer: [u8; 3] = [0; 3];
    stream.read(buffer.as_mut()).unwrap();
    println!("{}", String::from_utf8_lossy(buffer.as_ref()));

    stream.write(String::from("Hello World\n").as_bytes()).unwrap();
}

fn main() {
    let listener = TcpListener::bind("127.0.0.1:8080")\
                   .unwrap_or_else(|e| panic!("Failed to bind: {}", e));

    for stream in listener.incoming() {
        match stream {
            Ok(stream) => {
                handle_client(stream);
            }
            Err(e) => {println!("Connetion Failed: {}", e)}
        }
    }
}
