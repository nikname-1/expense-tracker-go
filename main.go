package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type expense struct {
	Name     string `json:"name"`
	Category string `json:"category"`
	Value    int    `json:"value"`
	Time     int64  `json:"time"`
}

type budget struct {
	Amount    int    `json:"amount"`
	Category  string `json:"category"`
	Remaining int    `json:"remaining"`
}

// Create objects in memory for the sake of this simple demo, in place of a database.
var budgets = map[string]budget{
	"Food":      {Amount: 400, Category: "Food", Remaining: 300},
	"Transport": {Amount: 200, Category: "Transport", Remaining: 200}}

var expenses = map[string]expense{
	"1": {Name: "McDonalds", Category: "Food", Value: 15, Time: time.Now().Unix()},
	"2": {Name: "Train", Category: "Transport", Value: 5, Time: time.Now().Unix()}}

func getExpenses(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, expenses)
}

func getBudgets(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, budgets)
}

func main() {
	router := gin.Default()
	router.GET("/expenses", getExpenses)
	router.GET("/budgets", getBudgets)
	router.Run("localhost:8080")

}
