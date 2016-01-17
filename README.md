# godota
The Godota package provides a quick way to grab data from the Dota2 API.

### Documentation
  [API Reference](https://godoc.org/github.com/bealox/godota)

### Installation 
 * go get github.com/bealox/godota
 * set an env variable called DOTA2API and assign api key to it

### Examples  
 ```golang
  hero := &Hero{}
  json, err := hero.GetJson() // this will return you an JSON struct.
  
  err = StoreHeroImage("npc_dota_hero_antimage", dir) // this will get you anti-mage's portrait.
 ```
### Update
Currently you can grab data for heroes, items, abiliies, leagues and teams. I might add live matchs and match history in the future. 
  


