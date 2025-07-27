package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/Danielyilma/controllers"
)

/**
 * BasicRoutes - Registers task-related HTTP routes
 *
 * @incomingRoutes: Pointer to the main Gin engine
 *
 * Description:
 * This function sets up the basic CRUD routes for the Task resource:
 * - GET    /tasks         -> Fetch all tasks
 * - GET    /tasks/:id     -> Fetch a single task by ID
 * - POST   /tasks         -> Create a new task
 * - PUT    /tasks/:id     -> Update an existing task by ID
 * - DELETE /tasks/:id     -> Delete a task by ID
 */


func BasicRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.GET("tasks", controllers.GetTasks())
	incomingRoutes.GET("tasks/:id", controllers.GetTask())
	incomingRoutes.PUT("tasks/:id", controllers.UpdateTask())
	incomingRoutes.DELETE("tasks/:id", controllers.DeleteTask())
	incomingRoutes.POST("tasks", controllers.CreateTask())
}