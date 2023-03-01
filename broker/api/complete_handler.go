package api

import (
	"database/sql"
	"net/http"
	db "todo/broker/db/sqlc"

	"github.com/gin-gonic/gin"
)

type CreateCompleteRequest struct {
	Title       string `json:"title" binding:"required"`
	ID          string `json:"id" binding:"required"`
	Description string `json:"description" binding:"required"`
	Date        string `json:"date" binding:"required"`
	IsDone      bool   `json:"is_done"`
	IsDelete    bool   `json:"is_delete"`
	IsFavorite  bool   `json:"is_favorite"`
}

type UpdateCompleteRequest struct {
	Title       string `json:"title,omitempty"`
	ID          string `json:"id" binding:"required"`
	Description string `json:"description,omitempty"`
	Date        string `json:"date,omitempty"`
	IsDone      bool   `json:"is_done,omitempty"`
	IsDelete    bool   `json:"is_delete,omitempty"`
	IsFavorite  bool   `json:"is_favorite,omitempty"`
}

type DeleteCompleteRequest struct {
	ID string `json:"id" binding:"required"`
}

type GetCompleteRequest struct {
	ID string `json:"id" binding:"required"`
}

func (server *Server) CreateCompleteAPI(ctx *gin.Context) {
	var req CreateCompleteRequest

	//binding
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateCompleteParams{
		Title:       req.Title,
		ID:          req.ID,
		Description: req.Description,
		Date:        req.Date,
		IsDone:      req.IsDone,
		IsDelete:    req.IsDelete,
		IsFavorite:  req.IsFavorite,
	}

	createTask, err := server.Queries.CreateComplete(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusForbidden, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, createTask)

}

func (server *Server) UpdateCompleteAPI(ctx *gin.Context) {
	var req UpdateCompleteRequest

	//binding
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateCompleteParams{
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

	updateTask, err := server.Queries.UpdateComplete(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusForbidden, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, updateTask)

}

func (server *Server) DeleteCompleteAPI(ctx *gin.Context) {
	var req DeleteCompleteRequest

	//binding
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := server.Queries.DeleteComplete(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusForbidden, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "delete successfully",
	})

}

func (server *Server) GetCompleteAPI(ctx *gin.Context) {
	var req GetCompleteRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	getCompleteTask, err := server.Queries.GetComplete(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusForbidden, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, getCompleteTask)
}

func (server *Server) GetAllCompleteAPI(ctx *gin.Context) {

	CompleteTaskList, err := server.Queries.GetAllComplete(ctx)
	if err != nil {
		ctx.JSON(http.StatusForbidden, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, CompleteTaskList)

}
