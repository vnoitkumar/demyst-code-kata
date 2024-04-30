package utils

import (
	"github.com/rs/zerolog/log"
)

func GetEvenNumberedSlice(sliceSize int) (evenNumberSlice []int) {
	log.Info().Msgf("Iterating and creating evenNumberSlice")
	for i := 1; true; i++ {
		if i%2 == 0 {
			log.Info().Msgf("Appending even number %d to evenNumberSlice", i)
			evenNumberSlice = append(evenNumberSlice, i)
		}

		if len(evenNumberSlice) == sliceSize {
			log.Info().Msgf("Breaking the loop as the len of evenNumberSlice (%d) and sliceSize (%d) are equal", len(evenNumberSlice), sliceSize)
			break
		}
	}
	return
}
