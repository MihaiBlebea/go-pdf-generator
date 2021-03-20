package main

import (
	"fmt"

	"github.com/MihaiBlebea/go-diploma/cmd"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load("./.env")
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	cmd.Execute()
}
