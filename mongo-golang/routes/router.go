package routes

import (
	"log"

	"github.com/FlerioEU/hello-world/mongo-golang/db"
	"github.com/gin-gonic/gin"
)

type Routes struct {
	router *gin.Engine
}

func New(db db.Database) Routes {
	r := Routes{
		router: gin.Default(),
	}

	api := r.router.Group("/api/v1")
	// fe := r.router.Group("/")

	r.registerBooks(api, *db.DB)

	routes := r.router.Routes()

	log.Println("Endpoints:  ")
	for _, route := range routes {
		log.Printf("%v:\t %v\n", route.Method, route.Path)
	}

	return r
}

func (r Routes) Run(addr string) {
	log.Printf("Listening on port %v", addr)
	r.router.Run(addr)
}
