package main

// =================================================================================================
// Problem
// =================================================================================================

// Let's say we have a function called `useProduct` that takes a product type as an argument and
// uses the product. The product type can be "A", "B", etc. We have different concrete products
// like `ConcreteProductA`, `ConcreteProductB`, etc. that implement the `useProduct`
// function. We can use a switch statement to check the product type and call the `useProduct` function
// on the concrete product. But this approach is not scalable. If we add more product types, we have to
// add more cases to the switch statement. This is a violation of the open/closed principle. We should
// be able to add new product types without modifying the existing code.

/*
type ConcreteProductA struct{}
func (p *ConcreteProductA) Use() {
	println("Using product A")
}

type ConcreteProductB struct{}
func (p *ConcreteProductB) Use() {
	println("Using product B")
}

func useProduct(productType string) {
 if productType == "A" {
  ConcreteProductA{}.Use()
 } else if productType == "B" {
  ConcreteProductB{}.Use()
 }
 // Add more product types here.
}
*/

// =================================================================================================
// Solution
// =================================================================================================

// We can use the factory method pattern to solve this problem. We can create a factory struct
// that has a method called `CreateProduct` that takes a product type as an argument and returns
// a product. The factory struct doesn't know about the concrete products. It only knows about the
// product interface. The concrete products implement the product interface. The `useProduct`
// function only knows about the factory and the product interface. It doesn't know about the
// concrete products. This way, we can add new product types without modifying the existing code.

// Product is the interface that all products should implement.
type Product interface {
	Use()
}

type ConcreteProductA struct{}

func (p *ConcreteProductA) Use() {
	println("Using product A")
}

type ConcreteProductB struct{}

func (p *ConcreteProductB) Use() {
	println("Using product B")
}

// ProductFactory is the struct that has a CreateProduct method.
type ProductFactory struct{}

// CreateProduct is the method that creates a product using the product type.
func (f *ProductFactory) CreateProduct(productType string) Product {
	if productType == "A" {
		return &ConcreteProductA{}
	} else if productType == "B" {
		return &ConcreteProductB{}
	}
	// Add more product types here.
	return nil
}

// =================================================================================================
// Application
// =================================================================================================

// This code only knows about the product factory and the product interface.
// It doesn't know about the concrete products like ConcreteProductA or ConcreteProductB.
func useProduct(productType string) {
	product := (&ProductFactory{}).CreateProduct(productType)
	product.Use()
}

func main() {
	useProduct("A")
	useProduct("B")
}

// Run the main function to see the output.
// go run main.go

// Output:
// Using product A
// Using product B
