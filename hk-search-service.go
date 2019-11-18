package main

import (
	"./api"
	"./cache"
	"flag"
	"fmt"
	"os"
)

func main() {

	query_name := flag.String("query", "", "Please Enter a Query, this program takes an argument -query = \"<your query>\" ")
	flag.Parse()

	if len(*query_name) == 0 {
		fmt.Println("Please enter a query, this program takens an argument -query = \"Your Search Query\"")
		os.Exit(1)
	}

	var req_data string
	var processed_data []string
	var status, stat bool
	var output_dict map[string]string

	res, flag := cache.Check_cache(*query_name)
	if flag == true {
		fmt.Println("(Cached Result) \n",
			"Heading :", res.Heading, "\n",
			"AbstractURL : ", res.AbstractURL, "\n",
			"AbstractText", res.AbstractText, "\n")
	} else {

		req_data, status = api.Request_data(*query_name)
		if status {
			processed_data, stat = api.Extract_script_text(req_data)
			if stat {
				output_dict = api.Extract_data(processed_data)
				data := cache.Detail{*query_name, output_dict["AbstractURL"], output_dict["AbstractText"], output_dict["Heading"]}
				cache.Insert_into_cache(data)
				for key, value := range output_dict {
					fmt.Println(key, ":", value)
				}
			} else {
				fmt.Println("Could not find any result")
			}
		} else {
			fmt.Printf("%s\n", req_data)
		}
	}
}
