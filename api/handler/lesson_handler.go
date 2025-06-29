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

type LessonHandler struct {
	lessonService *service.LessonService
}

func NewLessonHandler(svc *service.LessonService) *LessonHandler {
	return &LessonHandler{lessonService: svc}
}

func (h *LessonHandler) Create(c *gin.Context) {
	var req dto.LessonRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("Failed to bind lesson JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	lesson := req.ToModel()
	err := h.lessonService.CreateLesson(c.Request.Context(), &lesson, lesson.SkillID)
	if err != nil {
		if strings.Contains(err.Error(), "already exists") {
			c.JSON(
				http.StatusConflict,
				gin.H{"error": "Lesson title already exists for this skill"},
			)
			return
		}
		log.Println("Failed to create lesson:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create lesson"})
		return
	}

	c.JSON(http.StatusCreated, dto.FromLessonModel(lesson))
}

func (h *LessonHandler) GetByID(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid lesson ID"})
		return
	}

	lesson, err := h.lessonService.GetLessonByID(c, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Lesson not found"})
		return
	}

	c.JSON(http.StatusOK, dto.FromLessonModel(*lesson))
}

func (h *LessonHandler) ListBySkill(c *gin.Context) {
	skillID, err := uuid.Parse(c.Query("skill_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid skill ID"})
		return
	}

	lessons, err := h.lessonService.ListLessonsBySkillID(c, skillID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not list lessons"})
		return
	}

	var res []dto.LessonResponse
	for _, l := range lessons {
		res = append(res, dto.FromLessonModel(*l))
	}

	c.JSON(http.StatusOK, gin.H{"lessons": res})
}

func (h *LessonHandler) Update(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid lesson ID"})
		return
	}

	var req dto.LessonRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	lesson := req.ToModel()
	lesson.ID = id

	if err := h.lessonService.UpdateLesson(c.Request.Context(), &lesson); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update lesson"})
		return
	}

	c.JSON(http.StatusOK, dto.FromLessonModel(lesson))
}

func (h *LessonHandler) Delete(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid lesson ID"})
		return
	}

	if err := h.lessonService.DeleteLesson(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete lesson"})
		return
	}

	c.Status(http.StatusNoContent)
}

