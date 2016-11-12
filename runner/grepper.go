package runner

import "os/exec"

type Grepper struct {
}

func (Grepper) FileAlreadyHasStoryID(filePath string) bool {
	cmd := "grep"
	args := []string{"\\[*#[0-9]\\+\\]", filePath}
	grepMatchFound := exec.Command(cmd, args...).Run() == nil
	return grepMatchFound
}
