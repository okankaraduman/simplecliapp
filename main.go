package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"regexp"
	"runtime"
	"strings"
	"sync"
)

// wg is used to wait for the program to finish.
var wg sync.WaitGroup

func main() {

	num_of_parallel := flag.Int("parallel", 10, "Specify number of threads")
	flag.Parse()
	wg.Add(*num_of_parallel)
	argsWithProg := os.Args
	num_of_website := 0
	for idx, arg := range argsWithProg {
		if idx == 0 {
			continue
		}
		//Checking valid url
		re, _ := regexp.Compile(`(https?:\/\/(www\.))?[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b([-a-zA-Z0-9()@:%_\+.~#?&=]*)`)
		if re.MatchString(arg) {
			if num_of_website >= *num_of_parallel {
				wg.Add(1)
			}
			go makeRequest(arg)
			num_of_website += 1
		}
	}
	num_of_unused_parallels := *num_of_parallel - num_of_website
	for i := 0; i < num_of_unused_parallels; i++ {
		wg.Done()
	}

	wg.Wait()

}

func makeRequest(url string) {
	runtime.LockOSThread()
	defer wg.Done()

	client := &http.Client{}
	if !strings.Contains(url, "www") {
		url = "http://www." + url
	}
	resp, err := client.Get(url)
	if err != nil {
		fmt.Println("The requested URL cannot be reached at the moment!", url)
	}
	hash := GetMD5HashWithWrite(string(resp.Status))
	fmt.Println(url, hash)

	runtime.UnlockOSThread()
}
