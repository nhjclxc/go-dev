package main


import (
	"context"
	"go-micro.dev/v5"
)

type Request struct {
	Name string `json:"name"`
}

type Response struct {
	Message string `json:"message"`
}

type Say struct{}

func (h *Say) Hello(ctx context.Context, req *Request, rsp *Response) error {
	rsp.Message = "Hello " + req.Name
	return nil
}

func main() {
	// create the service
	service := micro.NewService(
		micro.Address(":8080"),
		micro.Name("helloworld"),
	)

	// register handler
	service.Handle(new(Say))

	// run the service
	service.Run()
}