package ui

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

type __mymainwindow struct{}

func (*__mymainwindow) init() {}

type MyMainWindow struct {
	*__mymainwindow
	*widgets.QMainWindow
	Centralwidget    *widgets.QWidget
	HorizontalLayout *widgets.QHBoxLayout
	SayHelloButton   *widgets.QPushButton
	Menubar          *widgets.QMenuBar
	Statusbar        *widgets.QStatusBar
}

func NewMyMainWindow(p widgets.QWidget_ITF) *MyMainWindow {
	var par *widgets.QWidget
	if p != nil {
		par = p.QWidget_PTR()
	}
	w := &MyMainWindow{QMainWindow: widgets.NewQMainWindow(par, 0)}
	w.setupUI()
	w.init()
	return w
}
func (w *MyMainWindow) setupUI() {
	if w.ObjectName() == "" {
		w.SetObjectName("MyMainWindow")
	}
	w.Resize2(607, 416)
	w.Centralwidget = widgets.NewQWidget(w, 0)
	w.Centralwidget.SetObjectName("centralwidget")
	w.HorizontalLayout = widgets.NewQHBoxLayout2(w.Centralwidget)
	w.HorizontalLayout.SetObjectName("horizontalLayout")
	w.SayHelloButton = widgets.NewQPushButton(w.Centralwidget)
	w.SayHelloButton.SetObjectName("sayHelloButton")
	w.HorizontalLayout.QLayout.AddWidget(w.SayHelloButton)
	w.SetCentralWidget(w.Centralwidget)
	w.Menubar = widgets.NewQMenuBar(w)
	w.Menubar.SetObjectName("menubar")
	w.Menubar.SetGeometry(core.NewQRect4(0, 0, 607, 24))
	w.SetMenuBar(w.Menubar)
	w.Statusbar = widgets.NewQStatusBar(w)
	w.Statusbar.SetObjectName("statusbar")
	w.SetStatusBar(w.Statusbar)
	w.retranslateUi()
	core.QMetaObject_ConnectSlotsByName(w)

}
func (w *MyMainWindow) retranslateUi() {
	w.SetWindowTitle(core.QCoreApplication_Translate("MyMainWindow", "MainWindow", "", 0))
	w.SayHelloButton.SetText(core.QCoreApplication_Translate("MyMainWindow", "Say Hello", "", 0))

}
