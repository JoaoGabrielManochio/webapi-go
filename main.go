package main

import (
	"github.com/JoaoGabrielManochio/webapi-go/server"
)

// @title API em GO
// @version openapi: 3.0.n
// @termsOfService http://swagger.io/terms/
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath http://localhost:8080/api/v1/
func main() {

	server := server.NewServer()

	server.Run()
}
