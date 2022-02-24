package test

/*
测试程序必须属于被测试的包，并且文件名满足这种形式 *_test.go
_test 程序不会被普通的 Go 编译器编译,只有 go test 会编译所有的程序：普通程序和测试程序。
测试文件中必须导入 "testing" 包，并写一些名字以 TestZzz 打头的全局函数:func TestAbcde(t *testing.T)
T 是传给测试函数的结构类型，用来管理测试状态，支持格式化测试日志，如 t.Log，t.Error，t.ErrorF 等。
用下面这些函数来通知测试失败：
1）func (t *T) Fail() 标记测试函数为失败，然后继续执行（剩下的测试）。
2）func (t *T) FailNow()  标记测试函数为失败并中止执行；文件中别的测试也被略过，继续执行下一个文件。
3）func (t *T) Log(args ...interface{}) args 被用默认的格式格式化并打印到错误日志中。
4）func (t *T) Fatal(args ...interface{}) 结合 先执行 3），然后执行 2）的效果。

testing 包中有一些类型和函数可以用来做简单的基准测试；测试代码中必须包含以 BenchmarkZzz 打头的函数并接收一个 *testing.B 类型的参数
func BenchmarkReverse(b *testing.B) {}

测试用例至少应该包括：
正常的用例
反面的用例（错误的输入，如用负数或字母代替数字，没有输入等）
边界检查用例（如果参数的取值范围是 0 到 1000，检查 0 和 1000 的情况）

*/
