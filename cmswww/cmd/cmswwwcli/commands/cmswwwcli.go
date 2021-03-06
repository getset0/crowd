package commands

import (
	"fmt"
	"strings"

	"github.com/decred/contractor-mgmt/cmswww/cmd/cmswwwcli/client"
	"github.com/decred/contractor-mgmt/cmswww/cmd/cmswwwcli/config"
)

type Options struct {
	// cli flags
	Host       func(string) error `long:"host" description:"cmswww host"`
	JSONOutput func()             `long:"jsonout" description:"Output only the last command's JSON output; use this option when writing scripts"`
	Verbose    func()             `short:"v" long:"verbose" description:"Print request and response details"`

	// cli commands
	Login                   LoginCmd                   `command:"login" description:"Login to the contractor mgmt system.\n\n           Parameters: <email> <password>\n  --------------------------------------"`
	Logout                  LogoutCmd                  `command:"logout" description:"Logout of the contractor mgmt system. Parameters: none\n  --------------------------------------"`
	NewIdentity             NewIdentityCmd             `command:"newidentity" description:"Generate a new identity. Parameters: none\n  --------------------------------------"`
	VerifyNewIdentity       VerifyIdentityCmd          `command:"verifyidentity" description:"Verify a newly generated identity.\n\n           Parameters: <token>\n  --------------------------------------"`
	Register                RegisterCmd                `command:"register" description:"Complete registration as a contractor.\n\n           Parameters: <email> <username> <password> <token>\n  --------------------------------------"`
	Policy                  PolicyCmd                  `command:"policy" description:"Fetch server policy. Parameters: none\n  --------------------------------------"`
	Version                 VersionCmd                 `command:"version" description:"Fetch server info and CSRF token. Parameters: none\n  --------------------------------------"`
	InviteNewUser           InviteNewUserCmd           `command:"invite" description:"Send a new contractor invitation.\n\n           Parameters: <email>\n  --------------------------------------"`
	UserDetails             UserDetailsCmd             `command:"user" description:"Fetch a user's details given the user id.\n\n           Parameters: <user id/email/username>\n  --------------------------------------"`
	ManageUser              ManageUserCmd              `command:"manageuser" description:"Manage a user by user id.\n\n           Parameters: <user id/email/username> <action> <reason>\n    Available actions: resendinvite, resendidentitytoken, lock, unlock\n  --------------------------------------"`
	EditUser                EditUserCmd                `command:"edituser" description:"Edit a user's details.\n\n           Parameters: [ --name <name> ] [ --location <location> ]\n  --------------------------------------"`
	UpdateExtendedPublicKey UpdateExtendedPublicKeyCmd `command:"updatexpublickey" description:"Edit a user's extended public key.\n\n           Parameters: [ --token <verification token> ] [ --xpubkey <xpubkey> ]\n  --------------------------------------"`
	ChangePassword          ChangePasswordCmd          `command:"changepassword" description:"Change your password.\n\n           Parameters: <current password> <new password>\n  --------------------------------------"`
	ResetPassword           ResetPasswordCmd           `command:"resetpassword" description:"Reset your password.\n\n           Parameters: <email> <new password>\n  --------------------------------------"`
	SubmitInvoice           SubmitInvoiceCmd           `command:"submitinvoice" description:"Submits an invoice for a given month and year.\n\n           Parameters: <month> <year>\n  --------------------------------------"`
	InvoiceDetails          InvoiceDetailsCmd          `command:"invoice" description:"Displays an invoice's details.\n\n           Parameters: <token>\n  --------------------------------------"`
	Invoices                InvoicesCmd                `command:"invoices" description:"Lists invoices with a particular status for a given month and year.\n\n           Parameters: <month> <year> [ --status <status> ]\n   Available statuses: unreviewed, rejected, approved, paid\n  --------------------------------------"`
	MyInvoices              MyInvoicesCmd              `command:"myinvoices" description:"Lists a user's invoices with a particular status.\n\n           Parameters: [status]\n   Available statuses: unreviewed, rejected, approved, paid\n  --------------------------------------"`
	SetInvoiceStatus        SetInvoiceStatusCmd        `command:"setinvoicestatus" description:"Changes an invoice's status.\n\n           Parameters: <token> <status>\n   Available statuses: rejected, approved, paid\n  --------------------------------------"`
	LogWork                 LogWorkCmd                 `command:"logwork" description:"Adds a line item to an invoice.\n\n           Parameters: <month> <year>\n  --------------------------------------"`
	DCRUSD                  DCRUSDCmd                  `command:"dcrusd" description:"Calculates the DCR-USD for a given month & year.\n\n           Parameters: <month> <year>\n  --------------------------------------"`
	ReviewInvoices          ReviewInvoicesCmd          `command:"reviewinvoices" description:"Generates a list of submitted invoices that are ready for initial review.\n\n           Parameters: <month> <year>\n  --------------------------------------"`
	PayInvoices             PayInvoicesCmd             `command:"payinvoices" description:"Generates a list of unpaid invoices that are ready for payment.\n\n           Parameters: <month> <year> <DCR-USD rate>\n  --------------------------------------"`
}

var Ctx *client.Ctx
var Opts Options

func SetupOptsFunctions() {
	Opts.Host = func(host string) error {
		if !strings.HasPrefix(host, "http://") && !strings.HasPrefix(host, "https://") {
			return fmt.Errorf("host must begin with http:// or https://")
		}

		config.Host = host

		if err := config.LoadCsrf(); err != nil {
			return err
		}

		return config.LoadCookies()
	}

	Opts.JSONOutput = func() {
		config.JSONOutput = true
	}

	Opts.Verbose = func() {
		config.Verbose = true
	}
}
