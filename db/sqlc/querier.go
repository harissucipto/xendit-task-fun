// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"context"
)

type Querier interface {
	CreateComment(ctx context.Context, arg CreateCommentParams) (Comment, error)
	DeleteComment(ctx context.Context, orgName string) error
	ListComments(ctx context.Context, orgName string) ([]Comment, error)
}

var _ Querier = (*Queries)(nil)