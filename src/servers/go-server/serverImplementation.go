package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "strconv"
)

// User struct to hold user data
type User struct {
    ID          int     `json:"id"`
    Name        string  `json:"name"`
    HoursWorked float64 `json:"hoursWorked"`
}

var users []User
var nextID = 1

func main() {
    router := gin.Default()
    router.Use(gin.Logger())
    router.Use(gin.Recovery())
    router.Use(corsMiddleware())

    router.GET("/users", getAllUsers)
    router.GET("/users/:id", getUserByID)
    router.POST("/users", addUser)
    router.PUT("/users/:id", updateUserByID)
    router.PATCH("/users/:id", updateUserHours)
    router.DELETE("/users", deleteAllUsers)
    router.DELETE("/users/:id", deleteUserByID)

    router.Run(":5004")
}

func corsMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(http.StatusNoContent)
            return
        }
        c.Next()
    }
}

func getAllUsers(c *gin.Context) {
    c.JSON(http.StatusOK, users)
}

func getUserByID(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
        return
    }
    for _, user := range users {
        if user.ID == id {
            c.JSON(http.StatusOK, user)
            return
        }
    }
    c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
}

func addUser(c *gin.Context) {
    var input struct {
        Name string `json:"name"`
    }
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }
    user := User{ID: nextID, Name: input.Name, HoursWorked: 0}
    users = append(users, user)
    nextID++
    c.JSON(http.StatusCreated, user)
}

func updateUserByID(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
        return
    }
    var input struct {
        Name string `json:"name"`
    }
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }
    for i, user := range users {
        if user.ID == id {
            users[i].Name = input.Name
            c.JSON(http.StatusOK, users[i])
            return
        }
    }
    c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
}

func updateUserHours(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
        return
    }
    var input struct {
        HoursToAdd float64 `json:"hoursToAdd"`
    }
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }
    for i, user := range users {
        if user.ID == id {
            users[i].HoursWorked += input.HoursToAdd
            c.JSON(http.StatusOK, users[i])
            return
        }
    }
    c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
}

func deleteAllUsers(c *gin.Context) {
    users = []User{}
    nextID = 1
    c.JSON(http.StatusOK, users)
}

func deleteUserByID(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
        return
    }
    for i, user := range users {
        if user.ID == id {
            users = append(users[:i], users[i+1:]...)
            c.JSON(http.StatusOK, user)
            return
        }
    }
    c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
}