package main

import (
	"log"
	"sort"
	"strconv"
	"sync"
	"sync/atomic"
)

func RunPipeline(cmds ...cmd) {
	wg := &sync.WaitGroup{}
	in := make(chan interface{})

	for _, c := range cmds {
		wg.Add(1)
		out := make(chan interface{})

		go func(c cmd, in, out chan interface{}) {
			c(in, out)
			wg.Done()
			close(out)
		}(c, in, out)

		in = out
	}

	wg.Wait()
}

func SelectUsers(in, out chan interface{}) {
	set := sync.Map{}
	wg := sync.WaitGroup{}
	for val := range in {
		email := val.(string)
		usrChan := make(chan User)

		go func(email string, out chan User) {
			out <- GetUser(email)
		}(email, usrChan)

		wg.Add(1)
		go func() {
			user := <-usrChan
			if _, ok := set.Load(user.Email); !ok {
				set.Store(user.Email, true)
				out <- user
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

func GetMsgWorker(butchUsr []User, out chan interface{}, wg *sync.WaitGroup) {
	defer wg.Done()
	messages, err := GetMessages(butchUsr...)
	if err != nil {
		log.Fatal(err)
	}
	for _, msg := range messages {
		out <- msg
	}
}

func SelectMessages(in, out chan interface{}) {
	butchUsr := make([]User, 0, GetMessagesMaxUsersBatch)
	wg := sync.WaitGroup{}
	for val := range in {
		usr := val.(User)
		butchUsr = append(butchUsr, usr)
		if len(butchUsr) == GetMessagesMaxUsersBatch {
			wg.Add(1)
			res := make([]User, len(butchUsr))
			copy(res, butchUsr)
			go GetMsgWorker(res, out, &wg)
			butchUsr = butchUsr[:0]
		}
	}
	if len(butchUsr) > 0 {
		wg.Add(1)
		go GetMsgWorker(butchUsr, out, &wg)
	}
	wg.Wait()
}

var checkSpamCounter int32
var mutex sync.Mutex

func checkBrut() bool {
	mutex.Lock()
	b := checkSpamCounter < int32(HasSpamMaxAsyncRequests)
	mutex.Unlock()
	return b
}

func CheckSpam(in, out chan interface{}) {
	wg := sync.WaitGroup{}
	for val := range in {
		msgId := val.(MsgID)
		for {
			if checkBrut() {
				atomic.AddInt32(&checkSpamCounter, 1)
				wg.Add(1)
				go func(id MsgID) {
					has, err := HasSpam(id)
					if err != nil {
						log.Fatal(err)
					}
					out <- MsgData{
						ID:      msgId,
						HasSpam: has,
					}
					atomic.AddInt32(&checkSpamCounter, -1)
					wg.Done()
				}(msgId)
				break
			}
		}
	}
	wg.Wait()
}

func CombineResults(in, out chan interface{}) {
	var data []MsgData
	for val := range in {
		data = append(data, val.(MsgData))
	}
	sort.Slice(data, func(i, j int) bool {
		if data[i].HasSpam == data[j].HasSpam {
			return data[i].ID < data[j].ID
		}
		// i < j
		return data[i].HasSpam && !data[j].HasSpam
	})

	for _, msg := range data {
		resStr := strconv.FormatBool(msg.HasSpam) + " " + strconv.FormatUint(uint64(msg.ID), 10)
		out <- resStr
	}
}
