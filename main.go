package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type expense struct {
	ID 	string `json:"id"`
	Name 	string `json:"name"`
	Type	string `json:"type"`
	Value	string `json:"value"`
}




func main() {
	router := gin.Default()

}


