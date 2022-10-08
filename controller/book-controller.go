package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/A-u-usman/golang_gin_clean_api/dto"
	"github.com/A-u-usman/golang_gin_clean_api/entity"
	"github.com/A-u-usman/golang_gin_clean_api/helper"
	"github.com/A-u-usman/golang_gin_clean_api/service"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type BookController interface {
	All(context *gin.Context)
	FindByID(context *gin.Context)
	Insert(context *gin.Context)
	Update(contect *gin.Context)
	Delete(context *gin.Context)
}

type bookController struct {
	bookService service.BookService
	jwtService  service.JWTService
}

func NewBookController(bookServ service.BookService, jwtServ service.JWTService) BookController {
	return &bookController{
		bookService: bookServ,
		jwtService:  jwtServ,
	}
}

func (c *bookController) All(context *gin.Context) {
	var books []entity.Book = c.bookService.All()
	res := helper.BuildResponse(true, "OK", books)
	context.JSON(http.StatusOK, res)
}

func (c *bookController) FindByID(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		res := helper.BuildErrorResponse("No param id w found", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	var book entity.Book = c.bookService.FindByID(id)
	if (book == entity.Book{}) {
		res := helper.BuildErrorResponse("Data not found", "No data with given id", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		res := helper.BuildResponse(true, "OK", book)
		context.JSON(http.StatusOK, res)
	}
}

func (c *bookController) Insert(context *gin.Context) {
	var bookCreateDTO dto.BookCreateDTO
	errDTO := context.ShouldBind(&bookCreateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	} else {
		authHeader := context.GetHeader("Authorization")
		userID := c.getUserIDByToken(authHeader)
		convertedUserID, err := strconv.ParseUint(userID, 10, 64)
		if err == nil {
			bookCreateDTO.UserID = convertedUserID
		}
		result := c.bookService.Insert(bookCreateDTO)
		response := helper.BuildResponse(true, "OK", result)
		context.JSON(http.StatusCreated, response)
	}
}

func (c *bookController) Update(context *gin.Context) {
	var bookUpdatedDTO dto.BookUpdateDTO
	errDTO := context.ShouldBind(&bookUpdatedDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
		return
	}

	authHeader := context.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])
	if c.bookService.IsAllowedToEdit(userID, bookUpdatedDTO.ID) {
		id, errID := strconv.ParseUint(userID, 10, 64)
		if errID == nil {
			bookUpdatedDTO.UserID = id
		}
		result := c.bookService.Update(bookUpdatedDTO)
		response := helper.BuildResponse(true, "OK", result)
		context.JSON(http.StatusOK, response)
	} else {
		response := helper.BuildErrorResponse("You dont have permission", "you are not the owner", helper.EmptyObj{})
		context.JSON(http.StatusForbidden, response)
	}
}

func (c *bookController) Delete(context *gin.Context) {
	var book entity.Book
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		response := helper.BuildErrorResponse("Failed to get id", "No param id was found", helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, response)
	}
	book.ID = id
	authHeader := context.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])
	if c.bookService.IsAllowedToEdit(userID, book.ID) {
		c.bookService.Delete(book)
		res := helper.BuildResponse(true, "Deleted", helper.EmptyObj{})
		context.JSON(http.StatusOK, res)
	} else {
		response := helper.BuildErrorResponse("You dont have permission", "you are not the owner", helper.EmptyObj{})
		context.JSON(http.StatusForbidden, response)
	}

}

func (c *bookController) getUserIDByToken(token string) string {
	aToken, err := c.jwtService.ValidateToken(token)
	if err != nil {
		panic(err.Error())
	}
	claims := aToken.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	return id
}
