package main

//to update the DB, simply shut down the server, delete the DB, and restart the web server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const (
	DBFILENAME  = "./DB/GeoLite2-Country_20200616/GeoLite2-Country.mmdb" //In the future, remove the data-specific directory and use a static or "latest" dir
	TARFILENAME = "./DB/GeoIP2-Country.tar.gz"
)

func main() {
	//check if DB is present, DL if missing
	if !fileExists(DBFILENAME) {
		DownloadDB()
		unpackDB()
		setDBLatest()
	}

	//start listener
	http.HandleFunc("/ipCheck/", handler)
	log.Fatal(http.ListenAndServe(":4567", nil))

}

//set up handler
func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Expecting 'POST'", http.StatusMethodNotAllowed)
		return
	}

	var contains bool
	//parse out IP and whitelisted countries
	bytes, err := ioutil.ReadAll(r.Body)
	var ipReq IPCheckRequest
	err = json.Unmarshal(bytes, &ipReq)
	if err != nil {
		http.Error(w, err.Error(), 400)
		log.Print(err)
		return
	}
	err = ipReq.validate()
	if err != nil {
		http.Error(w, err.Error(), 400)
		log.Print(err)
		return
	}
	//get country for IP
	country, err := lookupIP(ipReq.IP, ipReq.Lang)
	if err != nil {
		http.Error(w, err.Error(), 500)
		log.Print(err)
		return
	}

	//compare to list
	for _, elem := range ipReq.Whitelist {
		contains = contains || elem == country
	}

	out := map[string]string{
		"ip":      ipReq.IP,
		"Country": country,
	}
	outbytes, _ := json.Marshal(out)

	//send back IP with good or bad status code
	if contains {
		fmt.Fprint(w, string(outbytes)) //200 means the country is whitelisted
	} else {
		w.WriteHeader(422)
		fmt.Fprint(w, string(outbytes))
	}

}

func fileExists(name string) bool {

	info, err := os.Stat(name)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()

}

func dirExists(name string) bool {
	info, err := os.Stat(name)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()

}
