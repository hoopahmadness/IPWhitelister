package main

import (
	"github.com/mholt/archiver/v3"

	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func DownloadDB() {
	//use env variable for license

	license := os.Getenv("GEOIP_LICENSE")
	if license == "" {
		log.Fatal("Missing Env Variable 'GEOIP_LICENSE'; please set your license")
	}
	url := fmt.Sprintf("https://download.maxmind.com/app/geoip_download?edition_id=GeoLite2-Country&license_key=%s&suffix=tar.gz", license)
	if !dirExists("./DB/") {
		err := os.Mkdir("./DB/", 0755)
		if err != nil {
			fmt.Println('a')
			log.Fatal(err)
		}

	}
	DBtar, err := os.Create(TARFILENAME)
	if err != nil {
		log.Fatal(err)
	}
	defer DBtar.Close()

	resp, err := http.Get(url)
	if err != nil { //remove empty file so that it doesn't give a false positive on next run
		_ = os.Remove(TARFILENAME)
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		_ = os.Remove(TARFILENAME)
		log.Fatal(fmt.Sprintf("Got bad status from GeoLite; %d", resp.StatusCode))
	}
	_, err = io.Copy(DBtar, resp.Body)
	if err != nil {
		_ = os.Remove(TARFILENAME)
		log.Fatal(err)
	}
}

func unpackDB() {
	err := archiver.Unarchive(TARFILENAME, "./DB/")
	if err != nil {
		log.Fatal(err)
	}
	err = os.Rename(TARFILENAME, fmt.Sprintf("%s-%s", TARFILENAME, time.Now().String()))
	if err != nil {
		log.Print("Unable to rename leftover tarball")
	}
}

//Call this function to move the downloaded DB to a pre-defined directory
func setDBLatest() {
	//stub
	//For this codetest I'll just use the directory that the DB untars into
}
