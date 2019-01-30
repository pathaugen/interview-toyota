
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

  consoleOutput += cBold + cCyan + "Webserver Request" + cClr + " #" + cYellow + strconv.Itoa(webserverRequests) + cClr + breakspace + cBold + cCyan + "Request URL:" + cClr + " http://localhost:" + cClr + cYellow + "PORT" + cClr + "/stock/" + cClr + cYellow + "STOCK" + cClr + "/?stock_exchange=" + cClr + cYellow + "xxx" + cClr + breakspace

  // Set the request URL
  restUrl := "https://www.worldtradingdata.com/api/v1/stock?symbol=AAPL&stock_exchange=NYSE&api_token=" + apiToken
  response, err := http.Get( restUrl )
  if err != nil {
    consoleOutput += "The HTTP request failed with error:" + breakspace + cRed + err.Error() + cClr
  } else {
    data, _ := ioutil.ReadAll(response.Body)
    consoleOutput += string(data)
  }
  consoleOutput += breakspace + breakspace

  jsonData := map[string]string{"firstname": "Nic", "lastname": "Raboy"}
  jsonValue, _ := json.Marshal(jsonData)
  response, err = http.Post("https://httpbin.org/post", "application/json", bytes.NewBuffer(jsonValue))
  if err != nil {
    consoleOutput += "The HTTP request failed with error:" + breakspace + cRed + err.Error() + cClr
  } else {
    data, _ := ioutil.ReadAll(response.Body)
    consoleOutput += string(data)
  }
  consoleOutput += breakspace + breakspace

  fmt.Print( consoleOutput )
	fmt.Fprintf( w, webOutput )
}
