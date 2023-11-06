package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/moaabb/ecommerce/product_svc/domain/product"
	"github.com/moaabb/ecommerce/product_svc/infra/database/productdb"
	"go.uber.org/zap"
)

type ProductHandler struct {
	repository product.Repository
	l          *zap.Logger
}

func NewHandler(repo *productdb.Repository, z *zap.Logger) *ProductHandler {
	return &ProductHandler{
		repository: repo,
		l:          z,
	}

}

func (ph *ProductHandler) GetAll(c *gin.Context) {
	ph.l.Info("Fetching products on database")
	products, err := ph.repository.GetAll()
	if err != nil {
		ph.l.Error("error while fetching products", zap.Error(err))
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": "could not fetch products"})
		return
	}

	c.JSON(http.StatusOK, products)
}

func (ph *ProductHandler) GetTopProducts(c *gin.Context) {
	ph.l.Info("Fetching products on database")
	products, err := ph.repository.GetTopProducts()
	if err != nil {
		ph.l.Error("error while fetching products", zap.Error(err))
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": "could not fetch products"})
		return
	}

	c.JSON(http.StatusOK, products)
}

func (ph *ProductHandler) GetById(c *gin.Context) {
	ph.l.Info("getting product id...")
	productId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		ph.l.Error("error getting productId", zap.Error(err))
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": "could not fetch products"})
		return
	}

	ph.l.Info("Fetching product on database")
	product, err := ph.repository.GetById(uint(productId))
	if err != nil {
		ph.l.Error("error while fetching product", zap.Error(err))
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, "could not fetch product")
		return
	}

	c.JSON(http.StatusOK, product)
}

func (ph *ProductHandler) Create(c *gin.Context) {
	var p product.Product
	c.BindJSON(&p)

	ph.l.Info("Creating product on database")
	product, err := ph.repository.Create(p)
	if err != nil {
		ph.l.Error("error while creating product", zap.Error(err))
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, "could not create product")
		return
	}

	c.JSON(http.StatusCreated, product)
}
func (ph *ProductHandler) Update(c *gin.Context) {
	ph.l.Info("getting product id...")
	productId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		ph.l.Error("error getting productId", zap.Error(err))
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": "could not fetch products"})
		return
	}

	var p product.Product
	c.BindJSON(&p)

	ph.l.Info("updating product on database")
	product, err := ph.repository.Update(uint(productId), p)
	if err != nil {
		ph.l.Error("error while updating product", zap.Error(err))
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, "could not update product")
		return
	}

	c.JSON(http.StatusOK, product)
}
func (ph *ProductHandler) Delete(c *gin.Context) {
	ph.l.Info("getting product id...")
	productId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		ph.l.Error("error getting productId", zap.Error(err))
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": "could not fetch products"})
		return
	}

	ph.l.Info("deleting product on database")
	err = ph.repository.Delete(uint(productId))
	if err != nil {
		ph.l.Error("error while deleting product", zap.Error(err))
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, "could not delete product")
		return
	}

	c.Status(http.StatusNoContent)
}
