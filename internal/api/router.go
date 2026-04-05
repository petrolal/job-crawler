package api

import (
	"jobs-crawler/internal/api/handler"
	"jobs-crawler/internal/domain"

	"github.com/gin-gonic/gin"
)

func NewRouter(jobs []domain.Job) *gin.Engine {
	r := gin.Default()

	jobsHandler := &handler.JobsHandler{Jobs: jobs}

	v1 := r.Group("/v1")
	{
		v1.GET("/jobs", jobsHandler.List)
	}

	return r
}
