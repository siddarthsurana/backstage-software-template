package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Todo represents a task item with an ID, title, and completion status
type Todo struct {
	ID        string `json:"id"`        // Unique identifier for the todo
	Title     string `json:"title"`     // Title or description of the todo
	Completed bool   `json:"completed"` // Whether the todo is completed or not
}

// todos is an in-memory storage for todo items
var todos = []Todo{
	{ID: "1", Title: "Learn Go", Completed: false},
	{ID: "2", Title: "Build API", Completed: false},
}

// main is the entry point of the application
// It sets up the Gin router and defines all API endpoints
func main() {
	router := gin.Default()

	// Hello endpoint
	// GET /hello
	// Returns a greeting message
	router.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello from Todos API!",
		})
	})

	// Health check endpoint
	// GET /health
	// Returns the status of the API server
	router.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	// Get all todos
	// GET /todos
	// Returns a list of all todo items
	router.GET("/todos", func(c *gin.Context) {
		c.JSON(http.StatusOK, todos)
	})
	router.GET("/ready", func(c *gin.Context) {
		// Here you can add checks like DB connectivity if needed
		c.JSON(http.StatusOK, gin.H{"status": "ready"})
	})
	// Get a specific todo
	// GET /todos/:id
	// Returns a single todo item by its ID
	// Returns 404 if the todo is not found
	router.GET("/todos/:id", func(c *gin.Context) {
		id := c.Param("id")
		for _, todo := range todos {
			if todo.ID == id {
				c.JSON(http.StatusOK, todo)
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
	})

	// Create a new todo
	// POST /todos
	// Creates a new todo item
	// Expects a JSON body with title and completed fields
	// Returns the created todo item
	router.POST("/todos", func(c *gin.Context) {
		var newTodo Todo
		if err := c.ShouldBindJSON(&newTodo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		todos = append(todos, newTodo)
		c.JSON(http.StatusCreated, newTodo)
	})

	// Start the server on port 8080
	err := router.Run(":8080")
	if err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
