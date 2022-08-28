package utils

import (
	"fmt"

	"github.com/fatih/color"
)

type Colors struct {
	RED    string
	GREEN  string
	YELLOW string
	CYAN   string
}

var PrintColors = Colors{RED: "red", GREEN: "green", YELLOW: "yellow", CYAN: "cyan"}

func ColoredPrint(msg, clr string) {
	ColoredPrintf(msg, clr)
	fmt.Println()
}

func ColoredPrintf(msg, clr string) {
	var pr *color.Color
	switch clr {
	case "red":
		pr = color.New(color.FgRed)
	case "green":
		pr = color.New(color.FgGreen)
	case "yellow":
		pr = color.New(color.FgYellow)
	case "cyan":
		pr = color.New(color.FgCyan)
	default:
		pr = color.New(color.FgWhite)
	}

	pr.Printf(msg)
}
