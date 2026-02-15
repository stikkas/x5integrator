package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var routes = map[string]map[string]func(ctx *gin.Context){
	"/topic-completion": {
		http.MethodPost: topicCompletion,
	},
	"/study": {
		http.MethodPost: study,
	},
}

func (h *Handler) setupRouter(root string) {
	route := h.engine.Group(root)
	route.Use(addRequestUUID)
	for path, d := range routes {
		for meth, fn := range d {
			route.Handle(meth, path, fn)
		}
	}
}
