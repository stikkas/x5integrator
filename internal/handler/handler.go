package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stikkas/integrator/internal/handler/model"
	"github.com/stikkas/integrator/internal/handler/validate"
	"net/http"
)

type Handler struct {
	engine *gin.Engine
}

func Create(engine *gin.Engine) {
	h := &Handler{engine}
	h.setupRouter("/integration-service")
	validate.Bind()
}

func addRequestUUID(ctx *gin.Context) {
	ctx.Set("requestUUID", uuid.New().String())
}

func topicCompletion(ctx *gin.Context) {
	var body model.TopicCompletion

	if err := ctx.ShouldBindJSON(&body); err != nil {
		//ctx.JSON(http.StatusBadRequest, model.BadResponse(ctx, err))
		return
	}
	ctx.JSON(http.StatusOK, model.EmptyResponse(ctx))
}

func study(ctx *gin.Context) {
	var body model.StudyRequest

	if err := ctx.ShouldBindJSON(&body); err != nil {
		detail := validate.Study(err, ctx.Request.RequestURI)
		if detail != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, model.BadResponse(ctx, detail))
		} else {
			_ = ctx.AbortWithError(http.StatusBadRequest, err)
		}
		return
	}

	// do logic here

	ctx.JSON(http.StatusOK, model.EmptyResponse(ctx))
}
