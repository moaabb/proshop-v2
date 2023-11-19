package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin" // Import Gin instead of Fiber
	"github.com/moaabb/ecommerce/order_svc/domain/order"
	"github.com/moaabb/ecommerce/order_svc/infra/config"
	"github.com/moaabb/ecommerce/order_svc/infra/database/orderdb"
	"go.uber.org/zap"
)

type OrderHandler struct {
	repository order.Repository
	cfg        *config.Config
	l          *zap.Logger
}

func NewOrderHandler(repo *orderdb.Repository, z *zap.Logger, cfg *config.Config) *OrderHandler {
	return &OrderHandler{
		repository: repo,
		cfg:        cfg,
		l:          z,
	}
}

func (oh *OrderHandler) GetAll(c *gin.Context) {
	oh.l.Info("Fetching orders from the database")
	orders, err := oh.repository.GetAll()
	if err != nil {
		oh.l.Error("Error while fetching orders", zap.Error(err))
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Could not fetch orders"})
		return
	}

	c.JSON(http.StatusOK, orders)
}

func (oh *OrderHandler) GetByUserId(c *gin.Context) {
	oh.l.Info("Getting orders by user ID...")
	userId := c.GetUint("userId")
	if userId == 0 {
		oh.l.Error("coult not get userId")
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	oh.l.Info("Fetching orders by user ID from the database")
	orders, err := oh.repository.GetByUserId(uint(userId))
	if err != nil {
		oh.l.Error("Error while fetching orders by user ID", zap.Error(err))
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Could not fetch orders by user ID"})
		return
	}

	if len(orders) == 0 {
		orders = []order.Order{}
	}
	c.JSON(http.StatusOK, orders)
}

func (oh *OrderHandler) GetById(c *gin.Context) {
	oh.l.Info("Getting order by ID...")
	orderID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		oh.l.Error("Error getting order ID", zap.Error(err))
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Could not fetch order"})
		return
	}

	oh.l.Info("Fetching order from the database")
	order, err := oh.repository.GetById(uint(orderID))
	if err != nil {
		oh.l.Error("Error while fetching order", zap.Error(err))
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Could not fetch order"})
		return
	}

	c.JSON(http.StatusOK, order)
}

func (oh *OrderHandler) Create(c *gin.Context) {
	var o order.Order
	c.BindJSON(&o)
	userId := c.GetUint("userId")
	if userId == 0 {
		oh.l.Error("Error getting order ID")
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Could not create order"})
		return
	}
	o.UserID = int(userId)

	oh.l.Info("Creating order in the database")
	newOrder, err := oh.repository.Create(o)
	if err != nil {
		oh.l.Error("Error while creating order", zap.Error(err))
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": "Could not create order"})
		return
	}

	c.JSON(http.StatusCreated, newOrder)
}

func (oh *OrderHandler) Update(c *gin.Context) {
	oh.l.Info("Getting order ID...")
	orderID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		oh.l.Error("Error getting order ID", zap.Error(err))
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Could not update order"})
		return
	}

	var updatedOrder order.Order
	c.BindJSON(&updatedOrder)

	oh.l.Info("Updating order in the database")
	updatedOrder, err = oh.repository.Update(uint(orderID), updatedOrder)
	if err != nil {
		oh.l.Error("Error while updating order", zap.Error(err))
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Could not update order"})
		return
	}

	c.JSON(http.StatusCreated, updatedOrder)
}

func (oh *OrderHandler) Delete(c *gin.Context) {
	oh.l.Info("Getting order ID...")
	orderID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		oh.l.Error("Error getting order ID", zap.Error(err))
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Could not delete order"})
		return
	}

	oh.l.Info("Deleting order from the database")
	err = oh.repository.Delete(uint(orderID))
	if err != nil {
		oh.l.Error("Error while deleting order", zap.Error(err))
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Could not delete order"})
		return
	}

	c.Status(http.StatusNoContent)
}
