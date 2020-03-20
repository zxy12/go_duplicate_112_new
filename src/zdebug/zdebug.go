package zdebug

import (
	"fmt"
	"path/filepath"
	"runtime"
	"time"
)

// PrintTrace 是否开启trace
var PrintTrace = true

// T 调试输出，可以把printTrace改为false关闭输出
func T(format string, args ...interface{}) {
	if !PrintTrace {
		return
	}
	_, file1, line1, _ := runtime.Caller(1)
	file1 = filepath.Base(file1)
	x := time.Now().Format("[15:04:05]")
	prefix := fmt.Sprintf("[\033[41m%v\033[0m \033[32m%v:%v\033[0m ] ", x, file1, line1)
	fmt.Printf(prefix+format+"\n", args...)
}
