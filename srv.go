package main

import (
	"fmt"
	"context"
	"net"
	"google.golang.org/grpc"
	"agrpc"
	"strings"
	"log"
)
const (
	port=":5051"
)
type server struct{
	savedCustomers []*agrpc.CustomerRequest
}
func (s *server) CreateCustomer(ctx context.Context,in *agrpc.CustomerRequest)(*agrpc.CustomerResponse ,error){
	s.savedCustomers=append(s.savedCustomers,in)
	return &agrpc.CustomerResponse{Id:in.Id,Success:true},nil

}
func(s *server) GetCustomers(filter *agrpc.CustomerFilter,stream agrpc.Customer_GetCustomersServer)error{
	for _,customer:=range s.savedCustomers{

		if filter.Keyword != "" {
			if !strings.Contains(customer.Name, filter.Keyword) {
				continue
			}
		}
		if err := stream.Send(customer); err != nil {
			return err
		}
	}
	return nil

}
func main(){
	var objServer server
	fmt.Println("hello by customer server")
	lis,err:=net.Listen("tcp",port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s:=grpc.NewServer()
	agrpc.RegisterCustomerServer(s,&objServer)
	s.Serve(lis)
	
}