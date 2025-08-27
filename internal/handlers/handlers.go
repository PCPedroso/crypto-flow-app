package handlers

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

// HandleTransaction processes cryptocurrency transactions
func HandleTransaction(c *gin.Context) {
    // Logic for processing the transaction will go here
    c.JSON(http.StatusOK, gin.H{"message": "Transaction processed"})
}

// HandleHealthCheck checks the health of the application
func HandleHealthCheck(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"status": "healthy"})
}