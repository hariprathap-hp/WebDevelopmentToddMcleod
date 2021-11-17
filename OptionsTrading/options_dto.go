package main

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

type ToHtml struct {
	Expiry   string
	Fstrike1 int
	Fstrike2 int
	Strike1  int
	Strike2  int
	BankInfo Options
}
