package controller

import (
	"net/http"

	"github.com/alan890104/go-clean-arch-demo/domain"
	"github.com/gin-gonic/gin"
)

type BookController struct {
	bookUsecase   domain.BookUsecase
	recordUsecase domain.RecordUsecase
}

func NewBookController(bookUseCase domain.BookUsecase, recordUsecase domain.RecordUsecase) *BookController {
	return &BookController{
		bookUsecase:   bookUseCase,
		recordUsecase: recordUsecase,
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
	var borrowBookRequest domain.BorrowBookRequest
	if err := c.ShouldBindJSON(&borrowBookRequest); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, domain.Response{
			Msg: err.Error(),
		})
		return
	}
	if err := ctrl.bookUsecase.Borrow(c, &borrowBookRequest); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, domain.Response{
			Msg: err.Error(),
		})
		return
	}
	if err := ctrl.recordUsecase.Store(c, &domain.StoreRecordRequest{
		UserId: borrowBookRequest.UserId,
		BookId: borrowBookRequest.BookId,
	}); err != nil {
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
	var returnBookRequest domain.ReturnBookRequest
	if err := c.ShouldBindJSON(&returnBookRequest); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, domain.Response{
			Msg: err.Error(),
		})
		return
	}
	if err := ctrl.bookUsecase.Return(c, &returnBookRequest); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, domain.Response{
			Msg: err.Error(),
		})
		return
	}
	if err := ctrl.recordUsecase.UpdateEndDateByBookId(c, returnBookRequest.BookId); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, domain.Response{
			Msg: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, domain.Response{
		Msg: "return success",
	})
}

func (ctrl *BookController) UpdateBookByID(c *gin.Context) {
	var book domain.UpdateBookRequest
	if err := c.ShouldBindJSON(&book); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, domain.Response{
			Msg: err.Error(),
		})
		return
	}
	updatedBook, err := ctrl.bookUsecase.UpdateById(c, c.Param("bookID"), &book)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, domain.Response{
			Msg: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, updatedBook)
}

func (ctrl *BookController) DeleteBookByID(c *gin.Context) {
	if err := ctrl.bookUsecase.DeleteById(c, c.Param("bookID")); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, domain.Response{
			Msg: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, domain.Response{
		Msg: "delete success",
	})
}
