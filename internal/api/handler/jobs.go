package handler

import (
	"net/http"
	"strconv"
	"strings"

	"jobs-crawler/internal/api/render"
	"jobs-crawler/internal/domain"
	"jobs-crawler/internal/store"

	"github.com/gin-gonic/gin"
)

type JobsHandler struct {
	Store *store.JobStore
}

type paginatedResponse struct {
	Status   store.CrawlStatus `json:"status"`
	Data     []domain.Job      `json:"data"`
	Page     int               `json:"page"`
	PageSize int               `json:"page_size"`
	Total    int               `json:"total"`
}

func (h *JobsHandler) List(ctx *gin.Context) {
	page, err := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}

	pageSize, err := strconv.Atoi(ctx.DefaultQuery("page_size", "50"))
	if err != nil || pageSize < 1 || pageSize > 100 {
		pageSize = 50
	}

	title   := strings.ToUpper(ctx.Query("title"))
	company := strings.ToUpper(ctx.Query("company"))
	source  := strings.ToUpper(ctx.Query("source"))

	jobs, status := h.Store.Get()

	filtered := applyFilters(jobs, title, company, source)
	total := len(filtered)

	start := (page - 1) * pageSize
	if start >= total {
		render.JSON(ctx, http.StatusOK, paginatedResponse{
			Status:   status,
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

	render.JSON(ctx, http.StatusOK, paginatedResponse{
		Status:   status,
		Data:     filtered[start:end],
		Page:     page,
		PageSize: pageSize,
		Total:    total,
	})
}

func applyFilters(jobs []domain.Job, title, company, source string) []domain.Job {
	if title == "" && company == "" && source == "" {
		return jobs
	}

	out := make([]domain.Job, 0, len(jobs))
	for _, j := range jobs {
		if title != "" && !strings.Contains(strings.ToUpper(j.Title), title) {
			continue
		}
		if company != "" && !strings.Contains(strings.ToUpper(j.Company), company) {
			continue
		}
		if source != "" && !strings.Contains(strings.ToUpper(j.Source), source) {
			continue
		}
		out = append(out, j)
	}
	return out
}
