package handler

import (
	"log"
	"net/http"

	"github.com/bytebeatz/bandroom-cms/api/dto"
	"github.com/bytebeatz/bandroom-cms/core/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// UnitHandler defines HTTP handlers for unit operations.
type UnitHandler struct {
	unitService *service.UnitService
}

// NewUnitHandler initializes a new UnitHandler.
func NewUnitHandler(svc *service.UnitService) *UnitHandler {
	return &UnitHandler{unitService: svc}
}

// Create handles POST /api/units
func (h *UnitHandler) Create(c *gin.Context) {
	var req dto.UnitRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("Failed to bind JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	unit := req.ToModel()
	if err := h.unitService.CreateUnit(c.Request.Context(), &unit); err != nil {
		log.Println("Failed to create unit:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create unit"})
		return
	}

	c.JSON(http.StatusCreated, dto.FromUnitModel(unit))
}

// GetByID handles GET /api/units/:id
func (h *UnitHandler) GetByID(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid unit ID"})
		return
	}

	unit, err := h.unitService.GetUnitByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unit not found"})
		return
	}

	c.JSON(http.StatusOK, dto.FromUnitModel(*unit))
}

// Update handles PUT /api/units/:id
func (h *UnitHandler) Update(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid unit ID"})
		return
	}

	var req dto.UnitRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	unit := req.ToModel()
	unit.ID = id

	if err := h.unitService.UpdateUnit(c.Request.Context(), &unit); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update unit"})
		return
	}

	c.JSON(http.StatusOK, dto.FromUnitModel(unit))
}

// Delete handles DELETE /api/units/:id
func (h *UnitHandler) Delete(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid unit ID"})
		return
	}

	if err := h.unitService.DeleteUnit(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete unit"})
		return
	}

	c.Status(http.StatusNoContent)
}

// ListByCourse handles GET /api/courses/:courseId/units
func (h *UnitHandler) ListByCourse(c *gin.Context) {
	courseID, err := uuid.Parse(c.Param("courseId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course ID"})
		return
	}

	units, err := h.unitService.ListUnitsByCourseID(c.Request.Context(), courseID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch units"})
		return
	}

	var res []dto.UnitResponse
	for _, unit := range units {
		res = append(res, dto.FromUnitModel(*unit))
	}

	c.JSON(http.StatusOK, res)
}
