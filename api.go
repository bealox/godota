package godota

import (
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

var (
	apiKey = os.Getenv("DOTA2API")
)

//Dota2API an interface
type Dota2API interface {
	GetJSON() error
	getJSON() error
	getURL() string
}

type Dota2ParentAPI struct {
	Data []byte
	URL  string
}

//Get byte based on the http get data
func (h *Dota2ParentAPI) getData() ([]byte, error) {
	read, err := http.Get(h.URL)
	if err != nil {
		return nil, err
	}
	defer read.Body.Close()

	return ioutil.ReadAll(read.Body)
}

//RemoveContents This is mainly for testing purpose, remove all the files before store the imges again
func RemoveContents(dir string) error {
	d, err := os.Open(dir)
	if err != nil {
		return err
	}
	defer d.Close()
	names, err := d.Readdirnames(-1)
	if err != nil {
		return err
	}
	for _, name := range names {
		err = os.RemoveAll(filepath.Join(dir, name))
		if err != nil {
			return err
		}
	}
	return nil
}
