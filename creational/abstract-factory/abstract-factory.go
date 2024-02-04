package main

// =================================================================================================
// Problem
// =================================================================================================

// Suppose we have a system that supports multiple themes. Each theme has different types of UI elements,
// such as buttons, checkboxes, and menus. We want to be able to switch between themes easily and ensure
// that all UI elements in a theme are consistent.

/*
// Button is the interface that all buttons should implement.
type Button interface {
	Click()
}

// Checkbox is the interface that all checkboxes should implement.
type Checkbox interface {
	Check()
}

// Dark theme UI elements.

type DarkButton struct{}
func (b *DarkButton) Click() {
	println("Dark button clicked")
}

type DarkCheckbox struct{}
func (c *DarkCheckbox) Check() {
	println("Dark checkbox checked")
}

// Light theme UI elements.

type LightButton struct{}
func (b *LightButton) Click() {
	println("Light button clicked")
}

type LightCheckbox struct{}
func (c *LightCheckbox) Check() {
	println("Light checkbox checked")
}

// Theme represents a theme that has different types of UI elements.
type Theme struct {
	Button
	Checkbox
}

// createThemeElements creates UI elements for a theme.
func createThemeElements(name string) *Theme {
	if name == "dark" {
		// This is OK
		return &Theme{&DarkButton{}, &DarkCheckbox{}}
	} else if name == "light" {
		return &Theme{
			&LightButton{},
			&DarkCheckbox{}, // Opps! This is not OK
		}
	}
	return nil
}
*/

// =================================================================================================
// Solution
// =================================================================================================

// We can use the Abstract Factory design pattern to solve this problem. We can define an AbstractFactory
// interface that declares methods for creating each type of UI element. Then we can implement a concrete
// factory for each theme. Each concrete factory will create UI elements for its specific theme.

// We can then create a Theme struct that has fields for each type of UI element. We can use the AbstractFactory
// interface to create UI elements for a theme. This way, we can ensure that all UI elements in a theme are
// consistent. We can also easily switch between themes by using the appropriate concrete factory.

// Button is the interface that all buttons should implement.
type Button interface {
	Click()
}

// Checkbox is the interface that all checkboxes should implement.
type Checkbox interface {
	Check()
}

// Dark theme UI elements.

type DarkButton struct{}

func (b *DarkButton) Click() {
	println("Dark button clicked")
}

type DarkCheckbox struct{}

func (c *DarkCheckbox) Check() {
	println("Dark checkbox checked")
}

// Light theme UI elements.

type LightButton struct{}

func (b *LightButton) Click() {
	println("Light button clicked")
}

type LightCheckbox struct{}

func (c *LightCheckbox) Check() {
	println("Light checkbox checked")
}

// UIElementFactory is the abstract factory interface that declares methods for creating UI elements.
type UIElementFactory interface {
	CreateButton() Button
	CreateCheckbox() Checkbox
}

// DarkThemeFactory is a concrete factory that creates UI elements for the dark theme.
type DarkThemeFactory struct{}

func (f *DarkThemeFactory) CreateButton() Button {
	return &DarkButton{}
}
func (f *DarkThemeFactory) CreateCheckbox() Checkbox {
	return &DarkCheckbox{}
}

// LightThemeFactory is a concrete factory that creates UI elements for the light theme.
type LightThemeFactory struct{}

func (f *LightThemeFactory) CreateButton() Button {
	return &LightButton{}
}
func (f *LightThemeFactory) CreateCheckbox() Checkbox {
	return &LightCheckbox{}
}

// =================================================================================================
// Application
// =================================================================================================

// Theme represents a theme that has different types of UI elements.
type Theme struct {
	Button
	Checkbox
}

func createThemeElements(factory UIElementFactory) *Theme {
	return &Theme{
		Button:   factory.CreateButton(),
		Checkbox: factory.CreateCheckbox(),
	}
}

func main() {
	// Create a dark theme factory and use it to create UI elements.
	darkFactory := &DarkThemeFactory{}
	darkTheme := createThemeElements(darkFactory)
	darkTheme.Button.Click()
	darkTheme.Checkbox.Check()

	// Create a light theme factory and use it to create UI elements.
	lightFactory := &LightThemeFactory{}
	lightTheme := createThemeElements(lightFactory)
	lightTheme.Button.Click()
	lightTheme.Checkbox.Check()
}

// Run the main function to see the output.
// go run main.go

// Output:
// Dark button clicked
// Dark checkbox checked
// Light button clicked
// Light checkbox checked
