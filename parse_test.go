package main

import (
	"testing"
)

// [14:55:34] *** brd\ sets mode: +s

func TestParsing__Join(t *testing.T) {
	exs := []string{
		"[02:33:31] *** Joins: adam (adam@Snoonet-fdl.i3c.1b1g5k.IP)",
		"[19:07:43] *** Joins: RealJustinLong (RealJustinL@user/RealJustinLong)",
		"[13:47:51] *** Joins: toastedpe (~dchristen@104-185-27-156.lightspeed.cicril.sbcglobal.net)",
	}
	for i := range exs {
		if !isJoin(exs[i]) {
			t.Fatalf("join not matched -- %s", exs[i])
		}
	}

	// negative match
	if isJoin("bad") {
		t.Fatal("'bad' is not a join")
	}
}

func TestParsing__Part(t *testing.T) {
	exs := []string{
		"[03:14:57] *** Parts: bo4tdude[penis] (Bo4t@user/bo4tdude) (Leaving)",
		"[13:42:18] *** Parts: timmc (~timmc@209-6-68-29.c3-0.abr-ubr1.sbo-abr.ma.cable.rcn.com) ()",
	}
	for i := range exs {
		if !isPart(exs[i]) {
			t.Fatalf("part not matched -- %s", exs[i])
		}
	}

	// negative match
	if isJoin("bad") {
		t.Fatal("'bad' is not a part")
	}
}

func TestParsing__Quit(t *testing.T) {
	exs := []string{
		"[02:33:19] *** Quits: bo4tdude (Bo4t@user/bo4tdude) (Quit: Leaving)",
		"[19:07:32] *** Quits: RealJustinLong (RealJustinL@user/RealJustinLong) (Connection closed)",
		"[19:42:47] *** Quits: RealJustinLong (RealJustinL@user/RealJustinLong) (Ping timeout: 121 seconds)",
		"[14:15:07] *** Quits: toastedpe (~dchristen@104-185-27-156.lightspeed.cicril.sbcglobal.net) (Read error: Operation timed out)",
	}
	for i := range exs {
		if !isQuit(exs[i]) {
			t.Fatalf("quit not matched -- %s", exs[i])
		}
	}

	// negative match
	if isJoin("bad") {
		t.Fatal("'bad' is not a quit")
	}
}

func TestParsing__Rename(t *testing.T) {
	exs := []string{
		"[02:34:32] *** adam is now known as Snoo60230",
		"[07:17:05] *** ep0n is now known as epon",
	}
	for i := range exs {
		if !isRename(exs[i]) {
			t.Fatalf("rename not matched -- %s", exs[i])
		}
	}

	// negative match
	if isJoin("bad") {
		t.Fatal("'bad' is not a rename")
	}
}
