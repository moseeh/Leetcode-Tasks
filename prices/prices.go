package main

import (
    "fmt"
    "math"
)

// Function to return the maximum profit from buying and selling one stock
func maxProfit(prices []int) int {
    if len(prices) == 0 {
        return 0
    }

    minPrice := math.MaxInt32
    maxProfit := 0

    for _, price := range prices {
        if price < minPrice {
            minPrice = price
        } else if price - minPrice > maxProfit {
            maxProfit = price - minPrice
        }
    }

    return maxProfit
}

func main() {
    prices1 := []int{7, 1, 5, 3, 6, 4}
    fmt.Println("Max profit for prices1:", maxProfit(prices1)) // Output: 5

    prices2 := []int{7, 6, 4, 3, 1}
    fmt.Println("Max profit for prices2:", maxProfit(prices2)) // Output: 0
}
