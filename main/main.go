package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"mytest/src/scraler"
	"os"
	"sort"
)

func writeToCsv(s [][]string, csvName string) {
	file, err := os.OpenFile(csvName, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("open file is failed, err: ", err)
	}
	defer file.Close()
	// 写入UTF-8 BOM，防止中文乱码
	file.WriteString("\xEF\xBB\xBF")
	for index, value := range s {
		fmt.Println(index, value)
		w := csv.NewWriter(file)
		w.Write(value)
		// 写文件需要flush，不然缓存满了，后面的就写不进去了，只会写一部分
		w.Flush()
	}
}

func readFromCsv() [][]string {
	file, err := os.OpenFile("example.csv", os.O_RDONLY, 0644)
	if err != nil {
		fmt.Println("open file is failed, err: ", err)
	}
	defer file.Close()
	csvReader := csv.NewReader(file)
	rows, err := csvReader.ReadAll() // `rows` is of type [][]string
	if err != nil {
		panic(err)
	}
	return rows
}

func getName(s [][]string) []string {
	nameList := make([]string, 0)
	for _, row := range s {
		fmt.Println(row)
		name := bytes.TrimPrefix([]byte(row[0]), []byte{239, 187, 191})
		nameList = append(nameList, string(name))
	}
	return nameList
}

func RemoveRepByMap(arr []string) []string {
	result := []string{}
	tmpMap := make(map[string]interface{})
	for _, val := range arr {
		if _, ok := tmpMap[val]; !ok {
			result = append(result, val)
			tmpMap[val] = struct{}{}
		}
	}
	sort.Strings(result)
	return result
}

func genSortCsv(dataArr [][]string, sortNameArr []string)  {
	sortDataArr := make([][]string, 0)
	for _, name := range sortNameArr{
		for _, row:= range dataArr{
			if row[0] == name{
				sortDataArr = append(sortDataArr, row)
			}
		}
	}
	writeToCsv(sortDataArr, "sortExample.csv")
}



func main() {
	years := []string{"2000", "2001",  "2002",  "2003",
		"2004",  "2005",  "2006",  "2007",  "2008",  "2009",  "2010",  "2011",  "2012",  "2013",  "2014",
		"2015",  "2016", "2017"}
	for _, year := range years{
		country_data := scraler.Scwler(year)
		country_data = country_data[:11]
		writeToCsv(country_data, "example.csv")
	}


	csvData := readFromCsv()
	countryName := getName(csvData)
	uniqName := RemoveRepByMap(countryName)
	fmt.Println(uniqName)
	genSortCsv(csvData, uniqName)
}

