package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const (
	MAX_REDIRECTS int    = 10
	DOIURL        string = "http://dx.doi.org/"
)

// GetDOI takes DOI spec and returns its metadata
func GetDOI(doi string) (string, error) {
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			if len(via) >= MAX_REDIRECTS {
				msg := fmt.Sprintf("Stopped after %d redirects", MAX_REDIRECTS)
				return errors.New(msg)
			}
			// Move the 'Accept' header ahead
			req.Header.Set("Accept", via[0].Header.Get("Accept"))
			return nil
		}}

	req, err := http.NewRequest("GET", DOIURL+doi, nil)
	if err != nil {
		return "", err
	}

	// req.Header.Add(HDR_ACCEPT, HDR_JSON2)
	// req.Header.Add("Accept", "application/citeproc+json")
	// req.Header.Add("Accept", "application/vnd.citationstyles.csl+json")
	// req.Header.Add("Accept", "application/x-bibtex")
	// req.Header.Add("Accept", "application/x-datacite+xml")
	// req.Header.Add("Accept", "application/rdf+xml")
	// req.Header.Add("Accept", "application/vnd.datacite.datacite+xml")
	// req.Header.Add("Accept", "application/vnd.datacite.datacite+text")
	// req.Header.Add("Accept", "application/x-datacite+text")
	// req.Header.Add("Accept", "application/x-research-info-systems")
	// req.Header.Add("Accept", "text/x-bibliography")
	// req.Header.Add("Accept", "text/turtle")
	// req.Header.Add("Accept", "text/html")

	req.Header.Set("Accept", "application/citeproc+json")
	req.Header.Set("Accept", "application/vnd.datacite.datacite+xml")
	req.Header.Set("Accept", "application/vnd.citationstyles.csl+json")
	//req.Header.Add("Accept",
	//strings.Join([]string{
	// "application/citeproc+json",
	// "application/vnd.citationstyles.csl+json",
	// "application/x-bibtex",
	// "application/x-datacite+xml",
	// "application/rdf+xml",
	// "application/vnd.datacite.datacite+xml",
	//"application/vnd.datacite.datacite+text",
	//"application/x-datacite+text",
	//"application/x-research-info-systems",
	//"text/x-bibliography",
	//"text/turtle",
	//"text/html"
	// 	},
	// 		", "))

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	// fmt.Println("Responce content-type:", resp.Header.Get("contentType"))
	// fmt.Println("Request headers:")
	// for k, v := range req.Header {
	// 	fmt.Println("-->", k, ": ", v)
	// }
	// fmt.Printf("Responce status code: %d\n", resp.StatusCode)
	for k, v := range resp.Header {
		fmt.Println("<--", k, ": ", v)
	}
	// fmt.Println("Final URL:", resp.Request.URL.String())

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body[:]), nil
}

func main() {
	arglen := len(os.Args)
	for _, doi := range os.Args[1:] {
		fmt.Println("Processing " + DOIURL + doi)

		body, err := GetDOI(doi)
		if err != nil {
			log.Println("Error: ")
			log.Println(err)
		}
		fmt.Println(body)
		if arglen > 2 {
			fmt.Println("-----------------------------------------------")
		}

	}
}
