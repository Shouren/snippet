package main

import (
	"html/template"
	"log"
	"os"
)

var (
	TEST string = "Hello World"
	NAME string = "Hhello World"
)

type Upstream struct {
	IP     string `json:"ip-addr"`
	Port   string `json:"port"`
	Weight uint   `json:"weight"`
}

type APISpec struct {
	Name               string
	StripeUrl          bool
	ListenPath         string `json:"listen-path"`
	UriParamsUpstreams map[string]map[string][]Upstream
	HeadersUpstreams   map[string]map[string][]Upstream
	DefaultUpstreams   []Upstream
}

var tempData = `
APISpec: {{.Name}}

URI:
{{range $key, $elements := .UriParamsUpstreams}}
  key: {{$key}}
  {{range $value, $upstreams := $elements}}
    value: {{$value}}
      {{range $ups := $upstreams}}
        IPAddr: {{$ups.IP}}
      {{end}}
  {{end}}
{{end}}
HEADERS:
{{range $key, $elements := .HeadersUpstreams}}
  key: {{$key}}
  {{range $value, $upstreams := $elements}}
    value: {{$value}}
      {{range $ups := $upstreams}}
        IPAddr: {{$ups.IP}}
      {{end}}
  {{end}}
{{end}}
Default:
{{range $ups := .DefaultUpstreams}}
  IPAddr: {{$ups.IP}}
{{end}}
`

func main() {
	test := APISpec{
		Name:       "test",
		StripeUrl:  false,
		ListenPath: "/apis",
		UriParamsUpstreams: map[string]map[string][]Upstream{
			"target": map[string][]Upstream{
				"test": []Upstream{
					Upstream{
						IP:     "127.0.0.2",
						Port:   "80",
						Weight: 1,
					},
				},
			},
		},
		HeadersUpstreams: map[string]map[string][]Upstream{},
		DefaultUpstreams: []Upstream{
			Upstream{
				IP:     "127.0.0.1",
				Port:   "80",
				Weight: 1,
			},
			Upstream{
				IP:     "127.0.0.1",
				Port:   "8080",
				Weight: 1,
			},
		},
	}
	tp := template.New("test")
	tp.Parse(tempData)

	f, err := os.OpenFile("test.txt", os.O_RDWR|os.O_CREATE, 0755)
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}
	err = tp.Execute(f, test)
	if err != nil {
		log.Println(err)
	}
}
