package main

import (
	"Project_2_cleanArchitecture/database"
	route "Project_2_cleanArchitecture/routes"

	"github.com/gin-gonic/gin"
)

const PORT = ":8080"

func main() {
	router := gin.Default()

	database.StartDB()
	db := database.GetDB()

	route.SetupBookRoute(router,db)
	router.Run(PORT)
}