package report

import (
	"fmt"

	"github.com/burnt-toast/mlcc-data-analysis/program"
)

type Writer struct {
	ProgramData map[string]*program.Instance
}

func (w *Writer) GenerateAttendanceReport() {
	fmt.Println("test")
}
