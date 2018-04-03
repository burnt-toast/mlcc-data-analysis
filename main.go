package main

import (
	"fmt"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/burnt-toast/mlcc-data-analysis/program"
)

func main() {
	xlsx, err := excelize.OpenFile("./2017programdata.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Get all the rows in the Sheet1.
	rows := xlsx.GetRows("Sheet1")
	rowCount := 0
	programMap := make(map[string]*program.Instance)
	for _, row := range rows {
		cellCount := 0
		if rowCount != 0 { //do not process the header row
			programInstance := program.Instance{}
			for _, colCell := range row {
				processCell(&programInstance, colCell, cellCount)
				cellCount++
			}
			uniqueID := programInstance.EventName + programInstance.StartDate

			if val, ok := programMap[uniqueID]; ok {
				val.Attendance++
			} else {
				programMap[uniqueID] = &programInstance
			}
		}
		rowCount++
		cellCount = 0
	}
	//fmt.Println(programMap["Cold Photo02-18-17"].Attendance)
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
	}
}
