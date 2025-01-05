package handler

import (
	"net/http"

	entity "github.com/Genarodaniel/order-system/internal/entity"
	"github.com/Genarodaniel/order-system/internal/infra/api/response"
	usecase "github.com/Genarodaniel/order-system/internal/usecase"
	"github.com/gin-gonic/gin"
)

type OrderHandlerInterface interface {
	Create(ctx *gin.Context)
}

type ApiOrderHandler struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewApiOrderHandler(
	OrderRepository entity.OrderRepositoryInterface,
) *ApiOrderHandler {
	return &ApiOrderHandler{
		OrderRepository: OrderRepository,
	}
}

func (h *ApiOrderHandler) Create(ctx *gin.Context) {
	dto := usecase.CreateOrderInputDTO{}
	if err := ctx.ShouldBindJSON(&dto); err != nil {
		response.HasError(ctx, http.StatusBadRequest, err)
		return
	}

	orderUseCase := usecase.NewOrderUseCase(h.OrderRepository)
	output, err := orderUseCase.CreateOrder(dto)
	if err != nil {
		response.HasError(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusAccepted, output)
}
