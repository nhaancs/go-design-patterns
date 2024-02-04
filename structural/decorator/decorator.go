package main

// =================================================================================================
// Problem
// =================================================================================================

// Suppose we have a system that supports multiple notification types, such as email, SMS, and Slack.
// We want to be able to send notifications to users using different notification types. We also want to
// be able to add new notification types easily.

// The following code is an example of how we might implement this system.

/*
// Notifier is the interface that all notifications should implement.
type Notifier interface {
	Notify(message string)
}

type EmailNotifier struct{}

func (n *EmailNotifier) Notify(message string) {
	println("Sending email notification: ", message)
}

type SMSNotifier struct{}

func (n *SMSNotifier) Notify(message string) {
	println("Sending SMS notification: ", message)
}

type SlackNotifier struct{}

func (n *SlackNotifier) Notify(message string) {
	println("Sending Slack notification: ", message)
}

// If I have more notifiers, I have to define more types to combine them!!!

type EmailSMSSlackNotifier struct {
	emailNotifier EmailNotifier
	smsNotifier   SMSNotifier
	slackNotifier SlackNotifier
}

func (notifier EmailSMSSlackNotifier) Send(message string) {
	notifier.emailNotifier.Send(message)
	notifier.smsNotifier.Send(message)
	notifier.slackNotifier.Send(message)
}

// Maybe you want EmailSMSNotifier, EmailSlackNotifier, SMSSlackNotifier, etc.

func sendNotification(notifier Notifier, message string) {
	notifier.Notify(message)
}

func main() {
	sendNotification(&EmailSMSSlackNotifier{}, "Hello, User!")
}

*/

// =================================================================================================
// Solution
// =================================================================================================

// The decorator pattern allows us to add new functionality to an object without altering its structure.
// We can use the decorator pattern to add new notification types to our system without altering the existing
// notification types.

// Notifier is the interface that all notifications and notification decorators should implement.
type Notifier interface {
	Notify(message string)
}

type EmailNotifier struct{}

func (n *EmailNotifier) Notify(message string) {
	println("Sending email notification: ", message)
}

type SMSNotifier struct{}

func (n *SMSNotifier) Notify(message string) {
	println("Sending SMS notification: ", message)
}

type SlackNotifier struct{}

func (d *SlackNotifier) Notify(message string) {
	println("Sending Slack notification: ", message)
}

// NotificationDecorator is a decorator for notifications.
type NotificationDecorator struct {
	core     Notifier
	notifier Notifier
}

func (nd NotificationDecorator) Notify(message string) {
	nd.notifier.Notify(message)

	if nd.core != nil {
		nd.core.Notify(message)
	}
}

// Decorate use value receiver to avoid circular reference.
func (nd NotificationDecorator) Decorate(notifier Notifier) NotificationDecorator {
	return NotificationDecorator{
		core:     nd,
		notifier: notifier,
	}
}

// =================================================================================================
// Application
// =================================================================================================

func sendNotification(notifier Notifier, message string) {
	notifier.Notify(message)
}

func main() {
	// Create a notification with decorators.
	decoratedNotifier := NotificationDecorator{notifier: &EmailNotifier{}}.
		Decorate(&SMSNotifier{}).
		Decorate(&SlackNotifier{})

	sendNotification(&decoratedNotifier, "Hello, User!")
}

// Run the main function to see the output.
// go run main.go

// Output:
// Sending email notification:  Hello, User!
// Sending SMS notification:  Hello, User!
// Sending Slack notification:  Hello, User!
