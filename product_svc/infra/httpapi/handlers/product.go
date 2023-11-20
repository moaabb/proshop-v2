package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgconn"

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
	ph.l.Info("getting page number...")
	page := c.Query("pageNumber")
	pn, err := strconv.Atoi(page)
	if err != nil {
		pn = 1
	}

	ph.l.Info("Fetching products on database")
	products, err := ph.repository.GetAll(pn)
	if err != nil {
		ph.l.Error("error while fetching products", zap.Error(err))
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": "could not fetch products"})
		return
	}

	products.Page = pn

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
	err := c.BindJSON(&p)
	if err != nil {
		ph.l.Error("error binding dto", zap.Error(err))
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, "invalid request")
		return
	}

	ph.l.Info("Creating product on database")
	product, err := ph.repository.Create(p)
	if err != nil {
		ph.l.Error("error while creating product", zap.Error(err))
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, "could not create product")
		return
	}

	c.JSON(http.StatusCreated, product)
}

func (ph *ProductHandler) CreateReview(c *gin.Context) {
	var r product.Review
	err := c.BindJSON(&r)
	if err != nil {
		ph.l.Error("error binding dto", zap.Error(err))
		c.AbortWithStatusJSON(http.StatusBadRequest, "invalid request")
		return
	}

	ph.l.Info("getting user info")
	userId := c.GetUint("userId")
	if userId == 0 {
		ph.l.Error("could not get user identification")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "could not get user info",
		})
		return
	}

	r.UserID = userId

	ph.l.Info("Creating review on database")
	review, err := ph.repository.CreateReview(r)
	if err != nil {
		var pgErr *pgconn.PgError
		errors.As(err, &pgErr)
		ph.l.Error("error while creating review", zap.Error(pgErr))
		if pgErr.ConstraintName == "reviews_unique" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "Product already reviewed by you",
			})
			return
		} else {
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, "could not create review")
			return
		}
	}

	c.JSON(http.StatusCreated, review)
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
	err = c.BindJSON(&p)
	if err != nil {
		ph.l.Error("error binding dto", zap.Error(err))
		c.AbortWithStatusJSON(http.StatusBadRequest, "invalid request")
		return
	}

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
