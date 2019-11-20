package controllers

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"play/etl/services"
)

type ETL struct{}

func (etl *ETL) Run() {
	person := new(services.Person) // create instance
	rows, err := person.GetAll()
	if err != nil {
		log.Panic(err)
	}

	workerCount := 3
	users := make([]*services.Person, 0)
	ch := make(chan []*services.Person, workerCount)
	ctx, cancel := context.WithCancel(context.Background())

	var wg sync.WaitGroup
	wg.Add(3)
	for i := 0; i < workerCount; i++ {
		go process(ctx, ch, &wg)
	}

	for rows.Next() {
		// var p *services.Person
		p := services.Person{}
		rows.StructScan(&p)
		users = append(users, &p)

		// write to channel in batch
		if len(users) == 5 {
			ch <- users
			users = nil
		}
	}

	// flush
	ch <- users
	cancel()

	wg.Wait()

	fmt.Println("game over.")
}

func process(ctx context.Context, ch <-chan []*services.Person, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case data := <-ch:
			for _, d := range data {
				fmt.Printf("data: %v %v %v\n", d.ID, d.Phone, d.NickName)
			}
		case <-time.After(time.Second):
			fmt.Println("Break time")
		case <-ctx.Done():
			return
		}
	}
}
