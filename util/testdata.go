package util

import "os"

func GetTestDataFolder() string {
	cur, err := os.Getwd()
	if err != nil {
		return ""
	}
	result := cur + "/../test/data/"
	return result
}
