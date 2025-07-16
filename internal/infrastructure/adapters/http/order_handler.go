package http

import (
	"net/http"

	"ecommerce_order/internal/application/usecase"
	"ecommerce_order/internal/infrastructure/adapters/http/dto"
	"ecommerce_order/internal/infrastructure/adapters/http/errors"
	"ecommerce_order/internal/infrastructure/adapters/http/mapper"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	placeOrderUseCase usecase.PlaceOrderUseCase
	mapper            *mapper.OrderMapper
	errorHandler      *errors.ErrorHandler
}

func NewOrderHandler(placeOrderUseCase usecase.PlaceOrderUseCase) *OrderHandler {
	return &OrderHandler{
		placeOrderUseCase: placeOrderUseCase,
		mapper:            mapper.NewOrderMapper(),
		errorHandler:      errors.NewErrorHandler(),
	}
}

func (h *OrderHandler) Create(ctx *gin.Context) {
	var request dto.CreateOrderRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		h.errorHandler.HandleError(ctx, err)
		return
	}

	order := h.mapper.ToEntity(request)

	if err := h.placeOrderUseCase.Execute(&order); err != nil {
		h.errorHandler.HandleError(ctx, err)
		return
	}

	response := h.mapper.ToCreateOrderResponse(order)
	ctx.JSON(http.StatusCreated, response)
}
