package common

import (
	"fmt"

	"google.golang.org/grpc/codes"
	"k8s.io/apimachinery/pkg/util/sets"

	"github.com/nebulaclouds/nebula/nebulaadmin/pkg/errors"
	"github.com/nebulaclouds/nebula/nebulaidl/gen/pb-go/nebulaidl/admin"
)

const gormDescending = "%s desc"
const gormAscending = "%s asc"

type SortParameter interface {
	GetGormOrderExpr() string
}

type sortParamImpl struct {
	gormOrderExpression string
}

func (s *sortParamImpl) GetGormOrderExpr() string {
	return s.gormOrderExpression
}

func NewSortParameter(sort *admin.Sort, allowed sets.String) (SortParameter, error) {
	if sort == nil {
		return nil, nil
	}

	key := sort.Key
	if !allowed.Has(key) {
		return nil, errors.NewNebulaAdminErrorf(codes.InvalidArgument, "invalid sort key '%s'", key)
	}

	var gormOrderExpression string
	switch sort.Direction {
	case admin.Sort_DESCENDING:
		gormOrderExpression = fmt.Sprintf(gormDescending, key)
	case admin.Sort_ASCENDING:
		gormOrderExpression = fmt.Sprintf(gormAscending, key)
	default:
		return nil, errors.NewNebulaAdminErrorf(codes.InvalidArgument, "invalid sort order specified: %v", sort)
	}
	return &sortParamImpl{
		gormOrderExpression: gormOrderExpression,
	}, nil
}
