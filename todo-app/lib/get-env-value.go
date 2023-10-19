package lib

import (
	"log"
	"os"
	"strings"
)

func GetEnvFileValue(envPath string, key string) string {
	var valueToReturn string = ""
	data, err := os.ReadFile(envPath)

	if err != nil {
		panic(err)
	}
	arrayFile := strings.Split(string(data), "\n")

	for i := 0; i < len(arrayFile); i++ {
		splitedString := strings.Split(arrayFile[i], "=")

		if len(splitedString) == 2 && splitedString[0] == key {
			valueToReturn = strings.TrimSuffix(splitedString[1], "\n")
		}
	}
	if valueToReturn == "" {
		log.Fatal("Nothing found for the key ", key)
	}
	return valueToReturn
}
