package snapshoter

import (
	"bytes"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/nebulaclouds/nebula/nebulaadmin/pkg/repositories/errors"
	"github.com/nebulaclouds/nebula/nebulaadmin/pkg/repositories/mocks"
	adminModels "github.com/nebulaclouds/nebula/nebulaadmin/pkg/repositories/models"
	repositoryInterfaces "github.com/nebulaclouds/nebula/nebulaadmin/scheduler/repositories/interfaces"
	schedMocks "github.com/nebulaclouds/nebula/nebulaadmin/scheduler/repositories/mocks"
	"github.com/nebulaclouds/nebula/nebulaadmin/scheduler/repositories/models"
	"github.com/nebulaclouds/nebula/nebulastdlib/promutils"
)

var (
	db repositoryInterfaces.SchedulerRepoInterface
)

func setupSnapShoter(scope string) Persistence {
	db = mocks.NewMockRepository()
	return New(promutils.NewScope(scope), db)
}

func TestSnapShoterRead(t *testing.T) {

	t.Run("successful read", func(t *testing.T) {
		snapshoter := setupSnapShoter("TestSnapShoterReadSuccessfulRead")
		var bytesArray []byte
		f := bytes.NewBuffer(bytesArray)
		writer := VersionedSnapshot{}
		snapshot := &SnapshotV1{
			LastTimes: map[string]*time.Time{},
		}
		currTime := time.Now()
		snapshot.LastTimes["schedule1"] = &currTime
		err := writer.WriteSnapshot(f, snapshot)
		assert.Nil(t, err)

		snapshotRepo := db.ScheduleEntitiesSnapshotRepo().(*schedMocks.ScheduleEntitiesSnapShotRepoInterface)
		snapshotModel := models.ScheduleEntitiesSnapshot{
			BaseModel: adminModels.BaseModel{
				ID:        17,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			Snapshot: f.Bytes(),
		}
		snapshotRepo.OnRead(context.Background()).Return(snapshotModel, nil)

		reader := &VersionedSnapshot{}
		snapshotVal, err := snapshoter.Read(context.Background(), reader)
		assert.Nil(t, err)
		assert.NotNil(t, snapshotVal)
	})

	t.Run("unsuccessful read ignore error", func(t *testing.T) {
		snapshoter := setupSnapShoter("TestSnapShoterReadUnsuccessfulReadIgnoreError")
		snapshotRepo := db.ScheduleEntitiesSnapshotRepo().(*schedMocks.ScheduleEntitiesSnapShotRepoInterface)

		snapshotRepo.OnRead(context.Background()).Return(models.ScheduleEntitiesSnapshot{}, errors.GetSingletonMissingEntityError("schedule_entities_snapshots"))

		reader := &VersionedSnapshot{}
		snapshotVal, err := snapshoter.Read(context.Background(), reader)
		assert.Nil(t, err)
		assert.NotNil(t, snapshotVal)
	})

	t.Run("unsuccessful read dont ignore error", func(t *testing.T) {
		snapshoter := setupSnapShoter("TestSnapShoterReadUnsuccessfulReadDontIgnoreError")
		snapshotRepo := db.ScheduleEntitiesSnapshotRepo().(*schedMocks.ScheduleEntitiesSnapShotRepoInterface)

		snapshotRepo.OnRead(context.Background()).Return(models.ScheduleEntitiesSnapshot{}, errors.GetInvalidInputError("invalid input"))

		reader := &VersionedSnapshot{}
		_, err := snapshoter.Read(context.Background(), reader)
		assert.NotNil(t, err)
	})
}

func TestSnapShoterSave(t *testing.T) {
	snapshoter := setupSnapShoter("TestSnapShoterSave")
	writer := &VersionedSnapshot{}
	var bytesArray []byte
	f := bytes.NewBuffer(bytesArray)
	snapshot := &SnapshotV1{
		LastTimes: map[string]*time.Time{},
	}
	currTime := time.Now()
	snapshot.LastTimes["schedule1"] = &currTime
	err := writer.WriteSnapshot(f, snapshot)
	assert.Nil(t, err)

	snapshotRepo := db.ScheduleEntitiesSnapshotRepo().(*schedMocks.ScheduleEntitiesSnapShotRepoInterface)
	snapshotModel := models.ScheduleEntitiesSnapshot{
		BaseModel: adminModels.BaseModel{
			ID:        0,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
		},
		Snapshot: f.Bytes(),
	}
	snapshotRepo.OnWrite(context.Background(), snapshotModel).Return(nil)

	snapshoter.Save(context.Background(), writer, snapshot)
}
