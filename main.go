package main

import (
	"github.com/Shanedell/ga-api-go/api"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/personTimes", api.GetAllPersonTimes)
	router.POST("/personTimes", api.CreatePersonTime)
	router.GET("/personTimes/:id", api.FindPersonsTime)

	router.Run("localhost:8080")
	// router := mux.NewRouter()
	// router.HandleFunc("/personTimes", api.GetPersonTimes).Methods("GET")
	// log.Fatal(http.ListenAndServe(":8080", router))
}
