/* You have been asked to make a function called WebsiteRacer which takes two URLs and "races" them by hitting them with an HTTP GET and returning the URL which returned first. If none of them return within 10 seconds then it should return an error. */
package races

import (
	"net/http"
	"time"
)

func Racer(firstURL, secondURL string) (winner string) {
	aDuration := measureResponseTime(firstURL)
	bDuration := measureResponseTime(secondURL)

	if aDuration < bDuration {
		return firstURL
	}

	return secondURL
}

func measureResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}
