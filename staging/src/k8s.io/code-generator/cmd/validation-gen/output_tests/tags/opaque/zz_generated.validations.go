//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by validation-gen. DO NOT EDIT.

package opaque

import (
	context "context"
	fmt "fmt"

	equality "k8s.io/apimachinery/pkg/api/equality"
	operation "k8s.io/apimachinery/pkg/api/operation"
	safe "k8s.io/apimachinery/pkg/api/safe"
	validate "k8s.io/apimachinery/pkg/api/validate"
	field "k8s.io/apimachinery/pkg/util/validation/field"
	testscheme "k8s.io/code-generator/cmd/validation-gen/testscheme"
)

func init() { localSchemeBuilder.Register(RegisterValidations) }

// RegisterValidations adds validation functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterValidations(scheme *testscheme.Scheme) error {
	scheme.AddValidationFunc((*OtherString)(nil), func(ctx context.Context, op operation.Operation, obj, oldObj interface{}) field.ErrorList {
		switch op.Request.SubresourcePath() {
		case "/":
			return Validate_OtherString(ctx, op, nil /* fldPath */, obj.(*OtherString), safe.Cast[*OtherString](oldObj))
		}
		return field.ErrorList{field.InternalError(nil, fmt.Errorf("no validation found for %T, subresource: %v", obj, op.Request.SubresourcePath()))}
	})
	scheme.AddValidationFunc((*OtherStruct)(nil), func(ctx context.Context, op operation.Operation, obj, oldObj interface{}) field.ErrorList {
		switch op.Request.SubresourcePath() {
		case "/":
			return Validate_OtherStruct(ctx, op, nil /* fldPath */, obj.(*OtherStruct), safe.Cast[*OtherStruct](oldObj))
		}
		return field.ErrorList{field.InternalError(nil, fmt.Errorf("no validation found for %T, subresource: %v", obj, op.Request.SubresourcePath()))}
	})
	scheme.AddValidationFunc((*Struct)(nil), func(ctx context.Context, op operation.Operation, obj, oldObj interface{}) field.ErrorList {
		switch op.Request.SubresourcePath() {
		case "/":
			return Validate_Struct(ctx, op, nil /* fldPath */, obj.(*Struct), safe.Cast[*Struct](oldObj))
		}
		return field.ErrorList{field.InternalError(nil, fmt.Errorf("no validation found for %T, subresource: %v", obj, op.Request.SubresourcePath()))}
	})
	scheme.AddValidationFunc((TypedefMapOther)(nil), func(ctx context.Context, op operation.Operation, obj, oldObj interface{}) field.ErrorList {
		switch op.Request.SubresourcePath() {
		case "/":
			return Validate_TypedefMapOther(ctx, op, nil /* fldPath */, obj.(TypedefMapOther), safe.Cast[TypedefMapOther](oldObj))
		}
		return field.ErrorList{field.InternalError(nil, fmt.Errorf("no validation found for %T, subresource: %v", obj, op.Request.SubresourcePath()))}
	})
	scheme.AddValidationFunc((TypedefSliceOther)(nil), func(ctx context.Context, op operation.Operation, obj, oldObj interface{}) field.ErrorList {
		switch op.Request.SubresourcePath() {
		case "/":
			return Validate_TypedefSliceOther(ctx, op, nil /* fldPath */, obj.(TypedefSliceOther), safe.Cast[TypedefSliceOther](oldObj))
		}
		return field.ErrorList{field.InternalError(nil, fmt.Errorf("no validation found for %T, subresource: %v", obj, op.Request.SubresourcePath()))}
	})
	return nil
}

func Validate_OtherString(ctx context.Context, op operation.Operation, fldPath *field.Path, obj, oldObj *OtherString) (errs field.ErrorList) {
	// type OtherString
	if op.Type == operation.Update && (obj == oldObj || (obj != nil && oldObj != nil && *obj == *oldObj)) {
		return nil // no changes
	}
	errs = append(errs, validate.FixedResult(ctx, op, fldPath, obj, oldObj, false, "type OtherString")...)

	return errs
}

func Validate_OtherStruct(ctx context.Context, op operation.Operation, fldPath *field.Path, obj, oldObj *OtherStruct) (errs field.ErrorList) {
	// type OtherStruct
	if op.Type == operation.Update && (obj == oldObj || (obj != nil && oldObj != nil && *obj == *oldObj)) {
		return nil // no changes
	}
	errs = append(errs, validate.FixedResult(ctx, op, fldPath, obj, oldObj, false, "type OtherStruct")...)

	// field OtherStruct.StringField
	errs = append(errs,
		func(fldPath *field.Path, obj, oldObj *string) (errs field.ErrorList) {
			if op.Type == operation.Update && (obj == oldObj || (obj != nil && oldObj != nil && *obj == *oldObj)) {
				return nil // no changes
			}
			errs = append(errs, validate.FixedResult(ctx, op, fldPath, obj, oldObj, false, "field OtherStruct.StringField")...)
			return
		}(fldPath.Child("stringField"), &obj.StringField, safe.Field(oldObj, func(oldObj *OtherStruct) *string { return &oldObj.StringField }))...)

	return errs
}

func Validate_Struct(ctx context.Context, op operation.Operation, fldPath *field.Path, obj, oldObj *Struct) (errs field.ErrorList) {
	// field Struct.TypeMeta has no validation

	// field Struct.StructField
	errs = append(errs,
		func(fldPath *field.Path, obj, oldObj *OtherStruct) (errs field.ErrorList) {
			if op.Type == operation.Update && (obj == oldObj || (obj != nil && oldObj != nil && *obj == *oldObj)) {
				return nil // no changes
			}
			errs = append(errs, validate.FixedResult(ctx, op, fldPath, obj, oldObj, false, "field Struct.StructField")...)
			errs = append(errs, Validate_OtherStruct(ctx, op, fldPath, obj, oldObj)...)
			return
		}(fldPath.Child("structField"), &obj.StructField, safe.Field(oldObj, func(oldObj *Struct) *OtherStruct { return &oldObj.StructField }))...)

	// field Struct.StructPtrField
	errs = append(errs,
		func(fldPath *field.Path, obj, oldObj *OtherStruct) (errs field.ErrorList) {
			if op.Type == operation.Update && (obj == oldObj || (obj != nil && oldObj != nil && *obj == *oldObj)) {
				return nil // no changes
			}
			if e := validate.RequiredPointer(ctx, op, fldPath, obj, oldObj); len(e) != 0 {
				errs = append(errs, e...)
				return // do not proceed
			}
			errs = append(errs, validate.FixedResult(ctx, op, fldPath, obj, oldObj, false, "field Struct.StructPtrField")...)
			errs = append(errs, Validate_OtherStruct(ctx, op, fldPath, obj, oldObj)...)
			return
		}(fldPath.Child("structPtrField"), obj.StructPtrField, safe.Field(oldObj, func(oldObj *Struct) *OtherStruct { return oldObj.StructPtrField }))...)

	// field Struct.OpaqueStructField
	errs = append(errs,
		func(fldPath *field.Path, obj, oldObj *OtherStruct) (errs field.ErrorList) {
			if op.Type == operation.Update && (obj == oldObj || (obj != nil && oldObj != nil && *obj == *oldObj)) {
				return nil // no changes
			}
			errs = append(errs, validate.FixedResult(ctx, op, fldPath, obj, oldObj, false, "field Struct.OpaqueStructField")...)
			return
		}(fldPath.Child("opaqueStructField"), &obj.OpaqueStructField, safe.Field(oldObj, func(oldObj *Struct) *OtherStruct { return &oldObj.OpaqueStructField }))...)

	// field Struct.OpaqueStructPtrField
	errs = append(errs,
		func(fldPath *field.Path, obj, oldObj *OtherStruct) (errs field.ErrorList) {
			if op.Type == operation.Update && (obj == oldObj || (obj != nil && oldObj != nil && *obj == *oldObj)) {
				return nil // no changes
			}
			if e := validate.RequiredPointer(ctx, op, fldPath, obj, oldObj); len(e) != 0 {
				errs = append(errs, e...)
				return // do not proceed
			}
			errs = append(errs, validate.FixedResult(ctx, op, fldPath, obj, oldObj, false, "field Struct.OpaqueStructPtrField")...)
			return
		}(fldPath.Child("opaqueStructPtrField"), obj.OpaqueStructPtrField, safe.Field(oldObj, func(oldObj *Struct) *OtherStruct { return oldObj.OpaqueStructPtrField }))...)

	// field Struct.SliceOfStructField
	errs = append(errs,
		func(fldPath *field.Path, obj, oldObj []OtherStruct) (errs field.ErrorList) {
			if op.Type == operation.Update && equality.Semantic.DeepEqual(obj, oldObj) {
				return nil // no changes
			}
			errs = append(errs, validate.FixedResult(ctx, op, fldPath, obj, oldObj, false, "field Struct.SliceOfStructField")...)
			errs = append(errs, validate.EachSliceVal(ctx, op, fldPath, obj, oldObj, nil, nil, func(ctx context.Context, op operation.Operation, fldPath *field.Path, obj, oldObj *OtherStruct) field.ErrorList {
				return validate.FixedResult(ctx, op, fldPath, obj, oldObj, false, "field Struct.SliceOfStructField vals")
			})...)
			errs = append(errs, validate.EachSliceVal(ctx, op, fldPath, obj, oldObj, nil, nil, Validate_OtherStruct)...)
			return
		}(fldPath.Child("sliceOfStructField"), obj.SliceOfStructField, safe.Field(oldObj, func(oldObj *Struct) []OtherStruct { return oldObj.SliceOfStructField }))...)

	// field Struct.SliceOfOpaqueStructField
	errs = append(errs,
		func(fldPath *field.Path, obj, oldObj []OtherStruct) (errs field.ErrorList) {
			if op.Type == operation.Update && equality.Semantic.DeepEqual(obj, oldObj) {
				return nil // no changes
			}
			errs = append(errs, validate.FixedResult(ctx, op, fldPath, obj, oldObj, false, "field Struct.SliceOfOpaqueStructField")...)
			errs = append(errs, validate.EachSliceVal(ctx, op, fldPath, obj, oldObj, nil, nil, func(ctx context.Context, op operation.Operation, fldPath *field.Path, obj, oldObj *OtherStruct) field.ErrorList {
				return validate.FixedResult(ctx, op, fldPath, obj, oldObj, false, "field Struct.SliceOfOpaqueStructField vals")
			})...)
			return
		}(fldPath.Child("sliceOfOpaqueStructField"), obj.SliceOfOpaqueStructField, safe.Field(oldObj, func(oldObj *Struct) []OtherStruct { return oldObj.SliceOfOpaqueStructField }))...)

	// field Struct.ListMapOfStructField
	errs = append(errs,
		func(fldPath *field.Path, obj, oldObj []OtherStruct) (errs field.ErrorList) {
			if op.Type == operation.Update && equality.Semantic.DeepEqual(obj, oldObj) {
				return nil // no changes
			}
			errs = append(errs, validate.FixedResult(ctx, op, fldPath, obj, oldObj, false, "field Struct.ListMapOfStructField")...)
			errs = append(errs, validate.EachSliceVal(ctx, op, fldPath, obj, oldObj, func(a OtherStruct, b OtherStruct) bool { return a.StringField == b.StringField }, validate.DirectEqual, func(ctx context.Context, op operation.Operation, fldPath *field.Path, obj, oldObj *OtherStruct) field.ErrorList {
				return validate.FixedResult(ctx, op, fldPath, obj, oldObj, false, "field Struct.ListMapOfStructField vals")
			})...)
			errs = append(errs, validate.EachSliceVal(ctx, op, fldPath, obj, oldObj, func(a OtherStruct, b OtherStruct) bool { return a.StringField == b.StringField }, validate.DirectEqual, Validate_OtherStruct)...)
			return
		}(fldPath.Child("listMapOfStructField"), obj.ListMapOfStructField, safe.Field(oldObj, func(oldObj *Struct) []OtherStruct { return oldObj.ListMapOfStructField }))...)

	// field Struct.ListMapOfOpaqueStructField
	errs = append(errs,
		func(fldPath *field.Path, obj, oldObj []OtherStruct) (errs field.ErrorList) {
			if op.Type == operation.Update && equality.Semantic.DeepEqual(obj, oldObj) {
				return nil // no changes
			}
			errs = append(errs, validate.FixedResult(ctx, op, fldPath, obj, oldObj, false, "field Struct.ListMapOfOpaqueStructField")...)
			errs = append(errs, validate.EachSliceVal(ctx, op, fldPath, obj, oldObj, func(a OtherStruct, b OtherStruct) bool { return a.StringField == b.StringField }, validate.DirectEqual, func(ctx context.Context, op operation.Operation, fldPath *field.Path, obj, oldObj *OtherStruct) field.ErrorList {
				return validate.FixedResult(ctx, op, fldPath, obj, oldObj, false, "field Struct.ListMapOfOpaqueStructField vals")
			})...)
			return
		}(fldPath.Child("listMapOfOpaqueStructField"), obj.ListMapOfOpaqueStructField, safe.Field(oldObj, func(oldObj *Struct) []OtherStruct { return oldObj.ListMapOfOpaqueStructField }))...)

	// field Struct.MapOfStringToStructField
	errs = append(errs,
		func(fldPath *field.Path, obj, oldObj map[OtherString]OtherStruct) (errs field.ErrorList) {
			if op.Type == operation.Update && equality.Semantic.DeepEqual(obj, oldObj) {
				return nil // no changes
			}
			errs = append(errs, validate.EachMapKey(ctx, op, fldPath, obj, oldObj, func(ctx context.Context, op operation.Operation, fldPath *field.Path, obj, oldObj *OtherString) field.ErrorList {
				return validate.FixedResult(ctx, op, fldPath, obj, oldObj, false, "field Struct.MapOfStringToStructField keys")
			})...)
			errs = append(errs, validate.FixedResult(ctx, op, fldPath, obj, oldObj, false, "field Struct.MapOfStringToStructField")...)
			errs = append(errs, validate.EachMapVal(ctx, op, fldPath, obj, oldObj, validate.DirectEqual, func(ctx context.Context, op operation.Operation, fldPath *field.Path, obj, oldObj *OtherStruct) field.ErrorList {
				return validate.FixedResult(ctx, op, fldPath, obj, oldObj, false, "field Struct.MapOfStringToStructField vals")
			})...)
			errs = append(errs, validate.EachMapKey(ctx, op, fldPath, obj, oldObj, Validate_OtherString)...)
			errs = append(errs, validate.EachMapVal(ctx, op, fldPath, obj, oldObj, validate.DirectEqual, Validate_OtherStruct)...)
			return
		}(fldPath.Child("mapOfStringToStructField"), obj.MapOfStringToStructField, safe.Field(oldObj, func(oldObj *Struct) map[OtherString]OtherStruct { return oldObj.MapOfStringToStructField }))...)

	// field Struct.MapOfStringToOpaqueStructField
	errs = append(errs,
		func(fldPath *field.Path, obj, oldObj map[OtherString]OtherStruct) (errs field.ErrorList) {
			if op.Type == operation.Update && equality.Semantic.DeepEqual(obj, oldObj) {
				return nil // no changes
			}
			errs = append(errs, validate.EachMapKey(ctx, op, fldPath, obj, oldObj, func(ctx context.Context, op operation.Operation, fldPath *field.Path, obj, oldObj *OtherString) field.ErrorList {
				return validate.FixedResult(ctx, op, fldPath, obj, oldObj, false, "field Struct.MapOfStringToOpaqueStructField keys")
			})...)
			errs = append(errs, validate.FixedResult(ctx, op, fldPath, obj, oldObj, false, "field Struct.MapOfStringToOpaqueStructField")...)
			errs = append(errs, validate.EachMapVal(ctx, op, fldPath, obj, oldObj, validate.DirectEqual, func(ctx context.Context, op operation.Operation, fldPath *field.Path, obj, oldObj *OtherStruct) field.ErrorList {
				return validate.FixedResult(ctx, op, fldPath, obj, oldObj, false, "field Struct.MapOfStringToOpaqueStructField vals")
			})...)
			return
		}(fldPath.Child("mapOfStringToOpaqueStructField"), obj.MapOfStringToOpaqueStructField, safe.Field(oldObj, func(oldObj *Struct) map[OtherString]OtherStruct { return oldObj.MapOfStringToOpaqueStructField }))...)

	return errs
}

func Validate_TypedefMapOther(ctx context.Context, op operation.Operation, fldPath *field.Path, obj, oldObj TypedefMapOther) (errs field.ErrorList) {
	// type TypedefMapOther
	if op.Type == operation.Update && equality.Semantic.DeepEqual(obj, oldObj) {
		return nil // no changes
	}
	errs = append(errs, validate.FixedResult(ctx, op, fldPath, obj, oldObj, true, "type TypedefMapOther")...)

	return errs
}

func Validate_TypedefSliceOther(ctx context.Context, op operation.Operation, fldPath *field.Path, obj, oldObj TypedefSliceOther) (errs field.ErrorList) {
	// type TypedefSliceOther
	if op.Type == operation.Update && equality.Semantic.DeepEqual(obj, oldObj) {
		return nil // no changes
	}
	errs = append(errs, validate.FixedResult(ctx, op, fldPath, obj, oldObj, true, "type TypedefSliceOther")...)

	return errs
}
