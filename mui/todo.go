package mui

import (
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func GetTodoMainWindow() *widgets.Paragraph {
	p := widgets.NewParagraph()
	p.Text = "Hello World! this is todo main window"
	p.SetRect(0, 0, 25, 5)
	return p
}


func GetList() *widgets.List{
	listData := []string{
		"[0] gizak/termui",
		"[1] editbox.go",
		"[2] interrupt.go",
		"[3] keyboard.go",
		"[4] output.go",
		"[5] random_out.go",
		"[6] dashboard.go",
		"[7] nsf/termbox-go",
	}

	l := widgets.NewList()
	l.Title = "List"
	l.Rows = listData
	l.SetRect(0, 5, 25, 12)
	l.TextStyle.Fg = ui.ColorYellow
	return l
}