package lib

import (
	"fmt"
	"github.com/tealeg/xlsx"
)

func MapsToExcel(data []map[string]string, filename string) error {
	// 创建新的Excel文件
	file := xlsx.NewFile()

	// 创建一个工作表
	sheet, err := file.AddSheet("Sheet1")
	if err != nil {
		return fmt.Errorf("Error creating sheet: %v", err)
	}

	// 添加表头行
	if len(data) > 0 {
		headerRow := sheet.AddRow()
		for key := range data[0] {
			headerCell := headerRow.AddCell()
			headerCell.Value = key
		}

		// 添加数据行
		for _, rowData := range data {
			dataRow := sheet.AddRow()
			for _, value := range rowData {
				valueCell := dataRow.AddCell()
				valueCell.Value = value
			}
		}
	}

	// 保存文件
	err = file.Save(filename)
	if err != nil {
		return fmt.Errorf("Error saving Excel file: %v", err)
	}

	fmt.Println("Excel file saved successfully.")
	return nil
}
