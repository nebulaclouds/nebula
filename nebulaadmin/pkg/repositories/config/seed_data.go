package config

import (
	"context"
	"fmt"

	"gorm.io/gorm"

	"github.com/nebulaclouds/nebula/nebulaadmin/pkg/repositories/models"
	"github.com/nebulaclouds/nebula/nebulastdlib/logger"
)

// Returns a function to seed the database with default values.
func SeedProjects(db *gorm.DB, projects []string) error {
	tx := db.Begin()
	for _, project := range projects {
		projectModel := models.Project{
			Identifier:  project,
			Name:        project,
			Description: fmt.Sprintf("%s description", project),
		}
		if err := tx.Where(models.Project{Identifier: project}).Omit("id").FirstOrCreate(&projectModel).Error; err != nil {
			logger.Warningf(context.Background(), "failed to save project [%s]", project)
			tx.Rollback()
			return err
		}
	}
	return tx.Commit().Error
}
