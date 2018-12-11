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
	Guard     int
	Timestamp time.Time
	Action    string
	Text      string
}

type Guard struct {
	Minutes          [60]int
	MinutesAsleep    int
	MostCommonMinute int
	MostCommonValue  int
}

type Guards struct {
	List map[int]Guard
}

type By func(log1, log2 *LogEntry) bool

func (by By) Sort(logs []LogEntry) {
	logsorter := &logSorter{
		logs: logs,
		by:   by,
	}
	sort.Sort(logsorter)
}

type logSorter struct {
	logs []LogEntry
	by   func(log1, log2 *LogEntry) bool
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

func GenerateLogEntry(entry *string) (logentry LogEntry) {
	logentry.Text = *entry
	timeStampRegex := regexp.MustCompile("[\\d]{4}-[\\d]{2}-[\\d]{2} [\\d]{2}:[\\d]{2}")
	timestamp := timeStampRegex.FindString(logentry.Text)
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

func main() {
	inputList := []LogEntry{}
	reader := bufio.NewReader(os.Stdin)
	timestart := time.Now().UnixNano()
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
	fmt.Println(time.Now().UnixNano() - timestart)
	timeStart := time.Now().UnixNano()
	//	timesort := func(log1, log2 *LogEntry) bool {
	//		return log1.Timestamp.Before(log2.Timestamp)
	//	}
	guardList := new(Guards)
	guardList.List = make(map[int]Guard, 0)
	//	By(timesort).Sort(inputList)
	currentGuard := 0
	for i, entry := range inputList {
		if entry.Guard == 0 {
			inputList[i].Guard = currentGuard
		} else {
			guardList.List[entry.Guard] = *new(Guard)
			currentGuard = entry.Guard
		}
	}
	currentGuard = 0
	currentGuard = inputList[0].Guard
	state := 0 // awake
	prevMinute := 0
	for _, event := range inputList {
		currentMinute := event.Timestamp.Minute()
		minutesBetween := MinutesBetween(prevMinute, currentMinute)
		guard := guardList.List[currentGuard]
		for _, minute := range minutesBetween {
			guard.Minutes[minute] += state
		}
		guardList.List[currentGuard] = guard
		prevMinute = currentMinute
		switch action := event.Action; action {
		case "falls asleep":
			state = 1
		case "wakes up":
			state = 0
		case "begins shift":
			currentGuard = event.Guard
			state = 0
		}
	}
	for guardNum, guard := range guardList.List {
		sum := 0
		mostCommon := 0
		mostCommonVal := 0
		for i, minute := range guard.Minutes {
			if minute > mostCommonVal {
				mostCommon = i
				mostCommonVal = minute
			}
			sum += minute
		}
		guard.MinutesAsleep = sum
		guard.MostCommonMinute = mostCommon
		guard.MostCommonValue = mostCommonVal
		guardList.List[guardNum] = guard
	}

	maxTime := 0
	maxGuard := 0
	for guard, data := range guardList.List {
		if data.MinutesAsleep > maxTime {
			maxGuard = guard
			maxTime = data.MinutesAsleep
		}
	}
	fmt.Println(guardList.List[maxGuard].MostCommonMinute * maxGuard)

	mostSleptGuard := 0
	mostSleptMinute := 0
	mostSleptValue := 0
	for guard, data := range guardList.List {
		for i, value := range data.Minutes {
			if value > mostSleptValue {
				mostSleptGuard = guard
				mostSleptValue = value
				mostSleptMinute = i
			}
		}
	}
	fmt.Println(mostSleptGuard * mostSleptMinute)
	fmt.Println(time.Now().UnixNano() - timeStart)

}

func MinutesBetween(time1, time2 int) (times []int) {
	for {
		time1 = time1 % 60
		if time1 != time2 {
			times = append(times, time1)
			time1 += 1
		} else {
			break
		}
	}
	return
}
