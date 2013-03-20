Fields:
	boost::regex_constants::_match_flags format_all                                  // 
	boost::regex_constants::_match_flags format_default                              // 
	boost::regex_constants::_match_flags format_first_only                           // 
	boost::regex_constants::_match_flags format_no_copy                              // 
	boost::regex_constants::_match_flags format_perl                                 // 
	boost::regex_constants::_match_flags format_sed                                  // 
	boost::regex_constants::_match_flags match_all                                   // 
	boost::regex_constants::_match_flags match_any                                   // 
	boost::regex_constants::_match_flags match_continuous                            // 
	boost::regex_constants::_match_flags match_default                               // 
	boost::regex_constants::_match_flags match_extra                                 // 
	boost::regex_constants::_match_flags match_nosubs                                // 
	boost::regex_constants::_match_flags match_not_bob                               // 
	boost::regex_constants::_match_flags match_not_bol                               // 
	boost::regex_constants::_match_flags match_not_bow                               // 
	boost::regex_constants::_match_flags match_not_dot_newline                       // 
	boost::regex_constants::_match_flags match_not_dot_null                          // 
	boost::regex_constants::_match_flags match_not_eob                               // 
	boost::regex_constants::_match_flags match_not_eol                               // 
	boost::regex_constants::_match_flags match_not_eow                               // 
	boost::regex_constants::_match_flags match_not_null                              // 
	boost::regex_constants::_match_flags match_partial                               // 
	boost::regex_constants::_match_flags match_perl                                  // 
	boost::regex_constants::_match_flags match_posix                                 // 
	boost::regex_constants::_match_flags match_prev_avail                            // 
	boost::regex_constants::_match_flags match_single_line                           // 
	boost::memory_order memory_order_acq_rel                                         // 
	boost::memory_order memory_order_acquire                                         // 
	boost::memory_order memory_order_consume                                         // 
	boost::memory_order memory_order_relaxed                                         // 
	boost::memory_order memory_order_release                                         // 
	boost::memory_order memory_order_seq_cst                                         // 
Methods:
	T *  addressof(T &v )                                                            // 
	bool  atomic_compare_exchange(shared_ptr<T> *p , shared_ptr<T> *v , shared_ptr<T> w ) // 
	bool  atomic_compare_exchange_explicit(shared_ptr<T> *p , shared_ptr<T> *v , shared_ptr<T> w , boost::memory_order , boost::memory_order ) // 
	shared_ptr<T>  atomic_exchange(shared_ptr<T> *p , shared_ptr<T> r )              // 
	shared_ptr<T>  atomic_exchange_explicit(shared_ptr<T> *p , shared_ptr<T> r , boost::memory_order ) // 
	bool  atomic_is_lock_free(const shared_ptr<T> * )                                // 
	shared_ptr<T>  atomic_load(const shared_ptr<T> *p )                              // 
	shared_ptr<T>  atomic_load_explicit(const shared_ptr<T> *p , boost::memory_order ) // 
	void  atomic_store(shared_ptr<T> *p , shared_ptr<T> r )                          // 
	void  atomic_store_explicit(shared_ptr<T> *p , shared_ptr<T> r , boost::memory_order ) // 
	void  checked_array_delete(T *x )                                                // 
	void  checked_delete(T *x )                                                      // 
	shared_ptr<T>  const_pointer_cast<<#class T#>>(const shared_ptr<U> &r )          // 
	const reference_wrapper<const T>  cref(const T &t )                              // 
	shared_ptr<T>  dynamic_pointer_cast<<#class T#>>(const shared_ptr<U> &r )        // 
	exception_detail::clone_impl<T>  enable_current_exception(const T &x )           // 
	typename exception_detail::enable_error_info_return_type<T>::type  enable_error_info(const T &x ) // 
	T (&)[N]  get_c_array(boost::array<T, N> &arg )                                  // 
	const T (&)[N]  get_c_array(const boost::array<T, N> &arg )                      // 
	D *  get_deleter<<#class D#>>(const shared_ptr<T> &p )                           // 
	typename shared_ptr<T>::element_type *  get_pointer(const shared_ptr<T> &p )     // 
	T *  get_pointer(const scoped_ptr<T> &p )                                        // 
	T *  get_pointer(const reference_wrapper<T> &r )                                 // 
	void  hash_combine(std::size_t &seed , const T &v )                              // 
	std::size_t  hash_range(It first , It last )                                     // 
	void  hash_range(std::size_t &seed , It first , It last )                        // 
	std::size_t  hash_value(const array<T, N> &arr )                                 // 
	std::size_t  hash_value(const boost::shared_ptr<T> &p )                          // 
	typename boost::hash_detail::basic_numbers<T>::type  hash_value(T v )            // 
	typename boost::hash_detail::long_numbers<T>::type  hash_value(T v )             // 
	typename boost::hash_detail::ulong_numbers<T>::type  hash_value(T v )            // 
	typename boost::enable_if<boost::is_enum<T>, std::size_t>::type  hash_value(T v ) // 
	std::size_t  hash_value(T *const &v )                                            // 
	std::size_t  hash_value(const T (&x)[N] )                                        // 
	std::size_t  hash_value(T (&x)[N] )                                              // 
	std::size_t  hash_value(const std::basic_string<Ch, std::char_traits<Ch>, A> &v ) // 
	typename boost::hash_detail::float_numbers<T>::type  hash_value(T v )            // 
	std::size_t  hash_value(const std::pair<A, B> &v )                               // 
	std::size_t  hash_value(const std::vector<T, A> &v )                             // 
	std::size_t  hash_value(const std::list<T, A> &v )                               // 
	std::size_t  hash_value(const std::deque<T, A> &v )                              // 
	std::size_t  hash_value(const std::set<K, C, A> &v )                             // 
	std::size_t  hash_value(const std::multiset<K, C, A> &v )                        // 
	std::size_t  hash_value(const std::map<K, T, C, A> &v )                          // 
	std::size_t  hash_value(const std::multimap<K, T, C, A> &v )                     // 
	std::size_t  hash_value(const std::complex<T> &v )                               // 
	regex_iterator<const charT *, charT, traits>  make_regex_iterator(const charT *p , const basic_regex<charT, traits> &e , regex_constants::match_flag_type m ) // 
	regex_iterator<typename std::basic_string<charT, ST, SA>::const_iterator, charT, traits>  make_regex_iterator(const std::basic_string<charT, ST, SA> &p , const basic_regex<charT, traits> &e , regex_constants::match_flag_type m ) // 
	regex_token_iterator<const charT *, charT, traits>  make_regex_token_iterator(const charT *p , const basic_regex<charT, traits> &e , int submatch ) // 
	regex_token_iterator<typename std::basic_string<charT, ST, SA>::const_iterator, charT, traits>  make_regex_token_iterator(const std::basic_string<charT, ST, SA> &p , const basic_regex<charT, traits> &e , int submatch ) // 
	regex_token_iterator<const charT *, charT, traits>  make_regex_token_iterator(const charT *p , const basic_regex<charT, traits> &e , const int (&submatch)[N] , regex_constants::match_flag_type m ) // 
	regex_token_iterator<typename std::basic_string<charT, ST, SA>::const_iterator, charT, traits>  make_regex_token_iterator(const std::basic_string<charT, ST, SA> &p , const basic_regex<charT, traits> &e , const int (&submatch)[N] , regex_constants::match_flag_type m ) // 
	regex_token_iterator<const charT *, charT, traits>  make_regex_token_iterator(const charT *p , const basic_regex<charT, traits> &e , const std::vector<int> &submatch , regex_constants::match_flag_type m ) // 
	regex_token_iterator<typename std::basic_string<charT, ST, SA>::const_iterator, charT, traits>  make_regex_token_iterator(const std::basic_string<charT, ST, SA> &p , const basic_regex<charT, traits> &e , const std::vector<int> &submatch , regex_constants::match_flag_type m ) // 
	const reference_wrapper<T>  ref(T &t )                                           // 
	OutputIterator  regex_format(OutputIterator out , const match_results<Iterator, Allocator> &m , Functor fmt , match_flag_type flags ) // 
	std::basic_string<typename match_results<Iterator, Allocator>::char_type>  regex_format(const match_results<Iterator, Allocator> &m , Functor fmt , match_flag_type flags ) // 
	unsigned int  regex_grep(Predicate foo , BidiIterator first , BidiIterator last , const basic_regex<charT, traits> &e , match_flag_type flags ) // 
	unsigned int  regex_grep(Predicate foo , const charT *str , const basic_regex<charT, traits> &e , match_flag_type flags ) // 
	unsigned int  regex_grep(Predicate foo , const std::basic_string<charT, ST, SA> &s , const basic_regex<charT, traits> &e , match_flag_type flags ) // 
	bool  regex_match(BidiIterator first , BidiIterator last , match_results<BidiIterator, Allocator> &m , const basic_regex<charT, traits> &e , match_flag_type flags ) // 
	bool  regex_match(iterator first , iterator last , const basic_regex<charT, traits> &e , match_flag_type flags ) // 
	bool  regex_match(const charT *str , match_results<const charT *, Allocator> &m , const basic_regex<charT, traits> &e , match_flag_type flags ) // 
	bool  regex_match(const std::basic_string<charT, ST, SA> &s , match_results<typename std::basic_string<charT, ST, SA>::const_iterator, Allocator> &m , const basic_regex<charT, traits> &e , match_flag_type flags ) // 
	bool  regex_match(const charT *str , const basic_regex<charT, traits> &e , match_flag_type flags ) // 
	bool  regex_match(const std::basic_string<charT, ST, SA> &s , const basic_regex<charT, traits> &e , match_flag_type flags ) // 
	OutputIterator  regex_merge(OutputIterator out , Iterator first , Iterator last , const basic_regex<charT, traits> &e , const charT *fmt , match_flag_type flags ) // 
	OutputIterator  regex_merge(OutputIterator out , Iterator first , Iterator last , const basic_regex<charT, traits> &e , const std::basic_string<charT> &fmt , match_flag_type flags ) // 
	std::basic_string<charT>  regex_merge(const std::basic_string<charT> &s , const basic_regex<charT, traits> &e , const charT *fmt , match_flag_type flags ) // 
	std::basic_string<charT>  regex_merge(const std::basic_string<charT> &s , const basic_regex<charT, traits> &e , const std::basic_string<charT> &fmt , match_flag_type flags ) // 
	OutputIterator  regex_replace(OutputIterator out , BidirectionalIterator first , BidirectionalIterator last , const basic_regex<charT, traits> &e , Formatter fmt , match_flag_type flags ) // 
	std::basic_string<charT>  regex_replace(const std::basic_string<charT> &s , const basic_regex<charT, traits> &e , Formatter fmt , match_flag_type flags ) // 
	bool  regex_search(BidiIterator first , BidiIterator last , match_results<BidiIterator, Allocator> &m , const basic_regex<charT, traits> &e , match_flag_type flags ) // 
	bool  regex_search(BidiIterator first , BidiIterator last , match_results<BidiIterator, Allocator> &m , const basic_regex<charT, traits> &e , match_flag_type flags , BidiIterator base ) // 
	bool  regex_search(const charT *str , match_results<const charT *, Allocator> &m , const basic_regex<charT, traits> &e , match_flag_type flags ) // 
	bool  regex_search(const std::basic_string<charT, ST, SA> &s , match_results<typename std::basic_string<charT, ST, SA>::const_iterator, Allocator> &m , const basic_regex<charT, traits> &e , match_flag_type flags ) // 
	bool  regex_search(BidiIterator first , BidiIterator last , const basic_regex<charT, traits> &e , match_flag_type flags ) // 
	bool  regex_search(const charT *str , const basic_regex<charT, traits> &e , match_flag_type flags ) // 
	bool  regex_search(const std::basic_string<charT, ST, SA> &s , const basic_regex<charT, traits> &e , match_flag_type flags ) // 
	std::size_t  regex_split(OutputIterator out , std::basic_string<charT, Traits1, Alloc1> &s , const basic_regex<charT, Traits2> &e , match_flag_type flags , std::size_t max_split ) // 
	std::size_t  regex_split(OutputIterator out , std::basic_string<charT, Traits1, Alloc1> &s , const basic_regex<charT, Traits2> &e , match_flag_type flags ) // 
	std::size_t  regex_split(OutputIterator out , std::basic_string<charT, Traits1, Alloc1> &s ) // 
	shared_ptr<T>  reinterpret_pointer_cast<<#class T#>>(const shared_ptr<U> &r )    // 
	shared_ptr<T>  static_pointer_cast<<#class T#>>(const shared_ptr<U> &r )         // 
	void  swap(T1 &left , T2 &right )                                                // 
	void  swap(array<T, N> &x , array<T, N> &y )                                     // 
	void  swap(shared_ptr<T> &a , shared_ptr<T> &b )                                 // 
	void  swap(scoped_ptr<T> &a , scoped_ptr<T> &b )                                 // 
	void  swap(scoped_array<T> &a , scoped_array<T> &b )                             // 
	void  swap(basic_regex<charT, traits> &e1 , basic_regex<charT, traits> &e2 )     // 
	void  swap(match_results<BidiIterator, Allocator> &a , match_results<BidiIterator, Allocator> &b ) // 
	void  throw_exception(const E &e )                                               // 
	void  throw_exception_assert_compatibility(const std::exception & )              // 
	typename unwrap_reference<T>::type &  unwrap_ref(T &t )                          // 
