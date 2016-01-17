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
	//Hero is the structure
	Hero struct {
		Dota2ParentAPI
		JSON *HeroContent
	}

	//HeroContent is the json structure we are getting from dota2 api
	HeroContent struct {
		Result struct {
			Hereos []gHeroesInfo `json:"heroes"`
			Status int           `json:"status"`
			Count  int           `json:"count"`
		} `json:"result"`
	}

	gHeroesInfo struct {
		DotaName string `json:"name"`           // dota name
		DotaID   int    `json:"id"`             // dota id
		Name     string `json:"localized_name"` // language name (default is english)
	}
)

//GetJSON retrieve JSON from the API
func (h *Hero) GetJSON() error {

	h.Dota2ParentAPI.URL = h.getURL()

	b, err := h.getData()

	if err != nil {
		return err
	}

	h.Data = b

	err = h.getJSON()

	if err != nil {
		return err
	}

	return nil
}

func (h *Hero) getJSON() error {
	err := json.NewDecoder(bytes.NewBuffer(h.Data)).Decode(&h.JSON)
	if err != nil {
		return err
	}
	return nil
}

func (h *Hero) getURL() string {
	v := url.Values{}
	v.Add("key", apiKey)
	v.Add("language", "en_us")
	return "https://api.steampowered.com/IEconDOTA2_570/GetHeroes/v0001/?" + v.Encode()
}

//StoreImage store images from the API
func StoreHeroImage(dotaName string, dir string) error {
	//small horizontal portrait - 59x33px sb.png
	//large horizontal portrait - 205x11px lg.png
	//full quality horizontal portrait - 256x114px full.png
	//full quality vertical portrait - 234x272px vert.jpg

	extentions := []string{"sb.png", "lg.png", "full.png", "vert.jpg"}
	apiURL := "http://cdn.dota2.com/apps/dota2/images/heroes/"
	fileName := strings.Replace(dotaName, "npc_dota_hero_", "", -1)

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
