package main

import (
	"os"

	"github.com/talaniz/golang-ttd/mocks"
)

func main() {
	mocks.Countdown(os.Stdout)
}
