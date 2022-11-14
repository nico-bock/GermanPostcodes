package GermanPostcodes

import (
	"encoding/json"
	"github.com/pkg/errors"
	"os"
)

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

func InitZipcodeFinder() (ZipcodeFinder, error) {
	this := ZipcodeFinder{}
	var err error
	this.ZipcodeMap, err = GetZipCodeMap()
	if err != nil {
		return this, errors.Wrap(err, "Failed to loadZipcodeDate")
	}
	return this, nil
}

func GetZipCodeMap() (map[string][]*ZipcodeDateset, error) {

	data, err := os.ReadFile("DE/zipcodes.de.json")
	if err != nil {
		return nil, errors.Wrap(err, "Failed to Load ZipcodeJSON")
	}
	GermanZipcodes := make([]*ZipcodeDateset, 0)
	err = json.Unmarshal(data, GermanZipcodes)
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

func GetStateFromZipCode(postcode string) (string, error) {
	GermanZipcodeMap, err := GetZipCodeMap()
	if err != nil {
		return "", errors.Wrap(err, "Failed to load ZicodeData")
	}

	if ZipcodeData, containts := GermanZipcodeMap[postcode]; containts {
		return ZipcodeData[0].State, nil
	}
	return "", errors.New("Zipcode not found")
}

func (z ZipcodeFinder) GetStateFromZipCode(postcode string) (string, error) {

	if ZipcodeData, containts := z.ZipcodeMap[postcode]; containts {
		return ZipcodeData[0].State, nil
	}
	return "", errors.New("Zipcode not found")
}
