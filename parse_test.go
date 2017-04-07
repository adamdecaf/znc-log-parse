package main

import (
	"testing"
)

func TestParsing__Join(t *testing.T) {
	ex := "[02:33:31] *** Joins: adam (adam@Snoonet-fdl.i3c.1b1g5k.IP)"
	if !isJoin(ex) {
		t.Fatalf("join not matched")
	}
}

func TestParsing__Part(t *testing.T) {
	ex := "[03:14:57] *** Parts: bo4tdude[penis] (Bo4t@user/bo4tdude) (Leaving)"
	if !isPart(ex) {
		t.Fatalf("part not matched")
	}
}

func TestParsing__Quit(t *testing.T) {
	ex := "[02:33:19] *** Quits: bo4tdude (Bo4t@user/bo4tdude) (Quit: Leaving)"
	if !isQuit(ex) {
		t.Fatalf("quit not matched")
	}
}

func TestParsing__Rename(t *testing.T) {
	ex := "[02:34:32] *** adam is now known as Snoo60230"
	if !isRename(ex) {
		t.Fatalf("rename not matched")
	}
}
