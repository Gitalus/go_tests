/* You have been asked to make a function called WebsiteRacer which takes two URLs and "races" them by hitting them with an HTTP GET and returning the URL which returned first. If none of them return within 10 seconds then it should return an error. */
package races

import (
	"fmt"
	"net/http"
	"time"
)

func Racer(aURL, bURL string) (winner string, error error) {
	// select espera multiples channels, el primero en enviar un valor gana
	select {
	case <-ping(aURL):
		return aURL, nil
	case <-ping(bURL):
		return bURL, nil
	case <-time.After(10 * time.Second):
		// Errorf returns a created error
		return "", fmt.Errorf("timed out waiting fot %s and %s", aURL, bURL)
	}
}

func ping(url string) chan struct{} { // returns a empty struct because is the smallest type
	// Always use make for channel to avoid zero value for channles, that is nil
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}

// func measureResponseTime(url string) time.Duration {
// 	start := time.Now()
// 	http.Get(url)
// 	return time.Since(start)
// }
