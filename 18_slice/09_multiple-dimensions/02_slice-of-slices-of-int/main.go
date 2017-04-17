package main

import "fmt"

func main() {
	transactions := make([][]int, 0)

	for i := 0; i < 4; i++ {
		transaction := make([]int, 0)

		for j := 0; j < 4; j++ {
			transaction = append(transaction, i + j)
		}

		transactions = append(transactions, transaction)
	}

	fmt.Println(transactions)
}
