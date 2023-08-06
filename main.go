package main

import (
	"fmt"
	"log"
	"time"

	weatherrange "github.com/bcchapman/weather-range-check"
	"github.com/bcchapman/yeelight-client"
)

const (
	RED    = "FF0000"
	YELLOW = "FFFF00"
	GREEN  = "00FF00"
)

type logWriter struct {
}

func (writer logWriter) Write(bytes []byte) (int, error) {
	return fmt.Print(time.Now().UTC().Format("2006-01-02T15:04:05.999Z") + " [DEBUG] " + string(bytes))
}

func main() {
	log.SetFlags(0)
	log.SetOutput(new(logWriter))

	y := initializeLight()

	weatherrange.StartListening(onChangeHandler(y))
}

func onChangeHandler(y *yeelight.Yeelight) func(bool) {
	return func(lastReadingInRange bool) {
		if lastReadingInRange {
			y.SetColor(GREEN)
		} else {
			y.SetColor(RED)
		}
	}
}

func initializeLight() *yeelight.Yeelight {
	y, err := yeelight.Discover()
	checkError(err)

	on, err := y.GetProp("power")
	checkError(err)
	log.Printf("Power is %s\n", on[0].(string))

	err = y.SetPower(true)
	checkError(err)

	err = y.SetColor(RED) // default to RED
	checkError(err)

	return y
}

func checkError(err error) {
	if nil != err {
		log.Fatal(err)
	}
}
