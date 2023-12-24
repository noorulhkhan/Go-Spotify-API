package main

import (
	"fmt"
)

func init() {
	authorize()
}

// @contact.name   Noorul H. Khan
// @contact.url    http://www.swagger.io/support
// @contact.email  noorulhasan.khan@outlook.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	fmt.Println("Starting server ...")
	InitialMigration()
	InitializeRouter()
}
