package main

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"log"
	"net/http"
	"time"
)

type OperationType int

const (
	Study OperationType = iota + 1
	Status
	Subscribe
	Unsubscribe
)

type StudyType int

const (
	TOPIC StudyType = iota + 1
	TRACK
)

type StudyOperation struct {
	Type      OperationType `json:"type" binding:"required,oneof=1 2 3 4"`
	StudyType StudyType     `json:"studyType" binding:"required,oneof=1 2"`
	TN        int           `json:"tn" binding:"required"`
	StudyID   int           `json:"studyId" binding:"required"`
	Date      time.Time     `json:"date" binding:"datetime"`
}

type StudyRequest struct {
	SystemId    string            `json:"systemId" binding:"required"`
	RequestUUID string            `json:"requestUUID" binding:"required"`
	Operations  []*StudyOperation `json:"operations" binding:"required,gte=1,dive"`
}

type ErrorMsg struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func getErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "gte":
		return "This field shouldn't be empty " + fe.Param()
	case "oneof":
		return "Should be one of " + fe.Param()
	}
	return "Unknown error"
}

func main() {
	engine := gin.New()
	engine.POST("/study", func(context *gin.Context) {
		body := StudyRequest{}
		if err := context.ShouldBind(&body); err != nil {
			var ve validator.ValidationErrors
			if errors.As(err, &ve) {
				out := make([]ErrorMsg, len(ve))
				for i, fe := range ve {
					out[i] = ErrorMsg{fe.Field(), getErrorMsg(fe)}
				}
				context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": out})
			} else {
				_ = context.AbortWithError(http.StatusBadRequest, err)
			}
			fmt.Println(err)
			return
		}
		fmt.Println(body)
		context.JSON(http.StatusAccepted, &body)
	})
	if err := engine.Run(":3000"); err != nil {
		log.Fatal(err)
	}
}
