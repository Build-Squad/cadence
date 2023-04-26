/*
 * Cadence - The resource-oriented smart contract programming language
 *
 * Copyright Dapper Labs, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package interpreter

import (
	"github.com/onflow/atree"

	"github.com/onflow/cadence/runtime/common"
	"github.com/onflow/cadence/runtime/errors"
	"github.com/onflow/cadence/runtime/format"
	"github.com/onflow/cadence/runtime/sema"
)

// AccountCapabilityControllerValue

type AccountCapabilityControllerValue struct {
	BorrowType   StaticType
	CapabilityID UInt64Value
}

func NewUnmeteredAccountCapabilityControllerValue(
	staticType StaticType,
	capabilityID UInt64Value,
) *AccountCapabilityControllerValue {
	return &AccountCapabilityControllerValue{
		BorrowType:   staticType,
		CapabilityID: capabilityID,
	}
}

func NewAccountCapabilityControllerValue(
	memoryGauge common.MemoryGauge,
	staticType StaticType,
	capabilityID UInt64Value,
) *AccountCapabilityControllerValue {
	// Constant because its constituents are already metered.
	common.UseMemory(memoryGauge, common.AccountCapabilityControllerValueMemoryUsage)
	return NewUnmeteredAccountCapabilityControllerValue(
		staticType,
		capabilityID,
	)
}

var _ Value = &AccountCapabilityControllerValue{}
var _ atree.Value = &AccountCapabilityControllerValue{}
var _ EquatableValue = &AccountCapabilityControllerValue{}
var _ CapabilityControllerValue = &AccountCapabilityControllerValue{}
var _ MemberAccessibleValue = &AccountCapabilityControllerValue{}

func (*AccountCapabilityControllerValue) isValue() {}

func (*AccountCapabilityControllerValue) isCapabilityControllerValue() {}

func (v *AccountCapabilityControllerValue) Accept(interpreter *Interpreter, visitor Visitor) {
	visitor.VisitAccountCapabilityControllerValue(interpreter, v)
}

func (v *AccountCapabilityControllerValue) Walk(_ *Interpreter, walkChild func(Value)) {
	walkChild(v.CapabilityID)
}

func (v *AccountCapabilityControllerValue) StaticType(_ *Interpreter) StaticType {
	return PrimitiveStaticTypeAccountCapabilityController
}

func (*AccountCapabilityControllerValue) IsImportable(_ *Interpreter) bool {
	return false
}

func (v *AccountCapabilityControllerValue) String() string {
	return v.RecursiveString(SeenReferences{})
}

func (v *AccountCapabilityControllerValue) RecursiveString(seenReferences SeenReferences) string {
	return format.AccountCapabilityController(
		v.BorrowType.String(),
		v.CapabilityID.RecursiveString(seenReferences),
	)
}

func (v *AccountCapabilityControllerValue) MeteredString(memoryGauge common.MemoryGauge, seenReferences SeenReferences) string {
	common.UseMemory(memoryGauge, common.AccountCapabilityControllerValueStringMemoryUsage)

	return format.AccountCapabilityController(
		v.BorrowType.MeteredString(memoryGauge),
		v.CapabilityID.MeteredString(memoryGauge, seenReferences),
	)
}

func (v *AccountCapabilityControllerValue) ConformsToStaticType(
	_ *Interpreter,
	_ LocationRange,
	_ TypeConformanceResults,
) bool {
	return true
}

func (v *AccountCapabilityControllerValue) Equal(interpreter *Interpreter, locationRange LocationRange, other Value) bool {
	otherController, ok := other.(*AccountCapabilityControllerValue)
	if !ok {
		return false
	}

	return otherController.BorrowType.Equal(v.BorrowType) &&
		otherController.CapabilityID.Equal(interpreter, locationRange, v.CapabilityID)
}

func (*AccountCapabilityControllerValue) IsStorable() bool {
	return true
}

func (v *AccountCapabilityControllerValue) Storable(storage atree.SlabStorage, address atree.Address, maxInlineSize uint64) (atree.Storable, error) {
	return maybeLargeImmutableStorable(v, storage, address, maxInlineSize)
}

func (*AccountCapabilityControllerValue) NeedsStoreTo(_ atree.Address) bool {
	return false
}

func (*AccountCapabilityControllerValue) IsResourceKinded(_ *Interpreter) bool {
	return false
}

func (v *AccountCapabilityControllerValue) Transfer(
	interpreter *Interpreter,
	_ LocationRange,
	_ atree.Address,
	remove bool,
	storable atree.Storable,
) Value {
	if remove {
		interpreter.RemoveReferencedSlab(storable)
	}
	return v
}

func (v *AccountCapabilityControllerValue) Clone(interpreter *Interpreter) Value {
	return &AccountCapabilityControllerValue{
		BorrowType:   v.BorrowType,
		CapabilityID: v.CapabilityID,
	}
}

func (v *AccountCapabilityControllerValue) DeepRemove(_ *Interpreter) {
	// NO-OP
}

func (v *AccountCapabilityControllerValue) ByteSize() uint32 {
	return mustStorableSize(v)
}

func (v *AccountCapabilityControllerValue) StoredValue(_ atree.SlabStorage) (atree.Value, error) {
	return v, nil
}

func (v *AccountCapabilityControllerValue) ChildStorables() []atree.Storable {
	return []atree.Storable{
		v.CapabilityID,
	}
}

func (v *AccountCapabilityControllerValue) GetMember(inter *Interpreter, _ LocationRange, name string) Value {

	// TODO: sema.AccountCapabilityControllerTypeDeleteFunctionName

	switch name {
	case sema.AccountCapabilityControllerTypeCapabilityIDFieldName:
		return v.CapabilityID

	case sema.AccountCapabilityControllerTypeBorrowTypeFieldName:
		return NewTypeValue(inter, v.BorrowType)

	}

	return nil
}

func (*AccountCapabilityControllerValue) RemoveMember(_ *Interpreter, _ LocationRange, _ string) Value {
	// Storage capability controllers have no removable members (fields / functions)
	panic(errors.NewUnreachableError())
}

func (*AccountCapabilityControllerValue) SetMember(_ *Interpreter, _ LocationRange, _ string, _ Value) bool {
	// Storage capability controllers have no settable members (fields / functions)
	panic(errors.NewUnreachableError())
}
