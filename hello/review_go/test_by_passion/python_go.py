import ctypes # 导入ctypes包
from ctypes import CDLL # CDLL用于加载.so文件
# so文件加载方法,点出内部的函数
# 使用 09_go_python.go 编译
# 编译命令 go build -buildmode=c-shared -o go_python.so ./09_go_python.go
add = CDLL("./go_python.so").addstr
# 指明传入参数的类型  add.argstypes = []--->>列表
add.argtypes = [ctypes.c_char_p, ctypes.c_char_p]
# 指明返回值的数据类型
add.restype = ctypes.c_char_p
# 正常调用
print(add(b"hello", b"world"))
# 加载同一个so文件,内部含有两个函数  test addstr
test = CDLL("./go_python.so").test
# 调用函数
print(test())
