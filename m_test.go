package main

import "testing"

func RowsToStaffMap(rows []*SkillStaffUnion) map[int64][]*SkillStaffUnion {
	res := make(map[int64][]*SkillStaffUnion, 0)
	for _, v := range rows {
		n := v.StaffId
		if res[n] == nil {
			res[n] = make([]*SkillStaffUnion, 0)
		}
		res[n] = append(res[n], v)
	}

	return res
}

func RowsToStaffMap2(rows []*SkillStaffUnion) map[int64]*SkillStaffUnion {
	res := make(map[int64]*SkillStaffUnion, 0)
	for _, v := range rows {
		res[v.StaffId] = v
	}

	return res
}

type SkillStaffUnion struct {
	SkillName        string `json:"skill_name"`
	SkillId          int64  `json:"skill_id"`
	StaffId          int64  `json:"staff_id"`
	UserID           string `json:"user_id"`
	InnerName        string `json:"inner_name"`
	DisplayName      string `json:"display_name"`
	CurrentAcceptNum int64  `json:"current_accept_num"`
	MaxAccept        int64  `json:"max_accept"`
	OnlineStatus     int64  `json:"online_status"`
}

func ArrToMap2[T any](arr []T, f func(o T) int64) map[int64][]T {
	res := make(map[int64][]T, len(arr))
	for _, v := range arr {
		n := f(v)
		if res[n] == nil {
			res[n] = make([]T, 0)
		}
		res[n] = append(res[n], v)
	}
	return res
}

func ArrToMap[T any](arr []T, f func(o T) int64) map[int64]T {
	res := make(map[int64]T, len(arr))
	for _, v := range arr {
		res[f(v)] = v
	}
	return res
}

var arr = []*SkillStaffUnion{&SkillStaffUnion{
	SkillName:        "da",
	SkillId:          45,
	StaffId:          78,
	UserID:           "da",
	InnerName:        "da",
	DisplayName:      "da",
	CurrentAcceptNum: 77,
	MaxAccept:        89,
	OnlineStatus:     59,
}, &SkillStaffUnion{
	SkillName:        "da",
	SkillId:          45,
	StaffId:          728,
	UserID:           "da",
	InnerName:        "da",
	DisplayName:      "da",
	CurrentAcceptNum: 77,
	MaxAccept:        89,
	OnlineStatus:     59,
}, &SkillStaffUnion{
	SkillName:        "da",
	SkillId:          45,
	StaffId:          718,
	UserID:           "da",
	InnerName:        "da",
	DisplayName:      "da",
	CurrentAcceptNum: 77,
	MaxAccept:        89,
	OnlineStatus:     59,
}, &SkillStaffUnion{
	SkillName:        "da",
	SkillId:          45,
	StaffId:          78,
	UserID:           "da",
	InnerName:        "da",
	DisplayName:      "da",
	CurrentAcceptNum: 77,
	MaxAccept:        89,
	OnlineStatus:     59,
}, &SkillStaffUnion{
	SkillName:        "da",
	SkillId:          45,
	StaffId:          711118,
	UserID:           "dwrea",
	InnerName:        "da",
	DisplayName:      "da",
	CurrentAcceptNum: 77,
	MaxAccept:        89,
	OnlineStatus:     59,
}, &SkillStaffUnion{
	SkillName:        "da",
	SkillId:          4225,
	StaffId:          71133118,
	UserID:           "dwrea",
	InnerName:        "da",
	DisplayName:      "da",
	CurrentAcceptNum: 77,
	MaxAccept:        89,
	OnlineStatus:     59,
}}

func BenchmarkA2M(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RowsToStaffMap2(arr)
	}
}

func BenchmarkA2M2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ArrToMap[*SkillStaffUnion](arr, func(o *SkillStaffUnion) int64 { return o.StaffId })
	}
}

func BenchmarkA2M3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RowsToStaffMap(arr)
	}
}

func BenchmarkA2M4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ArrToMap2[*SkillStaffUnion](arr, func(o *SkillStaffUnion) int64 { return o.StaffId })
	}
}

func inter(a any) *SkillStaffUnion {
	switch v := a.(type) {
	case *SkillStaffUnion:
		return v
	default:
		return nil
	}
}

func ori(a *SkillStaffUnion) *SkillStaffUnion {
	var res *SkillStaffUnion
	res = a
	return res
}

func rs[T any](a any) T {
	var res T
	res, _ = a.(T)
	return res
}

// goos: darwin
// goarch: arm64
// pkg: tt
// cpu: Apple M1
// === RUN   BenchmarkOri
// BenchmarkOri
// BenchmarkOri-8          1000000000               0.3394 ns/op          0 B/op          0 allocs/op
// PASS
// ok      tt      0.936s
func BenchmarkOri(b *testing.B) {
	a := &SkillStaffUnion{SkillName: "da", SkillId: 45, StaffId: 78, UserID: "da", InnerName: "da", DisplayName: "da", CurrentAcceptNum: 77, MaxAccept: 89, OnlineStatus: 59}
	for i := 0; i < b.N; i++ {
		ori(a)
	}
}

// goos: darwin
// goarch: arm64
// pkg: tt
// cpu: Apple M1
// === RUN   BenchmarkInter
// BenchmarkInter
// BenchmarkInter-8        1000000000               0.3173 ns/op          0 B/op          0 allocs/op
// PASS
// ok      tt      0.806s
func BenchmarkInter(b *testing.B) {
	a := &SkillStaffUnion{SkillName: "da", SkillId: 45, StaffId: 78, UserID: "da", InnerName: "da", DisplayName: "da", CurrentAcceptNum: 77, MaxAccept: 89, OnlineStatus: 59}
	for i := 0; i < b.N; i++ {
		inter(a)
	}
}
func BenchmarkRs(b *testing.B) {
	a := &SkillStaffUnion{SkillName: "da", SkillId: 45, StaffId: 78, UserID: "da", InnerName: "da", DisplayName: "da", CurrentAcceptNum: 77, MaxAccept: 89, OnlineStatus: 59}
	for i := 0; i < b.N; i++ {
		rs[*SkillStaffUnion](a)
	}
}
