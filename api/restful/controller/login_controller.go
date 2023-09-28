package controller

import (
	"net/http"

	"github.com/alan890104/go-clean-arch-demo/bootstrap"
	"github.com/alan890104/go-clean-arch-demo/domain"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type LoginController struct {
	loginUsecase domain.LoginUsecase
	env          *bootstrap.Env
}

func NewLoginController(loginUsecase domain.LoginUsecase, env *bootstrap.Env) *LoginController {
	return &LoginController{
		loginUsecase: loginUsecase,
		env:          env,
	}
}

func (ctrl *LoginController) Login(c *gin.Context) {
	var request domain.LoginRequest

	if err := c.ShouldBind(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, domain.Response{
			Msg: err.Error(),
		})
		return
	}

	user, err := ctrl.loginUsecase.GetUserByEmail(c, request.Email)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, domain.Response{
			Msg: "user not found with email: " + request.Email,
		})
		return
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)) != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, domain.Response{
			Msg: "invalid password",
		})
		return
	}

	accessToken, err := ctrl.loginUsecase.CreateAccessToken(c, ctrl.env.JWT.AccessTokenSecret, ctrl.env.JWT.AccessTokenExpiry)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, domain.Response{
			Msg: err.Error(),
		})
		return
	}

	refreshToken, err := ctrl.loginUsecase.CreateRefreshToken(c, ctrl.env.JWT.RefreshTokenSecret, ctrl.env.JWT.RefreshTokenExpiry)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, domain.Response{
			Msg: err.Error(),
		})
		return
	}

	loginResponse := domain.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	c.JSON(http.StatusOK, domain.Response{
		Data: loginResponse,
	})
}
