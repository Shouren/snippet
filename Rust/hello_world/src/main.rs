extern crate reqwest;

use std::env::var;
use std::time::Instant;
use reqwest::Url;



fn get_proxy_addr() -> String {
    let default_proxy_addr= "http://127.0.0.1:1087".to_string();
    let key = "HTTP_PROXY";
    match var(key) {
        Ok(val) => return val,
        Err(_)=> return default_proxy_addr,
    }
}

fn main() {
    println!("Start proxy testing...");
    let proxy = reqwest::Proxy::all(Url::parse(get_proxy_addr().as_str()).unwrap()).unwrap();

    let client = reqwest::Client::builder().proxy(proxy).build().unwrap();

    let now = Instant::now();
    let resp = client.get("https://www.google.com/generate_204").send().unwrap();
    let latency = now.elapsed();

    println!("Status Code: {:?}", resp.status());
    println!("Time Duration: {:?}",  latency);
}
