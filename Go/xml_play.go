package main

import (
	// "encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type nmaptask struct {
	XMLName  xml.Name `xml:"nmaptask"`
	scanner  string   `xml:"scanner,attr"`
	args     string   `xml:"args,attr"`
	start    string   `xml:"start,attr"`
	startstr string   `xml:"startstr,attr"`
	version  string   `xml:"version,attr"`
}

type verbose struct {
	XMLName xml.Name `xml:"verbose"`
	level   string   `xml:"level,attr"`
}

type debugging struct {
	XMLName xml.Name `xml:"debugging"`
	level   string   `xml:"level,attr"`
}

func main() {
	f, err := os.Open("res.xml")
	if err != nil {
		fmt.Println("Error open file res.xml: %s", err)
	}
	defer f.Close()

	content, _ := ioutil.ReadAll(f)
	fmt.Printf("%s", content)
}
