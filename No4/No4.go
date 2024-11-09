package main

import "fmt"

func printSequence(n int) {
        if n == 1 {
                fmt.Print(n, " ")
                return
        }
        fmt.Print(n, " ")
        printSequence(n - 1)
        if n > 1 {
                fmt.Print(n, " ")
        }
}

func main() {
        var num int
        fmt.Print("Masukkan bilangan: ")
        fmt.Scan(&num)

        fmt.Println("Barisan bilangan:")
        printSequence(num)
        fmt.Println()
}