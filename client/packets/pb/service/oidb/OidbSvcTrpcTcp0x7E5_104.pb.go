// Code generated by protoc-gen-golite. DO NOT EDIT.
// source: pb/service/oidb/OidbSvcTrpcTcp0x7E5_104.proto

package oidb

import (
	proto "github.com/RomiChan/protobuf/proto"
)

// Friend Likes
type OidbSvcTrpcTcp0X7E5_104 struct {
	TargetUid proto.Option[string] `protobuf:"bytes,11,opt"`
	Source    uint32               `protobuf:"varint,12,opt"` // 71
	Count     uint32               `protobuf:"varint,13,opt"` // 1
	_         [0]func()
}
