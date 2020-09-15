package main

import (
	"os"
	"time"

	"github.com/ryanyogan/clockface/svg"
)

func main() {
	t := time.Now()
	svg.Write(os.Stdout, t)
}
