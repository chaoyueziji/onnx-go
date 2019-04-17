package onnxtest

// this file is auto-generated... DO NOT EDIT

import (
	"github.com/owulveryck/onnx-go/backend/testbackend"
	"gorgonia.org/tensor"
)

func init() {
	testbackend.Register("Div", "TestDivExample", NewTestDivExample)
}

// NewTestDivExample version: 3.
func NewTestDivExample() *testbackend.TestCase {
	return &testbackend.TestCase{
		OpType: "Div",
		Title:  "TestDivExample",
		ModelB: []byte{0x8, 0x3, 0x12, 0xc, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x3a, 0x55, 0xa, 0xe, 0xa, 0x1, 0x78, 0xa, 0x1, 0x79, 0x12, 0x1, 0x7a, 0x22, 0x3, 0x44, 0x69, 0x76, 0x12, 0x10, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x64, 0x69, 0x76, 0x5f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5a, 0xf, 0xa, 0x1, 0x78, 0x12, 0xa, 0xa, 0x8, 0x8, 0x1, 0x12, 0x4, 0xa, 0x2, 0x8, 0x2, 0x5a, 0xf, 0xa, 0x1, 0x79, 0x12, 0xa, 0xa, 0x8, 0x8, 0x1, 0x12, 0x4, 0xa, 0x2, 0x8, 0x2, 0x62, 0xf, 0xa, 0x1, 0x7a, 0x12, 0xa, 0xa, 0x8, 0x8, 0x1, 0x12, 0x4, 0xa, 0x2, 0x8, 0x2, 0x42, 0x2, 0x10, 0x9},

		/*

		   &pb.NodeProto{
		     Input:     []string{"x", "y"},
		     Output:    []string{"z"},
		     Name:      "",
		     OpType:    "Div",
		     Attributes: ([]*pb.AttributeProto) <nil>
		   ,
		   },


		*/

		Input: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(2),
				tensor.WithBacking([]float32{3, 4}),
			),

			tensor.New(
				tensor.WithShape(2),
				tensor.WithBacking([]float32{1, 2}),
			),
		},
		ExpectedOutput: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(2),
				tensor.WithBacking([]float32{3, 2}),
			),
		},
	}
}
