package services

import (
	"context"
	"fmt"
	"io"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type CalculatorService interface {
	Hello(name string) error
	Fibonacci(n uint32) error
	Average(numbers ...float64) error
	Sum(numbers ...int32) error
}

type calculatorService struct {
	calculatorClient CalculatorClient
}

func NewCalculatorService(calculatorClient CalculatorClient) CalculatorService {
	return calculatorService{calculatorClient}
}

func (base calculatorService) Hello(name string) error {
	req := HelloRequest{
		Name:        name,
		CreatedDate: timestamppb.Now(),
	}
	res, err := base.calculatorClient.Hello(context.Background(), &req)
	if err != nil {
		return err
	}
	fmt.Println("Service : Hello")
	fmt.Println("Request : ", req.Name)
	fmt.Println("Response : ", res.Result)
	return nil
}

func (base calculatorService) Fibonacci(n uint32) error {
	req := FibonacciRequest{
		N: n,
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	stream, err := base.calculatorClient.Fibonacci(ctx, &req)
	if err != nil {
		return err
	}
	fmt.Println("Service : Fibonacci")
	fmt.Println("Request : ", req.N)
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		fmt.Println("Response : ", res.Result)
	}
	return nil
}

func (base calculatorService) Average(numbers ...float64) error {
	strame, err := base.calculatorClient.Average(context.Background())
	if err != nil {
		return err
	}
	fmt.Println("Service : Average")
	for _, number := range numbers {
		req := AverageRequest{
			Number: number,
		}
		strame.Send(&req)
		fmt.Println("Request : ", req.Number)
		time.Sleep(time.Second)
	}

	res, err := strame.CloseAndRecv()
	if err != nil {
		return err
	}
	fmt.Println("Response : ", res.Result)
	return nil
}

func (base calculatorService) Sum(numbers ...int32) error {
	straem, err := base.calculatorClient.Sum(context.Background())
	if err != nil {
		return err
	}
	fmt.Println("Service : Sum")
	go func() {
		for _, number := range numbers {
			req := SumRequest{
				Number: number,
			}
			straem.Send(&req)
			fmt.Println("Request : ", req.Number)
			time.Sleep(time.Second * 2)
		}
		straem.CloseSend()
	}()
	done := make(chan bool)
	errs := make(chan error)
	go func() {
		for {
			res, err := straem.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				errs <- err
			}
			fmt.Println("Response : ", res.Result)
		}
		done <- true
	}()

	select {
	case <-done:
		return nil
	case err := <-errs:
		return err
	}
}
