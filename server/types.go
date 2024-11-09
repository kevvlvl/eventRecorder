package server

import "github.com/labstack/echo/v4"

type Server struct {
	EchoServer *echo.Echo
}

type FinanceReceipt struct {
	Name          string `json:"name"`
	Amount        string `json:"amount"`
	PaymentMethod string `json:"payment_method"`
}
