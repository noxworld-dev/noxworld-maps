module nw_ix

go 1.21

toolchain go1.22.4

replace github.com/noxworld-dev/noxworld-maps/noxworld => ../../noxworld

require (
	github.com/noxworld-dev/noxscript/ns/v4 v4.21.0
	github.com/noxworld-dev/noxworld-maps/noxworld v0.0.0-00010101000000-000000000000
	github.com/noxworld-dev/opennox-lib v0.0.0-20241111104024-77c44f5da21b
)
