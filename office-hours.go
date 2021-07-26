package main

import (
	"os"

	"github.com/algorand/go-algorand-sdk/types"
	"github.com/tmc/teal-examples/avm"
	"github.com/tmc/teal-examples/globals"
)

var txn types.Transaction

func main() {
	// ApplicationID is zero in inital creation txn
	// txn ApplicationID
	// bz handle_createapp
	if txn.ApplicationID == 0 {
		handle_createapp()
	}

	// Handle each possible OnCompletion type. We don't have to
	// worry about handling ClearState, because the
	// ClearStateProgram will execute in that case, not the
	// ApprovalProgram.

	// txn OnCompletion
	// int NoOp
	// ==
	// bnz handle_noop

	if txn.OnCompletion == types.NoOpOC {
		handle_noop()
	}

	// txn OnCompletion
	// int OptIn
	// ==
	// bnz handle_optin
	if txn.OnCompletion == types.OptInOC {
		handle_optin()
	}

	// txn OnCompletion
	// int CloseOut
	// ==
	// bnz handle_closeout
	if txn.OnCompletion == types.CloseOutOC {
		handle_closeout()
	}

	// txn OnCompletion
	// int UpdateApplication
	// ==
	// bnz handle_updateapp
	if txn.OnCompletion == types.UpdateApplicationOC {
		handle_updateapp()
	}

	// txn OnCompletion
	// int DeleteApplication
	// ==
	// bnz handle_deleteapp
	if txn.OnCompletion == types.DeleteApplicationOC {
		handle_deleteapp()
	}

	// Unexpected OnCompletion value. Should be unreachable.
	// err
	panic("Unexpected OnCompletion value. Should be unreachable.")
}

func handle_createapp() {
	// int 1
	// return
	os.Exit(1)
}
func handle_noop() {
	// txn ApplicationArgs 0
	// byte "triple"
	// ==
	// bnz triple
	if string(txn.ApplicationArgs[0]) == "triple" {
		triple()
	}

	// txn ApplicationArgs 0
	// byte "compare"
	// ==
	// bnz compare
	if string(txn.ApplicationArgs[0]) == "compare" {
		compare()
	}

	// unknown "method"
	// err
	panic("unknown method")
}

func handle_optin() {
	// A single txn with no args must be a simple optin.  Allow
	// all.  But if there are more transactions or arguments, fall
	// through to NOOP handling, as the user must want to use the
	// app immediately.

	// global GroupSize
	// int 1
	// ==
	// txn NumAppArgs
	// ||
	// bnz handle_noop
	if globals.GroupSize == 1 || len(txn.ApplicationArgs) != 0 {
		handle_noop()
	}

	// int 1
	// return
	os.Exit(1)
}
func handle_closeout() {
}
func handle_updateapp() {
}
func handle_deleteapp() {
}

func triple() {
	// txn ApplicationArgs 1
	// callsub triplearg
	x1 := triplearg(txn.ApplicationArgs[1])

	// txn ApplicationArgs 2
	// btoi
	x2 := avm.Btoi(txn.ApplicationArgs[2])

	// ==
	// return
	if x1 == x2 {
		os.Exit(1)
	}
	os.Exit(0)
}

func triplearg(arg []byte) int {
	// btoi
	// int 3
	// *
	// retsub
	x1 := avm.Btoi(arg)
	return x1 * 3
}

func compare() {}

/*

triple:

triplearg:

compare:
        txn ApplicationArgs 1   // --app-arg str:51
        callsub ctoi
        txn ApplicationArgs 2   // --app-arg int:51
        ==
        return

ctoi:
        int 0
        getbyte
        byte "0"
        int 0
        getbyte
        -
        retsub

lead_digit_value:               // "512"
        dup                     // "512" "512"
        callsub ctoi            // "512" 5
        swap                    // 5 "512"
        len                     // 5 3
        int 1                   // 5 3 1
        -                       // 5 2
        int 10                  // 5 2 10
        swap                    // 5 10 2
        exp                     // 5 100
        *                       // 500
        retsub

atoi:                           // "512"
        dup                     // "512"; "512"
        len                     // "512"; 3
        int 1                   // "512"; 3; 1
        ==                      // "512"; 0
        bz longer               // "512" -> longer
        callsub ctoi
        retsub
longer:
        dup                       // "512"; "512"
        callsub lead_digit_value  // "512"; 500
        swap                      // 500; "512"
        dup                       // 500; "512"; "512
        len                       // 500; "512"; 3
        int 1                     // 500; "512"; 3; 1
        swap                      // 500; "512"; 1; 3
        substring3                // 500; "12"
        callsub atoi              // 500; 12
        +                         // 512
        retsub

atoi_i:                         // "512"
        store 0                 // 0: string argument
        int 0
        store 1                 // 1: index
        int 0
        store 2                 // 2: answer

loop:   load 1
        load 0
        len
        >=
        bnz done
        load 2
        int 10
        *
        store 2

        load 0
        load 1
        getbyte
        int 48
        -

        load 2
        +
        store 2

        load 1
        int 1
        +
        store 1

        b loop

done:
        load 2
        retsub


handle_closeout:
        int 1
        return

        // Allow updates and deletes only from original creator
handle_updateapp:
handle_deleteapp:
        global CreatorAddress
        txn Sender
        ==
        return

bad:
        int 0
        return

*/
