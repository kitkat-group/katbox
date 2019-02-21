package kontent

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	log "github.com/Sirupsen/logrus"

	"github.com/kitkat-group/katbox/pkg/ksettings"
)

//Retrieve -
func Retrieve(preferences *ksettings.User) (*Articles, error) {
	log.Debugf("Retrieving content from [%s]", preferences.ContentURL)
	res, err := http.Get(preferences.ContentURL)

	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		panic(err.Error())
	}

	var a Articles
	json.Unmarshal(body, &a)
	return &a, nil
}
