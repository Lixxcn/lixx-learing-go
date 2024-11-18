package greetings // 此包名为greetings，一般包含与问候相关的函数和测试

import (
	"regexp"
	"testing" // 导入testing包，它提供Go语言中的单元测试功能
)

// TestHelloName 调用greetings.Hello函数并传入一个名字，检查返回值是否有效。
// 这是对Hello函数的一个单元测试。
func TestHelloName(t *testing.T) {
	name := "Gladys"                               // 定义一个名为Gladys的变量
	want := regexp.MustCompile(`\b` + name + `\b`) // 使用正则表达式创建一个期望的模式匹配对象
	// 调用Hello函数并传入name参数，获取返回的消息和错误
	msg, err := Hello("Gladys")
	// 如果返回的消息不匹配期望的正则表达式，或者存在错误，则测试失败
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
		// 使用t.Fatalf记录失败信息，如果测试失败，输出实际的msg和err以及期望的匹配模式
	}
}

// TestHelloEmpty 测试当向greetings.Hello函数传入空字符串时是否返回错误。
func TestHelloEmpty(t *testing.T) {
	// 调用Hello函数并传入空字符串，获取返回的消息和错误
	msg, err := Hello("")
	// 如果返回的消息非空或者错误为空，说明不符合预期，测试失败
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
		// 使用t.Fatalf记录失败信息，如果测试失败，输出实际的msg和err以及期望的结果（空字符串和错误）
	}
}
