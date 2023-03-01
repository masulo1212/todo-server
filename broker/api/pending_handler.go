package api

import (
	"database/sql"
	"net/http"
	db "todo/broker/db/sqlc"

	"github.com/gin-gonic/gin"
)

type CreatePendingRequest struct {
	Title       string `json:"title" binding:"required"`
	ID          string `json:"id" binding:"required"`
	Description string `json:"description" binding:"required"`
	Date        string `json:"date" binding:"required"`
	IsDone      bool   `json:"is_done"`
	IsDelete    bool   `json:"is_delete"`
	IsFavorite  bool   `json:"is_favorite"`
}

type UpdatePendingRequest struct {
	Title       string `json:"title,omitempty"`
	ID          string `json:"id" binding:"required"`
	Description string `json:"description,omitempty"`
	Date        string `json:"date,omitempty"`
	IsDone      bool   `json:"is_done,omitempty"`
	IsDelete    bool   `json:"is_delete,omitempty"`
	IsFavorite  bool   `json:"is_favorite,omitempty"`
}

type DeletePendingRequest struct {
	ID string `json:"id" binding:"required"`
}

type GetPendingRequest struct {
	ID string `json:"id" binding:"required"`
}

func (server *Server) CreatePendingAPI(ctx *gin.Context) {
	var req CreatePendingRequest

	//binding
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreatePendingParams{
		Title:       req.Title,
		ID:          req.ID,
		Description: req.Description,
		Date:        req.Date,
		IsDone:      req.IsDone,
		IsDelete:    req.IsDelete,
		IsFavorite:  req.IsFavorite,
	}

	createTask, err := server.Queries.CreatePending(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusForbidden, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, createTask)

}

func (server *Server) UpdatePendingAPI(ctx *gin.Context) {
	var req UpdatePendingRequest

	//binding
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdatePendingParams{
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

	updateTask, err := server.Queries.UpdatePending(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusForbidden, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, updateTask)

}

func (server *Server) DeletePendingAPI(ctx *gin.Context) {
	var req DeletePendingRequest

	//binding
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := server.Queries.DeletePending(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusForbidden, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "delete successfully",
	})

}

func (server *Server) GetPendingAPI(ctx *gin.Context) {
	var req GetPendingRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	getPendingTask, err := server.Queries.GetPending(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusForbidden, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, getPendingTask)
}

func (server *Server) GetAllPendingAPI(ctx *gin.Context) {

	pendingTaskList, err := server.Queries.GetAllPending(ctx)
	if err != nil {
		ctx.JSON(http.StatusForbidden, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, pendingTaskList)

}
