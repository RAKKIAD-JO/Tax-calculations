package main

import (
	. "test-backend/handlres"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/tax/calculations", Taxhandler)
	r.Run(":8080")

}
