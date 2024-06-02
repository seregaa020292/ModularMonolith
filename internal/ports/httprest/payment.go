package httprest

import (
	"context"

	"github.com/seregaa020292/ModularMonolith/internal/infrastructure/openapi"
)

type PaymentHandler struct{}

func NewPaymentHandler() *PaymentHandler {
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
