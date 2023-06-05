package services

import (
	"awesomeProject/internal/entities"
)

type ReceiptService struct {
	receiptRepo ReceiptRepostiory
}

func (e *ReceiptService) CreateReceipt() (*entities.Receipt, error) {
	// todo: validation を行う

	receipt := &entities.Receipt{}
	receiptRepo.CreateReceipt(receipt)
	return receipt, nil
}
