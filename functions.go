
package main

import (
  //"bytes"
  "encoding/json"
	"fmt"
  "io/ioutil"
  "net/http"
  "os"
  "os/exec"
  "strconv"
  "strings"
)


// Clear Screen
func clearScreen() {
  cmd := exec.Command("cmd", "/c", "cls")
  cmd.Stdout = os.Stdout
  cmd.Run()
}


// App Splash
func appSplash() {
  fmt.Print( appinfo )
}


// Debug Status
func debugStatus() {
  fmt.Print( breakspace + cBold + cCyan + "  Debug Status:" + cClr )
  fmt.Print( "      [" )
  if debug {
    fmt.Print( cBold + cGreen + "DEBUG ON" + cClr )
  } else {
    fmt.Print( cBold + cRed + "DEBUG OFF" + cClr )
  }
  fmt.Print( "]" + breakspace )
}


// Webserver Status
func webStatus() {
  fmt.Print( breakspace + cBold + cCyan + "  Webserver Status:" + cClr )
  fmt.Print( "  [" )
  if webserverRunning {
    fmt.Print( cBold + cGreen + "RUNNING" + cClr )
  } else {
    fmt.Print( cBold + cRed + "STOPPED" + cClr )
  }
  fmt.Print( "]" + breakspace )
}

// Webserver Setup
func webConfig() {
  http.HandleFunc("/", handlerRoot)
  http.HandleFunc("/stock/", handlerStock)
}


// Webserver Start
func webStart( port string ) {
  webserverPort = port
  fmt.Print( breakspace + cBold + cCyan + "  Starting Internal Webserver (port #" + cClr + cYellow + webserverPort + cBold + cCyan + "):" + cClr )
  fmt.Print( breakspace + "  http://localhost:" + cYellow + webserverPort + cClr + "/" + breakspace )
  go func() {
    http.ListenAndServe( ":" + port, nil )
  }()
  webserverRunning = true
  webStatus()
}


// Webserver Root Handler
// http://localhost:{port}/
func handlerRoot(w http.ResponseWriter, r *http.Request) {
  output := strings.Replace( htmlTemplate, "</head>", "<style>" + htmlStyle + "</style></head>", 2)
	fmt.Fprintf( w, output )
}


// Webserver Stock Handler
// http://localhost:{port}/stock/{symbol}/(?stock_exchange={NASDAQ,NYSE})
// Example Response:
//{
//  "NASDAQ": {
//    "symbol": "AAPL",
//    "name": "Apple Inc.",
//    "price": "154.94",
//    "close_yesterday": "154.94",
//    "currency": "USD",
//    "market_cap": "732835688367",
//    "volume": "142022",
//    "timezone": "EST",
//    "timezone_name": "America/New_York",
//    "gmt_offset": "-18000",
//    "last_trade_time": "2019-01-16 16:00:01"
//  }
//}
func handlerStock(w http.ResponseWriter, r *http.Request) {
  consoleOutput := ``
  webOutput := ``

  // Log each request made to /stock
  webserverRequests += 1
  consoleOutput += cBold + cCyan + "Webserver Request" + cClr + " #" + cYellow + strconv.Itoa(webserverRequests) + cClr
  consoleOutput += breakspace + cBold + cCyan + "Request URL:" + cClr + " http://localhost:" + cClr + cYellow + webserverPort + cClr + "/stock/" + cClr

  pageRequestedString := strings.TrimSuffix( strings.TrimPrefix(r.URL.Path, "/"), "/")
  pageRequestedStringArray := strings.Split( pageRequestedString, "/" )

  if len( pageRequestedStringArray ) == 1 {
    consoleOutput += breakspace + cRed + "Warning: No Stock Specified!" + cClr + breakspace + breakspace // Show Error
    webOutput += `{"warning":"You must specify a stock!"}`
  } else if len( pageRequestedStringArray ) > 1 {

    // Selected Stock
    selectedStock := pageRequestedStringArray[1]
    consoleOutput += cYellow + selectedStock + cClr + "/"

    // Selected Exchange
    selectedExchange := "AMEX" // e.g. AMEX,NASDAQ,NYSE
    selectedExchangeArr := make(map[string]bool)
    if r.URL.Query().Get("stock_exchange") != "" {
      selectedExchange = r.URL.Query().Get("stock_exchange")

      // For each CSV, append to array
    	selectedExchangeArray := strings.Split( selectedExchange, "," )
      for i := 0; i < len(selectedExchangeArray); i++ {
        selectedExchangeArr[selectedExchangeArray[i]] = true
      }

      consoleOutput += "?stock_exchange=" + cClr + cYellow + selectedExchange + cClr
    } else {
      selectedExchangeArr["AMEX"] = true
    }

    // User Selection Output
    consoleOutput += breakspace + cBold + cCyan + "Selected Stock: " + cClr + cYellow + selectedStock + cClr
    if r.URL.Query().Get("stock_exchange") != "" {
      consoleOutput += breakspace + cBold + cCyan + "Selected Exchange(s): " + cClr + cYellow + selectedExchange + cClr
    }

    // Set the request URL
    restUrl := "https://www.worldtradingdata.com/api/v1/stock" + "?api_token=" + apiToken + "&symbol=" + selectedStock
    if r.URL.Query().Get("stock_exchange") != "" {
      restUrl += "&stock_exchange=" + selectedExchange
    }

    // REST API via GET
    response, err := http.Get( restUrl )

    // REST API via POST
    //jsonData := map[string]string{ "field1": "value1", "field2": "value2" }
    //jsonValue, _ := json.Marshal(jsonData)
    //response, err = http.Post( restUrl, "application/json", bytes.NewBuffer(jsonValue) )

    if err != nil {
      consoleOutput += "The HTTP request failed with error:" + breakspace + cRed + err.Error() + cClr + breakspace
    } else {

      // Example Response from REST API:
      //{
      //  "symbols_requested": 1,
      //  "symbols_returned": 1,
      //  "data": [
      //    {
      //      "symbol": "AAPL",
      //      "name": "Apple Inc.",
      //      "currency": "USD",
      //      "price": "154.68",
      //      "price_open": "156.25",
      //      "day_high": "158.13",
      //      "day_low": "154.11",
      //      "52_week_high": "233.47",
      //      "52_week_low": "142.00",
      //      "day_change": "-1.62",
      //      "change_pct": "-1.04",
      //      "close_yesterday": "156.30",
      //      "market_cap": "731605893397",
      //      "volume": "25718",
      //      "volume_avg": "42370729",
      //      "shares": "4729803000",
      //      "stock_exchange_long": "NASDAQ Stock Exchange",
      //      "stock_exchange_short": "NASDAQ",
      //      "timezone": "EST",
      //      "timezone_name": "America/New_York",
      //      "gmt_offset": "-18000",
      //      "last_trade_time": "2019-01-29 16:00:01"
      //    }
      //  ]
      //}

      // Desired Output to End User:
      // symbol
      // name
      // price
      // close_yesterday
      // currency
      // market_cap
      // volume
      // timezone
      // timezone_name
      // gmt_offset
      // last_trade_time

      byteValue, _ := ioutil.ReadAll(response.Body)
      // fmt.Println( string(byteValue) ) // DEBUG

      // Initialize our Stock array
      var stock Stock
      json.Unmarshal( []byte(byteValue), &stock )
      //fmt.Println( stock ) // DEBUG

      if len(stock.StockData) > 0 {
        consoleOutput += breakspace + cBold + cCyan + "Response:" + cClr + breakspace
        for i := 0; i < len(stock.StockData); i++ {

          //consoleOutput += "symbols_requested: " + strconv.Itoa(stock.SymbolsRequested) + breakspace // DEBUG
          //consoleOutput += "symbols_returned: " + strconv.Itoa(stock.SymbolsRequested) + breakspace // DEBUG

          consoleOutput += "stock_exchange_short: " + cYellow + stock.StockData[i].StockExchangeShort + cClr + breakspace

          consoleOutput += "symbol: " + cYellow + stock.StockData[i].Symbol + cClr + breakspace
          consoleOutput += "name: " + cYellow + stock.StockData[i].Name + cClr + breakspace
          consoleOutput += "price: " + cYellow + stock.StockData[i].Price + cClr + breakspace
          consoleOutput += "close_yesterday: " + cYellow + stock.StockData[i].CloseYesterday + cClr + breakspace
          consoleOutput += "currency: " + cYellow + stock.StockData[i].Currency + cClr + breakspace
          consoleOutput += "market_cap: " + cYellow + stock.StockData[i].MarketCap + cClr + breakspace
          consoleOutput += "volume: " + cYellow + stock.StockData[i].Volume + cClr + breakspace
          consoleOutput += "timezone: " + cYellow + stock.StockData[i].Timezone + cClr + breakspace
          consoleOutput += "timezone_name: " + cYellow + stock.StockData[i].TimezoneName + cClr + breakspace
          consoleOutput += "gmt_offset: " + cYellow + stock.StockData[i].GmtOffset + cClr + breakspace
          consoleOutput += "last_trade_time: " + cYellow + stock.StockData[i].LastTradeTime + cClr + breakspace

          // Skip any data where stock_exchange_short isn't part of user's requested selectedExchange
          if selectedExchangeArr[stock.StockData[i].StockExchangeShort] {

            webOutput += `{"` + stock.StockData[i].StockExchangeShort + `":{
              "symbol":"` + stock.StockData[i].Symbol + `",
              "name":"` + stock.StockData[i].Name + `",
              "price":"` + stock.StockData[i].Price + `",
              "close_yesterday":"` + stock.StockData[i].CloseYesterday + `",
              "currency":"` + stock.StockData[i].Currency + `",
              "market_cap":"` + stock.StockData[i].MarketCap + `",
              "volume":"` + stock.StockData[i].Volume + `",
              "timezone":"` + stock.StockData[i].Timezone + `",
              "timezone_name":"` + stock.StockData[i].TimezoneName + `",
              "gmt_offset":"` + stock.StockData[i].GmtOffset + `",
              "last_trade_time":"` + stock.StockData[i].LastTradeTime + `"
              }}`

          } else {
            consoleOutput += cRed + "Warning: stock_exchange_short doesn't match user's exchange selection parameters!" + cClr + breakspace
            webOutput += `{ }`
          }

        }
      } else {
        consoleOutput += breakspace + cRed + string(byteValue) + cClr // Show Error
        webOutput += string(byteValue)
      }

    }
    consoleOutput += breakspace + breakspace
  }

  if debug { fmt.Print( consoleOutput ) }

  w.Header().Set("Content-Type", "application/json; charset=utf-8")
  //fmt.Fprint(w, "{" + webOutput + "}")
  fmt.Fprint(w, webOutput )
}
