// Copyright (c) 2017 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Code generated by thriftrw v1.10.0. DO NOT EDIT.
// @generated

package matching

import (
	"errors"
	"fmt"
	"github.com/uber/cadence/.gen/go/shared"
	"go.uber.org/thriftrw/wire"
	"strings"
)

// MatchingService_PollForActivityTask_Args represents the arguments for the MatchingService.PollForActivityTask function.
//
// The arguments for PollForActivityTask are sent and received over the wire as this struct.
type MatchingService_PollForActivityTask_Args struct {
	PollRequest *PollForActivityTaskRequest `json:"pollRequest,omitempty"`
}

// ToWire translates a MatchingService_PollForActivityTask_Args struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//   x, err := v.ToWire()
//   if err != nil {
//     return err
//   }
//
//   if err := binaryProtocol.Encode(x, writer); err != nil {
//     return err
//   }
func (v *MatchingService_PollForActivityTask_Args) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	if v.PollRequest != nil {
		w, err = v.PollRequest.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _PollForActivityTaskRequest_1_Read(w wire.Value) (*PollForActivityTaskRequest, error) {
	var v PollForActivityTaskRequest
	err := v.FromWire(w)
	return &v, err
}

// FromWire deserializes a MatchingService_PollForActivityTask_Args struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a MatchingService_PollForActivityTask_Args struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v MatchingService_PollForActivityTask_Args
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *MatchingService_PollForActivityTask_Args) FromWire(w wire.Value) error {
	var err error

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TStruct {
				v.PollRequest, err = _PollForActivityTaskRequest_1_Read(field.Value)
				if err != nil {
					return err
				}

			}
		}
	}

	return nil
}

// String returns a readable string representation of a MatchingService_PollForActivityTask_Args
// struct.
func (v *MatchingService_PollForActivityTask_Args) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [1]string
	i := 0
	if v.PollRequest != nil {
		fields[i] = fmt.Sprintf("PollRequest: %v", v.PollRequest)
		i++
	}

	return fmt.Sprintf("MatchingService_PollForActivityTask_Args{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this MatchingService_PollForActivityTask_Args match the
// provided MatchingService_PollForActivityTask_Args.
//
// This function performs a deep comparison.
func (v *MatchingService_PollForActivityTask_Args) Equals(rhs *MatchingService_PollForActivityTask_Args) bool {
	if !((v.PollRequest == nil && rhs.PollRequest == nil) || (v.PollRequest != nil && rhs.PollRequest != nil && v.PollRequest.Equals(rhs.PollRequest))) {
		return false
	}

	return true
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the arguments.
//
// This will always be "PollForActivityTask" for this struct.
func (v *MatchingService_PollForActivityTask_Args) MethodName() string {
	return "PollForActivityTask"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Call for this struct.
func (v *MatchingService_PollForActivityTask_Args) EnvelopeType() wire.EnvelopeType {
	return wire.Call
}

// MatchingService_PollForActivityTask_Helper provides functions that aid in handling the
// parameters and return values of the MatchingService.PollForActivityTask
// function.
var MatchingService_PollForActivityTask_Helper = struct {
	// Args accepts the parameters of PollForActivityTask in-order and returns
	// the arguments struct for the function.
	Args func(
		pollRequest *PollForActivityTaskRequest,
	) *MatchingService_PollForActivityTask_Args

	// IsException returns true if the given error can be thrown
	// by PollForActivityTask.
	//
	// An error can be thrown by PollForActivityTask only if the
	// corresponding exception type was mentioned in the 'throws'
	// section for it in the Thrift file.
	IsException func(error) bool

	// WrapResponse returns the result struct for PollForActivityTask
	// given its return value and error.
	//
	// This allows mapping values and errors returned by
	// PollForActivityTask into a serializable result struct.
	// WrapResponse returns a non-nil error if the provided
	// error cannot be thrown by PollForActivityTask
	//
	//   value, err := PollForActivityTask(args)
	//   result, err := MatchingService_PollForActivityTask_Helper.WrapResponse(value, err)
	//   if err != nil {
	//     return fmt.Errorf("unexpected error from PollForActivityTask: %v", err)
	//   }
	//   serialize(result)
	WrapResponse func(*shared.PollForActivityTaskResponse, error) (*MatchingService_PollForActivityTask_Result, error)

	// UnwrapResponse takes the result struct for PollForActivityTask
	// and returns the value or error returned by it.
	//
	// The error is non-nil only if PollForActivityTask threw an
	// exception.
	//
	//   result := deserialize(bytes)
	//   value, err := MatchingService_PollForActivityTask_Helper.UnwrapResponse(result)
	UnwrapResponse func(*MatchingService_PollForActivityTask_Result) (*shared.PollForActivityTaskResponse, error)
}{}

func init() {
	MatchingService_PollForActivityTask_Helper.Args = func(
		pollRequest *PollForActivityTaskRequest,
	) *MatchingService_PollForActivityTask_Args {
		return &MatchingService_PollForActivityTask_Args{
			PollRequest: pollRequest,
		}
	}

	MatchingService_PollForActivityTask_Helper.IsException = func(err error) bool {
		switch err.(type) {
		case *shared.BadRequestError:
			return true
		case *shared.InternalServiceError:
			return true
		default:
			return false
		}
	}

	MatchingService_PollForActivityTask_Helper.WrapResponse = func(success *shared.PollForActivityTaskResponse, err error) (*MatchingService_PollForActivityTask_Result, error) {
		if err == nil {
			return &MatchingService_PollForActivityTask_Result{Success: success}, nil
		}

		switch e := err.(type) {
		case *shared.BadRequestError:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for MatchingService_PollForActivityTask_Result.BadRequestError")
			}
			return &MatchingService_PollForActivityTask_Result{BadRequestError: e}, nil
		case *shared.InternalServiceError:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for MatchingService_PollForActivityTask_Result.InternalServiceError")
			}
			return &MatchingService_PollForActivityTask_Result{InternalServiceError: e}, nil
		}

		return nil, err
	}
	MatchingService_PollForActivityTask_Helper.UnwrapResponse = func(result *MatchingService_PollForActivityTask_Result) (success *shared.PollForActivityTaskResponse, err error) {
		if result.BadRequestError != nil {
			err = result.BadRequestError
			return
		}
		if result.InternalServiceError != nil {
			err = result.InternalServiceError
			return
		}

		if result.Success != nil {
			success = result.Success
			return
		}

		err = errors.New("expected a non-void result")
		return
	}

}

// MatchingService_PollForActivityTask_Result represents the result of a MatchingService.PollForActivityTask function call.
//
// The result of a PollForActivityTask execution is sent and received over the wire as this struct.
//
// Success is set only if the function did not throw an exception.
type MatchingService_PollForActivityTask_Result struct {
	// Value returned by PollForActivityTask after a successful execution.
	Success              *shared.PollForActivityTaskResponse `json:"success,omitempty"`
	BadRequestError      *shared.BadRequestError             `json:"badRequestError,omitempty"`
	InternalServiceError *shared.InternalServiceError        `json:"internalServiceError,omitempty"`
}

// ToWire translates a MatchingService_PollForActivityTask_Result struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//   x, err := v.ToWire()
//   if err != nil {
//     return err
//   }
//
//   if err := binaryProtocol.Encode(x, writer); err != nil {
//     return err
//   }
func (v *MatchingService_PollForActivityTask_Result) ToWire() (wire.Value, error) {
	var (
		fields [3]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	if v.Success != nil {
		w, err = v.Success.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 0, Value: w}
		i++
	}
	if v.BadRequestError != nil {
		w, err = v.BadRequestError.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}
	if v.InternalServiceError != nil {
		w, err = v.InternalServiceError.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 2, Value: w}
		i++
	}

	if i != 1 {
		return wire.Value{}, fmt.Errorf("MatchingService_PollForActivityTask_Result should have exactly one field: got %v fields", i)
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _PollForActivityTaskResponse_Read(w wire.Value) (*shared.PollForActivityTaskResponse, error) {
	var v shared.PollForActivityTaskResponse
	err := v.FromWire(w)
	return &v, err
}

// FromWire deserializes a MatchingService_PollForActivityTask_Result struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a MatchingService_PollForActivityTask_Result struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v MatchingService_PollForActivityTask_Result
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *MatchingService_PollForActivityTask_Result) FromWire(w wire.Value) error {
	var err error

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 0:
			if field.Value.Type() == wire.TStruct {
				v.Success, err = _PollForActivityTaskResponse_Read(field.Value)
				if err != nil {
					return err
				}

			}
		case 1:
			if field.Value.Type() == wire.TStruct {
				v.BadRequestError, err = _BadRequestError_Read(field.Value)
				if err != nil {
					return err
				}

			}
		case 2:
			if field.Value.Type() == wire.TStruct {
				v.InternalServiceError, err = _InternalServiceError_Read(field.Value)
				if err != nil {
					return err
				}

			}
		}
	}

	count := 0
	if v.Success != nil {
		count++
	}
	if v.BadRequestError != nil {
		count++
	}
	if v.InternalServiceError != nil {
		count++
	}
	if count != 1 {
		return fmt.Errorf("MatchingService_PollForActivityTask_Result should have exactly one field: got %v fields", count)
	}

	return nil
}

// String returns a readable string representation of a MatchingService_PollForActivityTask_Result
// struct.
func (v *MatchingService_PollForActivityTask_Result) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [3]string
	i := 0
	if v.Success != nil {
		fields[i] = fmt.Sprintf("Success: %v", v.Success)
		i++
	}
	if v.BadRequestError != nil {
		fields[i] = fmt.Sprintf("BadRequestError: %v", v.BadRequestError)
		i++
	}
	if v.InternalServiceError != nil {
		fields[i] = fmt.Sprintf("InternalServiceError: %v", v.InternalServiceError)
		i++
	}

	return fmt.Sprintf("MatchingService_PollForActivityTask_Result{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this MatchingService_PollForActivityTask_Result match the
// provided MatchingService_PollForActivityTask_Result.
//
// This function performs a deep comparison.
func (v *MatchingService_PollForActivityTask_Result) Equals(rhs *MatchingService_PollForActivityTask_Result) bool {
	if !((v.Success == nil && rhs.Success == nil) || (v.Success != nil && rhs.Success != nil && v.Success.Equals(rhs.Success))) {
		return false
	}
	if !((v.BadRequestError == nil && rhs.BadRequestError == nil) || (v.BadRequestError != nil && rhs.BadRequestError != nil && v.BadRequestError.Equals(rhs.BadRequestError))) {
		return false
	}
	if !((v.InternalServiceError == nil && rhs.InternalServiceError == nil) || (v.InternalServiceError != nil && rhs.InternalServiceError != nil && v.InternalServiceError.Equals(rhs.InternalServiceError))) {
		return false
	}

	return true
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the result.
//
// This will always be "PollForActivityTask" for this struct.
func (v *MatchingService_PollForActivityTask_Result) MethodName() string {
	return "PollForActivityTask"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Reply for this struct.
func (v *MatchingService_PollForActivityTask_Result) EnvelopeType() wire.EnvelopeType {
	return wire.Reply
}
