package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"sort"
)

func main() {

	var data services

	in, err := ioutil.ReadFile("services.yaml")
	if err != nil {
		log.Fatal(err)
	}

	if err := yaml.Unmarshal(in, &data); err != nil {
		log.Fatal(err)
	}

	sort.Sort(data)

	out, err := yaml.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}

	if err := ioutil.WriteFile("services.yaml", out, 0644); err != nil {
		log.Fatal(err)
	}

}

type services struct {
	Services []service `yaml:"services"`
}

func (s services) Len() int {
	return len(s.Services)
}

func (s services) Less(a, b int) bool {
	return s.Services[a].Name < s.Services[b].Name
}

func (s services) Swap(a, b int) {
	s.Services[a], s.Services[b] = s.Services[b], s.Services[a]
}

type service struct {
	Name                 string `yaml:"name"`
	Version              string `yaml:"version,omitempty"`
	Count                int    `yaml:"count,omitempty"`
	URI                  string `yaml:"uri,omitempty"`
	DesiredState         string `yaml:"desiredState,omitempty"`
	SequentialDeployment bool   `yaml:"sequentialDeployment,omitempty"`
}
