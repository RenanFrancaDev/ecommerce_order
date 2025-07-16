package errors

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"ecommerce_order/internal/infrastructure/adapters/http/dto"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ErrorHandler struct{}

func NewErrorHandler() *ErrorHandler {
	return &ErrorHandler{}
}

func (h *ErrorHandler) HandleError(ctx *gin.Context, err error) {
	if validationErr, ok := err.(validator.ValidationErrors); ok {
		h.handleValidationErrors(ctx, validationErr)
		return
	}

	if strings.Contains(err.Error(), "rabbitmq") || strings.Contains(err.Error(), "amqp") ||
		strings.Contains(err.Error(), "queue") || strings.Contains(err.Error(), "publish") {
		h.handleInfrastructureError(ctx, err)
		return
	}

	h.handleUnexpectedError(ctx, err)
}

func (h *ErrorHandler) handleValidationErrors(ctx *gin.Context, validationErrors validator.ValidationErrors) {
	var errors []dto.ValidationErrorDetail

	for _, err := range validationErrors {
		errors = append(errors, dto.ValidationErrorDetail{
			Field:   err.Field(),
			Message: h.getValidationMessage(err),
		})
	}

	ctx.JSON(http.StatusBadRequest, dto.StandardErrorResponse{
		Message:   "Validation failed",
		Status:    "error",
		ErrorCode: "VALIDATION_ERROR",
		Details:   errors,
	})
}

func (h *ErrorHandler) handleInfrastructureError(ctx *gin.Context, err error) {
	log.Printf("[feature:order] [msg:infrastructure error] [error: %v]", err)

	ctx.JSON(http.StatusServiceUnavailable, dto.StandardErrorResponse{
		Message:   "Unable to process order at this time. Please try again later.",
		Status:    "error",
		ErrorCode: "SERVICE_UNAVAILABLE",
	})
}

func (h *ErrorHandler) handleUnexpectedError(ctx *gin.Context, err error) {
	log.Printf("[feature:order] [msg:unexpected error] [error: %v]", err)

	ctx.JSON(http.StatusInternalServerError, dto.StandardErrorResponse{
		Message:   "An internal error occurred. Please try again.",
		Status:    "error",
		ErrorCode: "INTERNAL_ERROR",
	})
}

func (h *ErrorHandler) getValidationMessage(err validator.FieldError) string {
	switch err.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", err.Field())
	case "email":
		return fmt.Sprintf("%s must be a valid email address", err.Field())
	case "min":
		return fmt.Sprintf("%s must be at least %s", err.Field(), err.Param())
	case "max":
		return fmt.Sprintf("%s must be at most %s", err.Field(), err.Param())
	default:
		return fmt.Sprintf("%s is invalid", err.Field())
	}
}
