package factory_abs

import "fmt"

func printShoeDetail(shoe iShoe) {
	fmt.Printf("shoe brand is %s", shoe.getLogo())
	fmt.Println()
	fmt.Printf("shoe size is %d", shoe.getSize())
	fmt.Println()
}

func printShirtDetail(shirt iShirt) {
	fmt.Printf("shirt brand is %s", shirt.getLogo())
	fmt.Println()
	fmt.Printf("shirt size is %d", shirt.getSize())
	fmt.Println()
}

func TestAbsFactory() {
	adidas, _ := getSportsFactory("adidas")
	nike, _ := getSportsFactory("nike")

	adidasShoe := adidas.makeShoe()
	adidasShirt := adidas.makeShirt()

	nikeShoe := nike.makeShoe()
	nikeShirt := nike.makeShirt()

	printShoeDetail(adidasShoe)
	printShoeDetail(nikeShoe)

	printShirtDetail(adidasShirt)
	printShirtDetail(nikeShirt)
}
