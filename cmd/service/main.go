package main

import (
	"Backend_for_Capstone/api/database"
	routes "Backend_for_Capstone/api/router"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
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

	router.Use(cors.Default())
	routes.RouteFetch(router)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	err = router.Run(":" + port)
	if err != nil {
		fmt.Println("Error running router on port 8080")
	}
}
