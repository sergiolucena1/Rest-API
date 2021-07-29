package migrations

import (
	"github.com/sergiolucena1/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB)  {
	db.AutoMigrate(models.Product{})
}
