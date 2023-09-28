package controller

import (
	"net/http"

	"github.com/alan890104/go-clean-arch-demo/domain"
	"github.com/gin-gonic/gin"
)

type BookController struct {
	bookUsecase domain.BookUsecase
}

func NewBookController(bookUseCase domain.BookUsecase) *BookController {
	return &BookController{
		bookUsecase: bookUseCase,
	}
}

func (ctrl *BookController) GetBook(c *gin.Context) {
	books, err := ctrl.bookUsecase.GetAll(c)
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
	book, err := ctrl.bookUsecase.GetById(c, c.Param("bookID"))
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
	if err := ctrl.bookUsecase.Store(c, &book); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, domain.Response{
			Msg: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, domain.Response{
		Data: book,
	})
}

func (ctrl *BookController) BorrowBook(c *gin.Context) {
	bookID := c.Param("bookID")
	if bookID == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, domain.Response{
			Msg: "book id is required",
		})
		return
	}
	if err := ctrl.bookUsecase.Borrow(c, bookID); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, domain.Response{
			Msg: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, domain.Response{
		Msg: "borrow success",
	})
}

func (ctrl *BookController) ReturnBook(c *gin.Context) {
	bookID := c.Param("bookID")
	if bookID == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, domain.Response{
			Msg: "book id is required",
		})
		return
	}
	if err := ctrl.bookUsecase.Return(c, bookID); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, domain.Response{
			Msg: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, domain.Response{
		Msg: "return success",
	})
}
