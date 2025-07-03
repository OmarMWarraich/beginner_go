package exercise

import "fmt"

// item struct
type item struct {
	Name string
	Type string
}

// Player struct
type Player struct {
	Name string
	Inventory []item
}

// Pick up an item (modifies the player's inventory)
func (p *Player) PickUpItem(item item){
	// TODO: Implement function to add an item to inventory
	p.Inventory = append(p.Inventory, item)
	fmt.Printf("%s picked up %s!\n", p.Name, item.Name)
}

// Drop an item (removes from inventory)
func (p *Player) DropItem(itemName string){
	// TODO: Implement function to remove an item from inventory
	for i, item := range p.Inventory {
		if item.Name == itemName {
			p.Inventory = append(p.Inventory[:i], p.Inventory[i+1:]... ) // Remove Item
			fmt.Printf("%s dropped %s.\n", p.Name, item.Name)
		}
	}
	fmt.Printf("%s does not have %s in inventory.\n", p.Name, itemName)
}

// Use an item (if potion, remove it after use)
func (p *Player) UseItem(itemName string) {
	// TODO: Implement function to use an item
	for i, item := range p.Inventory {
		if item.Name == itemName {
			if item.Type == "potion" {
				fmt.Printf("%s used %s and feels rejuvenated!\n", p.Name, itemName)
				p.Inventory = append(p.Inventory[:i], p.Inventory[i+1:]... ) // Remove Item
				} else {
				fmt.Printf("%s used %s.\n", p.Name, item.Name)
			}
			return
		}
	}
	fmt.Printf("%s does not have %s in inventory.\n", p.Name, itemName)
}