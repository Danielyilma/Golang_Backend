package controllers

import (
	"fmt"
	"time"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/Danielyilma/database"
	"github.com/Danielyilma/models"
)


/**
 * GetTasks - Handler to fetch all tasks
 * Return: JSON response with a list of all tasks
 */

func GetTasks() gin.HandlerFunc {
	return func(c *gin.Context) {
		var tasks []models.Task

		for _, val := range database.Tasks {
			tasks = append(tasks, val)
		}

		c.JSON(http.StatusOK, gin.H{"tasks": tasks})
	}
}


/**
 * GetTask - Handler to fetch a single task by ID
 * Return: JSON response with the task details if found, error otherwise
 */

func GetTask() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		task, ok := database.Tasks[id]

		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("task with %s not found", id)})
			return
		}
	
		c.JSON(http.StatusOK, gin.H{"task": task})

	}     
}


/**
 * UpdateTask - Handler to update an existing task
 * Return: JSON response with the updated task or error
 */

func UpdateTask() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		
		existingTask, ok := database.Tasks[id]
		
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("task with %s not found", id)})
			return
		}

		var task models.Task

		if err := c.BindJSON(&task); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}


		task.Id = existingTask.Id
        task.CreatedAt = existingTask.CreatedAt
        task.UpdatedAt = time.Now()

		database.Tasks[id] = task

		c.JSON(http.StatusOK, gin.H{"sucess": task})
	}    
}


/**
 * DeleteTask - Handler to delete a task by ID
 * Return: JSON confirmation message or error
 */

func DeleteTask() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		_, ok := database.Tasks[id]

		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("task with %s not found", id)})
			return
		}
	
		delete(database.Tasks, id)

		c.JSON(http.StatusOK, gin.H{"sucess": fmt.Sprintf("task with %s deleted", id)})
	}
}


/**
 * CreateTask - Handler to create a new task
 * Return: JSON response with the created task or error
 */

func CreateTask() gin.HandlerFunc {
	return func(c *gin.Context) {
		var task models.Task

		if err := c.BindJSON(&task); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		task.Id = uuid.New().String()
		task.CreatedAt = time.Now()
		task.UpdatedAt = time.Now()
		task.Status = "IN PROGRESS"

		database.Tasks[task.Id] = task

		c.JSON(http.StatusCreated, gin.H{"success": task})
	}
}

