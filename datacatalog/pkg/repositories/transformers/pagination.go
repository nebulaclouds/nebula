package transformers

import (
	"strconv"
	"strings"

	"google.golang.org/grpc/codes"

	"github.com/nebulaclouds/nebula/datacatalog/pkg/common"
	"github.com/nebulaclouds/nebula/datacatalog/pkg/errors"
	"github.com/nebulaclouds/nebula/datacatalog/pkg/repositories/gormimpl"
	"github.com/nebulaclouds/nebula/datacatalog/pkg/repositories/models"
	"github.com/nebulaclouds/nebula/nebulaidl/gen/pb-go/nebulaidl/datacatalog"
)

func ApplyPagination(paginationOpts *datacatalog.PaginationOptions, input *models.ListModelsInput) error {
	var (
		offset    = common.DefaultPageOffset
		limit     = common.MaxPageLimit
		sortKey   = datacatalog.PaginationOptions_CREATION_TIME
		sortOrder = datacatalog.PaginationOptions_DESCENDING
	)

	if paginationOpts != nil {
		// if the token is empty, that is still valid input since it is optional
		if len(strings.Trim(paginationOpts.Token, " ")) == 0 {
			offset = common.DefaultPageOffset
		} else {
			parsedOffset, err := strconv.ParseInt(paginationOpts.Token, 10, 32)
			if err != nil {
				return errors.NewDataCatalogErrorf(codes.InvalidArgument, "Invalid token %v", offset)
			}
			offset = int(parsedOffset)
		}
		limit = int(paginationOpts.Limit)
		sortKey = paginationOpts.SortKey
		sortOrder = paginationOpts.SortOrder
	}

	input.Offset = offset
	input.Limit = limit
	input.SortParameter = gormimpl.NewGormSortParameter(sortKey, sortOrder)
	return nil
}
