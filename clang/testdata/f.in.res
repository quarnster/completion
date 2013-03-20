Fields:
	asIScriptModule * m_Module                                                       // 
	ev_stat mEvStat                                                                  // 
	long long mLastmod                                                               // 
	std::vector<SmartPointer<ResourceUpdatedListener> > mListeners                   // 
	Mutex mMutex                                                                     // 
	StringID mName                                                                   // 
	int mRefCount                                                                    // 
Methods:
	void  AddRef()                                                                   // 
	void  AddResourceUpdatedListener(ResourceUpdatedListener *listener )             // 
	void  DelRef()                                                                   // 
	asIScriptFunction *  GetFunction(const std::string &decl )                       // 
	StringID  GetName()                                                              // 
	int  GetReferenceCount()                                                         // 
	void  OnUpdated()                                                                // 
	std::istream *  Open()                                                           // 
	void  ReadAsString(char **data , int *size )                                     // 
	void  Reload()                                                                   // 
	void  Resource::Reload()                                                         // 
	void  RemoveResourceUpdatedListener(ResourceUpdatedListener *listener )          // 
	void  SetName(StringID )                                                         // 
	void  Update()                                                                   // 
