package googlesearch

import (
	"net/http"
	"io/ioutil"
	"log"
	"regexp"
	"net/url"
)

type Result struct {
	title, url string
}

// Search for query string... returns the max... of 8 results
func Search(query string) (results []Result) {

	query = url.QueryEscape(query)

	//	uri := "http://ajax.googleapis.com/ajax/services/search/web?v=1.0&rsz=8&q=" + query
	uri := "http://www.google.com/uds/GwebSearch?start=0&rsz=large&hl=en&key=notsupplied&v=1.0&q=" + query

	data, e := http.Get(uri)
	handle(e)
	defer data.Body.Close()

	body, e := ioutil.ReadAll(data.Body)
	handle(e)
	jsonRes := string(body)

	// that json package just screwed with my zen for too long
	// regexp it went.
	re := regexp.MustCompile("\"url\":\"(.+?)\"")
	links := re.FindAllStringSubmatch(jsonRes, -1)
	re = regexp.MustCompile("\"title\":\"(.+?)\"")
	titles := re.FindAllStringSubmatch(jsonRes, -1)

	for i, _ := range titles {
		res := Result{title:titles[i][1], url:links[i][1]}
		results = append(results, res)
	}

	return
}

func handle(err error) {
	if err != nil {
		log.Println("ERROR")
		log.Fatal(err)
	}
}
