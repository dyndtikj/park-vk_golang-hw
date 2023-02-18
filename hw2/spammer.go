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
	set := make(map[string]bool)
	for val := range in {
		email := val.(string)
		user := GetUser(email)
		if _, ok := set[user.Email]; !ok {
			out <- user
			set[user.Email] = true
		}
	}
}

func SelectMessages(in, out chan interface{}) {
	// TODO batches
	for val := range in {
		usr := val.(User)
		messages, err := GetMessages(usr)
		if err != nil {
			log.Fatal(err)
		}
		for _, msg := range messages {
			out <- msg
		}
	}
}

func CheckSpam(in, out chan interface{}) {
	// TODO antibrut
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
