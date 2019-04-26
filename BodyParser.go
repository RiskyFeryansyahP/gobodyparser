package gobodyparser

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
)

func BodyParser(data io.ReadCloser) (bodyParser map[string]interface{}) {
	var body interface{}

	datas, err := ioutil.ReadAll(data)
	if err != nil {
		log.Fatal("Error Read Body Request", err.Error())
	}

	err = json.Unmarshal(datas, &body)
	if err != nil {
		log.Fatal("Error Unmarshal Body Request", err.Error())
	}

	bodyParser = body.(map[string]interface{})
	return bodyParser
}
