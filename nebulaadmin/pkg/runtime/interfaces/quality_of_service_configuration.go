package interfaces

import (
	"github.com/nebulaclouds/nebula/nebulaidl/gen/pb-go/nebulaidl/core"
	"github.com/nebulaclouds/nebula/nebulastdlib/config"
)

type TierName = string

// Just incrementally start using mockery, replace with -all when working on https://github.com/nebulaclouds/nebula/issues/149
//go:generate mockery -name QualityOfServiceConfiguration -output=mocks -case=underscore

type QualityOfServiceSpec struct {
	QueueingBudget config.Duration `json:"queueingBudget"`
}

type QualityOfServiceConfig struct {
	TierExecutionValues map[TierName]QualityOfServiceSpec `json:"tierExecutionValues"`
	DefaultTiers        map[DomainName]TierName           `json:"defaultTiers"`
}

type QualityOfServiceConfiguration interface {
	GetTierExecutionValues() map[core.QualityOfService_Tier]core.QualityOfServiceSpec
	GetDefaultTiers() map[DomainName]core.QualityOfService_Tier
}
