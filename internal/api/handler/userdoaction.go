package handler

import (
	commonentity "go-server/internal/common/cmentity"
	"go-server/internal/entity"
	"go-server/internal/usecase/userdoaction"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserDoActionHandler interface {
	UserCreateAction(c *gin.Context)
}

type userDoActionHandler struct {
	UserUseCase userdoaction.UseCase
}

func NewUserDoActionHandler(uuc userdoaction.UseCase) UserDoActionHandler {
	return &userDoActionHandler{
		UserUseCase: uuc,
	}
}

// UserCreateAction	goc
// UserCreateAction	API
//
//	@Summary		User Creates Action
//	@Description	User Creates Action
//	@Tags			employee
//	@Accept      	json
//	@Security		ApiKeyAuth
//	@Produce		json
//	@Router			/v1/employees/{employee_id}/actions [post]
//	@Param			employee_id	path		string	true	"Employee ID"
//	@Param			request body entity.Action	true	"employee creates an action"
//	@Success		200					{object}	entity.Action
//	@Failure		500
func (h *userDoActionHandler) UserCreateAction(c *gin.Context) {
	var em entity.Action
	if err := c.ShouldBindJSON(&em); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)

		return
	}

	em.EmployeeID = commonentity.ID(c.Param("employee_id"))
	result, err := h.UserUseCase.UserCreateAction(c, &em)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)

		return
	}

	c.JSON(http.StatusOK, result)
}
