package stringutils

import "github.com/google/uuid"

func ConvertToUUID(input string) uuid.UUID {
	parsedUUID, err := uuid.Parse(input)
	if err != nil {
		return uuid.Nil
	}
	return parsedUUID
}
