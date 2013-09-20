Fields:
	// {"Name":{"Relative":"m_Module"},"Type":{"Name":{"Relative":"asIScriptModule *"}}}
	asIScriptModule * m_Module                                                       // 
	// {"Name":{"Relative":"mEvStat"},"Type":{"Name":{"Relative":"ev_stat"}}}
	ev_stat mEvStat                                                                  // 
	// {"Name":{"Relative":"mLastmod"},"Type":{"Name":{"Relative":"long long"}}}
	long long mLastmod                                                               // 
	// {"Name":{"Relative":"mListeners"},"Type":{"Name":{"Relative":"std::vector\u003cSmartPointer\u003cResourceUpdatedListener\u003e \u003e"}}}
	std::vector<SmartPointer<ResourceUpdatedListener> > mListeners                   // 
	// {"Name":{"Relative":"mMutex"},"Type":{"Name":{"Relative":"Mutex"}}}
	Mutex mMutex                                                                     // 
	// {"Name":{"Relative":"mName"},"Type":{"Name":{"Relative":"StringID"}}}
	StringID mName                                                                   // 
	// {"Name":{"Relative":"mRefCount"},"Type":{"Name":{"Relative":"int"}}}
	int mRefCount                                                                    // 
Methods:
	// {"Name":{"Relative":"AddRef"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"void"}}}]}
	void  AddRef()                                                                   // 
	// {"Name":{"Relative":"AddResourceUpdatedListener"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"void"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"ResourceUpdatedListener *listener"}}}]}
	void  AddResourceUpdatedListener(ResourceUpdatedListener *listener )             // 
	// {"Name":{"Relative":"DelRef"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"void"}}}]}
	void  DelRef()                                                                   // 
	// {"Name":{"Relative":"GetFunction"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"asIScriptFunction *"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"const std::string \u0026decl"}}}]}
	asIScriptFunction *  GetFunction(const std::string &decl )                       // 
	// {"Name":{"Relative":"GetName"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"StringID"}}}]}
	StringID  GetName()                                                              // 
	// {"Name":{"Relative":"GetReferenceCount"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"int"}}}]}
	int  GetReferenceCount()                                                         // 
	// {"Name":{"Relative":"OnUpdated"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"void"}}}]}
	void  OnUpdated()                                                                // 
	// {"Name":{"Relative":"Open"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"std::istream *"}}}]}
	std::istream *  Open()                                                           // 
	// {"Name":{"Relative":"ReadAsString"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"void"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"char **data"}}},{"Name":{},"Type":{"Name":{"Relative":"int *size"}}}]}
	void  ReadAsString(char **data , int *size )                                     // 
	// {"Name":{"Relative":"Reload"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"void"}}}]}
	void  Reload()                                                                   // 
	// {"Name":{"Relative":"Resource::Reload"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"void"}}}]}
	void  Resource::Reload()                                                         // 
	// {"Name":{"Relative":"RemoveResourceUpdatedListener"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"void"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"ResourceUpdatedListener *listener"}}}]}
	void  RemoveResourceUpdatedListener(ResourceUpdatedListener *listener )          // 
	// {"Name":{"Relative":"SetName"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"void"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"StringID"}}}]}
	void  SetName(StringID )                                                         // 
	// {"Name":{"Relative":"Update"},"Returns":[{"Name":{},"Type":{"Name":{"Relative":"void"}}}]}
	void  Update()                                                                   // 
