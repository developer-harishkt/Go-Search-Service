package api

import (
		"io/ioutil"
		"strings"
		"net/http"
		"github.com/PuerkitoBio/goquery"
		"bytes"
	)

	func Request_data(request_param string) (string, bool) {

		kp := -1 // Safe Search:	kp = 1 for On; kp = -1 for Moderate; kp = -2 for Off.
		kl := "in-en" // Region : in-en for India
		apiUrl := "https://duckduckgo.com/"

		url := apiUrl + "?q=" + string(request_param) + "&kp=" + string(kp) + "&kl=" + kl

		response, err := http.Get(url)

		if err != nil {
			return err.Error(), false
		} else {
			data, _ := ioutil.ReadAll(response.Body)
			return string(data), true
		}
	}

	func process_script_texts(script_texts []string) ([]string, bool) {

		status := false
		var remove_closing_paranthesis string

		for _, texts := range script_texts {
			if strings.Contains(texts, "DDG.ready(function () {DDG.") {
				status = true
				rtexts := strings.Replace(texts, "DDG.ready(function () DDG.duckbar.add(", "", -1)
				replace_comma := strings.Replace(rtexts, ", ", "-", -1) // To protect the slicing of details in AbstractText
				remove_opening_paranthesis := strings.Replace(replace_comma, "{", "", -1)
				remove_closing_paranthesis = strings.Replace(remove_opening_paranthesis, "}", "", -1)
				break
			}
		}

		processed_data := strings.Split(remove_closing_paranthesis, ",")
		return processed_data, status
	}

	func Extract_script_text(resp_data string) ([]string, bool) {

		resp_data_reader := bytes.NewReader([]byte(string(resp_data)))
		resp_data_doc, _ := goquery.NewDocumentFromReader(resp_data_reader)

		var script_texts []string
		resp_data_doc.Find("Script").Each(func(i int, script *goquery.Selection) {
			script_texts = append(script_texts, script.Text())
		})

		return process_script_texts(script_texts)
	}

	func slice_data_for_map(data string) (string, string) {

		sliced_data := strings.Split(data, ":\"")
		return strings.Replace(sliced_data[0], "\"", "", -1), strings.Replace(sliced_data[1], "\"", "", -1)
	}

	func Extract_data(processed_data []string) map[string]string {

		output_data := make(map[string]string)
		extract_count := 0
		key := ""
		value := ""

		for index, data := range processed_data {

			if extract_count < 3 {

				if strings.Contains(data, "Heading") {

					key, value = slice_data_for_map(data)
					output_data[key] = value
					extract_count += 1

				} else if strings.Contains(data, "AbstractURL") {

					key, value = slice_data_for_map(data)
					output_data[key] = value
					extract_count += 1

				} else if strings.Contains(data, "AbstractText") {


					for strings.Contains(processed_data[index + 1], "\":") != true {
							data += processed_data[index + 1]
							index += 1
					}

					key, value = slice_data_for_map(data)
					output_data[key] = value
					extract_count += 1
				}

			} else {
				break
			}
		}

		return output_data
	}
