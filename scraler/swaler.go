package scraler

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strings"
)

func Scwler(year string) [][]string {
	// Request the HTML page.
	url := "https://www.kuaiyilicai.com/stats/global/yearly/g_gross_saving_current_usd/" + year + ".html"
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	countrys := make([][]string, 0)
	doc.Find("tr").Each(func(i int, tr *goquery.Selection) {
		// For each item found, get the value
		trTmp := make([]string, 0)
		saveData := make([]string, 0)
		tr.Find("td").Each(func(i int, td *goquery.Selection) {
			trTmp = append(trTmp, td.Text())
		})

		if len(trTmp) == 4 {
			money := strings.Trim(strings.Split(trTmp[3], "(")[1], ")")
			money = strings.ReplaceAll(money, ",", "")
			countryName := bytes.TrimPrefix([]byte(trTmp[1]), []byte{239, 187, 191})
			fmt.Println("新的国家名字：", countryName)
			saveData = append(saveData, string(countryName))
			saveData = append(saveData, string(countryName))
			saveData = append(saveData, money)
			saveData = append(saveData, year)
			countrys = append(countrys, saveData)
		}
	})
	return countrys
}
