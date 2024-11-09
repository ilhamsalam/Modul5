package main

import "fmt"

func findFactors(number, divisor int) {
        if divisor > number {
                return
        }
        if number%divisor == 0 {
                fmt.Print(divisor, " ")
        }
        findFactors(number, divisor+1)
}

func main() {
        var num int
        fmt.Print("Masukkan bilangan: ")
        fmt.Scan(&num)

        fmt.Println("Faktor dari", num, "adalah:")
        findFactors(num, 1)
        fmt.Println()
}