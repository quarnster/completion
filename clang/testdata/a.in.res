Fields:
	// {"Name":{"Relative":"argc"},"Type":{"Name":{"Relative":"int"}}}
	int argc                                                                         // 
	// {"Name":{"Relative":"argv"},"Type":{"Name":{"Relative":"char **"}}}
	char ** argv                                                                     // 
	// {"Name":{"Relative":"sys_errlist"},"Type":{"Name":{"Relative":"const char *const []"}}}
	const char *const [] sys_errlist                                                 // 
	// {"Name":{"Relative":"sys_nerr"},"Type":{"Name":{"Relative":"const int"}}}
	const int sys_nerr                                                               // 
Methods:
	// {"Name":{"Relative":"asprintf"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"int"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"char **"}}},{"Name":{},"Type":{"Name":{"Relative":"const char *, ..."}}}]}
	int  asprintf(char ** , const char *, ... )                                      // 
	// {"Name":{"Relative":"clearerr"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"void"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"FILE *"}}}]}
	void  clearerr(FILE * )                                                          // 
	// {"Name":{"Relative":"ctermid"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"char *"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"char *"}}}]}
	char *  ctermid(char * )                                                         // 
	// {"Name":{"Relative":"ctermid_r"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"char *"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"char *"}}}]}
	char *  ctermid_r(char * )                                                       // 
	// {"Name":{"Relative":"dprintf"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"int"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"int"}}},{"Name":{},"Type":{"Name":{"Relative":"const char *, ..."}}}]}
	int  dprintf(int , const char *, ... )                                           // 
	// {"Name":{"Relative":"fclose"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"int"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"FILE *"}}}]}
	int  fclose(FILE * )                                                             // 
	// {"Name":{"Relative":"fdopen"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"FILE *"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"int"}}},{"Name":{},"Type":{"Name":{"Relative":"const char *"}}}]}
	FILE *  fdopen(int , const char * )                                              // 
	// {"Name":{"Relative":"feof"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"int"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"FILE *"}}}]}
	int  feof(FILE * )                                                               // 
	// {"Name":{"Relative":"ferror"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"int"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"FILE *"}}}]}
	int  ferror(FILE * )                                                             // 
	// {"Name":{"Relative":"fflush"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"int"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"FILE *"}}}]}
	int  fflush(FILE * )                                                             // 
	// {"Name":{"Relative":"fgetc"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"int"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"FILE *"}}}]}
	int  fgetc(FILE * )                                                              // 
	// {"Name":{"Relative":"fgetln"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"char *"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"FILE *"}}},{"Name":{},"Type":{"Name":{"Relative":"size_t *"}}}]}
	char *  fgetln(FILE * , size_t * )                                               // 
	// {"Name":{"Relative":"fgetpos"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"int"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"FILE *"}}},{"Name":{},"Type":{"Name":{"Relative":"fpos_t *"}}}]}
	int  fgetpos(FILE * , fpos_t * )                                                 // 
	// {"Name":{"Relative":"fgets"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"char *"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"char *"}}},{"Name":{},"Type":{"Name":{"Relative":"int"}}},{"Name":{},"Type":{"Name":{"Relative":"FILE *"}}}]}
	char *  fgets(char * , int , FILE * )                                            // 
	// {"Name":{"Relative":"fileno"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"int"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"FILE *"}}}]}
	int  fileno(FILE * )                                                             // 
	// {"Name":{"Relative":"flockfile"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"void"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"FILE *"}}}]}
	void  flockfile(FILE * )                                                         // 
	// {"Name":{"Relative":"fmtcheck"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"const char *"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"const char *"}}},{"Name":{},"Type":{"Name":{"Relative":"const char *"}}}]}
	const char *  fmtcheck(const char * , const char * )                             // 
	// {"Name":{"Relative":"fopen"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"FILE *"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"const char *"}}},{"Name":{},"Type":{"Name":{"Relative":"const char *"}}}]}
	FILE *  fopen(const char * , const char * )                                      // 
	// {"Name":{"Relative":"fprintf"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"int"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"FILE *"}}},{"Name":{},"Type":{"Name":{"Relative":"const char *, ..."}}}]}
	int  fprintf(FILE * , const char *, ... )                                        // 
	// {"Name":{"Relative":"fpurge"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"int"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"FILE *"}}}]}
	int  fpurge(FILE * )                                                             // 
	// {"Name":{"Relative":"fputc"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"int"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"int"}}},{"Name":{},"Type":{"Name":{"Relative":"FILE *"}}}]}
	int  fputc(int , FILE * )                                                        // 
	// {"Name":{"Relative":"fputs"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"int"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"const char *"}}},{"Name":{},"Type":{"Name":{"Relative":"FILE *"}}}]}
	int  fputs(const char * , FILE * )                                               // 
	// {"Name":{"Relative":"fread"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"size_t"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"void *"}}},{"Name":{},"Type":{"Name":{"Relative":"size_t"}}},{"Name":{},"Type":{"Name":{"Relative":"size_t"}}},{"Name":{},"Type":{"Name":{"Relative":"FILE *"}}}]}
	size_t  fread(void * , size_t , size_t , FILE * )                                // 
	// {"Name":{"Relative":"freopen"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"FILE *"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"const char *"}}},{"Name":{},"Type":{"Name":{"Relative":"const char *"}}},{"Name":{},"Type":{"Name":{"Relative":"FILE *"}}}]}
	FILE *  freopen(const char * , const char * , FILE * )                           // 
	// {"Name":{"Relative":"fscanf"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"int"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"FILE *"}}},{"Name":{},"Type":{"Name":{"Relative":"const char *, ..."}}}]}
	int  fscanf(FILE * , const char *, ... )                                         // 
	// {"Name":{"Relative":"fseek"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"int"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"FILE *"}}},{"Name":{},"Type":{"Name":{"Relative":"long"}}},{"Name":{},"Type":{"Name":{"Relative":"int"}}}]}
	int  fseek(FILE * , long , int )                                                 // 
	// {"Name":{"Relative":"fseeko"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"int"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"FILE *"}}},{"Name":{},"Type":{"Name":{"Relative":"off_t"}}},{"Name":{},"Type":{"Name":{"Relative":"int"}}}]}
	int  fseeko(FILE * , off_t , int )                                               // 
	// {"Name":{"Relative":"fsetpos"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"int"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"FILE *"}}},{"Name":{},"Type":{"Name":{"Relative":"const fpos_t *"}}}]}
	int  fsetpos(FILE * , const fpos_t * )                                           // 
	// {"Name":{"Relative":"ftell"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"long"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"FILE *"}}}]}
	long  ftell(FILE * )                                                             // 
	// {"Name":{"Relative":"ftello"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"off_t"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"FILE *"}}}]}
	off_t  ftello(FILE * )                                                           // 
	// {"Name":{"Relative":"ftrylockfile"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"int"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"FILE *"}}}]}
	int  ftrylockfile(FILE * )                                                       // 
	// {"Name":{"Relative":"funlockfile"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"void"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"FILE *"}}}]}
	void  funlockfile(FILE * )                                                       // 
	// {"Name":{"Relative":"funopen"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"FILE *"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"const void *"}}},{"Name":{},"Type":{"Name":{"Relative":"int (*)(void *, char *, int)"}}},{"Name":{},"Type":{"Name":{"Relative":"int (*)(void *, const char *, int)"}}},{"Name":{},"Type":{"Name":{"Relative":"fpos_t (*)(void *, fpos_t, int)"}}},{"Name":{},"Type":{"Name":{"Relative":"int (*)(void *)"}}}]}
	FILE *  funopen(const void * , int (*)(void *, char *, int) , int (*)(void *, const char *, int) , fpos_t (*)(void *, fpos_t, int) , int (*)(void *) ) // 
	// {"Name":{"Relative":"fwrite"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"size_t"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"const void *"}}},{"Name":{},"Type":{"Name":{"Relative":"size_t"}}},{"Name":{},"Type":{"Name":{"Relative":"size_t"}}},{"Name":{},"Type":{"Name":{"Relative":"FILE *"}}}]}
	size_t  fwrite(const void * , size_t , size_t , FILE * )                         // 
	// {"Name":{"Relative":"getc"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"int"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"FILE *"}}}]}
	int  getc(FILE * )                                                               // 
	// {"Name":{"Relative":"getc_unlocked"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"int"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"FILE *"}}}]}
	int  getc_unlocked(FILE * )                                                      // 
	// {"Name":{"Relative":"getchar"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"int"}}}]}
	int  getchar()                                                                   // 
	// {"Name":{"Relative":"getchar_unlocked"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"int"}}}]}
	int  getchar_unlocked()                                                          // 
	// {"Name":{"Relative":"getdelim"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"ssize_t"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"char **"}}},{"Name":{},"Type":{"Name":{"Relative":"size_t *"}}},{"Name":{},"Type":{"Name":{"Relative":"int"}}},{"Name":{},"Type":{"Name":{"Relative":"FILE *"}}}]}
	ssize_t  getdelim(char ** , size_t * , int , FILE * )                            // 
	// {"Name":{"Relative":"getline"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"ssize_t"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"char **"}}},{"Name":{},"Type":{"Name":{"Relative":"size_t *"}}},{"Name":{},"Type":{"Name":{"Relative":"FILE *"}}}]}
	ssize_t  getline(char ** , size_t * , FILE * )                                   // 
	// {"Name":{"Relative":"gets"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"char *"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"char *"}}}]}
	char *  gets(char * )                                                            // 
	// {"Name":{"Relative":"getw"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"int"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"FILE *"}}}]}
	int  getw(FILE * )                                                               // 
	// {"Name":{"Relative":"main"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"int"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"int argc"}}},{"Name":{},"Type":{"Name":{"Relative":"char **argv"}}}]}
	int  main(int argc , char **argv )                                               // 
	// {"Name":{"Relative":"pclose"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"int"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"FILE *"}}}]}
	int  pclose(FILE * )                                                             // 
	// {"Name":{"Relative":"perror"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"void"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"const char *"}}}]}
	void  perror(const char * )                                                      // 
	// {"Name":{"Relative":"popen"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"FILE *"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"const char *"}}},{"Name":{},"Type":{"Name":{"Relative":"const char *"}}}]}
	FILE *  popen(const char * , const char * )                                      // 
	// {"Name":{"Relative":"printf"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"int"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"const char *, ..."}}}]}
	int  printf(const char *, ... )                                                  // 
	// {"Name":{"Relative":"putc"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"int"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"int"}}},{"Name":{},"Type":{"Name":{"Relative":"FILE *"}}}]}
	int  putc(int , FILE * )                                                         // 
	// {"Name":{"Relative":"putc_unlocked"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"int"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"int"}}},{"Name":{},"Type":{"Name":{"Relative":"FILE *"}}}]}
	int  putc_unlocked(int , FILE * )                                                // 
	// {"Name":{"Relative":"putchar"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"int"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"int"}}}]}
	int  putchar(int )                                                               // 
	// {"Name":{"Relative":"putchar_unlocked"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"int"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"int"}}}]}
	int  putchar_unlocked(int )                                                      // 
	// {"Name":{"Relative":"puts"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"int"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"const char *"}}}]}
	int  puts(const char * )                                                         // 
	// {"Name":{"Relative":"putw"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"int"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"int"}}},{"Name":{},"Type":{"Name":{"Relative":"FILE *"}}}]}
	int  putw(int , FILE * )                                                         // 
	// {"Name":{"Relative":"remove"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"int"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"const char *"}}}]}
	int  remove(const char * )                                                       // 
	// {"Name":{"Relative":"rename"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"int"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"const char *"}}},{"Name":{},"Type":{"Name":{"Relative":"const char *"}}}]}
	int  rename(const char * , const char * )                                        // 
	// {"Name":{"Relative":"rewind"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"void"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"FILE *"}}}]}
	void  rewind(FILE * )                                                            // 
	// {"Name":{"Relative":"scanf"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"int"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"const char *, ..."}}}]}
	int  scanf(const char *, ... )                                                   // 
	// {"Name":{"Relative":"setbuf"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"void"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"FILE *"}}},{"Name":{},"Type":{"Name":{"Relative":"char *"}}}]}
	void  setbuf(FILE * , char * )                                                   // 
	// {"Name":{"Relative":"setbuffer"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"void"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"FILE *"}}},{"Name":{},"Type":{"Name":{"Relative":"char *"}}},{"Name":{},"Type":{"Name":{"Relative":"int"}}}]}
	void  setbuffer(FILE * , char * , int )                                          // 
	// {"Name":{"Relative":"setlinebuf"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"int"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"FILE *"}}}]}
	int  setlinebuf(FILE * )                                                         // 
	// {"Name":{"Relative":"setvbuf"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"int"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"FILE *"}}},{"Name":{},"Type":{"Name":{"Relative":"char *"}}},{"Name":{},"Type":{"Name":{"Relative":"int"}}},{"Name":{},"Type":{"Name":{"Relative":"size_t"}}}]}
	int  setvbuf(FILE * , char * , int , size_t )                                    // 
	// {"Name":{"Relative":"snprintf"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"int"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"char *"}}},{"Name":{},"Type":{"Name":{"Relative":"size_t"}}},{"Name":{},"Type":{"Name":{"Relative":"const char *, ..."}}}]}
	int  snprintf(char * , size_t , const char *, ... )                              // 
	// {"Name":{"Relative":"sprintf"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"int"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"char *"}}},{"Name":{},"Type":{"Name":{"Relative":"const char *, ..."}}}]}
	int  sprintf(char * , const char *, ... )                                        // 
	// {"Name":{"Relative":"sscanf"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"int"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"const char *"}}},{"Name":{},"Type":{"Name":{"Relative":"const char *, ..."}}}]}
	int  sscanf(const char * , const char *, ... )                                   // 
	// {"Name":{"Relative":"tempnam"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"char *"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"const char *"}}},{"Name":{},"Type":{"Name":{"Relative":"const char *"}}}]}
	char *  tempnam(const char * , const char * )                                    // 
	// {"Name":{"Relative":"tmpfile"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"FILE *"}}}]}
	FILE *  tmpfile()                                                                // 
	// {"Name":{"Relative":"tmpnam"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"char *"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"char *"}}}]}
	char *  tmpnam(char * )                                                          // 
	// {"Name":{"Relative":"ungetc"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"int"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"int"}}},{"Name":{},"Type":{"Name":{"Relative":"FILE *"}}}]}
	int  ungetc(int , FILE * )                                                       // 
	// {"Name":{"Relative":"vasprintf"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"int"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"char **"}}},{"Name":{},"Type":{"Name":{"Relative":"const char *"}}},{"Name":{},"Type":{"Name":{"Relative":"__va_list_tag *"}}}]}
	int  vasprintf(char ** , const char * , __va_list_tag * )                        // 
	// {"Name":{"Relative":"vdprintf"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"int"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"int"}}},{"Name":{},"Type":{"Name":{"Relative":"const char *"}}},{"Name":{},"Type":{"Name":{"Relative":"__va_list_tag *"}}}]}
	int  vdprintf(int , const char * , __va_list_tag * )                             // 
	// {"Name":{"Relative":"vfprintf"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"int"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"FILE *"}}},{"Name":{},"Type":{"Name":{"Relative":"const char *"}}},{"Name":{},"Type":{"Name":{"Relative":"__va_list_tag *"}}}]}
	int  vfprintf(FILE * , const char * , __va_list_tag * )                          // 
	// {"Name":{"Relative":"vfscanf"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"int"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"FILE *"}}},{"Name":{},"Type":{"Name":{"Relative":"const char *"}}},{"Name":{},"Type":{"Name":{"Relative":"__va_list_tag *"}}}]}
	int  vfscanf(FILE * , const char * , __va_list_tag * )                           // 
	// {"Name":{"Relative":"vprintf"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"int"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"const char *"}}},{"Name":{},"Type":{"Name":{"Relative":"__va_list_tag *"}}}]}
	int  vprintf(const char * , __va_list_tag * )                                    // 
	// {"Name":{"Relative":"vscanf"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"int"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"const char *"}}},{"Name":{},"Type":{"Name":{"Relative":"__va_list_tag *"}}}]}
	int  vscanf(const char * , __va_list_tag * )                                     // 
	// {"Name":{"Relative":"vsnprintf"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"int"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"char *"}}},{"Name":{},"Type":{"Name":{"Relative":"size_t"}}},{"Name":{},"Type":{"Name":{"Relative":"const char *"}}},{"Name":{},"Type":{"Name":{"Relative":"__va_list_tag *"}}}]}
	int  vsnprintf(char * , size_t , const char * , __va_list_tag * )                // 
	// {"Name":{"Relative":"vsprintf"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"int"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"char *"}}},{"Name":{},"Type":{"Name":{"Relative":"const char *"}}},{"Name":{},"Type":{"Name":{"Relative":"__va_list_tag *"}}}]}
	int  vsprintf(char * , const char * , __va_list_tag * )                          // 
	// {"Name":{"Relative":"vsscanf"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"int"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"const char *"}}},{"Name":{},"Type":{"Name":{"Relative":"const char *"}}},{"Name":{},"Type":{"Name":{"Relative":"__va_list_tag *"}}}]}
	int  vsscanf(const char * , const char * , __va_list_tag * )                     // 
	// {"Name":{"Relative":"zopen"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"FILE *"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"const char *"}}},{"Name":{},"Type":{"Name":{"Relative":"const char *"}}},{"Name":{},"Type":{"Name":{"Relative":"int"}}}]}
	FILE *  zopen(const char * , const char * , int )                                // 
