Fields:
	// {"Name":{"Relative":"npos"},"Type":{"Name":{"Relative":"const size_type"}}}
	const size_type npos                                                             // 
Methods:
	// {"Name":{"Relative":"append"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"std::basic_string\u003cchar\u003e \u0026"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"const std::basic_string\u003cchar\u003e \u0026__str"}}}]}
	std::basic_string<char> &  append(const std::basic_string<char> &__str )         // 
	// {"Name":{"Relative":"append"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"std::basic_string\u003cchar\u003e \u0026"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"const std::basic_string\u003cchar\u003e \u0026__str"}}},{"Name":{},"Type":{"Name":{"Relative":"size_type __pos"}}},{"Name":{},"Type":{"Name":{"Relative":"size_type __n"}}}]}
	std::basic_string<char> &  append(const std::basic_string<char> &__str , size_type __pos , size_type __n ) // 
	// {"Name":{"Relative":"append"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"std::basic_string\u003cchar\u003e \u0026"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"const char *__s"}}},{"Name":{},"Type":{"Name":{"Relative":"size_type __n"}}}]}
	std::basic_string<char> &  append(const char *__s , size_type __n )              // 
	// {"Name":{"Relative":"append"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"std::basic_string\u003cchar\u003e \u0026"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"const char *__s"}}}]}
	std::basic_string<char> &  append(const char *__s )                              // 
	// {"Name":{"Relative":"append"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"std::basic_string\u003cchar\u003e \u0026"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"size_type __n"}}},{"Name":{},"Type":{"Name":{"Relative":"char __c"}}}]}
	std::basic_string<char> &  append(size_type __n , char __c )                     // 
	// {"Name":{"Relative":"append"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"std::basic_string\u003cchar\u003e \u0026"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"_InputIterator __first"}}},{"Name":{},"Type":{"Name":{"Relative":"_InputIterator __last"}}}]}
	std::basic_string<char> &  append(_InputIterator __first , _InputIterator __last ) // 
	// {"Name":{"Relative":"assign"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"std::basic_string\u003cchar\u003e \u0026"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"const std::basic_string\u003cchar\u003e \u0026__str"}}}]}
	std::basic_string<char> &  assign(const std::basic_string<char> &__str )         // 
	// {"Name":{"Relative":"assign"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"std::basic_string\u003cchar\u003e \u0026"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"const std::basic_string\u003cchar\u003e \u0026__str"}}},{"Name":{},"Type":{"Name":{"Relative":"size_type __pos"}}},{"Name":{},"Type":{"Name":{"Relative":"size_type __n"}}}]}
	std::basic_string<char> &  assign(const std::basic_string<char> &__str , size_type __pos , size_type __n ) // 
	// {"Name":{"Relative":"assign"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"std::basic_string\u003cchar\u003e \u0026"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"const char *__s"}}},{"Name":{},"Type":{"Name":{"Relative":"size_type __n"}}}]}
	std::basic_string<char> &  assign(const char *__s , size_type __n )              // 
	// {"Name":{"Relative":"assign"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"std::basic_string\u003cchar\u003e \u0026"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"const char *__s"}}}]}
	std::basic_string<char> &  assign(const char *__s )                              // 
	// {"Name":{"Relative":"assign"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"std::basic_string\u003cchar\u003e \u0026"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"size_type __n"}}},{"Name":{},"Type":{"Name":{"Relative":"char __c"}}}]}
	std::basic_string<char> &  assign(size_type __n , char __c )                     // 
	// {"Name":{"Relative":"assign"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"std::basic_string\u003cchar\u003e \u0026"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"_InputIterator __first"}}},{"Name":{},"Type":{"Name":{"Relative":"_InputIterator __last"}}}]}
	std::basic_string<char> &  assign(_InputIterator __first , _InputIterator __last ) // 
	// {"Name":{"Relative":"at"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"const_reference"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"size_type __n"}}}]}
	const_reference  at(size_type __n )                                              // 
	// {"Name":{"Relative":"at"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"reference"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"size_type __n"}}}]}
	reference  at(size_type __n )                                                    // 
	// {"Name":{"Relative":"basic_string"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"void"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"_InputIterator __beg"}}},{"Name":{},"Type":{"Name":{"Relative":"_InputIterator __end"}}},{"Name":{},"Type":{"Name":{"Relative":"const std::allocator\u003cchar\u003e \u0026__a"}}}]}
	void  basic_string(_InputIterator __beg , _InputIterator __end , const std::allocator<char> &__a ) // 
	// {"Name":{"Relative":"begin"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"iterator"}}}]}
	iterator  begin()                                                                // 
	// {"Name":{"Relative":"begin"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"const_iterator"}}}]}
	const_iterator  begin()                                                          // 
	// {"Name":{"Relative":"c_str"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"const char *"}}}]}
	const char *  c_str()                                                            // 
	// {"Name":{"Relative":"capacity"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"size_type"}}}]}
	size_type  capacity()                                                            // 
	// {"Name":{"Relative":"clear"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"void"}}}]}
	void  clear()                                                                    // 
	// {"Name":{"Relative":"compare"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"int"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"const std::basic_string\u003cchar\u003e \u0026__str"}}}]}
	int  compare(const std::basic_string<char> &__str )                              // 
	// {"Name":{"Relative":"compare"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"int"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"size_type __pos"}}},{"Name":{},"Type":{"Name":{"Relative":"size_type __n"}}},{"Name":{},"Type":{"Name":{"Relative":"const std::basic_string\u003cchar\u003e \u0026__str"}}}]}
	int  compare(size_type __pos , size_type __n , const std::basic_string<char> &__str ) // 
	// {"Name":{"Relative":"compare"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"int"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"size_type __pos1"}}},{"Name":{},"Type":{"Name":{"Relative":"size_type __n1"}}},{"Name":{},"Type":{"Name":{"Relative":"const std::basic_string\u003cchar\u003e \u0026__str"}}},{"Name":{},"Type":{"Name":{"Relative":"size_type __pos2"}}},{"Name":{},"Type":{"Name":{"Relative":"size_type __n2"}}}]}
	int  compare(size_type __pos1 , size_type __n1 , const std::basic_string<char> &__str , size_type __pos2 , size_type __n2 ) // 
	// {"Name":{"Relative":"compare"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"int"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"const char *__s"}}}]}
	int  compare(const char *__s )                                                   // 
	// {"Name":{"Relative":"compare"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"int"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"size_type __pos"}}},{"Name":{},"Type":{"Name":{"Relative":"size_type __n1"}}},{"Name":{},"Type":{"Name":{"Relative":"const char *__s"}}}]}
	int  compare(size_type __pos , size_type __n1 , const char *__s )                // 
	// {"Name":{"Relative":"compare"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"int"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"size_type __pos"}}},{"Name":{},"Type":{"Name":{"Relative":"size_type __n1"}}},{"Name":{},"Type":{"Name":{"Relative":"const char *__s"}}},{"Name":{},"Type":{"Name":{"Relative":"size_type __n2"}}}]}
	int  compare(size_type __pos , size_type __n1 , const char *__s , size_type __n2 ) // 
	// {"Name":{"Relative":"copy"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"size_type"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"char *__s"}}},{"Name":{},"Type":{"Name":{"Relative":"size_type __n"}}},{"Name":{},"Type":{"Name":{"Relative":"size_type __pos"}}}]}
	size_type  copy(char *__s , size_type __n , size_type __pos )                    // 
	// {"Name":{"Relative":"data"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"const char *"}}}]}
	const char *  data()                                                             // 
	// {"Name":{"Relative":"empty"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"bool"}}}]}
	bool  empty()                                                                    // 
	// {"Name":{"Relative":"end"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"iterator"}}}]}
	iterator  end()                                                                  // 
	// {"Name":{"Relative":"end"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"const_iterator"}}}]}
	const_iterator  end()                                                            // 
	// {"Name":{"Relative":"erase"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"std::basic_string\u003cchar\u003e \u0026"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"size_type __pos"}}}]}
	std::basic_string<char> &  erase(size_type __pos )                               // 
	// {"Name":{"Relative":"erase"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"iterator"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"iterator __position"}}}]}
	iterator  erase(iterator __position )                                            // 
	// {"Name":{"Relative":"erase"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"iterator"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"iterator __first"}}},{"Name":{},"Type":{"Name":{"Relative":"iterator __last"}}}]}
	iterator  erase(iterator __first , iterator __last )                             // 
	// {"Name":{"Relative":"find"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"size_type"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"const char *__s"}}},{"Name":{},"Type":{"Name":{"Relative":"size_type __pos"}}},{"Name":{},"Type":{"Name":{"Relative":"size_type __n"}}}]}
	size_type  find(const char *__s , size_type __pos , size_type __n )              // 
	// {"Name":{"Relative":"find"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"size_type"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"const std::basic_string\u003cchar\u003e \u0026__str"}}},{"Name":{},"Type":{"Name":{"Relative":"size_type __pos"}}}]}
	size_type  find(const std::basic_string<char> &__str , size_type __pos )         // 
	// {"Name":{"Relative":"find"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"size_type"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"const char *__s"}}},{"Name":{},"Type":{"Name":{"Relative":"size_type __pos"}}}]}
	size_type  find(const char *__s , size_type __pos )                              // 
	// {"Name":{"Relative":"find"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"size_type"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"char __c"}}},{"Name":{},"Type":{"Name":{"Relative":"size_type __pos"}}}]}
	size_type  find(char __c , size_type __pos )                                     // 
	// {"Name":{"Relative":"find_first_not_of"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"size_type"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"const std::basic_string\u003cchar\u003e \u0026__str"}}},{"Name":{},"Type":{"Name":{"Relative":"size_type __pos"}}}]}
	size_type  find_first_not_of(const std::basic_string<char> &__str , size_type __pos ) // 
	// {"Name":{"Relative":"find_first_not_of"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"size_type"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"const char *__s"}}},{"Name":{},"Type":{"Name":{"Relative":"size_type __pos"}}},{"Name":{},"Type":{"Name":{"Relative":"size_type __n"}}}]}
	size_type  find_first_not_of(const char *__s , size_type __pos , size_type __n ) // 
	// {"Name":{"Relative":"find_first_not_of"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"size_type"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"const char *__s"}}},{"Name":{},"Type":{"Name":{"Relative":"size_type __pos"}}}]}
	size_type  find_first_not_of(const char *__s , size_type __pos )                 // 
	// {"Name":{"Relative":"find_first_not_of"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"size_type"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"char __c"}}},{"Name":{},"Type":{"Name":{"Relative":"size_type __pos"}}}]}
	size_type  find_first_not_of(char __c , size_type __pos )                        // 
	// {"Name":{"Relative":"find_first_of"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"size_type"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"const std::basic_string\u003cchar\u003e \u0026__str"}}},{"Name":{},"Type":{"Name":{"Relative":"size_type __pos"}}}]}
	size_type  find_first_of(const std::basic_string<char> &__str , size_type __pos ) // 
	// {"Name":{"Relative":"find_first_of"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"size_type"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"const char *__s"}}},{"Name":{},"Type":{"Name":{"Relative":"size_type __pos"}}},{"Name":{},"Type":{"Name":{"Relative":"size_type __n"}}}]}
	size_type  find_first_of(const char *__s , size_type __pos , size_type __n )     // 
	// {"Name":{"Relative":"find_first_of"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"size_type"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"const char *__s"}}},{"Name":{},"Type":{"Name":{"Relative":"size_type __pos"}}}]}
	size_type  find_first_of(const char *__s , size_type __pos )                     // 
	// {"Name":{"Relative":"find_first_of"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"size_type"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"char __c"}}},{"Name":{},"Type":{"Name":{"Relative":"size_type __pos"}}}]}
	size_type  find_first_of(char __c , size_type __pos )                            // 
	// {"Name":{"Relative":"find_last_not_of"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"size_type"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"const std::basic_string\u003cchar\u003e \u0026__str"}}},{"Name":{},"Type":{"Name":{"Relative":"size_type __pos"}}}]}
	size_type  find_last_not_of(const std::basic_string<char> &__str , size_type __pos ) // 
	// {"Name":{"Relative":"find_last_not_of"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"size_type"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"const char *__s"}}},{"Name":{},"Type":{"Name":{"Relative":"size_type __pos"}}},{"Name":{},"Type":{"Name":{"Relative":"size_type __n"}}}]}
	size_type  find_last_not_of(const char *__s , size_type __pos , size_type __n )  // 
	// {"Name":{"Relative":"find_last_not_of"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"size_type"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"const char *__s"}}},{"Name":{},"Type":{"Name":{"Relative":"size_type __pos"}}}]}
	size_type  find_last_not_of(const char *__s , size_type __pos )                  // 
	// {"Name":{"Relative":"find_last_not_of"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"size_type"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"char __c"}}},{"Name":{},"Type":{"Name":{"Relative":"size_type __pos"}}}]}
	size_type  find_last_not_of(char __c , size_type __pos )                         // 
	// {"Name":{"Relative":"find_last_of"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"size_type"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"const std::basic_string\u003cchar\u003e \u0026__str"}}},{"Name":{},"Type":{"Name":{"Relative":"size_type __pos"}}}]}
	size_type  find_last_of(const std::basic_string<char> &__str , size_type __pos ) // 
	// {"Name":{"Relative":"find_last_of"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"size_type"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"const char *__s"}}},{"Name":{},"Type":{"Name":{"Relative":"size_type __pos"}}},{"Name":{},"Type":{"Name":{"Relative":"size_type __n"}}}]}
	size_type  find_last_of(const char *__s , size_type __pos , size_type __n )      // 
	// {"Name":{"Relative":"find_last_of"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"size_type"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"const char *__s"}}},{"Name":{},"Type":{"Name":{"Relative":"size_type __pos"}}}]}
	size_type  find_last_of(const char *__s , size_type __pos )                      // 
	// {"Name":{"Relative":"find_last_of"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"size_type"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"char __c"}}},{"Name":{},"Type":{"Name":{"Relative":"size_type __pos"}}}]}
	size_type  find_last_of(char __c , size_type __pos )                             // 
	// {"Name":{"Relative":"get_allocator"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"allocator_type"}}}]}
	allocator_type  get_allocator()                                                  // 
	// {"Name":{"Relative":"insert"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"void"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"iterator __p"}}},{"Name":{},"Type":{"Name":{"Relative":"size_type __n"}}},{"Name":{},"Type":{"Name":{"Relative":"char __c"}}}]}
	void  insert(iterator __p , size_type __n , char __c )                           // 
	// {"Name":{"Relative":"insert"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"void"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"iterator __p"}}},{"Name":{},"Type":{"Name":{"Relative":"_InputIterator __beg"}}},{"Name":{},"Type":{"Name":{"Relative":"_InputIterator __end"}}}]}
	void  insert(iterator __p , _InputIterator __beg , _InputIterator __end )        // 
	// {"Name":{"Relative":"insert"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"std::basic_string\u003cchar\u003e \u0026"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"size_type __pos1"}}},{"Name":{},"Type":{"Name":{"Relative":"const std::basic_string\u003cchar\u003e \u0026__str"}}}]}
	std::basic_string<char> &  insert(size_type __pos1 , const std::basic_string<char> &__str ) // 
	// {"Name":{"Relative":"insert"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"std::basic_string\u003cchar\u003e \u0026"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"size_type __pos1"}}},{"Name":{},"Type":{"Name":{"Relative":"const std::basic_string\u003cchar\u003e \u0026__str"}}},{"Name":{},"Type":{"Name":{"Relative":"size_type __pos2"}}},{"Name":{},"Type":{"Name":{"Relative":"size_type __n"}}}]}
	std::basic_string<char> &  insert(size_type __pos1 , const std::basic_string<char> &__str , size_type __pos2 , size_type __n ) // 
	// {"Name":{"Relative":"insert"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"std::basic_string\u003cchar\u003e \u0026"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"size_type __pos"}}},{"Name":{},"Type":{"Name":{"Relative":"const char *__s"}}},{"Name":{},"Type":{"Name":{"Relative":"size_type __n"}}}]}
	std::basic_string<char> &  insert(size_type __pos , const char *__s , size_type __n ) // 
	// {"Name":{"Relative":"insert"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"std::basic_string\u003cchar\u003e \u0026"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"size_type __pos"}}},{"Name":{},"Type":{"Name":{"Relative":"const char *__s"}}}]}
	std::basic_string<char> &  insert(size_type __pos , const char *__s )            // 
	// {"Name":{"Relative":"insert"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"std::basic_string\u003cchar\u003e \u0026"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"size_type __pos"}}},{"Name":{},"Type":{"Name":{"Relative":"size_type __n"}}},{"Name":{},"Type":{"Name":{"Relative":"char __c"}}}]}
	std::basic_string<char> &  insert(size_type __pos , size_type __n , char __c )   // 
	// {"Name":{"Relative":"insert"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"iterator"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"iterator __p"}}},{"Name":{},"Type":{"Name":{"Relative":"char __c"}}}]}
	iterator  insert(iterator __p , char __c )                                       // 
	// {"Name":{"Relative":"length"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"size_type"}}}]}
	size_type  length()                                                              // 
	// {"Name":{"Relative":"max_size"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"size_type"}}}]}
	size_type  max_size()                                                            // 
	// {"Name":{"Relative":"push_back"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"void"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"char __c"}}}]}
	void  push_back(char __c )                                                       // 
	// {"Name":{"Relative":"rbegin"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"reverse_iterator"}}}]}
	reverse_iterator  rbegin()                                                       // 
	// {"Name":{"Relative":"rbegin"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"const_reverse_iterator"}}}]}
	const_reverse_iterator  rbegin()                                                 // 
	// {"Name":{"Relative":"rend"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"reverse_iterator"}}}]}
	reverse_iterator  rend()                                                         // 
	// {"Name":{"Relative":"rend"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"const_reverse_iterator"}}}]}
	const_reverse_iterator  rend()                                                   // 
	// {"Name":{"Relative":"replace"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"std::basic_string\u003cchar\u003e \u0026"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"size_type __pos"}}},{"Name":{},"Type":{"Name":{"Relative":"size_type __n"}}},{"Name":{},"Type":{"Name":{"Relative":"const std::basic_string\u003cchar\u003e \u0026__str"}}}]}
	std::basic_string<char> &  replace(size_type __pos , size_type __n , const std::basic_string<char> &__str ) // 
	// {"Name":{"Relative":"replace"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"std::basic_string\u003cchar\u003e \u0026"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"size_type __pos1"}}},{"Name":{},"Type":{"Name":{"Relative":"size_type __n1"}}},{"Name":{},"Type":{"Name":{"Relative":"const std::basic_string\u003cchar\u003e \u0026__str"}}},{"Name":{},"Type":{"Name":{"Relative":"size_type __pos2"}}},{"Name":{},"Type":{"Name":{"Relative":"size_type __n2"}}}]}
	std::basic_string<char> &  replace(size_type __pos1 , size_type __n1 , const std::basic_string<char> &__str , size_type __pos2 , size_type __n2 ) // 
	// {"Name":{"Relative":"replace"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"std::basic_string\u003cchar\u003e \u0026"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"size_type __pos"}}},{"Name":{},"Type":{"Name":{"Relative":"size_type __n1"}}},{"Name":{},"Type":{"Name":{"Relative":"const char *__s"}}},{"Name":{},"Type":{"Name":{"Relative":"size_type __n2"}}}]}
	std::basic_string<char> &  replace(size_type __pos , size_type __n1 , const char *__s , size_type __n2 ) // 
	// {"Name":{"Relative":"replace"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"std::basic_string\u003cchar\u003e \u0026"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"size_type __pos"}}},{"Name":{},"Type":{"Name":{"Relative":"size_type __n1"}}},{"Name":{},"Type":{"Name":{"Relative":"const char *__s"}}}]}
	std::basic_string<char> &  replace(size_type __pos , size_type __n1 , const char *__s ) // 
	// {"Name":{"Relative":"replace"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"std::basic_string\u003cchar\u003e \u0026"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"size_type __pos"}}},{"Name":{},"Type":{"Name":{"Relative":"size_type __n1"}}},{"Name":{},"Type":{"Name":{"Relative":"size_type __n2"}}},{"Name":{},"Type":{"Name":{"Relative":"char __c"}}}]}
	std::basic_string<char> &  replace(size_type __pos , size_type __n1 , size_type __n2 , char __c ) // 
	// {"Name":{"Relative":"replace"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"std::basic_string\u003cchar\u003e \u0026"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"iterator __i1"}}},{"Name":{},"Type":{"Name":{"Relative":"iterator __i2"}}},{"Name":{},"Type":{"Name":{"Relative":"const std::basic_string\u003cchar\u003e \u0026__str"}}}]}
	std::basic_string<char> &  replace(iterator __i1 , iterator __i2 , const std::basic_string<char> &__str ) // 
	// {"Name":{"Relative":"replace"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"std::basic_string\u003cchar\u003e \u0026"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"iterator __i1"}}},{"Name":{},"Type":{"Name":{"Relative":"iterator __i2"}}},{"Name":{},"Type":{"Name":{"Relative":"const char *__s"}}},{"Name":{},"Type":{"Name":{"Relative":"size_type __n"}}}]}
	std::basic_string<char> &  replace(iterator __i1 , iterator __i2 , const char *__s , size_type __n ) // 
	// {"Name":{"Relative":"replace"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"std::basic_string\u003cchar\u003e \u0026"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"iterator __i1"}}},{"Name":{},"Type":{"Name":{"Relative":"iterator __i2"}}},{"Name":{},"Type":{"Name":{"Relative":"const char *__s"}}}]}
	std::basic_string<char> &  replace(iterator __i1 , iterator __i2 , const char *__s ) // 
	// {"Name":{"Relative":"replace"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"std::basic_string\u003cchar\u003e \u0026"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"iterator __i1"}}},{"Name":{},"Type":{"Name":{"Relative":"iterator __i2"}}},{"Name":{},"Type":{"Name":{"Relative":"size_type __n"}}},{"Name":{},"Type":{"Name":{"Relative":"char __c"}}}]}
	std::basic_string<char> &  replace(iterator __i1 , iterator __i2 , size_type __n , char __c ) // 
	// {"Name":{"Relative":"replace"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"std::basic_string\u003cchar\u003e \u0026"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"iterator __i1"}}},{"Name":{},"Type":{"Name":{"Relative":"iterator __i2"}}},{"Name":{},"Type":{"Name":{"Relative":"_InputIterator __k1"}}},{"Name":{},"Type":{"Name":{"Relative":"_InputIterator __k2"}}}]}
	std::basic_string<char> &  replace(iterator __i1 , iterator __i2 , _InputIterator __k1 , _InputIterator __k2 ) // 
	// {"Name":{"Relative":"replace"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"std::basic_string\u003cchar\u003e \u0026"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"iterator __i1"}}},{"Name":{},"Type":{"Name":{"Relative":"iterator __i2"}}},{"Name":{},"Type":{"Name":{"Relative":"char *__k1"}}},{"Name":{},"Type":{"Name":{"Relative":"char *__k2"}}}]}
	std::basic_string<char> &  replace(iterator __i1 , iterator __i2 , char *__k1 , char *__k2 ) // 
	// {"Name":{"Relative":"replace"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"std::basic_string\u003cchar\u003e \u0026"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"iterator __i1"}}},{"Name":{},"Type":{"Name":{"Relative":"iterator __i2"}}},{"Name":{},"Type":{"Name":{"Relative":"const char *__k1"}}},{"Name":{},"Type":{"Name":{"Relative":"const char *__k2"}}}]}
	std::basic_string<char> &  replace(iterator __i1 , iterator __i2 , const char *__k1 , const char *__k2 ) // 
	// {"Name":{"Relative":"replace"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"std::basic_string\u003cchar\u003e \u0026"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"iterator __i1"}}},{"Name":{},"Type":{"Name":{"Relative":"iterator __i2"}}},{"Name":{},"Type":{"Name":{"Relative":"iterator __k1"}}},{"Name":{},"Type":{"Name":{"Relative":"iterator __k2"}}}]}
	std::basic_string<char> &  replace(iterator __i1 , iterator __i2 , iterator __k1 , iterator __k2 ) // 
	// {"Name":{"Relative":"replace"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"std::basic_string\u003cchar\u003e \u0026"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"iterator __i1"}}},{"Name":{},"Type":{"Name":{"Relative":"iterator __i2"}}},{"Name":{},"Type":{"Name":{"Relative":"const_iterator __k1"}}},{"Name":{},"Type":{"Name":{"Relative":"const_iterator __k2"}}}]}
	std::basic_string<char> &  replace(iterator __i1 , iterator __i2 , const_iterator __k1 , const_iterator __k2 ) // 
	// {"Name":{"Relative":"reserve"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"void"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"size_type __res_arg"}}}]}
	void  reserve(size_type __res_arg )                                              // 
	// {"Name":{"Relative":"resize"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"void"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"size_type __n"}}},{"Name":{},"Type":{"Name":{"Relative":"char __c"}}}]}
	void  resize(size_type __n , char __c )                                          // 
	// {"Name":{"Relative":"resize"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"void"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"size_type __n"}}}]}
	void  resize(size_type __n )                                                     // 
	// {"Name":{"Relative":"rfind"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"size_type"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"const std::basic_string\u003cchar\u003e \u0026__str"}}},{"Name":{},"Type":{"Name":{"Relative":"size_type __pos"}}}]}
	size_type  rfind(const std::basic_string<char> &__str , size_type __pos )        // 
	// {"Name":{"Relative":"rfind"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"size_type"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"const char *__s"}}},{"Name":{},"Type":{"Name":{"Relative":"size_type __pos"}}},{"Name":{},"Type":{"Name":{"Relative":"size_type __n"}}}]}
	size_type  rfind(const char *__s , size_type __pos , size_type __n )             // 
	// {"Name":{"Relative":"rfind"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"size_type"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"const char *__s"}}},{"Name":{},"Type":{"Name":{"Relative":"size_type __pos"}}}]}
	size_type  rfind(const char *__s , size_type __pos )                             // 
	// {"Name":{"Relative":"rfind"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"size_type"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"char __c"}}},{"Name":{},"Type":{"Name":{"Relative":"size_type __pos"}}}]}
	size_type  rfind(char __c , size_type __pos )                                    // 
	// {"Name":{"Relative":"size"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"size_type"}}}]}
	size_type  size()                                                                // 
	// {"Name":{"Relative":"substr"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"std::basic_string\u003cchar\u003e"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"size_type __pos"}}}]}
	std::basic_string<char>  substr(size_type __pos )                                // 
	// {"Name":{"Relative":"swap"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"void"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"std::basic_string\u003cchar\u003e \u0026__s"}}}]}
	void  swap(std::basic_string<char> &__s )                                        // 
