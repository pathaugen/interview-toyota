
package main

import (
  "bufio"
	"fmt"
	"os"
)


func main() {

  // Configure Webserver
  webConfig()

  // Clear Screen
  clearScreen()

  // Display Application Splash
  appSplash()

  // Webserver Status
  webStatus()

  // Start Webserver on PORT #
  webStart( "80" )

  // NOTE: Possible to insert a rolling activity Log if DEBUG is on in this space

  // Reading Input
  scanner := bufio.NewScanner(os.Stdin)
  var inputText string

  // Listen for input and break the loop if inputText == "q"
  for ( inputText != "q" ) {
    // fmt.Print( breakspace + "  [" + cYellow + "q" + cClr + "]" + cBold + cCyan + " + " + cClr + "[" + cYellow + "enter" + cClr + "]" + cBold + cCyan + " to Quit Application: " + cClr + breakspace + cYellow + "  > " + cClr )
    fmt.Print( breakspace + "  [" + cYellow + "q" + cClr + "]" + cBold + cCyan + " + " + cClr + "[" + cYellow + "enter" + cClr + "]" + cBold + cCyan + " to Quit Application (at any time)" + cClr +
      breakline +
      breakspace )

    scanner.Scan()
    inputText = scanner.Text()
  }

  fmt.Print( breakline + cBold + cGreen + "  Application Exited Without Errors" + cClr + breakspace )
}
