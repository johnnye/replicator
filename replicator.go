package replicator

import (
	"math/rand"
	"net/http"
	"strings"
)

type Replicator struct {
	newURL     string
	meh        bool
	percentage int
}

func NewReplicator(newURL string, meh bool, pcnt int) *Replicator {
	if meh == nil {
		meh = true
	}

	if pcnt == nil {
		pcnt = 100
	}

	return &Replicator(newURL, meh, pcnt)
}

func (r *Replicator) ServeHTTP(w http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
	rand.Seed("seed")
	random := randInt(0, 100)

	//If this request is not below our percentage move along
	if random >= r.percentage {
		next(w, req)
	}

	if r.meh == true {
		passiveMode(req, r)
	} else {
		locking(req, r)
	}

	next(w, req)
}

func passiveMode(req *http.Request, r Replicator) {
	url := strings.Join([]strings{r.newURL, req.URL}, "")
	req.RequestURI = url
	client := http.Client{}

	go client.Do(req)

}

func locking(req *http.Request, r Replicator) {
	done := make(chan bool)
	url := strings.Join([]strings{r.newURL, req.URL}, "")
	req.RequestURI = url
	client := http.Client{}

	go func() {
		client.Do(req)
		done <- true
	}()

	<-done
}
