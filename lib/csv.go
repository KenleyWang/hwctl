package lib

import (
	"encoding/csv"
	"fmt"
	"os"
)

func MapsToCSV(data []map[string]string, filename string) error {
	// 创建 CSV 文件
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("Error creating CSV file: %v", err)
	}
	defer file.Close()

	// 获取列顺序
	var keys []string
	if len(data) > 0 {
		for key := range (data)[0] {
			keys = append(keys, key)
		}
	}

	// 创建 CSV writer
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// 写入表头
	if len(keys) > 0 {
		err := writer.Write(keys)
		if err != nil {
			return fmt.Errorf("Error writing CSV header: %v", err)
		}

		// 写入数据行
		for _, row := range data {
			values := make([]string, 0, len(row))
			for _, key := range keys {
				values = append(values, row[key])
			}
			err := writer.Write(values)
			if err != nil {
				return fmt.Errorf("Error writing CSV row: %v", err)
			}
		}
	}

	fmt.Println("CSV file saved successfully.")
	return nil
}
