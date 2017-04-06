package utils

import (
	"log"
	"os"
	"fmt"
	"time"
	"runtime"
	"strconv"
	"github.com/astaxie/beego"
)

const (
	LOG_FATAL int = iota
	LOG_ERROR
	LOG_WARNING
	LOG_INFO
	LOG_DEBUG
)

var WriteFileSize int
var MaxFileSize int = 50*1024*1024
var logger *log.Logger
var file *os.File
var log_prefix string = "homework "

func init(){
	createlog()
}

func createlog()  {

	nowtime := time.Now()

	log_path := "logs/"+nowtime.Format("2006-01-02")

	os.MkdirAll(log_path,os.ModePerm)
	os.Chmod(log_path,os.ModePerm)

	log_file := nowtime.Format("15-04-05")+".log"

	fmt.Println(log_path," ",log_file)

	var err error
	file,err = os.Create(log_path+"/"+log_file)
	if err != nil {
		log.Fatalln("fail to create test.log file!")
		return
	}
	logger = log.New(file, log_prefix, log.LstdFlags)

	WriteFileSize = 0
}

func WriteLog(level int, v ...interface{}){

	if logger == nil {
		return
	}

	str_level := ""

	switch level {
	case LOG_FATAL:
		str_level = "[FETAL] "
		break
	case LOG_ERROR:
		str_level = "[ERROR] "
		break
	case LOG_WARNING:
		str_level = "[WARNING] "
		break
	case LOG_INFO:
		str_level = "[INFO] "
		break
	case LOG_DEBUG:
		str_level = "[DEBUG] "
		break
	default:
		str_level = "[INFO] "
	}

	var ok bool
	var line int
	var codefile string
	_, codefile, line, ok = runtime.Caller(1)
	if !ok {
		codefile = "???"
		line = 0
	}

	msg := codefile+":"+strconv.Itoa(line)+" "+str_level+fmt.Sprint(v...)

	WriteFileSize += (len(msg)+len(log_prefix)+len("2016/08/20 16:08:43"))

	logger.Println(msg)

	if "true" == beego.AppConfig.String("consolelog"){
		fmt.Println(msg)
	}

	if WriteFileSize > MaxFileSize{
		createlog()
	}

}
