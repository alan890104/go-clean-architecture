package controller

import (
	"net/http"

	"github.com/alan890104/go-clean-arch-demo/domain"
	"github.com/gin-gonic/gin"
)

type SignupController struct {
	signupUsecase domain.SignUpUsecase
}

func NewSignupController(signupUsecase domain.SignUpUsecase) *SignupController {
	return &SignupController{
		signupUsecase: signupUsecase,
	}
}

func (ctrl *SignupController) Signup(c *gin.Context) {
	var request domain.SignUpRequest

	if err := c.ShouldBind(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, domain.Response{
			Msg: err.Error(),
		})
		return
	}

	if err := ctrl.signupUsecase.Signup(c, &domain.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, domain.Response{
			Msg: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, domain.Response{
		Msg: "signup success",
	})
}
