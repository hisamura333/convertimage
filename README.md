# convertimage

convertimage is a CLI command to convert image.


## Build and Run

Get source files.
```$xslt
$ go get github.com/hisamura333/convertimage
```

Run using go run.
```
$ go run cmd/convertimage/main.go

```

## option
### -d string
   変換後の拡張子 (default "jpg")
   
### -s string
   変換前の拡張子 (default "png")
   
### -dir string
   変換したいディレクトリ配下 (default "./")
  
### -r bool   
   変換前の拡張子ファイルを削除するかのflag
  


## Running the tests

```cassandraql
$ go test -v 
```