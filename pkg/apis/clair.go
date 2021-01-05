package apis

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/elastic/go-elasticsearch/esapi"
	elasticsearch "github.com/elastic/go-elasticsearch/v7"
	"github.com/tmax-cloud/image-scanning-webhook/pkg/schemas"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
)

var logWebhook = logf.Log.WithName("webhook")

func CreateClairLog(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var m schemas.ScanningRequest

	// get body
	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		logWebhook.Error(err, "error occurs while decoding service instance body")
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
		logWebhook.Error(err, "Error creating the client: %s")
	}

	//change body data into byte
	d, err := json.Marshal(m.Body)
	if err != nil {
		logWebhook.Error(err, "change body data")
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
		logWebhook.Error(err, "Error getting response")
	}

	//response setting
	w.WriteHeader(res.StatusCode)
	w.Header().Set("Content-Type", res.Header.Get("Content-Type"))
	w.Header().Set("Content-Length", res.Header.Get("Content-Length"))
	io.Copy(w, res.Body)
	res.Body.Close()
}
