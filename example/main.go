package main

import (
	"fmt"
	"log"

	"github.com/zen37/apachelogparser"
)

/*
Apache Common Log
127.0.0.1 - frank [10/Oct/2000:13:55:36 -0700] "GET /apache_pb.gif HTTP/1.0" 200 2326

Apache Combined Log
127.0.0.1 - Scott [10/Dec/2019:13:55:36 -0700] "GET /server-status HTTP/1.1" 200 2326 "http://localhost/" "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.108 Safari/537.36"
64.180.189.161 - - [23/Dec/2021:18:41:18 -0600] "GET /cpsess6383042128/3rdparty/phpMyAdmin/themes/pmahomme/img/s_lang.png HTTP/2" 200 659 "https://cpanel.findroommates.net/cpsess6383042128/3rdparty/phpMyAdmin/phpmyadmin.css.php?nocache=6470895733ltr&server=1" "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36"
*/

//Apache Common Log
//const logApache = "127.0.0.1 - frank [10/Oct/2000:13:55:36 -0700] \"GET /apache_pb.gif HTTP/1.0\" 200 2326"

//Apache Combined Log
const logApache = "127.0.0.1 - Scott [10/Dec/2019:13:55:36 -0700] \"GET /server-status HTTP/1.1\" 200 2326 \"http://localhost/\" \"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.108 Safari/537.36\""

//const logApache = "154.53.43.164 - - [23/Dec/2021:14:41:34 -0600] \"GET /wp/installer-backup.php HTTP/1.1\" 301 707 \"-\" \"-\""
//const logApache = "154.53.43.164 - - [23/Dec/2021:14:41:37 -0600] \"GET /old/installer.php HTTP/1.1\" 301 707 \"-\" \"curl/7.68.0\""

func main() {

	fmt.Println(logApache)
	result, err := apachelogparser.ParseLogRecord(logApache)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(result)
}
