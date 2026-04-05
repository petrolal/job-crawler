// Package render provides HTTP response helpers for the API.
package render

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
)

// JSON writes a UTF-8 JSON response without escaping non-ASCII characters
// (e.g. á, ã, ç are kept as-is instead of being encoded as \uXXXX).
func JSON(ctx *gin.Context, code int, obj any) {
	ctx.Header("Content-Type", "application/json; charset=utf-8")
	ctx.Status(code)

	enc := json.NewEncoder(ctx.Writer)
	enc.SetEscapeHTML(false)
	enc.Encode(obj) //nolint:errcheck
}
