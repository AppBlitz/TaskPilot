// Package model
package model

import "time"

type StatusTaks string

type Tasks struct {
	ID          int        `json:"id"`
	Description string     `json:"description"`
	CreateAt    time.Time  `json:"createAt"`
	UpdateAt    time.Time  `json:"updateAt"`
	Status      StatusTaks `json:"status"`
}
