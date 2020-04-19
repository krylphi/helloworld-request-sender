package handler

import (
	"log"
	"sync"

	"github.com/valyala/fasthttp"
)

//Handler handler
type Handler struct {
	wg    *sync.WaitGroup
	pool  int
	limit int
	ctr   int
	url   string
	m     chan *Entry
}

// NewHandler handler constructor
func NewHandler(url string, pool int, limit int) *Handler {
	return &Handler{
		url:   url,
		pool:  pool,
		limit: limit,
		wg:    &sync.WaitGroup{},
		m:     make(chan *Entry, limit),
	}
}

// Handle handling function
func (h *Handler) Handle() {
	h.wg.Add(h.pool)
	for i := 0; i < h.pool; i++ {
		go func() {
			for e := range h.m {
				log.Print(e.ContentId)
				req := fasthttp.AcquireRequest()
				req.URI().Update(h.url)
				req.Header.Add("Content-Type", "application/json")
				req.Header.SetMethodBytes([]byte("POST"))
				req.AppendBody(e.Marshal())

				res := fasthttp.AcquireResponse()
				err := fasthttp.Do(req, res)
				if err != nil {
					log.Print(err.Error())
				}
				// when finish
				fasthttp.ReleaseRequest(req)
				fasthttp.ReleaseResponse(res)
			}
			h.wg.Done()
		}()
	}
	for i := 1; i <= h.limit; i++ {
		h.m <- EntryGen(i)
	}
	close(h.m)
	h.wg.Wait()
}
