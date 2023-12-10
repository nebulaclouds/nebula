package transformers

import (
	"time"

	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc/codes"

	"github.com/nebulaclouds/nebula/datacatalog/pkg/errors"
	"github.com/nebulaclouds/nebula/datacatalog/pkg/repositories/models"
	"github.com/nebulaclouds/nebula/nebulaidl/gen/pb-go/nebulaidl/datacatalog"
)

func FromReservationID(reservationID *datacatalog.ReservationID) models.ReservationKey {
	datasetID := reservationID.DatasetId

	return models.ReservationKey{
		DatasetProject: datasetID.Project,
		DatasetDomain:  datasetID.Domain,
		DatasetName:    datasetID.Name,
		DatasetVersion: datasetID.Version,
		TagName:        reservationID.TagName,
	}
}

func CreateReservation(reservation *models.Reservation, heartbeatInterval time.Duration) (datacatalog.Reservation, error) {
	expiresAtPb, err := ptypes.TimestampProto(reservation.ExpiresAt)
	if err != nil {
		return datacatalog.Reservation{}, errors.NewDataCatalogErrorf(codes.Internal, "failed to serialize expires at time")
	}

	heartbeatIntervalPb := ptypes.DurationProto(heartbeatInterval)
	return datacatalog.Reservation{
		ReservationId: &datacatalog.ReservationID{
			DatasetId: &datacatalog.DatasetID{
				Project: reservation.DatasetProject,
				Domain:  reservation.DatasetDomain,
				Name:    reservation.DatasetName,
				Version: reservation.DatasetVersion,
			},
			TagName: reservation.TagName,
		},
		OwnerId:           reservation.OwnerID,
		HeartbeatInterval: heartbeatIntervalPb,
		ExpiresAt:         expiresAtPb,
	}, nil
}
