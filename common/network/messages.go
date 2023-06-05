package network

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/thankala/diploma-thesis/common/components"
)

type MessageType int64

type Message struct {
	Type MessageType     `json:"type"`
	Task components.Task `json:"task"`
	Data json.RawMessage
}

const (
	UNUSED MessageType = iota
	CREATE_ARRANGEMENT
	START
	STOP
)

type CreateArrangement struct {
	ArrangementId uuid.UUID `json:"ArrangementId"`
	SenderName    string    `json:"SenderName"`
	TraceId       uuid.UUID `json:"TraceId"`
	Robot1Id      uuid.UUID `json:"Robot1Id"`
	Robot2Id      uuid.UUID `json:"Robot2Id"`
	Robot3Id      uuid.UUID `json:"Robot3Id"`
	Workbench1Id  uuid.UUID `json:"Workbench1Id"`
	Workbench2Id  uuid.UUID `json:"Workbench2Id"`
	SentAt        string    `json:"sentAt"`
}

func (c *CreateArrangement) WithTrace(id uuid.UUID) *CreateArrangement {
	if c.TraceId != uuid.Nil {
		return c
	}
	c.TraceId = id
	return c
}

type Start struct {
	SenderName    string    `json:"SenderName"`
	TraceId       uuid.UUID `json:"TraceId"`
	ArrangementId uuid.UUID `json:"ArrangementId"`
	SentAt        string    `json:"sentAt"`
}

func (start *Start) WithTraceId(id uuid.UUID) *Start {
	if start.TraceId != uuid.Nil {
		return start
	}
	start.TraceId = id
	return start
}

type Stop struct {
	SenderName    string    `json:"SenderName"`
	TraceId       uuid.UUID `json:"TraceId"`
	ArrangementId uuid.UUID `json:"ArrangementId"`
	SentAt        string    `json:"sentAt"`
}

func (stop *Stop) WithTraceId(id uuid.UUID) *Stop {
	if stop.TraceId != uuid.Nil {
		return stop
	}
	stop.TraceId = id
	return stop
}
