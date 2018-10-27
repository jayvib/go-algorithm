package main

import "fmt"

// Problem: Given a list, in which nth element is the price of the stock on nth day.
// You are asked to but once and sell once, on what date you will be buring and at what date
// you will be selling to get maximum profit.

func maxProfit(stocks []int) {
	size := len(stocks)
	buy := 0
	sell := 0
	curMin := 0
	currProfit := 0
	maxProfit := 0

	for i := 0; i < size; i++ {
		if stocks[i] < stocks[curMin] { // always check if the current item is a current minimum
			curMin = i
		}

		currProfit = stocks[i]-stocks[curMin] // given a current minimum calculate the profit on the particular element
		if currProfit > maxProfit {
			buy = curMin // track the buy date
			sell = i // track the sell date
			maxProfit = currProfit	 // this will be the currProfit for this particular buy-sell date.
		}
	}

	fmt.Println("Purchase day is-", buy, " at price ", stocks[buy])
	fmt.Println("Sell day is-", sell, " at price ", stocks[sell])
	fmt.Println("Max Profit:", maxProfit)
}
