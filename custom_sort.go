package main 

import (
	"fmt"
	"sort"
)

type byLength []string 
func (m byLength) Len() int {
	return len(m)
}

func (m byLength) Less(i, j int) bool {
	return len(m[i]) < len(m[j])
}

func (m byLength) Swap(i, j int)  {
	m[j], m[i] = m[i], m[j]
} 

func main() {
	m1 := []string{"apple", "man", "pear", "banana"}
	sort.Sort(byLength(m1))
	fmt.Println(m1)
}
