package miner

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func miner(
	ctx context.Context, 
	wg *sync.WaitGroup, 
	transferPoint chan<- int, 
	n int, power int,
	) {
	defer wg.Done();
	for {
		select{
		case <-ctx.Done(): 
		fmt.Println("Я шахтер № ", n, " мой рабочий день окончен!"); 
		return;
		default:
			fmt.Println("Я шахтер № ", n, " начал добывать уголь!");
			time.Sleep(1 * time.Second);
			fmt.Println("Я шахтер № ", n, " добыл уголь: ", power);
			
			transferPoint <- power;
			fmt.Println("Я шахтер № ", n, " передал уголь: ", power);
		}

	}
}

func MinerPool(ctx context.Context, minerCount int) <-chan int{

	coalTransferPoint := make(chan int);
	wg := &sync.WaitGroup{}

	for i := 1; i <= minerCount; i++ {
		wg.Add(1);
		go miner(ctx,wg, coalTransferPoint, i, i*10);
	}
	go func(){
	  wg.Wait();
	  close(coalTransferPoint);
	}()
	return coalTransferPoint;
}