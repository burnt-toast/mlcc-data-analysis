package report

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"

	"github.com/burnt-toast/mlcc-data-analysis/program"
)

type Writer struct {
	ProgramData map[string]*program.Instance
}

func (w *Writer) GenerateAttendanceReport() {
	if _, err := os.Stat("attendance-report.csv"); err == nil {
		os.Remove("attendance-report.csv")
	}
	file, err := os.Create("attendance-report.csv")
	checkError("Cannot create file", err)
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"Source", "Category", "GenericName", "EventName", "Capacity", "StartDate", "StartTime", "Attendence", "Percent full"})
	for _, v := range w.ProgramData {
		percent := (float64(v.Attendance) / float64(v.Capacity)) * 100
		percentString := strconv.FormatFloat(percent, 'f', 6, 64)
		err := writer.Write([]string{v.Source, v.Category, v.GenericName, v.EventName, strconv.Itoa(v.Capacity), v.StartDate, v.StartTime, strconv.Itoa(v.Attendance), percentString})
		checkError("Cannot write to file", err)
	}
}

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}
