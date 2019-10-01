package main

import (
	"fmt"
	"log"
	"os/exec"

	nmap "github.com/Shouren/go-nmap"
)

type HostStat struct {
	HostAddr string
	Latency  string
	HostUp   bool
}

type Result struct {
	HostResults map[string]HostStat
}

func main() {
	var cmd []string
	cmd = []string{
		"-oX", "-", "-v",
		"-sn",
		"--host-timeout", "1s",
	}
	hosts := []string{"103.37.149.30", "188.166.211.141", "172.104.102.93", "198.35.44.222"}

	cmd = append(cmd, hosts...)

	out, err := exec.Command("nmap", cmd...).Output()
	if err != nil {
		log.Fatal(err)
	}
	nmaprun, err := nmap.Parse(out)
	fmt.Printf("%s", out)
	res := Result{
		HostResults: make(map[string]HostStat),
	}
	for _, host := range nmaprun.Hosts {
		var hostup bool
		if host.Status.State == "up" {
			hostup = true
		}
		res.HostResults[host.Addresses[0].Addr] = HostStat{
			HostAddr: host.Addresses[0].Addr,
			Latency:  host.Times.SRTT,
			HostUp:   hostup,
		}
	}
	for _, ret := range res.HostResults {
		fmt.Println(ret.HostAddr, ret.Latency, ret.HostUp)
	}
}
