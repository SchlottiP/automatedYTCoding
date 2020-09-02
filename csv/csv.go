package csv

import (
	"encoding/csv"
	"os"
	"reflect"
	"strconv"
	"strings"
)

func CreateCSV(filePath string, values []interface{}) {
	file, e := os.Create(filePath)
	if e != nil {
		panic(e)
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()
	var row []string
	// HEADER
	for _, val := range values {
		if val == nil {
			continue
		}
		val := reflect.Indirect(reflect.ValueOf(val))
		row = make([]string, val.Type().NumField())
		for i := 0; i < val.Type().NumField(); i++ {
			row[i] = val.Type().Field(i).Name
		}
		writer.Write(row)
		break
	}
	// ROWS
	for _, val := range values {
		if val == nil {
			continue
		}
		val := reflect.Indirect(reflect.ValueOf(val))
		for i := 0; i < val.Type().NumField(); i++ {
			row = make([]string, val.Type().NumField())
			for i := 0; i < val.Type().NumField(); i++ {
				row[i] = getData(val, i)
			}
		}
		writer.Write(row)
	}

}

func getData(val reflect.Value, index int) string {
	switch val.Field(index).Kind() {
	case reflect.Uint64:
		return strconv.FormatInt(int64(val.Field(index).Uint()), 10)
	case reflect.Int:
		return strconv.FormatInt(val.Field(index).Int(), 10)
	case reflect.Float64:
		return strconv.FormatFloat(val.Field(index).Float(), 'f', 6, 64)
	case reflect.Slice:
		{
			sliceData := make([]string, val.Field(index).Len())
			for i := 0; i < val.Field(index).Len(); i++ {
				sliceData[i] = val.Field(index).Index(i).String()
			}
			return strings.Join(sliceData, ",")
		}
	default:
		return val.Field(index).String()
	}
}
