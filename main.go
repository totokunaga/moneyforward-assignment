package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/mfx/util/types"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Fprintln(os.Stderr, "Error: User ID must be given as the first argument")
		os.Exit(1)
	}

	userId, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: User ID must be integer")
		os.Exit(1)
	}

	var user types.User
	err = user.LoadData(userId)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	fmt.Println(fmt.Sprintf("Name: %s", user.Name))
	for _, account := range user.Accounts {
		fmt.Println(fmt.Sprintf("- %s: %d", account.Name, account.Balance))
	}
}
