package GermanPostcodes

import (
	"encoding/json"
	"github.com/pkg/errors"
	"os"
)

type state string

type ZipcodeDateset struct {
	CountryCode   string `json:"country_code"`
	Zipcode       string `json:"zipcode"`
	Place         string `json:"place"`
	State         string `json:"state"`
	StateCode     string `json:"state_code"`
	Province      string `json:"province"`
	ProvinceCode  string `json:"province_code"`
	Community     string `json:"community"`
	CommunityCode string `json:"community_code"`
	Latitude      string `json:"latitude"`
	Longitude     string `json:"longitude"`
}

func GetStateForPostCode(postcode string) (state, error) {

	data, err := os.ReadFile("DE/zipcodes.de.json")
	if err != nil {
		return "", errors.Wrap(err, "Failed to Load ZipcodeJSON")
	}
	GermanZipcodes := make([]*ZipcodeDateset, 0)
	err = json.Unmarshal(data, GermanZipcodes)
	if err != nil {
		return "", errors.Wrap(err, "Failed to unmarshall ZipcodeJSON")
	}

	GermanZipcodeMap := make(map[string]*ZipcodeDateset)
	for _, Zipcode := range GermanZipcodes {
		ger
	}

}
