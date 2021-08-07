package main

import "fmt"

func main() {
	anAk47, _ := getGun("ak47")
	anMaverick, _ := getGun("maverick")
	printDetails(anAk47)
	printDetails(anMaverick)
}

func printDetails(g iGun) {
	fmt.Printf("Gun: %s\n", g.getName())
	fmt.Printf("Power: %d\n", g.getPower())
}
