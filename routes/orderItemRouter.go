package routes

import (
	"github.com/gin-gonic/gin"
	controller "restaurant-api/controllers"
)

func OrderItemRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/order_items", controller.GetOrderItems())
	incomingRoutes.GET("/order_items/:order_items_id", controller.GetOrderItem())
	incomingRoutes.POST("/order_items", controller.CreateOrderItem())
	incomingRoutes.PATCH("/order_items/:orde_items_id", controller.UpdateOrderItem())
}
