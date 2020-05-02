package holiday

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
)

// key: 20200501, value: ["劳动节"]
var holidayCache = make(map[string][]string)

func init() {
	buff, err := ioutil.ReadFile("data.json")
	checkErr(err)
	json.Unmarshal(buff, &holidayCache)
}

// Query 查询
func Query(date string) ([]string, error) {
	if day, ok := holidayCache[date]; ok {
		return day, nil
	}
	return nil, errors.New("This date " + date + " is not a holiday")
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
