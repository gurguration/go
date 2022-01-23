package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"customName"`
	Age  int    `json:"age,omitempty"`
}

func main() {
	p := Person{Name: "tom"}
	pBytes, err := json.Marshal(p)
	if err != nil {
		fmt.Errorf("Encountered error: %w", err)
	}
	fmt.Print(string(pBytes))
}
