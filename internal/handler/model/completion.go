package model

import (
	"github.com/google/uuid"
	"time"
)

type TopicCompletion struct {
	SystemId           string    `json:"systemId" binding:"required"`
	RequestId          uuid.UUID `json:"requestId" binding:"required"`
	PersonnelNo        string    `json:"personnelNo" binding:"required"`
	TopicId            uint64    `json:"topicId" binding:"required"`
	CompletionDeadline time.Time `json:"completionDeadline" binding:"required"`
}
