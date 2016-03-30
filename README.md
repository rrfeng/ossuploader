# ossuploader

A super simple command to upload file to aliyun oss.

## Usage

```
Usage of ./ossuploader:
  -b string
        Bucket to store file.
  -c string
        Location of config file. Default is "$HOME/.ossuploader/config.json".
  -d    Delete action.
  -e string
        Endpoint of ali OSS service.
  -i string
        File to upload.
  -o string
        Name to save file in bucket. Default is the name of uploading file.
  -r int
        Number of routines for parallel upload. (default 1)
```

## Config File Format: JSON

```
{
  "accesskey": "MY_ACCESS_KEY",
  "secretkey": "MY_ACCESS_KEY_SECRET"
}
```

## Example

### UPLOAD:
`./ossuploader -e oss-cn-beijing-internal.aliyuncs.com -c ./config.json -b my-bucket -i test.file -o tmp/test.bin`

### DELETE:
`./ossuploader -e oss-cn-beijing-internal.aliyuncs.com -b my-bucket -d -o tmp/test.bin`

