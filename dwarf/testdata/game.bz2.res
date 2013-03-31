asCObjectType
Fields:
	// {"Name":{"Relative":"name","Absolute":"dwarf://field/./testdata/game.bz2;306871"},"Type":{"Name":{"Relative":"asCString"}},"Flags":128}
	public asCString name                                                            // dwarf://field/./testdata/game.bz2;306871
	// {"Name":{"Relative":"nameSpace","Absolute":"dwarf://field/./testdata/game.bz2;306886"},"Type":{"Name":{},"Specialization":[{"Name":{"Relative":"asSNameSpace"}}],"Flags":2560},"Flags":128}
	public asSNameSpace* nameSpace                                                   // dwarf://field/./testdata/game.bz2;306886
	// {"Name":{"Relative":"size","Absolute":"dwarf://field/./testdata/game.bz2;306901"},"Type":{"Name":{"Relative":"int"}},"Flags":128}
	public int size                                                                  // dwarf://field/./testdata/game.bz2;306901
	// {"Name":{"Relative":"properties","Absolute":"dwarf://field/./testdata/game.bz2;306916"},"Type":{"Name":{"Relative":"asCArray\u003casCObjectProperty *\u003e"}},"Flags":128}
	public asCArray<asCObjectProperty *> properties                                  // dwarf://field/./testdata/game.bz2;306916
	// {"Name":{"Relative":"methods","Absolute":"dwarf://field/./testdata/game.bz2;306931"},"Type":{"Name":{"Relative":"asCArray\u003cint\u003e"}},"Flags":128}
	public asCArray<int> methods                                                     // dwarf://field/./testdata/game.bz2;306931
	// {"Name":{"Relative":"interfaces","Absolute":"dwarf://field/./testdata/game.bz2;306946"},"Type":{"Name":{"Relative":"asCArray\u003casCObjectType *\u003e"}},"Flags":128}
	public asCArray<asCObjectType *> interfaces                                      // dwarf://field/./testdata/game.bz2;306946
	// {"Name":{"Relative":"enumValues","Absolute":"dwarf://field/./testdata/game.bz2;306961"},"Type":{"Name":{"Relative":"asCArray\u003casSEnumValue *\u003e"}},"Flags":128}
	public asCArray<asSEnumValue *> enumValues                                       // dwarf://field/./testdata/game.bz2;306961
	// {"Name":{"Relative":"derivedFrom","Absolute":"dwarf://field/./testdata/game.bz2;306977"},"Type":{"Name":{},"Specialization":[{"Name":{"Relative":"asCObjectType"}}],"Flags":2560},"Flags":128}
	public asCObjectType* derivedFrom                                                // dwarf://field/./testdata/game.bz2;306977
	// {"Name":{"Relative":"virtualFunctionTable","Absolute":"dwarf://field/./testdata/game.bz2;306993"},"Type":{"Name":{"Relative":"asCArray\u003casCScriptFunction *\u003e"}},"Flags":128}
	public asCArray<asCScriptFunction *> virtualFunctionTable                        // dwarf://field/./testdata/game.bz2;306993
	// {"Name":{"Relative":"flags","Absolute":"dwarf://field/./testdata/game.bz2;307009"},"Type":{"Name":{"Relative":"asDWORD"}},"Flags":128}
	public asDWORD flags                                                             // dwarf://field/./testdata/game.bz2;307009
	// {"Name":{"Relative":"accessMask","Absolute":"dwarf://field/./testdata/game.bz2;307025"},"Type":{"Name":{"Relative":"asDWORD"}},"Flags":128}
	public asDWORD accessMask                                                        // dwarf://field/./testdata/game.bz2;307025
	// {"Name":{"Relative":"beh","Absolute":"dwarf://field/./testdata/game.bz2;307041"},"Type":{"Name":{"Relative":"asSTypeBehaviour"}},"Flags":128}
	public asSTypeBehaviour beh                                                      // dwarf://field/./testdata/game.bz2;307041
	// {"Name":{"Relative":"templateSubTypes","Absolute":"dwarf://field/./testdata/game.bz2;307057"},"Type":{"Name":{"Relative":"asCArray\u003casCDataType\u003e"}},"Flags":128}
	public asCArray<asCDataType> templateSubTypes                                    // dwarf://field/./testdata/game.bz2;307057
	// {"Name":{"Relative":"acceptValueSubType","Absolute":"dwarf://field/./testdata/game.bz2;307073"},"Type":{"Name":{"Relative":"bool"}},"Flags":128}
	public bool acceptValueSubType                                                   // dwarf://field/./testdata/game.bz2;307073
	// {"Name":{"Relative":"acceptRefSubType","Absolute":"dwarf://field/./testdata/game.bz2;307089"},"Type":{"Name":{"Relative":"bool"}},"Flags":128}
	public bool acceptRefSubType                                                     // dwarf://field/./testdata/game.bz2;307089
	// {"Name":{"Relative":"engine","Absolute":"dwarf://field/./testdata/game.bz2;307105"},"Type":{"Name":{},"Specialization":[{"Name":{"Relative":"asCScriptEngine"}}],"Flags":2560},"Flags":128}
	public asCScriptEngine* engine                                                   // dwarf://field/./testdata/game.bz2;307105
	// {"Name":{"Relative":"module","Absolute":"dwarf://field/./testdata/game.bz2;307121"},"Type":{"Name":{},"Specialization":[{"Name":{"Relative":"asCModule"}}],"Flags":2560},"Flags":128}
	public asCModule* module                                                         // dwarf://field/./testdata/game.bz2;307121
	// {"Name":{"Relative":"userData","Absolute":"dwarf://field/./testdata/game.bz2;307137"},"Type":{"Name":{"Relative":"asCArray\u003cunsigned long\u003e"}},"Flags":128}
	public asCArray<unsigned long> userData                                          // dwarf://field/./testdata/game.bz2;307137
	// {"Name":{"Relative":"refCount","Absolute":"dwarf://field/./testdata/game.bz2;307153"},"Type":{"Name":{"Relative":"asCAtomic"}},"Flags":384}
	protected asCAtomic refCount                                                     // dwarf://field/./testdata/game.bz2;307153
	// {"Name":{"Relative":"gcFlag","Absolute":"dwarf://field/./testdata/game.bz2;307169"},"Type":{"Name":{"Relative":"bool"}},"Flags":384}
	protected bool gcFlag                                                            // dwarf://field/./testdata/game.bz2;307169
Methods:
	// {"Name":{"Relative":"GetEngine","Absolute":"dwarf://method/./testdata/game.bz2;307185"},"Flags":136,"Returns":[{"Name":{},"Type":{"Name":{},"Specialization":[{"Name":{"Relative":"asIScriptEngine"}}],"Flags":2560}}]}
	public asIScriptEngine*  GetEngine() const                                       // dwarf://method/./testdata/game.bz2;307185
	// {"Name":{"Relative":"GetConfigGroup","Absolute":"dwarf://method/./testdata/game.bz2;307218"},"Flags":136,"Returns":[{"Name":{},"Type":{"Name":{},"Specialization":[{"Name":{"Relative":"char"},"Flags":8}],"Flags":2560}}]}
	public const char*  GetConfigGroup() const                                       // dwarf://method/./testdata/game.bz2;307218
	// {"Name":{"Relative":"GetAccessMask","Absolute":"dwarf://method/./testdata/game.bz2;307251"},"Flags":136,"Returns":[{"Name":{},"Type":{"Name":{"Relative":"asDWORD"}}}]}
	public asDWORD  GetAccessMask() const                                            // dwarf://method/./testdata/game.bz2;307251
	// {"Name":{"Relative":"AddRef","Absolute":"dwarf://method/./testdata/game.bz2;307284"},"Flags":136,"Returns":[{"Name":{},"Type":{"Name":{"Relative":"int"}}}]}
	public int  AddRef() const                                                       // dwarf://method/./testdata/game.bz2;307284
	// {"Name":{"Relative":"Release","Absolute":"dwarf://method/./testdata/game.bz2;307317"},"Flags":136,"Returns":[{"Name":{},"Type":{"Name":{"Relative":"int"}}}]}
	public int  Release() const                                                      // dwarf://method/./testdata/game.bz2;307317
	// {"Name":{"Relative":"GetName","Absolute":"dwarf://method/./testdata/game.bz2;307350"},"Flags":136,"Returns":[{"Name":{},"Type":{"Name":{},"Specialization":[{"Name":{"Relative":"char"},"Flags":8}],"Flags":2560}}]}
	public const char*  GetName() const                                              // dwarf://method/./testdata/game.bz2;307350
	// {"Name":{"Relative":"GetNamespace","Absolute":"dwarf://method/./testdata/game.bz2;307383"},"Flags":136,"Returns":[{"Name":{},"Type":{"Name":{},"Specialization":[{"Name":{"Relative":"char"},"Flags":8}],"Flags":2560}}]}
	public const char*  GetNamespace() const                                         // dwarf://method/./testdata/game.bz2;307383
	// {"Name":{"Relative":"GetBaseType","Absolute":"dwarf://method/./testdata/game.bz2;307416"},"Flags":136,"Returns":[{"Name":{},"Type":{"Name":{},"Specialization":[{"Name":{"Relative":"asIObjectType"}}],"Flags":2560}}]}
	public asIObjectType*  GetBaseType() const                                       // dwarf://method/./testdata/game.bz2;307416
	// {"Name":{"Relative":"DerivesFrom","Absolute":"dwarf://method/./testdata/game.bz2;307449"},"Flags":136,"Returns":[{"Name":{},"Type":{"Name":{"Relative":"bool"}}}],"Parameters":[{"Name":{},"Type":{"Name":{},"Specialization":[{"Name":{"Relative":"asIObjectType"},"Flags":8}],"Flags":2560}}]}
	public bool  DerivesFrom(const asIObjectType* ) const                            // dwarf://method/./testdata/game.bz2;307449
	// {"Name":{"Relative":"GetFlags","Absolute":"dwarf://method/./testdata/game.bz2;307487"},"Flags":136,"Returns":[{"Name":{},"Type":{"Name":{"Relative":"asDWORD"}}}]}
	public asDWORD  GetFlags() const                                                 // dwarf://method/./testdata/game.bz2;307487
	// {"Name":{"Relative":"GetSize","Absolute":"dwarf://method/./testdata/game.bz2;307520"},"Flags":136,"Returns":[{"Name":{},"Type":{"Name":{"Relative":"asUINT"}}}]}
	public asUINT  GetSize() const                                                   // dwarf://method/./testdata/game.bz2;307520
	// {"Name":{"Relative":"GetTypeId","Absolute":"dwarf://method/./testdata/game.bz2;307553"},"Flags":136,"Returns":[{"Name":{},"Type":{"Name":{"Relative":"int"}}}]}
	public int  GetTypeId() const                                                    // dwarf://method/./testdata/game.bz2;307553
	// {"Name":{"Relative":"GetSubTypeId","Absolute":"dwarf://method/./testdata/game.bz2;307586"},"Flags":136,"Returns":[{"Name":{},"Type":{"Name":{"Relative":"int"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"asUINT"}}}]}
	public int  GetSubTypeId(asUINT ) const                                          // dwarf://method/./testdata/game.bz2;307586
	// {"Name":{"Relative":"GetSubType","Absolute":"dwarf://method/./testdata/game.bz2;307624"},"Flags":136,"Returns":[{"Name":{},"Type":{"Name":{},"Specialization":[{"Name":{"Relative":"asIObjectType"}}],"Flags":2560}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"asUINT"}}}]}
	public asIObjectType*  GetSubType(asUINT ) const                                 // dwarf://method/./testdata/game.bz2;307624
	// {"Name":{"Relative":"GetSubTypeCount","Absolute":"dwarf://method/./testdata/game.bz2;307662"},"Flags":136,"Returns":[{"Name":{},"Type":{"Name":{"Relative":"asUINT"}}}]}
	public asUINT  GetSubTypeCount() const                                           // dwarf://method/./testdata/game.bz2;307662
	// {"Name":{"Relative":"GetInterfaceCount","Absolute":"dwarf://method/./testdata/game.bz2;307695"},"Flags":136,"Returns":[{"Name":{},"Type":{"Name":{"Relative":"asUINT"}}}]}
	public asUINT  GetInterfaceCount() const                                         // dwarf://method/./testdata/game.bz2;307695
	// {"Name":{"Relative":"GetInterface","Absolute":"dwarf://method/./testdata/game.bz2;307728"},"Flags":136,"Returns":[{"Name":{},"Type":{"Name":{},"Specialization":[{"Name":{"Relative":"asIObjectType"}}],"Flags":2560}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"asUINT"}}}]}
	public asIObjectType*  GetInterface(asUINT ) const                               // dwarf://method/./testdata/game.bz2;307728
	// {"Name":{"Relative":"Implements","Absolute":"dwarf://method/./testdata/game.bz2;307766"},"Flags":136,"Returns":[{"Name":{},"Type":{"Name":{"Relative":"bool"}}}],"Parameters":[{"Name":{},"Type":{"Name":{},"Specialization":[{"Name":{"Relative":"asIObjectType"},"Flags":8}],"Flags":2560}}]}
	public bool  Implements(const asIObjectType* ) const                             // dwarf://method/./testdata/game.bz2;307766
	// {"Name":{"Relative":"GetFactoryCount","Absolute":"dwarf://method/./testdata/game.bz2;307804"},"Flags":136,"Returns":[{"Name":{},"Type":{"Name":{"Relative":"asUINT"}}}]}
	public asUINT  GetFactoryCount() const                                           // dwarf://method/./testdata/game.bz2;307804
	// {"Name":{"Relative":"GetFactoryByIndex","Absolute":"dwarf://method/./testdata/game.bz2;307837"},"Flags":136,"Returns":[{"Name":{},"Type":{"Name":{},"Specialization":[{"Name":{"Relative":"asIScriptFunction"}}],"Flags":2560}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"asUINT"}}}]}
	public asIScriptFunction*  GetFactoryByIndex(asUINT ) const                      // dwarf://method/./testdata/game.bz2;307837
	// {"Name":{"Relative":"GetFactoryByDecl","Absolute":"dwarf://method/./testdata/game.bz2;307875"},"Flags":136,"Returns":[{"Name":{},"Type":{"Name":{},"Specialization":[{"Name":{"Relative":"asIScriptFunction"}}],"Flags":2560}}],"Parameters":[{"Name":{},"Type":{"Name":{},"Specialization":[{"Name":{"Relative":"char"},"Flags":8}],"Flags":2560}}]}
	public asIScriptFunction*  GetFactoryByDecl(const char* ) const                  // dwarf://method/./testdata/game.bz2;307875
	// {"Name":{"Relative":"GetMethodCount","Absolute":"dwarf://method/./testdata/game.bz2;307913"},"Flags":136,"Returns":[{"Name":{},"Type":{"Name":{"Relative":"asUINT"}}}]}
	public asUINT  GetMethodCount() const                                            // dwarf://method/./testdata/game.bz2;307913
	// {"Name":{"Relative":"GetMethodByIndex","Absolute":"dwarf://method/./testdata/game.bz2;307946"},"Flags":136,"Returns":[{"Name":{},"Type":{"Name":{},"Specialization":[{"Name":{"Relative":"asIScriptFunction"}}],"Flags":2560}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"asUINT"}}},{"Name":{},"Type":{"Name":{"Relative":"bool"}}}]}
	public asIScriptFunction*  GetMethodByIndex(asUINT , bool ) const                // dwarf://method/./testdata/game.bz2;307946
	// {"Name":{"Relative":"GetMethodByName","Absolute":"dwarf://method/./testdata/game.bz2;307989"},"Flags":136,"Returns":[{"Name":{},"Type":{"Name":{},"Specialization":[{"Name":{"Relative":"asIScriptFunction"}}],"Flags":2560}}],"Parameters":[{"Name":{},"Type":{"Name":{},"Specialization":[{"Name":{"Relative":"char"},"Flags":8}],"Flags":2560}},{"Name":{},"Type":{"Name":{"Relative":"bool"}}}]}
	public asIScriptFunction*  GetMethodByName(const char* , bool ) const            // dwarf://method/./testdata/game.bz2;307989
	// {"Name":{"Relative":"GetMethodByDecl","Absolute":"dwarf://method/./testdata/game.bz2;308032"},"Flags":136,"Returns":[{"Name":{},"Type":{"Name":{},"Specialization":[{"Name":{"Relative":"asIScriptFunction"}}],"Flags":2560}}],"Parameters":[{"Name":{},"Type":{"Name":{},"Specialization":[{"Name":{"Relative":"char"},"Flags":8}],"Flags":2560}},{"Name":{},"Type":{"Name":{"Relative":"bool"}}}]}
	public asIScriptFunction*  GetMethodByDecl(const char* , bool ) const            // dwarf://method/./testdata/game.bz2;308032
	// {"Name":{"Relative":"GetPropertyCount","Absolute":"dwarf://method/./testdata/game.bz2;308075"},"Flags":136,"Returns":[{"Name":{},"Type":{"Name":{"Relative":"asUINT"}}}]}
	public asUINT  GetPropertyCount() const                                          // dwarf://method/./testdata/game.bz2;308075
	// {"Name":{"Relative":"GetProperty","Absolute":"dwarf://method/./testdata/game.bz2;308108"},"Flags":136,"Returns":[{"Name":{},"Type":{"Name":{"Relative":"int"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"asUINT"}}},{"Name":{},"Type":{"Name":{},"Specialization":[{"Name":{},"Specialization":[{"Name":{"Relative":"char"},"Flags":8}],"Flags":2560}],"Flags":2560}},{"Name":{},"Type":{"Name":{},"Specialization":[{"Name":{"Relative":"int"}}],"Flags":2560}},{"Name":{},"Type":{"Name":{},"Specialization":[{"Name":{"Relative":"bool"}}],"Flags":2560}},{"Name":{},"Type":{"Name":{},"Specialization":[{"Name":{"Relative":"int"}}],"Flags":2560}},{"Name":{},"Type":{"Name":{},"Specialization":[{"Name":{"Relative":"bool"}}],"Flags":2560}},{"Name":{},"Type":{"Name":{},"Specialization":[{"Name":{"Relative":"asDWORD"}}],"Flags":2560}}]}
	public int  GetProperty(asUINT , const char** , int* , bool* , int* , bool* , asDWORD* ) const // dwarf://method/./testdata/game.bz2;308108
	// {"Name":{"Relative":"GetPropertyDeclaration","Absolute":"dwarf://method/./testdata/game.bz2;308176"},"Flags":136,"Returns":[{"Name":{},"Type":{"Name":{},"Specialization":[{"Name":{"Relative":"char"},"Flags":8}],"Flags":2560}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"asUINT"}}}]}
	public const char*  GetPropertyDeclaration(asUINT ) const                        // dwarf://method/./testdata/game.bz2;308176
	// {"Name":{"Relative":"GetBehaviourCount","Absolute":"dwarf://method/./testdata/game.bz2;308214"},"Flags":136,"Returns":[{"Name":{},"Type":{"Name":{"Relative":"asUINT"}}}]}
	public asUINT  GetBehaviourCount() const                                         // dwarf://method/./testdata/game.bz2;308214
	// {"Name":{"Relative":"GetBehaviourByIndex","Absolute":"dwarf://method/./testdata/game.bz2;308247"},"Flags":136,"Returns":[{"Name":{},"Type":{"Name":{},"Specialization":[{"Name":{"Relative":"asIScriptFunction"}}],"Flags":2560}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"asUINT"}}},{"Name":{},"Type":{"Name":{},"Specialization":[{"Name":{"Relative":"asEBehaviours"}}],"Flags":2560}}]}
	public asIScriptFunction*  GetBehaviourByIndex(asUINT , asEBehaviours* ) const   // dwarf://method/./testdata/game.bz2;308247
	// {"Name":{"Relative":"SetUserData","Absolute":"dwarf://method/./testdata/game.bz2;308290"},"Flags":128,"Returns":[{"Name":{},"Type":{"Name":{},"Specialization":[{"Name":{"Relative":"void"}}],"Flags":2560}}],"Parameters":[{"Name":{},"Type":{"Name":{},"Specialization":[{"Name":{"Relative":"void"}}],"Flags":2560}},{"Name":{},"Type":{"Name":{"Relative":"asPWORD"}}}]}
	public void*  SetUserData(void* , asPWORD )                                      // dwarf://method/./testdata/game.bz2;308290
	// {"Name":{"Relative":"GetUserData","Absolute":"dwarf://method/./testdata/game.bz2;308333"},"Flags":136,"Returns":[{"Name":{},"Type":{"Name":{},"Specialization":[{"Name":{"Relative":"void"}}],"Flags":2560}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"asPWORD"}}}]}
	public void*  GetUserData(asPWORD ) const                                        // dwarf://method/./testdata/game.bz2;308333
	// {"Name":{"Relative":"asCObjectType","Absolute":"dwarf://method/./testdata/game.bz2;308371"},"Flags":128,"Returns":[{"Name":{},"Type":{"Name":{"Relative":"void"}}}],"Parameters":[{"Name":{},"Type":{"Name":{},"Specialization":[{"Name":{"Relative":"asCScriptEngine"}}],"Flags":2560}}]}
	public void  asCObjectType(asCScriptEngine* )                                    // dwarf://method/./testdata/game.bz2;308371
	// {"Name":{"Relative":"~asCObjectType","Absolute":"dwarf://method/./testdata/game.bz2;308393"},"Flags":128,"Returns":[{"Name":{},"Type":{"Name":{"Relative":"void"}}}]}
	public void  ~asCObjectType()                                                    // dwarf://method/./testdata/game.bz2;308393
	// {"Name":{"Relative":"Orphan","Absolute":"dwarf://method/./testdata/game.bz2;308418"},"Flags":128,"Returns":[{"Name":{},"Type":{"Name":{"Relative":"void"}}}],"Parameters":[{"Name":{},"Type":{"Name":{},"Specialization":[{"Name":{"Relative":"asCModule"}}],"Flags":2560}}]}
	public void  Orphan(asCModule* )                                                 // dwarf://method/./testdata/game.bz2;308418
	// {"Name":{"Relative":"GetRefCount","Absolute":"dwarf://method/./testdata/game.bz2;308444"},"Flags":128,"Returns":[{"Name":{},"Type":{"Name":{"Relative":"int"}}}]}
	public int  GetRefCount()                                                        // dwarf://method/./testdata/game.bz2;308444
	// {"Name":{"Relative":"SetGCFlag","Absolute":"dwarf://method/./testdata/game.bz2;308469"},"Flags":128,"Returns":[{"Name":{},"Type":{"Name":{"Relative":"void"}}}]}
	public void  SetGCFlag()                                                         // dwarf://method/./testdata/game.bz2;308469
	// {"Name":{"Relative":"GetGCFlag","Absolute":"dwarf://method/./testdata/game.bz2;308490"},"Flags":128,"Returns":[{"Name":{},"Type":{"Name":{"Relative":"bool"}}}]}
	public bool  GetGCFlag()                                                         // dwarf://method/./testdata/game.bz2;308490
	// {"Name":{"Relative":"EnumReferences","Absolute":"dwarf://method/./testdata/game.bz2;308515"},"Flags":128,"Returns":[{"Name":{},"Type":{"Name":{"Relative":"void"}}}],"Parameters":[{"Name":{},"Type":{"Name":{},"Specialization":[{"Name":{"Relative":"asIScriptEngine"}}],"Flags":2560}}]}
	public void  EnumReferences(asIScriptEngine* )                                   // dwarf://method/./testdata/game.bz2;308515
	// {"Name":{"Relative":"ReleaseAllHandles","Absolute":"dwarf://method/./testdata/game.bz2;308541"},"Flags":128,"Returns":[{"Name":{},"Type":{"Name":{"Relative":"void"}}}],"Parameters":[{"Name":{},"Type":{"Name":{},"Specialization":[{"Name":{"Relative":"asIScriptEngine"}}],"Flags":2560}}]}
	public void  ReleaseAllHandles(asIScriptEngine* )                                // dwarf://method/./testdata/game.bz2;308541
	// {"Name":{"Relative":"ReleaseAllFunctions","Absolute":"dwarf://method/./testdata/game.bz2;308567"},"Flags":128,"Returns":[{"Name":{},"Type":{"Name":{"Relative":"void"}}}]}
	public void  ReleaseAllFunctions()                                               // dwarf://method/./testdata/game.bz2;308567
	// {"Name":{"Relative":"IsInterface","Absolute":"dwarf://method/./testdata/game.bz2;308588"},"Flags":136,"Returns":[{"Name":{},"Type":{"Name":{"Relative":"bool"}}}]}
	public bool  IsInterface() const                                                 // dwarf://method/./testdata/game.bz2;308588
	// {"Name":{"Relative":"IsShared","Absolute":"dwarf://method/./testdata/game.bz2;308613"},"Flags":136,"Returns":[{"Name":{},"Type":{"Name":{"Relative":"bool"}}}]}
	public bool  IsShared() const                                                    // dwarf://method/./testdata/game.bz2;308613
	// {"Name":{"Relative":"AddPropertyToClass","Absolute":"dwarf://method/./testdata/game.bz2;308638"},"Flags":128,"Returns":[{"Name":{},"Type":{"Name":{},"Specialization":[{"Name":{"Relative":"asCObjectProperty"}}],"Flags":2560}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"asCString"},"Flags":40}},{"Name":{},"Type":{"Name":{"Relative":"asCDataType"},"Flags":40}},{"Name":{},"Type":{"Name":{"Relative":"bool"}}}]}
	public asCObjectProperty*  AddPropertyToClass(const & asCString , const & asCDataType , bool ) // dwarf://method/./testdata/game.bz2;308638
	// {"Name":{"Relative":"ReleaseAllProperties","Absolute":"dwarf://method/./testdata/game.bz2;308678"},"Flags":128,"Returns":[{"Name":{},"Type":{"Name":{"Relative":"void"}}}]}
	public void  ReleaseAllProperties()                                              // dwarf://method/./testdata/game.bz2;308678
	// {"Name":{"Relative":"asCObjectType","Absolute":"dwarf://method/./testdata/game.bz2;308699"},"Flags":384,"Returns":[{"Name":{},"Type":{"Name":{"Relative":"void"}}}]}
	protected void  asCObjectType()                                                  // dwarf://method/./testdata/game.bz2;308699
asCScriptFunction
Fields:
	// {"Name":{"Relative":"refCount","Absolute":"dwarf://field/./testdata/game.bz2;330664"},"Type":{"Name":{"Relative":"asCAtomic"}},"Flags":128}
	public asCAtomic refCount                                                        // dwarf://field/./testdata/game.bz2;330664
	// {"Name":{"Relative":"gcFlag","Absolute":"dwarf://field/./testdata/game.bz2;330679"},"Type":{"Name":{"Relative":"bool"}},"Flags":128}
	public bool gcFlag                                                               // dwarf://field/./testdata/game.bz2;330679
	// {"Name":{"Relative":"engine","Absolute":"dwarf://field/./testdata/game.bz2;330694"},"Type":{"Name":{},"Specialization":[{"Name":{"Relative":"asCScriptEngine"}}],"Flags":2560},"Flags":128}
	public asCScriptEngine* engine                                                   // dwarf://field/./testdata/game.bz2;330694
	// {"Name":{"Relative":"module","Absolute":"dwarf://field/./testdata/game.bz2;330709"},"Type":{"Name":{},"Specialization":[{"Name":{"Relative":"asCModule"}}],"Flags":2560},"Flags":128}
	public asCModule* module                                                         // dwarf://field/./testdata/game.bz2;330709
	// {"Name":{"Relative":"userData","Absolute":"dwarf://field/./testdata/game.bz2;330724"},"Type":{"Name":{},"Specialization":[{"Name":{"Relative":"void"}}],"Flags":2560},"Flags":128}
	public void* userData                                                            // dwarf://field/./testdata/game.bz2;330724
	// {"Name":{"Relative":"name","Absolute":"dwarf://field/./testdata/game.bz2;330739"},"Type":{"Name":{"Relative":"asCString"}},"Flags":128}
	public asCString name                                                            // dwarf://field/./testdata/game.bz2;330739
	// {"Name":{"Relative":"returnType","Absolute":"dwarf://field/./testdata/game.bz2;330754"},"Type":{"Name":{"Relative":"asCDataType"}},"Flags":128}
	public asCDataType returnType                                                    // dwarf://field/./testdata/game.bz2;330754
	// {"Name":{"Relative":"parameterTypes","Absolute":"dwarf://field/./testdata/game.bz2;330769"},"Type":{"Name":{"Relative":"asCArray\u003casCDataType\u003e"}},"Flags":128}
	public asCArray<asCDataType> parameterTypes                                      // dwarf://field/./testdata/game.bz2;330769
	// {"Name":{"Relative":"inOutFlags","Absolute":"dwarf://field/./testdata/game.bz2;330784"},"Type":{"Name":{"Relative":"asCArray\u003casETypeModifiers\u003e"}},"Flags":128}
	public asCArray<asETypeModifiers> inOutFlags                                     // dwarf://field/./testdata/game.bz2;330784
	// {"Name":{"Relative":"defaultArgs","Absolute":"dwarf://field/./testdata/game.bz2;330800"},"Type":{"Name":{"Relative":"asCArray\u003casCString *\u003e"}},"Flags":128}
	public asCArray<asCString *> defaultArgs                                         // dwarf://field/./testdata/game.bz2;330800
	// {"Name":{"Relative":"isReadOnly","Absolute":"dwarf://field/./testdata/game.bz2;330816"},"Type":{"Name":{"Relative":"bool"}},"Flags":128}
	public bool isReadOnly                                                           // dwarf://field/./testdata/game.bz2;330816
	// {"Name":{"Relative":"isPrivate","Absolute":"dwarf://field/./testdata/game.bz2;330832"},"Type":{"Name":{"Relative":"bool"}},"Flags":128}
	public bool isPrivate                                                            // dwarf://field/./testdata/game.bz2;330832
	// {"Name":{"Relative":"isFinal","Absolute":"dwarf://field/./testdata/game.bz2;330848"},"Type":{"Name":{"Relative":"bool"}},"Flags":128}
	public bool isFinal                                                              // dwarf://field/./testdata/game.bz2;330848
	// {"Name":{"Relative":"isOverride","Absolute":"dwarf://field/./testdata/game.bz2;330864"},"Type":{"Name":{"Relative":"bool"}},"Flags":128}
	public bool isOverride                                                           // dwarf://field/./testdata/game.bz2;330864
	// {"Name":{"Relative":"objectType","Absolute":"dwarf://field/./testdata/game.bz2;330880"},"Type":{"Name":{},"Specialization":[{"Name":{"Relative":"asCObjectType"}}],"Flags":2560},"Flags":128}
	public asCObjectType* objectType                                                 // dwarf://field/./testdata/game.bz2;330880
	// {"Name":{"Relative":"signatureId","Absolute":"dwarf://field/./testdata/game.bz2;330896"},"Type":{"Name":{"Relative":"int"}},"Flags":128}
	public int signatureId                                                           // dwarf://field/./testdata/game.bz2;330896
	// {"Name":{"Relative":"id","Absolute":"dwarf://field/./testdata/game.bz2;330912"},"Type":{"Name":{"Relative":"int"}},"Flags":128}
	public int id                                                                    // dwarf://field/./testdata/game.bz2;330912
	// {"Name":{"Relative":"funcType","Absolute":"dwarf://field/./testdata/game.bz2;330928"},"Type":{"Name":{"Relative":"asEFuncType"}},"Flags":128}
	public asEFuncType funcType                                                      // dwarf://field/./testdata/game.bz2;330928
	// {"Name":{"Relative":"accessMask","Absolute":"dwarf://field/./testdata/game.bz2;330944"},"Type":{"Name":{"Relative":"asDWORD"}},"Flags":128}
	public asDWORD accessMask                                                        // dwarf://field/./testdata/game.bz2;330944
	// {"Name":{"Relative":"isShared","Absolute":"dwarf://field/./testdata/game.bz2;330960"},"Type":{"Name":{"Relative":"bool"}},"Flags":128}
	public bool isShared                                                             // dwarf://field/./testdata/game.bz2;330960
	// {"Name":{"Relative":"nameSpace","Absolute":"dwarf://field/./testdata/game.bz2;330976"},"Type":{"Name":{},"Specialization":[{"Name":{"Relative":"asSNameSpace"}}],"Flags":2560},"Flags":128}
	public asSNameSpace* nameSpace                                                   // dwarf://field/./testdata/game.bz2;330976
	// {"Name":{"Relative":"byteCode","Absolute":"dwarf://field/./testdata/game.bz2;330992"},"Type":{"Name":{"Relative":"asCArray\u003cunsigned int\u003e"}},"Flags":128}
	public asCArray<unsigned int> byteCode                                           // dwarf://field/./testdata/game.bz2;330992
	// {"Name":{"Relative":"variableSpace","Absolute":"dwarf://field/./testdata/game.bz2;331008"},"Type":{"Name":{"Relative":"asDWORD"}},"Flags":128}
	public asDWORD variableSpace                                                     // dwarf://field/./testdata/game.bz2;331008
	// {"Name":{"Relative":"objVariableTypes","Absolute":"dwarf://field/./testdata/game.bz2;331024"},"Type":{"Name":{"Relative":"asCArray\u003casCObjectType *\u003e"}},"Flags":128}
	public asCArray<asCObjectType *> objVariableTypes                                // dwarf://field/./testdata/game.bz2;331024
	// {"Name":{"Relative":"funcVariableTypes","Absolute":"dwarf://field/./testdata/game.bz2;331040"},"Type":{"Name":{"Relative":"asCArray\u003casCScriptFunction *\u003e"}},"Flags":128}
	public asCArray<asCScriptFunction *> funcVariableTypes                           // dwarf://field/./testdata/game.bz2;331040
	// {"Name":{"Relative":"objVariablePos","Absolute":"dwarf://field/./testdata/game.bz2;331056"},"Type":{"Name":{"Relative":"asCArray\u003cint\u003e"}},"Flags":128}
	public asCArray<int> objVariablePos                                              // dwarf://field/./testdata/game.bz2;331056
	// {"Name":{"Relative":"objVariablesOnHeap","Absolute":"dwarf://field/./testdata/game.bz2;331072"},"Type":{"Name":{"Relative":"asUINT"}},"Flags":128}
	public asUINT objVariablesOnHeap                                                 // dwarf://field/./testdata/game.bz2;331072
	// {"Name":{"Relative":"objVariableInfo","Absolute":"dwarf://field/./testdata/game.bz2;331088"},"Type":{"Name":{"Relative":"asCArray\u003casSObjectVariableInfo\u003e"}},"Flags":128}
	public asCArray<asSObjectVariableInfo> objVariableInfo                           // dwarf://field/./testdata/game.bz2;331088
	// {"Name":{"Relative":"variables","Absolute":"dwarf://field/./testdata/game.bz2;331104"},"Type":{"Name":{"Relative":"asCArray\u003casSScriptVariable *\u003e"}},"Flags":128}
	public asCArray<asSScriptVariable *> variables                                   // dwarf://field/./testdata/game.bz2;331104
	// {"Name":{"Relative":"stackNeeded","Absolute":"dwarf://field/./testdata/game.bz2;331120"},"Type":{"Name":{"Relative":"int"}},"Flags":128}
	public int stackNeeded                                                           // dwarf://field/./testdata/game.bz2;331120
	// {"Name":{"Relative":"lineNumbers","Absolute":"dwarf://field/./testdata/game.bz2;331136"},"Type":{"Name":{"Relative":"asCArray\u003cint\u003e"}},"Flags":128}
	public asCArray<int> lineNumbers                                                 // dwarf://field/./testdata/game.bz2;331136
	// {"Name":{"Relative":"scriptSectionIdx","Absolute":"dwarf://field/./testdata/game.bz2;331152"},"Type":{"Name":{"Relative":"int"}},"Flags":128}
	public int scriptSectionIdx                                                      // dwarf://field/./testdata/game.bz2;331152
	// {"Name":{"Relative":"sectionIdxs","Absolute":"dwarf://field/./testdata/game.bz2;331168"},"Type":{"Name":{"Relative":"asCArray\u003cint\u003e"}},"Flags":128}
	public asCArray<int> sectionIdxs                                                 // dwarf://field/./testdata/game.bz2;331168
	// {"Name":{"Relative":"dontCleanUpOnException","Absolute":"dwarf://field/./testdata/game.bz2;331184"},"Type":{"Name":{"Relative":"bool"}},"Flags":128}
	public bool dontCleanUpOnException                                               // dwarf://field/./testdata/game.bz2;331184
	// {"Name":{"Relative":"vfTableIdx","Absolute":"dwarf://field/./testdata/game.bz2;331200"},"Type":{"Name":{"Relative":"int"}},"Flags":128}
	public int vfTableIdx                                                            // dwarf://field/./testdata/game.bz2;331200
	// {"Name":{"Relative":"sysFuncIntf","Absolute":"dwarf://field/./testdata/game.bz2;331216"},"Type":{"Name":{},"Specialization":[{"Name":{"Relative":"asSSystemFunctionInterface"}}],"Flags":2560},"Flags":128}
	public asSSystemFunctionInterface* sysFuncIntf                                   // dwarf://field/./testdata/game.bz2;331216
	// {"Name":{"Relative":"jitFunction","Absolute":"dwarf://field/./testdata/game.bz2;331232"},"Type":{"Name":{"Relative":"asJITFunction"}},"Flags":128}
	public asJITFunction jitFunction                                                 // dwarf://field/./testdata/game.bz2;331232
Methods:
	// {"Name":{"Relative":"GetEngine","Absolute":"dwarf://method/./testdata/game.bz2;331248"},"Flags":136,"Returns":[{"Name":{},"Type":{"Name":{},"Specialization":[{"Name":{"Relative":"asIScriptEngine"}}],"Flags":2560}}]}
	public asIScriptEngine*  GetEngine() const                                       // dwarf://method/./testdata/game.bz2;331248
	// {"Name":{"Relative":"AddRef","Absolute":"dwarf://method/./testdata/game.bz2;331281"},"Flags":136,"Returns":[{"Name":{},"Type":{"Name":{"Relative":"int"}}}]}
	public int  AddRef() const                                                       // dwarf://method/./testdata/game.bz2;331281
	// {"Name":{"Relative":"Release","Absolute":"dwarf://method/./testdata/game.bz2;331314"},"Flags":136,"Returns":[{"Name":{},"Type":{"Name":{"Relative":"int"}}}]}
	public int  Release() const                                                      // dwarf://method/./testdata/game.bz2;331314
	// {"Name":{"Relative":"GetId","Absolute":"dwarf://method/./testdata/game.bz2;331347"},"Flags":136,"Returns":[{"Name":{},"Type":{"Name":{"Relative":"int"}}}]}
	public int  GetId() const                                                        // dwarf://method/./testdata/game.bz2;331347
	// {"Name":{"Relative":"GetFuncType","Absolute":"dwarf://method/./testdata/game.bz2;331380"},"Flags":136,"Returns":[{"Name":{},"Type":{"Name":{"Relative":"asEFuncType"}}}]}
	public asEFuncType  GetFuncType() const                                          // dwarf://method/./testdata/game.bz2;331380
	// {"Name":{"Relative":"GetModuleName","Absolute":"dwarf://method/./testdata/game.bz2;331413"},"Flags":136,"Returns":[{"Name":{},"Type":{"Name":{},"Specialization":[{"Name":{"Relative":"char"},"Flags":8}],"Flags":2560}}]}
	public const char*  GetModuleName() const                                        // dwarf://method/./testdata/game.bz2;331413
	// {"Name":{"Relative":"GetScriptSectionName","Absolute":"dwarf://method/./testdata/game.bz2;331446"},"Flags":136,"Returns":[{"Name":{},"Type":{"Name":{},"Specialization":[{"Name":{"Relative":"char"},"Flags":8}],"Flags":2560}}]}
	public const char*  GetScriptSectionName() const                                 // dwarf://method/./testdata/game.bz2;331446
	// {"Name":{"Relative":"GetConfigGroup","Absolute":"dwarf://method/./testdata/game.bz2;331479"},"Flags":136,"Returns":[{"Name":{},"Type":{"Name":{},"Specialization":[{"Name":{"Relative":"char"},"Flags":8}],"Flags":2560}}]}
	public const char*  GetConfigGroup() const                                       // dwarf://method/./testdata/game.bz2;331479
	// {"Name":{"Relative":"GetAccessMask","Absolute":"dwarf://method/./testdata/game.bz2;331512"},"Flags":136,"Returns":[{"Name":{},"Type":{"Name":{"Relative":"asDWORD"}}}]}
	public asDWORD  GetAccessMask() const                                            // dwarf://method/./testdata/game.bz2;331512
	// {"Name":{"Relative":"GetObjectType","Absolute":"dwarf://method/./testdata/game.bz2;331545"},"Flags":136,"Returns":[{"Name":{},"Type":{"Name":{},"Specialization":[{"Name":{"Relative":"asIObjectType"}}],"Flags":2560}}]}
	public asIObjectType*  GetObjectType() const                                     // dwarf://method/./testdata/game.bz2;331545
	// {"Name":{"Relative":"GetObjectName","Absolute":"dwarf://method/./testdata/game.bz2;331578"},"Flags":136,"Returns":[{"Name":{},"Type":{"Name":{},"Specialization":[{"Name":{"Relative":"char"},"Flags":8}],"Flags":2560}}]}
	public const char*  GetObjectName() const                                        // dwarf://method/./testdata/game.bz2;331578
	// {"Name":{"Relative":"GetName","Absolute":"dwarf://method/./testdata/game.bz2;331611"},"Flags":136,"Returns":[{"Name":{},"Type":{"Name":{},"Specialization":[{"Name":{"Relative":"char"},"Flags":8}],"Flags":2560}}]}
	public const char*  GetName() const                                              // dwarf://method/./testdata/game.bz2;331611
	// {"Name":{"Relative":"GetNamespace","Absolute":"dwarf://method/./testdata/game.bz2;331644"},"Flags":136,"Returns":[{"Name":{},"Type":{"Name":{},"Specialization":[{"Name":{"Relative":"char"},"Flags":8}],"Flags":2560}}]}
	public const char*  GetNamespace() const                                         // dwarf://method/./testdata/game.bz2;331644
	// {"Name":{"Relative":"GetDeclaration","Absolute":"dwarf://method/./testdata/game.bz2;331677"},"Flags":136,"Returns":[{"Name":{},"Type":{"Name":{},"Specialization":[{"Name":{"Relative":"char"},"Flags":8}],"Flags":2560}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"bool"}}},{"Name":{},"Type":{"Name":{"Relative":"bool"}}}]}
	public const char*  GetDeclaration(bool , bool ) const                           // dwarf://method/./testdata/game.bz2;331677
	// {"Name":{"Relative":"IsReadOnly","Absolute":"dwarf://method/./testdata/game.bz2;331720"},"Flags":136,"Returns":[{"Name":{},"Type":{"Name":{"Relative":"bool"}}}]}
	public bool  IsReadOnly() const                                                  // dwarf://method/./testdata/game.bz2;331720
	// {"Name":{"Relative":"IsPrivate","Absolute":"dwarf://method/./testdata/game.bz2;331753"},"Flags":136,"Returns":[{"Name":{},"Type":{"Name":{"Relative":"bool"}}}]}
	public bool  IsPrivate() const                                                   // dwarf://method/./testdata/game.bz2;331753
	// {"Name":{"Relative":"IsFinal","Absolute":"dwarf://method/./testdata/game.bz2;331786"},"Flags":136,"Returns":[{"Name":{},"Type":{"Name":{"Relative":"bool"}}}]}
	public bool  IsFinal() const                                                     // dwarf://method/./testdata/game.bz2;331786
	// {"Name":{"Relative":"IsOverride","Absolute":"dwarf://method/./testdata/game.bz2;331819"},"Flags":136,"Returns":[{"Name":{},"Type":{"Name":{"Relative":"bool"}}}]}
	public bool  IsOverride() const                                                  // dwarf://method/./testdata/game.bz2;331819
	// {"Name":{"Relative":"IsShared","Absolute":"dwarf://method/./testdata/game.bz2;331852"},"Flags":136,"Returns":[{"Name":{},"Type":{"Name":{"Relative":"bool"}}}]}
	public bool  IsShared() const                                                    // dwarf://method/./testdata/game.bz2;331852
	// {"Name":{"Relative":"GetParamCount","Absolute":"dwarf://method/./testdata/game.bz2;331885"},"Flags":136,"Returns":[{"Name":{},"Type":{"Name":{"Relative":"asUINT"}}}]}
	public asUINT  GetParamCount() const                                             // dwarf://method/./testdata/game.bz2;331885
	// {"Name":{"Relative":"GetParamTypeId","Absolute":"dwarf://method/./testdata/game.bz2;331918"},"Flags":136,"Returns":[{"Name":{},"Type":{"Name":{"Relative":"int"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"asUINT"}}},{"Name":{},"Type":{"Name":{},"Specialization":[{"Name":{"Relative":"asDWORD"}}],"Flags":2560}}]}
	public int  GetParamTypeId(asUINT , asDWORD* ) const                             // dwarf://method/./testdata/game.bz2;331918
	// {"Name":{"Relative":"GetReturnTypeId","Absolute":"dwarf://method/./testdata/game.bz2;331961"},"Flags":136,"Returns":[{"Name":{},"Type":{"Name":{"Relative":"int"}}}]}
	public int  GetReturnTypeId() const                                              // dwarf://method/./testdata/game.bz2;331961
	// {"Name":{"Relative":"GetTypeId","Absolute":"dwarf://method/./testdata/game.bz2;331994"},"Flags":136,"Returns":[{"Name":{},"Type":{"Name":{"Relative":"int"}}}]}
	public int  GetTypeId() const                                                    // dwarf://method/./testdata/game.bz2;331994
	// {"Name":{"Relative":"IsCompatibleWithTypeId","Absolute":"dwarf://method/./testdata/game.bz2;332027"},"Flags":136,"Returns":[{"Name":{},"Type":{"Name":{"Relative":"bool"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"int"}}}]}
	public bool  IsCompatibleWithTypeId(int ) const                                  // dwarf://method/./testdata/game.bz2;332027
	// {"Name":{"Relative":"GetVarCount","Absolute":"dwarf://method/./testdata/game.bz2;332065"},"Flags":136,"Returns":[{"Name":{},"Type":{"Name":{"Relative":"asUINT"}}}]}
	public asUINT  GetVarCount() const                                               // dwarf://method/./testdata/game.bz2;332065
	// {"Name":{"Relative":"GetVar","Absolute":"dwarf://method/./testdata/game.bz2;332098"},"Flags":136,"Returns":[{"Name":{},"Type":{"Name":{"Relative":"int"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"asUINT"}}},{"Name":{},"Type":{"Name":{},"Specialization":[{"Name":{},"Specialization":[{"Name":{"Relative":"char"},"Flags":8}],"Flags":2560}],"Flags":2560}},{"Name":{},"Type":{"Name":{},"Specialization":[{"Name":{"Relative":"int"}}],"Flags":2560}}]}
	public int  GetVar(asUINT , const char** , int* ) const                          // dwarf://method/./testdata/game.bz2;332098
	// {"Name":{"Relative":"GetVarDecl","Absolute":"dwarf://method/./testdata/game.bz2;332146"},"Flags":136,"Returns":[{"Name":{},"Type":{"Name":{},"Specialization":[{"Name":{"Relative":"char"},"Flags":8}],"Flags":2560}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"asUINT"}}}]}
	public const char*  GetVarDecl(asUINT ) const                                    // dwarf://method/./testdata/game.bz2;332146
	// {"Name":{"Relative":"FindNextLineWithCode","Absolute":"dwarf://method/./testdata/game.bz2;332184"},"Flags":136,"Returns":[{"Name":{},"Type":{"Name":{"Relative":"int"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"int"}}}]}
	public int  FindNextLineWithCode(int ) const                                     // dwarf://method/./testdata/game.bz2;332184
	// {"Name":{"Relative":"GetByteCode","Absolute":"dwarf://method/./testdata/game.bz2;332222"},"Flags":128,"Returns":[{"Name":{},"Type":{"Name":{},"Specialization":[{"Name":{"Relative":"asDWORD"}}],"Flags":2560}}],"Parameters":[{"Name":{},"Type":{"Name":{},"Specialization":[{"Name":{"Relative":"asUINT"}}],"Flags":2560}}]}
	public asDWORD*  GetByteCode(asUINT* )                                           // dwarf://method/./testdata/game.bz2;332222
	// {"Name":{"Relative":"SetUserData","Absolute":"dwarf://method/./testdata/game.bz2;332260"},"Flags":128,"Returns":[{"Name":{},"Type":{"Name":{},"Specialization":[{"Name":{"Relative":"void"}}],"Flags":2560}}],"Parameters":[{"Name":{},"Type":{"Name":{},"Specialization":[{"Name":{"Relative":"void"}}],"Flags":2560}}]}
	public void*  SetUserData(void* )                                                // dwarf://method/./testdata/game.bz2;332260
	// {"Name":{"Relative":"GetUserData","Absolute":"dwarf://method/./testdata/game.bz2;332298"},"Flags":136,"Returns":[{"Name":{},"Type":{"Name":{},"Specialization":[{"Name":{"Relative":"void"}}],"Flags":2560}}]}
	public void*  GetUserData() const                                                // dwarf://method/./testdata/game.bz2;332298
	// {"Name":{"Relative":"asCScriptFunction","Absolute":"dwarf://method/./testdata/game.bz2;332331"},"Flags":128,"Returns":[{"Name":{},"Type":{"Name":{"Relative":"void"}}}],"Parameters":[{"Name":{},"Type":{"Name":{},"Specialization":[{"Name":{"Relative":"asCScriptEngine"}}],"Flags":2560}},{"Name":{},"Type":{"Name":{},"Specialization":[{"Name":{"Relative":"asCModule"}}],"Flags":2560}},{"Name":{},"Type":{"Name":{"Relative":"asEFuncType"}}}]}
	public void  asCScriptFunction(asCScriptEngine* , asCModule* , asEFuncType )     // dwarf://method/./testdata/game.bz2;332331
	// {"Name":{"Relative":"~asCScriptFunction","Absolute":"dwarf://method/./testdata/game.bz2;332363"},"Flags":128,"Returns":[{"Name":{},"Type":{"Name":{"Relative":"void"}}}]}
	public void  ~asCScriptFunction()                                                // dwarf://method/./testdata/game.bz2;332363
	// {"Name":{"Relative":"DestroyInternal","Absolute":"dwarf://method/./testdata/game.bz2;332388"},"Flags":128,"Returns":[{"Name":{},"Type":{"Name":{"Relative":"void"}}}]}
	public void  DestroyInternal()                                                   // dwarf://method/./testdata/game.bz2;332388
	// {"Name":{"Relative":"Orphan","Absolute":"dwarf://method/./testdata/game.bz2;332409"},"Flags":128,"Returns":[{"Name":{},"Type":{"Name":{"Relative":"void"}}}],"Parameters":[{"Name":{},"Type":{"Name":{},"Specialization":[{"Name":{"Relative":"asIScriptModule"}}],"Flags":2560}}]}
	public void  Orphan(asIScriptModule* )                                           // dwarf://method/./testdata/game.bz2;332409
	// {"Name":{"Relative":"AddVariable","Absolute":"dwarf://method/./testdata/game.bz2;332435"},"Flags":128,"Returns":[{"Name":{},"Type":{"Name":{"Relative":"void"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"asCString"},"Flags":32}},{"Name":{},"Type":{"Name":{"Relative":"asCDataType"},"Flags":32}},{"Name":{},"Type":{"Name":{"Relative":"int"}}}]}
	public void  AddVariable(& asCString , & asCDataType , int )                     // dwarf://method/./testdata/game.bz2;332435
	// {"Name":{"Relative":"GetSpaceNeededForArguments","Absolute":"dwarf://method/./testdata/game.bz2;332471"},"Flags":128,"Returns":[{"Name":{},"Type":{"Name":{"Relative":"int"}}}]}
	public int  GetSpaceNeededForArguments()                                         // dwarf://method/./testdata/game.bz2;332471
	// {"Name":{"Relative":"GetSpaceNeededForReturnValue","Absolute":"dwarf://method/./testdata/game.bz2;332496"},"Flags":128,"Returns":[{"Name":{},"Type":{"Name":{"Relative":"int"}}}]}
	public int  GetSpaceNeededForReturnValue()                                       // dwarf://method/./testdata/game.bz2;332496
	// {"Name":{"Relative":"GetDeclarationStr","Absolute":"dwarf://method/./testdata/game.bz2;332521"},"Flags":136,"Returns":[{"Name":{},"Type":{"Name":{"Relative":"asCString"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"bool"}}},{"Name":{},"Type":{"Name":{"Relative":"bool"}}}]}
	public asCString  GetDeclarationStr(bool , bool ) const                          // dwarf://method/./testdata/game.bz2;332521
	// {"Name":{"Relative":"GetLineNumber","Absolute":"dwarf://method/./testdata/game.bz2;332556"},"Flags":128,"Returns":[{"Name":{},"Type":{"Name":{"Relative":"int"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"int"}}},{"Name":{},"Type":{"Name":{},"Specialization":[{"Name":{"Relative":"int"}}],"Flags":2560}}]}
	public int  GetLineNumber(int , int* )                                           // dwarf://method/./testdata/game.bz2;332556
	// {"Name":{"Relative":"ComputeSignatureId","Absolute":"dwarf://method/./testdata/game.bz2;332591"},"Flags":128,"Returns":[{"Name":{},"Type":{"Name":{"Relative":"void"}}}]}
	public void  ComputeSignatureId()                                                // dwarf://method/./testdata/game.bz2;332591
	// {"Name":{"Relative":"IsSignatureEqual","Absolute":"dwarf://method/./testdata/game.bz2;332612"},"Flags":136,"Returns":[{"Name":{},"Type":{"Name":{"Relative":"bool"}}}],"Parameters":[{"Name":{},"Type":{"Name":{},"Specialization":[{"Name":{"Relative":"asCScriptFunction"},"Flags":8}],"Flags":2560}}]}
	public bool  IsSignatureEqual(const asCScriptFunction* ) const                   // dwarf://method/./testdata/game.bz2;332612
	// {"Name":{"Relative":"IsSignatureExceptNameEqual","Absolute":"dwarf://method/./testdata/game.bz2;332642"},"Flags":136,"Returns":[{"Name":{},"Type":{"Name":{"Relative":"bool"}}}],"Parameters":[{"Name":{},"Type":{"Name":{},"Specialization":[{"Name":{"Relative":"asCScriptFunction"},"Flags":8}],"Flags":2560}}]}
	public bool  IsSignatureExceptNameEqual(const asCScriptFunction* ) const         // dwarf://method/./testdata/game.bz2;332642
	// {"Name":{"Relative":"IsSignatureExceptNameEqual","Absolute":"dwarf://method/./testdata/game.bz2;332672"},"Flags":136,"Returns":[{"Name":{},"Type":{"Name":{"Relative":"bool"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"asCDataType"},"Flags":40}},{"Name":{},"Type":{"Name":{"Relative":"asCArray\u003casCDataType\u003e"},"Flags":32}},{"Name":{},"Type":{"Name":{"Relative":"asCArray\u003casETypeModifiers\u003e"},"Flags":32}},{"Name":{},"Type":{"Name":{},"Specialization":[{"Name":{"Relative":"asCObjectType"},"Flags":8}],"Flags":2560}},{"Name":{},"Type":{"Name":{"Relative":"bool"}}}]}
	public bool  IsSignatureExceptNameEqual(const & asCDataType , & asCArray<asCDataType> , & asCArray<asETypeModifiers> , const asCObjectType* , bool ) const // dwarf://method/./testdata/game.bz2;332672
	// {"Name":{"Relative":"IsSignatureExceptNameAndReturnTypeEqual","Absolute":"dwarf://method/./testdata/game.bz2;332722"},"Flags":136,"Returns":[{"Name":{},"Type":{"Name":{"Relative":"bool"}}}],"Parameters":[{"Name":{},"Type":{"Name":{},"Specialization":[{"Name":{"Relative":"asCScriptFunction"},"Flags":8}],"Flags":2560}}]}
	public bool  IsSignatureExceptNameAndReturnTypeEqual(const asCScriptFunction* ) const // dwarf://method/./testdata/game.bz2;332722
	// {"Name":{"Relative":"IsSignatureExceptNameAndReturnTypeEqual","Absolute":"dwarf://method/./testdata/game.bz2;332752"},"Flags":136,"Returns":[{"Name":{},"Type":{"Name":{"Relative":"bool"}}}],"Parameters":[{"Name":{},"Type":{"Name":{"Relative":"asCArray\u003casCDataType\u003e"},"Flags":32}},{"Name":{},"Type":{"Name":{"Relative":"asCArray\u003casETypeModifiers\u003e"},"Flags":32}},{"Name":{},"Type":{"Name":{},"Specialization":[{"Name":{"Relative":"asCObjectType"},"Flags":8}],"Flags":2560}},{"Name":{},"Type":{"Name":{"Relative":"bool"}}}]}
	public bool  IsSignatureExceptNameAndReturnTypeEqual(& asCArray<asCDataType> , & asCArray<asETypeModifiers> , const asCObjectType* , bool ) const // dwarf://method/./testdata/game.bz2;332752
	// {"Name":{"Relative":"DoesReturnOnStack","Absolute":"dwarf://method/./testdata/game.bz2;332797"},"Flags":136,"Returns":[{"Name":{},"Type":{"Name":{"Relative":"bool"}}}]}
	public bool  DoesReturnOnStack() const                                           // dwarf://method/./testdata/game.bz2;332797
	// {"Name":{"Relative":"JITCompile","Absolute":"dwarf://method/./testdata/game.bz2;332822"},"Flags":128,"Returns":[{"Name":{},"Type":{"Name":{"Relative":"void"}}}]}
	public void  JITCompile()                                                        // dwarf://method/./testdata/game.bz2;332822
	// {"Name":{"Relative":"AddReferences","Absolute":"dwarf://method/./testdata/game.bz2;332843"},"Flags":128,"Returns":[{"Name":{},"Type":{"Name":{"Relative":"void"}}}]}
	public void  AddReferences()                                                     // dwarf://method/./testdata/game.bz2;332843
	// {"Name":{"Relative":"ReleaseReferences","Absolute":"dwarf://method/./testdata/game.bz2;332864"},"Flags":128,"Returns":[{"Name":{},"Type":{"Name":{"Relative":"void"}}}]}
	public void  ReleaseReferences()                                                 // dwarf://method/./testdata/game.bz2;332864
	// {"Name":{"Relative":"GetPropertyByGlobalVarPtr","Absolute":"dwarf://method/./testdata/game.bz2;332885"},"Flags":128,"Returns":[{"Name":{},"Type":{"Name":{},"Specialization":[{"Name":{"Relative":"asCGlobalProperty"}}],"Flags":2560}}],"Parameters":[{"Name":{},"Type":{"Name":{},"Specialization":[{"Name":{"Relative":"void"}}],"Flags":2560}}]}
	public asCGlobalProperty*  GetPropertyByGlobalVarPtr(void* )                     // dwarf://method/./testdata/game.bz2;332885
	// {"Name":{"Relative":"GetRefCount","Absolute":"dwarf://method/./testdata/game.bz2;332915"},"Flags":128,"Returns":[{"Name":{},"Type":{"Name":{"Relative":"int"}}}]}
	public int  GetRefCount()                                                        // dwarf://method/./testdata/game.bz2;332915
	// {"Name":{"Relative":"SetFlag","Absolute":"dwarf://method/./testdata/game.bz2;332940"},"Flags":128,"Returns":[{"Name":{},"Type":{"Name":{"Relative":"void"}}}]}
	public void  SetFlag()                                                           // dwarf://method/./testdata/game.bz2;332940
	// {"Name":{"Relative":"GetFlag","Absolute":"dwarf://method/./testdata/game.bz2;332961"},"Flags":128,"Returns":[{"Name":{},"Type":{"Name":{"Relative":"bool"}}}]}
	public bool  GetFlag()                                                           // dwarf://method/./testdata/game.bz2;332961
	// {"Name":{"Relative":"EnumReferences","Absolute":"dwarf://method/./testdata/game.bz2;332986"},"Flags":128,"Returns":[{"Name":{},"Type":{"Name":{"Relative":"void"}}}],"Parameters":[{"Name":{},"Type":{"Name":{},"Specialization":[{"Name":{"Relative":"asIScriptEngine"}}],"Flags":2560}}]}
	public void  EnumReferences(asIScriptEngine* )                                   // dwarf://method/./testdata/game.bz2;332986
	// {"Name":{"Relative":"ReleaseAllHandles","Absolute":"dwarf://method/./testdata/game.bz2;333012"},"Flags":128,"Returns":[{"Name":{},"Type":{"Name":{"Relative":"void"}}}],"Parameters":[{"Name":{},"Type":{"Name":{},"Specialization":[{"Name":{"Relative":"asIScriptEngine"}}],"Flags":2560}}]}
	public void  ReleaseAllHandles(asIScriptEngine* )                                // dwarf://method/./testdata/game.bz2;333012
