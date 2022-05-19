package ui

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

type __sayhelloform struct{}

func (*__sayhelloform) init() {}

type SayHelloForm struct {
	*__sayhelloform
	*widgets.QWidget
	VerticalLayout *widgets.QVBoxLayout
	Label          *widgets.QLabel
}

func NewSayHelloForm(p widgets.QWidget_ITF) *SayHelloForm {
	var par *widgets.QWidget
	if p != nil {
		par = p.QWidget_PTR()
	}
	w := &SayHelloForm{QWidget: widgets.NewQWidget(par, 0)}
	w.setupUI()
	w.init()
	return w
}
func (w *SayHelloForm) setupUI() {
	if w.ObjectName() == "" {
		w.SetObjectName("SayHelloForm")
	}
	w.Resize2(400, 300)
	w.VerticalLayout = widgets.NewQVBoxLayout2(w)
	w.VerticalLayout.SetObjectName("verticalLayout")
	w.Label = widgets.NewQLabel(w, 0)
	w.Label.SetObjectName("label")
	w.Label.SetAlignment(core.Qt__AlignCenter)
	w.VerticalLayout.QLayout.AddWidget(w.Label)
	w.retranslateUi()
	core.QMetaObject_ConnectSlotsByName(w)

}
func (w *SayHelloForm) retranslateUi() {
	w.SetWindowTitle(core.QCoreApplication_Translate("SayHelloForm", "Form", "", 0))
	w.Label.SetText(core.QCoreApplication_Translate("SayHelloForm", "Hello World", "", 0))

}
