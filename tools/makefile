.PHONY: install-tools
install-tools:
	go install \
	../cmd/import-list \
	github.com/sanemat/go-xgoinstall/cmd/x-go-install \
	; \
	import-list -z tools.go | x-go-install -0
