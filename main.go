package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/burnt-toast/mlcc-data-analysis/program"
	"github.com/burnt-toast/mlcc-data-analysis/report"
)

func main() {
	programMap := make(map[string]*program.Instance)
	readFile("./2017NewSystemData.xlsx", programMap)
	fmt.Println("Unique instances of programs: {}", len(programMap))
	writer := report.Writer{ProgramData: programMap}
	writer.GenerateAttendanceReport()
}

//Reads the file into a map of program instances
func readFile(path string, programMap map[string]*program.Instance) {
	xlsx, err := excelize.OpenFile(path)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Get all the rows in the Sheet1.
	rows := xlsx.GetRows("Sheet1")
	rowCount := 0
	for _, row := range rows {
		cellCount := 0
		if rowCount != 0 { //do not process the header row
			programInstance := program.Instance{}
			for _, colCell := range row {
				processCell(&programInstance, colCell, cellCount)
				cellCount++
			}
			if !strings.Contains(programInstance.EventName, "CANCELLED") {
				uniqueID := programInstance.EventName + programInstance.StartDate
				if val, ok := programMap[uniqueID]; ok {
					val.Attendance++
				} else {
					programMap[uniqueID] = &programInstance
				}
			}
		}
		rowCount++
		cellCount = 0
	}
	fmt.Println("Rows Processed: {}", rowCount-1) //dont count the header
}

//Populates the appropriate program.Instance field based on the cellCount
func processCell(programInstance *program.Instance, colCell string, cellCount int) {
	switch cellCount {
	case 0:
		programInstance.Source = colCell
	case 1:
		programInstance.Category = colCell
	case 2:
		programInstance.GenericName = colCell
	case 3:
		programInstance.EventName = colCell
	case 4:
		programInstance.Capacity, _ = strconv.Atoi(colCell)
	case 7:
		programInstance.StartDate = colCell
	case 8:
		programInstance.StartTime = colCell
	}
}
