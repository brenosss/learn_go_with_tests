package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const spanish = "spanish"
const french = "french"

func Hello(name string, language string) string {
    if name == "" {
        name = "World"
    }
    switch language{
        case spanish:
            return spanishHelloPrefix + name
        case french:
            return frenchHelloPrefix + name
        default:
            return englishHelloPrefix + name
    }
}

func main() {
    fmt.Println(Hello("world", ""))
}
