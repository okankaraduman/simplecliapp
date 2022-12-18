run-parallel-flag:
	@go build
	@./cliapp -parallel 10 adjust.com google.com facebook.com yahoo.com yandex.com twitter.com
run-default:
	@go build
	@./cliapp adjust.com google.com facebook.com yahoo.com yandex.com twitter.com