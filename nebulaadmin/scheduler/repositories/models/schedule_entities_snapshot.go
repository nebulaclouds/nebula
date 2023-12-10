package models

import "github.com/nebulaclouds/nebula/nebulaadmin/pkg/repositories/models"

// Database model to save the snapshot for the schedulable entities in the db
type ScheduleEntitiesSnapshot struct {
	models.BaseModel
	Snapshot []byte `gorm:"column:snapshot" schema:"-"`
}

type ScheduleEntitiesSnapshotCollectionOutput struct {
	Snapshots []ScheduleEntitiesSnapshot
}
