package testing

import (
	"time"

	"github.com/nebulaclouds/nebula/nebulaplugins/go/tasks/config"
	nebulastdconfig "github.com/nebulaclouds/nebula/nebulastdlib/config"
)

//go:generate pflags Config --default-var=defaultConfig

var (
	defaultConfig = Config{
		SleepDuration: nebulastdconfig.Duration{Duration: 0 * time.Second},
	}

	ConfigSection = config.MustRegisterSubSection(echoTaskType, &defaultConfig)
)

type Config struct {
	// SleepDuration indicates the amount of time before transitioning to success
	SleepDuration nebulastdconfig.Duration `json:"sleep-duration" pflag:",Indicates the amount of time before transitioning to success"`
}
