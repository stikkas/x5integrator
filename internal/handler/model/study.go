package model

import "time"

var SystemIds = []string{"TORG.HR", "ADAPTATION", "LK2", "SKILLAZ"}

type OperationType int

const (
	// Status 1 - разово получить статус обучения
	Status OperationType = iota + 1
	// Study 2 - назначить обучение
	Study
	// Subscribe 3 - подписка на изменение статусов обучения
	Subscribe
	// Unsubscribe 4 - остановить подписку на получение изменения статусов обучения
	Unsubscribe
)

var OperationTypes = []OperationType{Status, Study, Subscribe, Unsubscribe}

type StudyType int

const (
	// Topic 1 - материал
	Topic StudyType = iota + 1
	// Track 2 - трек
	Track
)

var StudyTypes = []StudyType{Topic, Track}

type RequestOperation struct {
	OperationType OperationType `json:"operationType" binding:"operation"`
	Tn            uint          `json:"tn" binding:"required"`
	StudyType     StudyType     `json:"studyType" binding:"study"`
	StudyId       uint          `json:"studyId" binding:"required"`
	Date          time.Time     `json:"date"`
}

type StudyRequest struct {
	SystemId    string              `json:"systemId" binding:"systemId"`
	RequestUUID string              `json:"requestUUID" binding:"required"`
	Data        []*RequestOperation `json:"data" binding:"data,dive"`
}
