# IPWhitelister
Codetest

Forgive my simple README formatting;

This service has a single endpoint at '/ipCheck/' that expects a json object with the following fields:
-ip - a string representation of the requesting client's IP address
-whitelist - (Optional) an array of strings with country names
-lang - (Optional) a string representing a language

If Whitelister ascertains that the country of origin for that IP is in the white list, it will respond with a code 200 and a json body includeing the IP string and returned country name. If the whitelist does not contain that country, it will issue a 422 with the same body.

When the service is started, it will attempt to download the GeoLite2 Country database for local storage. If the DB is not already present at `./DB/GeoLite2-Country_20200616/GeoLite2-Country.mmdb`, you must export a license key to GEOIP_LICENSE to run the service. I added this feature because it simplifies the process of obtaining and updating the DB; if you have a license you can just run the service and not worry about it. If you want to make sure the DB is up to date, you can stop the service, delete the `./DB/` directory (or just rename the db file to be safe), and start it again. It's not a total solution, but it makes automation with something like a cron job much simpler. Note that I didn't want to spend too much time fooling around with the file system (this feature took up a significant fraction of the 4 hours allotted), so for now I'm using a hardwired directory `./DB/Country_20200616/` to store the DB. This works fine for now but when a new DB is released this will break. Next steps would be to dynamically link to the untarred DB or move it to a static folder.

I also added the ability to change the language used to query the DB, as seen in the request body above. This would not be useful if all of your customer DBs use the English country names, but I figured it was a fun bit of flexibility with little dev time. If you have whitelists in Spanish, for example, you can send "es". If a language code is blank or not included in the DB it just defaults to "en"