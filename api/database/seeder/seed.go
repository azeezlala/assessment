package seeder

import (
	"context"
	"github.com/azeezlala/assessment/api/database"
	"github.com/azeezlala/assessment/api/internal/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

func Seed() error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	db := database.GetDB()
	err := seedResources(ctx, db.DB)
	return err
}

func seedResources(ctx context.Context, db *gorm.DB) error {
	users := []model.Resources{
		{ID: uuid.NewString(), Name: "Lambda", Type: "function", Region: "us-west-1"},
		{ID: uuid.NewString(), Name: "Dynamic DB", Type: "database", Region: "us-west-1"},
		{ID: uuid.NewString(), Name: "Pubsub", Type: "event", Region: "us-central-1"},
	}
	return db.WithContext(ctx).Clauses(clause.OnConflict{DoNothing: true}).Create(&users).Error
}
