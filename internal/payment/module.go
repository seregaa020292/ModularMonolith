package payment

import (
	"github.com/google/wire"

	"github.com/seregaa020292/ModularMonolith/internal/payment/repository"
)

var ModuleSet = wire.NewSet(
	repository.NewPaymentRepo,
)
