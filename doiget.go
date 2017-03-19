package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const (
	maxRedirects int = 10

	// APIURL is the URL of the DOI
	APIURL string = "http://dx.doi.org/"
)

// FetchMeta takes DOI spec and returns its metadata
func FetchMeta(doi string) ([]byte, error) {
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			if len(via) >= maxRedirects {
				msg := fmt.Sprintf("Stopped after %d redirects", maxRedirects)
				return errors.New(msg)
			}
			// Move the 'Accept' header ahead
			req.Header.Set("Accept", via[0].Header.Get("Accept"))
			return nil
		}}

	req, err := http.NewRequest("GET", APIURL+doi, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/citeproc+json")
	// req.Header.Set("Accept", "application/vnd.datacite.datacite+xml")
	// req.Header.Set("Accept", "application/vnd.citationstyles.csl+json")
	// "application/citeproc+json",
	// "application/vnd.citationstyles.csl+json",
	// "application/x-bibtex",
	// "application/x-datacite+xml",
	// "application/rdf+xml",
	// "application/vnd.datacite.datacite+xml",
	// "application/vnd.datacite.datacite+text",
	// "application/x-datacite+text",
	// "application/x-research-info-systems",
	// "text/x-bibliography",
	// "text/turtle",
	// "text/html"

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
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
		return nil, err
	}
	return body, nil
}

func main() {
	arglen := len(os.Args)
	for _, doi := range os.Args[1:] {
		fmt.Println("Processing " + APIURL + doi)

		body, err := FetchMeta(doi)
		if err != nil {
			log.Println("Error: ")
			log.Println(err)
		}

		fmt.Println()

		var publication Publication
		err = json.Unmarshal(body, &publication)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%#v\n", publication)
		for _, a := range publication.Author {
			fmt.Printf("%#v\n", a)
		}
		fmt.Printf("%#v\n", publication.Created)
		fmt.Printf("%#v\n", publication.Deposited)
		fmt.Printf("%#v\n", publication.Indexed)

		// var ppi interface{}
		// err = json.Unmarshal(body, &ppi)
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// fmt.Println(ppi)
		// fmt.Printf("%#v\n\n", ppi)

		fmt.Println()
		fmt.Println(string(body[:]))

		if arglen > 2 {
			fmt.Println("-----------------------------------------------")
		}

	}
}
