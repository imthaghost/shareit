package upload

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

// Innermost ...
type Innermost struct {
	Full  string `json:"full"`
	Short string `json:"short"`
}

// Inner ...
type Inner struct {
	URL Innermost `json:"url"`
}

// Outer ...
type Outer struct {
	File Inner `json:"file"`
}

// Outmost ...
type Outmost struct {
	Status bool  `json:"status"`
	Data   Outer `json:"data"`
}

// FileUpload will take any arbitrary url and submit the provided file with a specified key argument
func FileUpload(url string, filepath string, key string) Outmost {

	// prepare file
	var b bytes.Buffer

	w := multipart.NewWriter(&b)
	// add the file
	f, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	// create form field for file
	fw, err := w.CreateFormFile(key, filepath)
	if err != nil {
		panic(err)
	}
	// copy the file
	if _, err = io.Copy(fw, f); err != nil {
		panic(err)
	}
	// Add the other fields
	if fw, err = w.CreateFormField(key); err != nil {
		panic(err)
	}
	if _, err = fw.Write([]byte(key)); err != nil {
		panic(err)
	}
	// close
	w.Close()
	// set up request for submission
	req, err := http.NewRequest("POST", url, &b)
	if err != nil {
		panic(err)
	}
	// set header content type
	req.Header.Set("Content-Type", w.FormDataContentType())

	// submit the request
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	// check the response
	if res.StatusCode != http.StatusOK {
		err = fmt.Errorf("bad status: %s", res.Status)
	}

	jsonData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	var re Outmost
	er := json.Unmarshal(jsonData, &re)
	if er != nil {
		panic(er)
	}
	return re
}
