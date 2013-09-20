Methods:
	// {"Name":{"Relative":"assign"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"void"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"size_type __n"}}},{"Name":{},"Type":{"Name":{"Relative":"const value_type \u0026__val"}}}]}
	void  assign(size_type __n , const value_type &__val )                           // 
	// {"Name":{"Relative":"assign"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"void"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"_InputIterator __first"}}},{"Name":{},"Type":{"Name":{"Relative":"_InputIterator __last"}}}]}
	void  assign(_InputIterator __first , _InputIterator __last )                    // 
	// {"Name":{"Relative":"at"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"reference"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"size_type __n"}}}]}
	reference  at(size_type __n )                                                    // 
	// {"Name":{"Relative":"at"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"const_reference"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"size_type __n"}}}]}
	const_reference  at(size_type __n )                                              // 
	// {"Name":{"Relative":"back"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"reference"}}}]}
	reference  back()                                                                // 
	// {"Name":{"Relative":"back"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"const_reference"}}}]}
	const_reference  back()                                                          // 
	// {"Name":{"Relative":"begin"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"iterator"}}}]}
	iterator  begin()                                                                // 
	// {"Name":{"Relative":"begin"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"const_iterator"}}}]}
	const_iterator  begin()                                                          // 
	// {"Name":{"Relative":"capacity"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"size_type"}}}]}
	size_type  capacity()                                                            // 
	// {"Name":{"Relative":"clear"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"void"}}}]}
	void  clear()                                                                    // 
	// {"Name":{"Relative":"data"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"pointer"}}}]}
	pointer  data()                                                                  // 
	// {"Name":{"Relative":"data"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"const_pointer"}}}]}
	const_pointer  data()                                                            // 
	// {"Name":{"Relative":"empty"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"bool"}}}]}
	bool  empty()                                                                    // 
	// {"Name":{"Relative":"end"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"iterator"}}}]}
	iterator  end()                                                                  // 
	// {"Name":{"Relative":"end"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"const_iterator"}}}]}
	const_iterator  end()                                                            // 
	// {"Name":{"Relative":"erase"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"iterator"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"iterator __position"}}}]}
	iterator  erase(iterator __position )                                            // 
	// {"Name":{"Relative":"erase"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"iterator"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"iterator __first"}}},{"Name":{},"Type":{"Name":{"Relative":"iterator __last"}}}]}
	iterator  erase(iterator __first , iterator __last )                             // 
	// {"Name":{"Relative":"front"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"reference"}}}]}
	reference  front()                                                               // 
	// {"Name":{"Relative":"front"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"const_reference"}}}]}
	const_reference  front()                                                         // 
	// {"Name":{"Relative":"get_allocator"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"allocator_type"}}}]}
	allocator_type  get_allocator()                                                  // 
	// {"Name":{"Relative":"insert"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"iterator"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"iterator __position"}}},{"Name":{},"Type":{"Name":{"Relative":"const value_type \u0026__x"}}}]}
	iterator  insert(iterator __position , const value_type &__x )                   // 
	// {"Name":{"Relative":"insert"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"void"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"iterator __position"}}},{"Name":{},"Type":{"Name":{"Relative":"size_type __n"}}},{"Name":{},"Type":{"Name":{"Relative":"const value_type \u0026__x"}}}]}
	void  insert(iterator __position , size_type __n , const value_type &__x )       // 
	// {"Name":{"Relative":"insert"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"void"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"iterator __position"}}},{"Name":{},"Type":{"Name":{"Relative":"_InputIterator __first"}}},{"Name":{},"Type":{"Name":{"Relative":"_InputIterator __last"}}}]}
	void  insert(iterator __position , _InputIterator __first , _InputIterator __last ) // 
	// {"Name":{"Relative":"max_size"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"size_type"}}}]}
	size_type  max_size()                                                            // 
	// {"Name":{"Relative":"pop_back"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"void"}}}]}
	void  pop_back()                                                                 // 
	// {"Name":{"Relative":"push_back"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"void"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"const value_type \u0026__x"}}}]}
	void  push_back(const value_type &__x )                                          // 
	// {"Name":{"Relative":"rbegin"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"reverse_iterator"}}}]}
	reverse_iterator  rbegin()                                                       // 
	// {"Name":{"Relative":"rbegin"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"const_reverse_iterator"}}}]}
	const_reverse_iterator  rbegin()                                                 // 
	// {"Name":{"Relative":"rend"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"reverse_iterator"}}}]}
	reverse_iterator  rend()                                                         // 
	// {"Name":{"Relative":"rend"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"const_reverse_iterator"}}}]}
	const_reverse_iterator  rend()                                                   // 
	// {"Name":{"Relative":"reserve"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"void"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"size_type __n"}}}]}
	void  reserve(size_type __n )                                                    // 
	// {"Name":{"Relative":"resize"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"void"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"size_type __new_size"}}},{"Name":{},"Type":{"Name":{"Relative":"value_type __x"}}}]}
	void  resize(size_type __new_size , value_type __x )                             // 
	// {"Name":{"Relative":"size"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"size_type"}}}]}
	size_type  size()                                                                // 
	// {"Name":{"Relative":"swap"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"void"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"std::vector\u003cStringID, std::allocator\u003cStringID\u003e \u003e \u0026__x"}}}]}
	void  swap(std::vector<StringID, std::allocator<StringID> > &__x )               // 
	// {"Name":{"Relative":"vector"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"void"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"_InputIterator __first"}}},{"Name":{},"Type":{"Name":{"Relative":"_InputIterator __last"}}},{"Name":{},"Type":{"Name":{"Relative":"const allocator_type \u0026__a"}}}]}
	void  vector(_InputIterator __first , _InputIterator __last , const allocator_type &__a ) // 
