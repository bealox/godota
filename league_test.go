package godota

import (
	"io/ioutil"
	"path/filepath"
	"testing"
)

func TestGetLeague(t *testing.T) {

	b, err := ioutil.ReadFile(filepath.Join("examples/", "league.json"))
	if err != nil {
		t.Error("Cannot read the file ", err)
	}

	league := &League{}
	league.Dota2ParentAPI.Data = b

	err = league.getJSON()
	if err != nil {
		t.Error("Unable to get JSON from API ", err)
	}
	if len(league.JSON.Result.Leagues) == 0 {
		t.Error("Unable to get JSON")
	}

}
