package pp

import (
    "sync"
    "fmt"
    "io/ioutil"
    "encoding/json"
    "os"
)

type DatabaseConnection struct {
    Host string `json:"host"`
    Port int `json:"port"`
    Username string `json:"username"`
    Password string `json:"password"`
}

type Configuration struct {
	Port int `json:"port"`
	Trackers map[string]Tracker `json:"trackers"`
    Database DatabaseConnection `json:"database"`
}

var instance *Configuration
var once sync.Once

func ReadConfiguration() *Configuration {
	file, e := ioutil.ReadFile("./config.json")
    if e != nil {
        fmt.Printf("File error: %v\n", e)
        os.Exit(1)
    }

    //m := new(Dispatch)
    //var m interface{}
    instance = &Configuration{}
    json.Unmarshal(file, &instance)
    return instance
}

func GetConfiguration() *Configuration {
    once.Do(func() {
        ReadConfiguration()
    })
    return instance
}