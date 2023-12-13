package parser

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
)

type Parser interface {
	GetCurrency(request string, response string) (float64, error)
}

type currencyParser struct {
	Date        string             `json:"date"`
	TargetRub   map[string]float64 `json:"rub"`
	TargetUsd   map[string]float64 `json:"usd"`
	TargetEur   map[string]float64 `json:"eur"`
	TargetBtc   map[string]float64 `json:"btc"`
	TargetUsdt  map[string]float64 `json:"usdt"`
	TargetEth   map[string]float64 `json:"eth"`
	TargetXrp   map[string]float64 `json:"xrp"`
	TargetLtc   map[string]float64 `json:"ltc"`
	TargetAda   map[string]float64 `json:"ada"`
	TargetDot   map[string]float64 `json:"dot"`
	TargetDoge  map[string]float64 `json:"doge"`
	TargetBnb   map[string]float64 `json:"bnb"`
	TargetLink  map[string]float64 `json:"link"`
	TargetXlm   map[string]float64 `json:"xlm"`
	TargetSol   map[string]float64 `json:"sol"`
	TargetAtom  map[string]float64 `json:"atom"`
	TargetMatic map[string]float64 `json:"matic"`
	TargetUni   map[string]float64 `json:"uni"`
	TargetFil   map[string]float64 `json:"fil"`

	mu sync.RWMutex
}

func NewCurrencyParser() Parser {
	return &currencyParser{}
}

func (c *currencyParser) GetCurrency(request string, response string) (float64, error) {
	resp, err := http.Get(fmt.Sprintf("https://cdn.jsdelivr.net/gh/fawazahmed0/currency-api@1/latest/currencies/%s.json", request))
	if err != nil {
		log.Println("No response from request")
		return 0, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	c.mu.Lock()
	var parser currencyParser
	if err := json.Unmarshal(body, &parser); err != nil {
		log.Println("Can not unmarshal JSON")
		return 0, err
	}
	c.mu.Unlock()

	currencyMap := c.typeCurrencySwitching(parser)
	if currencyMap == nil {
		return 0, ErrNilMap
	}
	c.mu.RLock()
	data, exist := currencyMap[response]
	if !exist {
		return 0, ErrNilMap
	}
	c.mu.RUnlock()

	return data, nil
}

func (c *currencyParser) typeCurrencySwitching(cp currencyParser) map[string]float64 {
	switch {
	case cp.TargetEur != nil:
		return cp.TargetEur
	case cp.TargetUsd != nil:
		return cp.TargetUsd
	case cp.TargetBtc != nil:
		return cp.TargetBtc
	case cp.TargetRub != nil:
		return cp.TargetRub
	case cp.TargetUsdt != nil:
		return cp.TargetUsdt
	case cp.TargetEth != nil:
		return cp.TargetEth
	case cp.TargetXrp != nil:
		return cp.TargetXrp
	case cp.TargetLtc != nil:
		return cp.TargetLtc
	case cp.TargetAda != nil:
		return cp.TargetAda
	case cp.TargetDot != nil:
		return cp.TargetDot
	case cp.TargetDoge != nil:
		return cp.TargetDoge
	case cp.TargetBnb != nil:
		return cp.TargetBnb
	case cp.TargetLink != nil:
		return cp.TargetLink
	case cp.TargetXlm != nil:
		return cp.TargetXlm
	case cp.TargetSol != nil:
		return cp.TargetSol
	case cp.TargetAtom != nil:
		return cp.TargetAtom
	case cp.TargetMatic != nil:
		return cp.TargetMatic
	case cp.TargetUni != nil:
		return cp.TargetUni
	case cp.TargetFil != nil:
		return cp.TargetFil
	default:
		return nil
	}
}
