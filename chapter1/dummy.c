#include <stdio.h>

void func1();
void func2(const char *s);

int main(int argc, char* argv[])
{
	func1();
	func2("World");
}

void func1()
{
	printf("This is func 1\n");
}

void func2(const char *s)
{
	printf("Hello %s\n", s);
}
