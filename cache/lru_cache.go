package cache

import (
		"encoding/json"
		"os"
		"io/ioutil"
	)

type Detail struct{
		Search string      	`json:Search`
		AbstractURL string 	`json:AbstractURL`
		AbstractText string `json:AbstractText`
		Heading string 		`json:Heading`
	}

type Details struct{
	Details []Detail 	`json:Details`
}

var file_name = ".lrucache.json"

func shift(details []Detail, index int, data Detail, length int) ([]Detail){
	copy(details[index:], details[index+1:])
	details = details[:length-1]
	details = append(details, data)
	return details
}

func cache_file() (Details){

	json_file, _ := os.OpenFile(file_name, os.O_RDWR|os.O_CREATE, 0666)
	byte_value, _ := ioutil.ReadAll(json_file)
	var details Details
	json.Unmarshal(byte_value, &details)
	return details
}

func Check_cache(searchString string) (*Detail, bool){
	file_descriptor := cache_file()
	length := len(file_descriptor.Details)
	details := file_descriptor.Details
	for i := 0; i < length; i++ {
		if searchString == details[i].Search {
			curr_det := details[i]
			details = shift(details, i, curr_det, length)
			json_file, _ := os.OpenFile(file_name, os.O_RDWR|os.O_CREATE, 0666)
			store, _ := json.Marshal(Details{details})
			json_file.Write(store)
			return &curr_det, true
		}
	}
	return nil, false
}

func Insert_into_cache(detail Detail){
	file_descriptor := cache_file()
	details := file_descriptor.Details
	length := len(details)
	if length == 5 {
		shift(details, 0, detail, length)
	}else{
		details = append(details, detail)
		}
	json_file, _ := os.OpenFile(file_name, os.O_RDWR|os.O_CREATE, 0666)
	store, _ := json.Marshal(Details{details})
	json_file.Write(store)

}
