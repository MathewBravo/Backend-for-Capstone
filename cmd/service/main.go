package main

import (
	"Backend_for_Capstone/api/database"
	routes "Backend_for_Capstone/api/router"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
)

func loadenv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}
}

func main() {
	loadenv()
	var err error
	database.EstablishConnection()

	router := gin.Default()

	routes.RouteFetch(router)

	err = router.Run("localhost:8080")
	if err != nil {
		fmt.Println("Error running router on port 8080")
	}
}
