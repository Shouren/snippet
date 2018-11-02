package hello

val PI = 3.14

fun main(args: Array<String>) {
    println("Hello World?")
    println(PI)
    stringReplace()
}

fun stringReplace() {
    var a = 1
    // simple name in template:
    val s1 = "a is $a"

    a = 2
    // arbitrary expression in template:
    val s2 = "${s1.replace("is", "was")}, but now is $a"
    println(s2)
}