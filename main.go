package main

import (
	"fmt"
)

const version = "0.1-dev"

func main() {
	c.Green("twitterDM-to-matrix")
	c.Yellow("version: " + version)
	c.Green("bridge twitter DM to Matrix")
	fmt.Println("configuring")

	c.Purple("------matrix------")
	readMatrixConfig()
	loginMatrix()
	fmt.Print("matrix token: ")
	c.Cyan(matrixToken.AccessToken)
	fmt.Println("")

	c.Purple("------twitter------")
	client := readTwitterConfigTokensAndConnect()
	user := getUserData(client)
	printTwitterUserData(user)
	fmt.Println("")

	c.Purple("------starting to listen------")

	stream(client)
}
