package www

import (
	"net/http"
	"log"
	"fmt"
	"os"
	"time"
	"strings"
	"io"
	"bytes"
	"net/url"
	"strconv"
	"errors"
)


func SaveResponse(response *http.Response) (string, error) {
	var fileName = fmt.Sprintf("./tmp/%d", time.Now().Unix())

	// TODO: check file existence first with io.IsExist
	output, err := os.Create(fileName)
	if err != nil {
		fmt.Print("Error while creating", fileName, "-", err)
		return 	"", err
	}
	defer output.Close()

	defer response.Body.Close()

	n, err := io.Copy(output, response.Body)
	if err != nil {
		log.Println("Error while downloading ", fileName, "-", err)
		return "", err
	}

	log.Println(n /1000, "Kbytes downloaded.")
	return fileName, nil
}

func DumpBody(res* http.Response, fn string) (string, error) {
	tokens := strings.Split(fn, "/")
	var fileName = fmt.Sprintf("./tmp/%d_%s", time.Now().Unix(), tokens[len(tokens)-1])
	// TODO: check file existence first with io.IsExist
	output, err := os.Create(fileName)
	if err != nil {
		fmt.Print("Error while creating", fileName, "-", err)
		return 	"", err
	}
	defer output.Close()
	defer res.Body.Close()

	count, err := io.Copy(output, res.Body)
	if err != nil {
		log.Println("Error while downloading", fileName, "-", err)
		return "", err
	}

	count += 1

	return fileName, nil
}

func Download(url string) (string, error) {
	tokens := strings.Split(url, "/")
	var fileName = fmt.Sprintf("./tmp/%d_%s", time.Now().Unix(), tokens[len(tokens)-1])
	log.Print("Downloading ", url, " to ", fileName)

	// TODO: check file existence first with io.IsExist
	output, err := os.Create(fileName)
	if err != nil {
		fmt.Print("Error while creating", fileName, "-", err)
		return 	"", err
	}
	defer output.Close()

	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error while downloading", url, "-", err)
		return "", err
	}
	defer response.Body.Close()

	n, err := io.Copy(output, response.Body)
	if err != nil {
		log.Println("Error while downloading", url, "-", err)
		return "", err
	}

	log.Println(n /1000, "Kbytes downloaded.")
	return fileName, nil
}


var RedirectAttemptedError = errors.New("--OK--")
func noRedirect(req *http.Request, via []*http.Request) error {
    return RedirectAttemptedError;
}

func PostForm(urlPath string, formValues map[string]string) (*Context, error) {
	form := ""
	for k, v := range formValues {
		form += k + "=" + v + "&";
	}
	form = form[:len(form) -1]
	client := &http.Client{ CheckRedirect: noRedirect }
	req, _ := http.NewRequest("POST", urlPath, bytes.NewBufferString(form))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(form) - 1))
	
	resp, _ := client.Do(req)

	cxt := &Context{}
	cxt.Cookies = map[string]string{}
	for _, c := range resp.Cookies() {
		cxt.Cookies[c.Name] = c.Value
	}

	return cxt, nil
}

func Do(cxt* Context) (*http.Response, error) {

	data := url.Values{}
	for k, v := range cxt.Data {
		data.Add(k, v)
	}

	client := &http.Client{}
	r, _ := http.NewRequest(cxt.Method, cxt.URL, bytes.NewBufferString(data.Encode()))
	for k, v := range cxt.Headers {
		r.Header.Add(k, v)
	}

	for k, v := range cxt.Cookies {
		r.AddCookie(&http.Cookie{Name: k, Value: v, HttpOnly: false})
	}

	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	return client.Do(r)
}