package models

import "time"

type Pagination struct {
	TotalItems int    `json:"totalItems"`
	Page       int    `json:"page"`
	Limit      int    `json:"limit"`
	OrderBy    string `json:"orderBy"`
}

type Conversation struct {
	ID           string     `json:"id"`
	Topic        string     `json:"topic"`
	Summary      string     `json:"summary"`
	Status       string     `json:"status"`
	Priority     string     `json:"priority"`
	QueueID      string     `json:"queueId"`
	AgentID      string     `json:"agentId"`
	CreatedAt    time.Time  `json:"createdAt"`
	UpdatedAt    time.Time  `json:"updatedAt"`
	ClosedAt     *time.Time `json:"closedAt,omitempty"`
	PendingSince time.Time  `json:"pendingSince"`
	FormID       string     `json:"formId"`
}
