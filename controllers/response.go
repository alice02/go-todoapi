package controllers

import "github.com/alice02/go-todoapi/models"

type (
	ResponseMessage struct {
		Status  string `json:"status"`
		Data    Data   `json:"data"`
		Message string `json:"message,omitempty"`
	}
	Data struct {
		Tasks []models.Task `json:"tasks,omitempty"`
		Task  *models.Task  `json:"task,omitempty"`
		Info  string        `json:"info,omitempty"`
	}
)
