package main

import(

)

const (
	TYPE_UNIX	= "unix"
	TYPE_SQL	= "sql"


	MODE_LIST	= "list"
	MODE_MODIFY	= "modify"
	MODE_ADD	= "add"
	MODE_REMOVE	= "remove"
	MODE_EXISTS	= "exists"
	MODE_AUTHENTICATE	= "authenticate"

	// generic errors pointing to mode
	EXIT_MAIN	= 99
	EXIT_LIST	= 100
	EXIT_MODIFY	= 101
	EXIT_ADD	= 102
	EXIT_REMOVE	= 103
	EXIT_AUTHENTICATE = 104
	EXIT_LIBRARY_ERROR = 105

	// usage error
	EXIT_USER_EMPTY		= 110
	EXIT_READ_PASSWORD	= 111
	EXIT_UNKNOWN_TYPE	= 112
	EXIT_UNKNOWN_MODE	= 113

	// edit errors
	EXIT_USER_UNKNOWN		= 120
	EXIT_CHANGE_PASSWORD_FAILED	= 121

	// add errors
	EXIT_USER_EXISTS		= 130
	EXIT_USER_CREATE		= 131

	// authenticate errors
	EXIT_WRONG_AUTHENTICATION	= 140
)
