// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: nft/nft.proto

package art_admin

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// NftClient is the client API for Nft service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NftClient interface {
	// Method used in ui for submitting drawing nft reference
	NewNFTMintRequest(ctx context.Context, in *NFTMintRequestToUpload, opts ...grpc.CallOption) (*NFTMintRequestWithStatus, error)
	// List paged mint requests by status
	ListNFTMintRequestsPaged(ctx context.Context, in *ListPagedRequest, opts ...grpc.CallOption) (*NFTMintRequestListArray, error)
	// Delete mint requests by internal id
	DeleteNFTMintRequestById(ctx context.Context, in *DeleteId, opts ...grpc.CallOption) (*DeleteStatus, error)
	// Upload resulted nft offchain from b64
	UpdateNFTOffchainUrl(ctx context.Context, in *UpdateNFTOffchainUrlRequest, opts ...grpc.CallOption) (*NFTMintRequestWithStatus, error)
	// Remove nft offchain url from mint request
	DeleteNFTOffchainUrl(ctx context.Context, in *DeleteId, opts ...grpc.CallOption) (*NFTMintRequestWithStatus, error)
	// Get all metadata with status StatusUploadedOffchain & StatusUploaded and create _metadata.json
	UploadOffchainMetadata(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*MetadataOffchainUrl, error)
	Burn(ctx context.Context, in *BurnRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	SetTrackingNumber(ctx context.Context, in *SetTrackingNumberRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// TODO: add rpc for getting metadata offchain url
	UploadIPFSMetadata(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type nftClient struct {
	cc grpc.ClientConnInterface
}

func NewNftClient(cc grpc.ClientConnInterface) NftClient {
	return &nftClient{cc}
}

func (c *nftClient) NewNFTMintRequest(ctx context.Context, in *NFTMintRequestToUpload, opts ...grpc.CallOption) (*NFTMintRequestWithStatus, error) {
	out := new(NFTMintRequestWithStatus)
	err := c.cc.Invoke(ctx, "/nft.Nft/NewNFTMintRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nftClient) ListNFTMintRequestsPaged(ctx context.Context, in *ListPagedRequest, opts ...grpc.CallOption) (*NFTMintRequestListArray, error) {
	out := new(NFTMintRequestListArray)
	err := c.cc.Invoke(ctx, "/nft.Nft/ListNFTMintRequestsPaged", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nftClient) DeleteNFTMintRequestById(ctx context.Context, in *DeleteId, opts ...grpc.CallOption) (*DeleteStatus, error) {
	out := new(DeleteStatus)
	err := c.cc.Invoke(ctx, "/nft.Nft/DeleteNFTMintRequestById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nftClient) UpdateNFTOffchainUrl(ctx context.Context, in *UpdateNFTOffchainUrlRequest, opts ...grpc.CallOption) (*NFTMintRequestWithStatus, error) {
	out := new(NFTMintRequestWithStatus)
	err := c.cc.Invoke(ctx, "/nft.Nft/UpdateNFTOffchainUrl", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nftClient) DeleteNFTOffchainUrl(ctx context.Context, in *DeleteId, opts ...grpc.CallOption) (*NFTMintRequestWithStatus, error) {
	out := new(NFTMintRequestWithStatus)
	err := c.cc.Invoke(ctx, "/nft.Nft/DeleteNFTOffchainUrl", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nftClient) UploadOffchainMetadata(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*MetadataOffchainUrl, error) {
	out := new(MetadataOffchainUrl)
	err := c.cc.Invoke(ctx, "/nft.Nft/UploadOffchainMetadata", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nftClient) Burn(ctx context.Context, in *BurnRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/nft.Nft/Burn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nftClient) SetTrackingNumber(ctx context.Context, in *SetTrackingNumberRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/nft.Nft/SetTrackingNumber", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nftClient) UploadIPFSMetadata(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/nft.Nft/UploadIPFSMetadata", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NftServer is the server API for Nft service.
// All implementations should embed UnimplementedNftServer
// for forward compatibility
type NftServer interface {
	// Method used in ui for submitting drawing nft reference
	NewNFTMintRequest(context.Context, *NFTMintRequestToUpload) (*NFTMintRequestWithStatus, error)
	// List paged mint requests by status
	ListNFTMintRequestsPaged(context.Context, *ListPagedRequest) (*NFTMintRequestListArray, error)
	// Delete mint requests by internal id
	DeleteNFTMintRequestById(context.Context, *DeleteId) (*DeleteStatus, error)
	// Upload resulted nft offchain from b64
	UpdateNFTOffchainUrl(context.Context, *UpdateNFTOffchainUrlRequest) (*NFTMintRequestWithStatus, error)
	// Remove nft offchain url from mint request
	DeleteNFTOffchainUrl(context.Context, *DeleteId) (*NFTMintRequestWithStatus, error)
	// Get all metadata with status StatusUploadedOffchain & StatusUploaded and create _metadata.json
	UploadOffchainMetadata(context.Context, *emptypb.Empty) (*MetadataOffchainUrl, error)
	Burn(context.Context, *BurnRequest) (*emptypb.Empty, error)
	SetTrackingNumber(context.Context, *SetTrackingNumberRequest) (*emptypb.Empty, error)
	// TODO: add rpc for getting metadata offchain url
	UploadIPFSMetadata(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
}

// UnimplementedNftServer should be embedded to have forward compatible implementations.
type UnimplementedNftServer struct {
}

func (UnimplementedNftServer) NewNFTMintRequest(context.Context, *NFTMintRequestToUpload) (*NFTMintRequestWithStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewNFTMintRequest not implemented")
}
func (UnimplementedNftServer) ListNFTMintRequestsPaged(context.Context, *ListPagedRequest) (*NFTMintRequestListArray, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListNFTMintRequestsPaged not implemented")
}
func (UnimplementedNftServer) DeleteNFTMintRequestById(context.Context, *DeleteId) (*DeleteStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteNFTMintRequestById not implemented")
}
func (UnimplementedNftServer) UpdateNFTOffchainUrl(context.Context, *UpdateNFTOffchainUrlRequest) (*NFTMintRequestWithStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateNFTOffchainUrl not implemented")
}
func (UnimplementedNftServer) DeleteNFTOffchainUrl(context.Context, *DeleteId) (*NFTMintRequestWithStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteNFTOffchainUrl not implemented")
}
func (UnimplementedNftServer) UploadOffchainMetadata(context.Context, *emptypb.Empty) (*MetadataOffchainUrl, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadOffchainMetadata not implemented")
}
func (UnimplementedNftServer) Burn(context.Context, *BurnRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Burn not implemented")
}
func (UnimplementedNftServer) SetTrackingNumber(context.Context, *SetTrackingNumberRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetTrackingNumber not implemented")
}
func (UnimplementedNftServer) UploadIPFSMetadata(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadIPFSMetadata not implemented")
}

// UnsafeNftServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NftServer will
// result in compilation errors.
type UnsafeNftServer interface {
	mustEmbedUnimplementedNftServer()
}

func RegisterNftServer(s grpc.ServiceRegistrar, srv NftServer) {
	s.RegisterService(&Nft_ServiceDesc, srv)
}

func _Nft_NewNFTMintRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NFTMintRequestToUpload)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NftServer).NewNFTMintRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nft.Nft/NewNFTMintRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NftServer).NewNFTMintRequest(ctx, req.(*NFTMintRequestToUpload))
	}
	return interceptor(ctx, in, info, handler)
}

func _Nft_ListNFTMintRequestsPaged_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPagedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NftServer).ListNFTMintRequestsPaged(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nft.Nft/ListNFTMintRequestsPaged",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NftServer).ListNFTMintRequestsPaged(ctx, req.(*ListPagedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Nft_DeleteNFTMintRequestById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NftServer).DeleteNFTMintRequestById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nft.Nft/DeleteNFTMintRequestById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NftServer).DeleteNFTMintRequestById(ctx, req.(*DeleteId))
	}
	return interceptor(ctx, in, info, handler)
}

func _Nft_UpdateNFTOffchainUrl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateNFTOffchainUrlRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NftServer).UpdateNFTOffchainUrl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nft.Nft/UpdateNFTOffchainUrl",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NftServer).UpdateNFTOffchainUrl(ctx, req.(*UpdateNFTOffchainUrlRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Nft_DeleteNFTOffchainUrl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NftServer).DeleteNFTOffchainUrl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nft.Nft/DeleteNFTOffchainUrl",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NftServer).DeleteNFTOffchainUrl(ctx, req.(*DeleteId))
	}
	return interceptor(ctx, in, info, handler)
}

func _Nft_UploadOffchainMetadata_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NftServer).UploadOffchainMetadata(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nft.Nft/UploadOffchainMetadata",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NftServer).UploadOffchainMetadata(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Nft_Burn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BurnRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NftServer).Burn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nft.Nft/Burn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NftServer).Burn(ctx, req.(*BurnRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Nft_SetTrackingNumber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetTrackingNumberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NftServer).SetTrackingNumber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nft.Nft/SetTrackingNumber",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NftServer).SetTrackingNumber(ctx, req.(*SetTrackingNumberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Nft_UploadIPFSMetadata_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NftServer).UploadIPFSMetadata(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nft.Nft/UploadIPFSMetadata",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NftServer).UploadIPFSMetadata(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Nft_ServiceDesc is the grpc.ServiceDesc for Nft service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Nft_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "nft.Nft",
	HandlerType: (*NftServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NewNFTMintRequest",
			Handler:    _Nft_NewNFTMintRequest_Handler,
		},
		{
			MethodName: "ListNFTMintRequestsPaged",
			Handler:    _Nft_ListNFTMintRequestsPaged_Handler,
		},
		{
			MethodName: "DeleteNFTMintRequestById",
			Handler:    _Nft_DeleteNFTMintRequestById_Handler,
		},
		{
			MethodName: "UpdateNFTOffchainUrl",
			Handler:    _Nft_UpdateNFTOffchainUrl_Handler,
		},
		{
			MethodName: "DeleteNFTOffchainUrl",
			Handler:    _Nft_DeleteNFTOffchainUrl_Handler,
		},
		{
			MethodName: "UploadOffchainMetadata",
			Handler:    _Nft_UploadOffchainMetadata_Handler,
		},
		{
			MethodName: "Burn",
			Handler:    _Nft_Burn_Handler,
		},
		{
			MethodName: "SetTrackingNumber",
			Handler:    _Nft_SetTrackingNumber_Handler,
		},
		{
			MethodName: "UploadIPFSMetadata",
			Handler:    _Nft_UploadIPFSMetadata_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "nft/nft.proto",
}
