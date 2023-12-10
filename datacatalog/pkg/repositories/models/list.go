package models

import "github.com/nebulaclouds/nebula/datacatalog/pkg/common"

// Inputs to specify to list models
type ListModelsInput struct {
	// The filters for the list
	ModelFilters []ModelFilter
	// The number of models to list
	Limit int
	// The token to offset results by
	Offset int
	// Parameter to sort by
	SortParameter SortParameter
}

type SortParameter interface {
	GetDBOrderExpression(tableName string) string
}

// Generates db filter expressions for model values
type ModelValueFilter interface {
	GetDBQueryExpression(tableName string) (DBQueryExpr, error)
}

// Generates the join expressions for filters that require other entities
type ModelJoinCondition interface {
	GetJoinOnDBQueryExpression(sourceTableName string, joiningTableName string, joiningTableAlias string) (string, error)
}

// A single filter for a model encompasses value filters and optionally a join condition if the filter is not on
// the source model
type ModelFilter struct {
	ValueFilters  []ModelValueFilter
	JoinCondition ModelJoinCondition
	Entity        common.Entity
}

// Encapsulates the query and necessary arguments to issue a DB query.
type DBQueryExpr struct {
	Query string
	Args  interface{}
}
