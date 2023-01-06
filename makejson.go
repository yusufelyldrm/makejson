package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {

	/*
		var s string
		fmt.Print("Enter your name: ")
		fmt.Scanf("%s", &s)*/

	//Get name and address information from user
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter a name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)
	fmt.Println("Enter address: ")
	address, _ := reader.ReadString('\n')
	address = strings.TrimSpace(address)

	//Add entries to a map
	data := map[string]string{
		"name":    name,
		"address": address,
	}
	// convert Map to JSON
	jsonData, err := json.Marshal(data)

	//error checking
	if err != nil {
		fmt.Println(err)
		return
	}

	//Print JSON
	fmt.Println(string(jsonData))

}
