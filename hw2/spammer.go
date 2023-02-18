package main

import (
	"log"
	"sort"
	"strconv"
	"sync"
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
			close(out)
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

func GetMsgWorker(butchUsr []User, out chan interface{}) {
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
	for val := range in {
		usr := val.(User)
		butchUsr = append(butchUsr, usr)
		if len(butchUsr) == GetMessagesMaxUsersBatch {
			GetMsgWorker(butchUsr, out)
			butchUsr = butchUsr[:0]
		}
	}
	if len(butchUsr) > 0 {
		GetMsgWorker(butchUsr, out)
	}
}

func CheckSpam(in, out chan interface{}) {
	for val := range in {
		msgId := val.(MsgID)
		has, err := HasSpam(msgId)
		if err != nil {
			log.Fatal(err)
		}
		out <- MsgData{
			ID:      msgId,
			HasSpam: has,
		}
	}
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
