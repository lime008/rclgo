/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package example_interfaces_msg
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	primitives "github.com/tiiuae/rclgo/pkg/rclgo/primitives"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lexample_interfaces__rosidl_typesupport_c -lexample_interfaces__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <example_interfaces/msg/u_int32_multi_array.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("example_interfaces/UInt32MultiArray", UInt32MultiArrayTypeSupport)
}

// Do not create instances of this type directly. Always use NewUInt32MultiArray
// function instead.
type UInt32MultiArray struct {
	Layout MultiArrayLayout `yaml:"layout"`// specification of data layout
	Data []uint32 `yaml:"data"`// array of data
}

// NewUInt32MultiArray creates a new UInt32MultiArray with default values.
func NewUInt32MultiArray() *UInt32MultiArray {
	self := UInt32MultiArray{}
	self.SetDefaults()
	return &self
}

func (t *UInt32MultiArray) Clone() *UInt32MultiArray {
	c := &UInt32MultiArray{}
	c.Layout = *t.Layout.Clone()
	if t.Data != nil {
		c.Data = make([]uint32, len(t.Data))
		copy(c.Data, t.Data)
	}
	return c
}

func (t *UInt32MultiArray) CloneMsg() types.Message {
	return t.Clone()
}

func (t *UInt32MultiArray) SetDefaults() {
	t.Layout.SetDefaults()
	
}

// CloneUInt32MultiArraySlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneUInt32MultiArraySlice(dst, src []UInt32MultiArray) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var UInt32MultiArrayTypeSupport types.MessageTypeSupport = _UInt32MultiArrayTypeSupport{}

type _UInt32MultiArrayTypeSupport struct{}

func (t _UInt32MultiArrayTypeSupport) New() types.Message {
	return NewUInt32MultiArray()
}

func (t _UInt32MultiArrayTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.example_interfaces__msg__UInt32MultiArray
	return (unsafe.Pointer)(C.example_interfaces__msg__UInt32MultiArray__create())
}

func (t _UInt32MultiArrayTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.example_interfaces__msg__UInt32MultiArray__destroy((*C.example_interfaces__msg__UInt32MultiArray)(pointer_to_free))
}

func (t _UInt32MultiArrayTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*UInt32MultiArray)
	mem := (*C.example_interfaces__msg__UInt32MultiArray)(dst)
	MultiArrayLayoutTypeSupport.AsCStruct(unsafe.Pointer(&mem.layout), &m.Layout)
	primitives.Uint32__Sequence_to_C((*primitives.CUint32__Sequence)(unsafe.Pointer(&mem.data)), m.Data)
}

func (t _UInt32MultiArrayTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*UInt32MultiArray)
	mem := (*C.example_interfaces__msg__UInt32MultiArray)(ros2_message_buffer)
	MultiArrayLayoutTypeSupport.AsGoStruct(&m.Layout, unsafe.Pointer(&mem.layout))
	primitives.Uint32__Sequence_to_Go(&m.Data, *(*primitives.CUint32__Sequence)(unsafe.Pointer(&mem.data)))
}

func (t _UInt32MultiArrayTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__example_interfaces__msg__UInt32MultiArray())
}

type CUInt32MultiArray = C.example_interfaces__msg__UInt32MultiArray
type CUInt32MultiArray__Sequence = C.example_interfaces__msg__UInt32MultiArray__Sequence

func UInt32MultiArray__Sequence_to_Go(goSlice *[]UInt32MultiArray, cSlice CUInt32MultiArray__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]UInt32MultiArray, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.example_interfaces__msg__UInt32MultiArray__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_example_interfaces__msg__UInt32MultiArray * uintptr(i)),
		))
		UInt32MultiArrayTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func UInt32MultiArray__Sequence_to_C(cSlice *CUInt32MultiArray__Sequence, goSlice []UInt32MultiArray) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.example_interfaces__msg__UInt32MultiArray)(C.malloc((C.size_t)(C.sizeof_struct_example_interfaces__msg__UInt32MultiArray * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.example_interfaces__msg__UInt32MultiArray)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_example_interfaces__msg__UInt32MultiArray * uintptr(i)),
		))
		UInt32MultiArrayTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func UInt32MultiArray__Array_to_Go(goSlice []UInt32MultiArray, cSlice []CUInt32MultiArray) {
	for i := 0; i < len(cSlice); i++ {
		UInt32MultiArrayTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func UInt32MultiArray__Array_to_C(cSlice []CUInt32MultiArray, goSlice []UInt32MultiArray) {
	for i := 0; i < len(goSlice); i++ {
		UInt32MultiArrayTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
