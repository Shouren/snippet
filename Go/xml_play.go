package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type XMLnmaprun struct {
	XMLName    xml.Name     `xml:"nmaprun"`
	Scanner    string       `xml:"scanner,attr"`
	Args       string       `xml:"args,attr"`
	Start      string       `xml:"start,attr"`
	Startstr   string       `xml:"startstr,attr"`
	Version    string       `xml:"version,attr"`
	xmlversion string       `xml:"xmloutputversion,attr"`
	Verbose    XMLverbose   `xml:"verbose"`
	Debugging  XMLdebugging `xml:"debugging"`
	Hosts      []XMLhost    `xml:"host"`
	Runstats   XMLrunstats  `xml:"runstats"`
}

type XMLverbose struct {
	XMLName xml.Name `xml:"verbose"`
	Level   string   `xml:"level,attr"`
}

type XMLdebugging struct {
	XMLName xml.Name `xml:"debugging"`
	level   string   `xml:"level,attr"`
}

type XMLhost struct {
	XMLName   xml.Name     `xml:"host"`
	status    XMLstatus    `xml:"status"`
	address   XMLadress    `xml:"address"`
	hostnames XMLhostnames `xml:"hostname"`
	times     XMLtimes     `xml:"times"`
}

type XMLstatus struct {
	XMLName    xml.Name `xml:"status"`
	state      string   `xml:"state,attr"`
	reason     string   `xml:"reason,attr"`
	reason_ttl string   `xml:"reason_tll"`
}

type XMLadress struct {
	XMLName  xml.Name `xml:"address"`
	addr     string   `xml:"addr"`
	addrtype string   `xml"addrtype"`
}

type XMLhostnames struct {
	XMLName xml.Name `xml:"hostnames"`
}

type XMLtimes struct {
	XMLName xml.Name `xml:"times"`
	srtt    string   `xml:"srtt"`
	rttvar  string   `xml:"rttvar"`
	to      string   `xml:"to"`
}

type XMLrunstats struct {
	XMLName xml.Name `xml:"runstats"`
	time
}

func main() {
	f, err := os.Open("res.xml")
	if err != nil {
		fmt.Println("Error open file res.xml: %s", err)
	}
	defer f.Close()

	content, _ := ioutil.ReadAll(f)
	var task XMLnmaprun
	paerr := xml.Unmarshal([]byte(content), &task)
	if err != nil {
		fmt.Printf("Error: %v", paerr)
		return
	}

	fmt.Printf("Name: %v\n", task.XMLName)
	fmt.Printf("scanner: %v\n", task.Scanner)
	fmt.Printf("verbose level: %v\n", task.Verbose.Level)
}
