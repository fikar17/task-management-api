package handlers

import (
	"net/http"
	"strconv"
	"task-management-api/models"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TaskHandler struct {
	DB *gorm.DB
}

func isValidStatus(status models.TaskStatus) bool {
	return status == models.StatusPending ||
		status == models.StatusInProgress ||
		status == models.StatusCompleted
}

func NewTaskHandler(db *gorm.DB) *TaskHandler {
	return &TaskHandler{DB: db}
}

func (h *TaskHandler) CreateTask(c *gin.Context) {
	var input models.CreateTaskInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	dueDate, err := time.Parse("2006-01-02", input.DueDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid due_date format. Use YYYY-MM-DD",
		})
		return
	}

	status := input.Status
	if status == "" {
		status = models.StatusPending
	}

	if !isValidStatus(status) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid status. Must be: pending, in-progress, or completed",
		})
		return
	}

	task := models.Task{
		Title:       input.Title,
		Description: input.Description,
		DueDate:     dueDate,
		Status:      status,
	}

	if err := h.DB.Create(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create task",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Task created successfully",
		"data":    task,
	})
}

func (h *TaskHandler) GetTasks(c *gin.Context) {
	var tasks []models.Task

	if err := h.DB.Order("created_at desc").Find(&tasks).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch tasks",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Tasks retrieved successfully",
		"data":    tasks,
	})
}

func (h *TaskHandler) GetTask(c *gin.Context) {
	id := c.Param("id")
	taskID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid task ID",
		})
		return
	}

	var task models.Task
	if err := h.DB.First(&task, taskID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Task not found",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch task",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Task retrieved successfully",
		"data":    task,
	})
}

func (h *TaskHandler) UpdateTask(c *gin.Context) {
	id := c.Param("id")
	taskID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid task ID",
		})
		return
	}

	var task models.Task
	if err := h.DB.First(&task, taskID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Task not found",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch task",
		})
		return
	}

	var input models.UpdateTaskInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Apply patch-like updates
	if input.Title != "" {
		task.Title = input.Title
	}
	if input.Description != "" {
		task.Description = input.Description
	}
	if input.DueDate != "" {
		d, err := time.Parse("2006-01-02", input.DueDate)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid due_date format. Use YYYY-MM-DD",
			})
			return
		}
		task.DueDate = d
	}
	if input.Status != "" {
		if !isValidStatus(input.Status) {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid status. Must be: pending, in-progress, or completed",
			})
			return
		}
		task.Status = input.Status
	}

	if err := h.DB.Save(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update task",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Task updated successfully",
		"data":    task,
	})
}

func (h *TaskHandler) DeleteTask(c *gin.Context) {
	id := c.Param("id")
	taskID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid task ID",
		})
		return
	}

	if err := h.DB.Delete(&models.Task{}, taskID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete task",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Task deleted successfully",
	})
}
