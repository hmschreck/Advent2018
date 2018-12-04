package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"time"
)

type LogEntry struct {
	Guard int
	Timestamp time.Time
	Action string
	Text string
}

type By func(log1, log2 *LogEntry) bool

func (by By) Sort(logs []LogEntry) {
	logsorter := &logSorter{
		logs: logs,
		by: by,
	}
	sort.Sort(logsorter)
}

type logSorter struct {
	logs []LogEntry
	by func(log1, log2 *LogEntry) bool
}

func (s *logSorter) Swap(i, j int) {
	s.logs[i], s.logs[j] = s.logs[j], s.logs[i]
}

func (s *logSorter) Len() int {
	return len(s.logs)
}

func (s *logSorter) Less(i, j int) bool {
	return s.by(&s.logs[i], &s.logs[j])
}

func GenerateLogEntry(entry *string) (logentry LogEntry){
	logentry.Text = *entry
	timeStampRegex := regexp.MustCompile("[\\d]{4}-[\\d]{2}-[\\d]{2} [\\d]{2}:[\\d]{2}")
	timestamp := timeStampRegex.FindString(logentry.Text)
	fmt.Println(timestamp)
	time, err := time.Parse("2006-01-02 15:04", timestamp)
	if err != nil {
		panic(err.Error())
	}
	logentry.Timestamp = time
	guardRegex := regexp.MustCompile("Guard #([\\d]{3,4})")
	guard := guardRegex.FindStringSubmatch(logentry.Text)
	if guard != nil {
		logentry.Guard, _ = strconv.Atoi(guard[1])
	}
	actionRegex := regexp.MustCompile("begins shift|falls asleep|wakes up")
	logentry.Action = actionRegex.FindString(logentry.Text)

	return
}


func main(){
inputList := []LogEntry{}
reader := bufio.NewReader(os.Stdin)
for {
		text, err := reader.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}
		newEntry := GenerateLogEntry(&text)
		inputList = append(inputList, newEntry)

	}
	timesort := func(log1, log2 *LogEntry) bool {
		return log1.Timestamp.Before(log2.Timestamp)
	}
	By(timesort).Sort(inputList)
	currentGuard := 0
	for i, entry := range inputList {
		if entry.Guard == 0 {
			inputList[i].Guard = currentGuard
		} else {
			currentGuard = entry.Guard
		}
	}
}