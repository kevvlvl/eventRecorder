package server

import (
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (s *Server) Echo() error {

	s.EchoServer = echo.New()

	if s.EchoServer == nil {
		return errors.New("echo server not created")
	}

	s.initRoutes()

	return nil
}

func (s *Server) initRoutes() {

	s.EchoServer.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "Up!")
	})

	s.EchoServer.POST("/finance/receipt", func(c echo.Context) error {
		r := new(FinanceReceipt)
		if err := c.Bind(r); err != nil {
			return err
		}

		return c.JSON(http.StatusCreated, r)
	})

	s.EchoServer.Logger.Fatal(s.EchoServer.Start(":8011"))
}
