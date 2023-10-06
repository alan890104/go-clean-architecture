package controller

import (
	"net/http"

	"github.com/alan890104/go-clean-arch-demo/domain"
	"github.com/gin-gonic/gin"
)

type RecordController struct {
	recordUsecase domain.RecordUsecase
}

func NewRecordController(recordUseCase domain.RecordUsecase) *RecordController {
	return &RecordController{
		recordUsecase: recordUseCase,
	}
}

func (ctrl *RecordController) GetRecord(c *gin.Context) {
	records, err := ctrl.recordUsecase.GetAll(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, domain.Response{
		Data: records,
	})
}

func (ctrl *RecordController) GetRecordByUserID(c *gin.Context) {
	records, err := ctrl.recordUsecase.GetByUserId(c, c.Param("userID"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, domain.Response{
			Msg: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, domain.Response{
		Data: records,
	})
}
