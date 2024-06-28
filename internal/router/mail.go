package router

import (
	"mail-service/internal/controller"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/google/uuid"
)

func (r *Router) SendMail(c *gin.Context) {
	var mailRequest []*controller.EmailRequest
	if err := c.ShouldBindJSON(&mailRequest); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := r.controller.PublishMails(mailRequest)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}
