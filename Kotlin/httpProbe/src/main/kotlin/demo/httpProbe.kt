package demo

import com.github.kittinunf.fuel.httpGet
import kotlin.system.measureTimeMillis

fun main(args: Array<String>) {
    println("Start http probing...")

    var msg = ""

    val latency = measureTimeMillis {
        val (req, resp, _ ) = "https://www.google.com/generate_204".httpGet().response()
        msg = "->\nURL: %s\n<-\nStatusCode : %s\n".format(req.url, resp.statusCode)
    }
    println(message = msg)
    println(message = "Latency: %sms".format((latency)))
}