package GermanPostcodes

import (
	"encoding/json"
	"github.com/pkg/errors"
	"os"
)

type State string

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
type ZipcodeFinder struct {
	ZipcodeMap map[string][]*ZipcodeDateset
}

func InitZipcodeFinder(zipcodePath string) (ZipcodeFinder, error) {
	this := ZipcodeFinder{}
	var err error
	this.ZipcodeMap, err = GetZipCodeMap(zipcodePath)
	if err != nil {
		return this, errors.Wrap(err, "Failed to loadZipcodeDate")
	}
	return this, nil
}

func GetZipCodeMap(zipcodePath string) (map[string][]*ZipcodeDateset, error) {

	data, err := os.ReadFile(zipcodePath)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to Load ZipcodeJSON")
	}
	GermanZipcodes := make([]*ZipcodeDateset, 0)
	err = json.Unmarshal(data, &GermanZipcodes)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to unmarshall ZipcodeJSON")
	}

	GermanZipcodeMap := make(map[string][]*ZipcodeDateset)
	for _, ZipcodeData := range GermanZipcodes {
		if _, contains := GermanZipcodeMap[ZipcodeData.Zipcode]; !contains {
			GermanZipcodeMap[ZipcodeData.Zipcode] = make([]*ZipcodeDateset, 0)
		}
		GermanZipcodeMap[ZipcodeData.Zipcode] = append(GermanZipcodeMap[ZipcodeData.Zipcode], ZipcodeData)
	}

	return GermanZipcodeMap, nil

}

func GetStateFromZipCode(zipcodePath, postcode string) (State, error) {
	GermanZipcodeMap, err := GetZipCodeMap(zipcodePath)
	if err != nil {
		return "", errors.Wrap(err, "Failed to load ZicodeData")
	}

	if ZipcodeData, containts := GermanZipcodeMap[postcode]; containts {
		return State(ZipcodeData[0].State), nil
	}
	return "", errors.New("Zipcode not found")
}

func (z ZipcodeFinder) GetStateFromZipCode(postcode string) (State, error) {

	if ZipcodeData, containts := z.ZipcodeMap[postcode]; containts {
		return State(ZipcodeData[0].State), nil
	}
	return "", errors.New("Zipcode not found")
}
