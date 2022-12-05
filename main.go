package main

import (
	"fmt"

	"github.com/bjartek/overflow"
)

func main() {

	o := overflow.Overflow(
		overflow.WithNetwork("embedded"),
		overflow.WithPanicOnError(),
		overflow.WithPrintResults(),
	)

	aliceSigner := overflow.WithSigner("alice")

	fmt.Println("We setup the Example token FT for the service account user")
	o.Tx("setup_account", overflow.WithSignerServiceAccount())

	fmt.Println("We setup the Example token FT for alice")
	o.Tx("setup_account", aliceSigner)

	fmt.Println("We create a Forwarder for Alice to send Example token back to service accont")
	o.Tx("create_forwarder", aliceSigner, overflow.WithArg("receiver", "account"))

	//Setup switchboard for alice
	fmt.Println("We setup switchboard for alice")
	o.Tx("switchboard/setup_account", aliceSigner)

	//Add alice vault example token to switchboard
	fmt.Println("We register our Example token forwader in alice switchboard, note the type of the receiver")
	o.Tx("switchboard/add_vault_capability", aliceSigner)

	fmt.Println("We mint some tokens to service account")
	o.Tx("mint_tokens",
		overflow.WithSignerServiceAccount(),
		overflow.WithArg("recipient", "account"),
		overflow.WithArg("amount", 42.0),
	)

	fmt.Println("Try to send funds to alice from serviceaccount using switchboard")
	o.Tx("switchboard/transfer_tokens",
		overflow.WithSignerServiceAccount(),
		overflow.WithArg("to", "alice"),
		overflow.WithArg("amount", 13.0),
		overflow.WithArg("receiverPath", "/public/GenericFTReceiver"),
	)

}
