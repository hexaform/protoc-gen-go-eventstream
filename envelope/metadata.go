package envelope

import (
	"time"

	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func NewMetadata() *Metadata {
	timestamp := time.Now()

	return &Metadata{
		EventId: uuid.New().String(),
		OccurredAt: &timestamppb.Timestamp{
			Seconds: timestamp.Unix(),
			Nanos:   int32(timestamp.Nanosecond()),
		},
	}
}
