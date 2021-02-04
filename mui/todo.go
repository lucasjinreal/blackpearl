package mui

import "github.com/gizak/termui/v3/widgets"

func GetTodoMainWindow() *widgets.Paragraph {
	p := widgets.NewParagraph()
	p.Text = "Hello World! this is todo main window"
	p.SetRect(0, 0, 25, 5)
	return p
}


func GetList() {

}