package main

import "github.com/bruh-boys/api-bruh-bot/src/smsSpam"

// run this for test the program please
func main() {
	phone := "12341234123412341"
	smsSpam.SendRequest(phone)
}
