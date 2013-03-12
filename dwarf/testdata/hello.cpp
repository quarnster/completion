#include <stdio.h>

class TestA {
public:
	void  publicVoidFunc();
	void  publicConstVoidFunc() const;
	FILE* publicFileFunc();
	FILE* publicConstFileFunc() const;

	static void publicStaticVoidFunc();
	static FILE* publicStaticFileFunc();

	int publicIntMember;
	FILE* publicFileMember;

private:
	void  privateVoidFunc();
	void  privateConstVoidFunc() const;
	FILE* privateFileFunc();
	FILE* privateConstFileFunc() const;

	static void privateStaticVoidFunc();
	static FILE* privateStaticFileFunc();

	int privateIntMember;
	FILE* privateFileMember;

protected:
	void  protectedVoidFunc();
	void  protectedConstVoidFunc() const;
	FILE* protectedFileFunc();
	FILE* protectedConstFileFunc() const;

	static void protectedStaticVoidFunc();
	static FILE* protectedStaticFileFunc();

	int protectedIntMember;
	FILE* protectedFileMember;
};


FILE* TestA::publicStaticFileFunc() {
	return NULL;
}

int main(int argc, char** argv) {
	TestA a;
	printf("Hello world!\n");
	return 0;
}
