// Based on https://github.com/hashicorp/terraform-plugin-sdk/blob/3819ed23c0/helper/logging/transport.go
package logging

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httputil"
	"strings"

	log "github.com/sirupsen/logrus"
)

type transport struct {
	name      string
	logger    *log.Logger
	transport http.RoundTripper
}

func (t *transport) RoundTrip(req *http.Request) (*http.Response, error) {
	if IsDebugOrHigher(t.logger) {
		reqData, err := httputil.DumpRequestOut(req, true)
		if err == nil {
			t.logger.Debugf(logReqMsg, t.name, prettyPrintJsonLines(reqData))
		} else {
			t.logger.Errorf("%s API Request error: %#v", t.name, err)
		}
	}

	resp, err := t.transport.RoundTrip(req)
	if err != nil {
		return resp, err
	}

	if IsDebugOrHigher(t.logger) {
		respData, err := httputil.DumpResponse(resp, true)
		if err == nil {
			t.logger.Debugf(logRespMsg, t.name, prettyPrintJsonLines(respData))
		} else {
			t.logger.Errorf("%s API Response error: %#v", t.name, err)
		}
	}

	return resp, nil
}

func NewTransport(name string, logger *log.Logger, t http.RoundTripper) *transport {
	return &transport{name, logger, t}
}

// prettyPrintJsonLines iterates through a []byte line-by-line,
// transforming any lines that are complete json into pretty-printed json.
func prettyPrintJsonLines(b []byte) string {
	parts := strings.Split(string(b), "\n")
	for i, p := range parts {
		if b := []byte(p); json.Valid(b) {
			var out bytes.Buffer
			json.Indent(&out, b, "", " ")
			parts[i] = out.String()
		}
	}
	return strings.Join(parts, "\n")
}

const logReqMsg = `%s API Request Details:
---[ REQUEST ]---------------------------------------
%s
-----------------------------------------------------`

const logRespMsg = `%s API Response Details:
---[ RESPONSE ]--------------------------------------
%s
-----------------------------------------------------`
