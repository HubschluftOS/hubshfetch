package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

const (
	hlos = `
   /$$       /$$     %s@%s 
  | $$      | $$     %s 
  | $$$$$$$ | $$     %s 
  | $$__  $$| $$     %s 
  | $$  \ $$| $$     %s 
  | $$  | $$| $$     %s 
  |__/  |__/|__/     %s 

`
)

var (
	Bold   = "\033[1m"
	Yellow = "\033[33m"

	Reset = "\033[0m"
)

var (
	Username_output string

	Hostname_output string
	GetOS_output    string
	Kernel_output   string
	Uptime_output   string
	Packages_output string
	Shell_output    string
	Wm_output       string
)

func Username() {
	username, err := exec.Command("whoami").Output()
	if err != nil {
		log.Fatal(err)
	}

	Username_output = fmt.Sprintf(strings.TrimSpace(string(username)))
}

func Hostname() {
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}
	Hostname_output = fmt.Sprintf(hostname)
}

func GetOS() {
	if runtime.GOOS == "linux" {
		GetOS_output = fmt.Sprintf("GNU/Linux")
	}
}

func Kernel() {
	kernel, err := ioutil.ReadFile("/proc/version")
	if err != nil {
		log.Fatal(err)
	}
	KernelPart := strings.Fields(string(kernel))

	Kernel_output = fmt.Sprintf(strings.TrimSpace(KernelPart[2]))
}

func Uptime() {
	uptime, err := ioutil.ReadFile("/proc/uptime")
	if err != nil {
		log.Fatal(err)
	}
	UptimeParts := strings.Fields(string(uptime))
	UptimeToInt, err := strconv.ParseFloat(UptimeParts[0], 64)
	if err != nil {
		log.Fatal(err)
	}
	Uptime_output = fmt.Sprintf(strings.TrimSpace("%.2f HOURS"), UptimeToInt/3600)
}

func Packages() {
	packages, err := exec.Command("sh", "-c", "pacman -Q | wc -l").Output()
	if err != nil {
		log.Fatal(err)
	}
	Packages_output = fmt.Sprintf(strings.TrimSpace(string(packages)))
}

func Shell() {
	shell, err := exec.Command("sh", "-c", "basename ${SHELL}").Output()
	if err != nil {
		log.Fatal(err)
	}
	Shell_output = fmt.Sprintf(strings.TrimSpace(string(shell)))
}

func WM() {
	wm, err := exec.Command("sh", "-c", "echo $XDG_SESSION_DESKTOP").Output()
	if err != nil {
		log.Fatal(err)
	}
	Wm_output = fmt.Sprintf(strings.TrimSpace(string(wm)))
}

func main() {
	Username()
	Hostname()
	GetOS()
	Kernel()
	Uptime()
	Packages()
	Shell()
	WM()
	fmt.Printf(hlos,
		Yellow+Username_output+Reset,
		Yellow+Hostname_output+Reset,
		Yellow+"OS: "+Reset+GetOS_output,
		Yellow+"KERNEL: "+Reset+Kernel_output,
		Yellow+"UPTIME: "+Reset+Uptime_output,
		Yellow+"PACKAGES: "+Reset+Packages_output,
		Yellow+"SHELL: "+Reset+Shell_output,
		Yellow+"WM: "+Reset+Wm_output)

}
