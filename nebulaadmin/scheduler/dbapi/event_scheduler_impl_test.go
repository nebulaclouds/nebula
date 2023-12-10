package dbapi

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/nebulaclouds/nebula/nebulaadmin/pkg/async/schedule/interfaces"
	repositoryInterfaces "github.com/nebulaclouds/nebula/nebulaadmin/pkg/repositories/interfaces"
	"github.com/nebulaclouds/nebula/nebulaadmin/pkg/repositories/mocks"
	schedMocks "github.com/nebulaclouds/nebula/nebulaadmin/scheduler/repositories/mocks"
	"github.com/nebulaclouds/nebula/nebulaidl/gen/pb-go/nebulaidl/admin"
	"github.com/nebulaclouds/nebula/nebulaidl/gen/pb-go/nebulaidl/core"
)

var (
	db repositoryInterfaces.Repository
)

func setupEventScheduler() interfaces.EventScheduler {
	db = mocks.NewMockRepository()
	return New(db)
}

func TestCreateScheduleInput(t *testing.T) {
	eventScheduler := setupEventScheduler()
	schedule := &admin.Schedule{
		ScheduleExpression: &admin.Schedule_CronSchedule{
			CronSchedule: &admin.CronSchedule{
				Schedule: "*/1 * * * *",
			},
		},
		KickoffTimeInputArg: "kickoff_time",
	}
	addScheduleInput, err := eventScheduler.CreateScheduleInput(context.Background(), nil, core.Identifier{
		Project: "project",
		Domain:  "domain",
		Name:    "scheduled_wroflow",
		Version: "v1",
	}, schedule)
	assert.Nil(t, err)
	assert.NotNil(t, addScheduleInput)
}

func TestRemoveSchedule(t *testing.T) {
	eventScheduler := setupEventScheduler()

	scheduleEntitiesRepo := db.SchedulableEntityRepo().(*schedMocks.SchedulableEntityRepoInterface)
	scheduleEntitiesRepo.OnDeactivateMatch(mock.Anything, mock.Anything).Return(nil)

	err := eventScheduler.RemoveSchedule(context.Background(), interfaces.RemoveScheduleInput{
		Identifier: core.Identifier{
			Project: "project",
			Domain:  "domain",
			Name:    "scheduled_wroflow",
			Version: "v1",
		},
	})
	assert.Nil(t, err)
}

func TestAddSchedule(t *testing.T) {
	t.Run("schedule_rate", func(t *testing.T) {
		eventScheduler := setupEventScheduler()
		schedule := admin.Schedule{
			ScheduleExpression: &admin.Schedule_Rate{
				Rate: &admin.FixedRate{
					Value: 1,
					Unit:  admin.FixedRateUnit_MINUTE,
				},
			},
			KickoffTimeInputArg: "kickoff_time",
		}

		scheduleEntitiesRepo := db.SchedulableEntityRepo().(*schedMocks.SchedulableEntityRepoInterface)
		scheduleEntitiesRepo.OnActivateMatch(mock.Anything, mock.Anything).Return(nil)

		err := eventScheduler.AddSchedule(context.Background(), interfaces.AddScheduleInput{
			Identifier: core.Identifier{
				Project: "project",
				Domain:  "domain",
				Name:    "scheduled_wroflow",
				Version: "v1",
			},
			ScheduleExpression: schedule,
		})
		assert.Nil(t, err)
	})

	t.Run("cron_schedule", func(t *testing.T) {
		eventScheduler := setupEventScheduler()
		schedule := admin.Schedule{
			ScheduleExpression: &admin.Schedule_CronSchedule{
				CronSchedule: &admin.CronSchedule{
					Schedule: "*/1 * * * *",
				},
			},
			KickoffTimeInputArg: "kickoff_time",
		}

		scheduleEntitiesRepo := db.SchedulableEntityRepo().(*schedMocks.SchedulableEntityRepoInterface)
		scheduleEntitiesRepo.OnActivateMatch(mock.Anything, mock.Anything).Return(nil)

		err := eventScheduler.AddSchedule(context.Background(), interfaces.AddScheduleInput{
			Identifier: core.Identifier{
				Project: "project",
				Domain:  "domain",
				Name:    "scheduled_wroflow",
				Version: "v1",
			},
			ScheduleExpression: schedule,
		})
		assert.Nil(t, err)
	})

	t.Run("cron_expression_unsupported", func(t *testing.T) {
		eventScheduler := setupEventScheduler()
		schedule := admin.Schedule{
			ScheduleExpression: &admin.Schedule_CronExpression{
				CronExpression: "* */1 * * * *",
			},
			KickoffTimeInputArg: "kickoff_time",
		}

		scheduleEntitiesRepo := db.SchedulableEntityRepo().(*schedMocks.SchedulableEntityRepoInterface)
		scheduleEntitiesRepo.OnActivateMatch(mock.Anything, mock.Anything).Return(nil)

		err := eventScheduler.AddSchedule(context.Background(), interfaces.AddScheduleInput{
			Identifier: core.Identifier{
				Project: "project",
				Domain:  "domain",
				Name:    "scheduled_wroflow",
				Version: "v1",
			},
			ScheduleExpression: schedule,
		})
		assert.NotNil(t, err)
	})
}
