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

// CourseHandler defines HTTP handlers for course operations.
type CourseHandler struct {
	courseService *service.CourseService
}

// NewCourseHandler initializes a new CourseHandler.
func NewCourseHandler(svc *service.CourseService) *CourseHandler {
	return &CourseHandler{courseService: svc}
}

// Create handles POST /api/courses
func (h *CourseHandler) Create(c *gin.Context) {
	var req dto.CourseRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("Failed to bind JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	log.Printf("Parsed CourseRequest: %+v\n", req)

	userID := c.GetString("user_id")
	role := c.GetString("role")

	log.Printf("Authenticated user: %s with role %s\n", userID, role)

	course := req.ToModel() // Returns model.Course
	if userID != "" {
		parsedID, err := uuid.Parse(userID)
		if err == nil {
			course.CreatorID = &parsedID
		}
	}

	err := h.courseService.CreateCourse(c.Request.Context(), &course) // Pass pointer
	if err != nil {
		// Check for title conflict
		if strings.Contains(err.Error(), "already exists") {
			c.JSON(http.StatusConflict, gin.H{"error": "Course title already exists"})
			return
		}
		log.Println("Failed to create course:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create course"})
		return
	}

	c.JSON(http.StatusCreated, dto.FromModel(course)) // No need to deref
}

// GetByID handles GET /api/courses/:id
func (h *CourseHandler) GetByID(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	course, err := h.courseService.GetCourseByID(c, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Course not found"})
		return
	}

	c.JSON(http.StatusOK, dto.FromModel(*course))
}

// Update handles PUT /api/courses/:id
func (h *CourseHandler) Update(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course ID"})
		return
	}

	var req dto.CourseRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	course := req.ToModel()
	course.ID = id
	userID := c.GetString("user_id")
	if userID != "" {
		if parsed, err := uuid.Parse(userID); err == nil {
			course.CreatorID = &parsed
		}
	}

	if err := h.courseService.UpdateCourse(c.Request.Context(), &course); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update course"})
		return
	}

	c.JSON(http.StatusOK, dto.FromModel(course))
}

// Delete handles DELETE /api/courses/:id
func (h *CourseHandler) Delete(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course ID"})
		return
	}

	if err := h.courseService.DeleteCourse(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete course"})
		return
	}

	c.Status(http.StatusNoContent)
}

// List handles GET /api/courses
func (h *CourseHandler) List(c *gin.Context) {
	courses, err := h.courseService.ListCourses(c, false)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not list courses"})
		return
	}

	var res []dto.CourseResponse
	for _, course := range courses {
		res = append(res, dto.FromModel(*course))
	}

	c.JSON(http.StatusOK, gin.H{
		"courses": res,
	})
}

