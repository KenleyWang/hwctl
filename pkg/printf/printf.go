package printf

import (
	"fmt"
	"os"
	"reflect"
	"text/tabwriter"
)

func PrintSlice(slice interface{}) error {
	sliceValue := reflect.ValueOf(slice)

	if sliceValue.Kind() != reflect.Ptr || sliceValue.Elem().Kind() != reflect.Slice {
		return fmt.Errorf("input is not a pointer to a slice")
	}

	// 创建一个 TabWriter
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)

	// 打印键
	if sliceValue.Elem().Len() > 0 {
		item := sliceValue.Elem().Index(0)
		typeOfItem := item.Type()

		for i := 0; i < typeOfItem.NumField(); i++ {
			fieldName := typeOfItem.Field(i).Name
			if _, err := fmt.Fprintf(w, "%s\t", fieldName); err != nil {
				return fmt.Errorf("error writing to TabWriter: %v", err)
			}
		}
		if _, err := fmt.Fprintln(w); err != nil {
			return fmt.Errorf("error writing to TabWriter: %v", err)
		}
	}

	// 打印值
	for i := 0; i < sliceValue.Elem().Len(); i++ {
		item := sliceValue.Elem().Index(i)
		for j := 0; j < item.NumField(); j++ {
			fieldValue := fmt.Sprintf("%v", item.Field(j).Interface())
			if _, err := fmt.Fprintf(w, "%s\t", fieldValue); err != nil {
				return fmt.Errorf("error writing to TabWriter: %v", err)
			}
		}
		if _, err := fmt.Fprintln(w); err != nil {
			return fmt.Errorf("error writing to TabWriter: %v", err)
		}
	}

	// 刷新 TabWriter
	if err := w.Flush(); err != nil {
		return fmt.Errorf("error flushing TabWriter: %v", err)
	}

	return nil
}
