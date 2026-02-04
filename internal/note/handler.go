package note

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Type of the handler
type Handler struct {
	service *Service
}

// Func to create handler
func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

// Create
// returns status code 201 on successfull creation
func (h *Handler) Create(c *gin.Context) {

	// Variable to take body of the req
	var req struct {
		Title       string	`json:"title"`
		Description string	`json:"description"`
	}

	// Taking title and description from the body of req
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Creating the handler using service
	note, err := h.service.Create(req.Title, req.Description)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// returning the note on successfull creation
	c.JSON(http.StatusCreated, gin.H{"note": note})

}
