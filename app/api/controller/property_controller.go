package controller

import (
	"esoft/app/domain"
	"esoft/app/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PropertyController struct {
	PropertyUsecase *usecase.PropertyUsecase
}

// CreateProperty godoc
// @Summary	Create of property
// @Tags Property
// @Accept json
// @Produce json
// @Param        data    body     domain.Property true  "scheme of property"
// @Success 200 {object} domain.Property
// @Failure 400 {object} domain.ErrorMessage
// @Failure 500 {object} domain.ErrorMessage
// @Router /properties [post]
func (u PropertyController) CreateProperty(c *gin.Context) {
	var property domain.Property

	err := c.BindJSON(&property)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorMessage{Title: "property schema", Body: err.Error()})
		return
	}
	//if property.FirstName == "" || property.LastName == "" || property.MiddleName == "" {
	//	c.JSON(http.StatusBadRequest, domain.ErrorMessage{Title: "property schema", Body: "Set first name, last name and middle name!"})
	//	return
	//}

	err = u.PropertyUsecase.CreateProperty(c, &property)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorMessage{
			Title: "Create new property",
			Body:  err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, property)
}

// GetProperties godoc
// @Summary	Get of properties
// @Tags Property
// @Accept json
// @Produce json
// @Success 200 {object} domain.Property
// @Failure 400 {object} domain.ErrorMessage
// @Failure 500 {object} domain.ErrorMessage
// @Router /properties [get]
func (u PropertyController) GetProperties(c *gin.Context) {
	properties, err := u.PropertyUsecase.GetProperties(c)

	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorMessage{
			Title: "Get all properties",
			Body:  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, properties)
}

// UpdateProperty godoc
// @Summary	Update of property
// @Tags Property
// @Accept json
// @Produce json
// @Param        data    body     domain.Property true  "scheme of property"
// @Success 200 {object} domain.Property
// @Failure 400 {object} domain.ErrorMessage
// @Failure 500 {object} domain.ErrorMessage
// @Router /properties [put]
func (u PropertyController) UpdateProperty(c *gin.Context) {
	var property domain.Property

	err := c.BindJSON(&property)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorMessage{Title: "property schema", Body: err.Error()})
		return
	}

	//if property.FirstName == "" || property.LastName == "" || property.MiddleName == "" {
	//	c.JSON(http.StatusBadRequest, domain.ErrorMessage{Title: "property schema", Body: "Set first name, last name and middle name!"})
	//	return
	//}
	err = u.PropertyUsecase.UpdateProperty(c, &property)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorMessage{
			Title: "Update of property",
			Body:  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, property)
}

// DeleteProperty godoc
// @Summary	Delete of property
// @Tags Property
// @Accept json
// @Produce json
// @Param        data    body     domain.Property true  "scheme of property"
// @Success 200 {object} domain.Property
// @Failure 400 {object} domain.ErrorMessage
// @Failure 500 {object} domain.ErrorMessage
// @Router /properties [delete]
func (u PropertyController) DeleteProperty(c *gin.Context) {
	var property domain.Property

	err := c.BindJSON(&property)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorMessage{Title: "property schema", Body: err.Error()})
		return
	}

	err = u.PropertyUsecase.DeleteProperty(c, &property)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorMessage{
			Title: "Delete of property",
			Body:  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, property)
}
