package main

import (
	"encoding/json"
	"fmt"
)

type Item struct {
	A int   `json:"a"`
	B []int `json:"b"`
}

func main() {
	dataArray := []Item{
		{A: 1, B: []int{1, 2}},
		{A: 1, B: []int{4, 6}},
		{A: 1, B: []int{1, 2}},
		{A: 3, B: []int{1, 2}},
		{A: 3, B: []int{1, 2}},
		{A: 2, B: []int{1, 2, 3}},
	}

	// Khởi tạo map để đếm số lần xuất hiện
	countMap := make(map[string]int)

	// Duyệt qua từng phần tử trong dataArray
	for _, item := range dataArray {
		jsonKey, _ := json.Marshal(item) // Chuyển struct thành chuỗi JSON
		countMap[string(jsonKey)]++
	}

	result, _ := json.MarshalIndent(countMap, "", "  ")
	fmt.Println(string(result))
}
