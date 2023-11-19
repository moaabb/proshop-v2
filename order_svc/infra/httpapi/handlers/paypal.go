package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/moaabb/ecommerce/order_svc/domain/order"
	"github.com/moaabb/ecommerce/order_svc/infra/utils"
	"go.uber.org/zap"
)

func (oh *OrderHandler) UpdateToPaid(c *gin.Context) {
	oh.l.Info("Getting order ID...")
	orderID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		oh.l.Error("Error getting order ID", zap.Error(err))
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Could not update order"})
		return
	}

	var paymentInfo order.PayPalPaymentDTO
	err = c.BindJSON(&paymentInfo)
	if err != nil {
		oh.l.Error("error binding dto", zap.Error(err))
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, "invalid request")
		return
	}

	o, err := oh.repository.GetById(uint(orderID))
	if err != nil {
		oh.l.Error("Error while fetching order", zap.Error(err))
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Could not fetch order"})
		return
	}

	err = oh.validatePaypalPayment(o, paymentInfo.ID)
	if err != nil {
		oh.l.Error("error updating order to paid", zap.Error(err))
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"error": "something went wrong, could not update order to paid",
		})
		return
	}

	var updatedOrder order.Order
	updatedOrder.PaymentID = paymentInfo.ID
	updatedOrder.PaymentStatus = paymentInfo.Status
	updatedOrder.PaymentUpdateTime = paymentInfo.UpdateTime.String()
	updatedOrder.PaymentEmailAddress = paymentInfo.Payer.EmailAddress
	updatedOrder.IsPaid = true
	updatedOrder.PaidAt = paymentInfo.CreateTime

	oh.l.Info("Updating order in the database")
	updatedOrder, err = oh.repository.UpdateToPaid(uint(orderID), updatedOrder)
	if err != nil {
		oh.l.Error("Error while updating order", zap.Error(err))
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Could not update order"})
		return
	}

	c.JSON(http.StatusOK, updatedOrder)
}

func (oh *OrderHandler) validatePaypalPayment(o order.Order, orderId string) error {
	token, err := utils.GeneratePaypalToken(oh.cfg)
	if err != nil {
		oh.l.Error("error generating paypal token", zap.Error(err))
		return err
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("https://%v/v2/checkout/orders/%s", oh.cfg.PaypalBaseUrl, orderId), nil)
	if err != nil {
		oh.l.Error("error mounting request for payment", zap.Error(err))
		return err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		oh.l.Error("error fetching payment data", zap.Error(err))
		return err
	}

	var paypalOrder order.PayPalGetOrderDto
	err = json.NewDecoder(resp.Body).Decode(&paypalOrder)
	if err != nil {
		oh.l.Error("error fetching payment data", zap.Error(err))
		return err
	}

	oh.l.Info(fmt.Sprintf("%v", paypalOrder))
	if paypalOrder.Status == "COMPLETED" && paypalOrder.PurchaseUnits[0].Amount.Value == o.TotalPrice {
		return nil
	}

	return errors.New("payment is not completed")

}
