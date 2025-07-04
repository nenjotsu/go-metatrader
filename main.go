package main

import (
	"fmt"
	"strings"

	"github.com/nenjotsu/go-metatrader/metatrader"
)

func main() {
	// symbols := []string{"EURUSD", "GBPUSD", "USDJPY", "XAUUSD", "AUDNZD"}
	symbols := []string{"EURUSD", "XAUUSD"}
	mt := metatrader.NewMetaTraderClient("0.0.0.0", 1122, false, true, "123456", symbols)
	// mt := metatrader.NewMetaTraderClient("0.0.0.0", 15557, false, true, "123456", symbols)

	// Connect to the server
	mt.Connect()
	defer mt.Disconnect()

	// Send a request to the server

	// get account info
	accountInfo, err := mt.GetAccountInfo()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Account info: %+v\n", accountInfo)

	balance, err := mt.GetBalance()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Balance: %+v\n", balance)

	// // Print the response

	// symbol := "XAUUSD"
	// currentPrice := 0.0

	// always get the current price
	// for {
	// 	// get current price
	// 	currentPrice, err := mt.GetCurrentPrice(symbol)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		return
	// 	}

	// 	// Print the response
	// 	fmt.Println(currentPrice)

	// 	time.Sleep(1 * time.Millisecond)
	// }

	// get historical data
	// history, err := mt.GetHistoricalData(symbol, timeframes.M1, actiontype.PRICE, "23-06-2025 00:00:00", "27-06-2025 00:00:00")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(history)

	// get orders. limit or stop orders
	orders, err := mt.GetOrders()
	if err != nil {
		fmt.Println(err)
		// return
	}
	fmt.Printf("Orders: %+v\n", orders)

	// // get positions
	// positions, err := mt.GetPositions()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(positions)

	// fmt.Printf(mt.)

	// create a var with dictionary of current prices per symbol
	// var currentPrices map[string]*models.CurrentPrice
	// []*metatrader.CurrentPrice

	// []*model.CurrentPrice

	// for _, symbol := range symbols {
	// 	// get current price
	// 	currentPrice, err := mt.GetCurrentPrice(symbol)

	// 	if err != nil {
	// 		fmt.Println(err)
	// 		return
	// 	}

	// 	// currentPrices[symbol] = currentPrice

	// 	// Print the response
	// 	fmt.Println(currentPrice)

	// }

	// // Print the response

	pips := 100
	stopLevel := 0.5
	normalizedPips := float64(pips) * stopLevel
	fmt.Printf("%.2f normalized pips \n", normalizedPips)
	// symbol := "XAUUSD"
	// fmt.Println(symbol)
	// fmt.Println(currentPrices[symbol])

	// fmt.Println(stopLoss, takeProfit)
	// err = mt.Buy(symbol, 0.01, stopLoss, takeProfit, 5)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// fmt.Println("BUY Order Sent")
	// separate go routine
	// check if market if closed
	// if !mt.() {
	// 	fmt.Println("Market is closed")
	// 	return
	// }

	// for {
	// 	// get current price
	// 	// get positions

	// 	// fmt.Println(currentPrice)
	positions, err := mt.GetPositions()
	if err != nil {
		fmt.Println(err)
		// return
	}
	fmt.Printf("Positions: %+v\n", positions)
	// 	time.Sleep(1 * time.Millisecond)

	currentPrice, err := mt.GetCurrentPrice("XAUUSD")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Current Price: %+v\n", currentPrice)

	// test buy
	// stopLoss := float64(currentPrice.Tick.Bid) - normalizedPips
	// takeProfit := float64(currentPrice.Tick.Bid) + normalizedPips
	// fmt.Println(stopLoss, takeProfit)
	// err = mt.Buy("XAUUSD", 0.01, stopLoss, takeProfit, 2)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// test Sell
	// stopLoss := float64(currentPrice.Tick.Ask) + normalizedPips
	// takeProfit := float64(currentPrice.Tick.Ask) - normalizedPips

	// fmt.Println(stopLoss, takeProfit)
	// err = mt.Sell("XAUUSD", 0.01, stopLoss, takeProfit, 2)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// test buy limit
	// priceLimit := float64(currentPrice.Tick.Bid) - (normalizedPips * 10)
	// stopLoss := float64(currentPrice.Tick.Bid) - (normalizedPips * 15)
	// takeProfit := float64(currentPrice.Tick.Bid) + (normalizedPips * 10)
	// fmt.Println(stopLoss, takeProfit)
	// err = mt.BuyLimit("XAUUSD", 0.01, stopLoss, takeProfit, priceLimit, 5)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// test sell limit
	// priceLimit := float64(currentPrice.Tick.Ask) + (normalizedPips * 2)
	// stopLoss := float64(currentPrice.Tick.Ask) + normalizedPips + (normalizedPips * 2)
	// takeProfit := float64(currentPrice.Tick.Ask) - (normalizedPips * 2)
	// fmt.Println(stopLoss, takeProfit)
	// err = mt.SellLimit("XAUUSD", 0.01, stopLoss, takeProfit, priceLimit, 2)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// test buy stop
	// priceLimit := float64(currentPrice.Tick.Bid) + (normalizedPips * 1.2)
	// stopLoss := float64(currentPrice.Tick.Bid) + (normalizedPips * 0.8)
	// takeProfit := float64(currentPrice.Tick.Bid) + (normalizedPips * 1.5)
	// fmt.Println(stopLoss, takeProfit)
	// err = mt.BuyStop("XAUUSD", 0.01, stopLoss, takeProfit, priceLimit, 20)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// test sell stop
	// priceLimit := float64(currentPrice.Tick.Ask) - (normalizedPips * 1.1)
	// stopLoss := float64(currentPrice.Tick.Ask) - (normalizedPips * 0.9)
	// takeProfit := float64(currentPrice.Tick.Ask) - (normalizedPips * 1.2)
	// fmt.Println(stopLoss, takeProfit)
	// err = mt.SellStop("XAUUSD", 0.01, stopLoss, takeProfit, priceLimit, 10)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// test cancel position
	// Orders: &[{Id:1105645678 Magic:0 Symbol:XAUUSD Type:ORDER_TYPE_BUY_STOP TimeSetup:0 Open:3361.352 Stoploss:3359.352 Takeprofit:3362.352 Volume:0.01} {Id:1105778156 Magic:0 Symbol:XAUUSD Type:ORDER_TYPE_BUY_LIMIT TimeSetup:0 Open:3343.637 Stoploss:0 Takeprofit:0 Volume:0.02} {Id:1105778368 Magic:0 Symbol:XAUUSD Type:ORDER_TYPE_SELL_LIMIT TimeSetup:0 Open:3355.284 Stoploss:0 Takeprofit:0 Volume:0.02} {Id:1105778463 Magic:0 Symbol:XAUUSD Type:ORDER_TYPE_BUY_LIMIT TimeSetup:0 Open:3343.791 Stoploss:0 Takeprofit:0 Volume:0.02}]
	if orders != nil {
		for _, order := range *orders {
			// if order.Type == "ORDER_TYPE_BUY_STOP" {
			fmt.Printf("order: %+v\n", order)
			// if order.Type == "ORDER_TYPE_BUY_STOP" || order.Type == "ORDER_TYPE_BUY_STOP_LIMIT" || order.Type == "ORDER_TYPE_BUY_LIMIT" {
			// 	id, err := strconv.Atoi(order.Id)
			// 	if err != nil {
			// 		fmt.Println(err)
			// 		return
			// 	}

			// 	err = mt.CancelOrderByTicketID(int64(id), order.Symbol, order.Open, order.Stoploss, order.Takeprofit, order.Volume)
			// 	if err != nil {
			// 		fmt.Println(err)
			// 		return
			// 	}
			// }
			// if order.Type == "ORDER_TYPE_SELL_STOP" || order.Type == "ORDER_TYPE_SELL_STOP_LIMIT" || order.Type == "ORDER_TYPE_SELL_LIMIT" {
			// 	id, err := strconv.Atoi(order.Id)
			// 	if err != nil {
			// 		fmt.Println(err)
			// 		return
			// 	}

			// 	err = mt.CancelOrderByTicketID(int64(id), order.Symbol, order.Open, order.Stoploss, order.Takeprofit, order.Volume)
			// 	if err != nil {
			// 		fmt.Println(err)
			// 		return
			// 	}
			// }
		}
	}
	// map of current prices by symbol
	currentPrices := make(map[string]float64)

	// loop on positions
	if positions != nil {
		// close all positions
		err := mt.CloseAllPositions()
		if err != nil {
			fmt.Println(err)
			return
		}
		for _, position := range *positions {
			// get current Price
			fmt.Printf("position: %+v\n", position)
			currentPrice, ok := currentPrices[position.Symbol]
			if !ok {
				symbolPrice, err := mt.GetCurrentPrice(position.Symbol)
				if err != nil {
					fmt.Println(err)
					return
				}
				currentPrices[position.Symbol] = symbolPrice.Tick.Bid
				if strings.Contains(position.Type, "BUY") {
					currentPrices[position.Symbol] = symbolPrice.Tick.Ask
				}
			}
			if currentPrice == 0 {
				currentPrice = position.Open
			}

			fmt.Printf("currentPrice: %+v\n", currentPrice)
			symbol := position.Symbol

			if symbol == "XAUUSD" {
				// err = mt.ClosePositionByTicketID(int64(position.Id), position.Symbol, currentPrice, position.Stoploss, position.Takeprofit, position.Volume)
				// if err != nil {
				// 	fmt.Println(err)
				// 	return
				// }

				err = mt.ClosePositionBySymbol(symbol, currentPrice, position.Stoploss, position.Takeprofit, position.Volume)
				if err != nil {
					fmt.Println(err)
					return
				}

				// err = mt.ClosePartialPosition(int64(position.Id), position.Symbol, currentPrice, position.Stoploss, position.Takeprofit, (position.Volume / 2))
				// if err != nil {
				// 	fmt.Println(err)
				// 	return
				// }
			}
		}
	}

	if orders != nil {
		err := mt.CancelAllOrders()
		if err != nil {
			fmt.Println(err)
			return
		}
		// err := mt.CancelAllOrdersBySymbol("XAUUSD")
		// if err != nil {
		// 	fmt.Println(err)
		// 	return
		// }
	}

	// for _, symbol := range symbols {
	// 	currentPrice, err := mt.GetCurrentPrice(symbol)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		return
	// 	}

	// 	history, err := mt.GetHistoricalData(symbol, timeframes.M1, actiontype.TRADES, "23-06-2025 00:00:00", "27-06-2025 00:00:00")
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		return
	// 	}
	// 	fmt.Println(history)
	// 	fmt.Println(currentPrice)
	// 	// log readable current time
	// 	fmt.Printf("%s\n", time.Now().Format("2006-01-02 15:04:05"))

	// 	// 	// Print the response
	// 	// 	// fmt.Println("Line 149")

	// 	// 	// if symbol == "XAUUSD" {
	// 	// 	// 	stopLoss := float64(currentPrice.Tick.Bid) - normalizedPips
	// 	// 	// 	takeProfit := float64(currentPrice.Tick.Bid) + normalizedPips

	// 	// 	// 	fmt.Println(stopLoss, takeProfit)
	// 	// 	// 	err = mt.Buy(symbol, 0.01, stopLoss, takeProfit, 5)
	// 	// 	// 	if err != nil {
	// 	// 	// 		fmt.Println(err)
	// 	// 	// 		return
	// 	// 	// 	}
	// 	// 	// }
	// 	// time.Sleep(10 * time.Millisecond)
	// 	// time.Sleep(5 * time.Second)
	// }
	// time.Sleep(10 * time.Second)
	// }

}
