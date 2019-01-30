
package main

import (
)


// Stock Struct
type Stock struct {
  SymbolsRequested  int         `json:"symbols_requested"`
  SymbolsReturned   int         `json:"symbols_returned"`
  StockData         []StockData `json:"data"`
}

type StockData struct {
  Symbol                string  `json:"symbol"`           // Return Value
  Name                  string  `json:"name"`             // Return Value
  Currency              string  `json:"currency"`         // Return Value
  Price                 string  `json:"price"`            // Return Value
  PriceOpen             string  `json:"price_open"`
  DayHigh               string  `json:"day_high"`
  DayLow                string  `json:"day_low"`
  WeekHigh              string  `json:"52_week_high"`
  WeekLow               string  `json:"52_week_low"`
  DayChange             string  `json:"day_change"`
  ChangePct             string  `json:"change_pct"`
  CloseYesterday        string  `json:"close_yesterday"`  // Return Value
  MarketCap             string  `json:"market_cap"`       // Return Value
  Volume                string  `json:"volume"`           // Return Value
  VolumeAvg             string  `json:"volume_avg"`
  Shares                string  `json:"shares"`
  StockExchangeLong     string  `json:"stock_exchange_long"`
  StockExchangeShort    string  `json:"stock_exchange_short"`
  Timezone              string  `json:"timezone"`         // Return Value
  TimezoneName          string  `json:"timezone_name"`    // Return Value
  GmtOffset             string  `json:"gmt_offset"`       // Return Value
  LastTradeTime         string  `json:"last_trade_time"`  // Return Value
}
