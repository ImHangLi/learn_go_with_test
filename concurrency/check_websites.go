/**
We have covered: 
- goroutines, the basic unit of concurrency in Go, which let us manage more than one website check request.
- anonymous functions, which we used to start each of the concurrent processes that check websites.
- channels, to help organize and control the communication between the different processes, allowing us to avoid a race condition bug.
- the race detector which helped us debug problems with concurrent code.
*/

package concurrency

type WebsiteCheck func(string) bool

type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteCheck, urls []string) map[string]bool {
	results := map[string]bool{}
	resultChannel := make(chan result)

	for _, url := range urls {
		go func() {
			resultChannel <- result{url, wc(url)}
		} ()
	}

	for i := 0; i < len(urls); i++ {
		r := <- resultChannel
		results[r.string] = r.bool
	}

	return results
}