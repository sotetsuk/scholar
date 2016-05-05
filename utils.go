package main

import (
	"strings"
	"regexp"
	"github.com/PuerkitoBio/goquery"
)

func parseAndInitializeArguments(arguments map[string]interface{}) (query, author, title, cluster_id, after, before, start, num string) {
	/*
		default: num="10", others=""
	*/

	if arguments["--author"] != nil {
		author = arguments["--author"].(string)
	}
	if arguments["--title"] != nil {
		title = arguments["--title"].(string)
	}
	if arguments["--query"] != nil {
		query = arguments["--query"].(string)
	}
	if arguments["<cluster-id>"] != nil {
		cluster_id = arguments["<cluster-id>"].(string)
	}
	if arguments["--after"] != nil {
		after = arguments["--after"].(string)
	}
	if arguments["--before"] != nil {
		before = arguments["--before"].(string)
	}
	if arguments["--num"] != nil {
		num = arguments["--num"].(string)
	}
	if arguments["--start"] != nil {
		start = arguments["--start"].(string)
	}

	if num == "" {
		num = "10"
	}

	return author, title, query, cluster_id, after, before, num, start
}

func getDoc(query func(map[string]interface{}) (string, error), arguments map[string]interface{}) (*goquery.Document, error) {
	url, err := query(arguments)
	if err != nil {
		return nil, err
	}

	doc, err := goquery.NewDocument(url)
	if err != nil {
		return nil, err
	}

	return doc, nil
}

func parseYear(s string) string {
	re, _ := regexp.Compile("\\d{4}")
	return strings.TrimSpace(string(re.Find([]byte(s))))
}

func parseClusterId(url string) string {
	url = url[15:]
	ix := strings.Index(url, "&")
	url = url[:ix]
	return strings.TrimSpace(url)
}

func parseNumberOfCitations(s string) string {
	return strings.TrimSpace(s[8:])
}

func parseNumberOfVersions(s string) string {
	return strings.TrimSpace(s[strings.Index(s, " "):strings.LastIndex(s, " ")])
}

func parseInfoId(url string) string {
	url = url[19:]
	ix := strings.Index(url, ":")
	url = url[:ix]
	return strings.TrimSpace(url)
}

func parsePDFSource(s string) string {
	// e.g., "[PDF] from arxiv.orgarxiv.org [PDF]"", => "PDFSource": "arxiv.org"
	prefix := "[PDF] from "
	suffix := " [PDF]"
	if strings.HasPrefix(s, prefix) && strings.HasSuffix(s, suffix) {
		s = strings.TrimSpace(s[len(prefix):len(s) - len(suffix)])
		return s[:len(s) / 2]
	}
	return ""
}

func StartAndEndWithDoubleQuotation(s string) bool {
	if strings.HasPrefix(s, "\"") && strings.HasSuffix(s, "\"") {
		return true
	} else {
		return false
	}
}