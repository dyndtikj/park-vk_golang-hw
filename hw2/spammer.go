package main

import (
	"log"
	"sort"
	"strconv"
	"strings"
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
	for val := range in {
		email := val.(string)
		out <- GetUser(email)
	}
}

func SelectMessages(in, out chan interface{}) {
	// TODO batches , and unique users
	for val := range in {
		usr := val.(User)
		msg, err := GetMessages(usr)
		if err != nil {
			log.Fatal(err)
		}
		out <- msg
	}
}

func CheckSpam(in, out chan interface{}) {
	for val := range in {
		msgIds := val.([]MsgID)
		for _, m := range msgIds {
			has, err := HasSpam(m)
			if err != nil {
				log.Fatal(err)
			}
			out <- MsgData{
				ID:      m,
				HasSpam: has,
			}

		}
	}
}

func CombineResults(in, out chan interface{}) {
	var results []string
	for val := range in {
		msgData := val.(MsgData)
		resStr := strconv.FormatBool(msgData.HasSpam) + strconv.Itoa(int(msgData.ID))
		results = append(results, resStr)
	}

	sort.Strings(results)
	separator := 0
	for idx, str := range results {
		sl := strings.Split(str, " ")
		if sl[0] == strconv.FormatBool(true) {
			separator = idx
			break
		}
	}
	result := results[separator:]
	result = append(result, results[:separator]...)

	out <- strings.Join(results, "\n")
}
