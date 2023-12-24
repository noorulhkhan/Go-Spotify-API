package main

import (
	"fmt"
)

func init() {
	authorize()
}

func main() {
	fmt.Println("Starting server ...")
	InitialMigration()
	InitializeRouter()
}
