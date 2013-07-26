
.SUFFIXES: .peg .go

.PHONY: rpc

EXE_EXTENSION = 
CP = cp
RM = rm -rf

buildPeg = ./parser_exe "-peg=$(1)" -notest -ignore="$(2)" -testfile "none silly" && $(CP) $(subst .go,,$(notdir $@))/$(notdir $@) $@

ignore_signatures = TypeSignature,TypeArgument,Entry,FormalTypeParameter
ignore_descriptors = Entry,ComponentType,FieldType,ObjectType
ignore_parser = Spacing,Pattern,Completion,Completions,ReturnType,Arguments,Argument,UnhandledStuff,KeyedStuff,Type,TemplateName,Function
ignore_expression = Spacing,Primary,Op,Expression,Grouping,BooleanOp
ignore_expand_path = Operation,File
ignore_csharp = Complete,Junk,UsingDirectives,Primary,BOM,Spacing,Reference,Grouping,InnerScope,EndOfLine,Comment,LineComment,BlockComment,Loops,Code,SpacedIdentifier,SpacedAccess
ignore_scopes = Scope,CommentScope,TextScope

PEGS = \
	../java/signatures/signatures.go \
	../java/descriptors/descriptors.go \
	../clang/parser/parser.go \
	../util/expression/expression.go \
	../util/expand_path/expand_path.go \
	../util/scopes/scopes.go \
	../net/csharp/csharp.go
ARCHIVE_DATA= ../3rdparty/jsonrpc.py ../editor/sublime/plugin.py
GENERATED = $(PEGS) ../editor/archive_data.go 


all: ../../parser parser_exe $(PEGS) ../editor/archive_data.go rpc

../../parser:
	go get github.com/quarnster/parser

../editor/archive_data.go: $(ARCHIVE_DATA)
	go run archive.go $(ARCHIVE_DATA)

rpc: $(GENERATED)
	go build -o rpc ../rpc

parser_exe: ../../parser
	go build -o parser_exe$(EXE_EXTENSION) github.com/quarnster/parser/exe

.peg.go: parser_exe
	$(call buildPeg,$<,$(ignore_$(subst .go,,$(notdir $@))))

test: rpc
	go run runtests.go

clean: 
	$(RM) $(GENERATED) rpc parser_exe $(subst .go,,$(notdir $(PEGS)))
