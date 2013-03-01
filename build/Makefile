# Maybe we just create a "dist" utility in the future?
PACKAGES = github.com/quarnster/completion \
	github.com/quarnster/completion/common \
	github.com/quarnster/completion/content \
	github.com/quarnster/completion/java \
	github.com/quarnster/completion/java/descriptors \
	github.com/quarnster/completion/java/parse \
	github.com/quarnster/completion/net \
	github.com/quarnster/completion/util

ifeq ($V, 1)
	VERBOSE = -v
endif

all: ../java/descriptors/descriptors.go test

parser_exe:
	go build -o parser_exe github.com/quarnster/parser/exe

descriptors/descriptors.go: parser_exe ../java/descriptors/descriptors.peg Makefile
	./parser_exe -peg=../java/descriptors/descriptors.peg -notest -ignore "Entry,ComponentType,FieldType,ObjectType" -testfile="none silly"

../java/descriptors/descriptors.go: descriptors/descriptors.go
	cp $< $@

test:
	go test $(VERBOSE) $(PACKAGES)
