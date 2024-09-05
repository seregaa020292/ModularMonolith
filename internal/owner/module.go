package owner

import (
	"github.com/google/wire"

	"github.com/seregaa020292/ModularMonolith/internal/owner/repository"
)

var Module = wire.NewSet(
	repository.NewOwnerRepo,
)
