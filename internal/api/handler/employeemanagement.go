package handler

import (
	"go-server/internal/entity"
	"go-server/internal/usecase/employeemanagement"
	"net/http"

	"github.com/gin-gonic/gin"
)

type EmployeemManagementHandler interface {
	CreateNewEmployee(c *gin.Context)
}

type employeemManagementHandler struct {
	EmployeeUseCase employeemanagement.UseCase
}

func NewEmployeeManagementHandler(euc employeemanagement.UseCase) EmployeemManagementHandler {
	return &employeemManagementHandler{
		EmployeeUseCase: euc,
	}
}

// CreateNewEmployee	godoc
// CreateNewEmployee	API
//
//	@Summary		Create New Employee
//	@Description	Create New Employee
//	@Tags			employee
//	@Accept      	json
//	@Security		ApiKeyAuth
//	@Produce		json
//	@Router			/admin/employees [post]
//	@Param			request body entity.Employee	true	"Create an employee"
//	@Success		200					{object}	entity.Employee
//	@Failure		500
func (h *employeemManagementHandler) CreateNewEmployee(c *gin.Context) {
	var em entity.Employee
	if err := c.ShouldBind(&em); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
	}

	result, err := h.EmployeeUseCase.CreateNewEmployee(c, &em)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, result)
}
