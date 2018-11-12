package demo

import com.github.kittinunf.fuel.httpGet

fun main(args: Array<String>) {
    println("Start http probing...")

    var startTime = System.currentTimeMillis()
    var (req, resp, _) = "https://www.google.com/generate_204".httpGet().response()
    var endTime = System.currentTimeMillis()

    println(message = "->")
    println(message = "URL: %s".format(req.url))
    println(message = "<-")
    println(message = "StatusCode: %s".format(resp.statusCode))
    println(message = "Latency: %sms".format((endTime - startTime)))
}
