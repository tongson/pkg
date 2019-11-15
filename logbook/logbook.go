package logbook

import (
	"fmt"
	"os"

	"aux"
)

// LogEntry is ...
type LogEntryT struct {
	Repo      string
	L1        string
	L2        string
	L3        string
	Stdout    string
	Stderr    string
	Arguments string
	Notes     string
}

// AddEntry is ...
func AddEntry(r LogEntryT) error {
	env := []string{"LC_ALL=C"}
	isDir := aux.StatPath("directory")
	if !isDir(r.L1) {
		clone := aux.RunArgs{Exe: "git", Env: env, Args: []string{"clone", r.Repo}}
		ret, _, stderr := aux.RunCmd(clone)
		if !ret {
			return fmt.Errorf("Failure cloning repository: ", stderr)
		}
	}
	fpath := fmt.Sprintf("%s/%s/%s", r.L1, r.L2, r.L3)
	if !isDir(fpath) {
		_ = os.MkdirAll(fpath, os.ModePerm)
	}
	ss := aux.RunArgs{Exe: "ss", Env: env, Args: []string{"-O", "-H", "-n", "-A", "tcp,udp"}}
	ret, sso, ssx := aux.RunCmd(ss)
	if !ret {
		sso = fmt.Sprintf("Failure running `ss`: %s", ssx)
	}
	_ = aux.StringToFile(fmt.Sprintf("%s/network", fpath), sso)
	pstree := aux.RunArgs{Exe: "pstree", Env: env, Args: []string{"-n", "-T", "-u", "-A", "-l", "-p", "-s"}}
	ret, pstreeo, pstreex := aux.RunCmd(pstree)
	if !ret {
		pstreeo = fmt.Sprintf("Failure running `pstree`: %s", pstreex)
	}
	_ = aux.StringToFile(fmt.Sprintf("%s/process_tree", fpath), pstreeo)
	ps := aux.RunArgs{Exe: "ps", Env: env, Args: []string{"-A", "-o", "pcpu,user,uid,pid,tty,args", "--sort=-pcpu"}}
	ret, pso, psx := aux.RunCmd(ps)
	if !ret {
		pso = fmt.Sprintf("Failure running `ps`: %s", psx)
	}
	_ = aux.StringToFile(fmt.Sprintf("%s/process_list", fpath), pso)
	_ = aux.StringToFile(fmt.Sprintf("%s/stdout", fpath), r.Stdout)
	_ = aux.StringToFile(fmt.Sprintf("%s/stderr", fpath), r.Stderr)
	if len(r.Arguments) == 0 {
		r.Arguments = ""
	}
	_ = aux.StringToFile(fmt.Sprintf("%s/arguments", fpath), r.Arguments)
	if len(r.Notes) == 0 {
		r.Notes = ""
	}
	_ = aux.StringToFile(fmt.Sprintf("%s/notes", fpath), r.Notes)

	add := aux.RunArgs{Exe: "git", Dir: r.L1, Env: env, Args: []string{"add", "-A"}}
	ret, _, stderr := aux.RunCmd(add)
	if !ret {
		return fmt.Errorf("Failure adding from git: ", stderr)
	}
	commit := aux.RunArgs{Exe: "git", Dir: r.L1, Env: env, Args: []string{"commit", "-m", fmt.Sprintf("%s(%s): %s", r.L2, r.L3, r.Arguments)}}
	ret, _, stderr = aux.RunCmd(commit)
	if !ret {
		return fmt.Errorf("Failure committing from git: ", stderr)
	}
	push := aux.RunArgs{Exe: "git", Dir: r.L1, Env: env, Args: []string{"push"}}
	ret, _, stderr = aux.RunCmd(push)
	if !ret {
		return fmt.Errorf("Failure pushing from git: ", stderr)
	}
	return nil
}
