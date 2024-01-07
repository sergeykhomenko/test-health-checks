package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/sergeykhomenko/test-health-checks/internal/app/services"
	"github.com/sergeykhomenko/test-health-checks/internal/types"
	"net/http"
)

type PingHandlers struct{}

func (h *PingHandlers) PingUrls(handlerCtx *gin.Context) {
	var (
		body         types.UrlsDto
		pingResponse types.PingResponse
	)

	err := handlerCtx.ShouldBindJSON(&body)
	if err != nil {
		handlerCtx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	switch body.Strategy {
	case string(types.PingStrategyFirstToFall):
		pingResponse = services.FirstToFall(body.Urls)
		break
	default:
		pingResponse = services.AtLeastOne(body.Urls)
	}

	handlerCtx.JSON(pingResponse.Status, pingResponse.Content)
}
