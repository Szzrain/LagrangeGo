package client

import (
	message2 "github.com/LagrangeDev/LagrangeGo/message"
	"github.com/LagrangeDev/LagrangeGo/packets/pb/action"
	"github.com/LagrangeDev/LagrangeGo/packets/pb/message"
	"github.com/LagrangeDev/LagrangeGo/utils"
	"github.com/RomiChan/protobuf/proto"
)

func (c *QQClient) SendRawMessage(route *message.RoutingHead, body *message.MessageBody) (resp *action.SendMessageResponse, err error) {
	seq := c.sig.Sequence + 1
	msg := &message.Message{
		RoutingHead: route,
		ContentHead: &message.ContentHead{
			Type:    1,
			SubType: proto.Some(uint32(0)),
			DivSeq:  proto.Some(uint32(0)),
		},
		Body: body,
		Seq:  proto.Some(uint32(seq)),
		Rand: proto.Some(utils.RandU32()),
	}
	// grp_id not null
	if (route.Grp != nil && route.Grp.GroupCode.IsSome()) || (route.GrpTmp != nil && route.GrpTmp.GroupUin.IsSome()) {
		msg.Ctrl = &message.MessageControl{MsgFlag: int32(utils.TimeStamp())}
	}

	data, err := proto.Marshal(msg)
	if err != nil {
		return
	}

	packet, err := c.SendUniPacketAndAwait("MessageSvc.PbSendMsg", data)
	if err != nil {
		return
	}
	err = proto.Unmarshal(packet.Data, resp)
	if err != nil {
		return
	}
	return
}

func (c *QQClient) SendGroupMessage(groupID uint32, msg []message2.IMessageElement) (resp *action.SendMessageResponse, err error) {
	body := message2.BuildMessageElements(msg)
	route := &message.RoutingHead{
		Grp: &message.Grp{GroupCode: proto.Some(groupID)},
	}
	return c.SendRawMessage(route, body)
}

func (c *QQClient) SendPrivateMessage(uin uint32, msg []message2.IMessageElement) (resp *action.SendMessageResponse, err error) {
	body := message2.BuildMessageElements(msg)
	route := &message.RoutingHead{
		C2C: &message.C2C{
			Uid: proto.Some(c.GetUid(uin)),
		},
	}
	return c.SendRawMessage(route, body)
}

func (c *QQClient) SendTempMessage(groupID uint32, uin uint32, msg []message2.IMessageElement) (resp *action.SendMessageResponse, err error) {
	body := message2.BuildMessageElements(msg)
	route := &message.RoutingHead{
		GrpTmp: &message.GrpTmp{
			GroupUin: proto.Some(groupID),
			ToUin:    proto.Some(uin),
		},
	}
	return c.SendRawMessage(route, body)
}