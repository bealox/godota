package godota

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type (
	Item struct {
		Dota2ParentAPI
		JSON *ItemContent
	}

	//ItemContent based on the JSON sturcture
	ItemContent struct {
		Result struct {
			Items  []gItemInfo `json:"items"`
			Status int         `json:"status"`
		} `json:"result"`
	}

	gItemInfo struct {
		DotaID     int    `json:"id"`             // the ID used to identify the item in the api
		DotaName   string `json:"name"`           // the code name of the item
		Cost       int    `json:"cost"`           // the gold cost of the item
		SecretShop int    `json:"secret_shop"`    // 1 if is available, 0 otherwise
		SideShop   int    `json:"side_shop"`      // 1 if is available, 0 otherwise
		Recipt     int    `json:"recipe"`         // 1 if is available, 0 otherwise
		Name       string `json:"localized_name"` // if a lang specified, this will show the in game name of the item for that language
	}
)

func (i *Item) GetJSON() error {

	i.Dota2ParentAPI.URL = i.getURL()

	b, err := i.getData()

	if err != nil {
		return err
	}

	i.Data = b

	err = i.getJSON()

	if err != nil {
		return err
	}

	return nil
}

func (h *Item) getJSON() error {
	err := json.NewDecoder(bytes.NewBuffer(h.Data)).Decode(&h.JSON)
	if err != nil {
		return err
	}
	return nil
}

func (i *Item) getURL() string {
	v := url.Values{}
	v.Add("key", apiKey)
	v.Add("language", "en_us")
	return "https://api.steampowered.com/IEconDOTA2_570/GetGameItems/V001/?" + v.Encode()
}

//StoreItemImage stores item's image into a dir
func StoreItemImage(dotaName string, dir string) error {
	//large horizontal portrait - 205x11px lg.png

	extentions := []string{"lg.png"}
	apiURL := "http://cdn.dota2.com/apps/dota2/images/items/"
	fileName := strings.Replace(dotaName, "item_", "", -1)

	for _, ext := range extentions {
		path := dir + fileName + "_" + ext
		if _, err := os.Stat(path); os.IsNotExist(err) {
			resp, err := http.Get(apiURL + fileName + "_" + ext)

			if err != nil {
				return err
			}

			file, err := os.Create(path)
			if err != nil {
				return err
			}

			_, err = io.Copy(file, resp.Body)
			resp.Body.Close()
			file.Close()
			if err != nil {
				return err
			}
		}

	}
	return nil
}
