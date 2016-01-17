package godota

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestGetItem(t *testing.T) {

	b, err := ioutil.ReadFile(filepath.Join("examples/", "item.json"))
	if err != nil {
		t.Error("Cannot read the file ", err)
	}

	item := &Item{}
	item.Dota2ParentAPI.Data = b

	err = item.getJSON()
	if err != nil {
		t.Error("Unable to get JSON from API ", err)
	}

	if item.JSON.Result.Status != 200 {
		t.Error("Unable to get JSON status: ", item.JSON.Result.Status)
	}

}

func TestStoreItemImage(t *testing.T) {

	pwd, err := os.Getwd()
	if err != nil {
		t.Error("unable to reach your current pwd ", err)
	}

	err = RemoveContents(pwd + "/items/")
	if err != nil {
		t.Error("Unable remove content")
	}

	err = StoreItemImage("item_blink", pwd+"/items/")

	if err != nil {
		t.Error("unable to store the image ", err)
	}
	_, err = os.Open(pwd + "/items/blink_lg.png")

	if err != nil {
		t.Error("unable to get the file ", err)
	}
}
