all:
	c-for-go vcx.yml
clean:
	rm -f vcx/cgo_helpers.go vcx/cgo_helpers.h vcx/cgo_helpers.c
	rm -r vcx/doc.go vcx/types.go vcx/const.go
	rm -f vcx/vcx.go
test:
	cd vcx && go build