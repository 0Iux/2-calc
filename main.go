package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var operation, nums string
	var nums_slice []int
	// В прошлом дз пришел коментарий, что лучше использовать одинаковые способы инициализации переменных
	// Мне удобней через := но я не могу же сделать такое присваивания для переменно которую я считываю из терминала
	// так как лучше быть ?
	for {
		operation, nums = user_input()
		nums_slice = make_slice(nums)
		user_output(operation, nums_slice)
		if !check_repeat() {
			break
		}
	}
}

func user_output(operation string, nums []int) {
	switch operation {
	case "AVG":
		fmt.Printf("AVG: %.2f\n", avg(nums))
	case "SUM":
		fmt.Printf("SUM: %.2f\n", sum(nums))
	case "MED":
		fmt.Printf("MED: %.2f\n", med(nums))
	}
}

func check_repeat() bool {
	fmt.Println("----------------------------Меню----------------------------")
	var choise string
	fmt.Println("Хотите ли вы еще раз рассчитать? (y/n): ")
	fmt.Scan(&choise)
	if choise == "n" || choise == "N" {
		return false
	}
	return true
}

func med(nums []int) float64 {
	lenght := len(nums)
	if lenght%2 != 0 {
		fmt.Printf("nums[lenght/2]: %v\n", nums[lenght/2])
		return float64(nums[lenght/2])
	} else {
		return (float64(nums[lenght/2]+nums[lenght/2-1]) / 2)
	}
}

func avg(nums []int) float64 {
	return float64(sum(nums) / float64(len(nums)))
}

func sum(nums []int) float64 {
	var ans int
	for _, elem := range nums {
		ans += elem
	}
	return float64(ans)
}

func make_slice(nums string) []int {
	var nums_slice []string
	var ans_slice []int
	nums_slice = strings.Split(nums, ",")
	for _, elem := range nums_slice {
		int_elem, _ := strconv.Atoi(elem)
		ans_slice = append(ans_slice, int_elem)
	}
	return ans_slice
}

func user_input() (string, string) {
	var operation, nums string
	fmt.Print("Введите операцию (AVG, SUM, MED): ")
	fmt.Scan(&operation)
	fmt.Print("Введите числа через запятую без пробела: ")
	fmt.Scan(&nums)
	return operation, nums
}
