//Package options... refer this link for roundtrip : https://www.reddit.com/r/golang/comments/d8wohs/http_get_403_forbidden/
package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

var url = "https://www.nseindia.com/api/option-chain-indices?symbol=BANKNIFTY"

//var url = "https://nseoptions.s3.ap-south-1.amazonaws.com/data1.json"
var tmpl *template.Template
var parse_err error
var expiry_Date string

type Options struct {
	Records struct {
		Expirydates []string `json:"expiryDates"`
		Data        []struct {
			Strikeprice int    `json:"strikePrice"`
			Expirydate  string `json:"expiryDate"`
			Pe          struct {
				Strikeprice           int     `json:"strikePrice"`
				Expirydate            string  `json:"expiryDate"`
				Openinterest          int     `json:"openInterest"`
				Changeinopeninterest  int     `json:"changeinOpenInterest"`
				Pchangeinopeninterest float64 `json:"pchangeinOpenInterest"`
				Totaltradedvolume     int     `json:"totalTradedVolume"`
				Impliedvolatility     float64 `json:"impliedVolatility"`
				Lastprice             float64 `json:"lastPrice"`
				Change                float64 `json:"change"`
				Pchange               float64 `json:"pChange"`
			} `json:"PE,omitempty"`
			Ce struct {
				Strikeprice           int     `json:"strikePrice"`
				Expirydate            string  `json:"expiryDate"`
				Openinterest          int     `json:"openInterest"`
				Changeinopeninterest  int     `json:"changeinOpenInterest"`
				Pchangeinopeninterest float64 `json:"pchangeinOpenInterest"`
				Totaltradedvolume     int     `json:"totalTradedVolume"`
				Impliedvolatility     float64 `json:"impliedVolatility"`
				Lastprice             float64 `json:"lastPrice"`
				Change                float64 `json:"change"`
				Pchange               float64 `json:"pChange"`
			} `json:"CE,omitempty"`
		} `json:"data"`
	} `json:"records"`
}

/*type roundTripperStripUserAgent struct{}

func (a roundTripperStripUserAgent) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("User-Agent", "")
	return http.DefaultTransport.RoundTrip(req)
}*/

func fetchURL(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside Fetch URL")
	client := http.DefaultClient
	//client := &http.Client{Transport: roundTripperStripUserAgent{}}
	req, req_err := http.NewRequest(http.MethodGet, url, nil)
	if req_err != nil {
		panic(req_err)
	}

	//either use this method by changing user-agent to an empty string "" or use the rountTripperStripUserAgent above
	req.Header.Set("User-Agent", "")
	//fmt.Println(req)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	result, rd_err := ioutil.ReadAll(resp.Body)
	if rd_err != nil {
		fmt.Println(rd_err)
	}

	var option Options

	json.Unmarshal([]byte(result), &option)

	if r.Method == "POST" {
		r.ParseForm()
		expiry_Date = r.FormValue("expiry")
	} else {
		expiry_Date = option.Records.Expirydates[0]
	}

	var toHtml = struct {
		Expiry   string
		BankInfo Options
	}{
		Expiry:   expiry_Date,
		BankInfo: option,
	}

	tmpl, parse_err = template.ParseGlob("css/*.gohtml")

	if parse_err != nil {
		fmt.Printf("Error while parsing html template file %s", parse_err)
	}
	tmpl.Execute(w, toHtml)
}

func main() {
	//url := "https://nseoptions.s3.ap-south-1.amazonaws.com/data.json"

	http.HandleFunc("/banknifty", fetchURL)
	http.Handle("/", http.FileServer(http.Dir("css/")))
	list_err := http.ListenAndServe(":8000", nil)
	if list_err != nil {
		fmt.Println(list_err)
	}

}
