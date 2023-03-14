package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

type Constructor struct {
	HP     int
	Damage int
}

func GenusMonster() string {
	firstLine := ReadInts()
	n := firstLine[0]
	k := firstLine[1]
	secondLine := ReadInts()
	thirdLine := ReadInts()
	mapCon := make(map[int]Constructor)
	for i := 0; i < len(mapCon); i++ {
		mapCon[i] = Constructor{
			HP:     secondLine[i],
			Damage: thirdLine[i],
		}
	}
	minDamage := mapCon[0].Damage
	for i := 0; i < n; i++ {
		for k > 0 {
			mapCon[i] = Constructor{
				HP:     secondLine[i] - k,
				Damage: thirdLine[i],
			}
			if minDamage > mapCon[i].Damage {
				minDamage = mapCon[i].Damage
			}
			k -= minDamage
		}
		if k <= 0 {
			return "NO"
		}
	}
	fmt.Println("mapCon => ", mapCon)

	return "YES"
}

func main() {
	t := ReadInt()
	for i := 0; i < t; i++ {
		fmt.Println(GenusMonster())
	}
}

var in = bufio.NewReader(os.Stdin)

type float = float32
type double = float64
type long = int64
type ulong = uint64

//
func TypeOf(arg interface{}) string {
	switch reflect.TypeOf(arg).Kind() {
	case reflect.Int:
		return "int"
	case reflect.Int8:
		return "int8"
	case reflect.Int16:
		return "int16"
	case reflect.Int32:
		return "int32"
	case reflect.Int64:
		return "int64"
	case reflect.Uint:
		return "uint"
	case reflect.Uint8:
		return "uint8"
	case reflect.Uint16:
		return "uint16"
	case reflect.Uint32:
		return "uint32"
	case reflect.Uint64:
		return "uint64"
	case reflect.Float32:
		return "float32"
	case reflect.Float64:
		return "float64"
	case reflect.Complex64:
		return "complex64"
	case reflect.Complex128:
		return "complex128"
	case reflect.Bool:
		return "bool"
	case reflect.String:
		return "string"
	case reflect.Chan:
		return "chan"
	case reflect.Ptr:
		return "ptr"
	case reflect.Uintptr:
		return "uintptr"
	case reflect.UnsafePointer:
		return "unsafeptr"
	case reflect.Interface:
		return "interface"
	case reflect.Array:
		return "array"
	case reflect.Slice:
		return "slice"
	case reflect.Map:
		return "map"
	case reflect.Struct:
		return "struct"
	case reflect.Func:
		return "func"
	case reflect.Invalid:
		return "invalid"
	default:
		return "nil"
	}
}

//
/* ~~~~~~ */
//
func _scln() string {
	ln, _ := in.ReadString('\n')
	return strings.Trim(ln, " \r\n")
}
func _sc(s []string) string {
	if len(s) == 0 {
		return _scln()
	}
	return s[0]
}

//
func ReadLine() string              { return _scln() }
func ReadString() string            { return _scln() }
func ReadStrings() []string         { return strings.Split(_scln(), " ") }
func ReadBool(s ...string) bool     { t, _ := strconv.ParseBool(_sc(s)); return t }
func ReadByte(s ...string) byte     { t, _ := strconv.ParseUint(_sc(s), 10, 8); return byte(t) }
func ReadDouble(s ...string) double { t, _ := strconv.ParseFloat(_sc(s), 64); return t }
func ReadFloat(s ...string) float   { t, _ := strconv.ParseFloat(_sc(s), 32); return float32(t) }
func ReadInt(s ...string) int       { t, _ := strconv.Atoi(_sc(s)); return t }
func ReadLong(s ...string) long     { t, _ := strconv.ParseInt(_sc(s), 10, 64); return t }
func ReadULong(s ...string) ulong   { t, _ := strconv.ParseUint(_sc(s), 10, 64); return t }

//
/* ~~~~~~ */
//
// Too much overhead with universal method using interface{}, wait for the generics,
// but for now — just don't look at the redundancy of the code after that commentary
//
func ReadBools() []bool {
	strs := ReadStrings()
	arr := make([]bool, len(strs))
	for i, s := range strs {
		arr[i] = ReadBool(s)
	}
	return arr
}
func ReadBytes() []byte {
	strs := ReadStrings()
	arr := make([]byte, len(strs))
	for i, s := range strs {
		arr[i] = ReadByte(s)
	}
	return arr
}
func ReadDoubles() []double {
	strs := ReadStrings()
	arr := make([]double, len(strs))
	for i, s := range strs {
		arr[i] = ReadDouble(s)
	}
	return arr
}
func ReadFloats() []float {
	strs := ReadStrings()
	arr := make([]float, len(strs))
	for i, s := range strs {
		arr[i] = ReadFloat(s)
	}
	return arr
}
func ReadInts() []int {
	strs := ReadStrings()
	arr := make([]int, len(strs))
	for i, s := range strs {
		arr[i] = ReadInt(s)
	}
	return arr
}
func ReadLongs() []long {
	strs := ReadStrings()
	arr := make([]long, len(strs))
	for i, s := range strs {
		arr[i] = ReadLong(s)
	}
	return arr
}
func ReadULongs() []ulong {
	strs := ReadStrings()
	arr := make([]ulong, len(strs))
	for i, s := range strs {
		arr[i] = ReadULong(s)
	}
	return arr
}

//
/* ~~~~~~ */
//
func Write(arg ...interface{})                 { fmt.Print(arg...) }
func WriteLine(arg ...interface{})             { fmt.Println(arg...) }
func WriteFormat(f string, arg ...interface{}) { fmt.Printf(f, arg...) }
