package logger

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

var (
	Debug   *logJson
	Warning *logJson
	Info    *logJson
	Error   *logJson
	Fatal   *logJson
)

type logJson struct {
	*log.Logger
	Keys map[string]interface{}
}

func (l *logJson) Json(val map[string]interface{}) {
	for k, v := range l.Keys {
		val[k] = v
	}
	val["create_time"] = time.Now().Format(time.RFC3339)

	data, err := json.Marshal(val)
	if err != nil {
		data = []byte(fmt.Sprintf(`{"json.Marshal error": "%s", "val": "%s"}`, err.Error(), val))
	}

	if val["level"] == "fatal" {
		l.Fatal(string(data[:]))
	} else {
		l.Print(string(data[:]))
	}
}

func (l *logJson) Set(key string, value interface{}) {
	l.Keys[key] = value
}

func init() {
	Debug = &logJson{Logger: log.New(io.MultiWriter(os.Stdout), "", 0), Keys: map[string]interface{}{}}
	Debug.Set("level", "debug")

	Warning = &logJson{Logger: log.New(io.MultiWriter(os.Stdout), "", 0), Keys: map[string]interface{}{}}
	Warning.Set("level", "warning")

	Info = &logJson{Logger: log.New(io.MultiWriter(os.Stdout), "", 0), Keys: map[string]interface{}{}}
	Info.Set("level", "info")

	Error = &logJson{Logger: log.New(io.MultiWriter(os.Stderr), "", 0), Keys: map[string]interface{}{}}
	Error.Set("level", "error")

	Fatal = &logJson{Logger: log.New(io.MultiWriter(os.Stderr), "", 0), Keys: map[string]interface{}{}}
	Fatal.Set("level", "fatal")
}
