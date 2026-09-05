package main

import (
	"concurrency/miner"
	"concurrency/postman"
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)



func main() {
	var coal atomic.Int64;
	mtx := sync.Mutex{}
	var mails []string;
	minerContext, minerCancel := context.WithCancel(context.Background())
	postmanContext, postmanCancel := context.WithCancel(context.Background())

	go func(){
	  time.Sleep(3 * time.Second);
	  minerCancel();
	}()
	go func(){
	  time.Sleep(6 * time.Second);
	  postmanCancel();
	}()	

	coalTransferPoint := miner.MinerPool(minerContext, 1000)
	mailTransferPoint := postman.PostmanPool(postmanContext, 1000)

	wg := &sync.WaitGroup{}

	wg.Add(1);
	go func(){
	  defer wg.Done();

	  for v := range coalTransferPoint{
		coal.Add(int64(v));
	  }
	}()
	wg.Add(1);
	go func(){
	  defer wg.Done();

	  for v := range mailTransferPoint{
		mtx.Lock();
		mails = append(mails, v);
		mtx.Unlock();
	  }
	}()	

	wg.Wait();

	fmt.Println("Сумма угля ", coal.Load());
	mtx.Lock();
	fmt.Println("Письма ", len(mails));
	mtx.Unlock()

	
}
