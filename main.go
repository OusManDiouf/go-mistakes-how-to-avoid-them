package main

import (
	"errors"
	"fmt"
	"strings"
)

type LogStore struct {
	UUID string
}

func main() {
	ls := LogStore{}
	logLine := "HJYU-JKJLK-KI975-9LK780-JLKJ0-889JKJ  2022-10-25 10:19:45 [GET] /healthcheck"
	err := ls.handleLog(logLine)
	if err != nil {
		return
	}
	fmt.Println(ls)
}

func (s *LogStore) handleLog(logLine string) error {
	if len(logLine) < 36 {
		return errors.New("log is not correctly formated")
	}
	// we mentionned that logLine message can be quite heavy
	// [:36] will create a new string referencing the same backing array as logLine !
	// therefore, each uuid that we store in mem will contain not only 36bytes,
	// but the number of bytes of the logLine
	//uuid := logLine[:36]

	// uuid := string([]byte(logLine[:36]))
	// go 1.18
	uuid := strings.Clone(logLine[:36])
	s.store(uuid)
	return nil
}

func (s *LogStore) store(uuid string) {
	s.UUID = uuid
}
