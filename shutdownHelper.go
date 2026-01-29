package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"
)

// 核心函数：执行关机命令（兼容跨平台）
func shutdownPC() error {
	var cmd *exec.Cmd
	osType := runtime.GOOS // 获取当前操作系统类型：windows、linux、darwin（macOS）

	switch osType {
	case "windows":
		// Windows 立即关机命令，/s 表示关机，/t 0 表示延迟0秒执行
		cmd = exec.Command("shutdown", "/s", "/t", "0")
	case "linux", "darwin":
		// Linux/macOS 立即关机命令，-h now 表示立即停止（关机）
		cmd = exec.Command("shutdown", "-h", "now")
	default:
		return fmt.Errorf("不支持的操作系统：%s", osType)
	}

	// 执行关机命令
	return cmd.Run()
}

// 倒计时展示函数
func countdown(minutes int) {
	totalSeconds := minutes * 60
	for i := totalSeconds; i > 0; i-- {
		// 格式化剩余时间（分:秒），覆盖当前行输出，更整洁
		fmt.Printf("\r剩余关机时间：%02d:%02d（按 Ctrl+C 取消关机）", i/60, i%60)
		time.Sleep(1 * time.Second)
	}
	// 倒计时结束后清空该行提示
	fmt.Println("\r倒计时结束，即将执行关机操作...")
}

func main() {
	fmt.Println("===== Go 电脑定时关机工具 =====")
	fmt.Println("说明：1. 输入倒计时分钟数（正整数）；2. 中途可按 Ctrl+C 取消关机；3. 需管理员/root 权限才能执行关机")

	// 1. 读取用户输入的倒计时分钟数
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("\n请输入倒计时分钟数：")
	scanner.Scan()
	input := strings.TrimSpace(scanner.Text())

	// 2. 合法性校验：转换为整数并判断是否为正整数
	minutes, err := strconv.Atoi(input)
	if err != nil || minutes <= 0 {
		fmt.Printf("输入错误！请输入大于 0 的整数，错误信息：%v\n", err)
		return
	}

	// 3. 提示用户即将开始倒计时
	fmt.Printf("\n已确认：%d 分钟后将自动关闭电脑\n", minutes)
	time.Sleep(2 * time.Second) // 给用户 2 秒反应时间

	// 4. 执行倒计时
	countdown(minutes)

	// 5. 倒计时结束，执行关机命令
	err = shutdownPC()
	if err != nil {
		fmt.Printf("执行关机失败！请检查是否拥有管理员/root 权限，错误信息：%v\n", err)
		return
	}

	// 6. 关机命令执行成功提示（若系统正常响应，此处可能不会被打印，因为电脑已开始关机流程）
	fmt.Println("关机命令已成功发送，电脑即将关闭...")
}