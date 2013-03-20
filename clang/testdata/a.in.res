Fields:
	int argc                                                                         // 
	char ** argv                                                                     // 
	const char *const [] sys_errlist                                                 // 
	const int sys_nerr                                                               // 
Methods:
	int  asprintf(char ** , const char *, ... )                                      // 
	void  clearerr(FILE * )                                                          // 
	char *  ctermid(char * )                                                         // 
	char *  ctermid_r(char * )                                                       // 
	int  dprintf(int , const char *, ... )                                           // 
	int  fclose(FILE * )                                                             // 
	FILE *  fdopen(int , const char * )                                              // 
	int  feof(FILE * )                                                               // 
	int  ferror(FILE * )                                                             // 
	int  fflush(FILE * )                                                             // 
	int  fgetc(FILE * )                                                              // 
	char *  fgetln(FILE * , size_t * )                                               // 
	int  fgetpos(FILE * , fpos_t * )                                                 // 
	char *  fgets(char * , int , FILE * )                                            // 
	int  fileno(FILE * )                                                             // 
	void  flockfile(FILE * )                                                         // 
	const char *  fmtcheck(const char * , const char * )                             // 
	FILE *  fopen(const char * , const char * )                                      // 
	int  fprintf(FILE * , const char *, ... )                                        // 
	int  fpurge(FILE * )                                                             // 
	int  fputc(int , FILE * )                                                        // 
	int  fputs(const char * , FILE * )                                               // 
	size_t  fread(void * , size_t , size_t , FILE * )                                // 
	FILE *  freopen(const char * , const char * , FILE * )                           // 
	int  fscanf(FILE * , const char *, ... )                                         // 
	int  fseek(FILE * , long , int )                                                 // 
	int  fseeko(FILE * , off_t , int )                                               // 
	int  fsetpos(FILE * , const fpos_t * )                                           // 
	long  ftell(FILE * )                                                             // 
	off_t  ftello(FILE * )                                                           // 
	int  ftrylockfile(FILE * )                                                       // 
	void  funlockfile(FILE * )                                                       // 
	FILE *  funopen(const void * , int (*)(void *, char *, int) , int (*)(void *, const char *, int) , fpos_t (*)(void *, fpos_t, int) , int (*)(void *) ) // 
	size_t  fwrite(const void * , size_t , size_t , FILE * )                         // 
	int  getc(FILE * )                                                               // 
	int  getc_unlocked(FILE * )                                                      // 
	int  getchar()                                                                   // 
	int  getchar_unlocked()                                                          // 
	ssize_t  getdelim(char ** , size_t * , int , FILE * )                            // 
	ssize_t  getline(char ** , size_t * , FILE * )                                   // 
	char *  gets(char * )                                                            // 
	int  getw(FILE * )                                                               // 
	int  main(int argc , char **argv )                                               // 
	int  pclose(FILE * )                                                             // 
	void  perror(const char * )                                                      // 
	FILE *  popen(const char * , const char * )                                      // 
	int  printf(const char *, ... )                                                  // 
	int  putc(int , FILE * )                                                         // 
	int  putc_unlocked(int , FILE * )                                                // 
	int  putchar(int )                                                               // 
	int  putchar_unlocked(int )                                                      // 
	int  puts(const char * )                                                         // 
	int  putw(int , FILE * )                                                         // 
	int  remove(const char * )                                                       // 
	int  rename(const char * , const char * )                                        // 
	void  rewind(FILE * )                                                            // 
	int  scanf(const char *, ... )                                                   // 
	void  setbuf(FILE * , char * )                                                   // 
	void  setbuffer(FILE * , char * , int )                                          // 
	int  setlinebuf(FILE * )                                                         // 
	int  setvbuf(FILE * , char * , int , size_t )                                    // 
	int  snprintf(char * , size_t , const char *, ... )                              // 
	int  sprintf(char * , const char *, ... )                                        // 
	int  sscanf(const char * , const char *, ... )                                   // 
	char *  tempnam(const char * , const char * )                                    // 
	FILE *  tmpfile()                                                                // 
	char *  tmpnam(char * )                                                          // 
	int  ungetc(int , FILE * )                                                       // 
	int  vasprintf(char ** , const char * , __va_list_tag * )                        // 
	int  vdprintf(int , const char * , __va_list_tag * )                             // 
	int  vfprintf(FILE * , const char * , __va_list_tag * )                          // 
	int  vfscanf(FILE * , const char * , __va_list_tag * )                           // 
	int  vprintf(const char * , __va_list_tag * )                                    // 
	int  vscanf(const char * , __va_list_tag * )                                     // 
	int  vsnprintf(char * , size_t , const char * , __va_list_tag * )                // 
	int  vsprintf(char * , const char * , __va_list_tag * )                          // 
	int  vsscanf(const char * , const char * , __va_list_tag * )                     // 
	FILE *  zopen(const char * , const char * , int )                                // 
