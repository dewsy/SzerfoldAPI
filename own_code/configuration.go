package own_code

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func ReadConfiguration() (conf map[string]string) {
	byteSlice, err := ioutil.ReadFile("conf.json")
	if err != nil {
		fmt.Printf("Could not open conf.json:")
		fmt.Printf(err.Error())
		os.Exit(1)
	}
	marshalErr := json.Unmarshal(byteSlice, &conf)
	if marshalErr != nil {
		fmt.Printf("Could not marshal json byteslice:")
		fmt.Printf(marshalErr.Error())
		os.Exit(1)
	}
	return
}
