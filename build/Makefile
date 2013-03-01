# Maybe we just create a "dist" utility in the future?
DIRS = $(shell find .. -type d -not -path "../.git*" -not -path "../build*" -not -path "*/testdata")
PACKAGES = $(patsubst ..%,github.com/quarnster/completion%, $(DIRS))

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
