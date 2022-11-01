package main

import (
	"encoding/json"
	"flag"
	"github.com/nico-bock/GermanPostcodes"
	"github.com/rs/zerolog/log"
	"os"
)

func main() {
	importFilename := flag.String("import-file", "DE/zipcodes.de.json", "The Path of the File to read from the Zipcode data.")
	exportFilename := flag.String("export-file", "./compressedZicodeStateMapping.json", "The Path of the File to write compressed Data to.")
	flag.Parse()

	data, err := os.ReadFile(*importFilename)
	if err != nil {
		log.Error().Err(err).Msg("Failed to LoadZipcodes")
	}
	GermanZipcodes := make([]*GermanPostcodes.ZipcodeDateset, 0)
	err = json.Unmarshal(data, &GermanZipcodes)
	if err != nil {
		log.Error().Err(err).Msg("Failed to LoadZipcodes")
	}

	GermanZipcodeMap := make(map[string]string)
	for _, ZipcodeData := range GermanZipcodes {
		if _, contains := GermanZipcodeMap[ZipcodeData.Zipcode]; contains {
			if GermanZipcodeMap[ZipcodeData.Zipcode] != ZipcodeData.State {
				storedState := GermanZipcodeMap[ZipcodeData.Zipcode]
				log.Info().Str("Zicode", ZipcodeData.Zipcode).Str("State1", storedState).Str("State2", ZipcodeData.State).Msg("Fould multiple States for Zipcode")
			}
		}
		GermanZipcodeMap[ZipcodeData.Zipcode] = ZipcodeData.State
	}

	data, err = json.Marshal(GermanZipcodeMap)
	err = os.WriteFile(*exportFilename, data, 777)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to Write to File")
	}
}
