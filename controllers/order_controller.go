package controllers

import (
	"go-backend/config"
	"go-backend/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func GetOrders(c *gin.Context) {
	var orders []models.Order
	var total int64

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset := (page - 1) * limit
	search := c.Query("search")
	query := config.DB.Preload("Customer")

	if search != "" {
		query = query.Joins("JOIN customers ON customers.id = orders.customer_id").
			Where("orders.status ILIKE ? OR customers.name ILIKE ?", "%"+search+"%", "%"+search+"%")
	}

	query.Model(&models.Order{}).Count(&total)
	query.Offset(offset).Limit(limit).Find(&orders)

	c.JSON(http.StatusOK, gin.H{
		"data":  orders,
		"total": total,
		"page":  page,
		"limit": limit,
	})
}

func GetOrder(c *gin.Context) {
	id := c.Param("id")
	var order models.Order

	if err := config.DB.Preload("Customer").First(&order, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	c.JSON(http.StatusOK, order)
}

func CreateOrder(c *gin.Context) {
	var input struct {
		CustomerID uint    `json:"customer_id" binding:"required"`
		OrderDate  string  `json:"order_date" binding:"required"` // format yyyy-mm-dd
		Status     string  `json:"status" binding:"required"`
		TotalPrice float64 `json:"total_price" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	orderDate, err := time.Parse("2006-01-02", input.OrderDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order_date format, expected yyyy-mm-dd"})
		return
	}

	order := models.Order{
		CustomerID: input.CustomerID,
		OrderDate:  orderDate,
		Status:     input.Status,
		TotalPrice: input.TotalPrice,
	}

	if err := config.DB.Create(&order).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Preload("Customer").First(&order, order.ID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load customer data"})
		return
	}

	c.JSON(http.StatusCreated, order)
}

func UpdateOrder(c *gin.Context) {
	id := c.Param("id")
	var order models.Order

	if err := config.DB.First(&order, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	var input struct {
		CustomerID uint    `json:"customer_id"`
		OrderDate  string  `json:"order_date"`
		Status     string  `json:"status"`
		TotalPrice float64 `json:"total_price"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.OrderDate != "" {
		orderDate, err := time.Parse("2006-01-02", input.OrderDate)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order_date format, expected yyyy-mm-dd"})
			return
		}
		order.OrderDate = orderDate
	}

	if input.CustomerID != 0 {
		order.CustomerID = input.CustomerID
	}
	if input.Status != "" {
		order.Status = input.Status
	}
	if input.TotalPrice != 0 {
		order.TotalPrice = input.TotalPrice
	}

	if err := config.DB.Save(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Preload Customer untuk mengisi detail customer dalam response
	if err := config.DB.Preload("Customer").First(&order, order.ID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load customer details"})
		return
	}

	c.JSON(http.StatusOK, order)
}

func DeleteOrder(c *gin.Context) {
	id := c.Param("id")
	if err := config.DB.Delete(&models.Order{}, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Order deleted"})
}
