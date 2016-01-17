package godota

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestGetTeam(t *testing.T) {

	b, err := ioutil.ReadFile(filepath.Join("examples/", "team.json"))
	if err != nil {
		t.Error("Cannot read the file ", err)
	}

	team := &Team{}
	team.Dota2ParentAPI.Data = b

	err = team.getJSON()
	if err != nil {
		t.Error("Unable to get JSON from API ", err)
	}
	if len(team.JSON.Result.Teams) == 0 {
		t.Error("Unable to get JSON")
	}

}

func TestGetTeamFromAPI(t *testing.T) {
	team := &Team{
		TeamID: 543897,
	}
	err := team.GetJSON()
	if err != nil {
		t.Error("Unable to get JSON From API ", err)
	}

	if len(team.JSON.Result.Teams) == 0 {
		t.Error("Unable to get data from the API")
	}
}

func TestStoreLogo(t *testing.T) {

	pwd, err := os.Getwd()
	if err != nil {
		t.Error("unable to reach your current pwd ", err)
	}

	logo, err := GetTeamLogo(612792385922460575)
	if err != nil {
		t.Error("Unable to get the logo ", err)
	}

	err = RemoveContents(pwd + "/teams/")
	if err != nil {
		t.Error("Unable remove content")
	}

	err = StoreTeamLogo(logo, pwd+"/teams/", "mski")
	if err != nil {
		t.Error("Unable to store image ", err)
	}

	_, err = os.Open(pwd + "/teams/mski.png")

	if err != nil {
		t.Error("unable to get the file ", err)
	}
}
