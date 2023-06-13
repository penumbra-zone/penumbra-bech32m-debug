// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: penumbra/core/transparent_proofs/v1alpha1/transparent_proofs.proto

package transparent_proofsv1alpha1

import (
	fmt "fmt"
	proto "github.com/cosmos/gogoproto/proto"
	io "io"
	math "math"
	math_bits "math/bits"
	v1alpha11 "penumbraprotos/penumbra/core/crypto/v1alpha1"
	v1alpha1 "penumbraprotos/penumbra/core/dex/v1alpha1"
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

// A Penumbra transparent SwapClaimProof.
type SwapClaimProof struct {
	// The swap being claimed
	SwapPlaintext *v1alpha1.SwapPlaintext `protobuf:"bytes,1,opt,name=swap_plaintext,json=swapPlaintext,proto3" json:"swap_plaintext,omitempty"`
	// Inclusion proof for the swap commitment
	SwapCommitmentProof *v1alpha11.StateCommitmentProof `protobuf:"bytes,4,opt,name=swap_commitment_proof,json=swapCommitmentProof,proto3" json:"swap_commitment_proof,omitempty"`
	// The nullifier key used to derive the swap nullifier
	Nk []byte `protobuf:"bytes,6,opt,name=nk,proto3" json:"nk,omitempty"`
	//*
	// @exclude
	// Describes output amounts
	Lambda_1I *v1alpha11.Amount `protobuf:"bytes,20,opt,name=lambda_1_i,json=lambda1I,proto3" json:"lambda_1_i,omitempty"`
	Lambda_2I *v1alpha11.Amount `protobuf:"bytes,21,opt,name=lambda_2_i,json=lambda2I,proto3" json:"lambda_2_i,omitempty"`
}

func (m *SwapClaimProof) Reset()         { *m = SwapClaimProof{} }
func (m *SwapClaimProof) String() string { return proto.CompactTextString(m) }
func (*SwapClaimProof) ProtoMessage()    {}
func (*SwapClaimProof) Descriptor() ([]byte, []int) {
	return fileDescriptor_1536b20e10cd99e5, []int{0}
}
func (m *SwapClaimProof) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SwapClaimProof) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SwapClaimProof.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SwapClaimProof) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SwapClaimProof.Merge(m, src)
}
func (m *SwapClaimProof) XXX_Size() int {
	return m.Size()
}
func (m *SwapClaimProof) XXX_DiscardUnknown() {
	xxx_messageInfo_SwapClaimProof.DiscardUnknown(m)
}

var xxx_messageInfo_SwapClaimProof proto.InternalMessageInfo

func (m *SwapClaimProof) GetSwapPlaintext() *v1alpha1.SwapPlaintext {
	if m != nil {
		return m.SwapPlaintext
	}
	return nil
}

func (m *SwapClaimProof) GetSwapCommitmentProof() *v1alpha11.StateCommitmentProof {
	if m != nil {
		return m.SwapCommitmentProof
	}
	return nil
}

func (m *SwapClaimProof) GetNk() []byte {
	if m != nil {
		return m.Nk
	}
	return nil
}

func (m *SwapClaimProof) GetLambda_1I() *v1alpha11.Amount {
	if m != nil {
		return m.Lambda_1I
	}
	return nil
}

func (m *SwapClaimProof) GetLambda_2I() *v1alpha11.Amount {
	if m != nil {
		return m.Lambda_2I
	}
	return nil
}

func init() {
	proto.RegisterType((*SwapClaimProof)(nil), "penumbra.core.transparent_proofs.v1alpha1.SwapClaimProof")
}

func init() {
	proto.RegisterFile("penumbra/core/transparent_proofs/v1alpha1/transparent_proofs.proto", fileDescriptor_1536b20e10cd99e5)
}

var fileDescriptor_1536b20e10cd99e5 = []byte{
	// 422 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xc1, 0xaa, 0xd3, 0x40,
	0x18, 0x85, 0x9b, 0x51, 0x2e, 0x32, 0x6a, 0x17, 0xd1, 0x2b, 0xe1, 0x2e, 0xc2, 0x45, 0x14, 0x7a,
	0x15, 0x13, 0xd2, 0xeb, 0x2a, 0xae, 0x4c, 0x16, 0xd2, 0x85, 0x30, 0xa4, 0xc5, 0x85, 0x04, 0xc2,
	0x34, 0x19, 0x35, 0x34, 0x93, 0x19, 0x92, 0xa9, 0xad, 0x6f, 0xe1, 0x33, 0xb8, 0xf4, 0x49, 0xc4,
	0x55, 0x97, 0x2e, 0x25, 0xc5, 0x8d, 0x2f, 0xe0, 0x56, 0x26, 0xe9, 0x34, 0x6d, 0x53, 0xb0, 0xb8,
	0x9c, 0xfc, 0xe7, 0x7c, 0xe7, 0xfc, 0x99, 0x81, 0x1e, 0x27, 0xf9, 0x9c, 0x4e, 0x0b, 0x6c, 0xc7,
	0xac, 0x20, 0xb6, 0x28, 0x70, 0x5e, 0x72, 0x5c, 0x90, 0x5c, 0x44, 0xbc, 0x60, 0xec, 0x5d, 0x69,
	0x7f, 0x74, 0x70, 0xc6, 0x3f, 0x60, 0xe7, 0xc8, 0xcc, 0xe2, 0x05, 0x13, 0x4c, 0xbf, 0x52, 0x0c,
	0x4b, 0x32, 0xac, 0x23, 0x3a, 0xc5, 0xb8, 0x78, 0xb2, 0x1f, 0x17, 0x17, 0x9f, 0xb8, 0x60, 0x6d,
	0x44, 0x73, 0x6e, 0xb0, 0x17, 0x8f, 0xf6, 0xb5, 0x09, 0x59, 0xb6, 0xc2, 0x84, 0x2c, 0x1b, 0xd5,
	0xc3, 0x5f, 0x00, 0xf6, 0xc7, 0x0b, 0xcc, 0xfd, 0x0c, 0xa7, 0x14, 0xc9, 0x38, 0x1d, 0xc1, 0x7e,
	0xb9, 0xc0, 0x3c, 0xe2, 0x19, 0x4e, 0x73, 0x41, 0x96, 0xc2, 0xd0, 0x2e, 0xb5, 0xc1, 0xed, 0xe1,
	0x95, 0xb5, 0x5f, 0x54, 0x42, 0x14, 0xd1, 0x92, 0x0c, 0xa4, 0x0c, 0xc1, 0xdd, 0x72, 0xf7, 0xa8,
	0xbf, 0x87, 0xe7, 0x35, 0x31, 0x66, 0x94, 0xa6, 0x82, 0x6e, 0x37, 0x33, 0x6e, 0xd6, 0xe0, 0xeb,
	0x03, 0xf0, 0x66, 0x8d, 0x96, 0x2d, 0xb0, 0x20, 0xfe, 0xd6, 0x5b, 0xb7, 0x0c, 0xee, 0x49, 0xe2,
	0xc1, 0x47, 0xbd, 0x0f, 0x41, 0x3e, 0x33, 0xce, 0x2e, 0xb5, 0xc1, 0x9d, 0x00, 0xe4, 0x33, 0xdd,
	0x87, 0x30, 0xc3, 0x74, 0x9a, 0xe0, 0xc8, 0x89, 0x52, 0xe3, 0x7e, 0x9d, 0xf6, 0xf8, 0x1f, 0x69,
	0x2f, 0x29, 0x9b, 0xe7, 0x22, 0xb8, 0xd5, 0x18, 0x9d, 0xd1, 0x0e, 0x64, 0x18, 0xa5, 0xc6, 0xf9,
	0x7f, 0x40, 0x86, 0x23, 0xef, 0x0f, 0xf8, 0x56, 0x99, 0xda, 0xaa, 0x32, 0xb5, 0x9f, 0x95, 0xa9,
	0x7d, 0x5e, 0x9b, 0xbd, 0xd5, 0xda, 0xec, 0xfd, 0x58, 0x9b, 0x3d, 0xf8, 0x2c, 0x66, 0xd4, 0x3a,
	0xf9, 0x0d, 0x78, 0x0f, 0x26, 0xed, 0xb0, 0xde, 0xba, 0x44, 0xf2, 0x26, 0x91, 0xf6, 0x76, 0xac,
	0x20, 0xf5, 0xd5, 0x96, 0xf6, 0xc9, 0x6f, 0xf3, 0x45, 0x77, 0xa6, 0x46, 0x5f, 0xc0, 0x0d, 0xe4,
	0x4f, 0xbe, 0x82, 0x01, 0x52, 0x05, 0x7d, 0x59, 0xb0, 0xd3, 0xc1, 0x7a, 0xb3, 0x31, 0x7c, 0x6f,
	0xa5, 0xa1, 0x94, 0x86, 0x1d, 0x69, 0xa8, 0xa4, 0x15, 0x78, 0x7e, 0xaa, 0x34, 0x7c, 0x85, 0xbc,
	0xd7, 0x44, 0xe0, 0x04, 0x0b, 0xfc, 0x1b, 0x3c, 0x55, 0x36, 0xd7, 0x95, 0x3e, 0xd7, 0xed, 0x18,
	0x5d, 0x57, 0x39, 0xa7, 0x67, 0xf5, 0xdf, 0xb8, 0xfe, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x1e, 0xb7,
	0x10, 0x9e, 0xab, 0x03, 0x00, 0x00,
}

func (m *SwapClaimProof) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SwapClaimProof) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SwapClaimProof) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Lambda_2I != nil {
		{
			size, err := m.Lambda_2I.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTransparentProofs(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0xaa
	}
	if m.Lambda_1I != nil {
		{
			size, err := m.Lambda_1I.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTransparentProofs(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0xa2
	}
	if len(m.Nk) > 0 {
		i -= len(m.Nk)
		copy(dAtA[i:], m.Nk)
		i = encodeVarintTransparentProofs(dAtA, i, uint64(len(m.Nk)))
		i--
		dAtA[i] = 0x32
	}
	if m.SwapCommitmentProof != nil {
		{
			size, err := m.SwapCommitmentProof.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTransparentProofs(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if m.SwapPlaintext != nil {
		{
			size, err := m.SwapPlaintext.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTransparentProofs(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTransparentProofs(dAtA []byte, offset int, v uint64) int {
	offset -= sovTransparentProofs(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *SwapClaimProof) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.SwapPlaintext != nil {
		l = m.SwapPlaintext.Size()
		n += 1 + l + sovTransparentProofs(uint64(l))
	}
	if m.SwapCommitmentProof != nil {
		l = m.SwapCommitmentProof.Size()
		n += 1 + l + sovTransparentProofs(uint64(l))
	}
	l = len(m.Nk)
	if l > 0 {
		n += 1 + l + sovTransparentProofs(uint64(l))
	}
	if m.Lambda_1I != nil {
		l = m.Lambda_1I.Size()
		n += 2 + l + sovTransparentProofs(uint64(l))
	}
	if m.Lambda_2I != nil {
		l = m.Lambda_2I.Size()
		n += 2 + l + sovTransparentProofs(uint64(l))
	}
	return n
}

func sovTransparentProofs(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTransparentProofs(x uint64) (n int) {
	return sovTransparentProofs(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SwapClaimProof) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTransparentProofs
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
			return fmt.Errorf("proto: SwapClaimProof: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SwapClaimProof: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SwapPlaintext", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransparentProofs
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
				return ErrInvalidLengthTransparentProofs
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTransparentProofs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.SwapPlaintext == nil {
				m.SwapPlaintext = &v1alpha1.SwapPlaintext{}
			}
			if err := m.SwapPlaintext.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SwapCommitmentProof", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransparentProofs
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
				return ErrInvalidLengthTransparentProofs
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTransparentProofs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.SwapCommitmentProof == nil {
				m.SwapCommitmentProof = &v1alpha11.StateCommitmentProof{}
			}
			if err := m.SwapCommitmentProof.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nk", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransparentProofs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTransparentProofs
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTransparentProofs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Nk = append(m.Nk[:0], dAtA[iNdEx:postIndex]...)
			if m.Nk == nil {
				m.Nk = []byte{}
			}
			iNdEx = postIndex
		case 20:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Lambda_1I", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransparentProofs
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
				return ErrInvalidLengthTransparentProofs
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTransparentProofs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Lambda_1I == nil {
				m.Lambda_1I = &v1alpha11.Amount{}
			}
			if err := m.Lambda_1I.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 21:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Lambda_2I", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransparentProofs
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
				return ErrInvalidLengthTransparentProofs
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTransparentProofs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Lambda_2I == nil {
				m.Lambda_2I = &v1alpha11.Amount{}
			}
			if err := m.Lambda_2I.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTransparentProofs(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTransparentProofs
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
func skipTransparentProofs(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTransparentProofs
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
					return 0, ErrIntOverflowTransparentProofs
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
					return 0, ErrIntOverflowTransparentProofs
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
				return 0, ErrInvalidLengthTransparentProofs
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTransparentProofs
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTransparentProofs
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTransparentProofs        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTransparentProofs          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTransparentProofs = fmt.Errorf("proto: unexpected end of group")
)
