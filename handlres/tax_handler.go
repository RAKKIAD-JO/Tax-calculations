package handlres

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"test-backend/models"
	"test-backend/service"
	"test-backend/utils"

	"github.com/gin-gonic/gin"
)

type TaxResponse struct {
	Message string `json:"message"`
	// Tax    float64 `json:"tax,omitempty"`
}

func Taxhandler(c *gin.Context) {
	var req models.TaxRequest
	// Read request body for logging and then reset it so ShouldBindJSON works
	bodyBytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Printf("failed to read request body: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to read request body"})
		return
	}
	log.Printf("Request Body: %s", string(bodyBytes))
	// reset the request body so gin can re-read it during binding
	c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON format"})
		return
	}

	if err := utils.ValidateTextRequest(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	calculate := service.CalculateTax(req)
	log.Printf("Calculated Tax: %+v", calculate)
	tax := calculate.Tax
	taxLevels := calculate.TaxLevels

	// log.Printf("Final Tax: %f, Tax Levels: %+v", tax, taxLevels)
	c.JSON(http.StatusOK, gin.H{
		"tax":       tax,
		"taxLevels": taxLevels,
	})
}
