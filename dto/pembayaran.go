package dto

import (
	"errors"
	"mime/multipart"
)

var (
	ErrCreatePembayaran   = errors.New("failed to create payment")
	ErrPembayaranNotFound = errors.New("payment not found")
)

type (
	PembayaranRequest struct {
		Name          string                `json:"name"`
		AccountNumber string                `json:"account_number"`
		BankName      string                `json:"bank_name"`
		TotalPrice    int                   `json:"total_price"`
		PaymentPhoto  *multipart.FileHeader `json:"payment_photo"`
	}

	PembayaranResponse struct {
		ID            string `json:"id"`
		Name          string `json:"name"`
		AccountNumber string `json:"account_number"`
		BankName      string `json:"bank_name"`
		TotalPrice    int    `json:"total_price"`
		PaymentUrl    string `json:"payment_url"`
	}
)
