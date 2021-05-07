package common

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CustomErrorHandler(err error, c echo.Context) {
	log.Println(err.Error())

	var statusCode int
	var errResponse ErrorResponse
	switch errType := err.(type) {
	case *ValidationError:
		statusCode = http.StatusBadRequest
		errResponse = ErrorResponse{
			BaseResponse: &BaseResponse{
				Status:  http.StatusBadRequest,
				Message: "Invalid Field Argument",
			},
			Errors: errType.Errors,
		}
	case *Error:
		statusCode = ErrorCodeToHTTPStatusCode(errType.Code)
		errResponse = ErrorResponse{
			BaseResponse: &BaseResponse{
				Status:  statusCode,
				Message: ErrorCode(err),
			},
			Errors: []string{ErrorMessage(errType)},
		}
	default:
		statusCode = http.StatusInternalServerError
		errResponse = ErrorResponse{
			&BaseResponse{http.StatusInternalServerError, "Internal Server Error"},
			[]string{"Errors happening in our system. Maybe try again later?"},
		}
	}

	c.JSON(statusCode, &errResponse)
}
