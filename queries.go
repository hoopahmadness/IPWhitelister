package main

/*
performance boost: use a lookup with a simplified struct to only get the fields I need
"If you only need several fields, you may get superior performance by using maxminddb's Lookup
directly with a result struct that only contains the required fields.
(See example_test.go in the maxminddb repository for an example of this.)"
*/
import (
	"github.com/oschwald/geoip2-golang"

	"errors"
	"log"
	"net"
)

func lookupIP(ipStr, lang string) (string, error) {
	db, err := geoip2.Open(DBFILENAME)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	ip := net.ParseIP(ipStr)
	if ip == nil {
		log.Print(INVALIDIP)
		return "", errors.New(INVALIDIP)
	}

	record, err := db.Country(ip)
	if err != nil {
		log.Print(err.Error())
		return "", err
	}

	country := record.Country.Names[lang]
	if country == "" {
		country = record.Country.Names["en"]
	}
	return country, nil

}
