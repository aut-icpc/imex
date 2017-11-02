/*
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 02-11-2017
 * |
 * | File Name:     main.go
 * +===============================================
 */

package main

import (
	"encoding/json"
	"fmt"
	"os"

	"./types"
)

func main() {
	var rs map[string]map[int]imex.Register

	f, err := os.Open("export.json")
	if err != nil {
		panic(err)
	}
	json.NewDecoder(f).Decode(&rs)

	// Onsite
	fmt.Println("Onsite:")
	for id, r := range rs["onsite"] {
		fmt.Printf("%d: %v\n", id, r)
	}

	// Online
	fmt.Println("Online:")
	for id, r := range rs["onsite"] {
		fmt.Printf("%d: %v\n", id, r)
	}
}
