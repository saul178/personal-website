package api

import "github.com/saul178/personal-website/internal/server"

func main() {
	routes := server.NewServer()
	server.SetupRoutes(routes)
	routes.Run()
}
