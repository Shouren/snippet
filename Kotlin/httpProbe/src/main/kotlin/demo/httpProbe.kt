package demo

import com.github.kittinunf.fuel.httpGet

fun main(args: Array<String>) {
    println("Start http probing...")

    var (req, resp, ret) = "https://www.google.com/generate_204".httpGet().responseString()

    println(req)
    println(resp)
    println(ret)

}
