// extern crate futures;
// extern crate hyper;
// extern crate tokio_core;

extern crate pretty_env_logger;

// use std::env;
// use std::io;;{self, Write};

// use futures::Future;
// use futures::stream::Stream;

extern crate hyper;
extern crate hyper_native_tls;

use hyper::Client;
use hyper::net::HttpsConnector;
use hyper_native_tls::NativeTlsClient;
use std::io::Read;

use hyper::Client;

fn main() {
    pretty_env_logger::init().unwrap();
    let ssl = NativeTlsClient::new().unwrap();
    let connector = HttpsConnector::new(ssl);
    let client = Client::with_connector(connector);
    let url = "http://www.google.com/robots.txt";
    println!("Get from {0} ...", url);
    let client = Client::new();
    let resp = client.get(url).send().unwrap();
    println!("Statut of Response: {}", resp.status);
    assert_eq!(resp.status, hyper::Ok);
}
