package handler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"strconv"
	"web-currency-parser/internal/parser"
)

type currencyHandler struct {
	parser parser.Parser
}

func NewCurrencyHandler() *currencyHandler {
	return &currencyHandler{
		parser: parser.NewCurrencyParser(),
	}
}

func (h *currencyHandler) CreateRequest(c *fiber.Ctx) error {
	reqCurrency := c.FormValue("currencyFrom")
	respCurrency := c.FormValue("currencyTo")
	amountCurrency := c.FormValue("amountFrom")

	if reqCurrency == "" && respCurrency == "" && amountCurrency == "" {
		log.Println("some param is zero value")
	}

	e := struct {
		Error string `json:"error"`
	}{}
	amount, err := strconv.ParseFloat(amountCurrency, 64)
	if err != nil {
		log.Println(err)
		e.Error = "invalid syntax"
		c.Status(fiber.StatusBadRequest)
		return c.JSON(e)
	}

	course, err := h.parser.GetCurrency(reqCurrency, respCurrency)
	if err != nil {
		log.Println(err)
		e.Error = err.Error()

		c.Status(fiber.StatusInternalServerError)
		return c.JSON(e)
	}

	totalAmount := amount * course
	total := fmt.Sprintf("%.2f", totalAmount)

	data := struct {
		Amount string `json:"amount"`
	}{
		Amount: total,
	}

	c.Status(fiber.StatusOK)
	return c.JSON(data)
}
