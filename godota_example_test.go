package godota

import (
	"log"
	"os"
)

func ExampleStoreHeroImage() {
	pwd, err := os.Getwd()
	if err != nil {
		log.Println("unable to reach your current pwd ", err)
	}

	err = RemoveContents(pwd + "/heroes/")
	if err != nil {
		log.Println("Unable remove content")
	}

	err = StoreHeroImage("npc_dota_hero_antimage", pwd+"/heroes/")

	if err != nil {
		log.Println("unable to store the image ", err)
	}
	_, err = os.Open(pwd + "/heroes/antimage_sb.png")

	if err != nil {
		log.Println("unable to get the file ", err)
	}
}
