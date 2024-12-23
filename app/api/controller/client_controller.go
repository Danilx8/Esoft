package controller

import (
	"esoft/app/domain"
	"esoft/app/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ClientController struct {
	ClientUsecase *usecase.ClientUsecase
}

// CreateClient godoc
// @Summary	Create of client
// @Tags Client
// @Accept json
// @Produce json
// @Param        data    body     domain.Client true  "scheme of client"
// @Success 200 {object} domain.Client
// @Failure 400 {object} domain.ErrorMessage
// @Failure 500 {object} domain.ErrorMessage
// @Router /clients [post]
func (u ClientController) CreateClient(c *gin.Context) {
	var client domain.Client

	err := c.BindJSON(&client)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorMessage{Title: "client schema", Body: err.Error()})
		return
	}
	if client.Email == "" && client.PhoneNumber == "" {
		c.JSON(http.StatusBadRequest, domain.ErrorMessage{Title: "client schema", Body: "Set email or phone number!"})
		return
	}

	err = u.ClientUsecase.CreateClient(c, &client)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorMessage{
			Title: "Create new client",
			Body:  err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, client)
}

// GetClients godoc
// @Summary	Get of clients
// @Tags Client
// @Accept json
// @Produce json
// @Success 200 {object} domain.Client
// @Failure 400 {object} domain.ErrorMessage
// @Failure 500 {object} domain.ErrorMessage
// @Router /clients [get]
func (u ClientController) GetClients(c *gin.Context) {
	clients, err := u.ClientUsecase.GetClients(c)

	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorMessage{
			Title: "Get all clients",
			Body:  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, clients)
}

// UpdateClient godoc
// @Summary	Update of client
// @Tags Client
// @Accept json
// @Produce json
// @Param        data    body     domain.Client true  "scheme of client"
// @Success 200 {object} domain.Client
// @Failure 400 {object} domain.ErrorMessage
// @Failure 500 {object} domain.ErrorMessage
// @Router /clients [put]
func (u ClientController) UpdateClient(c *gin.Context) {
	var client domain.Client

	err := c.BindJSON(&client)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorMessage{Title: "client schema", Body: err.Error()})
		return
	}

	if client.Email == "" && client.PhoneNumber == "" {
		c.JSON(http.StatusBadRequest, domain.ErrorMessage{Title: "client schema", Body: "Set email or phone number!"})
		return
	}
	err = u.ClientUsecase.UpdateClient(c, &client)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorMessage{
			Title: "Update of client",
			Body:  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, client)
}

// DeleteClient godoc
// @Summary	Delete of client
// @Tags Client
// @Accept json
// @Produce json
// @Param        data    body     domain.Client true  "scheme of client"
// @Success 200 {object} domain.Client
// @Failure 400 {object} domain.ErrorMessage
// @Failure 500 {object} domain.ErrorMessage
// @Router /clients [delete]
func (u ClientController) DeleteClient(c *gin.Context) {
	var client domain.Client

	err := c.BindJSON(&client)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorMessage{Title: "client schema", Body: err.Error()})
		return
	}

	err = u.ClientUsecase.DeleteClient(c, &client)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorMessage{
			Title: "Delete of client",
			Body:  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, client)
}

// SearchClient godoc
// @Summary	Search of client
// @Tags Client
// @Accept json
// @Produce json
// @Param        data    body     domain.Client true  "scheme of client"
// @Success 200 {object} domain.Client
// @Failure 400 {object} domain.ErrorMessage
// @Failure 500 {object} domain.ErrorMessage
// @Router /clients/search [post]
func (u ClientController) SearchClients(c *gin.Context) {
	var client domain.Client

	err := c.BindJSON(&client)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorMessage{Title: "client schema", Body: err.Error()})
		return
	}

	clients, err := u.ClientUsecase.SearchSimilaryClients(c, &client)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorMessage{
			Title: "Search of client",
			Body:  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, clients)
}
