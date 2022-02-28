module golib
go 1.17

require (
	jwt v0.0.0
	golang.org/x/text v0.3.7
	ini v0.0.0
)

replace (
	jwt => ./mod/jwt
	ini => ./mod/ini
)