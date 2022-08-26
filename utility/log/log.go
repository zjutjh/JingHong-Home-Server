package JHLog

import (
	"fmt"
	"io"
	"os"
	"time"
)

var LogMultiWriter io.Writer

func init() {
	offset := int(time.Monday - time.Now().Weekday())
	dataFrom := time.Now().AddDate(0, 0, offset)
	fileName := fmt.Sprintf("./log/%v.log", dataFrom.Format("2006-01-02"))
	f, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
	if err != nil {
		return
	}
	LogMultiWriter = io.MultiWriter(os.Stdout, f)
}
