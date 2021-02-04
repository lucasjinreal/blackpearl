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
