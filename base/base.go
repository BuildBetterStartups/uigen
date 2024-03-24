// base/base.go
package base

import (
	"fmt"
	"os"
)

// BaseFunction generates a simple HTML mockup and appends it to index.md
func BaseFunction() {
	html := `<div style="border: 2px solid black; width: 320px; height: 568px; position: relative; background-color: #f0f0f5; margin: 0 auto;">
	<div style="background-color: #000; color: #fff; text-align: center; padding: 10px;">
	  Top Nav
	</div>
  
	<div style="background-color: #000; color: #fff; text-align: center; padding: 10px;">
  Bottom Nav
</div>
  </div>
`

	file, err := os.OpenFile("index.md", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	file.WriteString(html)
}