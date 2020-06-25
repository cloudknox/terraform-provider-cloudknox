package utils

import (
	"crypto/md5"
	"crypto/sha1"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/url"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

func Truncate(str string, n int) string {
	if len(str) < n {
		return str
	} else {
		return str[0:n] + " ... <TRUNCATED>"
	}
}

func GetFileName(orgId, dataType, ext string, ts int64, pageIndex int) string {
	var fileName string
	if pageIndex == 0 {
		fileName = fmt.Sprintf("%s_%s_%d.%s", dataType, orgId, ts, ext)
	} else {
		fileName = fmt.Sprintf("%s_%s_%d_%d.%s", dataType, orgId, ts, pageIndex, ext)
	}
	return fileName
}

func GetFileSuffix(orgId, dataType string, ts int64) string {
	return fmt.Sprintf("%s_%s_%d", dataType, orgId, ts)
}

func DeferHandler(l *logrus.Logger, e interface{}) {
	var err error
	switch v := e.(type) {
	case *io.PipeReader:
		err = v.Close()
	case *io.PipeWriter:
		err = v.Close()
	case *multipart.Writer:
		err = v.Close()
	case *os.File:
		err = v.Close()
	default:
		l.Errorf("Failed To Close Deferred Entity [Unknown Entity]", err.Error())
		err = nil
	}
	if err != nil {
		l.Errorf("Failed To Close Deferred Entity [%s]", err.Error())
	}
}

func CheckIfPathExists(dir string) bool {
	_, err := os.Stat(dir)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func ChunkArrayIndexes(arraySize int, chunkSize int) [][]int {
	var chunkList []int
	var finalList [][]int
	for i := 0; i < arraySize; i++ {
		chunkList = append(chunkList, i)
	}
	for i := 0; i < arraySize; i = i + chunkSize {
		finalList = append(finalList, chunkList[i:getMin(arraySize, i+chunkSize)])
	}
	return finalList
}

func ConstructServiceResourceType(ast string, service string, resourceType string) string {
	var separator string
	switch ast {
	case "AZURE":
		separator = "/"
	case "AWS":
		separator = ":"
	case "GCP":
		separator = ":"
	case "NSX":
		separator = "/"
	default:
		separator = ""
	}

	if ast == "VCENTER" {
		if resourceType == "" {
			return "%"
		} else {
			return resourceType
		}
	} else {
		if resourceType == "" {
			return fmt.Sprintf("%s%s%s", service, separator, "%")
		} else {
			return fmt.Sprintf("%s%s%s", service, separator, resourceType)
		}
	}
}

func GetMD5(s string) string {
	data := []byte(s)
	return fmt.Sprintf("%x", md5.Sum(data))
}

func ConvertStringToHashDecimal(s string) int {
	var sum = 0
	for _, v := range GetMD5(s) {
		sum = sum + int(v)
	}
	return sum
}

func GetSHA1(s string) string {
	data := []byte(s)
	return fmt.Sprintf("%x", sha1.Sum(data))
}

func ParseJDBCUrl(s, username, password string) (string, error) {
	u, err := url.Parse(s)
	if err != nil {
		return "", err
	}
	path := strings.ReplaceAll(u.Opaque, "mysql://", "")
	pathTokens := strings.Split(path, "/")
	if len(pathTokens) == 2 {
		cs := fmt.Sprintf("%s:%s@tcp(%s)/%s?%s", username, password, pathTokens[0], pathTokens[1], u.RawQuery)
		return cs, nil
	} else {
		return "", errors.New("failed to parse components")
	}
}

func ArrayIntersection(a, b []string) (c []string) {
	m := make(map[string]bool)

	for _, item := range a {
		m[item] = true
	}

	for _, item := range b {
		if _, ok := m[item]; ok {
			c = append(c, item)
		}
	}
	return
}

func getMin(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
