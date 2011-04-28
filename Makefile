include $(GOROOT)/src/Make.inc

TARG=murmurhash3
CGOFILES=\
	murmurhash3.go

CGO_LDFLAGS=-lMurmurhash3

include $(GOROOT)/src/Make.pkg
