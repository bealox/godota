package godota

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestGetHero(t *testing.T) {

	b, err := ioutil.ReadFile(filepath.Join("examples/", "hero.json"))
	if err != nil {
		t.Error("Cannot read the file ", err)
	}

	hero := &Hero{}
	hero.Dota2ParentAPI.Data = b

	err = hero.getJSON()
	if err != nil {
		t.Error("Unable to get JSON from API ", err)
	}

	if hero.JSON.Result.Status != 200 {
		t.Error("Unable to get JSON status: ", hero.JSON.Result.Status)
	}

}

func TestStoreHeroImage(t *testing.T) {

	pwd, err := os.Getwd()
	if err != nil {
		t.Error("unable to reach your current pwd ", err)
	}

	err = RemoveContents(pwd + "/heroes/")
	if err != nil {
		t.Error("Unable remove content")
	}

	err = StoreHeroImage("npc_dota_hero_antimage", pwd+"/heroes/")

	if err != nil {
		t.Error("unable to store the image ", err)
	}
	_, err = os.Open(pwd + "/heroes/antimage_sb.png")

	if err != nil {
		t.Error("unable to get the file ", err)
	}
}
