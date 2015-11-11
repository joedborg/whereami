package whereami

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type IPData struct {
	As          string
	City        string
	Country     string
	CountryCode string
	Isp         string
	Lat         float64
	Lon         float64
	Org         string
	Query       string
	Region      string
	RegionName  string
	Status      string
	Timezone    string
	Zip         string
}

// GetLocationData returns the location data from ip-api.com
// based on your external IP address.
func GetLocationData() (IPData, error) {
	var data IPData
	ip, err := GetExternalIP()
	if err != nil {
		return data, err
	}
	url := "http://ip-api.com/json/" + ip
	resp, err := http.Get(url)
	if err != nil {
		return data, fmt.Errorf("GET %q: %v", url, err)
	}
	defer resp.Body.Close()

	var buf bytes.Buffer
	dec := json.NewDecoder(io.TeeReader(resp.Body, &buf))
	if err = dec.Decode(&data); err == nil {
		return data, nil
	}

	return data, fmt.Errorf("Parse %q: %v", buf.String(), err)
}

// GetExternalIP returns your external IP address as returned by http://echoip.com.
func GetExternalIP() (string, error) {
	url := "http://echoip.com"
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("GET %q: %v", url, err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return string(body), err
}
