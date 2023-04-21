package shorter

import (
	"id/projects/url-shortener/helper"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type handler struct {
	service Service
}

func NewHandler(service Service) *handler {
	return &handler{service}
}

func (h *handler) GenerateShorterUrl(c *gin.Context) {
	var input UrlInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Generate Shorten URL failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	shortUrl, err := h.service.GenerateShorterUrl(input)

	if err != nil {
		response := helper.APIResponse(err.Error(), http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := FormatShortUrl(shortUrl)
	response := helper.APIResponse("Generate Shorten URL success", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}

func (h *handler) SaveShorterUrl(c *gin.Context) {
	var input ShorterUrlInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Save Shorten URL failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	shortUrl, err := h.service.SaveShorterUrl(input)

	if err != nil {
		response := helper.APIResponse(err.Error(), http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := FormatShortUrl(shortUrl)
	response := helper.APIResponse("Save Shorten URL success", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}

func (h *handler) RedirectShorterUrl(c *gin.Context) {
	backHalf := c.Param("backHalf")

	shortUrl, err := h.service.FindShorterUrl(backHalf)
	if err != nil {
		response := helper.APIResponse("Redirect Shorten URL failed", http.StatusNotFound, "error", nil)
		c.JSON(http.StatusNotFound, response)
		return
	}

	redirectURL := shortUrl.DestinationUrl
	if !strings.HasPrefix(redirectURL, "http://") && !strings.HasPrefix(redirectURL, "https://") {
		redirectURL = "http://" + redirectURL
	}
	c.Redirect(http.StatusMovedPermanently, redirectURL)
}
