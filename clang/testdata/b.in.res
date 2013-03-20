Fields:
	std::float_denorm_style denorm_absent                                            // 
	std::float_denorm_style denorm_indeterminate                                     // 
	std::float_denorm_style denorm_present                                           // 
	const std::nothrow_t nothrow                                                     // 
	std::float_round_style round_indeterminate                                       // 
	std::float_round_style round_to_nearest                                          // 
	std::float_round_style round_toward_infinity                                     // 
	std::float_round_style round_toward_neg_infinity                                 // 
	std::float_round_style round_toward_zero                                         // 
Methods:
	void  abort()                                                                    // 
	int  abs(int )                                                                   // 
	long  abs(long __i )                                                             // 
	long long  abs(long long __x )                                                   // 
	_ForwardIterator  adjacent_find(_ForwardIterator __first , _ForwardIterator __last ) // 
	_ForwardIterator  adjacent_find(_ForwardIterator __first , _ForwardIterator __last , _BinaryPredicate __binary_pred ) // 
	void  advance(_InputIterator &__i , _Distance __n )                              // 
	char *  asctime(const struct tm * )                                              // 
	int  atexit(void (*)() )                                                         // 
	double  atof(const char * )                                                      // 
	int  atoi(const char * )                                                         // 
	long  atol(const char * )                                                        // 
	long long  atoll(const char * )                                                  // 
	back_insert_iterator<_Container>  back_inserter(_Container &__x )                // 
	bool  binary_search(_ForwardIterator __first , _ForwardIterator __last , const _Tp &__val ) // 
	bool  binary_search(_ForwardIterator __first , _ForwardIterator __last , const _Tp &__val , _Compare __comp ) // 
	binder1st<_Operation>  bind1st(const _Operation &__fn , const _Tp &__x )         // 
	binder2nd<_Operation>  bind2nd(const _Operation &__fn , const _Tp &__x )         // 
	void *  bsearch(const void * , const void * , size_t , size_t , int (*)(const void *, const void *) ) // 
	wint_t  btowc(int )                                                              // 
	void *  calloc(size_t , size_t )                                                 // 
	void  clearerr(FILE * )                                                          // 
	clock_t  clock()                                                                 // 
	_OutputIterator  copy(_InputIterator __first , _InputIterator __last , _OutputIterator __result ) // 
	typename __gnu_cxx::__enable_if<__is_char<_CharT>::__value, ostreambuf_iterator<_CharT> >::__type  copy(istreambuf_iterator<_CharT> , istreambuf_iterator<_CharT> , ostreambuf_iterator<_CharT> ) // 
	_BI2  copy_backward(_BI1 __first , _BI1 __last , _BI2 __result )                 // 
	typename iterator_traits<_InputIterator>::difference_type  count(_InputIterator __first , _InputIterator __last , const _Tp &__value ) // 
	typename iterator_traits<_InputIterator>::difference_type  count_if(_InputIterator __first , _InputIterator __last , _Predicate __pred ) // 
	char *  ctime(const time_t * )                                                   // 
	double  difftime(time_t , time_t )                                               // 
	typename iterator_traits<_InputIterator>::difference_type  distance(_InputIterator __first , _InputIterator __last ) // 
	div_t  div(int , int )                                                           // 
	ldiv_t  div(long __i , long __j )                                                // 
	lldiv_t  div(long long __n , long long __d )                                     // 
	bool  equal(_InputIterator1 __first1 , _InputIterator1 __last1 , _InputIterator2 __first2 ) // 
	bool  equal(_InputIterator1 __first1 , _InputIterator1 __last1 , _InputIterator2 __first2 , _BinaryPredicate __binary_pred ) // 
	pair<_ForwardIterator, _ForwardIterator>  equal_range(_ForwardIterator __first , _ForwardIterator __last , const _Tp &__val ) // 
	pair<_ForwardIterator, _ForwardIterator>  equal_range(_ForwardIterator __first , _ForwardIterator __last , const _Tp &__val , _Compare __comp ) // 
	void  exit(int )                                                                 // 
	int  fclose(FILE * )                                                             // 
	int  feof(FILE * )                                                               // 
	int  ferror(FILE * )                                                             // 
	int  fflush(FILE * )                                                             // 
	int  fgetc(FILE * )                                                              // 
	int  fgetpos(FILE * , fpos_t * )                                                 // 
	char *  fgets(char * , int , FILE * )                                            // 
	wint_t  fgetwc(FILE * )                                                          // 
	wchar_t *  fgetws(wchar_t * , int , FILE * )                                     // 
	void  fill(_ForwardIterator __first , _ForwardIterator __last , const _Tp &__value ) // 
	void  fill(unsigned char *__first , unsigned char *__last , const unsigned char &__c ) // 
	void  fill(signed char *__first , signed char *__last , const signed char &__c ) // 
	void  fill(char *__first , char *__last , const char &__c )                      // 
	void  fill(std::_Bit_iterator __first , std::_Bit_iterator __last , const bool &__x ) // 
	_OutputIterator  fill_n(_OutputIterator __first , _Size __n , const _Tp &__value ) // 
	unsigned char *  fill_n(unsigned char *__first , _Size __n , const unsigned char &__c ) // 
	signed char *  fill_n(signed char *__first , _Size __n , const signed char &__c ) // 
	char *  fill_n(char *__first , _Size __n , const char &__c )                     // 
	typename __gnu_cxx::__enable_if<__is_char<_CharT>::__value, istreambuf_iterator<_CharT> >::__type  find(istreambuf_iterator<_CharT> , istreambuf_iterator<_CharT> , const _CharT & ) // 
	_InputIterator  find(_InputIterator __first , _InputIterator __last , const _Tp &__val ) // 
	_ForwardIterator1  find_end(_ForwardIterator1 __first1 , _ForwardIterator1 __last1 , _ForwardIterator2 __first2 , _ForwardIterator2 __last2 ) // 
	_ForwardIterator1  find_end(_ForwardIterator1 __first1 , _ForwardIterator1 __last1 , _ForwardIterator2 __first2 , _ForwardIterator2 __last2 , _BinaryPredicate __comp ) // 
	_InputIterator  find_first_of(_InputIterator __first1 , _InputIterator __last1 , _ForwardIterator __first2 , _ForwardIterator __last2 ) // 
	_InputIterator  find_first_of(_InputIterator __first1 , _InputIterator __last1 , _ForwardIterator __first2 , _ForwardIterator __last2 , _BinaryPredicate __comp ) // 
	_InputIterator  find_if(_InputIterator __first , _InputIterator __last , _Predicate __pred ) // 
	FILE *  fopen(const char * , const char * )                                      // 
	_Function  for_each(_InputIterator __first , _InputIterator __last , _Function __f ) // 
	int  fprintf(FILE * , const char *, ... )                                        // 
	int  fputc(int , FILE * )                                                        // 
	int  fputs(const char * , FILE * )                                               // 
	wint_t  fputwc(wchar_t , FILE * )                                                // 
	int  fputws(const wchar_t * , FILE * )                                           // 
	size_t  fread(void * , size_t , size_t , FILE * )                                // 
	void  free(void * )                                                              // 
	FILE *  freopen(const char * , const char * , FILE * )                           // 
	front_insert_iterator<_Container>  front_inserter(_Container &__x )              // 
	int  fscanf(FILE * , const char *, ... )                                         // 
	int  fseek(FILE * , long , int )                                                 // 
	int  fsetpos(FILE * , const fpos_t * )                                           // 
	long  ftell(FILE * )                                                             // 
	int  fwide(FILE * , int )                                                        // 
	int  fwprintf(FILE * , const wchar_t *, ... )                                    // 
	size_t  fwrite(const void * , size_t , size_t , FILE * )                         // 
	int  fwscanf(FILE * , const wchar_t *, ... )                                     // 
	void  generate(_ForwardIterator __first , _ForwardIterator __last , _Generator __gen ) // 
	_OutputIterator  generate_n(_OutputIterator __first , _Size __n , _Generator __gen ) // 
	pair<_Tp *, ptrdiff_t>  get_temporary_buffer<<#typename _Tp#>>(ptrdiff_t __len ) // 
	int  getc(FILE * )                                                               // 
	int  getchar()                                                                   // 
	char *  getenv(const char * )                                                    // 
	basic_istream<_CharT, _Traits> &  getline(basic_istream<_CharT, _Traits> &__is , basic_string<_CharT, _Traits, _Alloc> &__str , _CharT __delim ) // 
	basic_istream<_CharT, _Traits> &  getline(basic_istream<_CharT, _Traits> &__is , basic_string<_CharT, _Traits, _Alloc> &__str ) // 
	char *  gets(char * )                                                            // 
	wint_t  getwc(FILE * )                                                           // 
	wint_t  getwchar()                                                               // 
	struct tm *  gmtime(const time_t * )                                             // 
	bool  includes(_InputIterator1 __first1 , _InputIterator1 __last1 , _InputIterator2 __first2 , _InputIterator2 __last2 ) // 
	bool  includes(_InputIterator1 __first1 , _InputIterator1 __last1 , _InputIterator2 __first2 , _InputIterator2 __last2 , _Compare __comp ) // 
	void  inplace_merge(_BidirectionalIterator __first , _BidirectionalIterator __middle , _BidirectionalIterator __last ) // 
	void  inplace_merge(_BidirectionalIterator __first , _BidirectionalIterator __middle , _BidirectionalIterator __last , _Compare __comp ) // 
	insert_iterator<_Container>  inserter(_Container &__x , _Iterator __i )          // 
	int  isalnum(int _c )                                                            // 
	int  isalpha(int _c )                                                            // 
	int  iscntrl(int _c )                                                            // 
	int  isdigit(int _c )                                                            // 
	int  isgraph(int _c )                                                            // 
	int  islower(int _c )                                                            // 
	int  isprint(int _c )                                                            // 
	int  ispunct(int _c )                                                            // 
	int  isspace(int _c )                                                            // 
	int  isupper(int _c )                                                            // 
	int  isxdigit(int _c )                                                           // 
	void  iter_swap(_ForwardIterator1 __a , _ForwardIterator2 __b )                  // 
	long  labs(long )                                                                // 
	ldiv_t  ldiv(long , long )                                                       // 
	bool  lexicographical_compare(_InputIterator1 __first1 , _InputIterator1 __last1 , _InputIterator2 __first2 , _InputIterator2 __last2 ) // 
	bool  lexicographical_compare(_InputIterator1 __first1 , _InputIterator1 __last1 , _InputIterator2 __first2 , _InputIterator2 __last2 , _Compare __comp ) // 
	bool  lexicographical_compare(const unsigned char *__first1 , const unsigned char *__last1 , const unsigned char *__first2 , const unsigned char *__last2 ) // 
	bool  lexicographical_compare(const char *__first1 , const char *__last1 , const char *__first2 , const char *__last2 ) // 
	long long  llabs(long long )                                                     // 
	lldiv_t  lldiv(long long , long long )                                           // 
	struct lconv *  localeconv()                                                     // 
	struct tm *  localtime(const time_t * )                                          // 
	_ForwardIterator  lower_bound(_ForwardIterator __first , _ForwardIterator __last , const _Tp &__val ) // 
	_ForwardIterator  lower_bound(_ForwardIterator __first , _ForwardIterator __last , const _Tp &__val , _Compare __comp ) // 
	void  make_heap(_RandomAccessIterator __first , _RandomAccessIterator __last )   // 
	void  make_heap(_RandomAccessIterator __first , _RandomAccessIterator __last , _Compare __comp ) // 
	pair<_T1, _T2>  make_pair(_T1 __x , _T2 __y )                                    // 
	void *  malloc(size_t )                                                          // 
	const _Tp &  max(const _Tp &__a , const _Tp &__b )                               // 
	const _Tp &  max(const _Tp &__a , const _Tp &__b , _Compare __comp )             // 
	_ForwardIterator  max_element(_ForwardIterator __first , _ForwardIterator __last ) // 
	_ForwardIterator  max_element(_ForwardIterator __first , _ForwardIterator __last , _Compare __comp ) // 
	int  mblen(const char * , size_t )                                               // 
	size_t  mbrlen(const char * , size_t , mbstate_t * )                             // 
	size_t  mbrtowc(wchar_t * , const char * , size_t , mbstate_t * )                // 
	int  mbsinit(const mbstate_t * )                                                 // 
	size_t  mbsrtowcs(wchar_t * , const char ** , size_t , mbstate_t * )             // 
	size_t  mbstowcs(wchar_t * , const char * , size_t )                             // 
	int  mbtowc(wchar_t * , const char * , size_t )                                  // 
	mem_fun_t<_Ret, _Tp>  mem_fun(_Ret (_Tp::*__f)() )                               // 
	const_mem_fun_t<_Ret, _Tp>  mem_fun(_Ret (_Tp::*__f)() const )                   // 
	mem_fun1_t<_Ret, _Tp, _Arg>  mem_fun(_Ret (_Tp::*__f)(_Arg) )                    // 
	const_mem_fun1_t<_Ret, _Tp, _Arg>  mem_fun(_Ret (_Tp::*__f)(_Arg) const )        // 
	mem_fun_ref_t<_Ret, _Tp>  mem_fun_ref(_Ret (_Tp::*__f)() )                       // 
	const_mem_fun_ref_t<_Ret, _Tp>  mem_fun_ref(_Ret (_Tp::*__f)() const )           // 
	mem_fun1_ref_t<_Ret, _Tp, _Arg>  mem_fun_ref(_Ret (_Tp::*__f)(_Arg) )            // 
	const_mem_fun1_ref_t<_Ret, _Tp, _Arg>  mem_fun_ref(_Ret (_Tp::*__f)(_Arg) const ) // 
	void *  memchr(const void * , int , size_t )                                     // 
	void *  memchr(void *__p , int __c , size_t __n )                                // 
	int  memcmp(const void * , const void * , size_t )                               // 
	void *  memcpy(void * , const void * , size_t )                                  // 
	void *  memmove(void * , const void * , size_t )                                 // 
	void *  memset(void * , int , size_t )                                           // 
	_OutputIterator  merge(_InputIterator1 __first1 , _InputIterator1 __last1 , _InputIterator2 __first2 , _InputIterator2 __last2 , _OutputIterator __result ) // 
	_OutputIterator  merge(_InputIterator1 __first1 , _InputIterator1 __last1 , _InputIterator2 __first2 , _InputIterator2 __last2 , _OutputIterator __result , _Compare __comp ) // 
	const _Tp &  min(const _Tp &__a , const _Tp &__b )                               // 
	const _Tp &  min(const _Tp &__a , const _Tp &__b , _Compare __comp )             // 
	_ForwardIterator  min_element(_ForwardIterator __first , _ForwardIterator __last ) // 
	_ForwardIterator  min_element(_ForwardIterator __first , _ForwardIterator __last , _Compare __comp ) // 
	pair<_InputIterator1, _InputIterator2>  mismatch(_InputIterator1 __first1 , _InputIterator1 __last1 , _InputIterator2 __first2 ) // 
	pair<_InputIterator1, _InputIterator2>  mismatch(_InputIterator1 __first1 , _InputIterator1 __last1 , _InputIterator2 __first2 , _BinaryPredicate __binary_pred ) // 
	time_t  mktime(struct tm * )                                                     // 
	bool  next_permutation(_BidirectionalIterator __first , _BidirectionalIterator __last ) // 
	bool  next_permutation(_BidirectionalIterator __first , _BidirectionalIterator __last , _Compare __comp ) // 
	unary_negate<_Predicate>  not1(const _Predicate &__pred )                        // 
	binary_negate<_Predicate>  not2(const _Predicate &__pred )                       // 
	void  nth_element(_RandomAccessIterator __first , _RandomAccessIterator __nth , _RandomAccessIterator __last ) // 
	void  nth_element(_RandomAccessIterator __first , _RandomAccessIterator __nth , _RandomAccessIterator __last , _Compare __comp ) // 
	void  partial_sort(_RandomAccessIterator __first , _RandomAccessIterator __middle , _RandomAccessIterator __last ) // 
	void  partial_sort(_RandomAccessIterator __first , _RandomAccessIterator __middle , _RandomAccessIterator __last , _Compare __comp ) // 
	_RandomAccessIterator  partial_sort_copy(_InputIterator __first , _InputIterator __last , _RandomAccessIterator __result_first , _RandomAccessIterator __result_last ) // 
	_RandomAccessIterator  partial_sort_copy(_InputIterator __first , _InputIterator __last , _RandomAccessIterator __result_first , _RandomAccessIterator __result_last , _Compare __comp ) // 
	_ForwardIterator  partition(_ForwardIterator __first , _ForwardIterator __last , _Predicate __pred ) // 
	void  perror(const char * )                                                      // 
	void  pop_heap(_RandomAccessIterator __first , _RandomAccessIterator __last )    // 
	void  pop_heap(_RandomAccessIterator __first , _RandomAccessIterator __last , _Compare __comp ) // 
	bool  prev_permutation(_BidirectionalIterator __first , _BidirectionalIterator __last ) // 
	bool  prev_permutation(_BidirectionalIterator __first , _BidirectionalIterator __last , _Compare __comp ) // 
	int  printf(const char *, ... )                                                  // 
	pointer_to_unary_function<_Arg, _Result>  ptr_fun(_Result (*__x)(_Arg) )         // 
	pointer_to_binary_function<_Arg1, _Arg2, _Result>  ptr_fun(_Result (*__x)(_Arg1, _Arg2) ) // 
	void  push_heap(_RandomAccessIterator __first , _RandomAccessIterator __last )   // 
	void  push_heap(_RandomAccessIterator __first , _RandomAccessIterator __last , _Compare __comp ) // 
	int  putc(int , FILE * )                                                         // 
	int  putchar(int )                                                               // 
	int  puts(const char * )                                                         // 
	wint_t  putwc(wchar_t , FILE * )                                                 // 
	wint_t  putwchar(wchar_t )                                                       // 
	void  qsort(void * , size_t , size_t , int (*)(const void *, const void *) )     // 
	int  rand()                                                                      // 
	void  random_shuffle(_RandomAccessIterator __first , _RandomAccessIterator __last ) // 
	void  random_shuffle(_RandomAccessIterator __first , _RandomAccessIterator __last , _RandomNumberGenerator &__rand ) // 
	void *  realloc(void * , size_t )                                                // 
	int  remove(const char * )                                                       // 
	_ForwardIterator  remove(_ForwardIterator __first , _ForwardIterator __last , const _Tp &__value ) // 
	_OutputIterator  remove_copy(_InputIterator __first , _InputIterator __last , _OutputIterator __result , const _Tp &__value ) // 
	_OutputIterator  remove_copy_if(_InputIterator __first , _InputIterator __last , _OutputIterator __result , _Predicate __pred ) // 
	_ForwardIterator  remove_if(_ForwardIterator __first , _ForwardIterator __last , _Predicate __pred ) // 
	int  rename(const char * , const char * )                                        // 
	void  replace(_ForwardIterator __first , _ForwardIterator __last , const _Tp &__old_value , const _Tp &__new_value ) // 
	_OutputIterator  replace_copy(_InputIterator __first , _InputIterator __last , _OutputIterator __result , const _Tp &__old_value , const _Tp &__new_value ) // 
	_OutputIterator  replace_copy_if(_InputIterator __first , _InputIterator __last , _OutputIterator __result , _Predicate __pred , const _Tp &__new_value ) // 
	void  replace_if(_ForwardIterator __first , _ForwardIterator __last , _Predicate __pred , const _Tp &__new_value ) // 
	void  return_temporary_buffer(_Tp *__p )                                         // 
	void  reverse(_BidirectionalIterator __first , _BidirectionalIterator __last )   // 
	_OutputIterator  reverse_copy(_BidirectionalIterator __first , _BidirectionalIterator __last , _OutputIterator __result ) // 
	void  rewind(FILE * )                                                            // 
	void  rotate(_ForwardIterator __first , _ForwardIterator __middle , _ForwardIterator __last ) // 
	_OutputIterator  rotate_copy(_ForwardIterator __first , _ForwardIterator __middle , _ForwardIterator __last , _OutputIterator __result ) // 
	int  scanf(const char *, ... )                                                   // 
	_ForwardIterator1  search(_ForwardIterator1 __first1 , _ForwardIterator1 __last1 , _ForwardIterator2 __first2 , _ForwardIterator2 __last2 ) // 
	_ForwardIterator1  search(_ForwardIterator1 __first1 , _ForwardIterator1 __last1 , _ForwardIterator2 __first2 , _ForwardIterator2 __last2 , _BinaryPredicate __predicate ) // 
	_ForwardIterator  search_n(_ForwardIterator __first , _ForwardIterator __last , _Integer __count , const _Tp &__val ) // 
	_ForwardIterator  search_n(_ForwardIterator __first , _ForwardIterator __last , _Integer __count , const _Tp &__val , _BinaryPredicate __binary_pred ) // 
	_OutputIterator  set_difference(_InputIterator1 __first1 , _InputIterator1 __last1 , _InputIterator2 __first2 , _InputIterator2 __last2 , _OutputIterator __result ) // 
	_OutputIterator  set_difference(_InputIterator1 __first1 , _InputIterator1 __last1 , _InputIterator2 __first2 , _InputIterator2 __last2 , _OutputIterator __result , _Compare __comp ) // 
	_OutputIterator  set_intersection(_InputIterator1 __first1 , _InputIterator1 __last1 , _InputIterator2 __first2 , _InputIterator2 __last2 , _OutputIterator __result ) // 
	_OutputIterator  set_intersection(_InputIterator1 __first1 , _InputIterator1 __last1 , _InputIterator2 __first2 , _InputIterator2 __last2 , _OutputIterator __result , _Compare __comp ) // 
	new_handler  set_new_handler(new_handler )                                       // 
	_OutputIterator  set_symmetric_difference(_InputIterator1 __first1 , _InputIterator1 __last1 , _InputIterator2 __first2 , _InputIterator2 __last2 , _OutputIterator __result ) // 
	_OutputIterator  set_symmetric_difference(_InputIterator1 __first1 , _InputIterator1 __last1 , _InputIterator2 __first2 , _InputIterator2 __last2 , _OutputIterator __result , _Compare __comp ) // 
	terminate_handler  set_terminate(terminate_handler )                             // 
	unexpected_handler  set_unexpected(unexpected_handler )                          // 
	_OutputIterator  set_union(_InputIterator1 __first1 , _InputIterator1 __last1 , _InputIterator2 __first2 , _InputIterator2 __last2 , _OutputIterator __result ) // 
	_OutputIterator  set_union(_InputIterator1 __first1 , _InputIterator1 __last1 , _InputIterator2 __first2 , _InputIterator2 __last2 , _OutputIterator __result , _Compare __comp ) // 
	void  setbuf(FILE * , char * )                                                   // 
	char *  setlocale(int , const char * )                                           // 
	int  setvbuf(FILE * , char * , int , size_t )                                    // 
	int  snprintf(char * , size_t , const char *, ... )                              // 
	void  sort(_RandomAccessIterator __first , _RandomAccessIterator __last )        // 
	void  sort(_RandomAccessIterator __first , _RandomAccessIterator __last , _Compare __comp ) // 
	void  sort_heap(_RandomAccessIterator __first , _RandomAccessIterator __last )   // 
	void  sort_heap(_RandomAccessIterator __first , _RandomAccessIterator __last , _Compare __comp ) // 
	int  sprintf(char * , const char *, ... )                                        // 
	void  srand(unsigned int )                                                       // 
	int  sscanf(const char * , const char *, ... )                                   // 
	_ForwardIterator  stable_partition(_ForwardIterator __first , _ForwardIterator __last , _Predicate __pred ) // 
	void  stable_sort(_RandomAccessIterator __first , _RandomAccessIterator __last ) // 
	void  stable_sort(_RandomAccessIterator __first , _RandomAccessIterator __last , _Compare __comp ) // 
	char *  strcat(char * , const char * )                                           // 
	char *  strchr(const char * , int )                                              // 
	char *  strchr(char *__s1 , int __n )                                            // 
	int  strcmp(const char * , const char * )                                        // 
	int  strcoll(const char * , const char * )                                       // 
	char *  strcpy(char * , const char * )                                           // 
	size_t  strcspn(const char * , const char * )                                    // 
	char *  strerror(int )                                                           // 
	size_t  strftime(char * , size_t , const char * , const struct tm * )            // 
	size_t  strlen(const char * )                                                    // 
	char *  strncat(char * , const char * , size_t )                                 // 
	int  strncmp(const char * , const char * , size_t )                              // 
	char *  strncpy(char * , const char * , size_t )                                 // 
	char *  strpbrk(const char * , const char * )                                    // 
	char *  strpbrk(char *__s1 , const char *__s2 )                                  // 
	char *  strrchr(const char * , int )                                             // 
	char *  strrchr(char *__s1 , int __n )                                           // 
	size_t  strspn(const char * , const char * )                                     // 
	char *  strstr(const char * , const char * )                                     // 
	char *  strstr(char *__s1 , const char *__s2 )                                   // 
	double  strtod(const char * , char ** )                                          // 
	float  strtof(const char * , char ** )                                           // 
	char *  strtok(char * , const char * )                                           // 
	long  strtol(const char * , char ** , int )                                      // 
	long double  strtold(const char * , char ** )                                    // 
	long long  strtoll(const char * , char ** , int )                                // 
	unsigned long  strtoul(const char * , char ** , int )                            // 
	unsigned long long  strtoull(const char * , char ** , int )                      // 
	size_t  strxfrm(char * , const char * , size_t )                                 // 
	void  swap(_Tp &__a , _Tp &__b )                                                 // 
	void  swap(vector<_Tp, _Alloc> &__x , vector<_Tp, _Alloc> &__y )                 // 
	void  swap(basic_string<_CharT, _Traits, _Alloc> &__lhs , basic_string<_CharT, _Traits, _Alloc> &__rhs ) // 
	_ForwardIterator2  swap_ranges(_ForwardIterator1 __first1 , _ForwardIterator1 __last1 , _ForwardIterator2 __first2 ) // 
	int  swprintf(wchar_t * , size_t , const wchar_t *, ... )                        // 
	int  swscanf(const wchar_t * , const wchar_t *, ... )                            // 
	int  system(const char * )                                                       // 
	void  terminate()                                                                // 
	time_t  time(time_t * )                                                          // 
	FILE *  tmpfile()                                                                // 
	char *  tmpnam(char * )                                                          // 
	int  tolower(int _c )                                                            // 
	int  toupper(int _c )                                                            // 
	_OutputIterator  transform(_InputIterator __first , _InputIterator __last , _OutputIterator __result , _UnaryOperation __unary_op ) // 
	_OutputIterator  transform(_InputIterator1 __first1 , _InputIterator1 __last1 , _InputIterator2 __first2 , _OutputIterator __result , _BinaryOperation __binary_op ) // 
	bool  uncaught_exception()                                                       // 
	void  unexpected()                                                               // 
	int  ungetc(int , FILE * )                                                       // 
	wint_t  ungetwc(wint_t , FILE * )                                                // 
	_ForwardIterator  uninitialized_copy(_InputIterator __first , _InputIterator __last , _ForwardIterator __result ) // 
	char *  uninitialized_copy(const char *__first , const char *__last , char *__result ) // 
	wchar_t *  uninitialized_copy(const wchar_t *__first , const wchar_t *__last , wchar_t *__result ) // 
	void  uninitialized_fill(_ForwardIterator __first , _ForwardIterator __last , const _Tp &__x ) // 
	void  uninitialized_fill_n(_ForwardIterator __first , _Size __n , const _Tp &__x ) // 
	_ForwardIterator  unique(_ForwardIterator __first , _ForwardIterator __last )    // 
	_ForwardIterator  unique(_ForwardIterator __first , _ForwardIterator __last , _BinaryPredicate __binary_pred ) // 
	_OutputIterator  unique_copy(_InputIterator __first , _InputIterator __last , _OutputIterator __result ) // 
	_OutputIterator  unique_copy(_InputIterator __first , _InputIterator __last , _OutputIterator __result , _BinaryPredicate __binary_pred ) // 
	_ForwardIterator  upper_bound(_ForwardIterator __first , _ForwardIterator __last , const _Tp &__val ) // 
	_ForwardIterator  upper_bound(_ForwardIterator __first , _ForwardIterator __last , const _Tp &__val , _Compare __comp ) // 
	int  vfprintf(FILE * , const char * , __va_list_tag * )                          // 
	int  vfscanf(FILE * , const char * , __va_list_tag * )                           // 
	int  vfwprintf(FILE * , const wchar_t * , __va_list_tag * )                      // 
	int  vfwscanf(FILE * , const wchar_t * , __va_list_tag * )                       // 
	int  vprintf(const char * , __va_list_tag * )                                    // 
	int  vscanf(const char * , __va_list_tag * )                                     // 
	int  vsnprintf(char * , size_t , const char * , __va_list_tag * )                // 
	int  vsprintf(char * , const char * , __va_list_tag * )                          // 
	int  vsscanf(const char * , const char * , __va_list_tag * )                     // 
	int  vswprintf(wchar_t * , size_t , const wchar_t * , __va_list_tag * )          // 
	int  vswscanf(const wchar_t * , const wchar_t * , __va_list_tag * )              // 
	int  vwprintf(const wchar_t * , __va_list_tag * )                                // 
	int  vwscanf(const wchar_t * , __va_list_tag * )                                 // 
	size_t  wcrtomb(char * , wchar_t , mbstate_t * )                                 // 
	wchar_t *  wcscat(wchar_t * , const wchar_t * )                                  // 
	wchar_t *  wcschr(const wchar_t * , wchar_t )                                    // 
	wchar_t *  wcschr(wchar_t *__p , wchar_t __c )                                   // 
	int  wcscmp(const wchar_t * , const wchar_t * )                                  // 
	int  wcscoll(const wchar_t * , const wchar_t * )                                 // 
	wchar_t *  wcscpy(wchar_t * , const wchar_t * )                                  // 
	size_t  wcscspn(const wchar_t * , const wchar_t * )                              // 
	size_t  wcsftime(wchar_t * , size_t , const wchar_t * , const struct tm * )      // 
	size_t  wcslen(const wchar_t * )                                                 // 
	wchar_t *  wcsncat(wchar_t * , const wchar_t * , size_t )                        // 
	int  wcsncmp(const wchar_t * , const wchar_t * , size_t )                        // 
	wchar_t *  wcsncpy(wchar_t * , const wchar_t * , size_t )                        // 
	wchar_t *  wcspbrk(const wchar_t * , const wchar_t * )                           // 
	wchar_t *  wcspbrk(wchar_t *__s1 , const wchar_t *__s2 )                         // 
	wchar_t *  wcsrchr(const wchar_t * , wchar_t )                                   // 
	wchar_t *  wcsrchr(wchar_t *__p , wchar_t __c )                                  // 
	size_t  wcsrtombs(char * , const wchar_t ** , size_t , mbstate_t * )             // 
	size_t  wcsspn(const wchar_t * , const wchar_t * )                               // 
	wchar_t *  wcsstr(const wchar_t * , const wchar_t * )                            // 
	wchar_t *  wcsstr(wchar_t *__s1 , const wchar_t *__s2 )                          // 
	double  wcstod(const wchar_t * , wchar_t ** )                                    // 
	float  wcstof(const wchar_t * , wchar_t ** )                                     // 
	wchar_t *  wcstok(wchar_t * , const wchar_t * , wchar_t ** )                     // 
	long  wcstol(const wchar_t * , wchar_t ** , int )                                // 
	long double  wcstold(const wchar_t * , wchar_t ** )                              // 
	long long  wcstoll(const wchar_t * , wchar_t ** , int )                          // 
	size_t  wcstombs(char * , const wchar_t * , size_t )                             // 
	unsigned long  wcstoul(const wchar_t * , wchar_t ** , int )                      // 
	unsigned long long  wcstoull(const wchar_t * , wchar_t ** , int )                // 
	size_t  wcsxfrm(wchar_t * , const wchar_t * , size_t )                           // 
	int  wctob(wint_t )                                                              // 
	int  wctomb(char * , wchar_t )                                                   // 
	wchar_t *  wmemchr(const wchar_t * , wchar_t , size_t )                          // 
	wchar_t *  wmemchr(wchar_t *__p , wchar_t __c , size_t __n )                     // 
	int  wmemcmp(const wchar_t * , const wchar_t * , size_t )                        // 
	wchar_t *  wmemcpy(wchar_t * , const wchar_t * , size_t )                        // 
	wchar_t *  wmemmove(wchar_t * , const wchar_t * , size_t )                       // 
	wchar_t *  wmemset(wchar_t * , wchar_t , size_t )                                // 
	int  wprintf(const wchar_t *, ... )                                              // 
	int  wscanf(const wchar_t *, ... )                                               // 
