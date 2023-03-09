// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tomahawk/erc20/v1/erc20.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/x/bank/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Owner enumerates the ownership of a ERC20 contract.
type Owner int32

const (
	// OWNER_UNSPECIFIED defines an invalid/undefined owner.
	OWNER_UNSPECIFIED Owner = 0
	// OWNER_MODULE erc20 is owned by the erc20 module account.
	OWNER_MODULE Owner = 1
	// EXTERNAL erc20 is owned by an external account.
	OWNER_EXTERNAL Owner = 2
)

var Owner_name = map[int32]string{
	0: "OWNER_UNSPECIFIED",
	1: "OWNER_MODULE",
	2: "OWNER_EXTERNAL",
}

var Owner_value = map[string]int32{
	"OWNER_UNSPECIFIED": 0,
	"OWNER_MODULE":      1,
	"OWNER_EXTERNAL":    2,
}

func (x Owner) String() string {
	return proto.EnumName(Owner_name, int32(x))
}

func (Owner) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5c364669f6882b8b, []int{0}
}

// TokenPair defines an instance that records a pairing consisting of a native
//  Cosmos Coin and an ERC20 token address.
type TokenPair struct {
	// address of ERC20 contract token
	Erc20Address string `protobuf:"bytes,1,opt,name=erc20_address,json=erc20Address,proto3" json:"erc20_address,omitempty"`
	// cosmos base denomination to be mapped to
	Denom string `protobuf:"bytes,2,opt,name=denom,proto3" json:"denom,omitempty"`
	// shows token mapping enable status
	Enabled bool `protobuf:"varint,3,opt,name=enabled,proto3" json:"enabled,omitempty"`
	// ERC20 owner address ENUM (0 invalid, 1 ModuleAccount, 2 external address)
	ContractOwner Owner `protobuf:"varint,4,opt,name=contract_owner,json=contractOwner,proto3,enum=tomahawk.erc20.v1.Owner" json:"contract_owner,omitempty"`
}

func (m *TokenPair) Reset()         { *m = TokenPair{} }
func (m *TokenPair) String() string { return proto.CompactTextString(m) }
func (*TokenPair) ProtoMessage()    {}
func (*TokenPair) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c364669f6882b8b, []int{0}
}
func (m *TokenPair) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TokenPair) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TokenPair.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TokenPair) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TokenPair.Merge(m, src)
}
func (m *TokenPair) XXX_Size() int {
	return m.Size()
}
func (m *TokenPair) XXX_DiscardUnknown() {
	xxx_messageInfo_TokenPair.DiscardUnknown(m)
}

var xxx_messageInfo_TokenPair proto.InternalMessageInfo

func (m *TokenPair) GetErc20Address() string {
	if m != nil {
		return m.Erc20Address
	}
	return ""
}

func (m *TokenPair) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

func (m *TokenPair) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

func (m *TokenPair) GetContractOwner() Owner {
	if m != nil {
		return m.ContractOwner
	}
	return OWNER_UNSPECIFIED
}

// RegisterCoinProposal is a gov Content type to register a token pair for a
// native Cosmos coin.
type RegisterCoinProposal struct {
	// title of the proposal
	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	// proposal description
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// metadata of the native Cosmos coin
	Metadata types.Metadata `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata"`
}

func (m *RegisterCoinProposal) Reset()         { *m = RegisterCoinProposal{} }
func (m *RegisterCoinProposal) String() string { return proto.CompactTextString(m) }
func (*RegisterCoinProposal) ProtoMessage()    {}
func (*RegisterCoinProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c364669f6882b8b, []int{1}
}
func (m *RegisterCoinProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RegisterCoinProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RegisterCoinProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RegisterCoinProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterCoinProposal.Merge(m, src)
}
func (m *RegisterCoinProposal) XXX_Size() int {
	return m.Size()
}
func (m *RegisterCoinProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterCoinProposal.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterCoinProposal proto.InternalMessageInfo

func (m *RegisterCoinProposal) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *RegisterCoinProposal) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *RegisterCoinProposal) GetMetadata() types.Metadata {
	if m != nil {
		return m.Metadata
	}
	return types.Metadata{}
}

// RegisterERC20Proposal is a gov Content type to register a token pair for an
// ERC20 token
type RegisterERC20Proposal struct {
	// title of the proposa  string title = 1;
	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	// proposal description
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// contract address of ERC20 token
	Erc20Address string `protobuf:"bytes,3,opt,name=erc20address,proto3" json:"erc20address,omitempty"`
}

func (m *RegisterERC20Proposal) Reset()         { *m = RegisterERC20Proposal{} }
func (m *RegisterERC20Proposal) String() string { return proto.CompactTextString(m) }
func (*RegisterERC20Proposal) ProtoMessage()    {}
func (*RegisterERC20Proposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c364669f6882b8b, []int{2}
}
func (m *RegisterERC20Proposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RegisterERC20Proposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RegisterERC20Proposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RegisterERC20Proposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterERC20Proposal.Merge(m, src)
}
func (m *RegisterERC20Proposal) XXX_Size() int {
	return m.Size()
}
func (m *RegisterERC20Proposal) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterERC20Proposal.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterERC20Proposal proto.InternalMessageInfo

func (m *RegisterERC20Proposal) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *RegisterERC20Proposal) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *RegisterERC20Proposal) GetErc20Address() string {
	if m != nil {
		return m.Erc20Address
	}
	return ""
}

// ToggleTokenConversionProposal is a gov Content type to toggle the conversion
// of a token pair.
type ToggleTokenConversionProposal struct {
	// title of the proposal
	Title string `protobuf:"bytes,1,opt,name=Title,proto3" json:"Title,omitempty"`
	// proposal description
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// token identifier can be either the hex contract address of the ERC20 or the
	// Cosmos base denomination
	Token string `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
}

func (m *ToggleTokenConversionProposal) Reset()         { *m = ToggleTokenConversionProposal{} }
func (m *ToggleTokenConversionProposal) String() string { return proto.CompactTextString(m) }
func (*ToggleTokenConversionProposal) ProtoMessage()    {}
func (*ToggleTokenConversionProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c364669f6882b8b, []int{3}
}
func (m *ToggleTokenConversionProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ToggleTokenConversionProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ToggleTokenConversionProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ToggleTokenConversionProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ToggleTokenConversionProposal.Merge(m, src)
}
func (m *ToggleTokenConversionProposal) XXX_Size() int {
	return m.Size()
}
func (m *ToggleTokenConversionProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_ToggleTokenConversionProposal.DiscardUnknown(m)
}

var xxx_messageInfo_ToggleTokenConversionProposal proto.InternalMessageInfo

func (m *ToggleTokenConversionProposal) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *ToggleTokenConversionProposal) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *ToggleTokenConversionProposal) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func init() {
	proto.RegisterEnum("tomahawk.erc20.v1.Owner", Owner_name, Owner_value)
	proto.RegisterType((*TokenPair)(nil), "tomahawk.erc20.v1.TokenPair")
	proto.RegisterType((*RegisterCoinProposal)(nil), "tomahawk.erc20.v1.RegisterCoinProposal")
	proto.RegisterType((*RegisterERC20Proposal)(nil), "tomahawk.erc20.v1.RegisterERC20Proposal")
	proto.RegisterType((*ToggleTokenConversionProposal)(nil), "tomahawk.erc20.v1.ToggleTokenConversionProposal")
}

func init() { proto.RegisterFile("tomahawk/erc20/v1/erc20.proto", fileDescriptor_5c364669f6882b8b) }

var fileDescriptor_5c364669f6882b8b = []byte{
	// 499 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x93, 0x41, 0x6b, 0xdb, 0x3e,
	0x18, 0xc6, 0xad, 0x36, 0xfd, 0xff, 0x5b, 0xb5, 0x0d, 0x99, 0x48, 0x20, 0x04, 0xea, 0x86, 0xec,
	0x12, 0x06, 0xb3, 0x9b, 0xec, 0x36, 0x06, 0xa3, 0x75, 0x3d, 0xc8, 0x68, 0x93, 0xe0, 0xb9, 0x6c,
	0xec, 0x12, 0x64, 0x5b, 0x78, 0x26, 0x89, 0xde, 0x20, 0x69, 0xee, 0xfa, 0x0d, 0x76, 0xdc, 0x65,
	0xf7, 0xc1, 0xf6, 0x61, 0x7a, 0xec, 0x71, 0xa7, 0x31, 0x92, 0xcb, 0x3e, 0xc6, 0xb0, 0xe4, 0x8c,
	0xf4, 0x36, 0xd8, 0x4d, 0xcf, 0xf3, 0xbe, 0x96, 0x7e, 0x7e, 0x1f, 0x09, 0xb7, 0x62, 0xca, 0x15,
	0xb8, 0x4c, 0xc4, 0xfd, 0x13, 0x37, 0xef, 0x99, 0x85, 0xb3, 0x10, 0xa0, 0x80, 0x54, 0x75, 0xcd,
	0x31, 0x56, 0xde, 0x6b, 0xd5, 0x53, 0x48, 0x41, 0x97, 0xdc, 0x62, 0x65, 0xba, 0x5a, 0x76, 0x0c,
	0x72, 0x0e, 0xd2, 0x8d, 0x28, 0x9f, 0xba, 0x79, 0x2f, 0x62, 0x8a, 0xf6, 0xb4, 0x30, 0xf5, 0xce,
	0x37, 0x84, 0xf7, 0x42, 0x98, 0x32, 0x3e, 0xa6, 0x99, 0x20, 0x0f, 0xf1, 0xa1, 0xde, 0x6f, 0x42,
	0x93, 0x44, 0x30, 0x29, 0x9b, 0xa8, 0x8d, 0xba, 0x7b, 0xc1, 0x81, 0x36, 0x4f, 0x8d, 0x47, 0xea,
	0x78, 0x27, 0x61, 0x1c, 0xe6, 0xcd, 0x2d, 0x5d, 0x34, 0x82, 0x34, 0xf1, 0xff, 0x8c, 0xd3, 0x68,
	0xc6, 0x92, 0xe6, 0x76, 0x1b, 0x75, 0x77, 0x83, 0xb5, 0x24, 0xcf, 0x70, 0x35, 0x06, 0xae, 0x04,
	0x8d, 0xd5, 0x04, 0xae, 0x39, 0x13, 0xcd, 0x4a, 0x1b, 0x75, 0xab, 0xfd, 0x86, 0x73, 0xff, 0x0f,
	0x9c, 0x51, 0x51, 0x0c, 0x0e, 0xd7, 0xcd, 0x5a, 0x3e, 0xad, 0xfc, 0xfa, 0x72, 0x8c, 0x3a, 0x9f,
	0x11, 0xae, 0x07, 0x2c, 0xcd, 0xa4, 0x62, 0xc2, 0x83, 0x8c, 0x8f, 0x05, 0x2c, 0x40, 0xd2, 0x59,
	0x01, 0xa3, 0x32, 0x35, 0x63, 0x25, 0xa9, 0x11, 0xa4, 0x8d, 0xf7, 0x13, 0x26, 0x63, 0x91, 0x2d,
	0x54, 0x06, 0xbc, 0x04, 0xdd, 0xb4, 0xc8, 0x73, 0xbc, 0x3b, 0x67, 0x8a, 0x26, 0x54, 0x51, 0xcd,
	0xbb, 0xdf, 0x3f, 0x72, 0xcc, 0xa8, 0x1c, 0x3d, 0x9d, 0x72, 0x54, 0xce, 0x65, 0xd9, 0x74, 0x56,
	0xb9, 0xfd, 0x71, 0x6c, 0x05, 0x7f, 0x3e, 0xd2, 0x5c, 0x56, 0xe7, 0x06, 0x37, 0xd6, 0x58, 0x7e,
	0xe0, 0xf5, 0x4f, 0xfe, 0x99, 0xab, 0x83, 0xcd, 0xb0, 0xd7, 0x01, 0x6c, 0x6f, 0x04, 0x50, 0x7a,
	0xe5, 0xd1, 0x12, 0x1f, 0x85, 0x90, 0xa6, 0x33, 0xa6, 0xe3, 0xf3, 0x80, 0xe7, 0x4c, 0xc8, 0x0c,
	0xee, 0x8d, 0x26, 0xdc, 0x44, 0x08, 0xff, 0x12, 0xa1, 0x40, 0x2f, 0xb6, 0x2c, 0xcf, 0x36, 0xc2,
	0xe4, 0xf0, 0xe8, 0x25, 0xde, 0xd1, 0xb1, 0x90, 0x06, 0x7e, 0x30, 0x7a, 0x3d, 0xf4, 0x83, 0xc9,
	0xd5, 0xf0, 0xd5, 0xd8, 0xf7, 0x06, 0x2f, 0x06, 0xfe, 0x79, 0xcd, 0x22, 0x35, 0x7c, 0x60, 0xec,
	0xcb, 0xd1, 0xf9, 0xd5, 0x85, 0x5f, 0x43, 0x84, 0xe0, 0xaa, 0x71, 0xfc, 0x37, 0xa1, 0x1f, 0x0c,
	0x4f, 0x2f, 0x6a, 0x5b, 0xad, 0xca, 0xc7, 0xaf, 0xb6, 0x75, 0x36, 0xb8, 0x5d, 0xda, 0xe8, 0x6e,
	0x69, 0xa3, 0x9f, 0x4b, 0x1b, 0x7d, 0x5a, 0xd9, 0xd6, 0xdd, 0xca, 0xb6, 0xbe, 0xaf, 0x6c, 0xeb,
	0xad, 0x9b, 0x66, 0xea, 0xdd, 0xfb, 0xc8, 0x89, 0x61, 0xee, 0x7a, 0xc5, 0x1d, 0x79, 0x3c, 0x64,
	0xea, 0x1a, 0xc4, 0xd4, 0x28, 0x37, 0xef, 0xbb, 0x1f, 0xca, 0x47, 0xa1, 0x6e, 0x16, 0x4c, 0x46,
	0xff, 0xe9, 0xcb, 0xfc, 0xe4, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x29, 0x55, 0xba, 0x56, 0x30,
	0x03, 0x00, 0x00,
}

func (this *TokenPair) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*TokenPair)
	if !ok {
		that2, ok := that.(TokenPair)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Erc20Address != that1.Erc20Address {
		return false
	}
	if this.Denom != that1.Denom {
		return false
	}
	if this.Enabled != that1.Enabled {
		return false
	}
	if this.ContractOwner != that1.ContractOwner {
		return false
	}
	return true
}
func (this *ToggleTokenConversionProposal) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ToggleTokenConversionProposal)
	if !ok {
		that2, ok := that.(ToggleTokenConversionProposal)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Title != that1.Title {
		return false
	}
	if this.Description != that1.Description {
		return false
	}
	if this.Token != that1.Token {
		return false
	}
	return true
}
func (m *TokenPair) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TokenPair) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TokenPair) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ContractOwner != 0 {
		i = encodeVarintErc20(dAtA, i, uint64(m.ContractOwner))
		i--
		dAtA[i] = 0x20
	}
	if m.Enabled {
		i--
		if m.Enabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintErc20(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Erc20Address) > 0 {
		i -= len(m.Erc20Address)
		copy(dAtA[i:], m.Erc20Address)
		i = encodeVarintErc20(dAtA, i, uint64(len(m.Erc20Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *RegisterCoinProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RegisterCoinProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RegisterCoinProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Metadata.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintErc20(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintErc20(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintErc20(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *RegisterERC20Proposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RegisterERC20Proposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RegisterERC20Proposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Erc20Address) > 0 {
		i -= len(m.Erc20Address)
		copy(dAtA[i:], m.Erc20Address)
		i = encodeVarintErc20(dAtA, i, uint64(len(m.Erc20Address)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintErc20(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintErc20(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ToggleTokenConversionProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ToggleTokenConversionProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ToggleTokenConversionProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Token) > 0 {
		i -= len(m.Token)
		copy(dAtA[i:], m.Token)
		i = encodeVarintErc20(dAtA, i, uint64(len(m.Token)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintErc20(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintErc20(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintErc20(dAtA []byte, offset int, v uint64) int {
	offset -= sovErc20(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *TokenPair) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Erc20Address)
	if l > 0 {
		n += 1 + l + sovErc20(uint64(l))
	}
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovErc20(uint64(l))
	}
	if m.Enabled {
		n += 2
	}
	if m.ContractOwner != 0 {
		n += 1 + sovErc20(uint64(m.ContractOwner))
	}
	return n
}

func (m *RegisterCoinProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovErc20(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovErc20(uint64(l))
	}
	l = m.Metadata.Size()
	n += 1 + l + sovErc20(uint64(l))
	return n
}

func (m *RegisterERC20Proposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovErc20(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovErc20(uint64(l))
	}
	l = len(m.Erc20Address)
	if l > 0 {
		n += 1 + l + sovErc20(uint64(l))
	}
	return n
}

func (m *ToggleTokenConversionProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovErc20(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovErc20(uint64(l))
	}
	l = len(m.Token)
	if l > 0 {
		n += 1 + l + sovErc20(uint64(l))
	}
	return n
}

func sovErc20(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozErc20(x uint64) (n int) {
	return sovErc20(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TokenPair) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowErc20
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TokenPair: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TokenPair: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Erc20Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowErc20
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthErc20
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthErc20
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Erc20Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowErc20
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthErc20
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthErc20
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Enabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowErc20
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Enabled = bool(v != 0)
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractOwner", wireType)
			}
			m.ContractOwner = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowErc20
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ContractOwner |= Owner(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipErc20(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthErc20
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *RegisterCoinProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowErc20
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RegisterCoinProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RegisterCoinProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowErc20
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthErc20
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthErc20
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowErc20
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthErc20
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthErc20
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowErc20
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthErc20
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthErc20
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Metadata.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipErc20(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthErc20
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *RegisterERC20Proposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowErc20
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RegisterERC20Proposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RegisterERC20Proposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowErc20
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthErc20
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthErc20
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowErc20
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthErc20
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthErc20
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Erc20Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowErc20
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthErc20
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthErc20
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Erc20Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipErc20(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthErc20
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ToggleTokenConversionProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowErc20
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ToggleTokenConversionProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ToggleTokenConversionProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowErc20
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthErc20
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthErc20
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowErc20
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthErc20
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthErc20
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Token", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowErc20
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthErc20
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthErc20
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Token = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipErc20(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthErc20
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipErc20(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowErc20
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowErc20
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowErc20
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthErc20
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupErc20
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthErc20
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthErc20        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowErc20          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupErc20 = fmt.Errorf("proto: unexpected end of group")
)
