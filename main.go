package main

import (
	"flag"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"log"
	"path/filepath"
)

var src *string = flag.String("i", "", "File to upload")
var key *string = flag.String("o", "", "Name to save file in bucket")
var bkt *string = flag.String("b", "", "Bucket to store file")
var accesskey *string = flag.String("a", "", "AccessKey of OSS account")
var secretkey *string = flag.String("s", "", "SecretKey of OSS account")
var routines *int = flag.Int("r", 1, "Number of routines for parallel upload")
var endpoint *string = flag.String("e", "", "Endpoint of ali OSS service")
var delete *bool = flag.Bool("d", false, "Delete action.")

func main() {
	flag.Parse()

	switch {
	case *src == "" && *delete == false:
		log.Panic("No input file!")
	case *bkt == "":
		log.Panic("No bucket name!")
	case *accesskey == "":
		log.Panic("No access key!")
	case *secretkey == "":
		log.Panic("No secret key!")
	case *endpoint == "":
		log.Panic("Please set the oss endpoint!")
	}

	client, err := oss.New(*endpoint, *accesskey, *secretkey)
	if err != nil {
		log.Panic(err.Error())
	}

	bucket, err := client.Bucket(*bkt)
	if err != nil {
		log.Panic(err.Error())
	}

	if *delete {
		err = bucket.DeleteObject(*key)
		if err != nil {
			log.Panic(err.Error())
		} else {
			log.Println(*key + " deleted!")
		}
	} else {
		f_name := filepath.Base(*src)

		err = bucket.UploadFile(*key, *src, 10*1024*1024, oss.Routines(*routines), oss.Checkpoint(true, "/tmp/."+f_name+".osstmp"))
		if err != nil {
			log.Panic(err.Error())
		} else {
			log.Println(*src + " upload success!")
		}
	}
}
