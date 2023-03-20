package main

import (
	"log"
	"toolkit"
)

func main() {
	var tools toolkit.Tools
	toSlug := "now!~23 is the time 123"
	slugify, err := tools.Slugify(toSlug)
	if err != nil {
		log.Println(err)
	}
	log.Println(slugify)
}
