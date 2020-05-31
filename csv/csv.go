package csv

import (
	"automatedYTCoding/apiadapter"
	"encoding/csv"
	"os"
	"reflect"
	"strconv"
	"strings"
)

func CreateVideoCSV(filePath string, videos map[string]*apiadapter.Video) {
	file, e := os.Create(filePath)
	if e != nil {
		panic(e)
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()
	var row []string
	// HEADER
	for _, video := range videos {
		val := reflect.Indirect(reflect.ValueOf(video))
		row = make([]string, val.Type().NumField())
		for i := 0; i < val.Type().NumField(); i++ {
			row[i] = val.Type().Field(i).Name
		}
		writer.Write(row)
		break
	}
	// ROWS
	for _, video := range videos {
		val := reflect.Indirect(reflect.ValueOf(video))
		for i := 0; i < val.Type().NumField(); i++ {
			row = make([]string, val.Type().NumField())
			for i := 0; i < val.Type().NumField(); i++ {
				row[i] = getData(val, i)
			}
		}
		writer.Write(row)
	}

}

func CreateCSV(filePath string, values map[string]int) {
	file, e := os.Create(filePath)
	if e != nil {
		panic(e)
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()
	var row []string
	// HEADER
	row = append(row, "id", "sentiment")
	writer.Write(row)
	for id, sentiment := range values {
		row = make([]string, 2)
		row[0] = id
		row[1] = strconv.Itoa(sentiment)
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
	return ""
}
