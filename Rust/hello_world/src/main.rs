use std::time::Instant;

use reqwest::Url;
use clap::{Arg, App};

fn main() {
    let params= App::new("Proxy time measure program")
        .version("0.1.0")
        .author("Shouren <yangshouren@gmail.com>")
        .arg(Arg::with_name("proxy")
            .short("p")
            .long("proxy")
            .takes_value(true)
            .help("Proxy address for measurement")
            .required(true))
        .arg(Arg::with_name("target")
            .short("t")
            .long("target")
            .takes_value(true)
            .help("Target address to send request for measurement")
            .default_value("https://www.google.com/generate_204"))
        .get_matches();

    let proxy_addr = params.value_of("proxy").expect("Missing proxy address");
    let target_addr = params.value_of("target").expect("Missing target address");
    println!("Start proxy testing...");
    println!("# Proxy: {}", proxy_addr);
    println!("# Target: {}", target_addr);

    let proxy = reqwest::Proxy::all(
        Url::parse(proxy_addr).expect("Invalid proxy address")
    ).expect("Create proxy failed");

    let client = reqwest::Client::builder().proxy(proxy).build().expect("Create client failed");

    println!("Getting proxy response...");
    let now = Instant::now();
    let resp = client.get(target_addr).send().expect(format!("Request URL({}) failed", target_addr).as_str());
    let latency = now.elapsed();

    println!("# Status Code: {:?}", resp.status());
    println!("# Time Duration: {:?}",  latency);
}