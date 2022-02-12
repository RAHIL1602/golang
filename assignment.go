package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var order string
	var quantity int
	var sales = []int{}
	counter := 0
	for true {
		fmt.Println("---------MENU----------")
		fmt.Print("C: coffee, 50rs\n D: dosa, 90 rs\n T: tomato soup, 30rs\n I : idli , 30rs\n V : vada, 35rs.\n B: bhature &chhole, 40rs\n P: paneer pakoda, 40rs\n M: manchurian, 100rs\n H: hakka noodle, 80rs.\n F: French fries, 40rs\n J: jalebi, 40rs\n L: lemonade, 25rs\n S: spring roll, 30rs\n")
		fmt.Println("-----------------------")
		fmt.Println("Press END For closing the day.")
		fmt.Println("Please place the order: ")
		fmt.Scan(&order)
		order = strings.ToUpper(order)

		if order == "END" {
			total_income(sales)
		} else {
			fmt.Scan(&quantity)
		}
		bill := bill(quantity, order)
		fmt.Println("####################")
		fmt.Println("Total bill: ", bill)
		fmt.Println("####################")
		sales = append(sales, bill)
		counter++

	}
}

func total_income(sale []int) {
	var sum int = 0

	for i := 0; i < len(sale); i++ {
		sum = sale[i] + sum
	}
	fmt.Println("Total Income for the day is : ", sum)
	os.Exit(0)
}

func bill(quantity int, order string) int {
	m := map[string]int{
		"C": 50,
		"D": 90,
		"T": 30,
		"I": 30,
		"V": 35,
		"B": 40,
		"P": 40,
		"M": 100,
		"H": 80,
		"F": 40,
		"J": 40,
		"L": 25,
		"S": 30,
	}
	billy := quantity * m[order]
	return billy
}
