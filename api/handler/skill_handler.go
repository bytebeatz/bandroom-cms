package handler

import (
	"log"
	"net/http"
	"strings"

	"github.com/bytebeatz/bandroom-cms/api/dto"
	"github.com/bytebeatz/bandroom-cms/core/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type SkillHandler struct {
	skillService *service.SkillService
}

func NewSkillHandler(svc *service.SkillService) *SkillHandler {
	return &SkillHandler{skillService: svc}
}

func (h *SkillHandler) Create(c *gin.Context) {
	var req dto.SkillRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("Failed to bind skill JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	log.Printf("Parsed SkillRequest: %+v\n", req)

	userIDStr := c.GetString("user_id")
	role := c.GetString("role")
	log.Printf("Authenticated user: %s with role %s\n", userIDStr, role)

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		log.Println("Invalid user_id in context:", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user"})
		return
	}

	skill := req.ToModel()
	skill.CreatorID = userID // ✅ This is the key line

	err = h.skillService.CreateSkill(c.Request.Context(), &skill, skill.CourseID)
	if err != nil {
		if strings.Contains(err.Error(), "already exists") {
			c.JSON(
				http.StatusConflict,
				gin.H{"error": "Skill title already exists for this course"},
			)
			return
		}
		log.Println("Failed to create skill:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create skill"})
		return
	}

	c.JSON(http.StatusCreated, dto.FromSkillModel(skill))
}

func (h *SkillHandler) GetByID(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid skill ID"})
		return
	}

	skill, err := h.skillService.GetSkillByID(c, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Skill not found"})
		return
	}

	c.JSON(http.StatusOK, dto.FromSkillModel(*skill))
}

func (h *SkillHandler) ListByUnit(c *gin.Context) {
	unitID, err := uuid.Parse(c.Query("unit_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid unit ID"})
		return
	}

	skills, err := h.skillService.ListSkillsByUnitID(c, unitID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not list skills"})
		return
	}

	var res []dto.SkillResponse
	for _, s := range skills {
		res = append(res, dto.FromSkillModel(*s))
	}

	c.JSON(http.StatusOK, res)

}

func (h *SkillHandler) Update(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid skill ID"})
		return
	}

	var req dto.SkillRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	skill := req.ToModel()
	skill.ID = id

	if err := h.skillService.UpdateSkill(c.Request.Context(), &skill); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update skill"})
		return
	}

	c.JSON(http.StatusOK, dto.FromSkillModel(skill))
}

func (h *SkillHandler) Delete(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid skill ID"})
		return
	}

	if err := h.skillService.DeleteSkill(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete skill"})
		return
	}

	c.Status(http.StatusNoContent)
}
