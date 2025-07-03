package metatrader

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/go-gota/gota/dataframe"
	"github.com/mitchellh/mapstructure"
	"github.com/nenjotsu/go-metatrader/metatrader/actiontype"
	"github.com/nenjotsu/go-metatrader/metatrader/models"
	"github.com/nenjotsu/go-metatrader/metatrader/timeframes"
	"github.com/nenjotsu/go-metatrader/metatrader/utils"
)

type MetaTrader struct {
	Host              string
	RealVolume        bool
	Debug             bool
	AuthorizationCode string
	InstrumentLookup  []string
	API               *MTFunctions
}

func NewMetaTraderClient(host string, port int, realVolume bool, debug bool, authorizationCode string, instrumentLookup []string) *MetaTrader {
	if debug {
		log.SetFlags(log.LstdFlags | log.Lmicroseconds)
		log.SetOutput(os.Stdout)
	}

	api := NewMTFunctions(host, port, debug, instrumentLookup, authorizationCode)

	return &MetaTrader{
		Host:              host,
		RealVolume:        realVolume,
		Debug:             debug,
		AuthorizationCode: authorizationCode,
		InstrumentLookup:  instrumentLookup,
		API:               api,
	}
}

func (mt *MetaTrader) Connect() {
	mt.API.Connect()
}

func (mt *MetaTrader) Disconnect() {
	mt.API.Disconnect()
}

func (mt *MetaTrader) GetCurrentPrice(symbol string) (*models.CurrentPrice, error) {
	// time, bid, ask
	response, err := mt.API.SendCommand("TICK|symbol=" + symbol)
	if err != nil {
		return nil, err
	}

	var tickEvent *models.TickEvent

	err = mapstructure.Decode(response, &tickEvent)
	if err != nil {
		return nil, fmt.Errorf("error decoding response: %s", err)
	}

	fmt.Println(tickEvent)

	if tickEvent.Data.Tick == nil {
		return nil, fmt.Errorf("no tick data found")
	}

	currentPrice := &models.CurrentPrice{
		Symbol: tickEvent.Data.Symbol,
		Tick: models.Tick{
			Timestamp: tickEvent.Data.Tick[0].(string),
			Bid:       tickEvent.Data.Tick[1].(float64),
			Ask:       tickEvent.Data.Tick[2].(float64),
		},
	}

	fmt.Println(currentPrice)

	return currentPrice, nil
}

func (mt *MetaTrader) GetOrders() (*models.Orders, error) {
	response, err := mt.API.SendCommand("ORDERS")
	if err != nil {
		return nil, err
	}

	responseMap := response.(map[string]interface{})
	dataMap := responseMap["orders"].([]interface{})

	if len(dataMap) == 0 {
		return nil, fmt.Errorf("no order data found")
	}

	var orders *models.Orders

	err = mapstructure.Decode(dataMap, &orders)
	if err != nil {
		return nil, fmt.Errorf("error decoding response: %s", err)
	}

	return orders, nil
}
func (mt *MetaTrader) GetOrdersBySymbol(symbol string) (*models.Orders, error) {
	response, err := mt.API.SendCommand("ORDERS|symbol=" + symbol)
	if err != nil {
		return nil, err
	}

	responseMap := response.(map[string]interface{})
	dataMap := responseMap["orders"].([]interface{})

	if len(dataMap) == 0 {
		return nil, fmt.Errorf("no order data found for symbol %s", symbol)
	}

	var orders *models.Orders

	err = mapstructure.Decode(dataMap, &orders)
	if err != nil {
		return nil, fmt.Errorf("error decoding response: %s", err)
	}

	fmt.Printf("Orders: %+v\n", orders)

	if orders == nil {
		return nil, fmt.Errorf("no order data found for symbol %s", symbol)
	}

	return orders, nil
}

func (mt *MetaTrader) GetPositions() (*models.Positions, error) {
	response, err := mt.API.SendCommand("POSITIONS")
	if err != nil {
		return nil, err
	}

	responseMap := response.(map[string]interface{})
	dataMap := responseMap["positions"].([]interface{})

	if len(dataMap) == 0 {
		return nil, fmt.Errorf("no position data found")
	}

	var positions *models.Positions

	err = mapstructure.Decode(dataMap, &positions)
	if err != nil {
		return nil, fmt.Errorf("error decoding response: %s", err)
	}

	return positions, nil
}

func (mt *MetaTrader) GetAccountInfo() (*models.AccountInfo, error) {
	response, err := mt.API.SendCommand("ACCOUNT")
	if err != nil {
		return nil, err
	}

	var accountInfo *models.AccountInfo

	err = mapstructure.Decode(response, &accountInfo)
	if err != nil {
		return nil, fmt.Errorf("error decoding response: %s", err)
	}

	fmt.Printf("Account info: %+v\n", accountInfo)

	return accountInfo, nil
}

func (mt *MetaTrader) GetBalance() (*models.Balance, error) {
	response, err := mt.API.SendCommand("BALANCE")
	if err != nil {
		return nil, err
	}

	var balance *models.Balance

	err = mapstructure.Decode(response, &balance)
	if err != nil {
		return nil, fmt.Errorf("error decoding response: %s", err)
	}

	return balance, nil
}

func (mt *MetaTrader) GetHistoricalData(symbol string, timeFrame string, actionType string, from string, to string) (*dataframe.DataFrame, error) {
	fromDate, err := utils.ConvertDateToUTC(from, "02-01-2006 15:04:05")
	if err != nil {
		return nil, err
	}

	toDate, err := utils.ConvertDateToUTC(to, "02-01-2006 15:04:05")
	if err != nil {
		return nil, err
	}

	command := "HISTORY|symbol=" + symbol + "|timeFrame=" + timeFrame + "|actionType=" + actionType + "|from=" + fromDate + "|to=" + toDate

	response, err := mt.API.SendCommand(command)
	if err != nil {
		return nil, err
	}

	// fmt.Println(response)

	return mt.ProcessHistoricalData(response.(map[string]interface{}), timeFrame, actionType)
}

func (mt *MetaTrader) ProcessHistoricalData(data map[string]interface{}, timeFrame string, actionType string) (*dataframe.DataFrame, error) {
	var df dataframe.DataFrame

	// Create an array of structs
	var rates []models.Rate

	var ticks []models.Tick

	var trades []models.Trade

	switch actionType {
	case actiontype.TRADES:
		dataMap := data["trades"].([]interface{})

		if len(dataMap) == 0 {
			return nil, fmt.Errorf("no trade data found")
		}

		for _, trade := range dataMap {
			// Split the input into individual records
			fmt.Sprintf("%v", trade)
			fields := strings.Split(trade.(string), "||")

			trades = append(trades, models.Trade{
				Ticket:    fields[0],
				Timestamp: fields[1],
				Price:     fields[2],
				Volume:    fields[3],
				Symbol:    fields[4],
				Type:      fields[5],
				Entry:     fields[6],
				// Profit:    fields[7],
			})
		}

		df = dataframe.LoadStructs(trades)

	case actiontype.PRICE:
		if timeFrame != timeframes.CURRENT {
			dataMap := data["rates"].([]interface{})

			if len(dataMap) == 0 {
				return nil, fmt.Errorf("no rate data found")
			}

			for _, rate := range dataMap {
				// Split the input into individual records
				fields := strings.Split(rate.(string), "||")

				// date, _ := time.Parse("2006.01.02 15:04:05", fields[0])
				date := fields[0] // time.Parse("2006.01.02 15:04:05", fields[0])
				open := utils.ParseFloat(fields[1])
				high := utils.ParseFloat(fields[2])
				low := utils.ParseFloat(fields[3])
				close := utils.ParseFloat(fields[4])
				tickVolume := utils.ParseFloat(fields[5])
				realVolume := utils.ParseFloat(fields[6])
				spread := utils.ParseInt64(fields[7])

				rates = append(rates, models.Rate{
					Date:       date,
					Open:       open,
					High:       high,
					Low:        low,
					Close:      close,
					TickVolume: tickVolume,
					RealVolume: realVolume,
					Spread:     spread,
				})
			}

			df = dataframe.LoadStructs(rates)
		} else {
			dataMap := data["ticks"].([]interface{})

			if len(dataMap) == 0 {
				return nil, fmt.Errorf("no tick data found")
			}

			for _, tick := range dataMap {
				// Split the input into individual records
				fields := strings.Split(tick.(string), "||")

				timestamp := fields[0]
				bid := utils.ParseFloat(fields[1])
				ask := utils.ParseFloat(fields[2])

				ticks = append(ticks, models.Tick{
					Timestamp: timestamp,
					Bid:       bid,
					Ask:       ask,
				})
			}

			df = dataframe.LoadStructs(ticks)
		}

	default:
		return nil, fmt.Errorf("invalid action type")
	}

	return &df, nil
}

func (mt *MetaTrader) Trade(symbol string, actionType string, volume float64, stoploss float64, takeprofit float64, price float64, deviation float64, id int64) error {

	if len(symbol) == 0 {
		return fmt.Errorf("symbol is required")
	}

	if volume <= 0 {
		return fmt.Errorf("volume must be greater than 0")
	}

	if stoploss <= 0 {
		return fmt.Errorf("stoploss must be greater than 0")
	}

	if takeprofit <= 0 {
		return fmt.Errorf("takeprofit must be greater than 0")
	}

	if deviation < 0 {
		return fmt.Errorf("deviation must be greater than or equal to 0")
	}

	if !utils.Contains(actiontype.ACTIONS, actionType) {
		return fmt.Errorf("invalid action type")

	}

	var command string
	trade_id := strconv.FormatInt(time.Now().Unix(), 10)
	if id > 0 {
		trade_id = strconv.FormatInt(id, 10)
	}
	// expiration = 0 // int(time.time()) + 60 * 60 * 24  # 1 day
	// expiration in 1 hour
	expiration := 0 // int(time.Now().Unix()) + 60*60 // 1 hr
	// expiration = int(time.Now().Unix()) + 60*60 // 1 hr
	command = "TRADE|id=" + trade_id + "|actionType=" + actionType + "|symbol=" + symbol + "|volume=" + strconv.FormatFloat(volume, 'f', -1, 64) + "|price=" + strconv.FormatFloat(price, 'f', -1, 64) + "|stoploss=" + strconv.FormatFloat(stoploss, 'f', -1, 64) + "|takeprofit=" + strconv.FormatFloat(takeprofit, 'f', -1, 64) + "|expiration=" + strconv.Itoa(expiration) + "|deviation=" + strconv.FormatFloat(deviation, 'f', -1, 64)
	// return mt.API.SendCommand("TRADE|id=" + id + "|actionType=" + actionType + "|symbol=" + symbol + "|volume=" + strconv.FormatFloat(volume, 'f', -1, 64) + "|price=" + strconv.FormatFloat(price, 'f', -1, 64) + "|stoploss=" + strconv.FormatFloat(stoploss, 'f', -1, 64) + "|takeprofit=" + strconv.FormatFloat(takeprofit, 'f', -1, 64) + "|deviation=" + strconv.FormatFloat(deviation, 'f', -1, 64))
	fmt.Println(command)

	_, err := mt.API.SendCommand(command)
	if err != nil {
		return err
	}

	return nil
}

func (mt *MetaTrader) Buy(symbol string, volume float64, stoploss float64, takeprofit float64, deviation float64) error {
	return mt.Trade(symbol, "ORDER_TYPE_BUY", volume, stoploss, takeprofit, 0, deviation, 0)
}

func (mt *MetaTrader) Sell(symbol string, volume float64, stoploss float64, takeprofit float64, deviation float64) error {
	return mt.Trade(symbol, "ORDER_TYPE_SELL", volume, stoploss, takeprofit, 0, deviation, 0)
}

func (mt *MetaTrader) BuyLimit(symbol string, volume float64, stoploss float64, takeprofit float64, price float64, deviation float64) error {
	return mt.Trade(symbol, "ORDER_TYPE_BUY_LIMIT", volume, stoploss, takeprofit, price, deviation, 0)
}

func (mt *MetaTrader) SellLimit(symbol string, volume float64, stoploss float64, takeprofit float64, price float64, deviation float64) error {
	return mt.Trade(symbol, "ORDER_TYPE_SELL_LIMIT", volume, stoploss, takeprofit, price, deviation, 0)
}

func (mt *MetaTrader) BuyStop(symbol string, volume float64, stoploss float64, takeprofit float64, price float64, deviation float64) error {
	return mt.Trade(symbol, "ORDER_TYPE_BUY_STOP_LIMIT", volume, stoploss, takeprofit, price, deviation, 0)
}

func (mt *MetaTrader) SellStop(symbol string, volume float64, stoploss float64, takeprofit float64, price float64, deviation float64) error {
	return mt.Trade(symbol, "ORDER_TYPE_SELL_STOP_LIMIT", volume, stoploss, takeprofit, price, deviation, 0)
}

func (mt *MetaTrader) CancelOrderByTicketID(id int64, symbol string, priceOpen float64, stoploss float64, takeprofit float64, volume float64) error {
	// symbol := ""
	// volume := 0.0
	price := 0.0
	if stoploss == 0 || takeprofit == 0 {
		stoploss = priceOpen - 0.5
		takeprofit = priceOpen + 0.5
	}
	// stoploss := 3359.352
	// takeprofit := 3360.352
	// expiration := 0
	deviation := 10.0
	// comment := "cancel order"

	return mt.Trade(symbol, "ORDER_CANCEL", volume, stoploss, takeprofit, price, deviation, id)
}

func (mt *MetaTrader) ClosePositionByTicketID(id int64, symbol string, price float64, stoploss float64, takeprofit float64, volume float64) error {
	// symbol := ""
	// volume := 0.0
	// price :=
	if stoploss == 0 || takeprofit == 0 {
		stoploss = price - 0.5
		takeprofit = price + 0.5
	}
	// stoploss := 0.0
	// takeprofit := 0.0
	// expiration := 0
	deviation := 10.0
	// comment := "close position"

	return mt.Trade(symbol, "POSITION_CLOSE_ID", volume, stoploss, takeprofit, price, deviation, id)
}

func (mt *MetaTrader) ClosePositionBySymbol(symbol string, price float64, stoploss float64, takeprofit float64, volume float64) error {
	// id := ""
	// volume := 0.0
	if stoploss == 0 || takeprofit == 0 {
		stoploss = price - 0.5
		takeprofit = price + 0.5
	}
	// price := 0.0
	// stoploss := 0.0
	// takeprofit := 0.0
	// expiration := 0
	deviation := 10.0
	// comment := "close position"

	return mt.Trade(symbol, "POSITION_CLOSE_SYMBOL", volume, stoploss, takeprofit, price, deviation, 0)
}

func (mt *MetaTrader) ClosePartialPosition(positionID int64, symbol string, price float64, stoploss float64, takeprofit float64, volume float64) error {
	// symbol := ""
	// price := 0.0
	// stoploss := 0.0
	// takeprofit := 0.0
	// expiration := 0
	if stoploss == 0 || takeprofit == 0 {
		stoploss = price - 0.5
		takeprofit = price + 0.5
	}
	deviation := 10.0
	// comment := "close position"

	return mt.Trade(symbol, "POSITION_PARTIAL", volume, stoploss, takeprofit, price, deviation, positionID)
}

func (mt *MetaTrader) CancelAllOrders() error {
	orders, err := mt.GetOrders()
	if err != nil {
		return err
	}

	for _, order := range *orders {
		// parse id string to int

		id, err := strconv.ParseInt(order.Id, 10, 64)
		if err != nil {
			return err
		}

		err = mt.CancelOrderByTicketID(id, order.Symbol, order.Open, order.Stoploss, order.Takeprofit, order.Volume)
		if err != nil {
			return err
		}
	}

	return nil
}
func (mt *MetaTrader) CancelAllOrdersBySymbol(symbol string) error {
	orders, err := mt.GetOrdersBySymbol(symbol)
	if err != nil {
		return err
	}

	for _, order := range *orders {
		// parse id string to int
		if order.Symbol != symbol {
			continue
		}
		id, err := strconv.ParseInt(order.Id, 10, 64)
		if err != nil {
			return err
		}

		err = mt.CancelOrderByTicketID(id, order.Symbol, order.Open, order.Stoploss, order.Takeprofit, order.Volume)
		if err != nil {
			return err
		}
	}

	return nil
}

func (mt *MetaTrader) CloseAllPositions() error {
	positions, err := mt.GetPositions()
	if err != nil {
		return err
	}
	if positions != nil {
		for _, position := range *positions {
			//  parse int to int64
			id := int64(position.Id)
			err := mt.ClosePositionByTicketID(id, position.Symbol, position.Open, position.Stoploss, position.Takeprofit, position.Volume)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
