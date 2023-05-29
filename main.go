package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type QueryResultStruct struct {
	ID           string `json:"id"`
	Manufacturer string `json:"manufacturer"`
	Model        string `json:"model"`
	Count        int    `json:"count"`
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.POST("/query", func(c *gin.Context) {
		result := []QueryResultStruct{}
		result = append(result, QueryResultStruct{
			ID:           "4502-3484-ccid",
			Manufacturer: "Fortinet",
			Model:        "FortiGate 60E",
			Count:        10,
		})
		result = append(result, QueryResultStruct{
			ID:           "4502-3484-kims",
			Manufacturer: "Fortinet",
			Model:        "FortiGate 80E",
			Count:        11,
		})
		result = append(result, QueryResultStruct{
			ID:           "4502-3484-fd0a",
			Manufacturer: "Fortinet",
			Model:        "FortiGate 100E",
			Count:        4,
		})
		c.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "ok",
			"data":    result,
		})
	})
	r.Run()
}
