package configs

import ( 
	"fmt"
	"io/ioutil"
	"gopkg.in/yaml.v3"
	"log"
)

func GetCurrentEpisode() {
	fmt.Println("GetCurrentEpisode called")
	
	// https://zetcode.com/golang/yaml/
	yfile, err := ioutil.ReadFile("configs.yaml")
	
	if err != nil {
          log.Fatal(err)
  	}


	data := make(map[interface{}]interface{})

	yaml.Unmarshal(yfile, &data)

	for k, v := range data {
		fmt.Printf("%s -> %d\n", k, v)
	}


}	
