package schemas

type ScanningRequest struct {
	Index      string                 `json:"index"`
	DocumentID string                 `json:"documentId"`
	Body       map[string]interface{} `json:"body"`
}
