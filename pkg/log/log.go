package log

import (
	"log"
	"os"
	"runtime"
	"strconv"
	"strings"
)

const (
	DEFAULT_MESSAGE_DELIM = " - "
	PREFIX_INFO           = "[ INFO ]"
	PREFIX_WARN           = "[ WARN ]"
	PREFIX_ERR            = "[ ERROR ]"
)

func Info(messages ...string) {
	log.Println(PREFIX_INFO, strings.Join(messages, DEFAULT_MESSAGE_DELIM))
}

func Infof(fstring string, variables ...interface{}) {
	infoString := strings.Join([]string{PREFIX_INFO, fstring}, " ")
	log.Printf(infoString, variables)
}

func Error(err error, messages ...string) {
	logErrorWarning(PREFIX_ERR, err, messages)
}

func Warn(err error, messages ...string) {
	logErrorWarning(PREFIX_WARN, err, messages)
}

func logErrorWarning(prefix string, err error, messages []string) {
	workDir, _ := os.Getwd()
	_, file, line, _ := runtime.Caller(2)
	file = strings.TrimPrefix(file, workDir+"/")

	errLocation := file + ":" + strconv.Itoa(line)

	messagesString := strings.Join(messages, DEFAULT_MESSAGE_DELIM)
	errString := strings.Join(
		[]string{
			errLocation,
			messagesString,
			err.Error()}, DEFAULT_MESSAGE_DELIM)

	log.Println(prefix, errString)
}
