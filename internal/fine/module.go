package fine

import (
	"github.com/google/wire"

	"github.com/seregaa020292/ModularMonolith/internal/fine/repository"
)

var ModuleSet = wire.NewSet(
	repository.NewFineRepo,
)
