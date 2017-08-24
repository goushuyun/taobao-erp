package export

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

type Csv struct {
	mp       map[string]string //fileName --> title
	names    []string
	hasStart bool
	writer   *bufio.Writer
	needBom  bool

	separator string
	newline   string
}

const (
	OUT_ENCODING = "gbk" //输出编码
)

func NewCsv(w io.Writer) *Csv {
	var csv = &Csv{}
	csv.writer = bufio.NewWriter(w)
	csv.needBom = true
	csv.separator = `,`
	csv.newline = "\n"
	csv.SetHasStart(false)
	csv.names = make([]string, 0, 8)
	return csv
}

func (this *Csv) GetWriter() *bufio.Writer {
	return this.writer
}

func (this *Csv) writeLine(str ...string) {
	var ss = strings.Join(str, this.separator)
	for _, s := range ss {
		this.writer.WriteRune(s)
	}
	this.writer.WriteString(this.newline)
}

func (this *Csv) SetBoom(flag bool) {
	this.needBom = flag
}

func (this *Csv) SetHasStart(start bool) {
	this.hasStart = start
}

func (this *Csv) init() {
	this.mp = make(map[string]string)
	//this.writer.C
}

func (this *Csv) buildMap(entity interface{}) error {
	var v = reflect.ValueOf(entity)
	var t = reflect.TypeOf(entity)

	switch t.Kind() {
	case reflect.Slice:
		if v.Len() < 1 {
			return errors.New("[CSV:error] Slice cannot be empty")
		}
		return this.buildMap(v.Index(0).Interface())

	case reflect.Ptr:
		return this.buildMap(v.Elem().Interface())

	case reflect.Struct:

		for i := 0; i < t.NumField(); i++ {
			if !v.Field(i).CanInterface() {
				continue
			}

			if t.Field(i).Anonymous && v.Field(i).CanInterface() {
				this.buildAnonymous(v.Field(i).Interface())
			}

			var name = t.Field(i).Name
			var titleName = name
			var tag = string(t.Field(i).Tag)
			var regex = regexp.MustCompile(`csv:"([^\"]+)"`)
			var res = regex.FindSubmatch([]byte(tag))
			if len(res) > 1 {
				titleName = string(res[1])
			}

			if titleName != "-" && couldCsv(v.Field(i).Interface()) {
				this.names = append(this.names, name)
				this.mp[name] = titleName
			}

		}
		return nil

	default:
		return errors.New(fmt.Sprintf("[CSV:error] cannot suppor this struct -> %#v", t.Kind()))
	}
}

func (this *Csv) buildAnonymous(entity interface{}) {
	var v = reflect.ValueOf(entity)
	var t = reflect.TypeOf(entity)

	switch t.Kind() {

	case reflect.Slice:
		return

	case reflect.Ptr:
		this.buildAnonymous(v.Elem().Interface())
		return
	case reflect.Struct:

		for i := 0; i < t.NumField(); i++ {
			if !v.Field(i).CanInterface() {
				continue
			}

			if t.Field(i).Anonymous && v.Field(i).CanInterface() {
				this.buildAnonymous(v.Field(i).Interface())
			}

			var name = t.Field(i).Name
			var titleName = name
			var tag = string(t.Field(i).Tag)
			var regex = regexp.MustCompile(`csv:"([^\s]+)"`)
			var res = regex.FindSubmatch([]byte(tag))
			if len(res) > 1 {
				titleName = string(res[1])
			}

			if titleName != "-" && couldCsv(v.Field(i).Interface()) {
				this.names = append(this.names, name)
				this.mp[name] = titleName
			}

		}
	}
}

func (this *Csv) Parse(entity interface{}) (err error) {
	this.init()
	err = this.buildMap(entity)
	if err != nil {
		return err
	}
	this.writeTitle()
	err = this.parseCsv(entity)
	this.writer.Flush()
	return err
}

func (this *Csv) parseCsv(entity interface{}) (err error) {
	var v = reflect.ValueOf(entity)
	var t = reflect.TypeOf(entity)

	switch t.Kind() {
	case reflect.Slice:
		for i := 0; i < v.Len(); i++ {
			if err := this.parseCsv(v.Index(i).Interface()); err != nil {
				return err
			}
		}
		return nil
	case reflect.Ptr:
		return this.parseCsv(v.Elem().Interface())

	case reflect.Struct:
		var mp = this.mp
		var strs = make([]string, 0, t.NumField())
		for _, name := range this.names {
			_, ok := mp[name]
			if ok {
				strs = append(strs, bean2Str(v.FieldByName(name).Interface()))
			}
		}
		this.writeLine(strs...)
		return nil

	}
	return errors.New(fmt.Sprintf("[CSV:error] cannot suppor this struct -> %#v", t.Kind()))
}

func (this *Csv) writeTitle() {
	if this.hasStart {
		return
	}
	if this.needBom {
		this.writer.WriteString("\xEF\xBB\xBF")
	}

	var strs = make([]string, 0, len(this.names))
	for _, name := range this.names {
		title, ok := this.mp[name]
		// cd, err := iconv.Open("gbk", "utf-8") // convert gbk to utf8
		// if err != nil {
		// 	fmt.Println("iconv.Open failed!")
		// } else {
		// 	defer cd.Close()
		// 	title = cd.ConvString(title)
		// }

		if ok {
			strs = append(strs, title)
		}
	}

	this.hasStart = true
	this.writeLine(strs...)
}

func couldCsv(bean interface{}) bool {
	switch reflect.TypeOf(bean).Kind() {
	case reflect.Float32, reflect.Float64, reflect.Int, reflect.Int16, reflect.Int8, reflect.Int32, reflect.Uint, reflect.Int64, reflect.String, reflect.Bool:
		return true
	}
	return false
}

func bean2Str(bean interface{}) string {
	var v = reflect.ValueOf(bean)
	switch reflect.TypeOf(bean).Kind() {
	case reflect.Float32, reflect.Float64:
		return strconv.FormatFloat(v.Float(), 'f', -1, 64)
	case reflect.Int, reflect.Int16, reflect.Int8, reflect.Int32, reflect.Uint, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.String:
		value := v.String()
		return value
	case reflect.Bool:
		if v.Bool() {
			return "true"
		}
		return "false"
	default:
		return "UNKNOW"
	}
}
