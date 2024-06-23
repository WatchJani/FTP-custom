package server

import "testing"

func TestParser(t *testing.T) {
	payload := []byte("janko")

	cmd, args := "janko", ""

	if realCmd, realArgs := Parser(payload); cmd != realCmd && args != realArgs {
		t.Errorf("Args [%s : %s] | Cmd [%s : %s]", args, realArgs, cmd, realCmd)
	}
}
