.PHONY: build
build:
	@echo building
	@GOARM=6 GOARCH=arm GOOS=linux go build -o bin/out .

.PHONY: scp
scp:
	@echo copying
	@sudo scp bin/out victor@192.168.6.115:/home/victor