package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"
	"time"
)

type Results struct {
	Results map[string]bool
	Source string
}

const regexStr  = " ([$A-Z\\_a-z0-9]+\\s*=[^=])"

func main()  {
	scanTime := strings.Replace(time.Now().Format("15:04:05"),":","-",-1)
	saveResultsFlag := flag.Bool("s", false, "saves output to file named from base url of source")
	saveResultsLongFlag := flag.Bool("save", false, "saves output to file named from base url of source")
	paramLengthFilterFlag := flag.Int("length",0, "Minimum length variable to collect")
	flag.Parse()
	r, _ := regexp.Compile(regexStr)
	scan := bufio.NewScanner(os.Stdin)
	for scan.Scan() {
		extractParams(r, scan.Text(),*saveResultsFlag||*saveResultsLongFlag, scanTime, *paramLengthFilterFlag)
	}



}

func extractParams(r *regexp.Regexp, source string, saveOutput bool, scanTime string, lengthFilter int) {
	resultsMap := &Results{ make(map[string]bool),source}
	var client http.Client
	resp, err := client.Get(source)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		scanner := bufio.NewScanner(resp.Body)
		for scanner.Scan() {
			line := scanner.Text()
			if strings.Contains(line, "var ") {
				results := r.FindAllString(line, -1)
				if len(results) != 0 {
					for _, variable := range r.FindAllString(line, -1) {
						cleaned := strings.TrimSpace(strings.Split(variable, "=")[0])
						if len(cleaned) > lengthFilter {
							if !resultsMap.Results[cleaned] {
								resultsMap.Results[cleaned] = true
							}
						}
					}
				}
			}
		}
		parsedSource, _ := url.Parse(source)
		count := 0
		if !saveOutput {
			for k := range resultsMap.Results {
				if count == len(resultsMap.Results)-1 {
					fmt.Print(k)
				} else {
					fmt.Println(k)
				}
				count+=1
			}
		} else {
			f, err := os.OpenFile(parsedSource.Host+"_"+scanTime,
				os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				log.Println(err)
			}
			defer f.Close()
			for k := range resultsMap.Results {
				if count == len(resultsMap.Results)-1 {
					if _, err := f.WriteString(k); err != nil {
						log.Fatalln(err)
					}
				} else {
					if _, err := f.WriteString(k+"\n"); err != nil {
						log.Fatalln(err)
					}
				}
				count+=1
			}
		}
	}
}