package api

import (
	"database/sql"
	db "todo/broker/db/sqlc"

	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
	*db.Queries
	pg *sql.DB
}

func NewServer(pg *sql.DB) Server {
	router := gin.Default()

	queries := db.New(pg)

	server := Server{
		router:  router,
		Queries: queries,
		pg:      pg, //用在刪除全部的bin 只能手動使用execContext
	}
	server.Routes()

	return server

}

func (server *Server) Routes() {
	server.router.POST("/pending/add-task", server.CreatePendingAPI)
	server.router.POST("/pending/update-task", server.UpdatePendingAPI)
	server.router.POST("/pending/delete-task", server.DeletePendingAPI)
	server.router.GET("/pending/get-task", server.GetPendingAPI)
	server.router.GET("/pending/get-alltask", server.GetAllPendingAPI)

	server.router.POST("/complete/add-task", server.CreateCompleteAPI)
	server.router.POST("/complete/update-task", server.UpdateCompleteAPI)
	server.router.POST("/complete/delete-task", server.DeleteCompleteAPI)
	server.router.GET("/complete/get-task", server.GetCompleteAPI)
	server.router.GET("/complete/get-alltask", server.GetAllCompleteAPI)
	//
	server.router.POST("/fav/add-task", server.CreateFavAPI)
	server.router.POST("/fav/update-task", server.UpdateFavAPI)
	server.router.POST("/fav/delete-task", server.DeleteFavAPI)
	server.router.GET("/fav/get-task", server.GetFavAPI)
	server.router.GET("/fav/get-alltask", server.GetAllFavAPI)
	//
	server.router.POST("/bin/add-task", server.CreateBinAPI)
	server.router.POST("/bin/delete-task", server.DeleteBinAPI)
	server.router.POST("/bin/delete-alltask", server.DeleteAllBinAPI)
	server.router.GET("/bin/get-task", server.GetBinAPI)
	server.router.GET("/bin/get-alltask", server.GetAllBinAPI)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func (server *Server) Start() error {
	err := server.router.Run(":9090")
	if err != nil {
		return err
	}
	return nil
}
