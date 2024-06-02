package httprest

import (
	"github.com/seregaa020292/ModularMonolith/internal/infrastructure/openapi"
)

var _ openapi.StrictServerInterface = (*HttpRest)(nil)

type HttpRest struct{}
