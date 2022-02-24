// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"time"
)

type Comment struct {
	ID        int64     `json:"id"`
	Content   string    `json:"content"`
	OrgName   string    `json:"org_name"`
	CreatedAt time.Time `json:"created_at"`
}

type DeletedComment struct {
	ID        int64     `json:"id"`
	Content   string    `json:"content"`
	OrgName   string    `json:"org_name"`
	CreatedAt time.Time `json:"created_at"`
}
