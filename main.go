package main

import (
	_ "primeskills-test-api/docs"
	"primeskills-test-api/internal/infrastructure"
)

//	@title			Primeskills Test API Documentation
//	@version		1.0
//	@description	Simple Todo App with JWT authentication

//	@host		localhost:3000
//	@BasePath	/api/v1
//	@schemes	http https

//	@accept		application/json
//	@produce	application/json

// @securityDefinitions.apikey	Bearer
// @in							header
// @name						Authorization
func main() {
	infrastructure.Run()
}
