// Package convert is image conversion processing
//
// 画像変換処理。
// 変換前と変換後の拡張子を指定し、画像変換を行います。
// ディレクトリを選択し、再帰的にディレクトリ探索を行い変換します。
// 変換前のファイルを削除したい場合は、remove(-r)フラグを指定します。
package convertimage

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
)

// FlagOps is flag option
type FlagOps struct {
	Dir    string
	Src    string
	Dest   string
	Remove bool


}

// FileDetails have Extionsion and FileName info
type FileDetails struct {
	Extension string
	FileName  string
}

// Convert is image conversion processing
func Convert(flagOps FlagOps) (error) {

	var size int64
	count := 0

	err := filepath.Walk(flagOps.Dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		fileDetails := FileDetails{
			filepath.Ext(path),
			filepath.Clean(path),
		}

		if strings.HasSuffix(fileDetails.Extension, flagOps.Src) {

			decodeImg, err := decodeImage(fileDetails)
			if err != nil {
				return  err
			}

			if flagOps.Remove {
				defer os.Remove(fileDetails.FileName)
			}

			fileNameRemoveExt := strings.Replace(fileDetails.FileName, flagOps.Src, "", 1)
			dstFile, err := os.Create(fmt.Sprintf(fileNameRemoveExt + flagOps.Dest))
			if err != nil {
				return  err
			}

			defer dstFile.Close()

			err = encodeImage(decodeImg, flagOps, dstFile)
			if err != nil {
				return  err
			}
		}

		if !info.IsDir() {
			count += 1
			size += info.Size()
		}
		return nil
	})
	if err != nil {
		return err
	}

	return nil

}

func decodeImage(fileDetails FileDetails) (image.Image, error) {
	srcFile, err := os.Open(fileDetails.FileName)

	defer srcFile.Close()

	if err != nil {
		return nil, err
	}

	decodeImg, _, err := image.Decode(srcFile)
	return decodeImg, err
}

func encodeImage(decodeImg image.Image, flagOps FlagOps, dstFile *os.File) error {
	switch flagOps.Dest {
	case "jpeg", "jpg":
		err := jpeg.Encode(dstFile, decodeImg, nil)
		return err

	case "gif":
		err := gif.Encode(dstFile, decodeImg, nil)
		return err

	case "png":
		err := png.Encode(dstFile, decodeImg)
		return err

	default:
		return fmt.Errorf("Error: invalid extension")
	}
}

