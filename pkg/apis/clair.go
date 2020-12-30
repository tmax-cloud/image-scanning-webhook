package apis

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/elastic/go-elasticsearch/esapi"
	elasticsearch "github.com/elastic/go-elasticsearch/v7"
	"github.com/jitaeyun/image-scanning-webhook/pkg/schemas"
)

func CreateClairLog(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var m schemas.ScanningRequest

	// get body
	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		log.Fatal(err, "error occurs while decoding service instance body")
		return
	}

	//setting elasticsearch config
	elasticsearchUrl := os.Getenv("ELASTIC_SEARCH_URL")
	cfg := elasticsearch.Config{
		Addresses: []string{
			elasticsearchUrl,
		},
	}

	//set elasticsearch client
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	//change body data into byte
	d, err := json.Marshal(m.Body)
	if err != nil {
		log.Fatalf("change body data: %s", err)
	}

	//change data into io.reader
	data := bytes.NewReader(d)
	req := esapi.IndexRequest{
		Index:      m.Index,
		DocumentID: m.DocumentID,
		Body:       data,
		Refresh:    "true",
	}

	// Perform the request with the client.
	res, err := req.Do(context.Background(), es)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}

	//response setting
	w.WriteHeader(res.StatusCode)
	w.Header().Set("Content-Type", res.Header.Get("Content-Type"))
	w.Header().Set("Content-Length", res.Header.Get("Content-Length"))
	io.Copy(w, res.Body)
	res.Body.Close()
}
