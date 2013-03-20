asCObjectType
Fields:
	public asCString name                                                            // dwarf://field/./testdata/game.bz2;306871
	public asSNameSpace* nameSpace                                                   // dwarf://field/./testdata/game.bz2;306886
	public int size                                                                  // dwarf://field/./testdata/game.bz2;306901
	public asCArray<asCObjectProperty *> properties                                  // dwarf://field/./testdata/game.bz2;306916
	public asCArray<int> methods                                                     // dwarf://field/./testdata/game.bz2;306931
	public asCArray<asCObjectType *> interfaces                                      // dwarf://field/./testdata/game.bz2;306946
	public asCArray<asSEnumValue *> enumValues                                       // dwarf://field/./testdata/game.bz2;306961
	public asCObjectType* derivedFrom                                                // dwarf://field/./testdata/game.bz2;306977
	public asCArray<asCScriptFunction *> virtualFunctionTable                        // dwarf://field/./testdata/game.bz2;306993
	public asDWORD flags                                                             // dwarf://field/./testdata/game.bz2;307009
	public asDWORD accessMask                                                        // dwarf://field/./testdata/game.bz2;307025
	public asSTypeBehaviour beh                                                      // dwarf://field/./testdata/game.bz2;307041
	public asCArray<asCDataType> templateSubTypes                                    // dwarf://field/./testdata/game.bz2;307057
	public bool acceptValueSubType                                                   // dwarf://field/./testdata/game.bz2;307073
	public bool acceptRefSubType                                                     // dwarf://field/./testdata/game.bz2;307089
	public asCScriptEngine* engine                                                   // dwarf://field/./testdata/game.bz2;307105
	public asCModule* module                                                         // dwarf://field/./testdata/game.bz2;307121
	public asCArray<unsigned long> userData                                          // dwarf://field/./testdata/game.bz2;307137
	protected asCAtomic refCount                                                     // dwarf://field/./testdata/game.bz2;307153
	protected bool gcFlag                                                            // dwarf://field/./testdata/game.bz2;307169
Methods:
	public asIScriptEngine*  GetEngine() const                                       // dwarf://method/./testdata/game.bz2;307185
	public const char*  GetConfigGroup() const                                       // dwarf://method/./testdata/game.bz2;307218
	public asDWORD  GetAccessMask() const                                            // dwarf://method/./testdata/game.bz2;307251
	public int  AddRef() const                                                       // dwarf://method/./testdata/game.bz2;307284
	public int  Release() const                                                      // dwarf://method/./testdata/game.bz2;307317
	public const char*  GetName() const                                              // dwarf://method/./testdata/game.bz2;307350
	public const char*  GetNamespace() const                                         // dwarf://method/./testdata/game.bz2;307383
	public asIObjectType*  GetBaseType() const                                       // dwarf://method/./testdata/game.bz2;307416
	public bool  DerivesFrom(const asIObjectType* ) const                            // dwarf://method/./testdata/game.bz2;307449
	public asDWORD  GetFlags() const                                                 // dwarf://method/./testdata/game.bz2;307487
	public asUINT  GetSize() const                                                   // dwarf://method/./testdata/game.bz2;307520
	public int  GetTypeId() const                                                    // dwarf://method/./testdata/game.bz2;307553
	public int  GetSubTypeId(asUINT ) const                                          // dwarf://method/./testdata/game.bz2;307586
	public asIObjectType*  GetSubType(asUINT ) const                                 // dwarf://method/./testdata/game.bz2;307624
	public asUINT  GetSubTypeCount() const                                           // dwarf://method/./testdata/game.bz2;307662
	public asUINT  GetInterfaceCount() const                                         // dwarf://method/./testdata/game.bz2;307695
	public asIObjectType*  GetInterface(asUINT ) const                               // dwarf://method/./testdata/game.bz2;307728
	public bool  Implements(const asIObjectType* ) const                             // dwarf://method/./testdata/game.bz2;307766
	public asUINT  GetFactoryCount() const                                           // dwarf://method/./testdata/game.bz2;307804
	public asIScriptFunction*  GetFactoryByIndex(asUINT ) const                      // dwarf://method/./testdata/game.bz2;307837
	public asIScriptFunction*  GetFactoryByDecl(const char* ) const                  // dwarf://method/./testdata/game.bz2;307875
	public asUINT  GetMethodCount() const                                            // dwarf://method/./testdata/game.bz2;307913
	public asIScriptFunction*  GetMethodByIndex(asUINT , bool ) const                // dwarf://method/./testdata/game.bz2;307946
	public asIScriptFunction*  GetMethodByName(const char* , bool ) const            // dwarf://method/./testdata/game.bz2;307989
	public asIScriptFunction*  GetMethodByDecl(const char* , bool ) const            // dwarf://method/./testdata/game.bz2;308032
	public asUINT  GetPropertyCount() const                                          // dwarf://method/./testdata/game.bz2;308075
	public int  GetProperty(asUINT , const char** , int* , bool* , int* , bool* , asDWORD* ) const // dwarf://method/./testdata/game.bz2;308108
	public const char*  GetPropertyDeclaration(asUINT ) const                        // dwarf://method/./testdata/game.bz2;308176
	public asUINT  GetBehaviourCount() const                                         // dwarf://method/./testdata/game.bz2;308214
	public asIScriptFunction*  GetBehaviourByIndex(asUINT , asEBehaviours* ) const   // dwarf://method/./testdata/game.bz2;308247
	public void*  SetUserData(void* , asPWORD )                                      // dwarf://method/./testdata/game.bz2;308290
	public void*  GetUserData(asPWORD ) const                                        // dwarf://method/./testdata/game.bz2;308333
	public void  asCObjectType(asCScriptEngine* )                                    // dwarf://method/./testdata/game.bz2;308371
	public void  ~asCObjectType()                                                    // dwarf://method/./testdata/game.bz2;308393
	public void  Orphan(asCModule* )                                                 // dwarf://method/./testdata/game.bz2;308418
	public int  GetRefCount()                                                        // dwarf://method/./testdata/game.bz2;308444
	public void  SetGCFlag()                                                         // dwarf://method/./testdata/game.bz2;308469
	public bool  GetGCFlag()                                                         // dwarf://method/./testdata/game.bz2;308490
	public void  EnumReferences(asIScriptEngine* )                                   // dwarf://method/./testdata/game.bz2;308515
	public void  ReleaseAllHandles(asIScriptEngine* )                                // dwarf://method/./testdata/game.bz2;308541
	public void  ReleaseAllFunctions()                                               // dwarf://method/./testdata/game.bz2;308567
	public bool  IsInterface() const                                                 // dwarf://method/./testdata/game.bz2;308588
	public bool  IsShared() const                                                    // dwarf://method/./testdata/game.bz2;308613
	public asCObjectProperty*  AddPropertyToClass(const & asCString , const & asCDataType , bool ) // dwarf://method/./testdata/game.bz2;308638
	public void  ReleaseAllProperties()                                              // dwarf://method/./testdata/game.bz2;308678
	protected void  asCObjectType()                                                  // dwarf://method/./testdata/game.bz2;308699
asCScriptFunction
Fields:
	public asCAtomic refCount                                                        // dwarf://field/./testdata/game.bz2;330664
	public bool gcFlag                                                               // dwarf://field/./testdata/game.bz2;330679
	public asCScriptEngine* engine                                                   // dwarf://field/./testdata/game.bz2;330694
	public asCModule* module                                                         // dwarf://field/./testdata/game.bz2;330709
	public void* userData                                                            // dwarf://field/./testdata/game.bz2;330724
	public asCString name                                                            // dwarf://field/./testdata/game.bz2;330739
	public asCDataType returnType                                                    // dwarf://field/./testdata/game.bz2;330754
	public asCArray<asCDataType> parameterTypes                                      // dwarf://field/./testdata/game.bz2;330769
	public asCArray<asETypeModifiers> inOutFlags                                     // dwarf://field/./testdata/game.bz2;330784
	public asCArray<asCString *> defaultArgs                                         // dwarf://field/./testdata/game.bz2;330800
	public bool isReadOnly                                                           // dwarf://field/./testdata/game.bz2;330816
	public bool isPrivate                                                            // dwarf://field/./testdata/game.bz2;330832
	public bool isFinal                                                              // dwarf://field/./testdata/game.bz2;330848
	public bool isOverride                                                           // dwarf://field/./testdata/game.bz2;330864
	public asCObjectType* objectType                                                 // dwarf://field/./testdata/game.bz2;330880
	public int signatureId                                                           // dwarf://field/./testdata/game.bz2;330896
	public int id                                                                    // dwarf://field/./testdata/game.bz2;330912
	public asEFuncType funcType                                                      // dwarf://field/./testdata/game.bz2;330928
	public asDWORD accessMask                                                        // dwarf://field/./testdata/game.bz2;330944
	public bool isShared                                                             // dwarf://field/./testdata/game.bz2;330960
	public asSNameSpace* nameSpace                                                   // dwarf://field/./testdata/game.bz2;330976
	public asCArray<unsigned int> byteCode                                           // dwarf://field/./testdata/game.bz2;330992
	public asDWORD variableSpace                                                     // dwarf://field/./testdata/game.bz2;331008
	public asCArray<asCObjectType *> objVariableTypes                                // dwarf://field/./testdata/game.bz2;331024
	public asCArray<asCScriptFunction *> funcVariableTypes                           // dwarf://field/./testdata/game.bz2;331040
	public asCArray<int> objVariablePos                                              // dwarf://field/./testdata/game.bz2;331056
	public asUINT objVariablesOnHeap                                                 // dwarf://field/./testdata/game.bz2;331072
	public asCArray<asSObjectVariableInfo> objVariableInfo                           // dwarf://field/./testdata/game.bz2;331088
	public asCArray<asSScriptVariable *> variables                                   // dwarf://field/./testdata/game.bz2;331104
	public int stackNeeded                                                           // dwarf://field/./testdata/game.bz2;331120
	public asCArray<int> lineNumbers                                                 // dwarf://field/./testdata/game.bz2;331136
	public int scriptSectionIdx                                                      // dwarf://field/./testdata/game.bz2;331152
	public asCArray<int> sectionIdxs                                                 // dwarf://field/./testdata/game.bz2;331168
	public bool dontCleanUpOnException                                               // dwarf://field/./testdata/game.bz2;331184
	public int vfTableIdx                                                            // dwarf://field/./testdata/game.bz2;331200
	public asSSystemFunctionInterface* sysFuncIntf                                   // dwarf://field/./testdata/game.bz2;331216
	public asJITFunction jitFunction                                                 // dwarf://field/./testdata/game.bz2;331232
Methods:
	public asIScriptEngine*  GetEngine() const                                       // dwarf://method/./testdata/game.bz2;331248
	public int  AddRef() const                                                       // dwarf://method/./testdata/game.bz2;331281
	public int  Release() const                                                      // dwarf://method/./testdata/game.bz2;331314
	public int  GetId() const                                                        // dwarf://method/./testdata/game.bz2;331347
	public asEFuncType  GetFuncType() const                                          // dwarf://method/./testdata/game.bz2;331380
	public const char*  GetModuleName() const                                        // dwarf://method/./testdata/game.bz2;331413
	public const char*  GetScriptSectionName() const                                 // dwarf://method/./testdata/game.bz2;331446
	public const char*  GetConfigGroup() const                                       // dwarf://method/./testdata/game.bz2;331479
	public asDWORD  GetAccessMask() const                                            // dwarf://method/./testdata/game.bz2;331512
	public asIObjectType*  GetObjectType() const                                     // dwarf://method/./testdata/game.bz2;331545
	public const char*  GetObjectName() const                                        // dwarf://method/./testdata/game.bz2;331578
	public const char*  GetName() const                                              // dwarf://method/./testdata/game.bz2;331611
	public const char*  GetNamespace() const                                         // dwarf://method/./testdata/game.bz2;331644
	public const char*  GetDeclaration(bool , bool ) const                           // dwarf://method/./testdata/game.bz2;331677
	public bool  IsReadOnly() const                                                  // dwarf://method/./testdata/game.bz2;331720
	public bool  IsPrivate() const                                                   // dwarf://method/./testdata/game.bz2;331753
	public bool  IsFinal() const                                                     // dwarf://method/./testdata/game.bz2;331786
	public bool  IsOverride() const                                                  // dwarf://method/./testdata/game.bz2;331819
	public bool  IsShared() const                                                    // dwarf://method/./testdata/game.bz2;331852
	public asUINT  GetParamCount() const                                             // dwarf://method/./testdata/game.bz2;331885
	public int  GetParamTypeId(asUINT , asDWORD* ) const                             // dwarf://method/./testdata/game.bz2;331918
	public int  GetReturnTypeId() const                                              // dwarf://method/./testdata/game.bz2;331961
	public int  GetTypeId() const                                                    // dwarf://method/./testdata/game.bz2;331994
	public bool  IsCompatibleWithTypeId(int ) const                                  // dwarf://method/./testdata/game.bz2;332027
	public asUINT  GetVarCount() const                                               // dwarf://method/./testdata/game.bz2;332065
	public int  GetVar(asUINT , const char** , int* ) const                          // dwarf://method/./testdata/game.bz2;332098
	public const char*  GetVarDecl(asUINT ) const                                    // dwarf://method/./testdata/game.bz2;332146
	public int  FindNextLineWithCode(int ) const                                     // dwarf://method/./testdata/game.bz2;332184
	public asDWORD*  GetByteCode(asUINT* )                                           // dwarf://method/./testdata/game.bz2;332222
	public void*  SetUserData(void* )                                                // dwarf://method/./testdata/game.bz2;332260
	public void*  GetUserData() const                                                // dwarf://method/./testdata/game.bz2;332298
	public void  asCScriptFunction(asCScriptEngine* , asCModule* , asEFuncType )     // dwarf://method/./testdata/game.bz2;332331
	public void  ~asCScriptFunction()                                                // dwarf://method/./testdata/game.bz2;332363
	public void  DestroyInternal()                                                   // dwarf://method/./testdata/game.bz2;332388
	public void  Orphan(asIScriptModule* )                                           // dwarf://method/./testdata/game.bz2;332409
	public void  AddVariable(& asCString , & asCDataType , int )                     // dwarf://method/./testdata/game.bz2;332435
	public int  GetSpaceNeededForArguments()                                         // dwarf://method/./testdata/game.bz2;332471
	public int  GetSpaceNeededForReturnValue()                                       // dwarf://method/./testdata/game.bz2;332496
	public asCString  GetDeclarationStr(bool , bool ) const                          // dwarf://method/./testdata/game.bz2;332521
	public int  GetLineNumber(int , int* )                                           // dwarf://method/./testdata/game.bz2;332556
	public void  ComputeSignatureId()                                                // dwarf://method/./testdata/game.bz2;332591
	public bool  IsSignatureEqual(const asCScriptFunction* ) const                   // dwarf://method/./testdata/game.bz2;332612
	public bool  IsSignatureExceptNameEqual(const asCScriptFunction* ) const         // dwarf://method/./testdata/game.bz2;332642
	public bool  IsSignatureExceptNameEqual(const & asCDataType , & asCArray<asCDataType> , & asCArray<asETypeModifiers> , const asCObjectType* , bool ) const // dwarf://method/./testdata/game.bz2;332672
	public bool  IsSignatureExceptNameAndReturnTypeEqual(const asCScriptFunction* ) const // dwarf://method/./testdata/game.bz2;332722
	public bool  IsSignatureExceptNameAndReturnTypeEqual(& asCArray<asCDataType> , & asCArray<asETypeModifiers> , const asCObjectType* , bool ) const // dwarf://method/./testdata/game.bz2;332752
	public bool  DoesReturnOnStack() const                                           // dwarf://method/./testdata/game.bz2;332797
	public void  JITCompile()                                                        // dwarf://method/./testdata/game.bz2;332822
	public void  AddReferences()                                                     // dwarf://method/./testdata/game.bz2;332843
	public void  ReleaseReferences()                                                 // dwarf://method/./testdata/game.bz2;332864
	public asCGlobalProperty*  GetPropertyByGlobalVarPtr(void* )                     // dwarf://method/./testdata/game.bz2;332885
	public int  GetRefCount()                                                        // dwarf://method/./testdata/game.bz2;332915
	public void  SetFlag()                                                           // dwarf://method/./testdata/game.bz2;332940
	public bool  GetFlag()                                                           // dwarf://method/./testdata/game.bz2;332961
	public void  EnumReferences(asIScriptEngine* )                                   // dwarf://method/./testdata/game.bz2;332986
	public void  ReleaseAllHandles(asIScriptEngine* )                                // dwarf://method/./testdata/game.bz2;333012
