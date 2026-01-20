package eventstream

import (
	"time"

	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func NewMessageMetadata() *MessageMetadata {
	timestamp := time.Now()

	return &MessageMetadata{
		EventId: uuid.New().String(),
		OccurredAt: &timestamppb.Timestamp{
			Seconds: timestamp.Unix(),
			Nanos:   int32(timestamp.Nanosecond()),
		},
	}
}
