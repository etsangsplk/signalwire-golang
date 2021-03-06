GO_BIN ?= go
RM_BIN ?= rm

export PATH := $(PATH):/usr/local/go/bin

linter-install:
	$(RM_BIN) -f go.mod go.sum
	$(GO_BIN) mod init linter
	$(GO_BIN) get github.com/golangci/golangci-lint/cmd/golangci-lint@v1.21.0
	$(GO_BIN) get -u github.com/mgechev/revive
	$(RM_BIN) -f go.mod go.sum

test:
	make test -C signalwire

lint:
	make lint -C signalwire
	make lint -C RelayTests/CallComplex
	make lint -C RelayTests/CallOutbound
	make lint -C RelayTests/CallInbound
	make lint -C RelayTests/CallRecord
	make lint -C RelayExamples/Outbound
	make lint -C RelayExamples/Inbound
	make lint -C RelayExamples/PlayAsync
	make lint -C RelayExamples/RecordAsync
	make lint -C RelayExamples/RecordMultipleAsync
	make lint -C RelayExamples/PlayMultipleAsync
	make lint -C RelayExamples/Detect
	make lint -C RelayExamples/ReceiveFaxAsync
	make lint -C RelayExamples/SendFax
	make lint -C RelayExamples/RecordBlocking
	make lint -C RelayExamples/ReceiveFaxBlocking
	make lint -C RelayExamples/Connect
	make lint -C RelayExamples/Tap
	make lint -C RelayExamples/SendDigits
	make lint -C RelayExamples/MessageSend
	make lint -C RelayExamples/MessageReceive
	make lint -C RelayExamples/DeliverTask
	make lint -C RelayExamples/ClientConnectStress

update:
	make update -C signalwire
	make update -C RelayTests/CallComplex
	make update -C RelayTests/CallOutbound
	make update -C RelayTests/CallInbound
	make update -C RelayTests/CallRecord
	make update -C RelayExamples/Outbound
	make update -C RelayExamples/Inbound
	make update -C RelayExamples/PlayAsync
	make update -C RelayExamples/RecordAsync
	make update -C RelayExamples/RecordMultipleAsync
	make update -C RelayExamples/PlayMultipleAsync
	make update -C RelayExamples/Detect
	make update -C RelayExamples/ReceiveFaxAsync
	make update -C RelayExamples/SendFax
	make update -C RelayExamples/RecordBlocking
	make update -C RelayExamples/ReceiveFaxBlocking
	make update -C RelayExamples/Connect
	make update -C RelayExamples/Tap
	make update -C RelayExamples/SendDigits
	make update -C RelayExamples/MessageSend
	make update -C RelayExamples/MessageReceive
	make update -C RelayExamples/DeliverTask
	make update -C RelayExamples/ClientConnectStress

