package logging

import(

)

const (
	TIME_FORMAT = "2006-01-02T15:04:05Z07:00"
	DATE_FORMAT = "2006-01-02"
	CLOCK_FORMAT = "15:04:05"


	// cannot continue, aborting immidiatly
	EMERGENCY =		0x01
	EMERGENCY_SYSLOG =	1
	// require immidiate attention
	CRITICAL =		0x02
	CRITICAL_SYSLOG =	2
	// not used
	ALERT =			CRITICAL
	ALERT_SYSLOG =		3
	// an error that was recovered
	ERROR =			0x04
	ERROR_SYSLOG =		4
	// warning. space nearly full etc.
	WARNING =		0x08
	WARNING_SYSLOG =	5
	// notice. behaviour will change, found this driver etc.
	NOTICE =		0x16
	NOTICE_SYSLOG =		6
	// info. version info, something succeded etc.
	INFO =			0x32
	INFO_SYSLOG =		7
	// debug. favourite time of programming
	DEBUG =			0x64
	DEBUG_SYSLOG =		8


	// EMERGENCY & CRITICAL
	FATAL =		0x03
	// FATAL & ERROR
	ERRORS =	0x07
	// ERRORS & WARNING
	WARNINGS =	0x15
	// WARNINGS & NOTICE
	VERBOSE =	0x31
	// VERBOSE & INFO
	ATTENTION =	0x63
	// VERBOSE & DEBUG
	DEBUGGING =	0x7f
	// max
	EVERYTHING =	0xff
)

