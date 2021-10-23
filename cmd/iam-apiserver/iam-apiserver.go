package main

import (
	iam_apiserver "github.com/mingyuans/go-layout/internal/iam-apiserver"
	"math/rand"
	"os"
	"runtime"
	"time"
)

func main() {
	// 全局初始化一次 Seed, 保障后续使用 rand 进行随机数生成不会有重复的问题；
	rand.Seed(time.Now().UTC().UnixNano())
	// 设置最大进程数量；
	if (len(os.Getenv("GOMAXPROCS"))) == 0 {
		runtime.GOMAXPROCS(runtime.NumCPU())
	}

	//basename 最好保持跟这个  internal/iam-apiserver 这个文件夹名一致；
	//因为这个 basename 会作为 cmd 命令，而 MK 编译处理的 exec 文件是按文件名来的。
	iam_apiserver.NewApp("iam-apiserver").Run()
}
