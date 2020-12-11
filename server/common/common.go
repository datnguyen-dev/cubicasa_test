package common

import "github.com/google/uuid"

//GenUID -
func GenUID() string {
	uid := uuid.New()
	return uid.String()
}
