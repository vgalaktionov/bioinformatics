package lib

import "log"

func Check(err error) {
	if err != nil {
		log.Fatalf("%s", err)
	}
}
