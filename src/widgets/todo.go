package widgets

import (
	ui "blackpearl/src/termui"
	"blackpearl/src/utils"
	"strconv"
	"time"
)


type TodoWidget struct {
	*ui.Table
	updateInterval time.Duration
}

func NewTodoWidget() *TodoWidget {
	self := &TodoWidget{
		Table: ui.NewTable(),
		updateInterval: time.Second,
	}


	self.Title = " Daybreak Todos"
	self.Header = []string{"ID", "Title", "Project", "DueTime", "CreateTime", "Done"}
	self.ColGap = 2
	self.ColResizer = func() {
		self.ColWidths = []int{
			utils.MaxInt(4, (self.Inner.Dx()-29)/2),
			utils.MaxInt(5, (self.Inner.Dx()-29)/2),
			4, 5, 5, 5,
		}
	}

	self.update()

	go func() {
		for range time.NewTicker(self.updateInterval).C {
			self.Lock()
			self.update()
			self.Unlock()
		}
	}()
	return self
}

func (self *TodoWidget) update() {
	self.Rows = make([][]string, 12)

	// update data?
	for i := 0; i < 12; i++ {
		self.Rows[i] = make([]string, 6)
		self.Rows[i][0] = strconv.Itoa(i)
		self.Rows[i][1] = "Do some trick, z做吧"
		self.Rows[i][2] = "工作"
		self.Rows[i][3] = "2021/03/16"
		self.Rows[i][4] = "2021/03/16 08:23"
		self.Rows[i][5] = "√"
	}

}