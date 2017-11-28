package main

import (
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	t := "http://localhost:6633/"
	target, err := url.Parse(t)
	if err != nil {
		log.Fatal(err)
	}
	path := target.String()
	fmt.Println(target.Host)

	mp, err := regexp.Compile("(.*)")
	if err != nil {
		log.Fatal(err)
	}
	result_slice := mp.FindAllStringSubmatch(path, -1)
	fmt.Printf("%q\n", result_slice)

	newpath := "/apollo$1"
	dollarMatch := regexp.MustCompile(`\$\d+`)
	replace_slice := dollarMatch.FindAllStringSubmatch(newpath, -1)
	fmt.Printf("%q\n", replace_slice)

	mapped_replace := make(map[string]string)
	for mI, replacementVal := range result_slice[0] {
		indexVal := "$" + strconv.Itoa(mI)
		mapped_replace[indexVal] = replacementVal
	}

	for _, v := range replace_slice {
		fmt.Println(v[0])
		newpath = strings.Replace(newpath, v[0], mapped_replace[v[0]], -1)
	}

	fmt.Println(newpath)
}
