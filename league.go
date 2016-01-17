package godota

import (
	"bytes"
	"encoding/json"
	"net/url"
)

//leagueInfoJson array of json
type (
	League struct {
		Dota2ParentAPI
		JSON *LeagueContent
	}

	leagueInfoJSON struct {
		Name          string `json:"name"`
		LeagueID      uint   `json:"leagueid"`
		Desc          string `json:"description"`
		TournamentURL string `json:"tournament_url"`
		ItemDef       uint   `json:"itemdef"`
		FileName      string `json:"-"`
	}

	// LeagueResult getting json data back
	LeagueContent struct {
		Result struct {
			Leagues []leagueInfoJSON `json:"leagues"`
		} `json:"result"`
	}
)

func (h *League) GetJSON() error {

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

func (h *League) getJSON() error {
	err := json.NewDecoder(bytes.NewBuffer(h.Data)).Decode(&h.JSON)
	if err != nil {
		return err
	}
	return nil
}

func (a *League) getURL() string {
	v := url.Values{}
	v.Add("key", apiKey)
	v.Add("language", "en_us")
	return "https://api.steampowered.com/IDOTA2Match_570/GetLeagueListing/v001/?" + v.Encode()
}
