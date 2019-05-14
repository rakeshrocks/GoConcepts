package main

import "fmt"
import "time"

func main() {

    reqs := make(chan int, 5)
    for i := 1; i <= 5; i++ {
        reqs <- i
    }
    close(reqs)

    limiter := time.Tick(2 * time.Second)

    for req := range reqs {
        <-limiter
        fmt.Println("Request being processed:", req, time.Now())
    }
}
