package router

import (
	"github.com/bytebeatz/bandroom-cms/api/handler"
	"github.com/bytebeatz/bandroom-cms/api/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter(
	courseHandler *handler.CourseHandler,
	unitHandler *handler.UnitHandler,
	skillHandler *handler.SkillHandler,
	lessonHandler *handler.LessonHandler,
) *gin.Engine {
	r := gin.New()

	// Core middleware
	r.Use(middleware.RecoveryMiddleware())
	r.Use(middleware.LoggerMiddleware())
	r.Use(middleware.CORSMiddleware())

	// Health check
	r.GET("/health", handler.HealthCheck)

	// Protected API
	api := r.Group("/api", middleware.AuthMiddleware())
	{
		// Course routes
		courses := api.Group("/courses")
		courses.Use(middleware.RequireAdmin())
		{
			courses.POST("", courseHandler.Create)
			courses.GET("", courseHandler.List)
			courses.GET("/:id", courseHandler.GetByID)
			courses.PUT("/:id", courseHandler.Update)
			courses.DELETE("/:id", courseHandler.Delete)
		}

		// Unit routes
		units := api.Group("/units")
		units.Use(middleware.RequireAdmin())
		{
			units.POST("", unitHandler.Create)
			units.GET("/course/:courseId", unitHandler.ListByCourse)
			units.GET("/:id", unitHandler.GetByID)
			units.PUT("/:id", unitHandler.Update)
			units.DELETE("/:id", unitHandler.Delete)
		}

		// Skill routes
		skills := api.Group("/skills")
		skills.Use(middleware.RequireAdmin())
		{
			skills.POST("", skillHandler.Create)
			skills.GET("", skillHandler.ListByUnit) // expects ?unit_id= query param
			skills.GET("/:id", skillHandler.GetByID)
			skills.PUT("/:id", skillHandler.Update)
			skills.DELETE("/:id", skillHandler.Delete)
		}

		// Lesson routes
		lessons := api.Group("/lessons")
		lessons.Use(middleware.RequireAdmin())
		{
			lessons.POST("", lessonHandler.Create)
			lessons.GET("", lessonHandler.ListBySkill) // expects ?skill_id= query param
			lessons.GET("/:id", lessonHandler.GetByID)
			lessons.PUT("/:id", lessonHandler.Update)
			lessons.DELETE("/:id", lessonHandler.Delete)
		}
	}

	return r
}

