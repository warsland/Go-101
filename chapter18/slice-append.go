package main

import "fmt"

func dunp(label string, slice []string) {
	fmt.Printf("%v: length %v, vapacity %v %v\n", label, len(slice), cap(slice), slice)
}

func main() {
	dwarf1 := []string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}
	dwarf2 := append(dwarf1, "Orcus")
	dwarf3 := append(dwarf1, "Salacia", "Quaoar", "Sedna")
	dunp("dwarf1", dwarf1)
	dunp("dwarf2", dwarf2)
	dunp("dwarf3", dwarf3)
	dwarf2[0] = "Sun"
	fmt.Printf("%v\n%v\n%v", dwarf1, dwarf2, dwarf3)
}
