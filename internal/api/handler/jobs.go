package handler

import (
	"net/http"
	"strconv"

	"jobs-crawler/internal/domain"

	"github.com/gin-gonic/gin"
)

type JobsHandler struct {
	Jobs []domain.Job
}

type paginatedResponse struct {
	Data     []domain.Job `json:"data"`
	Page     int          `json:"page"`
	PageSize int          `json:"page_size"`
	Total    int          `json:"total"`
}

func (h *JobsHandler) List(ctx *gin.Context) {
	page, err := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}

	pageSize, err := strconv.Atoi(ctx.DefaultQuery("page_size", "20"))
	if err != nil || pageSize < 1 || pageSize > 100 {
		pageSize = 20
	}

	total := len(h.Jobs)
	start := (page - 1) * pageSize
	if start >= total {
		ctx.JSON(http.StatusOK, paginatedResponse{
			Data:     []domain.Job{},
			Page:     page,
			PageSize: pageSize,
			Total:    total,
		})
		return
	}

	end := start + pageSize
	if end > total {
		end = total
	}

	ctx.JSON(http.StatusOK, paginatedResponse{
		Data:     h.Jobs[start:end],
		Page:     page,
		PageSize: pageSize,
		Total:    total,
	})
}
