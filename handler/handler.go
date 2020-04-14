package handler

import (
	"bytes"
	"log"
	"net/http"
	"sync"
)

//Handler handler
type Handler struct {
	wg    *sync.WaitGroup
	pool  int
	limit int
	ctr   int
	url   string
	m     chan int
}

// NewHandler handler constructor
func NewHandler(url string, pool int, limit int) *Handler {
	return &Handler{
		url:   url,
		pool:  pool,
		limit: limit,
		wg:    &sync.WaitGroup{},
		m:     make(chan int, 1),
	}
}

// Handle handling function
func (h *Handler) Handle() {
	for i := 0; i < h.limit/h.pool; i += 1 {
		i := int64(i)
		h.wg.Add(h.pool / 100)
		for j := 0; j < h.pool/100; j += 1 {
			j := int64(j)
			go func() {
				for k := 1; k <= 100; k++ {
					k := int64(k)
					var contentId int64 = i*int64(h.pool) + j*100 + k
					log.Print(contentId)
					req, err := http.NewRequest("POST", h.url, bytes.NewBuffer(EntryGen(contentId).Marshal()))
					if err != nil {
						continue
					}
					req.Header.Set("Content-Type", "application/json")
					//req, _ := http.NewRequest("GET", h.url, nil)
					client := &http.Client{}
					res, err := client.Do(req)
					if err != nil {
						print("err: ")
						println(client)
						println(err.Error())
					} else {
						if res.StatusCode != http.StatusOK {
							print("err: ")
							println(client)
							println(res.Status)
						}
						res.Body.Close()
					}
				}
				h.wg.Done()
			}()
		}
		h.wg.Wait()
	}
}
