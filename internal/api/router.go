// Package api wires together the HTTP router and all route handlers.
package api

import (
	"jobs-crawler/internal/api/handler"
	"jobs-crawler/internal/store"

	"github.com/gin-gonic/gin"
)

// NewRouter creates and returns a configured Gin engine with all routes registered.
func NewRouter(jobStore *store.JobStore) *gin.Engine {
	r := gin.Default()

	jobsHandler := &handler.JobsHandler{Store: jobStore}

	v1 := r.Group("/v1")
	{
		v1.GET("/jobs", jobsHandler.List)
	}

	return r
}
