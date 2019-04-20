package main

import (
	"context"

	"github.com/IBM/cloudland/web/sca/clients"
	"github.com/IBM/cloudland/web/sca/pkgs"
	"google.golang.org/grpc"
)

func init() {
	RootCmd.AddCommand(pkgs.Commands(context.Background,
		func() *grpc.ClientConn {
			return clients.GetClientConn("admin")
		})...)
}
