package fine

import (
	"github.com/google/wire"

	"github.com/seregaa020292/ModularMonolith/internal/fine/query"
	"github.com/seregaa020292/ModularMonolith/internal/fine/repository"
)

var ModuleSet = wire.NewSet(
	repository.NewFineRepo,
	wire.Bind(new(query.Repository), new(*repository.FineRepo)),
	query.NewGetList,
)
