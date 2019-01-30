
package main

import (
  "bytes"
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
  http.HandleFunc("/stock", handlerStock)
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

  webserverRequests += 1

  consoleOutput += cBold + cCyan + "Webserver Request" + cClr + " #" + cYellow + strconv.Itoa(webserverRequests) + cClr
  consoleOutput += breakspace + cBold + cCyan + "Request URL:" + cClr + " http://localhost:" + cClr + cYellow + "PORT" + cClr + "/stock/" + cClr + cYellow + "STOCK" + cClr + "/?stock_exchange=" + cClr + cYellow + "xxx" + cClr
  consoleOutput += breakspace + cBold + cCyan + "Response:" + cClr + breakspace

  // Set the request URL
  restUrl := "https://www.worldtradingdata.com/api/v1/stock" + "?api_token=" + apiToken + "&symbol=" + "AAPL" + "&stock_exchange=" + "NYSE"
  response, err := http.Get( restUrl )
  if err != nil {
    consoleOutput += "The HTTP request failed with error:" + breakspace + cRed + err.Error() + cClr + breakspace
  } else {
    // Example Response:
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
    // Desired Responses:
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
    data, _ := ioutil.ReadAll(response.Body)
    consoleOutput += string(data)
  }
  consoleOutput += breakspace + breakspace

  jsonData := map[string]string{"firstname": "Nic", "lastname": "Raboy"}
  jsonValue, _ := json.Marshal(jsonData)
  response, err = http.Post("https://httpbin.org/post", "application/json", bytes.NewBuffer(jsonValue))
  if err != nil {
    consoleOutput += "The HTTP request failed with error:" + breakspace + cRed + err.Error() + cClr + breakspace
  } else {
    data, _ := ioutil.ReadAll(response.Body)
    consoleOutput += string(data)
  }
  consoleOutput += breakspace + breakspace

  if debug { fmt.Print( consoleOutput ) }
	fmt.Fprintf( w, webOutput )
}
