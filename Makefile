.PHONY: gci

gci:
	@GO111MODULE=off go get github.com/daixiang0/gci
	gci -w -local github.com/phanletrunghieu/pkg ./..