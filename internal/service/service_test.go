package service

import (
	"fmt"
	"net/url"
	"testing"
)

func TestOpenGraphClient(tst *testing.T) {
	// ls := NewLinkService(nil, opengraph.NewOGClient())
	urls := []string{
		"http://time.com/5299000/donald-trump-june-12-meeting-kim-jong-un-back-on/",
		"https://www.cbsnews.com/news/alexander-gauland-calls-nazi-era-a-speck-of-bird-poop-in-german-history/",
		"https://www.sbs.com.au/news/france-gives-us-a-few-days-to-avoid-trade-war",
		"https://abcnews.go.com/Health/wireStory/health-agency-backs-call-tax-sugary-drinks-55581825",
	}

	//t := time.Now()
	// var wg sync.WaitGroup
	// for _, v := range urls {
	// 	wg.Add(1)
	// 	go func(url string) {
	// 		defer wg.Done()
	// 		_, err := ls.GetLinkFromURL(context.TODO(), url)
	// 		if err != nil {
	// 			tst.Error(err)
	// 			return
	// 		}
	// 	}(v)
	// }
	// wg.Wait()
	// fmt.Println(time.Since(t))

	for _, v := range urls {
		e := url.QueryEscape(v)
		fmt.Println(e)
	}

}
