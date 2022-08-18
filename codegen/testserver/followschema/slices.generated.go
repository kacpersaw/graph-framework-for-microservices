// Code generated by gitlab.eng.vmware.com/nsx-allspark_users/nexus-sdk/gqlgen.git, DO NOT EDIT.

package followschema

import (
	"context"
	"errors"
	"strconv"

	"gitlab.eng.vmware.com/nsx-allspark_users/nexus-sdk/gqlgen.git/graphql"
	"github.com/vektah/gqlparser/v2/ast"
)

// region    ************************** generated!.gotpl **************************

// endregion ************************** generated!.gotpl **************************

// region    ***************************** args.gotpl *****************************

// endregion ***************************** args.gotpl *****************************

// region    ************************** directives.gotpl **************************

// endregion ************************** directives.gotpl **************************

// region    **************************** field.gotpl *****************************

func (ec *executionContext) _Slices_test1(ctx context.Context, field graphql.CollectedField, obj *Slices) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Slices_test1(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Test1, nil
	})

	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.([]*string)
	fc.Result = res
	return ec.marshalOString2ᚕᚖstring(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Slices_test1(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Slices",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _Slices_test2(ctx context.Context, field graphql.CollectedField, obj *Slices) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Slices_test2(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Test2, nil
	})

	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.([]string)
	fc.Result = res
	return ec.marshalOString2ᚕstringᚄ(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Slices_test2(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Slices",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _Slices_test3(ctx context.Context, field graphql.CollectedField, obj *Slices) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Slices_test3(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Test3, nil
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.([]*string)
	fc.Result = res
	return ec.marshalNString2ᚕᚖstring(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Slices_test3(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Slices",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _Slices_test4(ctx context.Context, field graphql.CollectedField, obj *Slices) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Slices_test4(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Test4, nil
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.([]string)
	fc.Result = res
	return ec.marshalNString2ᚕstringᚄ(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Slices_test4(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Slices",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

// endregion **************************** field.gotpl *****************************

// region    **************************** input.gotpl *****************************

// endregion **************************** input.gotpl *****************************

// region    ************************** interface.gotpl ***************************

// endregion ************************** interface.gotpl ***************************

// region    **************************** object.gotpl ****************************

var slicesImplementors = []string{"Slices"}

func (ec *executionContext) _Slices(ctx context.Context, sel ast.SelectionSet, obj *Slices) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, slicesImplementors)
	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("Slices")
		case "test1":

			out.Values[i] = ec._Slices_test1(ctx, field, obj)

		case "test2":

			out.Values[i] = ec._Slices_test2(ctx, field, obj)

		case "test3":

			out.Values[i] = ec._Slices_test3(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "test4":

			out.Values[i] = ec._Slices_test4(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalids > 0 {
		return graphql.Null
	}
	return out
}

// endregion **************************** object.gotpl ****************************

// region    ***************************** type.gotpl *****************************

func (ec *executionContext) unmarshalNBytes2ᚕbyte(ctx context.Context, v interface{}) ([]byte, error) {
	res, err := UnmarshalBytes(v)
	return res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) marshalNBytes2ᚕbyte(ctx context.Context, sel ast.SelectionSet, v []byte) graphql.Marshaler {
	if v == nil {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "the requested element is null which the schema does not allow")
		}
		return graphql.Null
	}
	res := MarshalBytes(v)
	if res == graphql.Null {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "the requested element is null which the schema does not allow")
		}
	}
	return res
}

func (ec *executionContext) marshalOSlices2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐSlices(ctx context.Context, sel ast.SelectionSet, v *Slices) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return ec._Slices(ctx, sel, v)
}

// endregion ***************************** type.gotpl *****************************
