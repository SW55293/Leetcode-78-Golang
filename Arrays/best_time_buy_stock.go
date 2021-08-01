package main

func maxProfit(stock []int) int {

	if len(stock) <= 1 {
		return 0
	}

	min := stock[0]
	profit := 0

	for _, value := range stock {

		if value < min {
			min = value
		} else {
			profit = largerValue(profit, (value - min))
		}

	}
	return profit
}

func largerValue(a, b int) int {
	if a > b {
		return a
	}
	return b
}
