package config

import "github.com/gin-contrib/cors"

var CorsConfig = cors.Config{
	AllowAllOrigins:  true,
	AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
	AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
	AllowCredentials: true,
}
