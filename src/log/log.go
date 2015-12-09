//local functions
package log

import (
	"os"
)

func Check(e error, m string) {
	if e != nil {
		panic(m + "\n" + e.Error())
		os.Exit(1)
	}
}
