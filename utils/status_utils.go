package utils

import "github.com/vnoitkumar/demyst-code-kata/constants"

func GetStatus(isCompleted bool) (status string) {
	statusMap := map[bool]string{
		true:  constants.COMPLETED,
		false: constants.NOT_COMPLETED,
	}

	return statusMap[isCompleted]
}
