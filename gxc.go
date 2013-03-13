package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"os"
)

type Param struct {
	Name  string `xml:"name,attr"`
	Value string `xml:"value,attr"`
}

type Section struct {
	Name        string `xml:"name,attr"`
	Description string `xml:"description,attr"`
	Params      []Param
}

type Conf struct {
	XMLName xml.Name `xml:"document"`
	Section []Section
}

func main() {
	conf := Conf{}

	fin, err := os.Open("/home/rif/Documents/prog/go/src/github.com/rif/gxc/cgrates.xml")
	if err != nil {
		log.Fatal(err)
	}

	data := make([]byte, 100)
	_, err = fin.Read(data)
	if err != nil {
		log.Fatal(err)
	}

	err = xml.Unmarshal([]byte(data), &conf)
	if err != nil {
		fmt.Printf("xml error: %v\n", err)
		return
	}

}
