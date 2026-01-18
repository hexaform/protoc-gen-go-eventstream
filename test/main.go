package main

import pb "pkg.hexaform.dev/protogen/eventstream/test/gen/go"

func main() {
	e := &pb.UserRegistered{
		UserId:   357,
		Username: "StElijah",
	}

	m := pb.NewUserEvent(e)

	m.Un
}
