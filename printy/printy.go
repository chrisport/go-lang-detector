package printy

import (
	"net/http/httputil"
	"net/http"
	"fmt"
	"encoding/json"
	"os"
	"io/ioutil"
	"io"
)

func JsonReader(r io.Reader, pretty bool) {
	bytes, err := ioutil.ReadAll(r)
	if noError(err) {
		JsonBytes(bytes, pretty)
	}
}

func JsonBytes(bytes []byte, pretty bool) {
	var t map[string]interface{}
	err := json.Unmarshal(bytes, &t)
	if noError(err) {
		AsJson(t, pretty)
	}
}

func AsJson(i interface{}, pretty bool) {
	enc := json.NewEncoder(os.Stdout)
	if pretty {
		enc.SetIndent("", "  ")
	}
	noError(enc.Encode(i))
}

func noError(err error) bool {
	if err != nil {
		fmt.Println("Printy error:", err)
	}
	return err == nil
}

func PrettyResponse(resp *http.Response) {
	b, err := httputil.DumpResponse(resp, true)

	if noError(err) {
		fmt.Println(string(b))
	}
}

func PrettyFullRequest(req *http.Request) {
	b, err := httputil.DumpRequest(req, true)

	if noError(err) {
		fmt.Println(string(b))
	}
}

func RequestHeaders(req *http.Request) {
	AsJson(req.Header, false)
}

func RequestBody(req *http.Request, pretty bool) {
	if req.Body == nil {
		return
	}

	bytes, err := ioutil.ReadAll(req.Body)
	if !noError(err) {
		return
	}
	JsonBytes(bytes, pretty)
}
