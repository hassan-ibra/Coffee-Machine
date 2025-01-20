package main

import (
	"fmt"
)

type coffee_machine struct{
	water int
	milk int
	coffee_beans int
	disposable_cups int
	money int
	
}

func (cm *coffee_machine)buy(){
	// For one espresso, the coffee machine needs 250 ml of water and 16 g of coffee beans. It costs $4.
	espresso_water := 250
	espresso_coffeeB := 16
	espresso_costs := 4
	// For a latte, the coffee machine needs 350 ml of water, 75 ml of milk, and 20 g of coffee beans. It costs $7.
	latte_water := 350
	latte_milk := 75
	latte_coffeeB := 20
	latte_costs := 7
	// And for a cappuccino, the coffee machine needs 200 ml of water, 100 ml of milk, and 12 g of coffee beans. It costs $6.
	cappuccino_water := 200
	cappuccino_milk := 100
	cappuccino_coffeeB := 12
	cappuccino_costs := 6
	
	fmt.Println("What do you want to buy? 1 - espresso, 2 - latte, 3 - cappuccino, back - to main menu:")
	
	var coffee_option int
	fmt.Scan(&coffee_option)
	switch coffee_option {
	case 1:
		if cm.water >= espresso_water && cm.coffee_beans >= espresso_coffeeB && cm.disposable_cups >0{
			fmt.Println("I have enough resources, making you a coffee!")
			cm.water -= espresso_water
			cm.coffee_beans -= espresso_coffeeB
			cm.disposable_cups -= 1
			cm.money += espresso_costs
		}else{
			if cm.water < cappuccino_water{
				fmt.Println("Sorry, not enough water!")
			}else if cm.milk < cappuccino_milk{
				fmt.Println("Sorry, not enough milk!")
			}else if cm.coffee_beans >= cappuccino_coffeeB {
				fmt.Println("Sorry, not enough coffee bean!")
			}else{
				fmt.Println("Sorry, not enough disposable cups!")
			}
		}	
	case 2:
		if cm.water >= latte_water && cm.milk >= latte_milk && cm.coffee_beans >= latte_coffeeB && cm.disposable_cups > 0{
			fmt.Println("I have enough resources, making you a coffee!")
			cm.water -= latte_water
			cm.milk -= latte_milk
			cm.coffee_beans -= latte_coffeeB
			cm.disposable_cups -= 1
			cm.money+= latte_costs
			
		}else{
			if cm.water < cappuccino_water{
				fmt.Println("Sorry, not enough water!")
			}else if cm.milk < cappuccino_milk{
				fmt.Println("Sorry, not enough milk!")
			}else if cm.coffee_beans >= cappuccino_coffeeB {
				fmt.Println("Sorry, not enough coffee bean!")
			}else{
				fmt.Println("Sorry, not enough disposable cups!")
			}
		}	
	case 3:
		if cm.water >= cappuccino_water && cm.milk >= cappuccino_milk && cm.coffee_beans >= cappuccino_coffeeB && cm.disposable_cups > 0{
			fmt.Println("I have enough resources, making you a coffee!")
			cm.water -= cappuccino_water
			cm.milk -= cappuccino_milk
			cm.coffee_beans -= cappuccino_coffeeB
			cm.disposable_cups -= 1
			cm.money+= cappuccino_costs	
		}else{
			if cm.water < cappuccino_water{
				fmt.Println("Sorry, not enough water!")
			}else if cm.milk < cappuccino_milk{
				fmt.Println("Sorry, not enough milk!")
			}else if cm.coffee_beans >= cappuccino_coffeeB {
				fmt.Println("Sorry, not enough coffee bean!")
			}else{
				fmt.Println("Sorry, not enough disposable cups!")
			}
		}		
	}
	
}
func (cm *coffee_machine) fill(){
	var fill_milk, fill_water, fill_coffee_beans, fill_disposable_cups int
	fmt.Println("Write how many ml of water you want to add:")
	fmt.Scan(&fill_water)
	fmt.Println("Write how many ml of milk you want to add:")
	fmt.Scan(&fill_milk )
	fmt.Println("Write how many grams of coffee beans you want to add:")
	fmt.Scan(&fill_coffee_beans)
	fmt.Println("Write how many disposable cups you want to add:")
	fmt.Scan(&fill_disposable_cups)

	cm.milk+=fill_milk
	cm.water+=fill_water
	cm.coffee_beans+=fill_coffee_beans
	cm.disposable_cups+= fill_disposable_cups
}
func(cm *coffee_machine)take() {
	fmt.Printf("I gave you $%v\n", cm.money)
	cm.money = 0
	
	

	
	
}


func main() {	
	machine1 := coffee_machine{400, 540, 120, 9,550}

	var choice string
	for choice != "exit"{
		fmt.Println("\nWrite action (buy, fill, take, remaining, exit):")
		fmt.Scan(&choice)
		switch choice {
		case "buy":
			machine1.buy()
		case "fill":
			machine1.fill()
		case "take":
			machine1.take()
		case "remaining":
			fmt.Printf("The coffee machine has:\n")
			fmt.Printf("%v ml of water\n", machine1.water)
			fmt.Printf("%v ml of milk\n", machine1.milk)
			fmt.Printf("%v g of coffee beans\n",machine1.coffee_beans)
			fmt.Printf("%v disposable cups\n", machine1.disposable_cups)
			fmt.Printf("$%v of money\n",machine1.money)
			
		case "exit":
			break

		}
	}


	
}
	