package main

import "fmt"

// DateStruct is used in Publication
type DateStruct struct {
	DateParts [][]int `json:"date-parts"`
	DateTime  string  `json:"date-time"`
	Timestamp int     `json:"timestamp"`
}

// Author
type Author struct {
	Affiliation []string `json:"affiliation"`
	Family      string   `json:"family"`
	Given       string   `json:"given"`
}

func (a *Author) String() string {
	return fmt.Sprintf("%s %s from %s", a.Given, a.Family, a.Affiliation)
}

// Paper metainformation struct contains most fields that come with responce JSON
type Publication struct {
	DOI            string      `json:"DOI"`
	ISSN           []string    `json:"ISSN"`
	URL            string      `json:"URL"`
	Author         []*Author   `json:"author"`
	ContainerTitle string      `json:"container-title"`
	Created        *DateStruct `json:"created"`
	Deposited      *DateStruct `json:"deposited"`
	Indexed        *DateStruct `json:"indexed"`
	Issue          string      `json:"issue"`
	Issued         *struct {
		DateParts [][]int `json:"date-parts"`
	} `json:"issued"`
	Member         string   `json:"member"`
	Page           string   `json:"page"`
	Prefix         string   `json:"prefix"`
	Publisher      string   `json:"publisher"`
	ReferenceCount int      `json:"reference-count"`
	Score          float64  `json:"score"`
	Source         string   `json:"source"`
	Subject        []string `json:"subject"`
	Subtitle       []string `json:"subtitle"`
	Title          string   `json:"title"`
	Type           string   `json:"type"`
	Volume         string   `json:"volume"`
}
