use std::time::Instant;

extern crate reqwest;


fn main() {
    println!("Hello, world!");

    let proxy = reqwest::Proxy::https("http://localhost:1087").unwrap();

    let client = reqwest::Client::builder().proxy(proxy).build().unwrap();

    let now = Instant::now();
    let resp = client.get("https://www.google.com/generate_204").send().unwrap();
    let latency = now.elapsed();

    println!("Status Code: {:?}", resp.status());
    println!("Time Duration: {:?}",  latency);
}
