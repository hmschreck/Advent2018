package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
	"time"
)

type Step struct {
	Letter    string
	Prev      []*Step
	Next      []*Step
	Completed bool
	Index     int
	Queued    bool
	Duration  int
	Locked    bool
}

type Worker struct {
	WorkingOn *Step
}

type WorkerPool struct {
	Workers []*Worker
}

var CAPITALA = int([]rune("A")[0])

func (worker *Worker) Tick(queue *PriorityQueue) {
	if worker.WorkingOn == nil || worker.WorkingOn.Duration == 0 || worker.WorkingOn.Completed == true {
		if queue.Len() > 0 {
			worker.WorkingOn = heap.Pop(queue).(*Step)
		} else {
			worker.WorkingOn = nil
			return
		}
	}
	worker.WorkingOn.Duration -= 1
	if worker.WorkingOn.Duration == 0 {
		worker.WorkingOn.Completed = true
	}
}

func CreateWorkerPool(len int) (pool WorkerPool) {
	for i := 0; i < len; i++ {
		pool.Workers = append(pool.Workers, &Worker{})
	}
	return
}

func (pool *WorkerPool) Tick(queue *PriorityQueue) {
	for _, worker := range pool.Workers {
		worker.Tick(queue)
	}
}

type PriorityQueue []*Step

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Letter < pq[j].Letter
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	step := x.(*Step)
	step.Index = n
	*pq = append(*pq, step)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	step := old[n-1]
	step.Index = n - 1
	*pq = old[0 : n-1]
	return step
}

func (pq *PriorityQueue) Update(step *Step) {
	heap.Fix(pq, step.Index)
}

var StepsList = make(map[string]*Step, 0)
var StepOrder = []string{}

func main() {
	//timeStart := time.Now().UnixNano()
	stepsRegex := regexp.MustCompile("Step (.) must be finished before step (.) can begin.")
	inputList := []string{}
	reader := bufio.NewReader(os.Stdin)
	for {
		text, err := reader.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}
		inputList = append(inputList, text)
	}
	for _, line := range inputList {
		steps := stepsRegex.FindStringSubmatch(line)
		StepsList[steps[1]] = &Step{Letter: steps[1], Duration: int([]rune(steps[1])[0]) - CAPITALA + 1 + 60}
		StepsList[steps[2]] = &Step{Letter: steps[2], Duration: int([]rune(steps[2])[0]) - CAPITALA + 1 + 60}
	}
	for _, line := range inputList {
		steps := stepsRegex.FindStringSubmatch(line)
		StepsList[steps[1]].Next = append(StepsList[steps[1]].Next, StepsList[steps[2]])
		StepsList[steps[2]].Prev = append(StepsList[steps[2]].Prev, StepsList[steps[1]])
	}

	// Find start
	firstStep := Step{Letter: "", Duration: 0}
	for _, step := range StepsList {
		if len(step.Prev) == 0 {
			firstStep.Next = append(firstStep.Next, step)
		}
	}
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	StepThrough(&firstStep, &pq)
	fmt.Println(strings.Join(StepOrder, ""))

	for _, step := range StepsList {
		step.Completed = false
	}
	pq2 := make(PriorityQueue, 0)
	heap.Init(&pq2)
	for _, step := range StepsList {
		step.Completed = false
		step.Queued = false
	}
	pool := CreateWorkerPool(5)
	tick := 0
	startTimePart2 := time.Now().UnixNano()
OuterLoop:
	for {
		allCompleted := true
		for _, step := range StepsList {
			if step.Completed == false && step.Queued == false {
				prevDone := true
				for _, prev := range step.Prev {
					if prev.Completed == false {
						prevDone = false
					}
				}
				if prevDone == false {
					continue
				}
				step.Queued = true
				heap.Push(&pq2, step)
			}
			if step.Completed == false {
				allCompleted = false
			}
		}
		if allCompleted {
			break OuterLoop
		}
		tick += 1
		pool.Tick(&pq2)
	}
	fmt.Println(tick)
	fmt.Println("Part two completed in ", float64((time.Now().UnixNano()-startTimePart2))/1000000.0, " milliseconds")
}

func StepThrough(step *Step, queue *PriorityQueue) {
	StepOrder = append(StepOrder, step.Letter)
	step.Completed = true
	for _, nextStep := range step.Next {
		allCompleted := true
		for _, prev := range nextStep.Prev {
			if prev.Completed == false {
				allCompleted = false
			}
		}
		if allCompleted == true {
			if nextStep.Queued == false {
				nextStep.Queued = true
				heap.Push(queue, nextStep)

			}
		}

	}
	if queue.Len() > 0 {
		NewStep := heap.Pop(queue).(*Step)
		StepThrough(NewStep, queue)
	}
}
