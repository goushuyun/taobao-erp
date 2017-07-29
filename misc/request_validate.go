/**
 * Copyright 2015-2016, Wothing Co., Ltd.
 * All rights reserved.
 *
 * Created by Elvizlai on 2016/04/17 11:10
 */

// validate regex stolen from github.com/asaskevich/govalidator

package misc

import (
	"bytes"
	"net/http"
	"reflect"
	"regexp"
	"strings"

	"github.com/golang/protobuf/proto"
	"github.com/goushuyun/weixin-golang/errs"
	"github.com/goushuyun/weixin-golang/misc/jsonpb"
)

var jsu = jsonpb.Unmarshaler{AllowUnknownFields: true}

var funcMap = make(map[string]interface{})

var (
	_uuid         = regexp.MustCompile("^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$")
	_alpha        = regexp.MustCompile("^[a-zA-Z]+$")
	_numeric      = regexp.MustCompile("^[-+]?[0-9]+$")
	_float        = regexp.MustCompile("^(?:[-+]?(?:[0-9]+))?(?:\\.[0-9]*)?(?:[eE][\\+\\-]?(?:[0-9]+))?$")
	_alphaNumeric = regexp.MustCompile("^[a-zA-Z0-9]+$")
	_base64       = regexp.MustCompile("^(?:[A-Za-z0-9+\\/]{4})*(?:[A-Za-z0-9+\\/]{2}==|[A-Za-z0-9+\\/]{3}=|[A-Za-z0-9+\\/]{4})$")
	_email        = regexp.MustCompile("^(((([a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+(\\.([a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+)*)|((\\x22)((((\\x20|\\x09)*(\\x0d\\x0a))?(\\x20|\\x09)+)?(([\\x01-\\x08\\x0b\\x0c\\x0e-\\x1f\\x7f]|\\x21|[\\x23-\\x5b]|[\\x5d-\\x7e]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(\\([\\x01-\\x09\\x0b\\x0c\\x0d-\\x7f]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}]))))*(((\\x20|\\x09)*(\\x0d\\x0a))?(\\x20|\\x09)+)?(\\x22)))@((([a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(([a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])([a-zA-Z]|\\d|-|\\.|_|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*([a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.)+(([a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(([a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])([a-zA-Z]|\\d|-|\\.|_|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*([a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.?$")
)

func init() {
	funcMap["null"] = isNull
	funcMap["uuid"] = isUUID
	funcMap["alpha"] = isAlpha
	funcMap["numeric"] = isNumeric
	funcMap["float"] = isFloat
	funcMap["alphanum"] = isAlphanumeric
	funcMap["base64"] = isBase64
	funcMap["email"] = isEmail
}

func isNull(str string) bool {
	return len(str) == 0
}

func isUUID(str string) bool {
	return _uuid.MatchString(str)
}

func isAlpha(str string) bool {
	if isNull(str) {
		return true
	}
	return _alpha.MatchString(str)
}

func isAlphanumeric(str string) bool {
	if isNull(str) {
		return true
	}
	return _alphaNumeric.MatchString(str)
}

func isBase64(str string) bool {
	return _base64.MatchString(str)
}

// TODO uppercase letters are not supported
func isEmail(str string) bool {
	return _email.MatchString(str)
}

func isFloat(str string) bool {
	return str != "" && _float.MatchString(str)
}

func isNumeric(str string) bool {
	if isNull(str) {
		return true
	}
	return _numeric.MatchString(str)
}

// TODO
func inRange(value interface{}, left, right float64) bool {
	return false
}

// TODO length and matches Ex:"length(min|max)": ByteLength, "matches(pattern)": StringMatches,

// abc_xyz to AbcXyz
func underscoreToCamelCase(s string) string {
	return strings.Replace(strings.Title(strings.Replace(strings.ToLower(s), "_", " ", -1)), " ", "", -1)
}

func Request2Struct(r *http.Request, req proto.Message, constraints ...string) error {
	body := r.Context().Value("body").([]byte)
	if body == nil || len(body) == 0 || bytes.Equal(body, []byte("null")) {
		return errs.NewError(errs.ErrRequestFormat, "request check failed: body is blank or 'null'")
	}
	err := jsu.Unmarshal(bytes.NewBuffer(body), req)
	if err != nil {
		return errs.NewError(errs.ErrRequestFormat, "request unmarshal failed: "+err.Error())
	}

	if len(constraints) == 0 {
		return nil
	}

	rv := reflect.Indirect(reflect.ValueOf(req))

	for _, v := range constraints {
		if err = validate(rv, v); err != nil {
			return err
		}
	}

	return nil
}

// usage:
// req := &pb.Version{}
// err:=misc.Bytes2Struct(body, req, "id:uuid", "build_version:alphanum")
// judge err
// Deprecated, using Request2Struct instead, only use in some unnormal case
func Bytes2Struct(body []byte, req proto.Message, constraints ...string) error {
	if body == nil || len(body) == 0 || bytes.Equal(body, []byte("null")) {
		return errs.NewError(errs.ErrRequestFormat, "request check failed: body is blank or 'null'")
	}

	err := jsu.Unmarshal(bytes.NewBuffer(body), req)
	if err != nil {
		return errs.NewError(errs.ErrRequestFormat, "request unmarshal failed: "+err.Error())
	}

	if len(constraints) == 0 {
		return nil
	}

	r := reflect.Indirect(reflect.ValueOf(req))

	for _, v := range constraints {
		if err = validate(r, v); err != nil {
			return err
		}
	}

	return nil
}

func validate(rv reflect.Value, constraint string) error {
	//fmt.Println(rv.Kind())
	switch rv.Kind() {
	case reflect.Invalid:
		return errs.NewError(errs.ErrRequestFormat, `validate failed for parent of '`+constraint+`' is invalid`)
	case reflect.Ptr:
		return validate(rv.Elem(), constraint)
	}

	// multi walker Ex: SA.SB
	if i := strings.Index(constraint, "."); i != -1 {
		return validate(rv.FieldByName(underscoreToCamelCase(constraint[:i])), constraint[i+1:])
	}

	x := strings.SplitN(constraint, ":", 2)

	fieldName := underscoreToCamelCase(x[0])

	value := rv.FieldByName(fieldName)
	if value.IsValid() {
		if len(x) == 2 {
			if !constraintValidate(value, funcMap[x[1]]) {
				return errs.NewError(errs.ErrRequestFormat, `require '`+x[0]+`' with constraint '`+x[1]+`'`)
			}
		} else {
			if !normalValidate(value) {
				return errs.NewError(errs.ErrRequestFormat, `require '`+x[0]+`' but missed`)
			}
		}
	} else {
		return errs.NewError(errs.ErrInternal, "no field found for `"+x[0]+"`")
	}

	return nil
}

func constraintValidate(v reflect.Value, constraintFunc interface{}) bool {
	fv := reflect.ValueOf(constraintFunc)
	if fv.Kind() == reflect.Func {
		return fv.Call([]reflect.Value{v})[0].Bool()
	}
	return false
}

func normalValidate(v reflect.Value) bool {
	k := v.Kind()
	switch k {
	case reflect.String:
		if v.String() == "" {
			return false
		}
	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int, reflect.Int64:
		if v.Int() == 0 {
			return false
		}
	case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint, reflect.Uint64:
		if v.Uint() == 0 {
			return false
		}
	case reflect.Float32, reflect.Float64:
		if v.Float() == 0 {
			return false
		}
	case reflect.Slice:
		if v.Len() == 0 {
			return false
		}
	case reflect.Ptr:
		if v.IsNil() {
			return false
		}
	}

	return true
}
