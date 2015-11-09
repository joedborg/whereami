package whereami

import (
	"encoding/json"
	"fmt"
	"github.com/joedborg/getexternalip"
	"io/ioutil"
	"net/http"
)

type IPData struct {
	As          string
	City        string
	Country     string
	CountryCode string
	Isp         string
	Lat         string
	Lon         string
	Org         string
	Query       string
	Region      string
	RegionName  string
	Status      string
	Timezone    string
	Zip         string
}

func errorCheck(err error) {
	if err != nil {
		panic(err)
	}
}

func GetLocationData() IPData {
	ip := getexternalip.GetExternalIP()
	url := fmt.Sprintf("http://ip-api.com/json/%s", ip)
	resp, err := http.Get(url)
	errorCheck(err)
	body, err := ioutil.ReadAll(resp.Body)
	errorCheck(err)

	var data IPData
	err = json.Unmarshal(body, &data)

	return data
}
