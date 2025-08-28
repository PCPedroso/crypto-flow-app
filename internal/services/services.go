package services

import (
	"crypto-flow-app/internal/models"
	"errors"
)

func ProcessTransaction(transaction models.Transaction) (string, error) {
	// Business logic for processing a cryptocurrency transaction
	if err := ValidateTransaction(transaction); err != nil {
		return "", err
	}
	// Process the transaction (e.g., save to database, call external APIs, etc.)
	return "Transaction processed successfully", nil
}

func ValidateTransaction(transaction models.Transaction) error {
	// Validate the transaction (e.g., check amounts, user existence, etc.)
	if transaction.Amount <= 0 {
		return errors.New("invalid transaction amount")
	}
	if transaction.UserID == "" {
		return errors.New("user ID cannot be empty")
	}
	// Additional validation logic can be added here
	return nil
}
