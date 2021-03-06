// Code generated by protoc-gen-go. DO NOT EDIT.
// source: trillian_map_api.proto

package trillian

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import google_protobuf "github.com/golang/protobuf/ptypes/any"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// MapLeaf represents the data behind Map leaves.
type MapLeaf struct {
	// index is the location of this leaf.
	// All indexes for a given Map must contain a constant number of bits.
	Index []byte `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	// leaf_hash is the tree hash of leaf_value.  This does not need to be set
	// on SetMapLeavesRequest; the server will fill it in.
	LeafHash []byte `protobuf:"bytes,2,opt,name=leaf_hash,json=leafHash,proto3" json:"leaf_hash,omitempty"`
	// leaf_value is the data the tree commits to.
	LeafValue []byte `protobuf:"bytes,3,opt,name=leaf_value,json=leafValue,proto3" json:"leaf_value,omitempty"`
	// extra_data holds related contextual data, but is not covered by any hash.
	ExtraData []byte `protobuf:"bytes,4,opt,name=extra_data,json=extraData,proto3" json:"extra_data,omitempty"`
}

func (m *MapLeaf) Reset()                    { *m = MapLeaf{} }
func (m *MapLeaf) String() string            { return proto.CompactTextString(m) }
func (*MapLeaf) ProtoMessage()               {}
func (*MapLeaf) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *MapLeaf) GetIndex() []byte {
	if m != nil {
		return m.Index
	}
	return nil
}

func (m *MapLeaf) GetLeafHash() []byte {
	if m != nil {
		return m.LeafHash
	}
	return nil
}

func (m *MapLeaf) GetLeafValue() []byte {
	if m != nil {
		return m.LeafValue
	}
	return nil
}

func (m *MapLeaf) GetExtraData() []byte {
	if m != nil {
		return m.ExtraData
	}
	return nil
}

type MapLeafInclusion struct {
	Leaf      *MapLeaf `protobuf:"bytes,1,opt,name=leaf" json:"leaf,omitempty"`
	Inclusion [][]byte `protobuf:"bytes,2,rep,name=inclusion,proto3" json:"inclusion,omitempty"`
}

func (m *MapLeafInclusion) Reset()                    { *m = MapLeafInclusion{} }
func (m *MapLeafInclusion) String() string            { return proto.CompactTextString(m) }
func (*MapLeafInclusion) ProtoMessage()               {}
func (*MapLeafInclusion) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *MapLeafInclusion) GetLeaf() *MapLeaf {
	if m != nil {
		return m.Leaf
	}
	return nil
}

func (m *MapLeafInclusion) GetInclusion() [][]byte {
	if m != nil {
		return m.Inclusion
	}
	return nil
}

type GetMapLeavesRequest struct {
	MapId int64    `protobuf:"varint,1,opt,name=map_id,json=mapId" json:"map_id,omitempty"`
	Index [][]byte `protobuf:"bytes,2,rep,name=index,proto3" json:"index,omitempty"`
}

func (m *GetMapLeavesRequest) Reset()                    { *m = GetMapLeavesRequest{} }
func (m *GetMapLeavesRequest) String() string            { return proto.CompactTextString(m) }
func (*GetMapLeavesRequest) ProtoMessage()               {}
func (*GetMapLeavesRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *GetMapLeavesRequest) GetMapId() int64 {
	if m != nil {
		return m.MapId
	}
	return 0
}

func (m *GetMapLeavesRequest) GetIndex() [][]byte {
	if m != nil {
		return m.Index
	}
	return nil
}

// This message replaces the current implementation of GetMapLeavesRequest
// with the difference that revision must be >=0.
type GetMapLeavesByRevisionRequest struct {
	MapId int64    `protobuf:"varint,1,opt,name=map_id,json=mapId" json:"map_id,omitempty"`
	Index [][]byte `protobuf:"bytes,2,rep,name=index,proto3" json:"index,omitempty"`
	// revision >= 0.
	Revision int64 `protobuf:"varint,3,opt,name=revision" json:"revision,omitempty"`
}

func (m *GetMapLeavesByRevisionRequest) Reset()                    { *m = GetMapLeavesByRevisionRequest{} }
func (m *GetMapLeavesByRevisionRequest) String() string            { return proto.CompactTextString(m) }
func (*GetMapLeavesByRevisionRequest) ProtoMessage()               {}
func (*GetMapLeavesByRevisionRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *GetMapLeavesByRevisionRequest) GetMapId() int64 {
	if m != nil {
		return m.MapId
	}
	return 0
}

func (m *GetMapLeavesByRevisionRequest) GetIndex() [][]byte {
	if m != nil {
		return m.Index
	}
	return nil
}

func (m *GetMapLeavesByRevisionRequest) GetRevision() int64 {
	if m != nil {
		return m.Revision
	}
	return 0
}

type GetMapLeavesResponse struct {
	MapLeafInclusion []*MapLeafInclusion `protobuf:"bytes,2,rep,name=map_leaf_inclusion,json=mapLeafInclusion" json:"map_leaf_inclusion,omitempty"`
	MapRoot          *SignedMapRoot      `protobuf:"bytes,3,opt,name=map_root,json=mapRoot" json:"map_root,omitempty"`
}

func (m *GetMapLeavesResponse) Reset()                    { *m = GetMapLeavesResponse{} }
func (m *GetMapLeavesResponse) String() string            { return proto.CompactTextString(m) }
func (*GetMapLeavesResponse) ProtoMessage()               {}
func (*GetMapLeavesResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func (m *GetMapLeavesResponse) GetMapLeafInclusion() []*MapLeafInclusion {
	if m != nil {
		return m.MapLeafInclusion
	}
	return nil
}

func (m *GetMapLeavesResponse) GetMapRoot() *SignedMapRoot {
	if m != nil {
		return m.MapRoot
	}
	return nil
}

type SetMapLeavesRequest struct {
	MapId  int64      `protobuf:"varint,1,opt,name=map_id,json=mapId" json:"map_id,omitempty"`
	Leaves []*MapLeaf `protobuf:"bytes,2,rep,name=leaves" json:"leaves,omitempty"`
	// Metadata that the Map should associate with the new Map root after
	// incorporating the leaf changes.  The metadata will be reflected in the
	// Map Root returned in the map's SetLeaves response.
	// Map personalities should use metadata to persist any state needed later
	// to continue mapping from an external data source.
	Metadata *google_protobuf.Any `protobuf:"bytes,4,opt,name=metadata" json:"metadata,omitempty"`
}

func (m *SetMapLeavesRequest) Reset()                    { *m = SetMapLeavesRequest{} }
func (m *SetMapLeavesRequest) String() string            { return proto.CompactTextString(m) }
func (*SetMapLeavesRequest) ProtoMessage()               {}
func (*SetMapLeavesRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{5} }

func (m *SetMapLeavesRequest) GetMapId() int64 {
	if m != nil {
		return m.MapId
	}
	return 0
}

func (m *SetMapLeavesRequest) GetLeaves() []*MapLeaf {
	if m != nil {
		return m.Leaves
	}
	return nil
}

func (m *SetMapLeavesRequest) GetMetadata() *google_protobuf.Any {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type SetMapLeavesResponse struct {
	MapRoot *SignedMapRoot `protobuf:"bytes,2,opt,name=map_root,json=mapRoot" json:"map_root,omitempty"`
}

func (m *SetMapLeavesResponse) Reset()                    { *m = SetMapLeavesResponse{} }
func (m *SetMapLeavesResponse) String() string            { return proto.CompactTextString(m) }
func (*SetMapLeavesResponse) ProtoMessage()               {}
func (*SetMapLeavesResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{6} }

func (m *SetMapLeavesResponse) GetMapRoot() *SignedMapRoot {
	if m != nil {
		return m.MapRoot
	}
	return nil
}

type GetSignedMapRootRequest struct {
	MapId int64 `protobuf:"varint,1,opt,name=map_id,json=mapId" json:"map_id,omitempty"`
}

func (m *GetSignedMapRootRequest) Reset()                    { *m = GetSignedMapRootRequest{} }
func (m *GetSignedMapRootRequest) String() string            { return proto.CompactTextString(m) }
func (*GetSignedMapRootRequest) ProtoMessage()               {}
func (*GetSignedMapRootRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{7} }

func (m *GetSignedMapRootRequest) GetMapId() int64 {
	if m != nil {
		return m.MapId
	}
	return 0
}

type GetSignedMapRootByRevisionRequest struct {
	MapId    int64 `protobuf:"varint,1,opt,name=map_id,json=mapId" json:"map_id,omitempty"`
	Revision int64 `protobuf:"varint,2,opt,name=revision" json:"revision,omitempty"`
}

func (m *GetSignedMapRootByRevisionRequest) Reset()         { *m = GetSignedMapRootByRevisionRequest{} }
func (m *GetSignedMapRootByRevisionRequest) String() string { return proto.CompactTextString(m) }
func (*GetSignedMapRootByRevisionRequest) ProtoMessage()    {}
func (*GetSignedMapRootByRevisionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor1, []int{8}
}

func (m *GetSignedMapRootByRevisionRequest) GetMapId() int64 {
	if m != nil {
		return m.MapId
	}
	return 0
}

func (m *GetSignedMapRootByRevisionRequest) GetRevision() int64 {
	if m != nil {
		return m.Revision
	}
	return 0
}

type GetSignedMapRootResponse struct {
	MapRoot *SignedMapRoot `protobuf:"bytes,2,opt,name=map_root,json=mapRoot" json:"map_root,omitempty"`
}

func (m *GetSignedMapRootResponse) Reset()                    { *m = GetSignedMapRootResponse{} }
func (m *GetSignedMapRootResponse) String() string            { return proto.CompactTextString(m) }
func (*GetSignedMapRootResponse) ProtoMessage()               {}
func (*GetSignedMapRootResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{9} }

func (m *GetSignedMapRootResponse) GetMapRoot() *SignedMapRoot {
	if m != nil {
		return m.MapRoot
	}
	return nil
}

type InitMapRequest struct {
	MapId int64 `protobuf:"varint,1,opt,name=map_id,json=mapId" json:"map_id,omitempty"`
}

func (m *InitMapRequest) Reset()                    { *m = InitMapRequest{} }
func (m *InitMapRequest) String() string            { return proto.CompactTextString(m) }
func (*InitMapRequest) ProtoMessage()               {}
func (*InitMapRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{10} }

func (m *InitMapRequest) GetMapId() int64 {
	if m != nil {
		return m.MapId
	}
	return 0
}

type InitMapResponse struct {
	Created *SignedMapRoot `protobuf:"bytes,1,opt,name=created" json:"created,omitempty"`
}

func (m *InitMapResponse) Reset()                    { *m = InitMapResponse{} }
func (m *InitMapResponse) String() string            { return proto.CompactTextString(m) }
func (*InitMapResponse) ProtoMessage()               {}
func (*InitMapResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{11} }

func (m *InitMapResponse) GetCreated() *SignedMapRoot {
	if m != nil {
		return m.Created
	}
	return nil
}

func init() {
	proto.RegisterType((*MapLeaf)(nil), "trillian.MapLeaf")
	proto.RegisterType((*MapLeafInclusion)(nil), "trillian.MapLeafInclusion")
	proto.RegisterType((*GetMapLeavesRequest)(nil), "trillian.GetMapLeavesRequest")
	proto.RegisterType((*GetMapLeavesByRevisionRequest)(nil), "trillian.GetMapLeavesByRevisionRequest")
	proto.RegisterType((*GetMapLeavesResponse)(nil), "trillian.GetMapLeavesResponse")
	proto.RegisterType((*SetMapLeavesRequest)(nil), "trillian.SetMapLeavesRequest")
	proto.RegisterType((*SetMapLeavesResponse)(nil), "trillian.SetMapLeavesResponse")
	proto.RegisterType((*GetSignedMapRootRequest)(nil), "trillian.GetSignedMapRootRequest")
	proto.RegisterType((*GetSignedMapRootByRevisionRequest)(nil), "trillian.GetSignedMapRootByRevisionRequest")
	proto.RegisterType((*GetSignedMapRootResponse)(nil), "trillian.GetSignedMapRootResponse")
	proto.RegisterType((*InitMapRequest)(nil), "trillian.InitMapRequest")
	proto.RegisterType((*InitMapResponse)(nil), "trillian.InitMapResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for TrillianMap service

type TrillianMapClient interface {
	// GetLeaves returns an inclusion proof for each index requested.
	// For indexes that do not exist, the inclusion proof will use nil for the empty leaf value.
	GetLeaves(ctx context.Context, in *GetMapLeavesRequest, opts ...grpc.CallOption) (*GetMapLeavesResponse, error)
	GetLeavesByRevision(ctx context.Context, in *GetMapLeavesByRevisionRequest, opts ...grpc.CallOption) (*GetMapLeavesResponse, error)
	SetLeaves(ctx context.Context, in *SetMapLeavesRequest, opts ...grpc.CallOption) (*SetMapLeavesResponse, error)
	GetSignedMapRoot(ctx context.Context, in *GetSignedMapRootRequest, opts ...grpc.CallOption) (*GetSignedMapRootResponse, error)
	GetSignedMapRootByRevision(ctx context.Context, in *GetSignedMapRootByRevisionRequest, opts ...grpc.CallOption) (*GetSignedMapRootResponse, error)
	InitMap(ctx context.Context, in *InitMapRequest, opts ...grpc.CallOption) (*InitMapResponse, error)
}

type trillianMapClient struct {
	cc *grpc.ClientConn
}

func NewTrillianMapClient(cc *grpc.ClientConn) TrillianMapClient {
	return &trillianMapClient{cc}
}

func (c *trillianMapClient) GetLeaves(ctx context.Context, in *GetMapLeavesRequest, opts ...grpc.CallOption) (*GetMapLeavesResponse, error) {
	out := new(GetMapLeavesResponse)
	err := grpc.Invoke(ctx, "/trillian.TrillianMap/GetLeaves", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trillianMapClient) GetLeavesByRevision(ctx context.Context, in *GetMapLeavesByRevisionRequest, opts ...grpc.CallOption) (*GetMapLeavesResponse, error) {
	out := new(GetMapLeavesResponse)
	err := grpc.Invoke(ctx, "/trillian.TrillianMap/GetLeavesByRevision", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trillianMapClient) SetLeaves(ctx context.Context, in *SetMapLeavesRequest, opts ...grpc.CallOption) (*SetMapLeavesResponse, error) {
	out := new(SetMapLeavesResponse)
	err := grpc.Invoke(ctx, "/trillian.TrillianMap/SetLeaves", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trillianMapClient) GetSignedMapRoot(ctx context.Context, in *GetSignedMapRootRequest, opts ...grpc.CallOption) (*GetSignedMapRootResponse, error) {
	out := new(GetSignedMapRootResponse)
	err := grpc.Invoke(ctx, "/trillian.TrillianMap/GetSignedMapRoot", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trillianMapClient) GetSignedMapRootByRevision(ctx context.Context, in *GetSignedMapRootByRevisionRequest, opts ...grpc.CallOption) (*GetSignedMapRootResponse, error) {
	out := new(GetSignedMapRootResponse)
	err := grpc.Invoke(ctx, "/trillian.TrillianMap/GetSignedMapRootByRevision", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trillianMapClient) InitMap(ctx context.Context, in *InitMapRequest, opts ...grpc.CallOption) (*InitMapResponse, error) {
	out := new(InitMapResponse)
	err := grpc.Invoke(ctx, "/trillian.TrillianMap/InitMap", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for TrillianMap service

type TrillianMapServer interface {
	// GetLeaves returns an inclusion proof for each index requested.
	// For indexes that do not exist, the inclusion proof will use nil for the empty leaf value.
	GetLeaves(context.Context, *GetMapLeavesRequest) (*GetMapLeavesResponse, error)
	GetLeavesByRevision(context.Context, *GetMapLeavesByRevisionRequest) (*GetMapLeavesResponse, error)
	SetLeaves(context.Context, *SetMapLeavesRequest) (*SetMapLeavesResponse, error)
	GetSignedMapRoot(context.Context, *GetSignedMapRootRequest) (*GetSignedMapRootResponse, error)
	GetSignedMapRootByRevision(context.Context, *GetSignedMapRootByRevisionRequest) (*GetSignedMapRootResponse, error)
	InitMap(context.Context, *InitMapRequest) (*InitMapResponse, error)
}

func RegisterTrillianMapServer(s *grpc.Server, srv TrillianMapServer) {
	s.RegisterService(&_TrillianMap_serviceDesc, srv)
}

func _TrillianMap_GetLeaves_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMapLeavesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrillianMapServer).GetLeaves(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/trillian.TrillianMap/GetLeaves",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrillianMapServer).GetLeaves(ctx, req.(*GetMapLeavesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrillianMap_GetLeavesByRevision_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMapLeavesByRevisionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrillianMapServer).GetLeavesByRevision(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/trillian.TrillianMap/GetLeavesByRevision",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrillianMapServer).GetLeavesByRevision(ctx, req.(*GetMapLeavesByRevisionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrillianMap_SetLeaves_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetMapLeavesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrillianMapServer).SetLeaves(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/trillian.TrillianMap/SetLeaves",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrillianMapServer).SetLeaves(ctx, req.(*SetMapLeavesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrillianMap_GetSignedMapRoot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSignedMapRootRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrillianMapServer).GetSignedMapRoot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/trillian.TrillianMap/GetSignedMapRoot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrillianMapServer).GetSignedMapRoot(ctx, req.(*GetSignedMapRootRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrillianMap_GetSignedMapRootByRevision_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSignedMapRootByRevisionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrillianMapServer).GetSignedMapRootByRevision(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/trillian.TrillianMap/GetSignedMapRootByRevision",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrillianMapServer).GetSignedMapRootByRevision(ctx, req.(*GetSignedMapRootByRevisionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrillianMap_InitMap_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InitMapRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrillianMapServer).InitMap(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/trillian.TrillianMap/InitMap",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrillianMapServer).InitMap(ctx, req.(*InitMapRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TrillianMap_serviceDesc = grpc.ServiceDesc{
	ServiceName: "trillian.TrillianMap",
	HandlerType: (*TrillianMapServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetLeaves",
			Handler:    _TrillianMap_GetLeaves_Handler,
		},
		{
			MethodName: "GetLeavesByRevision",
			Handler:    _TrillianMap_GetLeavesByRevision_Handler,
		},
		{
			MethodName: "SetLeaves",
			Handler:    _TrillianMap_SetLeaves_Handler,
		},
		{
			MethodName: "GetSignedMapRoot",
			Handler:    _TrillianMap_GetSignedMapRoot_Handler,
		},
		{
			MethodName: "GetSignedMapRootByRevision",
			Handler:    _TrillianMap_GetSignedMapRootByRevision_Handler,
		},
		{
			MethodName: "InitMap",
			Handler:    _TrillianMap_InitMap_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "trillian_map_api.proto",
}

func init() { proto.RegisterFile("trillian_map_api.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 711 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x55, 0x4d, 0x4f, 0xdb, 0x4a,
	0x14, 0x7d, 0x4e, 0x02, 0x49, 0x6e, 0x9e, 0x78, 0x79, 0x43, 0x5a, 0x8c, 0x21, 0x15, 0x18, 0x21,
	0x8a, 0x90, 0x62, 0x48, 0x77, 0xec, 0x40, 0x48, 0x7c, 0x08, 0x10, 0x72, 0x2a, 0x2a, 0x75, 0x93,
	0x4e, 0x92, 0x21, 0x19, 0xc9, 0xf6, 0xb8, 0xf1, 0x24, 0x82, 0x22, 0x36, 0x5d, 0x74, 0xdb, 0x45,
	0xbb, 0xea, 0xa2, 0x7f, 0xaa, 0x7f, 0xa1, 0x3f, 0xa4, 0x9a, 0xf1, 0xd8, 0xf9, 0x72, 0xd2, 0xa8,
	0xdd, 0x65, 0xe6, 0xdc, 0x7b, 0xcf, 0xb9, 0xf7, 0x9e, 0x89, 0xe1, 0x39, 0xef, 0x52, 0xc7, 0xa1,
	0xd8, 0xab, 0xbb, 0xd8, 0xaf, 0x63, 0x9f, 0x56, 0xfc, 0x2e, 0xe3, 0x0c, 0xe5, 0xa2, 0x7b, 0x63,
	0x29, 0xfa, 0x15, 0x22, 0xc6, 0x7a, 0x9b, 0xb1, 0xb6, 0x43, 0x2c, 0xec, 0x53, 0x0b, 0x7b, 0x1e,
	0xe3, 0x98, 0x53, 0xe6, 0x05, 0x0a, 0x5d, 0x55, 0xa8, 0x3c, 0x35, 0x7a, 0x77, 0x16, 0xf6, 0x1e,
	0x42, 0xc8, 0xfc, 0x00, 0xd9, 0x2b, 0xec, 0x5f, 0x12, 0x7c, 0x87, 0x4a, 0xb0, 0x40, 0xbd, 0x16,
	0xb9, 0xd7, 0xb5, 0x0d, 0xed, 0xe5, 0xbf, 0x76, 0x78, 0x40, 0x6b, 0x90, 0x77, 0x08, 0xbe, 0xab,
	0x77, 0x70, 0xd0, 0xd1, 0x53, 0x12, 0xc9, 0x89, 0x8b, 0x33, 0x1c, 0x74, 0x50, 0x19, 0x40, 0x82,
	0x7d, 0xec, 0xf4, 0x88, 0x9e, 0x96, 0xa8, 0x0c, 0xbf, 0x15, 0x17, 0x02, 0x26, 0xf7, 0xbc, 0x8b,
	0xeb, 0x2d, 0xcc, 0xb1, 0x9e, 0x09, 0x61, 0x79, 0x73, 0x82, 0x39, 0x36, 0xdf, 0x40, 0x51, 0x71,
	0x9f, 0x7b, 0x4d, 0xa7, 0x17, 0x50, 0xe6, 0xa1, 0x6d, 0xc8, 0x88, 0x7c, 0xa9, 0xa1, 0x50, 0xfd,
	0xbf, 0x12, 0xf7, 0xa9, 0x22, 0x6d, 0x09, 0xa3, 0x75, 0xc8, 0xd3, 0x28, 0x47, 0x4f, 0x6d, 0xa4,
	0x45, 0xe1, 0xf8, 0xc2, 0x3c, 0x83, 0xe5, 0x53, 0xc2, 0xc3, 0x8c, 0x3e, 0x09, 0x6c, 0xf2, 0xbe,
	0x47, 0x02, 0x8e, 0x9e, 0xc1, 0xa2, 0x98, 0x27, 0x6d, 0xc9, 0xea, 0x69, 0x7b, 0xc1, 0xc5, 0xfe,
	0x79, 0x6b, 0xd0, 0x77, 0x58, 0x27, 0x3c, 0x5c, 0x64, 0x72, 0xe9, 0x62, 0xc6, 0xec, 0x40, 0x79,
	0xb8, 0xd2, 0xf1, 0x83, 0x4d, 0xfa, 0x54, 0x70, 0xfc, 0x49, 0x4d, 0x64, 0x40, 0xae, 0xab, 0xf2,
	0xe5, 0xb0, 0xd2, 0x76, 0x7c, 0x36, 0xbf, 0x6a, 0x50, 0x1a, 0x15, 0x1d, 0xf8, 0xcc, 0x0b, 0x08,
	0x3a, 0x03, 0x24, 0x18, 0xe4, 0x9c, 0x47, 0x7b, 0x2e, 0x54, 0x8d, 0x89, 0xf9, 0xc4, 0x93, 0xb4,
	0x8b, 0xee, 0xf8, 0x6c, 0xab, 0x90, 0x13, 0x95, 0xba, 0x8c, 0x71, 0x49, 0x5f, 0xa8, 0xae, 0x0c,
	0xf2, 0x6b, 0xb4, 0xed, 0x91, 0xd6, 0x15, 0xf6, 0x6d, 0xc6, 0xb8, 0x9d, 0x75, 0xc3, 0x1f, 0xe6,
	0x67, 0x0d, 0x96, 0x6b, 0xf3, 0xcf, 0x72, 0x17, 0x16, 0x1d, 0x19, 0xa7, 0x04, 0x26, 0x2c, 0x50,
	0x05, 0xa0, 0x7d, 0xc8, 0xb9, 0x84, 0xe3, 0xd8, 0x1a, 0x85, 0x6a, 0xa9, 0x12, 0xfa, 0xb4, 0x12,
	0xf9, 0xb4, 0x72, 0xe4, 0x3d, 0xd8, 0x71, 0x94, 0x5a, 0xc9, 0x05, 0x94, 0x6a, 0x49, 0x73, 0x1a,
	0xee, 0x2e, 0x35, 0x67, 0x77, 0xfb, 0xb0, 0x72, 0x4a, 0xf8, 0x28, 0x38, 0xb3, 0x41, 0xf3, 0x16,
	0x36, 0xc7, 0x33, 0xe6, 0x36, 0xc5, 0xf0, 0xfa, 0x53, 0x63, 0xeb, 0xbf, 0x06, 0x7d, 0x52, 0xc9,
	0x5f, 0x74, 0xb6, 0x03, 0x4b, 0xe7, 0x1e, 0x15, 0x63, 0xfa, 0x4d, 0x43, 0x27, 0xf0, 0x5f, 0x1c,
	0xa8, 0xf8, 0x0e, 0x20, 0xdb, 0xec, 0x12, 0xcc, 0x49, 0x4b, 0x3d, 0xc3, 0xe9, 0x74, 0x2a, 0xae,
	0xfa, 0x6d, 0x01, 0x0a, 0xaf, 0x55, 0xcc, 0x15, 0xf6, 0xd1, 0x25, 0xe4, 0x4f, 0x09, 0x0f, 0x37,
	0x84, 0xca, 0x83, 0xf4, 0x84, 0x67, 0x69, 0xbc, 0x98, 0x06, 0x87, 0x72, 0xcc, 0x7f, 0xd0, 0x3b,
	0xf9, 0x9e, 0xc7, 0x9f, 0x20, 0xda, 0x49, 0x4e, 0x9c, 0xd8, 0xc7, 0x1c, 0x0c, 0x97, 0x90, 0xaf,
	0x25, 0xe9, 0xad, 0xcd, 0xd6, 0x5b, 0x4b, 0xae, 0xf6, 0x49, 0x83, 0xe2, 0xf8, 0x36, 0xd1, 0xe6,
	0x88, 0x88, 0x24, 0xcf, 0x19, 0xe6, 0xac, 0x10, 0x55, 0x7d, 0xef, 0xe3, 0x8f, 0x9f, 0x5f, 0x52,
	0xdb, 0x68, 0xcb, 0xea, 0x1f, 0x34, 0x08, 0xc7, 0x07, 0x96, 0x8b, 0xfd, 0xc0, 0x7a, 0x0c, 0x77,
	0xfb, 0x64, 0x09, 0x97, 0x04, 0x87, 0x0e, 0xe6, 0x62, 0xe7, 0xdf, 0x35, 0x30, 0xa6, 0xdb, 0x15,
	0xed, 0x4d, 0xe7, 0x9b, 0x1c, 0xe2, 0x3c, 0xe2, 0x2c, 0x29, 0x6e, 0x17, 0xed, 0xcc, 0x12, 0x67,
	0x3d, 0x46, 0xae, 0x7f, 0x42, 0x4d, 0xc8, 0x2a, 0xf7, 0x21, 0x7d, 0x50, 0x7f, 0xd4, 0xb9, 0xc6,
	0x6a, 0x02, 0xa2, 0x08, 0xb7, 0x24, 0x61, 0xd9, 0x5c, 0x4b, 0x26, 0x3c, 0xa4, 0x1e, 0xe5, 0xc7,
	0xd7, 0xb0, 0xda, 0x64, 0x6e, 0xf4, 0xe7, 0x32, 0xfa, 0xe5, 0x3c, 0x5e, 0x1e, 0xb2, 0xed, 0x91,
	0x4f, 0x6f, 0xc4, 0xe5, 0x8d, 0xf6, 0xd6, 0x68, 0x53, 0xde, 0xe9, 0x35, 0x2a, 0x4d, 0xe6, 0x5a,
	0xea, 0xeb, 0x19, 0x25, 0x36, 0x16, 0x65, 0xe6, 0xab, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x60,
	0x7f, 0xb2, 0x65, 0xa7, 0x07, 0x00, 0x00,
}
