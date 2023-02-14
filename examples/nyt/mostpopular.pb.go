// Code generated by protoc-gen-go.
// source: mostpopular.proto
// DO NOT EDIT!

/*
Package nyt is a generated protocol buffer package.

It is generated from these files:
	mostpopular.proto
	semanticconcept.proto

It has these top-level messages:
	MostPopularResponse
	MostPopularResult
	SemanticConceptResponse
	SemanticConceptResult
	SemanticConceptArticleList
	SemanticConceptArticle
*/
package nyt

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type MostPopularResponse struct {
	Status     string               `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Copyright  string               `protobuf:"bytes,2,opt,name=copyright" json:"copyright,omitempty"`
	NumResults uint32               `protobuf:"varint,3,opt,name=num_results,json=numResults" json:"num_results,omitempty"`
	Results    []*MostPopularResult `protobuf:"bytes,4,rep,name=results" json:"results,omitempty"`
}

func (m *MostPopularResponse) Reset()                    { *m = MostPopularResponse{} }
func (m *MostPopularResponse) String() string            { return proto.CompactTextString(m) }
func (*MostPopularResponse) ProtoMessage()               {}
func (*MostPopularResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *MostPopularResponse) GetResults() []*MostPopularResult {
	if m != nil {
		return m.Results
	}
	return nil
}

type MostPopularResult struct {
	Url           string `protobuf:"bytes,1,opt,name=url" json:"url,omitempty"`
	Id            uint64 `protobuf:"varint,2,opt,name=id" json:"id,omitempty"`
	AsssetId      uint64 `protobuf:"varint,3,opt,name=assset_id,json=asssetId" json:"assset_id,omitempty"`
	AdxKeywords   string `protobuf:"bytes,4,opt,name=adx_keywords,json=adxKeywords" json:"adx_keywords,omitempty"`
	Column        string `protobuf:"bytes,5,opt,name=column" json:"column,omitempty"`
	Section       string `protobuf:"bytes,6,opt,name=section" json:"section,omitempty"`
	Byline        string `protobuf:"bytes,7,opt,name=byline" json:"byline,omitempty"`
	Type          string `protobuf:"bytes,8,opt,name=type" json:"type,omitempty"`
	Title         string `protobuf:"bytes,9,opt,name=title" json:"title,omitempty"`
	Abstract      string `protobuf:"bytes,10,opt,name=abstract" json:"abstract,omitempty"`
	PublishedDate string `protobuf:"bytes,11,opt,name=published_date,json=publishedDate" json:"published_date,omitempty"`
	Source        string `protobuf:"bytes,12,opt,name=source" json:"source,omitempty"`
	Views         uint32 `protobuf:"varint,13,opt,name=views" json:"views,omitempty"`
}

func (m *MostPopularResult) Reset()                    { *m = MostPopularResult{} }
func (m *MostPopularResult) String() string            { return proto.CompactTextString(m) }
func (*MostPopularResult) ProtoMessage()               {}
func (*MostPopularResult) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto.RegisterType((*MostPopularResponse)(nil), "nyt.MostPopularResponse")
	proto.RegisterType((*MostPopularResult)(nil), "nyt.MostPopularResult")
}

var fileDescriptor0 = []byte{
	// 346 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x64, 0x92, 0xed, 0x4a, 0xf3, 0x30,
	0x14, 0xc7, 0xd9, 0xba, 0xb7, 0x9e, 0x6e, 0xe3, 0xd9, 0x79, 0x64, 0x04, 0x15, 0x9c, 0x03, 0x61,
	0x9f, 0x86, 0xe8, 0x2d, 0xf8, 0x45, 0x44, 0x90, 0xde, 0x40, 0xe9, 0x4b, 0x70, 0xc5, 0xae, 0x09,
	0xc9, 0x89, 0xb3, 0x37, 0xe3, 0x85, 0x79, 0x35, 0xa6, 0x49, 0x37, 0x11, 0xbf, 0x9d, 0xff, 0xef,
	0x97, 0x6c, 0x7f, 0x72, 0x0a, 0x8b, 0xbd, 0xd0, 0x24, 0x85, 0x34, 0x55, 0xaa, 0xb6, 0x52, 0x09,
	0x12, 0x18, 0xd4, 0x0d, 0xad, 0x3f, 0x7b, 0xf0, 0xff, 0xd9, 0xaa, 0x17, 0xaf, 0x62, 0xae, 0xa5,
	0xa8, 0x35, 0xc7, 0x25, 0x8c, 0x34, 0xa5, 0x64, 0x34, 0xeb, 0xad, 0x7a, 0x9b, 0x30, 0xee, 0x12,
	0x5e, 0x42, 0x98, 0x0b, 0xd9, 0xa8, 0xf2, 0x75, 0x47, 0xac, 0xef, 0xd4, 0x0f, 0xc0, 0x2b, 0x88,
	0x6a, 0xb3, 0x4f, 0x14, 0xd7, 0xa6, 0x22, 0xcd, 0x02, 0xeb, 0x67, 0x31, 0x58, 0x14, 0x7b, 0x82,
	0xb7, 0x30, 0x3e, 0xca, 0xc1, 0x2a, 0xd8, 0x44, 0x77, 0xcb, 0xad, 0x6d, 0xb1, 0xfd, 0xdd, 0xc0,
	0xea, 0xf8, 0x78, 0x6c, 0xfd, 0xd5, 0x87, 0xc5, 0x1f, 0x8d, 0xff, 0x20, 0x30, 0xaa, 0xea, 0xba,
	0xb5, 0x23, 0xce, 0xa1, 0x5f, 0x16, 0xae, 0xd1, 0x20, 0xb6, 0x13, 0x5e, 0x40, 0x98, 0x6a, 0xad,
	0x39, 0x25, 0x16, 0x07, 0x0e, 0x4f, 0x3c, 0x78, 0x2c, 0xf0, 0x1a, 0xa6, 0x69, 0xf1, 0x91, 0xbc,
	0xf1, 0xe6, 0x20, 0x54, 0xd1, 0x76, 0x69, 0x7f, 0x27, 0xb2, 0xec, 0xa9, 0x43, 0xed, 0x03, 0xe4,
	0xa2, 0x32, 0xfb, 0x9a, 0x0d, 0xfd, 0x03, 0xf8, 0x84, 0x0c, 0xc6, 0x9a, 0xe7, 0x54, 0x8a, 0x9a,
	0x8d, 0x9c, 0x38, 0xc6, 0xf6, 0x46, 0xd6, 0x54, 0x65, 0xcd, 0xd9, 0xd8, 0xdf, 0xf0, 0x09, 0x11,
	0x06, 0xd4, 0x48, 0xce, 0x26, 0x8e, 0xba, 0x19, 0xcf, 0x60, 0x48, 0x25, 0x55, 0x9c, 0x85, 0x0e,
	0xfa, 0x80, 0xe7, 0x30, 0x49, 0x33, 0x4d, 0x2a, 0xcd, 0x89, 0x81, 0x13, 0xa7, 0x8c, 0x37, 0x30,
	0x97, 0x26, 0xab, 0x4a, 0xbd, 0xe3, 0x45, 0x52, 0xa4, 0xc4, 0x59, 0xe4, 0x4e, 0xcc, 0x4e, 0xf4,
	0xc1, 0x42, 0xb7, 0x37, 0x61, 0x54, 0xce, 0xd9, 0xb4, 0xdb, 0x9b, 0x4b, 0xed, 0x1f, 0xbe, 0x97,
	0xfc, 0xa0, 0xd9, 0xcc, 0xed, 0xc4, 0x87, 0x6c, 0xe4, 0xbe, 0x84, 0xfb, 0xef, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xdb, 0xe0, 0x3a, 0x3e, 0x1e, 0x02, 0x00, 0x00,
}