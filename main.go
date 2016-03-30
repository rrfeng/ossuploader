package main

import (
	"encoding/json"
	"flag"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

var src *string = flag.String("i", "", "File to upload.")
var key *string = flag.String("o", "", "Name to save file in bucket. Default is the name of uploading file.")
var bkt *string = flag.String("b", "", "Bucket to store file.")
var delete *bool = flag.Bool("d", false, "Delete action.")
var cfgfile *string = flag.String("c", "", "Location of config file. Default is \"$HOME/.ossuploader/config.json\".")
var routines *int = flag.Int("r", 1, "Number of routines for parallel upload.")
var endpoint *string = flag.String("e", "", "Endpoint of ali OSS service.")

type config struct {
	AccessKey string `json: accesskey`
	SecretKey string `json: secretkey`
}

func main() {
	flag.Parse()

	switch {
	case *delete == false && *src == "":
		log.Fatalln("No input file!")
	case *delete == false && *key == "":
		*key = *src
	case *delete && *key == "":
		log.Fatalln("Delete action need objectKey!")
	case *bkt == "":
		log.Fatalln("No bucket name!")
	case *endpoint == "":
		log.Fatalln("Please set the oss endpoint!")
	case *cfgfile == "":
		*cfgfile = os.Getenv("HOME") + "/.ossuploader/config.json"
	}

	var cfg config
	if _, err := os.Stat(*cfgfile); os.IsNotExist(err) {
		log.Fatalln(err.Error())
	} else {
		data, err := ioutil.ReadFile(*cfgfile)
		if err != nil {
			log.Fatalln(err.Error())
		}
		err = json.Unmarshal(data, &cfg)
		if err != nil {
			log.Fatalln(err.Error())
		}
	}

	client, err := oss.New(*endpoint, cfg.AccessKey, cfg.SecretKey)
	if err != nil {
		log.Fatalln(err.Error())
	}

	bucket, err := client.Bucket(*bkt)
	if err != nil {
		log.Fatalln(err.Error())
	}

	if *delete {
		err = bucket.DeleteObject(*key)
		if err != nil {
			log.Fatalln(err.Error())
		} else {
			log.Println(*key + " deleted!")
		}
	} else {
		f_name := filepath.Base(*src)

		err = bucket.UploadFile(*key, *src, 10*1024*1024, oss.Routines(*routines), oss.Checkpoint(true, "/tmp/."+f_name+".osstmp"))
		if err != nil {
			log.Fatalln(err.Error())
		} else {
			log.Println(*src + " upload success!")
		}
	}
}
