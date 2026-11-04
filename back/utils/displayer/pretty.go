package displayer

import (
	"encoding/json"
	"fmt"
	"log"
)

func PrettyDisplay(message string, v interface{}) {
	log.Printf("====%s====\n", message)
	empJSON, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println(string(empJSON))
	fmt.Println("==================================")
}
