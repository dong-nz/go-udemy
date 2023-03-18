package hello

import (
	"fmt"
	"greeting"
	"log"
)

func main() {
	log.SetFlags(3)
	log.SetPrefix("greeting: ")
	hello, _ := greeting.Hello("Dong")
	fmt.Println(hello)

	message, err := greeting.Hello("")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)

}
