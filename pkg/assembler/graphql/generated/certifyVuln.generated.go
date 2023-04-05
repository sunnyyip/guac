// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/guacsec/guac/pkg/assembler/graphql/model"
	"github.com/vektah/gqlparser/v2/ast"
)

// region    ************************** generated!.gotpl **************************

// endregion ************************** generated!.gotpl **************************

// region    ***************************** args.gotpl *****************************

// endregion ***************************** args.gotpl *****************************

// region    ************************** directives.gotpl **************************

// endregion ************************** directives.gotpl **************************

// region    **************************** field.gotpl *****************************

func (ec *executionContext) _CertifyVuln_id(ctx context.Context, field graphql.CollectedField, obj *model.CertifyVuln) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_CertifyVuln_id(ctx, field)
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
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.ID, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNID2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_CertifyVuln_id(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "CertifyVuln",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type ID does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _CertifyVuln_package(ctx context.Context, field graphql.CollectedField, obj *model.CertifyVuln) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_CertifyVuln_package(ctx, field)
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
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Package, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(*model.Package)
	fc.Result = res
	return ec.marshalNPackage2ᚖgithubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐPackage(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_CertifyVuln_package(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "CertifyVuln",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "id":
				return ec.fieldContext_Package_id(ctx, field)
			case "type":
				return ec.fieldContext_Package_type(ctx, field)
			case "namespaces":
				return ec.fieldContext_Package_namespaces(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type Package", field.Name)
		},
	}
	return fc, nil
}

func (ec *executionContext) _CertifyVuln_vulnerability(ctx context.Context, field graphql.CollectedField, obj *model.CertifyVuln) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_CertifyVuln_vulnerability(ctx, field)
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
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Vulnerability, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(model.Vulnerability)
	fc.Result = res
	return ec.marshalNVulnerability2githubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐVulnerability(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_CertifyVuln_vulnerability(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "CertifyVuln",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type Vulnerability does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _CertifyVuln_metadata(ctx context.Context, field graphql.CollectedField, obj *model.CertifyVuln) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_CertifyVuln_metadata(ctx, field)
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
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Metadata, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(*model.VulnerabilityMetaData)
	fc.Result = res
	return ec.marshalNVulnerabilityMetaData2ᚖgithubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐVulnerabilityMetaData(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_CertifyVuln_metadata(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "CertifyVuln",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "timeScanned":
				return ec.fieldContext_VulnerabilityMetaData_timeScanned(ctx, field)
			case "dbUri":
				return ec.fieldContext_VulnerabilityMetaData_dbUri(ctx, field)
			case "dbVersion":
				return ec.fieldContext_VulnerabilityMetaData_dbVersion(ctx, field)
			case "scannerUri":
				return ec.fieldContext_VulnerabilityMetaData_scannerUri(ctx, field)
			case "scannerVersion":
				return ec.fieldContext_VulnerabilityMetaData_scannerVersion(ctx, field)
			case "origin":
				return ec.fieldContext_VulnerabilityMetaData_origin(ctx, field)
			case "collector":
				return ec.fieldContext_VulnerabilityMetaData_collector(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type VulnerabilityMetaData", field.Name)
		},
	}
	return fc, nil
}

func (ec *executionContext) _NoVuln_id(ctx context.Context, field graphql.CollectedField, obj *model.NoVuln) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_NoVuln_id(ctx, field)
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
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.ID, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNID2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_NoVuln_id(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "NoVuln",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type ID does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _VulnerabilityMetaData_timeScanned(ctx context.Context, field graphql.CollectedField, obj *model.VulnerabilityMetaData) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_VulnerabilityMetaData_timeScanned(ctx, field)
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
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.TimeScanned, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(time.Time)
	fc.Result = res
	return ec.marshalNTime2timeᚐTime(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_VulnerabilityMetaData_timeScanned(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "VulnerabilityMetaData",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type Time does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _VulnerabilityMetaData_dbUri(ctx context.Context, field graphql.CollectedField, obj *model.VulnerabilityMetaData) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_VulnerabilityMetaData_dbUri(ctx, field)
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
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.DbURI, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_VulnerabilityMetaData_dbUri(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "VulnerabilityMetaData",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _VulnerabilityMetaData_dbVersion(ctx context.Context, field graphql.CollectedField, obj *model.VulnerabilityMetaData) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_VulnerabilityMetaData_dbVersion(ctx, field)
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
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.DbVersion, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_VulnerabilityMetaData_dbVersion(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "VulnerabilityMetaData",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _VulnerabilityMetaData_scannerUri(ctx context.Context, field graphql.CollectedField, obj *model.VulnerabilityMetaData) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_VulnerabilityMetaData_scannerUri(ctx, field)
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
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.ScannerURI, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_VulnerabilityMetaData_scannerUri(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "VulnerabilityMetaData",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _VulnerabilityMetaData_scannerVersion(ctx context.Context, field graphql.CollectedField, obj *model.VulnerabilityMetaData) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_VulnerabilityMetaData_scannerVersion(ctx, field)
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
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.ScannerVersion, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_VulnerabilityMetaData_scannerVersion(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "VulnerabilityMetaData",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _VulnerabilityMetaData_origin(ctx context.Context, field graphql.CollectedField, obj *model.VulnerabilityMetaData) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_VulnerabilityMetaData_origin(ctx, field)
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
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Origin, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_VulnerabilityMetaData_origin(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "VulnerabilityMetaData",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _VulnerabilityMetaData_collector(ctx context.Context, field graphql.CollectedField, obj *model.VulnerabilityMetaData) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_VulnerabilityMetaData_collector(ctx, field)
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
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Collector, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_VulnerabilityMetaData_collector(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "VulnerabilityMetaData",
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

func (ec *executionContext) unmarshalInputCertifyVulnSpec(ctx context.Context, obj interface{}) (model.CertifyVulnSpec, error) {
	var it model.CertifyVulnSpec
	asMap := map[string]interface{}{}
	for k, v := range obj.(map[string]interface{}) {
		asMap[k] = v
	}

	fieldsInOrder := [...]string{"id", "package", "vulnerability", "timeScanned", "dbUri", "dbVersion", "scannerUri", "scannerVersion", "origin", "collector"}
	for _, k := range fieldsInOrder {
		v, ok := asMap[k]
		if !ok {
			continue
		}
		switch k {
		case "id":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("id"))
			it.ID, err = ec.unmarshalOID2ᚖstring(ctx, v)
			if err != nil {
				return it, err
			}
		case "package":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("package"))
			it.Package, err = ec.unmarshalOPkgSpec2ᚖgithubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐPkgSpec(ctx, v)
			if err != nil {
				return it, err
			}
		case "vulnerability":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("vulnerability"))
			it.Vulnerability, err = ec.unmarshalOVulnerabilitySpec2ᚖgithubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐVulnerabilitySpec(ctx, v)
			if err != nil {
				return it, err
			}
		case "timeScanned":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("timeScanned"))
			it.TimeScanned, err = ec.unmarshalOTime2ᚖtimeᚐTime(ctx, v)
			if err != nil {
				return it, err
			}
		case "dbUri":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("dbUri"))
			it.DbURI, err = ec.unmarshalOString2ᚖstring(ctx, v)
			if err != nil {
				return it, err
			}
		case "dbVersion":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("dbVersion"))
			it.DbVersion, err = ec.unmarshalOString2ᚖstring(ctx, v)
			if err != nil {
				return it, err
			}
		case "scannerUri":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("scannerUri"))
			it.ScannerURI, err = ec.unmarshalOString2ᚖstring(ctx, v)
			if err != nil {
				return it, err
			}
		case "scannerVersion":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("scannerVersion"))
			it.ScannerVersion, err = ec.unmarshalOString2ᚖstring(ctx, v)
			if err != nil {
				return it, err
			}
		case "origin":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("origin"))
			it.Origin, err = ec.unmarshalOString2ᚖstring(ctx, v)
			if err != nil {
				return it, err
			}
		case "collector":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("collector"))
			it.Collector, err = ec.unmarshalOString2ᚖstring(ctx, v)
			if err != nil {
				return it, err
			}
		}
	}

	return it, nil
}

func (ec *executionContext) unmarshalInputVulnerabilityInput(ctx context.Context, obj interface{}) (model.VulnerabilityInput, error) {
	var it model.VulnerabilityInput
	asMap := map[string]interface{}{}
	for k, v := range obj.(map[string]interface{}) {
		asMap[k] = v
	}

	fieldsInOrder := [...]string{"osv", "cve", "ghsa", "noVuln"}
	for _, k := range fieldsInOrder {
		v, ok := asMap[k]
		if !ok {
			continue
		}
		switch k {
		case "osv":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("osv"))
			it.Osv, err = ec.unmarshalOOSVInputSpec2ᚖgithubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐOSVInputSpec(ctx, v)
			if err != nil {
				return it, err
			}
		case "cve":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("cve"))
			it.Cve, err = ec.unmarshalOCVEInputSpec2ᚖgithubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐCVEInputSpec(ctx, v)
			if err != nil {
				return it, err
			}
		case "ghsa":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("ghsa"))
			it.Ghsa, err = ec.unmarshalOGHSAInputSpec2ᚖgithubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐGHSAInputSpec(ctx, v)
			if err != nil {
				return it, err
			}
		case "noVuln":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("noVuln"))
			it.NoVuln, err = ec.unmarshalOBoolean2ᚖbool(ctx, v)
			if err != nil {
				return it, err
			}
		}
	}

	return it, nil
}

func (ec *executionContext) unmarshalInputVulnerabilityMetaDataInput(ctx context.Context, obj interface{}) (model.VulnerabilityMetaDataInput, error) {
	var it model.VulnerabilityMetaDataInput
	asMap := map[string]interface{}{}
	for k, v := range obj.(map[string]interface{}) {
		asMap[k] = v
	}

	fieldsInOrder := [...]string{"timeScanned", "dbUri", "dbVersion", "scannerUri", "scannerVersion", "origin", "collector"}
	for _, k := range fieldsInOrder {
		v, ok := asMap[k]
		if !ok {
			continue
		}
		switch k {
		case "timeScanned":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("timeScanned"))
			it.TimeScanned, err = ec.unmarshalNTime2timeᚐTime(ctx, v)
			if err != nil {
				return it, err
			}
		case "dbUri":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("dbUri"))
			it.DbURI, err = ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
		case "dbVersion":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("dbVersion"))
			it.DbVersion, err = ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
		case "scannerUri":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("scannerUri"))
			it.ScannerURI, err = ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
		case "scannerVersion":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("scannerVersion"))
			it.ScannerVersion, err = ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
		case "origin":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("origin"))
			it.Origin, err = ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
		case "collector":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("collector"))
			it.Collector, err = ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
		}
	}

	return it, nil
}

func (ec *executionContext) unmarshalInputVulnerabilitySpec(ctx context.Context, obj interface{}) (model.VulnerabilitySpec, error) {
	var it model.VulnerabilitySpec
	asMap := map[string]interface{}{}
	for k, v := range obj.(map[string]interface{}) {
		asMap[k] = v
	}

	fieldsInOrder := [...]string{"osv", "cve", "ghsa", "noVuln"}
	for _, k := range fieldsInOrder {
		v, ok := asMap[k]
		if !ok {
			continue
		}
		switch k {
		case "osv":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("osv"))
			it.Osv, err = ec.unmarshalOOSVSpec2ᚖgithubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐOSVSpec(ctx, v)
			if err != nil {
				return it, err
			}
		case "cve":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("cve"))
			it.Cve, err = ec.unmarshalOCVESpec2ᚖgithubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐCVESpec(ctx, v)
			if err != nil {
				return it, err
			}
		case "ghsa":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("ghsa"))
			it.Ghsa, err = ec.unmarshalOGHSASpec2ᚖgithubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐGHSASpec(ctx, v)
			if err != nil {
				return it, err
			}
		case "noVuln":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("noVuln"))
			it.NoVuln, err = ec.unmarshalOBoolean2ᚖbool(ctx, v)
			if err != nil {
				return it, err
			}
		}
	}

	return it, nil
}

// endregion **************************** input.gotpl *****************************

// region    ************************** interface.gotpl ***************************

func (ec *executionContext) _Vulnerability(ctx context.Context, sel ast.SelectionSet, obj model.Vulnerability) graphql.Marshaler {
	switch obj := (obj).(type) {
	case nil:
		return graphql.Null
	case model.Osv:
		return ec._OSV(ctx, sel, &obj)
	case *model.Osv:
		if obj == nil {
			return graphql.Null
		}
		return ec._OSV(ctx, sel, obj)
	case model.Cve:
		return ec._CVE(ctx, sel, &obj)
	case *model.Cve:
		if obj == nil {
			return graphql.Null
		}
		return ec._CVE(ctx, sel, obj)
	case model.Ghsa:
		return ec._GHSA(ctx, sel, &obj)
	case *model.Ghsa:
		if obj == nil {
			return graphql.Null
		}
		return ec._GHSA(ctx, sel, obj)
	case model.NoVuln:
		return ec._NoVuln(ctx, sel, &obj)
	case *model.NoVuln:
		if obj == nil {
			return graphql.Null
		}
		return ec._NoVuln(ctx, sel, obj)
	default:
		panic(fmt.Errorf("unexpected type %T", obj))
	}
}

// endregion ************************** interface.gotpl ***************************

// region    **************************** object.gotpl ****************************

var certifyVulnImplementors = []string{"CertifyVuln", "Node"}

func (ec *executionContext) _CertifyVuln(ctx context.Context, sel ast.SelectionSet, obj *model.CertifyVuln) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, certifyVulnImplementors)
	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("CertifyVuln")
		case "id":

			out.Values[i] = ec._CertifyVuln_id(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "package":

			out.Values[i] = ec._CertifyVuln_package(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "vulnerability":

			out.Values[i] = ec._CertifyVuln_vulnerability(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "metadata":

			out.Values[i] = ec._CertifyVuln_metadata(ctx, field, obj)

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

var noVulnImplementors = []string{"NoVuln", "Vulnerability", "Node"}

func (ec *executionContext) _NoVuln(ctx context.Context, sel ast.SelectionSet, obj *model.NoVuln) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, noVulnImplementors)
	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("NoVuln")
		case "id":

			out.Values[i] = ec._NoVuln_id(ctx, field, obj)

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

var vulnerabilityMetaDataImplementors = []string{"VulnerabilityMetaData"}

func (ec *executionContext) _VulnerabilityMetaData(ctx context.Context, sel ast.SelectionSet, obj *model.VulnerabilityMetaData) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, vulnerabilityMetaDataImplementors)
	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("VulnerabilityMetaData")
		case "timeScanned":

			out.Values[i] = ec._VulnerabilityMetaData_timeScanned(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "dbUri":

			out.Values[i] = ec._VulnerabilityMetaData_dbUri(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "dbVersion":

			out.Values[i] = ec._VulnerabilityMetaData_dbVersion(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "scannerUri":

			out.Values[i] = ec._VulnerabilityMetaData_scannerUri(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "scannerVersion":

			out.Values[i] = ec._VulnerabilityMetaData_scannerVersion(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "origin":

			out.Values[i] = ec._VulnerabilityMetaData_origin(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "collector":

			out.Values[i] = ec._VulnerabilityMetaData_collector(ctx, field, obj)

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

func (ec *executionContext) marshalNCertifyVuln2githubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐCertifyVuln(ctx context.Context, sel ast.SelectionSet, v model.CertifyVuln) graphql.Marshaler {
	return ec._CertifyVuln(ctx, sel, &v)
}

func (ec *executionContext) marshalNCertifyVuln2ᚕᚖgithubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐCertifyVulnᚄ(ctx context.Context, sel ast.SelectionSet, v []*model.CertifyVuln) graphql.Marshaler {
	ret := make(graphql.Array, len(v))
	var wg sync.WaitGroup
	isLen1 := len(v) == 1
	if !isLen1 {
		wg.Add(len(v))
	}
	for i := range v {
		i := i
		fc := &graphql.FieldContext{
			Index:  &i,
			Result: &v[i],
		}
		ctx := graphql.WithFieldContext(ctx, fc)
		f := func(i int) {
			defer func() {
				if r := recover(); r != nil {
					ec.Error(ctx, ec.Recover(ctx, r))
					ret = nil
				}
			}()
			if !isLen1 {
				defer wg.Done()
			}
			ret[i] = ec.marshalNCertifyVuln2ᚖgithubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐCertifyVuln(ctx, sel, v[i])
		}
		if isLen1 {
			f(i)
		} else {
			go f(i)
		}

	}
	wg.Wait()

	for _, e := range ret {
		if e == graphql.Null {
			return graphql.Null
		}
	}

	return ret
}

func (ec *executionContext) marshalNCertifyVuln2ᚖgithubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐCertifyVuln(ctx context.Context, sel ast.SelectionSet, v *model.CertifyVuln) graphql.Marshaler {
	if v == nil {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "the requested element is null which the schema does not allow")
		}
		return graphql.Null
	}
	return ec._CertifyVuln(ctx, sel, v)
}

func (ec *executionContext) marshalNVulnerability2githubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐVulnerability(ctx context.Context, sel ast.SelectionSet, v model.Vulnerability) graphql.Marshaler {
	if v == nil {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "the requested element is null which the schema does not allow")
		}
		return graphql.Null
	}
	return ec._Vulnerability(ctx, sel, v)
}

func (ec *executionContext) unmarshalNVulnerabilityInput2githubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐVulnerabilityInput(ctx context.Context, v interface{}) (model.VulnerabilityInput, error) {
	res, err := ec.unmarshalInputVulnerabilityInput(ctx, v)
	return res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) marshalNVulnerabilityMetaData2ᚖgithubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐVulnerabilityMetaData(ctx context.Context, sel ast.SelectionSet, v *model.VulnerabilityMetaData) graphql.Marshaler {
	if v == nil {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "the requested element is null which the schema does not allow")
		}
		return graphql.Null
	}
	return ec._VulnerabilityMetaData(ctx, sel, v)
}

func (ec *executionContext) unmarshalNVulnerabilityMetaDataInput2githubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐVulnerabilityMetaDataInput(ctx context.Context, v interface{}) (model.VulnerabilityMetaDataInput, error) {
	res, err := ec.unmarshalInputVulnerabilityMetaDataInput(ctx, v)
	return res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) unmarshalOCertifyVulnSpec2ᚖgithubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐCertifyVulnSpec(ctx context.Context, v interface{}) (*model.CertifyVulnSpec, error) {
	if v == nil {
		return nil, nil
	}
	res, err := ec.unmarshalInputCertifyVulnSpec(ctx, v)
	return &res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) unmarshalOVulnerabilitySpec2ᚖgithubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐVulnerabilitySpec(ctx context.Context, v interface{}) (*model.VulnerabilitySpec, error) {
	if v == nil {
		return nil, nil
	}
	res, err := ec.unmarshalInputVulnerabilitySpec(ctx, v)
	return &res, graphql.ErrorOnPath(ctx, err)
}

// endregion ***************************** type.gotpl *****************************
