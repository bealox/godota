package godota

import (
	"bytes"
	"encoding/json"
)

type (
	Ability struct {
		Dota2ParentAPI
		JSON *AbilityContent
	}

	gAbility struct {
		Name string `json:"name"`
		ID   string `json"id"`
	}

	AbilityContent struct {
		Abilities []gAbility `json:"abilities"`
	}
)

//GetJSON retrieve JSON from the API
func (a *Ability) GetJSON() error {

	a.Dota2ParentAPI.URL = a.getURL()

	b, err := a.getData()

	if err != nil {
		return err
	}

	a.Data = b

	err = a.getJSON()

	if err != nil {
		return err
	}

	return nil
}

func (h *Ability) getJSON() error {
	err := json.NewDecoder(bytes.NewBuffer(h.Data)).Decode(&h.JSON)
	if err != nil {
		return err
	}
	return nil
}

func (a *Ability) getURL() string {
	return "https://raw.githubusercontent.com/kronusme/dota2-api/master/data/abilities.json"
}
