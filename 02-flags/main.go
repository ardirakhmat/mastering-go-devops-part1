package main

import (
    "flag"
    "fmt"
)

func main() {
    name := flag.String("name", "world", "specify the name to greet")
    age := flag.Int("age", 0, "specify the age")
    flag.Parse()
    fmt.Printf("Hello, %s! You are %d years old.\n", *name, *age)
}
