/* ----------------------------------------------------------------------------
 * This file was automatically generated by SWIG (http://www.swig.org).
 * Version 4.0.1
 *
 * This file is not intended to be easily readable and contains a number of
 * coding conventions designed to improve portability and efficiency. Do not make
 * changes to this file unless you know what you are doing--modify the SWIG
 * interface file instead.
 * ----------------------------------------------------------------------------- */

// source: SGP4.i

package sgp4

/*
#define intgo swig_intgo
typedef void *swig_voidp;

#include <stdint.h>


typedef long long intgo;
typedef unsigned long long uintgo;



typedef struct { char *p; intgo n; } _gostring_;
typedef struct { void* array; intgo len; intgo cap; } _goslice_;


typedef _gostring_ swig_type_1;
typedef _gostring_ swig_type_2;
extern void _wrap_Swig_free_sgp4_380a74ed070f3dce(uintptr_t arg1);
extern uintptr_t _wrap_Swig_malloc_sgp4_380a74ed070f3dce(swig_intgo arg1);
extern uintptr_t _wrap_new_SGP4_sgp4_380a74ed070f3dce(uintptr_t arg1);
extern uintptr_t _wrap_SGP4_FindPosition_sgp4_380a74ed070f3dce(uintptr_t arg1, uintptr_t arg2);
extern void _wrap_delete_SGP4_sgp4_380a74ed070f3dce(uintptr_t arg1);
extern uintptr_t _wrap_new_Tle_sgp4_380a74ed070f3dce(swig_type_1 arg1, swig_type_2 arg2);
extern void _wrap_delete_Tle_sgp4_380a74ed070f3dce(uintptr_t arg1);
extern uintptr_t _wrap_new_Observer_sgp4_380a74ed070f3dce(double arg1, double arg2, double arg3);
extern void _wrap_delete_Observer_sgp4_380a74ed070f3dce(uintptr_t arg1);
extern uintptr_t _wrap_new_Eci_sgp4_380a74ed070f3dce(uintptr_t arg1, double arg2, double arg3, double arg4);
extern uintptr_t _wrap_Eci_ToGeodetic_sgp4_380a74ed070f3dce(uintptr_t arg1);
extern void _wrap_delete_Eci_sgp4_380a74ed070f3dce(uintptr_t arg1);
extern uintptr_t _wrap_new_DateTime_sgp4_380a74ed070f3dce(swig_intgo arg1, swig_intgo arg2, swig_intgo arg3, swig_intgo arg4, swig_intgo arg5, swig_intgo arg6);
extern uintptr_t _wrap_DateTime_Now__SWIG_0_sgp4_380a74ed070f3dce(_Bool arg1);
extern uintptr_t _wrap_DateTime_Now__SWIG_1_sgp4_380a74ed070f3dce(void);
extern void _wrap_delete_DateTime_sgp4_380a74ed070f3dce(uintptr_t arg1);
extern void _wrap_CoordGeodetic_latitude_set_sgp4_380a74ed070f3dce(uintptr_t arg1, double arg2);
extern double _wrap_CoordGeodetic_latitude_get_sgp4_380a74ed070f3dce(uintptr_t arg1);
extern void _wrap_CoordGeodetic_longitude_set_sgp4_380a74ed070f3dce(uintptr_t arg1, double arg2);
extern double _wrap_CoordGeodetic_longitude_get_sgp4_380a74ed070f3dce(uintptr_t arg1);
extern void _wrap_CoordGeodetic_altitude_set_sgp4_380a74ed070f3dce(uintptr_t arg1, double arg2);
extern double _wrap_CoordGeodetic_altitude_get_sgp4_380a74ed070f3dce(uintptr_t arg1);
extern uintptr_t _wrap_new_CoordGeodetic_sgp4_380a74ed070f3dce(void);
extern void _wrap_delete_CoordGeodetic_sgp4_380a74ed070f3dce(uintptr_t arg1);
#undef intgo
*/
import "C"

import "unsafe"
import _ "runtime/cgo"
import "sync"


type _ unsafe.Pointer



var Swig_escape_always_false bool
var Swig_escape_val interface{}


type _swig_fnptr *byte
type _swig_memberptr *byte


type _ sync.Mutex

func Swig_free(arg1 uintptr) {
	_swig_i_0 := arg1
	C._wrap_Swig_free_sgp4_380a74ed070f3dce(C.uintptr_t(_swig_i_0))
}

func Swig_malloc(arg1 int) (_swig_ret uintptr) {
	var swig_r uintptr
	_swig_i_0 := arg1
	swig_r = (uintptr)(C._wrap_Swig_malloc_sgp4_380a74ed070f3dce(C.swig_intgo(_swig_i_0)))
	return swig_r
}

type SwigcptrSGP4 uintptr

func (p SwigcptrSGP4) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrSGP4) SwigIsSGP4() {
}

func NewSGP4(arg1 Tle) (_swig_ret SGP4) {
	var swig_r SGP4
	_swig_i_0 := arg1.Swigcptr()
	swig_r = (SGP4)(SwigcptrSGP4(C._wrap_new_SGP4_sgp4_380a74ed070f3dce(C.uintptr_t(_swig_i_0))))
	return swig_r
}

func (arg1 SwigcptrSGP4) FindPosition(arg2 DateTime) (_swig_ret Eci) {
	var swig_r Eci
	_swig_i_0 := arg1
	_swig_i_1 := arg2.Swigcptr()
	swig_r = (Eci)(SwigcptrEci(C._wrap_SGP4_FindPosition_sgp4_380a74ed070f3dce(C.uintptr_t(_swig_i_0), C.uintptr_t(_swig_i_1))))
	return swig_r
}

func DeleteSGP4(arg1 SGP4) {
	_swig_i_0 := arg1.Swigcptr()
	C._wrap_delete_SGP4_sgp4_380a74ed070f3dce(C.uintptr_t(_swig_i_0))
}

type SGP4 interface {
	Swigcptr() uintptr
	SwigIsSGP4()
	FindPosition(arg2 DateTime) (_swig_ret Eci)
}

type SwigcptrTle uintptr

func (p SwigcptrTle) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrTle) SwigIsTle() {
}

func NewTle(arg1 string, arg2 string) (_swig_ret Tle) {
	var swig_r Tle
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	swig_r = (Tle)(SwigcptrTle(C._wrap_new_Tle_sgp4_380a74ed070f3dce(*(*C.swig_type_1)(unsafe.Pointer(&_swig_i_0)), *(*C.swig_type_2)(unsafe.Pointer(&_swig_i_1)))))
	if Swig_escape_always_false {
		Swig_escape_val = arg1
	}
	if Swig_escape_always_false {
		Swig_escape_val = arg2
	}
	return swig_r
}

func DeleteTle(arg1 Tle) {
	_swig_i_0 := arg1.Swigcptr()
	C._wrap_delete_Tle_sgp4_380a74ed070f3dce(C.uintptr_t(_swig_i_0))
}

type Tle interface {
	Swigcptr() uintptr
	SwigIsTle()
}

type SwigcptrObserver uintptr

func (p SwigcptrObserver) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrObserver) SwigIsObserver() {
}

func NewObserver(arg1 float64, arg2 float64, arg3 float64) (_swig_ret Observer) {
	var swig_r Observer
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	_swig_i_2 := arg3
	swig_r = (Observer)(SwigcptrObserver(C._wrap_new_Observer_sgp4_380a74ed070f3dce(C.double(_swig_i_0), C.double(_swig_i_1), C.double(_swig_i_2))))
	return swig_r
}

func DeleteObserver(arg1 Observer) {
	_swig_i_0 := arg1.Swigcptr()
	C._wrap_delete_Observer_sgp4_380a74ed070f3dce(C.uintptr_t(_swig_i_0))
}

type Observer interface {
	Swigcptr() uintptr
	SwigIsObserver()
}

type SwigcptrEci uintptr

func (p SwigcptrEci) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrEci) SwigIsEci() {
}

func NewEci(arg1 DateTime, arg2 float64, arg3 float64, arg4 float64) (_swig_ret Eci) {
	var swig_r Eci
	_swig_i_0 := arg1.Swigcptr()
	_swig_i_1 := arg2
	_swig_i_2 := arg3
	_swig_i_3 := arg4
	swig_r = (Eci)(SwigcptrEci(C._wrap_new_Eci_sgp4_380a74ed070f3dce(C.uintptr_t(_swig_i_0), C.double(_swig_i_1), C.double(_swig_i_2), C.double(_swig_i_3))))
	return swig_r
}

func (arg1 SwigcptrEci) ToGeodetic() (_swig_ret CoordGeodetic) {
	var swig_r CoordGeodetic
	_swig_i_0 := arg1
	swig_r = (CoordGeodetic)(SwigcptrCoordGeodetic(C._wrap_Eci_ToGeodetic_sgp4_380a74ed070f3dce(C.uintptr_t(_swig_i_0))))
	return swig_r
}

func DeleteEci(arg1 Eci) {
	_swig_i_0 := arg1.Swigcptr()
	C._wrap_delete_Eci_sgp4_380a74ed070f3dce(C.uintptr_t(_swig_i_0))
}

type Eci interface {
	Swigcptr() uintptr
	SwigIsEci()
	ToGeodetic() (_swig_ret CoordGeodetic)
}

type SwigcptrDateTime uintptr

func (p SwigcptrDateTime) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrDateTime) SwigIsDateTime() {
}

func NewDateTime(arg1 int, arg2 int, arg3 int, arg4 int, arg5 int, arg6 int) (_swig_ret DateTime) {
	var swig_r DateTime
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	_swig_i_2 := arg3
	_swig_i_3 := arg4
	_swig_i_4 := arg5
	_swig_i_5 := arg6
	swig_r = (DateTime)(SwigcptrDateTime(C._wrap_new_DateTime_sgp4_380a74ed070f3dce(C.swig_intgo(_swig_i_0), C.swig_intgo(_swig_i_1), C.swig_intgo(_swig_i_2), C.swig_intgo(_swig_i_3), C.swig_intgo(_swig_i_4), C.swig_intgo(_swig_i_5))))
	return swig_r
}

func DateTimeNow__SWIG_0(arg1 bool) (_swig_ret DateTime) {
	var swig_r DateTime
	_swig_i_0 := arg1
	swig_r = (DateTime)(SwigcptrDateTime(C._wrap_DateTime_Now__SWIG_0_sgp4_380a74ed070f3dce(C._Bool(_swig_i_0))))
	return swig_r
}

func DateTimeNow__SWIG_1() (_swig_ret DateTime) {
	var swig_r DateTime
	swig_r = (DateTime)(SwigcptrDateTime(C._wrap_DateTime_Now__SWIG_1_sgp4_380a74ed070f3dce()))
	return swig_r
}

func DateTimeNow(a ...interface{}) DateTime {
	argc := len(a)
	if argc == 0 {
		return DateTimeNow__SWIG_1()
	}
	if argc == 1 {
		return DateTimeNow__SWIG_0(a[0].(bool))
	}
	panic("No match for overloaded function call")
}

func DeleteDateTime(arg1 DateTime) {
	_swig_i_0 := arg1.Swigcptr()
	C._wrap_delete_DateTime_sgp4_380a74ed070f3dce(C.uintptr_t(_swig_i_0))
}

type DateTime interface {
	Swigcptr() uintptr
	SwigIsDateTime()
}

type SwigcptrCoordGeodetic uintptr

func (p SwigcptrCoordGeodetic) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrCoordGeodetic) SwigIsCoordGeodetic() {
}

func (arg1 SwigcptrCoordGeodetic) SetLatitude(arg2 float64) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	C._wrap_CoordGeodetic_latitude_set_sgp4_380a74ed070f3dce(C.uintptr_t(_swig_i_0), C.double(_swig_i_1))
}

func (arg1 SwigcptrCoordGeodetic) GetLatitude() (_swig_ret float64) {
	var swig_r float64
	_swig_i_0 := arg1
	swig_r = (float64)(C._wrap_CoordGeodetic_latitude_get_sgp4_380a74ed070f3dce(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrCoordGeodetic) SetLongitude(arg2 float64) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	C._wrap_CoordGeodetic_longitude_set_sgp4_380a74ed070f3dce(C.uintptr_t(_swig_i_0), C.double(_swig_i_1))
}

func (arg1 SwigcptrCoordGeodetic) GetLongitude() (_swig_ret float64) {
	var swig_r float64
	_swig_i_0 := arg1
	swig_r = (float64)(C._wrap_CoordGeodetic_longitude_get_sgp4_380a74ed070f3dce(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrCoordGeodetic) SetAltitude(arg2 float64) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	C._wrap_CoordGeodetic_altitude_set_sgp4_380a74ed070f3dce(C.uintptr_t(_swig_i_0), C.double(_swig_i_1))
}

func (arg1 SwigcptrCoordGeodetic) GetAltitude() (_swig_ret float64) {
	var swig_r float64
	_swig_i_0 := arg1
	swig_r = (float64)(C._wrap_CoordGeodetic_altitude_get_sgp4_380a74ed070f3dce(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func NewCoordGeodetic() (_swig_ret CoordGeodetic) {
	var swig_r CoordGeodetic
	swig_r = (CoordGeodetic)(SwigcptrCoordGeodetic(C._wrap_new_CoordGeodetic_sgp4_380a74ed070f3dce()))
	return swig_r
}

func DeleteCoordGeodetic(arg1 CoordGeodetic) {
	_swig_i_0 := arg1.Swigcptr()
	C._wrap_delete_CoordGeodetic_sgp4_380a74ed070f3dce(C.uintptr_t(_swig_i_0))
}

type CoordGeodetic interface {
	Swigcptr() uintptr
	SwigIsCoordGeodetic()
	SetLatitude(arg2 float64)
	GetLatitude() (_swig_ret float64)
	SetLongitude(arg2 float64)
	GetLongitude() (_swig_ret float64)
	SetAltitude(arg2 float64)
	GetAltitude() (_swig_ret float64)
}


