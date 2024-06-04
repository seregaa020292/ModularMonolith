package owner

import (
	"github.com/google/wire"

	"github.com/seregaa020292/ModularMonolith/internal/owner/repository"
)

var ModuleSet = wire.NewSet(
	repository.NewOwnerRepo,
)
