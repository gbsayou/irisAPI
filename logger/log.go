package logger

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"time"
)

/**
 * 将日期作为日志文件的名字
 * @method todayFileName
 * @returns {string} 文件名
 */
func todayFileName() string {
	today := time.Now().Format("Jan 02 2006")
	return "./logger/logs/" + today + ".txt"
}

/**
 * 生成新文件
 * @method newLogFile
 * @returns {*os.File} 文件
 */
func newLogFile() *os.File {
	filename := todayFileName()
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("创建日志文件失败")
		panic(err)
	}
	return f
}

//写日志
func Logger(content interface{}) {
	f := newLogFile()
	w := bufio.NewWriter(f)
	contentStr, err := json.Marshal(content)//将内容转换成字符串
	if err != nil {
		fmt.Println("转字符串失败")
	}
	_, err1 := w.WriteString(string(contentStr) + "\n")
	if err1 != nil {
		fmt.Println("写入文件错误", err1)
	}
	// fmt.Println(wr)
	w.Flush()
}
