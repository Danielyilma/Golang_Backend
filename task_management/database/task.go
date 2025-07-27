package database

import (
	"github.com/Danielyilma/models"
)


/**
 * Tasks - In-memory store for tasks
 * 
 * This map acts as a simple in-memory database to store task data.
 * The key is the task's unique ID (string), and the value is a Task model.
 * 
 * NOTE: This storage is temporary and will be reset when the application restarts.
 */

var Tasks map[string]models.Task = map[string]models.Task{}


