package main

import (
	"fmt"

	"github.com/dghubble/go-twitter/twitter"
)

func getUserData(client *twitter.Client) *twitter.User {
	// Verify Credentials
	verifyParams := &twitter.AccountVerifyParams{
		SkipStatus:   twitter.Bool(true),
		IncludeEmail: twitter.Bool(true),
	}
	user, _, _ := client.Accounts.VerifyCredentials(verifyParams)
	return user
}
func printTwitterUserData(user *twitter.User) {
	fmt.Print("username: ")
	c.Cyan(user.Name + " @" + user.ScreenName)
	if user.Email != "" {
		fmt.Print("Email ")
		c.Red(user.Email)
	}
	if user.Location != "" {
		fmt.Print("Location: ")
		c.Red(user.Location)
	}
	fmt.Print("user created on: ")
	c.Cyan(user.CreatedAt)
}
