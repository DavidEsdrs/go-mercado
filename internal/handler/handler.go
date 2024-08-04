package handler

import (
	"net/http"
	"strconv"

	"github.com/DavidEsdrs/go-mercado/internal/model"
	service "github.com/DavidEsdrs/go-mercado/internal/services"
	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	service *service.ProductService
}

func NewProductHandler(service *service.ProductService) *ProductHandler {
	return &ProductHandler{
		service: service,
	}
}

func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var product model.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.InsertProduct(&product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task"})
		return
	}

	c.JSON(http.StatusCreated, product)
}

func (h *ProductHandler) ReadProduct(c *gin.Context) {
	idAsString := c.Param("id")
	id, err := strconv.ParseUint(idAsString, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "given param is invalid",
		})
		return
	}

	product, err := h.service.ReadProduct(uint(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":          "Unable to create product",
			"internal_error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, product)
}

func (h *ProductHandler) ReadProducts(c *gin.Context) {
	products, err := h.service.ReadProducts()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":          "Unable to read products",
			"internal_error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, products)
}
