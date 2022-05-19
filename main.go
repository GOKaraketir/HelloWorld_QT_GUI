package main

import (
	"fmt"
	"github.com/GOKaraketir/HelloWorld_QT_GUI/ui"
	"github.com/therecipe/qt/widgets"
	"os"
)

func main() {

	fmt.Println("Hello World")
	widgets.NewQApplication(len(os.Args), os.Args)

	mw := ui.NewMyMainWindow(nil)

	mw.SayHelloButton.ConnectClicked(func(bool) {
		ui.NewSayHelloForm(nil).Show()
	})

	mw.Show()

	widgets.QApplication_Exec()
}
