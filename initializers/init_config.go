package initializers

import (
	"encoding/json"
	"os"

	"github.com/go-playground/validator"
	"github.com/rs/zerolog/log"
	"github.com/vnoitkumar/demyst-code-kata/configurations"
)

func LoadConfig(filePath string, configStruct *configurations.Config) (err error) {
	log.Info().Msgf("Opening config file from path %s", filePath)
	file, err := os.Open(filePath)
	if err != nil {
		log.Err(err).Msg("Error opening config file")
		return err
	}

	log.Info().Msg("Decoding config file to configStruct")
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&configStruct)
	if err != nil {
		log.Err(err).Msg("Error decoding config file")
		return err
	}

	log.Info().Msgf("Successfully loaded config from path %s", filePath)
	return nil
}

func ValidateConfig(configStruct configurations.Config) (err error) {
	log.Info().Msg("Validating configStruct")

	validate := validator.New()
	err = validate.Struct(configStruct)
	if err != nil {
		log.Err(err).Msg("Error validating config file")
		return err
	}

	log.Info().Msg("Successfully validated configStruct")
	return nil
}
