/**
 *@Author luojunying
 *@Date 2022-01-09 22:21
 */
package main

import "fmt"

const (
	StandardLogLevel = iota

	InfoLogLevel

	ErrorLogLevel
)

type BaseLogger interface {
	PrintLog(level int, message string)
	Write(message string)
}

//实现三个日志类
type StandardLogger struct {
	Level      int
	NextLogger BaseLogger
}

func NewStandardLogger() *StandardLogger {
	return &StandardLogger{
		Level:      StandardLogLevel,
		NextLogger: nil,
	}
}

//PrintLog 标准日志类输入日志，并且流向下一个对象方法
func (sl *StandardLogger) PrintLog(level int, message string) {
	if sl.Level == level {
		sl.Write(message)
	}
	if sl.NextLogger != nil {
		sl.NextLogger.PrintLog(level, message)
	}
}

func (sl *StandardLogger) Write(message string) {
	fmt.Printf("Standard logger out: %s", message)
}

//SetNextLogger 标准日志类设置下一个对象方法
func (sl *StandardLogger) SetNextLogger(logger BaseLogger) {
	sl.NextLogger = logger
}

//----------- Standard end

//InfoLogger 提示日志类
type InfoLogger struct {
	Level      int
	NextLogger BaseLogger
}

//NewInfoLogger 实例化提示日志类
func NewInfoLogger() *InfoLogger {
	return &InfoLogger{
		Level:      InfoLogLevel,
		NextLogger: nil,
	}
}

//Write 提示日志类的写方法
func (infoL *InfoLogger) Write(message string) {
	fmt.Printf("Info Logger out: %s.\n", message)
}

//PrintLog 提示日志类的输入日志方法
func (infoL *InfoLogger) PrintLog(level int, message string) {
	if infoL.Level == level {
		infoL.Write(message)
	}
	if infoL.NextLogger != nil {
		infoL.NextLogger.PrintLog(level, message)
	}
}

//SetNextLogger 提示日志类设置下一个对象
func (infoL *InfoLogger) SetNextLogger(logger BaseLogger) {
	infoL.NextLogger = logger
}

//ErrorLogger 错误日志类
type ErrorLogger struct {
	Level      int
	NextLogger BaseLogger
}

//NewErrorLogger 实例化错误日志类
func NewErrorLogger() *ErrorLogger {
	return &ErrorLogger{
		Level:      ErrorLogLevel,
		NextLogger: nil,
	}
}

//Write 错误日志类写方法
func (el *ErrorLogger) Write(message string) {
	fmt.Printf("Error logger out: %s.\n", message)
}

//PrintLog 错误日志类输入日志方法
func (el *ErrorLogger) PrintLog(level int, message string) {
	if el.Level == level {
		el.Write(message)
	}
	if el.NextLogger != nil {
		el.NextLogger.PrintLog(level, message)
	}
}

//SetNextLogger 错误日志类设置下一个对象
func (el *ErrorLogger) SetNextLogger(logger BaseLogger) {
	el.NextLogger = logger
}


//GetChainOfLoggers 获取日志对象链
func GetChainOfLoggers() BaseLogger {
	errLog := NewErrorLogger()
	infoLog := NewInfoLogger()
	standardLog := NewStandardLogger()

	errLog.SetNextLogger(infoLog)
	infoLog.SetNextLogger(standardLog)

	return errLog
}

func main() {
	log := GetChainOfLoggers()
	log.PrintLog(1, "err")
	log.Write("123123")
}

/**
1. 简介
Chain of Responsibility Pattern为请求创建一个接受者对象的链，这样可以使得请求和发送者解耦。

2. 责任链模式解决的问题
为了避免请求者和发送者耦合在一起，让多个对象都有可能接收数据，我们将这些接受者对象连城一个链，并且沿着这条链传递请求。直到有对象处理这个请求为止。

2.1 使用场景
js中的事件冒泡
Linux内核中的软件中断
2.2 优点
降低耦合度
简化接收对象，增加和减少请求的处理很容易
通过改变链内的成员或者调动他们的顺序，可以动态新增和删除责任。
2.3 缺点
不能保证请求一定被处理
调试代码时不方便，不容易观察运行时的特征，排查错误不方便。
*/
