package main

import "fmt"

func printOddNumbers(n int) {
        if n < 1 {
                return
        }
        printOddNumbers(n - 2)
        fmt.Print(n, " ")
}

func main() {
        var num int
        fmt.Print("Masukkan bilangan: ")
        fmt.Scan(&num)

        if num % 2 == 0 {
                num-- // Jika bilangan genap, kurangi 1 agar dimulai dari bilangan ganjil terdekat
        }

        fmt.Println("Barisan bilangan ganjil:")
        printOddNumbers(num)
        fmt.Println()
}