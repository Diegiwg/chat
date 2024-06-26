package chat_ai

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func Req(payload Payload) (Response, error) {
	url := "https://duckduckgo.com/duckchat/v1/chat"
	method := "POST"

	payloadData, _ := json.Marshal(payload)
	body := bytes.NewBuffer([]byte(payloadData))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return Response{}, err
	}

	req.Header.Add("accept", "text/event-stream")
	req.Header.Add("accept-language", "en-US,en;q=0.9")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("cookie", "dcm=6")
	req.Header.Add("origin", "https://duckduckgo.com")
	req.Header.Add("priority", "u=1, i")
	req.Header.Add("referer", "https://duckduckgo.com/")
	req.Header.Add("sec-ch-ua", `"Not)A;Brand";v="99", "Google Chrome";v="127", "Chromium";v="127"`)
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("sec-ch-ua-platform", `"Linux"`)
	req.Header.Add("sec-fetch-dest", "empty")
	req.Header.Add("sec-fetch-mode", "cors")
	req.Header.Add("sec-fetch-site", "same-origin")
	req.Header.Add("user-agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.0.0 Safari/537.36")
	req.Header.Add("x-vqd-4", "4-241511294805875811649864678132297207880")

	res, err := client.Do(req)
	if err != nil {
		return Response{}, err
	}
	defer res.Body.Close()

	response := Response{
		Data: []ResponseMessage{},
		Text: "",
	}

	scanner := bufio.NewScanner(res.Body)
	for scanner.Scan() {

		content := scanner.Text()

		if content == "data: [DONE]" {
			break
		}

		if content == "" {
			continue
		}

		if strings.Contains(content, "status") {
			fmt.Println(content)
			continue
		}

		passed, _ := strings.CutPrefix(content, "data: ")

		var data ResponseMessage
		json.Unmarshal([]byte(passed), &data)
		response.Data = append(response.Data, data)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	final := ""
	for _, v := range response.Data {
		final += v.Message
	}

	response.Text = strings.TrimSpace(final)

	return response, nil
}
