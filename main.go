package main

import (
	"context"
	"log"
	"net"

	subscriptionpb "github.com/grpc-example/subscription"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type MembershipGRPCServer struct {
	subscriptionpb.UnsafeMembershipServiceServer
}

func (m *MembershipGRPCServer) GetMembership(ctx context.Context, param *subscriptionpb.Membership) (*subscriptionpb.Membership, error) {
	// get data from DB
	mem := subscriptionpb.Membership{
		Guid: "guid",
		Name: "David",
	}

	return &subscriptionpb.Membership{
		Guid: mem.Guid,
		Name: mem.Name,
	}, nil
}

func NewMembershipGRPCServer() *MembershipGRPCServer {
	return &MembershipGRPCServer{}
}

func main() {
	lis, err := net.Listen("tcp", ":2345")
	if err != nil {
		log.Println(err)
	}

	s := grpc.NewServer()
	reflection.Register(s)

	member := NewMembershipGRPCServer()

	subscriptionpb.RegisterMembershipServiceServer(s, member)

	log.Println(`service grpc ready`)
	err = s.Serve(lis)
	if err != nil {
		log.Println(err)
	}
}

// step create grpc server
// - create struct MembershipGRPCServer dan NewMembershipGRPCServer
// - create UnsafeMembershipServiceServer
// - create implementation of GetMembership
