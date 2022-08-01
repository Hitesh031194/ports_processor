package repository

//go:generate mockery --name=PortRepository

import (
	"context"
	dbModel "github.com/hiteshpattanayak-tw/ports_processor/internal/app/db/migrations/model"
)

type PortRepository interface {
	UpsertPort(ctx context.Context, port *dbModel.Port) error
}
