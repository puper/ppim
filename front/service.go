package front

import (
	proto "github.com/puper/ppim/proto/front"
)

type Service struct {
}

func (this *Service) Auth(args *proto.AuthArgs, reply *proto.AuthReply) error {
	storage.Put()
	return nil
}

func (this *Service) Quit(args *proto.QuitArgs, reply *proto.QuitReply) error {
	return nil
}
