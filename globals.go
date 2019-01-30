
package main

import (
)


// Debug Switch
// Always have a quick switch to display debug data.
// Setting this will show requests in the console as they are served.
var debug = true


// Golang Console Colors
// Example: fmt.Print( cRed + "HelloWorld" + cClr )
var cClr				= "\u001b[0m"

var cBold				= "\u001b[1m"

var cBlack			= "\u001b[30m"
var cRed				= "\u001b[31m"
var cGreen			= "\u001b[32m"
var cYellow			= "\u001b[33m"
var cBlue				= "\u001b[34m"
var cMagenta		= "\u001b[35m"
var cCyan				= "\u001b[36m"
var cWhite			= "\u001b[37m"

var cBlackBG		= "\u001b[40m"
var cRedBG			= "\u001b[41m"
var cGreenBG		= "\u001b[42m"
var cYellowBG		= "\u001b[43m"
var cBlueBG			= "\u001b[44m"
var cMagentaBG	= "\u001b[45m"
var cCyanBG			= "\u001b[46m"
var cWhiteBG		= "\u001b[47m"


// Output Simplification
var breakspace = "\n"
var breakline = breakspace + cBlue + "  ====================================================" + cClr + breakspace


// Console Splash
var appinfo = `
  ` + cBlue + `====================================================` + cBold + cCyan + `
   _____      _   _    _
  |  __ \    | | | |  | |
  | |__) |_ _| |_| |__| | __ _ _   _  __ _  ___ _ __
  |  ___/ _`+"`"+` | __|  __  |/ _`+"`"+` | | | |/ _`+"`"+` |/ _ \ '_ \
  | |  | (_| | |_| |  | | (_| | |_| | (_| |  __/ | | |
  |_|   \__,_|\__|_|  |_|\__,_|\__,_|\__, |\___|_| |_|
                                     __/ |
                                    |___/` + cClr + `
  ` + cCyan + `Interview: ` + cWhite + `Toyota` + cClr + `
  ` + cCyan + `https://github.com/` + cYellow + `pathaugen` + cCyan + `/interview-toyota` + cClr + `
  ` + cBlue + `====================================================` + cClr + `
`


// Webserver Globals
var webserverRunning = false
var webserverPort = "UNKNOWN"
var webserverRequests = 0


// Website Template
var htmlTemplate = `
  <html>
    <head>
      <title>PatHaugen Interview: Toyota</title>
    </head>
    <body style="padding:2%;">
      <h1>PatHaugen Interview: Toyota</h1>
      <div id="menu">
        <a href="/">Home</a>
        <a href="/stock">Stock</a>
      </div>
    </body>
  </html>
`


// Website Style
var htmlStyle = `
  body {
    background-color: #3C3C3C;
    color: white;
    font-family: sans-serif;
    padding: 2%%;
  }
  body #menu a {
    background-color: darkblue;
    border-radius: 2px;
    color: white;
    display: inline-block;
    padding: 5px 10px;
    text-decoration: none;
  }
`
