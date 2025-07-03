package models

// 2025.07.03 22:54:54.759	mt-server (XAUUSD,M1)	ServerSocketSend -> client.responseData: {"error":false,"positions":[{"id":1107579898,"magic":0,"symbol":"XAUUSD","type":"POSITION_TYPE_BUY","time_setup":1751554044,"open":3328.11400,"stoploss":0.00000,"takeprofit":0.00000,"volume":0.10000},{"id":1107579898,"magic":0,"symbol":"XAUUSD","type":"POSITION_TYPE_BUY","time_setup":1751554044,"open":3328.11400,"stoploss":0.00000,"takeprofit":0.00000,"volume":0.10000},{"id":1107579898,"magic":0,"symbol":"XAUUSD","type":"POSITION_TYPE_BUY","time_setup":1751554044,"o

type Position struct {
	Id           int     `json:"id"`
	Magic        int     `json:"magic"`
	Symbol       string  `json:"symbol"`
	Type         string  `json:"type"`
	TimeSetup    int     `json:"time_setup"`
	Open         float64 `json:"open"`
	CurrentPrice float64 `json:"current"`
	Stoploss     float64 `json:"stoploss"`
	Takeprofit   float64 `json:"takeprofit"`
	Volume       float64 `json:"volume"`
}

type Positions []Position
