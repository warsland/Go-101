package main

import (
	"fmt"
	"sort"
)

func sortStrings(s []string, less func(i, j int) bool) {
	if less == nil {
		less = func(i, j int) bool { return s[i] < s[j] }
	}
	sort.Slice(s, less)
}

func main() {
	foods := []string{"onion", "carrot", "celery"}
	sortStrings(foods, nil)
	fmt.Println(foods)
}
