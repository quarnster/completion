
.PHONY: rpc

CP = cp
RM = rm -rf

ignore_signatures = TypeSignature,TypeArgument,Entry,FormalTypeParameter
ignore_descriptors = Entry,ComponentType,FieldType,ObjectType
ignore_parser = Spacing,Pattern,Completion,Completions,ReturnType,Arguments,Argument,UnhandledStuff,KeyedStuff,Type,TemplateName,Function
ignore_expand_path = Operation,File
ignore_csharp = Complete,Junk,UsingDirectives,Primary,BOM,Spacing,Reference,Grouping,InnerScope,EndOfLine,Comment,LineComment,BlockComment,Loops,Code,SpacedIdentifier,SpacedAccess
ignore_scopes = Scope,SingleLineComment,MultiLineComment
ignore_simplify = NewLine,EOF,Spacing,End,Impl,Contents,Simplified

PEGS = \
	../java/signatures/signatures.go \
	../java/descriptors/descriptors.go \
	../clang/parser/parser.go \
	../util/expand_path/expand_path.go \
	../util/scopes/scopes.go \
	../util/simplify/simplify.go \
	../net/csharp/csharp.go
ARCHIVE_DATA= ../3rdparty/jsonrpc.py $(wildcard ../editor/sublime/*.py) $(wildcard ../editor/sublime/*.sublime-keymap) $(wildcard ../editor/sublime/*.sublime-menu) ../README.md
GENERATED = $(PEGS) ../editor/archive_data.go


all: ../../parser $(PEGS) ../editor/archive_data.go completion

../../parser:
	go get github.com/quarnster/parser
	go install github.com/quarnster/parser/pegparser

../editor/archive_data.go: $(ARCHIVE_DATA)
	go run archive.go ${ARCHIVE_DATA}

completion: $(GENERATED)
	go build ../

test: completion
	go test -tags=nolibclang github.com/quarnster/completion/...

clean:
	$(RM) $(GENERATED) completion $(subst .go,,$(notdir $(PEGS)))

include ../../parser/PegRules.make
