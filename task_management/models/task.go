package models

import (
	"time"
)


/**
 * Task - Struct representing a task model
 * 
 * Fields:
 * - Id:          Unique identifier for the task (UUID as string)
 * - Title:       Title of the task
 * - Description: Detailed explanation of the task
 * - Status:      Current status of the task (e.g., "IN PROGRESS", "COMPLETED")
 * - CreatedAt:   Timestamp when the task was created
 * - UpdatedAt:   Timestamp of the last update to the task
 * 
 * JSON tags are used to map the struct fields to JSON keys during
 * serialization and deserialization.
 */

type Task struct {
    Id          string  `json:"id"`
    Title       string     `json:"title"`
    Description string     `json:"description"`
    Status      string	   `json:"status"`
    CreatedAt   time.Time  `json:"created_at"`
    UpdatedAt   time.Time  `json:"updated_at"`
}