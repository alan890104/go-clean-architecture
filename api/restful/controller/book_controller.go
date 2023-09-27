package controller

import (
	"net/http"

	"github.com/alan890104/go-clean-arch-demo/domain"
	"github.com/gin-gonic/gin"
)

type BookController struct {
	bookUseCase domain.BookUsecase
}

func NewBookController(bookUseCase domain.BookUsecase) *BookController {
	return &BookController{
		bookUseCase: bookUseCase,
	}
}

func (ctrl *BookController) GetBook(c *gin.Context) {
	books, err := ctrl.bookUseCase.GetAll(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, domain.Response{
		Data: books,
	})
}

func (ctrl *BookController) GetBookByID(c *gin.Context) {
	book, err := ctrl.bookUseCase.GetById(c, c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, domain.Response{
			Msg: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, book)
}

func (ctrl *BookController) CreateBook(c *gin.Context) {
	var book domain.StoreBookRequest
	if err := c.ShouldBindJSON(&book); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, domain.Response{
			Msg: err.Error(),
		})
		return
	}
	if err := ctrl.bookUseCase.Store(c, &book); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, domain.Response{
			Msg: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, domain.Response{
		Data: book,
	})
}
