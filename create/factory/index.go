package factory

import "fmt"

func printDetails(g iGun) {
	fmt.Printf("Gun %s", g.getName())
	fmt.Println()
	fmt.Printf("Power %d", g.getPower())
}

func TestFactory() {
	ak47, _ := getGun("ak47")
	musket, _ := getGun("musket")

	printDetails(ak47)
	printDetails(musket)
}