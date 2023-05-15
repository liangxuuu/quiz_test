package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/buger/jsonparser"
	"github.com/hotstar/hs-core-api-go/v2/request"
	"github.com/hotstar/hs-core-ui-models-go/feature/form"
	"google.golang.org/protobuf/types/known/anypb"
)

type QuizData struct {
	Question struct {
		Id    string `json:"id"`
		Title string `json:"title"`
	} `json:"question"`
	Options []struct {
		Id    string `json:"id"`
		Title string `json:"title"`
	} `json:"options"`
}

func main() {
	data := getQuizPage()
	fmt.Println("##################")
	submitQuiz(data)
}

func getQuizPage() (data QuizData) {
	url := "http://localhost:8080/v2/pages/quiz?content_id=1100010232&client_capabilities=%7B%25E2%2580%259Dpackage%25E2%2580%259D%253A%5B%2522dash%25E2%2580%259D%252C%2522hls%25E2%2580%259D%5D%252C%2522container%25E2%2580%259D%253A%5B%2522fmp4%2522%252C%2522fmp4br%25E2%2580%259D%252C%2522ts%25E2%2580%259D%5D%252C%2522ads%25E2%2580%259D%253A%5B%2522ssai%25E2%2580%259D%252C%2522non_ssai%25E2%2580%259D%5D%252C%2522audio_channel%25E2%2580%259D%253A%5B%2522stereo%25E2%2580%259D%5D%252C%2522encryption%25E2%2580%259D%253A%5B%2522plain%25E2%2580%259D%252C%2522widevine%25E2%2580%259D%5D%252C%2522video_codec%25E2%2580%259D%253A%5B%2522h264%2522%5D%252C%2522ladder%25E2%2580%259D%253A%5B%2522phone%25E2%2580%259D%5D%252C%2522resolution%25E2%2580%259D%253A%5B%2522sd%25E2%2580%259D%252C%2522hd%25E2%2580%259D%252C%2522fhd%25E2%2580%259D%5D%252C%2522dynamic_range%25E2%2580%259D%253A%5B%2522sdr%25E2%2580%259D%5D%7D&drm_parameters=%7B%250A%2520%2520%2520%2520%2520%2520%2520%2520%2522widevine_security_level%25E2%2580%259D%253A%2520%5B%250A%2520%2520%2520%2520%2520%2520%2520%2520%2520%2520%2520%2520%2522HW_SECURE_CRYPTO%25E2%2580%259D%252C%250A%2520%2520%2520%2520%2520%2520%2520%2520%2520%2520%2520%2520%2522HW_SECURE_DECODE%25E2%2580%259D%252C%250A%2520%2520%2520%2520%2520%2520%2520%2520%2520%2520%2520%2520%2522HW_SECURE_ALL%25E2%2580%259D%250A%2520%2520%2520%2520%2520%2520%2520%2520%5D%252C%250A%2520%2520%2520%2520%2520%2520%2520%2520%2522hdcp_version%25E2%2580%259D%253A%2520%5B%250A%2520%2520%2520%2520%2520%2520%2520%2520%2520%2520%2520%2520%2522HDCP_V2_1%2522%252C%250A%2520%2520%2520%2520%2520%2520%2520%2520%2520%2520%2520%2520%2522HDCP_V2_2%25E2%2580%259D%252C%250A%2520%2520%2520%2520%2520%2520%2520%2520%2520%2520%2520%2520%2522HDCP_NO_DIGITAL_OUTPUT%25E2%2580%259D%250A%2520%2520%2520%2520%2520%2520%2520%2520%5D%250A%2520%2520%2520%2520%7D&context=ongoing"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	addHeaders(req)
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
	dataBlob, _, _, err := jsonparser.Get(body, "success", "page", "spaces", "tray", "widget_wrappers", "[0]", "widget", "data")
	if err != nil {
		fmt.Println(err)
		return
	}
	if err := json.Unmarshal(dataBlob, &data); err != nil {
		fmt.Println(err)
		return
	}
	return
}

func submitQuiz(data QuizData) {
	url := "http://localhost:8080/v2/pages/quiz?content_id=1100010232&client_capabilities=%7B%25E2%2580%259Dpackage%25E2%2580%259D%253A%5B%2522dash%25E2%2580%259D%252C%2522hls%25E2%2580%259D%5D%252C%2522container%25E2%2580%259D%253A%5B%2522fmp4%2522%252C%2522fmp4br%25E2%2580%259D%252C%2522ts%25E2%2580%259D%5D%252C%2522ads%25E2%2580%259D%253A%5B%2522ssai%25E2%2580%259D%252C%2522non_ssai%25E2%2580%259D%5D%252C%2522audio_channel%25E2%2580%259D%253A%5B%2522stereo%25E2%2580%259D%5D%252C%2522encryption%25E2%2580%259D%253A%5B%2522plain%25E2%2580%259D%252C%2522widevine%25E2%2580%259D%5D%252C%2522video_codec%25E2%2580%259D%253A%5B%2522h264%2522%5D%252C%2522ladder%25E2%2580%259D%253A%5B%2522phone%25E2%2580%259D%5D%252C%2522resolution%25E2%2580%259D%253A%5B%2522sd%25E2%2580%259D%252C%2522hd%25E2%2580%259D%252C%2522fhd%25E2%2580%259D%5D%252C%2522dynamic_range%25E2%2580%259D%253A%5B%2522sdr%25E2%2580%259D%5D%7D&drm_parameters=%7B%250A%2520%2520%2520%2520%2520%2520%2520%2520%2522widevine_security_level%25E2%2580%259D%253A%2520%5B%250A%2520%2520%2520%2520%2520%2520%2520%2520%2520%2520%2520%2520%2522HW_SECURE_CRYPTO%25E2%2580%259D%252C%250A%2520%2520%2520%2520%2520%2520%2520%2520%2520%2520%2520%2520%2522HW_SECURE_DECODE%25E2%2580%259D%252C%250A%2520%2520%2520%2520%2520%2520%2520%2520%2520%2520%2520%2520%2522HW_SECURE_ALL%25E2%2580%259D%250A%2520%2520%2520%2520%2520%2520%2520%2520%5D%252C%250A%2520%2520%2520%2520%2520%2520%2520%2520%2522hdcp_version%25E2%2580%259D%253A%2520%5B%250A%2520%2520%2520%2520%2520%2520%2520%2520%2520%2520%2520%2520%2522HDCP_V2_1%2522%252C%250A%2520%2520%2520%2520%2520%2520%2520%2520%2520%2520%2520%2520%2522HDCP_V2_2%25E2%2580%259D%252C%250A%2520%2520%2520%2520%2520%2520%2520%2520%2520%2520%2520%2520%2522HDCP_NO_DIGITAL_OUTPUT%25E2%2580%259D%250A%2520%2520%2520%2520%2520%2520%2520%2520%5D%250A%2520%2520%2520%2520%7D&context=ongoing"
	method := "POST"

	anyBody, err := anypb.New(&form.FormRequest{
		FormInputs: []*form.FormInput{
			{
				FormInputId: data.Question.Id,
				FormData: &form.FormData{
					IsValid: true,
					FormValue: &form.FormData_FormValue{
						Value: &form.FormData_FormValue_OptionValue{
							OptionValue: &form.FormData_OptionValue{
								Options: []string{data.Options[0].Id},
							},
						},
					},
				},
			},
		},
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	r := &request.FetchPageRequest{
		Body: anyBody,
	}

	b, _ := json.Marshal(r)

	// b, err := proto.Marshal(r)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewReader(b))

	if err != nil {
		fmt.Println(err)
		return
	}
	addHeaders(req)
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}

func addHeaders(req *http.Request) {
	req.Header.Add("x-hs-request-id", "61a19c1a-5f9e-4061-9a55-645f6bcc0c00")
	req.Header.Add("x-country-code", "in")
	req.Header.Add("x-hs-hid", "c1eaef9b221e4806a9601eb17e5ae30d")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("x-hs-platform", "android")
	req.Header.Add("x-hs-usertoken", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJ1bV9hY2Nlc3MiLCJleHAiOjE2ODY0Njk3MzYsImlhdCI6MTY4Mzg3NzczNiwiaXNzIjoiVFMiLCJqdGkiOiJmNGIxMmVhMjI5ZGU0OGJhYWJkZGYxOGJjOGMyNzAzYiIsInN1YiI6IntcImhJZFwiOlwiMjA2ZWM1MDJlNWY2NGI5Nzk2MTAwMGUwZDdhYmEwNDZcIixcInBJZFwiOlwiOTk3OTM1OWRhNzkwNDRmNGIzNmExNjU5NmY5OWYwMjRcIixcIm5hbWVcIjpcIkFEVUxUXCIsXCJwaG9uZVwiOlwiNzAzNDcyMjY3NFwiLFwiaXBcIjpcIjEzLjIzMy42LjE5OFwiLFwiY291bnRyeUNvZGVcIjpcImluXCIsXCJjdXN0b21lclR5cGVcIjpcIm51XCIsXCJ0eXBlXCI6XCJwaG9uZVwiLFwiaXNFbWFpbFZlcmlmaWVkXCI6ZmFsc2UsXCJpc1Bob25lVmVyaWZpZWRcIjp0cnVlLFwiZGV2aWNlSWRcIjpcIjY0YjJkZDYzLTI2ZDYtNDJiMi1hMzQyLTZlODQ1Mzc0MTM0YVwiLFwicHJvZmlsZVwiOlwiQURVTFRcIixcInZlcnNpb25cIjpcInYyXCIsXCJzdWJzY3JpcHRpb25zXCI6e1wiaW5cIjp7XCJIb3RzdGFyVklQXCI6e1wic3RhdHVzXCI6XCJTXCIsXCJleHBpcnlcIjpcIjIwMjQtMDUtMDhUMjA6MDE6MjcuMDAwWlwiLFwic2hvd0Fkc1wiOlwiMVwiLFwiY250XCI6XCIxXCJ9fX0sXCJlbnRcIjpcIkNxa0JDZ1VLQXdvQkJSS2ZBUklIWVc1a2NtOXBaQklEYVc5ekVnbGhibVJ5YjJsa2RIWVNCbVpwY21WMGRoSUhZWEJ3YkdWMGRoSUVjbTlyZFJJRGQyVmlFZ1J0ZDJWaUVnZDBhWHBsYm5SMkVnVjNaV0p2Y3hJR2FtbHZjM1JpRWdwamFISnZiV1ZqWVhOMEVnUjBkbTl6RWdSd1kzUjJFZ05xYVc4U0IycHBieTFzZVdZYUFuTmtHZ0pvWkJvRFptaGtJZ056WkhJcUJuTjBaWEpsYnlvSVpHOXNZbmsxTGpGWUFRb05FZ3NJWkRnQlFBRlEwQVZZQVFvaUNob0tEaElGTlRVNE16WVNCVFkwTURRNUNnZ2lCbVpwY21WMGRoSUVPR1JZQVFxcEFRb0ZDZ01LQVFBU253RVNCMkZ1WkhKdmFXUVNBMmx2Y3hJSllXNWtjbTlwWkhSMkVnWm1hWEpsZEhZU0IyRndjR3hsZEhZU0JISnZhM1VTQTNkbFloSUViWGRsWWhJSGRHbDZaVzUwZGhJRmQyVmliM01TQm1wcGIzTjBZaElLWTJoeWIyMWxZMkZ6ZEJJRWRIWnZjeElFY0dOMGRoSURhbWx2RWdkcWFXOHRiSGxtR2dKelpCb0NhR1FhQTJab1pDSURjMlJ5S2daemRHVnlaVzhxQ0dSdmJHSjVOUzR4V0FFU1l3Z0JFTmpUcjg3MU1SbzBDaEpFU0ZOV1NWQXVTVTR1V1dWaGNpNHpPVGtTQ2todmRITjBZWEpXU1ZBYUJGTmxiR1lnd0lQcmtJQXhLTmpUcjg3MU1TZ0JNQUU2SGdvYVJFaFRWa2xRTGtGa2MwWnlaV1V1U1U0dVdXVmhjaTQyT1RrUUFRPT1cIixcImlzc3VlZEF0XCI6MTY4Mzg3NzczNjM1NCxcIm1hdHVyaXR5TGV2ZWxcIjpcIjE4K1wiLFwiaW1nXCI6XCI3XCIsXCJkcGlkXCI6XCI5OTc5MzU5ZGE3OTA0NGY0YjM2YTE2NTk2Zjk5ZjAyNFwiLFwic3RcIjoxLFwiZGF0YVwiOlwiQ2dRSUFCSUFDZ1FJQURvQUNnb0lBQ0lHZ0FFU2lBRUJDZ1FJQURJQUNnUUlBRUlBQ25rSUFDcDFDZ0lLQUFvRUNnSUlBZ3BwQ2djSUFSVUFBQUJBRWdvS0EyMWhjaVdFQ3h3OUVnb0tBMnRoYmlWU1JPYzdFZ29LQTJKbGJpVk5EMjA5RWdvS0EzUmhiU1c4TWZJOUVnb0tBM1JsYkNXS1N2QTlFZ29LQTIxaGJDVnljZTg4RWdvS0EyaHBiaVVpelJJL0Vnb0tBMlZ1WnlWQ2VYQTlcIn0iLCJ2ZXJzaW9uIjoiMV8wIn0.oiNzanOOmFdZnHCKjuI6azMQ-3Yz8w1bvoCW72afYww")
	req.Header.Add("User-Agent", "android")
	req.Header.Add("hotstarauth", "st=1637856223~exp=16378562236000~acl=/*~hmac=3240e5bfef8fe8566cb49670cf91032507f221c282582c7db215758f5b975814")
	req.Header.Add("X-HS-ForceRefresh", "true")
	req.Header.Add("x-hs-client", "platform:android;app_id:in.startv.hotstar;app_version:18.5.22.1;os:Android;os_version:9.0")
}
