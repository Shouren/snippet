package demo

import org.springframework.web.bind.annotation.RestController
import org.springframework.web.bind.annotation.GetMapping
import java.util.concurrent.atomic.AtomicInteger


data class TestData(val key: String, val value: AtomicInteger)

@RestController
class DemoController {
    var counter = AtomicInteger()

    @GetMapping("/couting")
    fun counting(): TestData{
        counter.getAndIncrement()
        return  TestData("Hello", counter)
    }
}