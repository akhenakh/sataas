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

package cppsgp4

/*
#define intgo swig_intgo
typedef void *swig_voidp;

#include <stdint.h>


typedef long long intgo;
typedef unsigned long long uintgo;



typedef struct { char *p; intgo n; } _gostring_;
typedef struct { void* array; intgo len; intgo cap; } _goslice_;


typedef long long swig_type_1;
typedef long long swig_type_2;
typedef long long swig_type_3;
typedef long long swig_type_4;
typedef _gostring_ swig_type_5;
typedef _gostring_ swig_type_6;
typedef _gostring_ swig_type_7;
typedef _gostring_ swig_type_8;
typedef _gostring_ swig_type_9;
typedef _gostring_ swig_type_10;
extern void _wrap_Swig_free_cppsgp4_d13c9b05ebdf32a2(uintptr_t arg1);
extern uintptr_t _wrap_Swig_malloc_cppsgp4_d13c9b05ebdf32a2(swig_intgo arg1);
extern uintptr_t _wrap_new_PassDetailsVector__SWIG_0_cppsgp4_d13c9b05ebdf32a2(void);
extern uintptr_t _wrap_new_PassDetailsVector__SWIG_1_cppsgp4_d13c9b05ebdf32a2(swig_type_1 arg1);
extern uintptr_t _wrap_new_PassDetailsVector__SWIG_2_cppsgp4_d13c9b05ebdf32a2(uintptr_t arg1);
extern swig_type_2 _wrap_PassDetailsVector_size_cppsgp4_d13c9b05ebdf32a2(uintptr_t arg1);
extern swig_type_3 _wrap_PassDetailsVector_capacity_cppsgp4_d13c9b05ebdf32a2(uintptr_t arg1);
extern void _wrap_PassDetailsVector_reserve_cppsgp4_d13c9b05ebdf32a2(uintptr_t arg1, swig_type_4 arg2);
extern _Bool _wrap_PassDetailsVector_isEmpty_cppsgp4_d13c9b05ebdf32a2(uintptr_t arg1);
extern void _wrap_PassDetailsVector_clear_cppsgp4_d13c9b05ebdf32a2(uintptr_t arg1);
extern void _wrap_PassDetailsVector_add_cppsgp4_d13c9b05ebdf32a2(uintptr_t arg1, uintptr_t arg2);
extern uintptr_t _wrap_PassDetailsVector_get_cppsgp4_d13c9b05ebdf32a2(uintptr_t arg1, swig_intgo arg2);
extern void _wrap_PassDetailsVector_set_cppsgp4_d13c9b05ebdf32a2(uintptr_t arg1, swig_intgo arg2, uintptr_t arg3);
extern void _wrap_delete_PassDetailsVector_cppsgp4_d13c9b05ebdf32a2(uintptr_t arg1);
extern uintptr_t _wrap_new_SGP4_cppsgp4_d13c9b05ebdf32a2(uintptr_t arg1);
extern uintptr_t _wrap_SGP4_FindPosition_cppsgp4_d13c9b05ebdf32a2(uintptr_t arg1, uintptr_t arg2);
extern void _wrap_delete_SGP4_cppsgp4_d13c9b05ebdf32a2(uintptr_t arg1);
extern uintptr_t _wrap_new_Tle_cppsgp4_d13c9b05ebdf32a2(swig_type_5 arg1, swig_type_6 arg2, swig_type_7 arg3);
extern swig_type_8 _wrap_Tle_Line1_cppsgp4_d13c9b05ebdf32a2(uintptr_t arg1);
extern swig_type_9 _wrap_Tle_Line2_cppsgp4_d13c9b05ebdf32a2(uintptr_t arg1);
extern swig_type_10 _wrap_Tle_Name_cppsgp4_d13c9b05ebdf32a2(uintptr_t arg1);
extern swig_intgo _wrap_Tle_NoradNumber_cppsgp4_d13c9b05ebdf32a2(uintptr_t arg1);
extern void _wrap_delete_Tle_cppsgp4_d13c9b05ebdf32a2(uintptr_t arg1);
extern uintptr_t _wrap_new_Observer_cppsgp4_d13c9b05ebdf32a2(double arg1, double arg2, double arg3);
extern uintptr_t _wrap_Observer_GetLookAngle_cppsgp4_d13c9b05ebdf32a2(uintptr_t arg1, uintptr_t arg2);
extern void _wrap_delete_Observer_cppsgp4_d13c9b05ebdf32a2(uintptr_t arg1);
extern uintptr_t _wrap_new_Eci_cppsgp4_d13c9b05ebdf32a2(uintptr_t arg1, double arg2, double arg3, double arg4);
extern uintptr_t _wrap_Eci_ToGeodetic_cppsgp4_d13c9b05ebdf32a2(uintptr_t arg1);
extern void _wrap_delete_Eci_cppsgp4_d13c9b05ebdf32a2(uintptr_t arg1);
extern uintptr_t _wrap_new_DateTime_cppsgp4_d13c9b05ebdf32a2(swig_intgo arg1, swig_intgo arg2, swig_intgo arg3, swig_intgo arg4, swig_intgo arg5, swig_intgo arg6);
extern uintptr_t _wrap_DateTime_Now__SWIG_0_cppsgp4_d13c9b05ebdf32a2(_Bool arg1);
extern uintptr_t _wrap_DateTime_Now__SWIG_1_cppsgp4_d13c9b05ebdf32a2(void);
extern double _wrap_DateTime_ToJulian_cppsgp4_d13c9b05ebdf32a2(uintptr_t arg1);
extern void _wrap_delete_DateTime_cppsgp4_d13c9b05ebdf32a2(uintptr_t arg1);
extern uintptr_t _wrap_GeneratePassList_cppsgp4_d13c9b05ebdf32a2(double arg1, double arg2, double arg3, uintptr_t arg4, uintptr_t arg5, uintptr_t arg6, swig_intgo arg7);
extern void _wrap_CoordGeodetic_latitude_set_cppsgp4_d13c9b05ebdf32a2(uintptr_t arg1, double arg2);
extern double _wrap_CoordGeodetic_latitude_get_cppsgp4_d13c9b05ebdf32a2(uintptr_t arg1);
extern void _wrap_CoordGeodetic_longitude_set_cppsgp4_d13c9b05ebdf32a2(uintptr_t arg1, double arg2);
extern double _wrap_CoordGeodetic_longitude_get_cppsgp4_d13c9b05ebdf32a2(uintptr_t arg1);
extern void _wrap_CoordGeodetic_altitude_set_cppsgp4_d13c9b05ebdf32a2(uintptr_t arg1, double arg2);
extern double _wrap_CoordGeodetic_altitude_get_cppsgp4_d13c9b05ebdf32a2(uintptr_t arg1);
extern uintptr_t _wrap_new_CoordGeodetic_cppsgp4_d13c9b05ebdf32a2(void);
extern void _wrap_delete_CoordGeodetic_cppsgp4_d13c9b05ebdf32a2(uintptr_t arg1);
extern void _wrap_CoordTopocentric_azimuth_set_cppsgp4_d13c9b05ebdf32a2(uintptr_t arg1, double arg2);
extern double _wrap_CoordTopocentric_azimuth_get_cppsgp4_d13c9b05ebdf32a2(uintptr_t arg1);
extern void _wrap_CoordTopocentric_elevation_set_cppsgp4_d13c9b05ebdf32a2(uintptr_t arg1, double arg2);
extern double _wrap_CoordTopocentric_elevation_get_cppsgp4_d13c9b05ebdf32a2(uintptr_t arg1);
extern void _wrap_CoordTopocentric_Xrange_set_cppsgp4_d13c9b05ebdf32a2(uintptr_t arg1, double arg2);
extern double _wrap_CoordTopocentric_Xrange_get_cppsgp4_d13c9b05ebdf32a2(uintptr_t arg1);
extern void _wrap_CoordTopocentric_range_rate_set_cppsgp4_d13c9b05ebdf32a2(uintptr_t arg1, double arg2);
extern double _wrap_CoordTopocentric_range_rate_get_cppsgp4_d13c9b05ebdf32a2(uintptr_t arg1);
extern uintptr_t _wrap_new_CoordTopocentric_cppsgp4_d13c9b05ebdf32a2(void);
extern void _wrap_delete_CoordTopocentric_cppsgp4_d13c9b05ebdf32a2(uintptr_t arg1);
extern void _wrap_EventHorizonDetails_time_set_cppsgp4_d13c9b05ebdf32a2(uintptr_t arg1, uintptr_t arg2);
extern uintptr_t _wrap_EventHorizonDetails_time_get_cppsgp4_d13c9b05ebdf32a2(uintptr_t arg1);
extern void _wrap_EventHorizonDetails_azimuth_set_cppsgp4_d13c9b05ebdf32a2(uintptr_t arg1, double arg2);
extern double _wrap_EventHorizonDetails_azimuth_get_cppsgp4_d13c9b05ebdf32a2(uintptr_t arg1);
extern uintptr_t _wrap_new_EventHorizonDetails_cppsgp4_d13c9b05ebdf32a2(void);
extern void _wrap_delete_EventHorizonDetails_cppsgp4_d13c9b05ebdf32a2(uintptr_t arg1);
extern void _wrap_PassDetails_aos_set_cppsgp4_d13c9b05ebdf32a2(uintptr_t arg1, uintptr_t arg2);
extern uintptr_t _wrap_PassDetails_aos_get_cppsgp4_d13c9b05ebdf32a2(uintptr_t arg1);
extern void _wrap_PassDetails_los_set_cppsgp4_d13c9b05ebdf32a2(uintptr_t arg1, uintptr_t arg2);
extern uintptr_t _wrap_PassDetails_los_get_cppsgp4_d13c9b05ebdf32a2(uintptr_t arg1);
extern void _wrap_PassDetails_aos_azimuth_set_cppsgp4_d13c9b05ebdf32a2(uintptr_t arg1, double arg2);
extern double _wrap_PassDetails_aos_azimuth_get_cppsgp4_d13c9b05ebdf32a2(uintptr_t arg1);
extern void _wrap_PassDetails_los_azimuth_set_cppsgp4_d13c9b05ebdf32a2(uintptr_t arg1, double arg2);
extern double _wrap_PassDetails_los_azimuth_get_cppsgp4_d13c9b05ebdf32a2(uintptr_t arg1);
extern void _wrap_PassDetails_max_elevation_set_cppsgp4_d13c9b05ebdf32a2(uintptr_t arg1, double arg2);
extern double _wrap_PassDetails_max_elevation_get_cppsgp4_d13c9b05ebdf32a2(uintptr_t arg1);
extern void _wrap_PassDetails_aos_range_rate_set_cppsgp4_d13c9b05ebdf32a2(uintptr_t arg1, double arg2);
extern double _wrap_PassDetails_aos_range_rate_get_cppsgp4_d13c9b05ebdf32a2(uintptr_t arg1);
extern void _wrap_PassDetails_los_range_rate_set_cppsgp4_d13c9b05ebdf32a2(uintptr_t arg1, double arg2);
extern double _wrap_PassDetails_los_range_rate_get_cppsgp4_d13c9b05ebdf32a2(uintptr_t arg1);
extern uintptr_t _wrap_new_PassDetails_cppsgp4_d13c9b05ebdf32a2(void);
extern void _wrap_delete_PassDetails_cppsgp4_d13c9b05ebdf32a2(uintptr_t arg1);
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

type swig_gostring struct {
	p uintptr
	n int
}

func swigCopyString(s string) string {
	p := *(*swig_gostring)(unsafe.Pointer(&s))
	r := string((*[0x7fffffff]byte)(unsafe.Pointer(p.p))[:p.n])
	Swig_free(p.p)
	return r
}

func Swig_free(arg1 uintptr) {
	_swig_i_0 := arg1
	C._wrap_Swig_free_cppsgp4_d13c9b05ebdf32a2(C.uintptr_t(_swig_i_0))
}

func Swig_malloc(arg1 int) (_swig_ret uintptr) {
	var swig_r uintptr
	_swig_i_0 := arg1
	swig_r = (uintptr)(C._wrap_Swig_malloc_cppsgp4_d13c9b05ebdf32a2(C.swig_intgo(_swig_i_0)))
	return swig_r
}

type SwigcptrPassDetailsVector uintptr

func (p SwigcptrPassDetailsVector) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrPassDetailsVector) SwigIsPassDetailsVector() {
}

func NewPassDetailsVector__SWIG_0() (_swig_ret PassDetailsVector) {
	var swig_r PassDetailsVector
	swig_r = (PassDetailsVector)(SwigcptrPassDetailsVector(C._wrap_new_PassDetailsVector__SWIG_0_cppsgp4_d13c9b05ebdf32a2()))
	return swig_r
}

func NewPassDetailsVector__SWIG_1(arg1 int64) (_swig_ret PassDetailsVector) {
	var swig_r PassDetailsVector
	_swig_i_0 := arg1
	swig_r = (PassDetailsVector)(SwigcptrPassDetailsVector(C._wrap_new_PassDetailsVector__SWIG_1_cppsgp4_d13c9b05ebdf32a2(C.swig_type_1(_swig_i_0))))
	return swig_r
}

func NewPassDetailsVector__SWIG_2(arg1 PassDetailsVector) (_swig_ret PassDetailsVector) {
	var swig_r PassDetailsVector
	_swig_i_0 := arg1.Swigcptr()
	swig_r = (PassDetailsVector)(SwigcptrPassDetailsVector(C._wrap_new_PassDetailsVector__SWIG_2_cppsgp4_d13c9b05ebdf32a2(C.uintptr_t(_swig_i_0))))
	return swig_r
}

func NewPassDetailsVector(a ...interface{}) PassDetailsVector {
	argc := len(a)
	if argc == 0 {
		return NewPassDetailsVector__SWIG_0()
	}
	if argc == 1 {
		if _, ok := a[0].(int64); !ok {
			goto check_2
		}
		return NewPassDetailsVector__SWIG_1(a[0].(int64))
	}
check_2:
	if argc == 1 {
		return NewPassDetailsVector__SWIG_2(a[0].(PassDetailsVector))
	}
	panic("No match for overloaded function call")
}

func (arg1 SwigcptrPassDetailsVector) Size() (_swig_ret int64) {
	var swig_r int64
	_swig_i_0 := arg1
	swig_r = (int64)(C._wrap_PassDetailsVector_size_cppsgp4_d13c9b05ebdf32a2(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrPassDetailsVector) Capacity() (_swig_ret int64) {
	var swig_r int64
	_swig_i_0 := arg1
	swig_r = (int64)(C._wrap_PassDetailsVector_capacity_cppsgp4_d13c9b05ebdf32a2(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrPassDetailsVector) Reserve(arg2 int64) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	C._wrap_PassDetailsVector_reserve_cppsgp4_d13c9b05ebdf32a2(C.uintptr_t(_swig_i_0), C.swig_type_4(_swig_i_1))
}

func (arg1 SwigcptrPassDetailsVector) IsEmpty() (_swig_ret bool) {
	var swig_r bool
	_swig_i_0 := arg1
	swig_r = (bool)(C._wrap_PassDetailsVector_isEmpty_cppsgp4_d13c9b05ebdf32a2(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrPassDetailsVector) Clear() {
	_swig_i_0 := arg1
	C._wrap_PassDetailsVector_clear_cppsgp4_d13c9b05ebdf32a2(C.uintptr_t(_swig_i_0))
}

func (arg1 SwigcptrPassDetailsVector) Add(arg2 PassDetails) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2.Swigcptr()
	C._wrap_PassDetailsVector_add_cppsgp4_d13c9b05ebdf32a2(C.uintptr_t(_swig_i_0), C.uintptr_t(_swig_i_1))
}

func (arg1 SwigcptrPassDetailsVector) Get(arg2 int) (_swig_ret PassDetails) {
	var swig_r PassDetails
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	swig_r = (PassDetails)(SwigcptrPassDetails(C._wrap_PassDetailsVector_get_cppsgp4_d13c9b05ebdf32a2(C.uintptr_t(_swig_i_0), C.swig_intgo(_swig_i_1))))
	return swig_r
}

func (arg1 SwigcptrPassDetailsVector) Set(arg2 int, arg3 PassDetails) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	_swig_i_2 := arg3.Swigcptr()
	C._wrap_PassDetailsVector_set_cppsgp4_d13c9b05ebdf32a2(C.uintptr_t(_swig_i_0), C.swig_intgo(_swig_i_1), C.uintptr_t(_swig_i_2))
}

func DeletePassDetailsVector(arg1 PassDetailsVector) {
	_swig_i_0 := arg1.Swigcptr()
	C._wrap_delete_PassDetailsVector_cppsgp4_d13c9b05ebdf32a2(C.uintptr_t(_swig_i_0))
}

type PassDetailsVector interface {
	Swigcptr() uintptr
	SwigIsPassDetailsVector()
	Size() (_swig_ret int64)
	Capacity() (_swig_ret int64)
	Reserve(arg2 int64)
	IsEmpty() (_swig_ret bool)
	Clear()
	Add(arg2 PassDetails)
	Get(arg2 int) (_swig_ret PassDetails)
	Set(arg2 int, arg3 PassDetails)
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
	swig_r = (SGP4)(SwigcptrSGP4(C._wrap_new_SGP4_cppsgp4_d13c9b05ebdf32a2(C.uintptr_t(_swig_i_0))))
	return swig_r
}

func (arg1 SwigcptrSGP4) FindPosition(arg2 DateTime) (_swig_ret Eci) {
	var swig_r Eci
	_swig_i_0 := arg1
	_swig_i_1 := arg2.Swigcptr()
	swig_r = (Eci)(SwigcptrEci(C._wrap_SGP4_FindPosition_cppsgp4_d13c9b05ebdf32a2(C.uintptr_t(_swig_i_0), C.uintptr_t(_swig_i_1))))
	return swig_r
}

func DeleteSGP4(arg1 SGP4) {
	_swig_i_0 := arg1.Swigcptr()
	C._wrap_delete_SGP4_cppsgp4_d13c9b05ebdf32a2(C.uintptr_t(_swig_i_0))
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

func NewTle(arg1 string, arg2 string, arg3 string) (_swig_ret Tle) {
	var swig_r Tle
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	_swig_i_2 := arg3
	swig_r = (Tle)(SwigcptrTle(C._wrap_new_Tle_cppsgp4_d13c9b05ebdf32a2(*(*C.swig_type_5)(unsafe.Pointer(&_swig_i_0)), *(*C.swig_type_6)(unsafe.Pointer(&_swig_i_1)), *(*C.swig_type_7)(unsafe.Pointer(&_swig_i_2)))))
	if Swig_escape_always_false {
		Swig_escape_val = arg1
	}
	if Swig_escape_always_false {
		Swig_escape_val = arg2
	}
	if Swig_escape_always_false {
		Swig_escape_val = arg3
	}
	return swig_r
}

func (arg1 SwigcptrTle) Line1() (_swig_ret string) {
	var swig_r string
	_swig_i_0 := arg1
	swig_r_p := C._wrap_Tle_Line1_cppsgp4_d13c9b05ebdf32a2(C.uintptr_t(_swig_i_0))
	swig_r = *(*string)(unsafe.Pointer(&swig_r_p))
	var swig_r_1 string
	swig_r_1 = swigCopyString(swig_r)
	return swig_r_1
}

func (arg1 SwigcptrTle) Line2() (_swig_ret string) {
	var swig_r string
	_swig_i_0 := arg1
	swig_r_p := C._wrap_Tle_Line2_cppsgp4_d13c9b05ebdf32a2(C.uintptr_t(_swig_i_0))
	swig_r = *(*string)(unsafe.Pointer(&swig_r_p))
	var swig_r_1 string
	swig_r_1 = swigCopyString(swig_r)
	return swig_r_1
}

func (arg1 SwigcptrTle) Name() (_swig_ret string) {
	var swig_r string
	_swig_i_0 := arg1
	swig_r_p := C._wrap_Tle_Name_cppsgp4_d13c9b05ebdf32a2(C.uintptr_t(_swig_i_0))
	swig_r = *(*string)(unsafe.Pointer(&swig_r_p))
	var swig_r_1 string
	swig_r_1 = swigCopyString(swig_r)
	return swig_r_1
}

func (arg1 SwigcptrTle) NoradNumber() (_swig_ret uint) {
	var swig_r uint
	_swig_i_0 := arg1
	swig_r = (uint)(C._wrap_Tle_NoradNumber_cppsgp4_d13c9b05ebdf32a2(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func DeleteTle(arg1 Tle) {
	_swig_i_0 := arg1.Swigcptr()
	C._wrap_delete_Tle_cppsgp4_d13c9b05ebdf32a2(C.uintptr_t(_swig_i_0))
}

type Tle interface {
	Swigcptr() uintptr
	SwigIsTle()
	Line1() (_swig_ret string)
	Line2() (_swig_ret string)
	Name() (_swig_ret string)
	NoradNumber() (_swig_ret uint)
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
	swig_r = (Observer)(SwigcptrObserver(C._wrap_new_Observer_cppsgp4_d13c9b05ebdf32a2(C.double(_swig_i_0), C.double(_swig_i_1), C.double(_swig_i_2))))
	return swig_r
}

func (arg1 SwigcptrObserver) GetLookAngle(arg2 Eci) (_swig_ret CoordTopocentric) {
	var swig_r CoordTopocentric
	_swig_i_0 := arg1
	_swig_i_1 := arg2.Swigcptr()
	swig_r = (CoordTopocentric)(SwigcptrCoordTopocentric(C._wrap_Observer_GetLookAngle_cppsgp4_d13c9b05ebdf32a2(C.uintptr_t(_swig_i_0), C.uintptr_t(_swig_i_1))))
	return swig_r
}

func DeleteObserver(arg1 Observer) {
	_swig_i_0 := arg1.Swigcptr()
	C._wrap_delete_Observer_cppsgp4_d13c9b05ebdf32a2(C.uintptr_t(_swig_i_0))
}

type Observer interface {
	Swigcptr() uintptr
	SwigIsObserver()
	GetLookAngle(arg2 Eci) (_swig_ret CoordTopocentric)
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
	swig_r = (Eci)(SwigcptrEci(C._wrap_new_Eci_cppsgp4_d13c9b05ebdf32a2(C.uintptr_t(_swig_i_0), C.double(_swig_i_1), C.double(_swig_i_2), C.double(_swig_i_3))))
	return swig_r
}

func (arg1 SwigcptrEci) ToGeodetic() (_swig_ret CoordGeodetic) {
	var swig_r CoordGeodetic
	_swig_i_0 := arg1
	swig_r = (CoordGeodetic)(SwigcptrCoordGeodetic(C._wrap_Eci_ToGeodetic_cppsgp4_d13c9b05ebdf32a2(C.uintptr_t(_swig_i_0))))
	return swig_r
}

func DeleteEci(arg1 Eci) {
	_swig_i_0 := arg1.Swigcptr()
	C._wrap_delete_Eci_cppsgp4_d13c9b05ebdf32a2(C.uintptr_t(_swig_i_0))
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
	swig_r = (DateTime)(SwigcptrDateTime(C._wrap_new_DateTime_cppsgp4_d13c9b05ebdf32a2(C.swig_intgo(_swig_i_0), C.swig_intgo(_swig_i_1), C.swig_intgo(_swig_i_2), C.swig_intgo(_swig_i_3), C.swig_intgo(_swig_i_4), C.swig_intgo(_swig_i_5))))
	return swig_r
}

func DateTimeNow__SWIG_0(arg1 bool) (_swig_ret DateTime) {
	var swig_r DateTime
	_swig_i_0 := arg1
	swig_r = (DateTime)(SwigcptrDateTime(C._wrap_DateTime_Now__SWIG_0_cppsgp4_d13c9b05ebdf32a2(C._Bool(_swig_i_0))))
	return swig_r
}

func DateTimeNow__SWIG_1() (_swig_ret DateTime) {
	var swig_r DateTime
	swig_r = (DateTime)(SwigcptrDateTime(C._wrap_DateTime_Now__SWIG_1_cppsgp4_d13c9b05ebdf32a2()))
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

func (arg1 SwigcptrDateTime) ToJulian() (_swig_ret float64) {
	var swig_r float64
	_swig_i_0 := arg1
	swig_r = (float64)(C._wrap_DateTime_ToJulian_cppsgp4_d13c9b05ebdf32a2(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func DeleteDateTime(arg1 DateTime) {
	_swig_i_0 := arg1.Swigcptr()
	C._wrap_delete_DateTime_cppsgp4_d13c9b05ebdf32a2(C.uintptr_t(_swig_i_0))
}

type DateTime interface {
	Swigcptr() uintptr
	SwigIsDateTime()
	ToJulian() (_swig_ret float64)
}

func GeneratePassList(arg1 float64, arg2 float64, arg3 float64, arg4 SGP4, arg5 DateTime, arg6 DateTime, arg7 int) (_swig_ret PassDetailsVector) {
	var swig_r PassDetailsVector
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	_swig_i_2 := arg3
	_swig_i_3 := arg4.Swigcptr()
	_swig_i_4 := arg5.Swigcptr()
	_swig_i_5 := arg6.Swigcptr()
	_swig_i_6 := arg7
	swig_r = (PassDetailsVector)(SwigcptrPassDetailsVector(C._wrap_GeneratePassList_cppsgp4_d13c9b05ebdf32a2(C.double(_swig_i_0), C.double(_swig_i_1), C.double(_swig_i_2), C.uintptr_t(_swig_i_3), C.uintptr_t(_swig_i_4), C.uintptr_t(_swig_i_5), C.swig_intgo(_swig_i_6))))
	return swig_r
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
	C._wrap_CoordGeodetic_latitude_set_cppsgp4_d13c9b05ebdf32a2(C.uintptr_t(_swig_i_0), C.double(_swig_i_1))
}

func (arg1 SwigcptrCoordGeodetic) GetLatitude() (_swig_ret float64) {
	var swig_r float64
	_swig_i_0 := arg1
	swig_r = (float64)(C._wrap_CoordGeodetic_latitude_get_cppsgp4_d13c9b05ebdf32a2(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrCoordGeodetic) SetLongitude(arg2 float64) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	C._wrap_CoordGeodetic_longitude_set_cppsgp4_d13c9b05ebdf32a2(C.uintptr_t(_swig_i_0), C.double(_swig_i_1))
}

func (arg1 SwigcptrCoordGeodetic) GetLongitude() (_swig_ret float64) {
	var swig_r float64
	_swig_i_0 := arg1
	swig_r = (float64)(C._wrap_CoordGeodetic_longitude_get_cppsgp4_d13c9b05ebdf32a2(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrCoordGeodetic) SetAltitude(arg2 float64) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	C._wrap_CoordGeodetic_altitude_set_cppsgp4_d13c9b05ebdf32a2(C.uintptr_t(_swig_i_0), C.double(_swig_i_1))
}

func (arg1 SwigcptrCoordGeodetic) GetAltitude() (_swig_ret float64) {
	var swig_r float64
	_swig_i_0 := arg1
	swig_r = (float64)(C._wrap_CoordGeodetic_altitude_get_cppsgp4_d13c9b05ebdf32a2(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func NewCoordGeodetic() (_swig_ret CoordGeodetic) {
	var swig_r CoordGeodetic
	swig_r = (CoordGeodetic)(SwigcptrCoordGeodetic(C._wrap_new_CoordGeodetic_cppsgp4_d13c9b05ebdf32a2()))
	return swig_r
}

func DeleteCoordGeodetic(arg1 CoordGeodetic) {
	_swig_i_0 := arg1.Swigcptr()
	C._wrap_delete_CoordGeodetic_cppsgp4_d13c9b05ebdf32a2(C.uintptr_t(_swig_i_0))
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

type SwigcptrCoordTopocentric uintptr

func (p SwigcptrCoordTopocentric) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrCoordTopocentric) SwigIsCoordTopocentric() {
}

func (arg1 SwigcptrCoordTopocentric) SetAzimuth(arg2 float64) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	C._wrap_CoordTopocentric_azimuth_set_cppsgp4_d13c9b05ebdf32a2(C.uintptr_t(_swig_i_0), C.double(_swig_i_1))
}

func (arg1 SwigcptrCoordTopocentric) GetAzimuth() (_swig_ret float64) {
	var swig_r float64
	_swig_i_0 := arg1
	swig_r = (float64)(C._wrap_CoordTopocentric_azimuth_get_cppsgp4_d13c9b05ebdf32a2(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrCoordTopocentric) SetElevation(arg2 float64) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	C._wrap_CoordTopocentric_elevation_set_cppsgp4_d13c9b05ebdf32a2(C.uintptr_t(_swig_i_0), C.double(_swig_i_1))
}

func (arg1 SwigcptrCoordTopocentric) GetElevation() (_swig_ret float64) {
	var swig_r float64
	_swig_i_0 := arg1
	swig_r = (float64)(C._wrap_CoordTopocentric_elevation_get_cppsgp4_d13c9b05ebdf32a2(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrCoordTopocentric) SetXrange(arg2 float64) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	C._wrap_CoordTopocentric_Xrange_set_cppsgp4_d13c9b05ebdf32a2(C.uintptr_t(_swig_i_0), C.double(_swig_i_1))
}

func (arg1 SwigcptrCoordTopocentric) GetXrange() (_swig_ret float64) {
	var swig_r float64
	_swig_i_0 := arg1
	swig_r = (float64)(C._wrap_CoordTopocentric_Xrange_get_cppsgp4_d13c9b05ebdf32a2(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrCoordTopocentric) SetRange_rate(arg2 float64) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	C._wrap_CoordTopocentric_range_rate_set_cppsgp4_d13c9b05ebdf32a2(C.uintptr_t(_swig_i_0), C.double(_swig_i_1))
}

func (arg1 SwigcptrCoordTopocentric) GetRange_rate() (_swig_ret float64) {
	var swig_r float64
	_swig_i_0 := arg1
	swig_r = (float64)(C._wrap_CoordTopocentric_range_rate_get_cppsgp4_d13c9b05ebdf32a2(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func NewCoordTopocentric() (_swig_ret CoordTopocentric) {
	var swig_r CoordTopocentric
	swig_r = (CoordTopocentric)(SwigcptrCoordTopocentric(C._wrap_new_CoordTopocentric_cppsgp4_d13c9b05ebdf32a2()))
	return swig_r
}

func DeleteCoordTopocentric(arg1 CoordTopocentric) {
	_swig_i_0 := arg1.Swigcptr()
	C._wrap_delete_CoordTopocentric_cppsgp4_d13c9b05ebdf32a2(C.uintptr_t(_swig_i_0))
}

type CoordTopocentric interface {
	Swigcptr() uintptr
	SwigIsCoordTopocentric()
	SetAzimuth(arg2 float64)
	GetAzimuth() (_swig_ret float64)
	SetElevation(arg2 float64)
	GetElevation() (_swig_ret float64)
	SetXrange(arg2 float64)
	GetXrange() (_swig_ret float64)
	SetRange_rate(arg2 float64)
	GetRange_rate() (_swig_ret float64)
}

type SwigcptrEventHorizonDetails uintptr

func (p SwigcptrEventHorizonDetails) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrEventHorizonDetails) SwigIsEventHorizonDetails() {
}

func (arg1 SwigcptrEventHorizonDetails) SetTime(arg2 DateTime) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2.Swigcptr()
	C._wrap_EventHorizonDetails_time_set_cppsgp4_d13c9b05ebdf32a2(C.uintptr_t(_swig_i_0), C.uintptr_t(_swig_i_1))
}

func (arg1 SwigcptrEventHorizonDetails) GetTime() (_swig_ret DateTime) {
	var swig_r DateTime
	_swig_i_0 := arg1
	swig_r = (DateTime)(SwigcptrDateTime(C._wrap_EventHorizonDetails_time_get_cppsgp4_d13c9b05ebdf32a2(C.uintptr_t(_swig_i_0))))
	return swig_r
}

func (arg1 SwigcptrEventHorizonDetails) SetAzimuth(arg2 float64) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	C._wrap_EventHorizonDetails_azimuth_set_cppsgp4_d13c9b05ebdf32a2(C.uintptr_t(_swig_i_0), C.double(_swig_i_1))
}

func (arg1 SwigcptrEventHorizonDetails) GetAzimuth() (_swig_ret float64) {
	var swig_r float64
	_swig_i_0 := arg1
	swig_r = (float64)(C._wrap_EventHorizonDetails_azimuth_get_cppsgp4_d13c9b05ebdf32a2(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func NewEventHorizonDetails() (_swig_ret EventHorizonDetails) {
	var swig_r EventHorizonDetails
	swig_r = (EventHorizonDetails)(SwigcptrEventHorizonDetails(C._wrap_new_EventHorizonDetails_cppsgp4_d13c9b05ebdf32a2()))
	return swig_r
}

func DeleteEventHorizonDetails(arg1 EventHorizonDetails) {
	_swig_i_0 := arg1.Swigcptr()
	C._wrap_delete_EventHorizonDetails_cppsgp4_d13c9b05ebdf32a2(C.uintptr_t(_swig_i_0))
}

type EventHorizonDetails interface {
	Swigcptr() uintptr
	SwigIsEventHorizonDetails()
	SetTime(arg2 DateTime)
	GetTime() (_swig_ret DateTime)
	SetAzimuth(arg2 float64)
	GetAzimuth() (_swig_ret float64)
}

type SwigcptrPassDetails uintptr

func (p SwigcptrPassDetails) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrPassDetails) SwigIsPassDetails() {
}

func (arg1 SwigcptrPassDetails) SetAos(arg2 DateTime) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2.Swigcptr()
	C._wrap_PassDetails_aos_set_cppsgp4_d13c9b05ebdf32a2(C.uintptr_t(_swig_i_0), C.uintptr_t(_swig_i_1))
}

func (arg1 SwigcptrPassDetails) GetAos() (_swig_ret DateTime) {
	var swig_r DateTime
	_swig_i_0 := arg1
	swig_r = (DateTime)(SwigcptrDateTime(C._wrap_PassDetails_aos_get_cppsgp4_d13c9b05ebdf32a2(C.uintptr_t(_swig_i_0))))
	return swig_r
}

func (arg1 SwigcptrPassDetails) SetLos(arg2 DateTime) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2.Swigcptr()
	C._wrap_PassDetails_los_set_cppsgp4_d13c9b05ebdf32a2(C.uintptr_t(_swig_i_0), C.uintptr_t(_swig_i_1))
}

func (arg1 SwigcptrPassDetails) GetLos() (_swig_ret DateTime) {
	var swig_r DateTime
	_swig_i_0 := arg1
	swig_r = (DateTime)(SwigcptrDateTime(C._wrap_PassDetails_los_get_cppsgp4_d13c9b05ebdf32a2(C.uintptr_t(_swig_i_0))))
	return swig_r
}

func (arg1 SwigcptrPassDetails) SetAos_azimuth(arg2 float64) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	C._wrap_PassDetails_aos_azimuth_set_cppsgp4_d13c9b05ebdf32a2(C.uintptr_t(_swig_i_0), C.double(_swig_i_1))
}

func (arg1 SwigcptrPassDetails) GetAos_azimuth() (_swig_ret float64) {
	var swig_r float64
	_swig_i_0 := arg1
	swig_r = (float64)(C._wrap_PassDetails_aos_azimuth_get_cppsgp4_d13c9b05ebdf32a2(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrPassDetails) SetLos_azimuth(arg2 float64) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	C._wrap_PassDetails_los_azimuth_set_cppsgp4_d13c9b05ebdf32a2(C.uintptr_t(_swig_i_0), C.double(_swig_i_1))
}

func (arg1 SwigcptrPassDetails) GetLos_azimuth() (_swig_ret float64) {
	var swig_r float64
	_swig_i_0 := arg1
	swig_r = (float64)(C._wrap_PassDetails_los_azimuth_get_cppsgp4_d13c9b05ebdf32a2(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrPassDetails) SetMax_elevation(arg2 float64) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	C._wrap_PassDetails_max_elevation_set_cppsgp4_d13c9b05ebdf32a2(C.uintptr_t(_swig_i_0), C.double(_swig_i_1))
}

func (arg1 SwigcptrPassDetails) GetMax_elevation() (_swig_ret float64) {
	var swig_r float64
	_swig_i_0 := arg1
	swig_r = (float64)(C._wrap_PassDetails_max_elevation_get_cppsgp4_d13c9b05ebdf32a2(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrPassDetails) SetAos_range_rate(arg2 float64) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	C._wrap_PassDetails_aos_range_rate_set_cppsgp4_d13c9b05ebdf32a2(C.uintptr_t(_swig_i_0), C.double(_swig_i_1))
}

func (arg1 SwigcptrPassDetails) GetAos_range_rate() (_swig_ret float64) {
	var swig_r float64
	_swig_i_0 := arg1
	swig_r = (float64)(C._wrap_PassDetails_aos_range_rate_get_cppsgp4_d13c9b05ebdf32a2(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrPassDetails) SetLos_range_rate(arg2 float64) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	C._wrap_PassDetails_los_range_rate_set_cppsgp4_d13c9b05ebdf32a2(C.uintptr_t(_swig_i_0), C.double(_swig_i_1))
}

func (arg1 SwigcptrPassDetails) GetLos_range_rate() (_swig_ret float64) {
	var swig_r float64
	_swig_i_0 := arg1
	swig_r = (float64)(C._wrap_PassDetails_los_range_rate_get_cppsgp4_d13c9b05ebdf32a2(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func NewPassDetails() (_swig_ret PassDetails) {
	var swig_r PassDetails
	swig_r = (PassDetails)(SwigcptrPassDetails(C._wrap_new_PassDetails_cppsgp4_d13c9b05ebdf32a2()))
	return swig_r
}

func DeletePassDetails(arg1 PassDetails) {
	_swig_i_0 := arg1.Swigcptr()
	C._wrap_delete_PassDetails_cppsgp4_d13c9b05ebdf32a2(C.uintptr_t(_swig_i_0))
}

type PassDetails interface {
	Swigcptr() uintptr
	SwigIsPassDetails()
	SetAos(arg2 DateTime)
	GetAos() (_swig_ret DateTime)
	SetLos(arg2 DateTime)
	GetLos() (_swig_ret DateTime)
	SetAos_azimuth(arg2 float64)
	GetAos_azimuth() (_swig_ret float64)
	SetLos_azimuth(arg2 float64)
	GetLos_azimuth() (_swig_ret float64)
	SetMax_elevation(arg2 float64)
	GetMax_elevation() (_swig_ret float64)
	SetAos_range_rate(arg2 float64)
	GetAos_range_rate() (_swig_ret float64)
	SetLos_range_rate(arg2 float64)
	GetLos_range_rate() (_swig_ret float64)
}
