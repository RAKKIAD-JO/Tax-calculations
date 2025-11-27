package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin, gin.Default()
	r.post("/tax/calculations", Taxhandler)
	r.Run(":8080")

}
