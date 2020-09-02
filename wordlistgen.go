package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"sync"
	"time"
)

func main() {
	var concurrency int
	var parameters string
	var domains string
	var domain string
	var appendF bool
	flag.IntVar(&concurrency, "c", 30, "The concurrency for speed")
	flag.StringVar(&domains, "dL", "", "The domains file")
	flag.StringVar(&domain, "d", "", "The domain")
	flag.StringVar(&parameters, "p", "", "The parameters to gen wordlist")
	flag.BoolVar(&appendF, "a", true, "Append / at the end of the endpoint")
	flag.Parse()

	if (parameters != "" && domains != "") || (parameters != "" && domain != "") {
		var wg sync.WaitGroup
		for i := 0; i <= concurrency; i++ {
			wg.Add(1)
			go func() {
				gen(domains, domain,parameters, appendF)
				wg.Done()
			}()
			wg.Wait()
		}
	}
}

func gen(ds string, d string, p string, a bool) {

	// Regex to find relative links
	REGEX := `(https?|ftp|file)://[-A-Za-z0-9\+&@#/%?=~_|!:,.;]*[-A-Za-z0-9\+&@#/%=~_|]`

	time.Sleep(time.Millisecond * 10)

	var parameters = make([]string, 0)
	var links = make([]string, 0)
	var domains = make([]string, 0)

	client:=&http.Client{}

	jScanner := bufio.NewScanner(os.Stdin)
	for jScanner.Scan() {
		req, err := http.NewRequest("GET",jScanner.Text(), nil)
		if err != nil { return }
		resp,err := client.Do(req)
		if err != nil { return }
		bodyBytes,err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return
		}

		bodyString := string(bodyBytes)
		re := regexp.MustCompile(REGEX)
		match := re.FindStringSubmatch(bodyString)
		if match != nil {
			links = append(links, string(match[0]))
		}
	}

	fp, err := os.Open(p)
	if err != nil {
		return
	}

	if d == "" {

		dp, err := os.Open(ds)
		if err != nil {
			return
		}
	
		dScanner := bufio.NewScanner(dp)
		for dScanner.Scan() {
			domains=append(domains, dScanner.Text())
		}
	
		pScanner := bufio.NewScanner(fp)
		for pScanner.Scan() {
			parameters = append(parameters, pScanner.Text())
		}

		for _,domain := range domains {
			for  _,link := range links {
				for _,parameter := range parameters {
					u, err := url.Parse(link)
					if err != nil {
						return
					}
					if a == false {
						fmt.Printf("%s%s?%s=FUZZ\n", domain, u.Path, parameter)
					}else {
						fmt.Printf("%s%s/?%s=FUZZ\n", domain, u.Path, parameter)
					}
				}
			}
		}
	}else {

		pScanner := bufio.NewScanner(fp)
                for pScanner.Scan() {
                        parameters = append(parameters, pScanner.Text())
                }
                for  _,link := range links {
            		for _,parameter := range parameters {
                      		u, err := url.Parse(link)
                                if err != nil {
                              		 return
                                }
				if a == false {
                                	fmt.Printf("%s%s?%s=FUZZ\n", d, u.Path, parameter)
                                }else {
                             		 fmt.Printf("%s%s/?%s=FUZZ\n", d, u.Path, parameter)
                               	}
                        }
                }

	}
}
