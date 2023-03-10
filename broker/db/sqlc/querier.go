// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package db

import (
	"context"
)

type Querier interface {
	CreateBin(ctx context.Context, arg CreateBinParams) (BinTask, error)
	CreateComplete(ctx context.Context, arg CreateCompleteParams) (CompleteTask, error)
	CreateFav(ctx context.Context, arg CreateFavParams) (FavTask, error)
	CreatePending(ctx context.Context, arg CreatePendingParams) (PendingTask, error)
	DeleteBin(ctx context.Context, id string) error
	DeleteComplete(ctx context.Context, id string) error
	DeleteFav(ctx context.Context, id string) error
	DeletePending(ctx context.Context, id string) error
	GetAllBin(ctx context.Context) ([]BinTask, error)
	GetAllComplete(ctx context.Context) ([]CompleteTask, error)
	GetAllFav(ctx context.Context) ([]FavTask, error)
	GetAllPending(ctx context.Context) ([]PendingTask, error)
	GetBin(ctx context.Context, id string) (BinTask, error)
	GetComplete(ctx context.Context, id string) (CompleteTask, error)
	GetFav(ctx context.Context, id string) (FavTask, error)
	GetPending(ctx context.Context, id string) (PendingTask, error)
	UpdateComplete(ctx context.Context, arg UpdateCompleteParams) (CompleteTask, error)
	UpdateFav(ctx context.Context, arg UpdateFavParams) (FavTask, error)
	UpdatePending(ctx context.Context, arg UpdatePendingParams) (PendingTask, error)
}

var _ Querier = (*Queries)(nil)
