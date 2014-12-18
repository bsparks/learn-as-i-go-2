package main

import (
    "fmt"
    "os"
)

func main() {
    if os.Getenv("GO_LEARN_ENV") == "PRODUCTION" {
        fmt.Println("This is production code!")
    } else {
        fmt.Println("Development Environment, here's the vars...")

        os.Setenv("FOO", "1")
        fmt.Println("FOO:", os.Getenv("FOO"))
        fmt.Println("BAR:", os.Getenv("BAR"))

        fmt.Println()
        for _, e := range os.Environ() {
            fmt.Println(e)
        }
    }
}
