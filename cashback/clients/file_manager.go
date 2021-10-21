package clients

import (
	"encoding/json"
	"errors"
	"io/ioutil"

	"sleep.com/cashback/entities"
)

type FileManager struct {
}

func NewFileManager() *FileManager {
	fm := &FileManager{}
	return fm
}

func (*FileManager) GetDeals() ([]entities.Deal, error) {
	buf, err := ioutil.ReadFile("resources/deals.json")
	if err != nil {
		return nil, errors.New("Error retrieving deals")
	}
	s := string(buf)
	var deals []entities.Deal
	err = json.Unmarshal([]byte(s), &deals)
	if err != nil {
		return nil, errors.New("Error retrieving deals")
	}
	return deals, nil
}
