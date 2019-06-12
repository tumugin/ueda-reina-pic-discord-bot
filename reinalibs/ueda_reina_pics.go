package reinalibs

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"time"
)

type UedaReinaPics struct {
	Urls []string
}

func NewUedaReinaPics(jsonpath string) UedaReinaPics {
	bytes, err := ioutil.ReadFile(jsonpath)
	if err != nil {
		log.Fatal(err)
	}
	var reina UedaReinaPics
	if err := json.Unmarshal(bytes, &reina); err != nil {
		log.Fatal(err)
	}
	return reina
}

func (reina UedaReinaPics) GetRandomReinaPic() string {
	rand.Seed(time.Now().UnixNano())
	length := len(reina.Urls)
	randomKey := rand.Intn(length - 1)
	return reina.Urls[randomKey]
}
