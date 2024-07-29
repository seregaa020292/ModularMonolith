package reqbody

import (
	"bytes"
	"io"
	"net/http"
)

func CopyBody(r *http.Request) []byte {
	if r.Body == nil {
		return nil
	}

	defer r.Body.Close()

	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		return nil
	}

	r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	return bodyBytes
}
