package main

// =================================================================================================
// Problem
// =================================================================================================

// We have a `Payment` struct that has a `ProcessPayment` method. The `ProcessPayment` method is responsible
// for making the payment. The `ProcessPayment` method is implemented in a way that it checks the payment type
// and then makes the payment.
// The problem is that the `ProcessPayment` method is getting bigger and bigger as we add more payment types.
// We need to find a way to make the `ProcessPayment` method smaller and more maintainable.

/*
type Payment struct {
	paymentType string
}

func (p *Payment) Pay(amount int64) {
	if p.paymentType == "credit_card" {
		println("Paying with credit card: ", amount)
	} else if p.paymentType == "paypal" {
		println("Paying with PayPal: ", amount)
	}
	// Add more payment types here.
}
*/

// =================================================================================================
// Solution
// =================================================================================================

// We can use the strategy pattern to solve this problem. We can create  a `PaymentStrategy` interface
// that has a `ProcessPayment` method. Then we can create different payment strategies that implement the
// `PaymentStrategy` interface. Then we can pass the payment strategy to the `Payment` struct and call the
// `ProcessPayment` method on the payment strategy.

// PaymentStrategy is the interface that all payment strategies should implement.
type PaymentStrategy interface {
	ProcessPayment(amount int64)
}

// CreditCardStrategy is a concrete payment strategy using a credit card.
type CreditCardStrategy struct{}

func (c *CreditCardStrategy) ProcessPayment(amount int64) {
	println("Paying with credit card: ", amount)
}

// PayPalStrategy is a concrete payment strategy using PayPal.
type PayPalStrategy struct{}

func (p *PayPalStrategy) ProcessPayment(amount int64) {
	println("Paying with PayPal: ", amount)
}

// Payment is the struct that has a payment strategy.
type Payment struct {
	strategy PaymentStrategy
}

// NewPayment is a factory function that creates a new Payment.
func NewPayment(strategy PaymentStrategy) *Payment {
	return &Payment{strategy: strategy}
}

// Pay is the method that makes the payment using the payment strategy.
func (p *Payment) Pay(amount int64) {
	p.strategy.ProcessPayment(amount)
}

// =================================================================================================
// Application
// =================================================================================================

func main() {
	// Create a new payment with a credit card strategy.
	payment := NewPayment(&CreditCardStrategy{})
	// Make the payment.
	payment.Pay(1000)

	// Create a new payment with a PayPal strategy.
	payment = NewPayment(&PayPalStrategy{})
	// Make the payment.
	payment.Pay(10000)
}

// Run the main function to see the output.
// go run main.go

// Output:
// Paying with credit card: 1000
// Paying with PayPal: 10000
