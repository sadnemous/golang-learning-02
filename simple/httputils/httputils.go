package httputils

import (
	"bytes"
	"context"
	"io"
	"net/http"
)

type HttpRouter struct {
	URL         string
	Method      string
	Accept      string
	ContentType string
	Payload     []byte
}

func NewHttpRouter(url, method, accept, contentType string, payload []byte) *HttpRouter {
	return &HttpRouter{
		URL:         url,
		Method:      method,
		Accept:      accept,
		ContentType: contentType,
		Payload:     payload,
	}
}

func (h *HttpRouter) Send(ctx context.Context) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, h.Method, h.URL, bytes.NewBuffer(h.Payload))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", h.Accept)
	req.Header.Set("Content-Type", h.ContentType)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
