package main

import (
	"fmt"
	"github.com/bxcodec/faker/v4"
	"github.com/filatkinen/cassandra/internal"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"sync/atomic"
	"syscall"
	"time"
)

const (
	workers = 20
	records = 100_000
)

func main() {
	defer internal.Session.Close()
	var wg sync.WaitGroup
	var count int32
	closeChan := make(chan struct{})
	var lock sync.Mutex

	mStudentID, err := internal.MaxStudentID()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(mStudentID)

	timeStart := time.Now()
	atomic.AddInt32(&count, int32(mStudentID))

	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			var student internal.Student
			s1 := rand.NewSource(time.Now().UnixNano())
			r1 := rand.New(s1)
			for {
				student.ID = int(atomic.AddInt32(&count, 1))
				//count++
				//student.ID = int(count)
				student.Age = r1.Intn(8) + 18
				lock.Lock()
				student.Firstname = faker.FirstName()
				student.Lastname = faker.LastName()
				lock.Unlock()
				select {
				case <-closeChan:
					return
				default:
				}
				if student.ID > records+mStudentID {
					return
				}
				err = internal.AddStudent(student)
				if err != nil {
					fmt.Println(err)
					return
				}
			}
		}()
	}
	signalStop := make(chan os.Signal, 1)
	signal.Notify(signalStop, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-signalStop
		close(closeChan)
	}()

	wg.Wait()
	fmt.Println(time.Since(timeStart))
}
