package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sandronister/cassandra-go/internal/infra/web/dto"
	"github.com/sandronister/cassandra-go/internal/usecase"
)

type DriverTruckHandler struct {
	useCase *usecase.DriverUseCase
}

func NewDriverTruckHandler(useCase *usecase.DriverUseCase) *DriverTruckHandler {
	return &DriverTruckHandler{useCase: useCase}
}

func (h *DriverTruckHandler) GetDrivers(c *gin.Context) {
	drivers, err := h.useCase.GetDrivers()
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.Message{Message: err.Error()})
	}
	c.JSON(http.StatusOK, drivers)
}

func (h *DriverTruckHandler) GetDriver(c *gin.Context) {
	id := c.Param("id")
	driver, err := h.useCase.GetDriver(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.Message{Message: err.Error()})
	}
	c.JSON(http.StatusOK, driver)
}

func (h *DriverTruckHandler) CreateDriver(c *gin.Context) {
	driver := new(dto.DriversTruck)
	if err := c.Bind(driver); err != nil {
		c.JSON(http.StatusBadRequest, dto.Message{Message: err.Error()})
	}

	err := h.useCase.CreateDriver(driver)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.Message{Message: err.Error()})
	}
	c.JSON(http.StatusCreated, driver)
}

func (h *DriverTruckHandler) UpdateDriver(c *gin.Context) {
	id := c.Param("id")
	driver := new(dto.DriversTruck)
	driver.LicenseId = id

	if err := c.Bind(driver); err != nil {
		c.JSON(http.StatusBadRequest, dto.Message{Message: err.Error()})
	}

	err := h.useCase.UpdateDriver(driver)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.Message{Message: err.Error()})
	}
	c.JSON(http.StatusOK, driver)
}

func (h *DriverTruckHandler) DeleteDriver(c *gin.Context) {
	id := c.Param("id")
	err := h.useCase.DeleteDriver(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.Message{Message: err.Error()})
	}
	c.JSON(http.StatusNoContent, nil)
}
