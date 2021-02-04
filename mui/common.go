package mui

import (
	ui "github.com/gizak/termui/v3"
	"log"
)


func InitUI () {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()
}

func HoldUI () {
	// quit until user trigger Q/q
	uiEvents := ui.PollEvents()
	for {
		select {
		case e := <-uiEvents:
			switch e.ID {
			case "q", "<C-c>":
				return
			}
		}
	}
}
