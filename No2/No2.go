package main

import "fmt"

func printStars(n int) {
        if n == 1 {
                fmt.Println("*")
                return
        }

        printStars(n - 1)
        for i := 0; i < n; i++ {
                fmt.Print("*")
        }
        fmt.Println()
}

func main() {
        var n int
        fmt.Print("Masukkan nilai N: ")
        fmt.Scan(&n)

        printStars(n)
}