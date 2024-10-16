package reqbody

import (
	"bytes"
	"io"
	"net/http"
)

func Copy(r *http.Request) []byte {
	if r.Body == nil {
		return nil
	}

	defer r.Body.Close()

	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		return nil
	}

	r.Body = io.NopCloser(bytes.NewReader(bodyBytes))

	return bodyBytes
}
