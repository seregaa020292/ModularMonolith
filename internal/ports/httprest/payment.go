package httprest

import (
	"context"

	"github.com/seregaa020292/ModularMonolith/internal/infrastructure/openapi"
	"github.com/seregaa020292/ModularMonolith/internal/payment/repository"
)

type PaymentHandler struct{}

func NewPaymentHandler(repo *repository.PaymentRepo) *PaymentHandler {
	return &PaymentHandler{}
}

func (h PaymentHandler) CreatePayment(ctx context.Context, request openapi.CreatePaymentRequestObject) (openapi.CreatePaymentResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (h PaymentHandler) ListPayments(ctx context.Context, request openapi.ListPaymentsRequestObject) (openapi.ListPaymentsResponseObject, error) {
	//TODO implement me
	panic("implement me")
}
