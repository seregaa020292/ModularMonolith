package report

import (
	"github.com/google/wire"

	"github.com/seregaa020292/ModularMonolith/internal/report/repository"
)

var ModuleSet = wire.NewSet(
	repository.NewReportRepo,
)
