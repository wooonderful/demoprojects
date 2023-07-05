package main

import (
	"bytes"
	"crypto/tls"
	"net/http"
	"strings"
)

var registryUrl = "https://121.43.118.69:30003"

func main() {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	cli := &http.Client{Transport: tr}

	if !strings.HasPrefix(registryUrl, "http") {
		registryUrl = "https://" + strings.TrimSpace(registryUrl)
	}

	registryUrl = strings.TrimSuffix(registryUrl, "/") + "/v2"

	{ // harbor 场景
		catalogUrl := registryUrl + "/_catalog"
		harborReq, _ := http.NewRequest(http.MethodGet, catalogUrl, bytes.NewReader([]byte{}))

		harborReq.SetBasicAuth("admin", "Default1!")

		var harborRes *http.Response
		harborRes, err := cli.Do(harborReq)
		if err != nil {
			return
		}
		defer harborRes.Body.Close()
		if harborRes.StatusCode >= 200 && harborRes.StatusCode < 300 {
			return
		}
	}
}
