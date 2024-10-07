package main

import (
	"time"

	"github.com/google/uuid"
)

type Todo struct {
	Id          uuid.UUID `json:"id"`
	Todo        string    `json:"todo"`
	Done        bool      `json:"done"`
	CreatedDate time.Time `json:"createdDate"`
	UpdateDate  time.Time `json:"updateDate"`
}
