package __cookiecutter_project_name__

import "time"

type CreateObject struct {
	ID       string      `json:"type"`
	Metadata interface{} `json:"metadata"`
}

type Object struct {
	ID       string      `json:"id"`
	Object   string      `json:"object"`
	Status   string      `json:"status"`
	Metadata interface{} `json:"metadata"`
}

type StatusTransitions []struct {
	Status string    `json:"status"`
	Date   time.Time `json:"date"`
}

type CurrentStatus struct {
	Status string    `json:"status"`
	Date   time.Time `json:"date"`
}
