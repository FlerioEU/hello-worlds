package routes

import (
	"log"

	"github.com/FlerioEU/hello-world/mongo-golang/db"
	"github.com/gin-gonic/gin"
)

type Routes struct {
	Router *gin.Engine
}

func New(db db.Database) Routes {
	r := Routes{
		Router: gin.Default(),
	}

	api := r.Router.Group("/api/v1")
	// fe := r.router.Group("/")

	r.registerBooks(api, *db.DB)

	routes := r.Router.Routes()

	log.Println("Endpoints:  ")
	for _, route := range routes {
		log.Printf("%v:\t %v\n", route.Method, route.Path)
	}

	return r
}

func (r Routes) Run(addr string) {
	log.Printf("Listening on port %v", addr)
	r.Router.Run(addr)
}
