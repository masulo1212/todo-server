package api

import (
	"database/sql"
	"net/http"
	db "todo/broker/db/sqlc"

	"github.com/gin-gonic/gin"
)

type CreateFavRequest struct {
	Title       string `json:"title" binding:"required"`
	ID          string `json:"id" binding:"required"`
	Description string `json:"description" binding:"required"`
	Date        string `json:"date" binding:"required"`
	IsDone      bool   `json:"is_done"`
	IsDelete    bool   `json:"is_delete"`
	IsFavorite  bool   `json:"is_favorite"`
}

type UpdateFavRequest struct {
	Title       string `json:"title,omitempty"`
	ID          string `json:"id" binding:"required"`
	Description string `json:"description,omitempty"`
	Date        string `json:"date,omitempty"`
	IsDone      bool   `json:"is_done,omitempty"`
	IsDelete    bool   `json:"is_delete,omitempty"`
	IsFavorite  bool   `json:"is_favorite,omitempty"`
}

type DeleteFavRequest struct {
	ID string `json:"id" binding:"required"`
}

type GetFavRequest struct {
	ID string `json:"id" binding:"required"`
}

func (server *Server) CreateFavAPI(ctx *gin.Context) {
	var req CreateFavRequest

	//binding
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateFavParams{
		Title:       req.Title,
		ID:          req.ID,
		Description: req.Description,
		Date:        req.Date,
		IsDone:      req.IsDone,
		IsDelete:    req.IsDelete,
		IsFavorite:  req.IsFavorite,
	}

	createTask, err := server.Queries.CreateFav(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusForbidden, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, createTask)

}

func (server *Server) UpdateFavAPI(ctx *gin.Context) {
	var req UpdateFavRequest

	//binding
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateFavParams{
		ID:         req.ID,
		IsDone:     req.IsDone,
		IsDelete:   req.IsDelete,
		IsFavorite: req.IsFavorite,
	}

	if req.Title != "" {
		arg.Title = sql.NullString{
			String: req.Title,
			Valid:  true,
		}
	}

	if req.Description != "" {
		arg.Description = sql.NullString{
			String: req.Description,
			Valid:  true,
		}
	}

	if req.Date != "" {
		arg.Date = sql.NullString{
			String: req.Date,
			Valid:  true,
		}
	}

	updateTask, err := server.Queries.UpdateFav(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusForbidden, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, updateTask)

}

func (server *Server) DeleteFavAPI(ctx *gin.Context) {
	var req DeleteFavRequest

	//binding
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := server.Queries.DeleteFav(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusForbidden, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "delete successfully",
	})

}

func (server *Server) GetFavAPI(ctx *gin.Context) {
	var req GetFavRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	getFavTask, err := server.Queries.GetFav(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusForbidden, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, getFavTask)
}

func (server *Server) GetAllFavAPI(ctx *gin.Context) {

	FavTaskList, err := server.Queries.GetAllFav(ctx)
	if err != nil {
		ctx.JSON(http.StatusForbidden, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, FavTaskList)

}
