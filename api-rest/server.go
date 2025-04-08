package main

import (
	"net/http"
	"strconv"

	"slices"

	"github.com/gin-gonic/gin"
)

type Person struct {
	Id      int
	Name    string  `json:"name" binding:"required"`
	Age     int     `json:"age" binding:"number"`
	Address Address `json:"address" binding:"required"`
}

type Address struct {
	Street       string `json:"street" binding:"required"`
	Neighborhood string `json:"neighborhood" binding:"required"`
	HouseNumber  int    `json:"houseNumber" binding:"required,number"`
}

func RemoveIndex(s []Person, index int) []Person {
	return slices.Delete(s, index, index+1)
}

func main() {
	var persons []Person
	router := gin.Default()

	router.POST("/", func(c *gin.Context) {
		var req Person
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		person := Person{
			Id:      len(persons) + 1,
			Name:    req.Name,
			Age:     req.Age,
			Address: req.Address,
		}

		persons = append(persons, person)

		c.JSON(http.StatusCreated, gin.H{
			"message": "Created with success!",
			"data":    person,
		})
	})

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"persons": persons,
		})
	})

	router.GET("/:id", func(c *gin.Context) {
		idStr := c.Param("id")

		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
			return
		}

		index := id - 1

		if index >= 0 && index < len(persons) {
			c.JSON(http.StatusOK, gin.H{"person": persons[index]})
		} else {
			c.JSON(http.StatusNotFound, gin.H{"error": "Person not found"})
		}
	})

	router.PUT("/:id", func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
			return
		}

		var updateReq Person
		if err := c.ShouldBindJSON(&updateReq); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		found := false
		for _, person := range persons {
			if person.Id == id {
				found = true

				person := Person{
					Id:      id,
					Name:    updateReq.Name,
					Age:     updateReq.Age,
					Address: updateReq.Address,
				}

				index := id - 1
				persons[index] = person

				c.JSON(http.StatusOK, gin.H{"message": "Person updated successfully"})
				return
			}
		}

		if !found {
			c.JSON(http.StatusNotFound, gin.H{"error": "Person not found"})
		}
	})

	router.Run(":3000")
}
