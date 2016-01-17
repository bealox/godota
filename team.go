package godota

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

//TeamInfoJSON Json structure

type (
	//Team
	Team struct {
		Dota2ParentAPI
		JSON   *TeamContent
		TeamID int64
	}

	gTeam struct {
		TeamID      int64       `json:"team_id"`
		Name        string      `json:"name"`
		Tag         string      `json:"tag"`
		CreatedBy   int64       `json:"time_created"`
		Rating      interface{} `json:"rating"`
		Logo        int64       `json:"logo"`
		LogoSponsor int64       `json:"logo_sponsor"`
		CountryCode string      `json:"country_code"`
		URL         string      `json:"url"`
		Player0     int64       `json:"player_0_account_id"`
		Player1     int64       `json:"player_1_account_id"`
		Player2     int64       `json:"player_2_account_id"`
		Player3     int64       `json:"player_3_account_id"`
		Player4     int64       `json:"player_4_account_id"`
		SubPlayer   int64       `json:"player_5_account_id"` //subtitude
	}

	//TeamContent json structure
	TeamContent struct {
		Result struct {
			Teams []gTeam `json:"teams"`
		} `json:"result"`
	}

	TeamLogoJSON struct {
		Data struct {
			FileName string `json:"filename"`
			URL      string `json:"url"`
			Size     int    `json:"size"`
		} `json:"data"`
	}
)

func (t *Team) GetJSON() error {

	t.Dota2ParentAPI.URL = t.getURL()
	//https://api.steampowered.com/IDOTA2Match_570/GetTeamInfoByTeamID/v001/?key=?&start_at_team_id=543897&teams_requested=1
	b, err := t.getData()

	if err != nil {
		return err
	}

	t.Data = b

	err = t.getJSON()

	if err != nil {
		return err
	}

	return nil
}

func (t *Team) getJSON() error {
	err := json.NewDecoder(bytes.NewBuffer(t.Data)).Decode(&t.JSON)
	if err != nil {
		return err
	}
	return nil
}

//getURL only give you show one team at a time, and team id is required
func (a *Team) getURL() string {
	v := url.Values{}
	v.Add("key", apiKey)
	v.Add("start_at_team_id", strconv.FormatInt(a.TeamID, 10))
	v.Add("teams_requested", "1")
	return "https://api.steampowered.com/IDOTA2Match_570/GetTeamInfoByTeamID/v001/?" + v.Encode()
}

//GetTeamLogo team logo details
func GetTeamLogo(id int64) (*TeamLogoJSON, error) {

	v := url.Values{}
	v.Add("ugcid", strconv.FormatInt(id, 10))
	v.Add("key", apiKey)
	v.Add("appid", "570")
	api := "http://api.steampowered.com/ISteamRemoteStorage/GetUGCFileDetails/v1/?" + v.Encode()
	resp, err := http.Get(api)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var document *TeamLogoJSON
	if err = json.NewDecoder(resp.Body).Decode(&document); err != nil {
		return nil, err
	}

	return document, nil
}

// StoreTeamLogo is to store image into the filesystem
func StoreTeamLogo(logo *TeamLogoJSON, dir string, fileName string) error {
	rep, err := http.Get(logo.Data.URL)
	if err != nil {
		return err
	}
	defer rep.Body.Close()

	//if logo dir doesn't exist then create one

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.Mkdir(dir, 0755)
	}

	file := dir + fileName + ".png"
	// if team logo doesn't exist then create one
	if _, err := os.Stat(file); os.IsNotExist(err) {
		file, err := os.Create(file)
		defer file.Close()
		if err != nil {
			return err
		}
		_, err = io.Copy(file, rep.Body)
		if err != nil {
			return err
		}
	}

	return nil
}
