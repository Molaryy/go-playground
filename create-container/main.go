package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func runContainer(args []string) {
	fmt.Printf("Running %v as PID %d\n", args[2:], os.Getppid())
	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, args[2:]...)...)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID,
	}

	isFuncPanic(cmd.Run())
}

func runContainerChild(args []string) {

	cmd := exec.Command(args[2], args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	isFuncPanic(syscall.Chroot(os.Getenv("HOME")))
	isFuncPanic(os.Chdir("/"))
	isFuncPanic(syscall.Mount("proc", "proc", "proc", 0, ""))
	isFuncPanic(cmd.Run())

	isFuncPanic(cmd.Run())
}

func messageHelper(cmds []string, opts []string) {
	cmdsDescriptions := []string{
		"you can start a container",
		"run as a child",
	}
	optsDescriptions := []string{
		"provides message helper",
	}

	fmt.Println("Usage: jb-docker [COMMAND] [OPTIONS]")
	fmt.Println("\nAvailable commands:")
	for idxCmd := range cmds {
		fmt.Println(fmt.Sprintf("[%s] - %s", cmds[idxCmd], cmdsDescriptions[idxCmd]))
	}
	fmt.Println("\nAvailable options:")
	for idxOpt := range opts {
		fmt.Println(fmt.Sprintf("[%s] - %s", opts[idxOpt], optsDescriptions[idxOpt]))
	}
}

func startContainer(args []string) {
	cmds := []string{"run", "child"}
	opts := []string{"--help"}

	if len(args) < 2 {
		panic("Provide at least 1 argument")
	}
	switch args[1] {
	case cmds[0]:
		runContainer(args)
	case cmds[1]:
		runContainerChild(args)
	case opts[0]:
		messageHelper(cmds, opts)
	default:
		fmt.Println(fmt.Sprintf("jb-docker: '%s' is not a jb-docker command.\nSee 'jb-docker --help'", args[1]))
	}
}

func isFuncPanic(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	startContainer(os.Args)
}
