package api

import (
	"net/http"
	db "todo/broker/db/sqlc"

	"github.com/gin-gonic/gin"
)

type CreateBinRequest struct {
	Title       string `json:"title" binding:"required"`
	ID          string `json:"id" binding:"required"`
	Description string `json:"description" binding:"required"`
	Date        string `json:"date" binding:"required"`
	IsDone      bool   `json:"is_done"`
	IsDelete    bool   `json:"is_delete"`
	IsFavorite  bool   `json:"is_favorite"`
}

type DeleteBinRequest struct {
	ID string `json:"id" binding:"required"`
}

type GetBinRequest struct {
	ID string `json:"id" binding:"required"`
}

func (server *Server) CreateBinAPI(ctx *gin.Context) {
	var req CreateBinRequest

	//binding
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateBinParams{
		Title:       req.Title,
		ID:          req.ID,
		Description: req.Description,
		Date:        req.Date,
		IsDone:      req.IsDone,
		IsDelete:    req.IsDelete,
		IsFavorite:  req.IsFavorite,
	}

	createTask, err := server.Queries.CreateBin(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusForbidden, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, createTask)

}

func (server *Server) DeleteBinAPI(ctx *gin.Context) {
	var req DeleteBinRequest

	//binding
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := server.Queries.DeleteBin(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusForbidden, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "delete successfully",
	})

}

func (server *Server) GetBinAPI(ctx *gin.Context) {
	var req GetBinRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	getBinTask, err := server.Queries.GetBin(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusForbidden, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, getBinTask)
}

func (server *Server) GetAllBinAPI(ctx *gin.Context) {

	BinTaskList, err := server.Queries.GetAllBin(ctx)
	if err != nil {
		ctx.JSON(http.StatusForbidden, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, BinTaskList)

}

func (server *Server) DeleteAllBinAPI(ctx *gin.Context) {
	_, err := server.pg.ExecContext(ctx, "DELETE FROM bin_tasks")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "delete all successfully",
	})
}
