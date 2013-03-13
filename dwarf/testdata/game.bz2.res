asCObjectType
Fields:
	public asCString name
	public asSNameSpace* nameSpace
	public int size
	public asCArray<asCObjectProperty *> properties
	public asCArray<int> methods
	public asCArray<asCObjectType *> interfaces
	public asCArray<asSEnumValue *> enumValues
	public asCObjectType* derivedFrom
	public asCArray<asCScriptFunction *> virtualFunctionTable
	public asDWORD flags
	public asDWORD accessMask
	public asSTypeBehaviour beh
	public asCArray<asCDataType> templateSubTypes
	public bool acceptValueSubType
	public bool acceptRefSubType
	public asCScriptEngine* engine
	public asCModule* module
	public asCArray<unsigned long> userData
	protected asCAtomic refCount
	protected bool gcFlag
Methods:
	public asIScriptEngine*  GetEngine() const
	public const char*  GetConfigGroup() const
	public asDWORD  GetAccessMask() const
	public int  AddRef() const
	public int  Release() const
	public const char*  GetName() const
	public const char*  GetNamespace() const
	public asIObjectType*  GetBaseType() const
	public bool  DerivesFrom(const asIObjectType* ) const
	public asDWORD  GetFlags() const
	public asUINT  GetSize() const
	public int  GetTypeId() const
	public int  GetSubTypeId(asUINT ) const
	public asIObjectType*  GetSubType(asUINT ) const
	public asUINT  GetSubTypeCount() const
	public asUINT  GetInterfaceCount() const
	public asIObjectType*  GetInterface(asUINT ) const
	public bool  Implements(const asIObjectType* ) const
	public asUINT  GetFactoryCount() const
	public asIScriptFunction*  GetFactoryByIndex(asUINT ) const
	public asIScriptFunction*  GetFactoryByDecl(const char* ) const
	public asUINT  GetMethodCount() const
	public asIScriptFunction*  GetMethodByIndex(asUINT , bool ) const
	public asIScriptFunction*  GetMethodByName(const char* , bool ) const
	public asIScriptFunction*  GetMethodByDecl(const char* , bool ) const
	public asUINT  GetPropertyCount() const
	public int  GetProperty(asUINT , const char** , int* , bool* , int* , bool* , asDWORD* ) const
	public const char*  GetPropertyDeclaration(asUINT ) const
	public asUINT  GetBehaviourCount() const
	public asIScriptFunction*  GetBehaviourByIndex(asUINT , asEBehaviours* ) const
	public void*  SetUserData(void* , asPWORD )
	public void*  GetUserData(asPWORD ) const
	public void  asCObjectType(asCScriptEngine* )
	public void  ~asCObjectType()
	public void  Orphan(asCModule* )
	public int  GetRefCount()
	public void  SetGCFlag()
	public bool  GetGCFlag()
	public void  EnumReferences(asIScriptEngine* )
	public void  ReleaseAllHandles(asIScriptEngine* )
	public void  ReleaseAllFunctions()
	public bool  IsInterface() const
	public bool  IsShared() const
	public asCObjectProperty*  AddPropertyToClass(const & asCString , const & asCDataType , bool )
	public void  ReleaseAllProperties()
	protected void  asCObjectType()
asCScriptFunction
Fields:
	public asCAtomic refCount
	public bool gcFlag
	public asCScriptEngine* engine
	public asCModule* module
	public void* userData
	public asCString name
	public asCDataType returnType
	public asCArray<asCDataType> parameterTypes
	public asCArray<asETypeModifiers> inOutFlags
	public asCArray<asCString *> defaultArgs
	public bool isReadOnly
	public bool isPrivate
	public bool isFinal
	public bool isOverride
	public asCObjectType* objectType
	public int signatureId
	public int id
	public asEFuncType funcType
	public asDWORD accessMask
	public bool isShared
	public asSNameSpace* nameSpace
	public asCArray<unsigned int> byteCode
	public asDWORD variableSpace
	public asCArray<asCObjectType *> objVariableTypes
	public asCArray<asCScriptFunction *> funcVariableTypes
	public asCArray<int> objVariablePos
	public asUINT objVariablesOnHeap
	public asCArray<asSObjectVariableInfo> objVariableInfo
	public asCArray<asSScriptVariable *> variables
	public int stackNeeded
	public asCArray<int> lineNumbers
	public int scriptSectionIdx
	public asCArray<int> sectionIdxs
	public bool dontCleanUpOnException
	public int vfTableIdx
	public asSSystemFunctionInterface* sysFuncIntf
	public asJITFunction jitFunction
Methods:
	public asIScriptEngine*  GetEngine() const
	public int  AddRef() const
	public int  Release() const
	public int  GetId() const
	public asEFuncType  GetFuncType() const
	public const char*  GetModuleName() const
	public const char*  GetScriptSectionName() const
	public const char*  GetConfigGroup() const
	public asDWORD  GetAccessMask() const
	public asIObjectType*  GetObjectType() const
	public const char*  GetObjectName() const
	public const char*  GetName() const
	public const char*  GetNamespace() const
	public const char*  GetDeclaration(bool , bool ) const
	public bool  IsReadOnly() const
	public bool  IsPrivate() const
	public bool  IsFinal() const
	public bool  IsOverride() const
	public bool  IsShared() const
	public asUINT  GetParamCount() const
	public int  GetParamTypeId(asUINT , asDWORD* ) const
	public int  GetReturnTypeId() const
	public int  GetTypeId() const
	public bool  IsCompatibleWithTypeId(int ) const
	public asUINT  GetVarCount() const
	public int  GetVar(asUINT , const char** , int* ) const
	public const char*  GetVarDecl(asUINT ) const
	public int  FindNextLineWithCode(int ) const
	public asDWORD*  GetByteCode(asUINT* )
	public void*  SetUserData(void* )
	public void*  GetUserData() const
	public void  asCScriptFunction(asCScriptEngine* , asCModule* , asEFuncType )
	public void  ~asCScriptFunction()
	public void  DestroyInternal()
	public void  Orphan(asIScriptModule* )
	public void  AddVariable(& asCString , & asCDataType , int )
	public int  GetSpaceNeededForArguments()
	public int  GetSpaceNeededForReturnValue()
	public asCString  GetDeclarationStr(bool , bool ) const
	public int  GetLineNumber(int , int* )
	public void  ComputeSignatureId()
	public bool  IsSignatureEqual(const asCScriptFunction* ) const
	public bool  IsSignatureExceptNameEqual(const asCScriptFunction* ) const
	public bool  IsSignatureExceptNameEqual(const & asCDataType , & asCArray<asCDataType> , & asCArray<asETypeModifiers> , const asCObjectType* , bool ) const
	public bool  IsSignatureExceptNameAndReturnTypeEqual(const asCScriptFunction* ) const
	public bool  IsSignatureExceptNameAndReturnTypeEqual(& asCArray<asCDataType> , & asCArray<asETypeModifiers> , const asCObjectType* , bool ) const
	public bool  DoesReturnOnStack() const
	public void  JITCompile()
	public void  AddReferences()
	public void  ReleaseReferences()
	public asCGlobalProperty*  GetPropertyByGlobalVarPtr(void* )
	public int  GetRefCount()
	public void  SetFlag()
	public bool  GetFlag()
	public void  EnumReferences(asIScriptEngine* )
	public void  ReleaseAllHandles(asIScriptEngine* )
