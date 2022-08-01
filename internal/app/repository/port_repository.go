package repository

//go:generate mockery --name=PortRepository

import (
	"context"
	"github.com/hiteshpattanayak-tw/ports_processor/internal/app/config"
	dbModel "github.com/hiteshpattanayak-tw/ports_processor/internal/app/db/migrations/model"
	"gorm.io/gorm"
)

type PortRepository interface {
	UpsertPort(ctx context.Context, port *dbModel.Port) error
}

type portRepository struct {
	portDb *gorm.DB
	appConfig  config.AppConfig
}

func ProvidePortRepository(db *gorm.DB, appConfig config.AppConfig) PortRepository {
	return &portRepository{portDb: db, appConfig: appConfig}
}

func (pr *portRepository) UpsertPort(ctx context.Context, port *dbModel.Port) error {
	err := pr.portDb.WithContext(ctx).Table(pr.appConfig.GetDatabaseConfig().PortsTableName).Where(dbModel.Port{Name: port.Name}).Save(port).Error
	if err != nil {
		return err
	}

	return nil
}
