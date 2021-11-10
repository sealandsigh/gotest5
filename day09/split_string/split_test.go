package split_string

import (
	"reflect"
	"testing"
)

// func TestSplit(t *testing.T) {
// 	ret := Split("babcbef", "b")
// 	want := []string{"", "a", "c", "ef"}
// 	if !reflect.DeepEqual(ret, want) {
// 		// 测试用例失败了
// 		t.Errorf("want:%v but got:%v\n", want, ret)
// 	}
// }

// func Test2Split(t *testing.T) {
// 	ret := Split("a:b:c", ":")
// 	want := []string{"a", "b", "c"}
// 	if !reflect.DeepEqual(ret, want) {
// 		// 测试用例失败了
// 		t.Errorf("want:%v but got:%v\n", want, ret)
// 	}
// }

// func Test3Split(t *testing.T) {
// 	ret := Split("abcef", "bc")
// 	want := []string{"a", "ef"}
// 	if !reflect.DeepEqual(ret, want) {
// 		// 测试用例失败了
// 		t.Errorf("want:%v but got:%v\n", want, ret)
// 	}
// }

// 用struct整合
// func TestSplit(t *testing.T) {
// 	type testCase struct {
// 		str  string
// 		sep  string
// 		want []string
// 	}

// 	testGroup := []testCase{
// 		{"babcbef", "b", []string{"", "a", "c", "ef"}},
// 		{"a:b:c", ":", []string{"a", "b", "c"}},
// 		{"abcef", "bc", []string{"a", "ef"}},
// 		{"沙河有沙又有河", "有", []string{"沙河", "沙又", "hehe"}},
// 	}

// 	for _, tc := range testGroup {
// 		got := Split(tc.str, tc.sep)
// 		if !reflect.DeepEqual(got, tc.want) {
// 			t.Fatalf("want:%#v got:%#v\n", tc.want, got)
// 		}
// 	}
// }

// 子测试
func TestSplit(t *testing.T) {
	type testCase struct {
		str  string
		sep  string
		want []string
	}

	testGroup := map[string]testCase{
		"case_1": {"babcbef", "b", []string{"", "a", "c", "ef"}},
		"case_2": {"a:b:c", ":", []string{"a", "b", "c"}},
		"case_3": {"abcef", "bc", []string{"a", "ef"}},
		"case_4": {"沙河有沙又有河", "有", []string{"沙河", "沙又", "hehe"}},
	}

	for name, tc := range testGroup {
		t.Run(name, func(t *testing.T) {
			got := Split(tc.str, tc.sep)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("want:%#v got:%#v\n", tc.want, got)
			}
		})
	}
}
