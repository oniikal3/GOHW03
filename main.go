package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/api/todos", getTodosHandler)
	r.GET("/api/todos/:id", getTodoHandler)
	r.POST("/api/todos", postTodosHandler)
	r.PUT("/api/todos/:id", putTodoHandler)
	r.DELETE("/api/todos/:id", deleteTodoHandler)
	r.Run(":1234")
}

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

var todos = map[int]Todo{}

func postTodosHandler(c *gin.Context) {
	fmt.Println("postTodosHandler")
	t := Todo{}
	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	id := len(todos)
	id++
	t.ID = id
	todos[id] = t

	fmt.Println(t)

	response := gin.H{
		"id":     fmt.Sprintf("%v", t.ID),
		"title":  fmt.Sprintf("%v", t.Title),
		"status": fmt.Sprintf("%v", t.Status),
	}
	c.JSON(http.StatusCreated, response)
}

func getTodoHandler(c *gin.Context) {
	i, _ := strconv.Atoi(c.Param("id"))
	t := todos[i]
	c.JSON(http.StatusOK, t)
}

func getTodosHandler(c *gin.Context) {
	all := []Todo{}
	for _, v := range todos {
		all = append(all, v)
	}
	c.JSON(http.StatusOK, all)
}

func putTodoHandler(c *gin.Context) {
	i, _ := strconv.Atoi(c.Param("id"))
	t := todos[i]
	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	fmt.Println(t)
	todos[i] = t
	c.JSONP(http.StatusOK, t)
}

func deleteTodoHandler(c *gin.Context) {
	i, _ := strconv.Atoi(c.Param("id"))
	delete(todos, i)
	response := gin.H{
		"todos":  todos,
		"status": "success",
	}
	c.JSON(http.StatusOK, response)
}
