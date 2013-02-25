# Maybe we just create a "dist" utility in the future?

all: ../java/descriptors/descriptors.go

parser_exe:
	go build -o parser_exe github.com/quarnster/parser/exe

descriptors/descriptors.go: parser_exe ../java/descriptors/descriptors.peg Makefile
	./parser_exe -peg=../java/descriptors/descriptors.peg -notest -ignore "Entry,ComponentType,FieldType,ObjectType" -testfile="none silly"

../java/descriptors/descriptors.go: descriptors/descriptors.go
	cp $< $@
