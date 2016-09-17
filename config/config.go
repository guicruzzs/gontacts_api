package config

import (
  "encoding/json"
  "os"
  "fmt"
)

type Configuration struct {
  Db DatabaseConfig
}

type DatabaseConfig struct {
  Host string
  Port string
  User string
  Password string
  Schema string
}

func Get() Configuration {
  var env string = os.Getenv("GO_ENV")
  if (env == ""){ env = "development" }

  file, _ := os.Open(fmt.Sprintf("config/%s.json", env))
  decoder := json.NewDecoder(file)
  configuration := Configuration{}
  err := decoder.Decode(&configuration)
  if err != nil {
    fmt.Println("error:", err)
  }
  return configuration
}
