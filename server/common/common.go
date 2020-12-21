package common

import "github.com/google/uuid"

//GenUID - Generation UUID string
func GenUID() string {
	uid := uuid.New()
	return uid.String()
}
