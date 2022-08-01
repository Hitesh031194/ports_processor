package repository

import (
	"github.com/hiteshpattanayak-tw/ports_processor/internal/app/config"
	dbModel "github.com/hiteshpattanayak-tw/ports_processor/internal/app/db/migrations/model"
	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type portDbOperator struct {
	portDb *gorm.DB
	appConfig  config.AppConfig
}

func ProvidePortRepository(db *gorm.DB, appConfig config.AppConfig) PortRepository {
	return &portDbOperator{portDb: db, appConfig: appConfig}
}

func (pr *portDbOperator) UpsertPort(ctx context.Context, port *dbModel.Port) error {
	err := pr.portDb.WithContext(ctx).Table(pr.appConfig.GetDatabaseConfig().PortsTableName).Where(dbModel.Port{Name: port.Name}).Save(port).Error
	if err != nil {
		return err
	}

	return nil
}
