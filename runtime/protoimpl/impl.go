// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package protoimpl contains the default implementation for messages
// generated by protoc-gen-go.
//
// WARNING: This package should only ever be imported by generated messages.
// The compatibility agreement covers nothing except for functionality needed
// to keep existing generated messages operational. Breakages that occur due
// to unauthorized usages of this package are not the author's responsibility.
package protoimpl

import (
	"google.golang.org/protobuf/internal/fileinit"
	"google.golang.org/protobuf/internal/impl"
)

const (
	// MaxVersion is the maximum supported version for generated .pb.go files;
	// which is the current version of the package.
	// This is incremented when the functionality of this package expands.
	MaxVersion = 0

	// MinVersion is the minimum supported version for generated .pb.go files.
	// This is incremented when the runtime drops support for old code.
	MinVersion = 0

	// Version is the current minor version of the runtime.
	Version = MaxVersion // v2.{Version}.x

	// TODO: Encode a date instead of the minor version?
)

type (
	// EnforceVersion is used by code generated by protoc-gen-go
	// to statically enforce minimum and maximum versions of this package.
	// A compilation failure implies either that:
	//	* the runtime package is too old and needs to be updated OR
	//	* the generated code is too old and needs to be regenerated.
	//
	// The runtime package can be upgraded by running:
	//	go get google.golang.org/protobuf
	//
	// The generated code can be regenerated by running:
	//	protoc --go_out=${PROTOC_GEN_GO_ARGS} ${PROTO_FILES}
	//
	// Example usage by generated code:
	//	const (
	//		// Verify that runtime/protoimpl is sufficiently up-to-date.
	//		_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - genVersion)
	//		// Verify that this generated code is sufficiently up-to-date.
	//		_ = protoimpl.EnforceVersion(genVersion - protoimpl.MinVersion)
	//	)
	//
	// The genVersion is the current version used to generated the code.
	// This compile-time check relies on negative integer overflow of a uint
	// being a compilation failure (guaranteed by the Go specification).
	EnforceVersion uint

	MessageInfo = impl.MessageInfo
	FileBuilder = fileinit.FileBuilder

	// TODO: Change these to more efficient data structures.
	ExtensionFields = map[int32]impl.ExtensionField
	UnknownFields   = []byte
	SizeCache       = int32

	ExtensionFieldV1 = impl.ExtensionField
)

var X impl.Export
