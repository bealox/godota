package godota

import (
	"io/ioutil"
	"log"
	"path/filepath"
	"testing"
)

func TestGetAbility(t *testing.T) {

	b, err := ioutil.ReadFile(filepath.Join("examples/", "ability.json"))
	if err != nil {
		t.Error("Cannot read the file ", err)
	}

	ability := &Ability{}
	ability.Dota2ParentAPI.Data = b

	err = ability.getJSON()
	if err != nil {
		t.Error("Unable to get JSON from API ", err)
	}
	log.Println(len(ability.JSON.Abilities))
	if len(ability.JSON.Abilities) == 0 {
		t.Error("Unable to get JSON")
	}

}
