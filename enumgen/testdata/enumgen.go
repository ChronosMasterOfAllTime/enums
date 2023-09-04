// Code generated by "enumgen.test.exe -test.paniconexit0 -test.timeout=10m0s"; DO NOT EDIT.

package testdata

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
	"sync/atomic"

	"goki.dev/enums"
)

const (
	_DaysName_0      = "DAY_SUNDAY"
	_DaysLowerName_0 = "day_sunday"
	_DaysName_1      = "DAY_MONDAY"
	_DaysLowerName_1 = "day_monday"
	_DaysName_2      = "DAY_TUESDAY"
	_DaysLowerName_2 = "day_tuesday"
	_DaysName_3      = "DAY_WEDNESDAY"
	_DaysLowerName_3 = "day_wednesday"
	_DaysName_4      = "DAY_THURSDAY"
	_DaysLowerName_4 = "day_thursday"
	_DaysName_5      = "DAY_FRIDAY"
	_DaysLowerName_5 = "day_friday"
	_DaysName_6      = "DAY_SATURDAY"
	_DaysLowerName_6 = "day_saturday"
)

var (
	_DaysIndex_0 = [...]uint8{0, 10}
	_DaysIndex_1 = [...]uint8{0, 10}
	_DaysIndex_2 = [...]uint8{0, 11}
	_DaysIndex_3 = [...]uint8{0, 13}
	_DaysIndex_4 = [...]uint8{0, 12}
	_DaysIndex_5 = [...]uint8{0, 10}
	_DaysIndex_6 = [...]uint8{0, 12}
)

// String returns the string representation
// of this Days value.
func (i Days) String() string {
	switch {
	case i == 1:
		return _DaysName_0
	case i == 3:
		return _DaysName_1
	case i == 5:
		return _DaysName_2
	case i == 7:
		return _DaysName_3
	case i == 9:
		return _DaysName_4
	case i == 11:
		return _DaysName_5
	case i == 13:
		return _DaysName_6
	default:
		return "Days(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the enumgen command to generate them again.
func _DaysNoOp() {
	var x [1]struct{}
	_ = x[Sunday-(1)]
	_ = x[Monday-(3)]
	_ = x[Tuesday-(5)]
	_ = x[Wednesday-(7)]
	_ = x[Thursday-(9)]
	_ = x[Friday-(11)]
	_ = x[Saturday-(13)]
}

var _DaysValues = []Days{Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday}

// DaysN is the highest valid value
// for type Days, plus one.
const DaysN Days = 14

var _DaysNameToValueMap = map[string]Days{
	_DaysName_0[0:10]:      Sunday,
	_DaysLowerName_0[0:10]: Sunday,
	_DaysName_1[0:10]:      Monday,
	_DaysLowerName_1[0:10]: Monday,
	_DaysName_2[0:11]:      Tuesday,
	_DaysLowerName_2[0:11]: Tuesday,
	_DaysName_3[0:13]:      Wednesday,
	_DaysLowerName_3[0:13]: Wednesday,
	_DaysName_4[0:12]:      Thursday,
	_DaysLowerName_4[0:12]: Thursday,
	_DaysName_5[0:10]:      Friday,
	_DaysLowerName_5[0:10]: Friday,
	_DaysName_6[0:12]:      Saturday,
	_DaysLowerName_6[0:12]: Saturday,
}

var _DaysNames = []string{
	_DaysName_0[0:10],
	_DaysName_1[0:10],
	_DaysName_2[0:11],
	_DaysName_3[0:13],
	_DaysName_4[0:12],
	_DaysName_5[0:10],
	_DaysName_6[0:12],
}

var _DaysDescMap = map[Days]string{
	1:  _DaysDescs[0],
	3:  _DaysDescs[1],
	5:  _DaysDescs[2],
	7:  _DaysDescs[3],
	9:  _DaysDescs[4],
	11: _DaysDescs[5],
	13: _DaysDescs[6],
}

var _DaysDescs = []string{
	`Sunday is the first day of the week`,
	`Monday is the second day of the week`,
	`Tuesday is the third day of the week`,
	`Wednesday is the fourth day of the week`,
	`Thursday is the fifth day of the week`,
	`Friday is the sixth day of the week`,
	`Saturday is the seventh day of the week`,
}

// SetString sets the Days value from its
// string representation, and returns an
// error if the string is invalid.
func (i *Days) SetString(s string) error {
	if val, ok := _DaysNameToValueMap[s]; ok {
		*i = val
		return nil
	}

	if val, ok := _DaysNameToValueMap[strings.ToLower(s)]; ok {
		*i = val
		return nil
	}
	return errors.New(s + " does not belong to Days values")
}

// Int64 returns the Days value as an int64.
func (i Days) Int64() int64 {
	return int64(i)
}

// SetInt64 sets the Days value from an int64.
func (i *Days) SetInt64(in int64) {
	*i = Days(in)
}

// Desc returns the description of the Days value.
func (i Days) Desc() string {
	if str, ok := _DaysDescMap[i]; ok {
		return str
	}
	return i.String()
}

// DaysValues returns all possible values of
// the type Days. This slice will be in the
// same order as those returned by the Values,
// Strings, and Descs methods on Days.
func DaysValues() []Days {
	return _DaysValues
}

// Values returns all possible values of
// type Days. This slice will be in the
// same order as those returned by Strings and Descs.
func (i Days) Values() []enums.Enum {
	res := make([]enums.Enum, len(_DaysValues))
	for i, d := range _DaysValues {
		res[i] = d
	}
	return res
}

// Strings returns the string representations of
// all possible values of type Days.
// This slice will be in the same order as
// those returned by Values and Descs.
func (i Days) Strings() []string {
	return _DaysNames
}

// Descs returns the descriptions of all
// possible values of type Days.
// This slice will be in the same order as
// those returned by Values and Strings.
func (i Days) Descs() []string {
	return _DaysDescs
}

// IsValid returns whether the value is a
// valid option for type Days.
func (i Days) IsValid() bool {
	for _, v := range _DaysValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalText implements the [encoding.TextMarshaler] interface.
func (i Days) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (i *Days) UnmarshalText(text []byte) error {
	return i.SetString(string(text))
}

// MarshalGQL implements the [graphql.Marshaler] interface.
func (i Days) MarshalGQL(w io.Writer) {
	w.Write([]byte(strconv.Quote(i.String())))
}

// UnmarshalGQL implements the [graphql.Unmarshaler] interface.
func (i *Days) UnmarshalGQL(value any) error {
	str, ok := value.(string)
	if !ok {
		return fmt.Errorf("Days should be a string, but got a value of type %T instead", value)
	}
	return i.SetString(str)
}

const (
	_StatesName_0      = "enabled"
	_StatesLowerName_0 = "enabled"
	_StatesName_1      = "not-enabled"
	_StatesLowerName_1 = "not-enabled"
	_StatesName_2      = "focused"
	_StatesLowerName_2 = "focused"
	_StatesName_3      = "vered"
	_StatesLowerName_3 = "vered"
	_StatesName_4      = "currently-being-pressed-by-user"
	_StatesLowerName_4 = "currently-being-pressed-by-user"
	_StatesName_5      = "actively-focused"
	_StatesLowerName_5 = "actively-focused"
	_StatesName_6      = "selected"
	_StatesLowerName_6 = "selected"
)

var (
	_StatesIndex_0 = [...]uint8{0, 7}
	_StatesIndex_1 = [...]uint8{0, 11}
	_StatesIndex_2 = [...]uint8{0, 7}
	_StatesIndex_3 = [...]uint8{0, 5}
	_StatesIndex_4 = [...]uint8{0, 31}
	_StatesIndex_5 = [...]uint8{0, 16}
	_StatesIndex_6 = [...]uint8{0, 8}
)

// BitIndexString returns the string
// representation of this States value
// if it is a bit index value
// (typically an enum constant), and
// not an actual bit flag value.
func (i States) BitIndexString() string {
	switch {
	case i == 1:
		return _StatesName_0
	case i == 3:
		return _StatesName_1
	case i == 5:
		return _StatesName_2
	case i == 7:
		return _StatesName_3
	case i == 9:
		return _StatesName_4
	case i == 11:
		return _StatesName_5
	case i == 13:
		return _StatesName_6
	default:
		return "States(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the enumgen command to generate them again.
func _StatesNoOp() {
	var x [1]struct{}
	_ = x[Enabled-(1)]
	_ = x[Disabled-(3)]
	_ = x[Focused-(5)]
	_ = x[Hovered-(7)]
	_ = x[Active-(9)]
	_ = x[ActivelyFocused-(11)]
	_ = x[Selected-(13)]
}

var _StatesValues = []States{Enabled, Disabled, Focused, Hovered, Active, ActivelyFocused, Selected}

// StatesN is the highest valid value
// for type States, plus one.
const StatesN States = 14

var _StatesNameToValueMap = map[string]States{
	_StatesName_0[0:7]:       Enabled,
	_StatesLowerName_0[0:7]:  Enabled,
	_StatesName_1[0:11]:      Disabled,
	_StatesLowerName_1[0:11]: Disabled,
	_StatesName_2[0:7]:       Focused,
	_StatesLowerName_2[0:7]:  Focused,
	_StatesName_3[0:5]:       Hovered,
	_StatesLowerName_3[0:5]:  Hovered,
	_StatesName_4[0:31]:      Active,
	_StatesLowerName_4[0:31]: Active,
	_StatesName_5[0:16]:      ActivelyFocused,
	_StatesLowerName_5[0:16]: ActivelyFocused,
	_StatesName_6[0:8]:       Selected,
	_StatesLowerName_6[0:8]:  Selected,
}

var _StatesNames = []string{
	_StatesName_0[0:7],
	_StatesName_1[0:11],
	_StatesName_2[0:7],
	_StatesName_3[0:5],
	_StatesName_4[0:31],
	_StatesName_5[0:16],
	_StatesName_6[0:8],
}

var _StatesDescMap = map[States]string{
	1:  _StatesDescs[0],
	3:  _StatesDescs[1],
	5:  _StatesDescs[2],
	7:  _StatesDescs[3],
	9:  _StatesDescs[4],
	11: _StatesDescs[5],
	13: _StatesDescs[6],
}

var _StatesDescs = []string{
	`Enabled indicates the widget is enabled`,
	`Disabled indicates the widget is disabled`,
	`Focused indicates the widget has keyboard focus`,
	`Hovered indicates the widget is being hovered over`,
	`Active indicates the widget is being interacted with`,
	`ActivelyFocused indicates the widget has active keyboard focus`,
	`Selected indicates the widget is selected`,
}

// SetString sets the States value from its
// string representation, and returns an
// error if the string is invalid.
func (i *States) SetString(s string) error {
	*i = 0
	flgs := strings.Split(s, "|")
	for _, flg := range flgs {
		if val, ok := _StatesNameToValueMap[flg]; ok {
			i.SetFlag(true, &val)
		} else if val, ok := _StatesNameToValueMap[strings.ToLower(flg)]; ok {
			i.SetFlag(true, &val)
		} else {
			return errors.New(flg + " is not a valid value for type States")
		}
	}
	return nil
}

// Int64 returns the States value as an int64.
func (i States) Int64() int64 {
	return int64(i)
}

// SetInt64 sets the States value from an int64.
func (i *States) SetInt64(in int64) {
	*i = States(in)
}

// Desc returns the description of the States value.
func (i States) Desc() string {
	if str, ok := _StatesDescMap[i]; ok {
		return str
	}
	return i.String()
}

// StatesValues returns all possible values of
// the type States. This slice will be in the
// same order as those returned by the Values,
// Strings, and Descs methods on States.
func StatesValues() []States {
	return _StatesValues
}

// Values returns all possible values of
// type States. This slice will be in the
// same order as those returned by Strings and Descs.
func (i States) Values() []enums.Enum {
	res := make([]enums.Enum, len(_StatesValues))
	for i, d := range _StatesValues {
		res[i] = d
	}
	return res
}

// Strings returns the string representations of
// all possible values of type States.
// This slice will be in the same order as
// those returned by Values and Descs.
func (i States) Strings() []string {
	return _StatesNames
}

// Descs returns the descriptions of all
// possible values of type States.
// This slice will be in the same order as
// those returned by Values and Strings.
func (i States) Descs() []string {
	return _StatesDescs
}

// IsValid returns whether the value is a
// valid option for type States.
func (i States) IsValid() bool {
	for _, v := range _StatesValues {
		if i == v {
			return true
		}
	}
	return false
}

// String returns the string representation
// of this States value.
func (i States) String() string {
	str := ""
	for _, ie := range _StatesValues {
		if i.HasFlag(ie) {
			ies := ie.BitIndexString()
			if str == "" {
				str = ies
			} else {
				str += "|" + ies
			}
		}
	}
	return str
}

// Has returns whether these
// bit flags have the given bit flag set.
func (i States) HasFlag(f enums.BitFlag) bool {
	return i&(1<<uint32(f.Int64())) != 0
}

// Set sets the value of the given
// flags in these flags to the given value.
func (i *States) SetFlag(on bool, f ...enums.BitFlag) {
	var mask int64
	for _, v := range f {
		mask |= 1 << v.Int64()
	}
	in := int64(*i)
	if on {
		in |= mask
		atomic.StoreInt64((*int64)(i), in)
	} else {
		in &^= mask
		atomic.StoreInt64((*int64)(i), in)
	}
}

// MarshalJSON implements the [json.Marshaler] interface.
func (i States) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the [json.Unmarshaler] interface.
func (i *States) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return errors.New("States should be a string, but got " + string(data) + "instead")
	}
	return i.SetString(s)
}

// Scan implements the [driver.Valuer] interface.
func (i States) Value() (driver.Value, error) {
	return i.String(), nil
}

// Value implements the [sql.Scanner] interface.
func (i *States) Scan(value any) error {
	if value == nil {
		return nil
	}

	var str string
	switch v := value.(type) {
	case []byte:
		str = string(v)
	case string:
		str = v
	case fmt.Stringer:
		str = v.String()
	default:
		return fmt.Errorf("invalid value for type States: %[1]T(%[1]v)", value)
	}

	return i.SetString(str)
}

const _LanguagesName = "GoPythonJavaScriptDartRustRubyCCPPObjectiveCJavaTypeScriptKotlinSwift"
const _LanguagesLowerName = "gopythonjavascriptdartrustrubyccppobjectivecjavatypescriptkotlinswift"

var _LanguagesMap = map[Languages]string{
	6:  _LanguagesName[0:2],
	10: _LanguagesName[2:8],
	14: _LanguagesName[8:18],
	18: _LanguagesName[18:22],
	22: _LanguagesName[22:26],
	26: _LanguagesName[26:30],
	30: _LanguagesName[30:31],
	34: _LanguagesName[31:34],
	38: _LanguagesName[34:44],
	42: _LanguagesName[44:48],
	46: _LanguagesName[48:58],
	50: _LanguagesName[58:64],
	54: _LanguagesName[64:69],
}

// BitIndexString returns the string
// representation of this Languages value
// if it is a bit index value
// (typically an enum constant), and
// not an actual bit flag value.
func (i Languages) BitIndexString() string {
	if str, ok := _LanguagesMap[i]; ok {
		return str
	}
	return "Languages(" + strconv.FormatInt(int64(i), 10) + ")"
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the enumgen command to generate them again.
func _LanguagesNoOp() {
	var x [1]struct{}
	_ = x[Go-(6)]
	_ = x[Python-(10)]
	_ = x[JavaScript-(14)]
	_ = x[Dart-(18)]
	_ = x[Rust-(22)]
	_ = x[Ruby-(26)]
	_ = x[C-(30)]
	_ = x[CPP-(34)]
	_ = x[ObjectiveC-(38)]
	_ = x[Java-(42)]
	_ = x[TypeScript-(46)]
	_ = x[Kotlin-(50)]
	_ = x[Swift-(54)]
}

var _LanguagesValues = []Languages{Go, Python, JavaScript, Dart, Rust, Ruby, C, CPP, ObjectiveC, Java, TypeScript, Kotlin, Swift}

// LanguagesN is the highest valid value
// for type Languages, plus one.
const LanguagesN Languages = 55

var _LanguagesNameToValueMap = map[string]Languages{
	_LanguagesName[0:2]:        Go,
	_LanguagesLowerName[0:2]:   Go,
	_LanguagesName[2:8]:        Python,
	_LanguagesLowerName[2:8]:   Python,
	_LanguagesName[8:18]:       JavaScript,
	_LanguagesLowerName[8:18]:  JavaScript,
	_LanguagesName[18:22]:      Dart,
	_LanguagesLowerName[18:22]: Dart,
	_LanguagesName[22:26]:      Rust,
	_LanguagesLowerName[22:26]: Rust,
	_LanguagesName[26:30]:      Ruby,
	_LanguagesLowerName[26:30]: Ruby,
	_LanguagesName[30:31]:      C,
	_LanguagesLowerName[30:31]: C,
	_LanguagesName[31:34]:      CPP,
	_LanguagesLowerName[31:34]: CPP,
	_LanguagesName[34:44]:      ObjectiveC,
	_LanguagesLowerName[34:44]: ObjectiveC,
	_LanguagesName[44:48]:      Java,
	_LanguagesLowerName[44:48]: Java,
	_LanguagesName[48:58]:      TypeScript,
	_LanguagesLowerName[48:58]: TypeScript,
	_LanguagesName[58:64]:      Kotlin,
	_LanguagesLowerName[58:64]: Kotlin,
	_LanguagesName[64:69]:      Swift,
	_LanguagesLowerName[64:69]: Swift,
}

var _LanguagesNames = []string{
	_LanguagesName[0:2],
	_LanguagesName[2:8],
	_LanguagesName[8:18],
	_LanguagesName[18:22],
	_LanguagesName[22:26],
	_LanguagesName[26:30],
	_LanguagesName[30:31],
	_LanguagesName[31:34],
	_LanguagesName[34:44],
	_LanguagesName[44:48],
	_LanguagesName[48:58],
	_LanguagesName[58:64],
	_LanguagesName[64:69],
}

var _LanguagesDescMap = map[Languages]string{
	6:  _LanguagesDescs[0],
	10: _LanguagesDescs[1],
	14: _LanguagesDescs[2],
	18: _LanguagesDescs[3],
	22: _LanguagesDescs[4],
	26: _LanguagesDescs[5],
	30: _LanguagesDescs[6],
	34: _LanguagesDescs[7],
	38: _LanguagesDescs[8],
	42: _LanguagesDescs[9],
	46: _LanguagesDescs[10],
	50: _LanguagesDescs[11],
	54: _LanguagesDescs[12],
}

var _LanguagesDescs = []string{
	`Go is the best programming language`,
	``,
	`JavaScript is the worst programming language`,
	``,
	``,
	``,
	``,
	``,
	``,
	``,
	``,
	``,
	``,
}

// SetString sets the Languages value from its
// string representation, and returns an
// error if the string is invalid.
func (i *Languages) SetString(s string) error {
	*i = 0
	flgs := strings.Split(s, "|")
	for _, flg := range flgs {
		if val, ok := _LanguagesNameToValueMap[flg]; ok {
			i.SetFlag(true, &val)
		} else if val, ok := _LanguagesNameToValueMap[strings.ToLower(flg)]; ok {
			i.SetFlag(true, &val)
		} else {
			return errors.New(flg + " is not a valid value for type Languages")
		}
	}
	return nil
}

// Int64 returns the Languages value as an int64.
func (i Languages) Int64() int64 {
	return int64(i)
}

// SetInt64 sets the Languages value from an int64.
func (i *Languages) SetInt64(in int64) {
	*i = Languages(in)
}

// Desc returns the description of the Languages value.
func (i Languages) Desc() string {
	if str, ok := _LanguagesDescMap[i]; ok {
		return str
	}
	return i.String()
}

// LanguagesValues returns all possible values of
// the type Languages. This slice will be in the
// same order as those returned by the Values,
// Strings, and Descs methods on Languages.
func LanguagesValues() []Languages {
	return _LanguagesValues
}

// Values returns all possible values of
// type Languages. This slice will be in the
// same order as those returned by Strings and Descs.
func (i Languages) Values() []enums.Enum {
	res := make([]enums.Enum, len(_LanguagesValues))
	for i, d := range _LanguagesValues {
		res[i] = d
	}
	return res
}

// Strings returns the string representations of
// all possible values of type Languages.
// This slice will be in the same order as
// those returned by Values and Descs.
func (i Languages) Strings() []string {
	return _LanguagesNames
}

// Descs returns the descriptions of all
// possible values of type Languages.
// This slice will be in the same order as
// those returned by Values and Strings.
func (i Languages) Descs() []string {
	return _LanguagesDescs
}

// IsValid returns whether the value is a
// valid option for type Languages.
func (i Languages) IsValid() bool {
	_, ok := _LanguagesMap[i]
	return ok
}

// String returns the string representation
// of this Languages value.
func (i Languages) String() string {
	str := ""
	for _, ie := range _LanguagesValues {
		if i.HasFlag(ie) {
			ies := ie.BitIndexString()
			if str == "" {
				str = ies
			} else {
				str += "|" + ies
			}
		}
	}
	return str
}

// Has returns whether these
// bit flags have the given bit flag set.
func (i Languages) HasFlag(f enums.BitFlag) bool {
	return i&(1<<uint32(f.Int64())) != 0
}

// Set sets the value of the given
// flags in these flags to the given value.
func (i *Languages) SetFlag(on bool, f ...enums.BitFlag) {
	var mask int64
	for _, v := range f {
		mask |= 1 << v.Int64()
	}
	in := int64(*i)
	if on {
		in |= mask
		atomic.StoreInt64((*int64)(i), in)
	} else {
		in &^= mask
		atomic.StoreInt64((*int64)(i), in)
	}
}

// MarshalJSON implements the [json.Marshaler] interface.
func (i Languages) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the [json.Unmarshaler] interface.
func (i *Languages) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return errors.New("Languages should be a string, but got " + string(data) + "instead")
	}
	return i.SetString(s)
}

// MarshalText implements the [encoding.TextMarshaler] interface.
func (i Languages) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (i *Languages) UnmarshalText(text []byte) error {
	return i.SetString(string(text))
}
