package handler

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"userop_srv/global"
	"userop_srv/model"
	"userop_srv/proto"
)

func (*UserOpServer) MessageList(ctx context.Context, req *proto.MessageRequest) (*proto.MessageListResponse, error) {
	var rsp proto.MessageListResponse
	var messages []model.LeavingMessages
	var messageList  []*proto.MessageResponse

	result := global.DB.Where(&model.LeavingMessages{User:req.UserId}).Find(&messages)
	rsp.Total = int32(result.RowsAffected)

	for _, message := range messages {
		messageList = append(messageList, &proto.MessageResponse{
			Id:          message.ID,
			UserId:      message.User,
			MessageType: message.MessageType,
			Subject:     message.Subject,
			Message:     message.Message,
			File:        message.File,
		})
	}

	rsp.Data = messageList
	return &rsp, nil
}

func (*UserOpServer) CreateMessage(ctx context.Context, req *proto.MessageRequest) (*proto.MessageResponse, error) {
	var message model.LeavingMessages

	message.User = req.UserId
	message.MessageType = req.MessageType
	message.Subject = req.Subject
	message.Message = req.Message
	message.File = req.File

	global.DB.Save(&message)

	return &proto.MessageResponse{Id:message.ID}, nil
}

func (*UserOpServer) DeleteMessage(ctx context.Context, req *proto.MessageRequest) (*emptypb.Empty, error) {
	if result := global.DB.Where("id=? and user=?", req.Id, req.UserId).Delete(&model.LeavingMessages{}); result.RowsAffected == 0{
		return nil, status.Errorf(codes.NotFound, "用户留言不存在")
	}
	return &emptypb.Empty{}, nil
}