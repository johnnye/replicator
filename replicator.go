package replicator

import (
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Replicator struct {
	newURL     string
	meh        bool
	percentage int
}

func NewReplicator(newURL string, meh bool, pcnt int) *Replicator {
	// Set some default parameters
	rand.Seed(time.Now().UTC().UnixNano())

	return &Replicator{newURL, meh, pcnt}
}

func (r *Replicator) ServeHTTP(w http.ResponseWriter, req *http.Request, next http.HandlerFunc) {

	random := rand.Intn(r.percentage)

	//If this request is not below our percentage move along
	if random >= r.percentage {
		next(w, req)
	}

	newUrl := strings.Join([]string{r.newURL, req.URL.String()}, "")
	req.URL, _ = url.Parse(newUrl)
	req.RequestURI = ""

	if r.meh == true {
		passiveMode(req, r)
	} else {
		locking(req, r)
	}

	next(w, req)
}

func passiveMode(req *http.Request, r *Replicator) {
	c := http.Client{}
 	go c.Do(req)
}

func locking(req *http.Request, r *Replicator) {
	done := make(chan bool)
	client := http.Client{}

	go func() {
		client.Do(req)

		done <- true
	}()
	<-done
}
