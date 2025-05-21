package controllers

import (
	"go-backend/config"
	"go-backend/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetCustomers(c *gin.Context) {
	userIDInterface, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	userID, ok := userIDInterface.(uint)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID"})
		return
	}

	var customers []models.Customer
	var total int64

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset := (page - 1) * limit
	search := c.Query("search")

	query := config.DB.Model(&models.Customer{}).Where("user_id = ?", userID)

	if search != "" {
		query = query.Where("name ILIKE ? OR email ILIKE ?", "%"+search+"%", "%"+search+"%")
	}

	query.Count(&total)
	query.Offset(offset).Limit(limit).Find(&customers)

	c.JSON(http.StatusOK, gin.H{
		"data":  customers,
		"total": total,
		"page":  page,
		"limit": limit,
	})
}

func GetCustomer(c *gin.Context) {
	id := c.Param("id")

	userIDInterface, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	userID, ok := userIDInterface.(uint)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID"})
		return
	}

	var customer models.Customer

	if err := config.DB.Where("id = ? AND user_id = ?", id, userID).First(&customer).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found or you don't have access"})
		return
	}

	c.JSON(http.StatusOK, customer)
}

func CreateCustomer(c *gin.Context) {
	var input models.Customer

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userIDInterface, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	userID, ok := userIDInterface.(uint)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID"})
		return
	}
	input.UserID = userID

	if err := config.DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, input)
}

func UpdateCustomer(c *gin.Context) {
	id := c.Param("id")
	var customer models.Customer

	if err := config.DB.First(&customer, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
		return
	}

	var input models.Customer
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Model(&customer).Updates(map[string]interface{}{
		"Name":    input.Name,
		"Email":   input.Email,
		"Phone":   input.Phone,
		"Address": input.Address,
	})

	c.JSON(http.StatusOK, customer)
}

func DeleteCustomer(c *gin.Context) {
	id := c.Param("id")
	if err := config.DB.Delete(&models.Customer{}, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Customer deleted"})
}
