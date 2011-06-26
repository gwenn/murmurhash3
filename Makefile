include $(GOROOT)/src/Make.inc

TARG=murmurhash3
CGOFILES=\
	murmurhash3.go

# TODO Create two files (one for the amd64 platform and another for th 386 platform
# See pkg/big...

CGO_LDFLAGS=-lMurmurHash3

include $(GOROOT)/src/Make.pkg
