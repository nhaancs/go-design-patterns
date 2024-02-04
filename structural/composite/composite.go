package main

// =================================================================================================
// Problem
// =================================================================================================

// Suppose we have a system where we need to represent a hierarchy of boxes and real items. Each real item
// and box has a cost, and we want to calculate the total cost. We want to be able to treat boxes and real items uniformly.
// However, using traditional object-oriented design, we would have to create separate classes for Box and RealItem,
// and client code would have to treat them differently.

/*
type RealItem struct {
 name string
 cost int
}

func (ri *RealItem) Cost() int {
 return ri.cost
}

type Box struct {
 items []RealItem
}

func (b *Box) Cost() int {
 totalCost := 0
 for _, item := range b.items {
  totalCost += item.Cost()
 }
 return totalCost
}

// Client code would have to treat Box and RealItem differently.
func calculateCost(ri *RealItem, b *Box) int {
 return ri.Cost() + b.Cost()
}
*/

// =================================================================================================
// Solution
// =================================================================================================

// We can use the Composite design pattern to solve this problem. We can create an Item interface that
// declares the Cost method. Then we can create RealItem and Box structs that implement the Item
// interface. The Box struct can have a slice of Items.

// Item is the interface that all items should implement.
type Item interface {
	Cost() int
}

type RealItem struct {
	name string
	cost int
}

func (ri *RealItem) Cost() int {
	return ri.cost
}

type Box struct {
	items []Item
}

func (b *Box) Cost() int {
	totalCost := 0
	for _, item := range b.items {
		totalCost += item.Cost()
	}
	return totalCost
}

// =================================================================================================
// Application
// =================================================================================================

func main() {
	box := &Box{
		items: []Item{
			&RealItem{name: "Item 1", cost: 10},
			&RealItem{name: "Item 2", cost: 20},
			&Box{
				items: []Item{
					&RealItem{name: "Item 3", cost: 30},
					&RealItem{name: "Item 4", cost: 40},
				},
			},
		},
	}

	println("Total cost: ", box.Cost())
}

// Run the main function to see the output.
// go run main.go

// Output:
// Total cost:  100
