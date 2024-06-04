package handler

import (
	commonentity "go-server/internal/common/cmentity"
	"go-server/internal/usecase/usergetinfo"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserGetInfoHandler interface {
	UserGetAttendanceList(c *gin.Context)
}

type userGetInfoHandler struct {
	UserGetInfoUC usergetinfo.UseCase
}

func NewUserGetInfoHandler(uuc usergetinfo.UseCase) UserGetInfoHandler {
	return &userGetInfoHandler{
		UserGetInfoUC: uuc,
	}
}

// UserCreateAction	goc
// UserCreateAction	API
//
//	@Summary		User Gets Attendance List
//	@Description	User Gets Attendance List
//	@Tags			employee
//	@Accept      	json
//	@Security		ApiKeyAuth
//	@Produce		json
//	@Router			/v1/employees/{employee_id}/attendances [get]
//	@Param			employee_id	path		string	true	"Employee ID"
//	@Success		200					{object}	[]entity.Attendance
//	@Failure		500
func (h *userGetInfoHandler) UserGetAttendanceList(c *gin.Context) {
	result, err := h.UserGetInfoUC.UserGetAttendanceList(c, commonentity.ID(c.Param("employee_id")))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, result)
}
