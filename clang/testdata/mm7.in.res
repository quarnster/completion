Fields:
	NSArray * accessibilityActionNames
	NSArray * accessibilityAttributeNames
	id accessibilityFocusedUIElement
	BOOL accessibilityIsIgnored
	NSArray * accessibilityParameterizedAttributeNames
	BOOL allowsWeakReference
	NSArray * attributeKeys
	id autoContentAccessingProxy
	id autorelease
	void awakeFromNib
	BOOL boolValue
	NSString * capitalizedString
	Class class
	FourCharCode classCode
	NSClassDescription * classDescription
	Class classForArchiver
	Class classForCoder
	Class classForKeyedArchiver
	Class classForPortCoder
	NSString * className
	BOOL commitEditing
	id copy
	const char * cString
	NSUInteger cStringLength
	void dealloc
	NSString * debugDescription
	NSString * decomposedStringWithCanonicalMapping
	NSString * decomposedStringWithCompatibilityMapping
	NSString * description
	void discardEditing
	double doubleValue
	NSArray * exposedBindings
	NSStringEncoding fastestEncoding
	const char * fileSystemRepresentation
	void finalize
	float floatValue
	NSUInteger hash
	BOOL ignoreModifierKeysWhileDragging
	id init
	NSInteger integerValue
	int intValue
	BOOL isAbsolutePath
	BOOL isExplicitlyIncluded
	BOOL isProxy
	NSString * key
	NSString * lastPathComponent
	NSUInteger length
	NSString * localizedKey
	long long longLongValue
	const char * lossyCString
	NSString * lowercaseString
	id mutableCopy
	NSScriptObjectSpecifier * objectSpecifier
	void * observationInfo
	NSArray * pathComponents
	NSString * pathExtension
	NSString * precomposedStringWithCanonicalMapping
	NSString * precomposedStringWithCompatibilityMapping
	id propertyList
	NSDictionary * propertyListFromStringsFileFormat
	void release
	id retain
	NSUInteger retainCount
	BOOL retainWeakReference
	NSDictionary * scriptingProperties
	id self
	NSStringEncoding smallestEncoding
	NSString * stringByAbbreviatingWithTildeInPath
	NSString * stringByDeletingLastPathComponent
	NSString * stringByDeletingPathExtension
	NSString * stringByExpandingTildeInPath
	NSString * stringByResolvingSymlinksInPath
	NSString * stringByStandardizingPath
	Class superclass
	NSArray * toManyRelationshipKeys
	NSArray * toOneRelationshipKeys
	NSString * uppercaseString
	const char * UTF8String
	id value
	NSZone * zone
Methods:
	NSString *  ((NSString *) accessibilityActionDescription)
	NSUInteger  ((NSString *) accessibilityArrayAttributeCount)
	id  ((NSString *) accessibilityAttributeValue)
	id  ((NSPoint) accessibilityHitTest)
	NSUInteger  ((id) accessibilityIndexOfChild)
	BOOL  ((NSString *) accessibilityIsAttributeSettable)
	void  ((NSString *) accessibilityPerformAction)
	id  ((NSCoder *) awakeAfterUsingCoder)
	BOOL  ((NSStringEncoding) canBeConvertedToEncoding)
	NSString *  ((NSLocale *) capitalizedStringWithLocale)
	NSComparisonResult  ((NSString *) caseInsensitiveCompare)
	void  ((id) changeColor)
	void  ((id) changeFont)
	unichar  ((NSUInteger) characterAtIndex)
	BOOL  ((NSError **) commitEditingAndReturnError)
	NSComparisonResult  ((NSString *) compare)
	NSArray *  ((NSCharacterSet *) componentsSeparatedByCharactersInSet)
	NSArray *  ((NSString *) componentsSeparatedByString)
	BOOL  ((Protocol *) conformsToProtocol)
	void  ((NSNotification *) controlTextDidBeginEditing)
	void  ((NSNotification *) controlTextDidChange)
	void  ((NSNotification *) controlTextDidEndEditing)
	id  ((NSZone *) copyWithZone)
	const char *  ((NSStringEncoding) cStringUsingEncoding)
	NSData *  ((NSStringEncoding) dataUsingEncoding)
	NSDictionary *  ((NSArray *) dictionaryWithValuesForKeys)
	void  ((NSString *) didChangeValueForKey)
	void  ((CALayer *) displayLayer)
	BOOL  ((id) doesContain)
	void  ((SEL) doesNotRecognizeSelector)
	NSDragOperation  ((BOOL) draggingSourceOperationMaskForLocal)
	void  ((NSCoder *) encodeWithCoder)
	void  (^(NSString *line, BOOL *stop)block enumerateLinesUsingBlock)
	id  ((SEL) forwardingTargetForSelector)
	void  ((NSInvocation *) forwardInvocation)
	void  ((unichar *) getCharacters)
	void  ((char *) getCString)
	id  ((NSString *) handleQueryWithUnboundKey)
	BOOL  ((NSString *) hasPrefix)
	BOOL  ((NSString *) hasSuffix)
	NSArray *  ((NSScriptObjectSpecifier *) indicesOfObjectsByEvaluatingObjectSpecifier)
	NSDictionary *  ((NSString *) infoForBinding)
	id  ((NSCoder *) initWithCoder)
	id  ((NSString *) initWithContentsOfFile)
	id  ((NSURL *) initWithContentsOfURL)
	id  ((const char *) initWithCString)
	id  ((NSString *), ... initWithFormat)
	id  ((NSString *) initWithString)
	id  ((const char *) initWithUTF8String)
	void  ((CALayer *) invalidateLayoutOfLayer)
	NSString *  ((NSString *) inverseForRelationshipKey)
	BOOL  ((NSString *) isCaseInsensitiveLike)
	BOOL  ((id) isEqual)
	BOOL  ((id) isEqualTo)
	BOOL  ((NSString *) isEqualToString)
	BOOL  ((id) isGreaterThan)
	BOOL  ((id) isGreaterThanOrEqualTo)
	BOOL  ((Class) isKindOfClass)
	BOOL  ((id) isLessThan)
	BOOL  ((id) isLessThanOrEqualTo)
	BOOL  ((NSString *) isLike)
	BOOL  ((Class) isMemberOfClass)
	BOOL  ((id) isNotEqualTo)
	void  ((CALayer *) layoutSublayersOfLayer)
	NSUInteger  ((NSStringEncoding) lengthOfBytesUsingEncoding)
	NSRange  ((NSRange) lineRangeForRange)
	NSComparisonResult  ((NSString *) localizedCaseInsensitiveCompare)
	NSComparisonResult  ((NSString *) localizedCompare)
	NSComparisonResult  ((NSString *) localizedStandardCompare)
	NSString *  ((NSLocale *) lowercaseStringWithLocale)
	NSUInteger  ((NSStringEncoding) maximumLengthOfBytesUsingEncoding)
	IMP  ((SEL) methodForSelector)
	NSMethodSignature *  ((SEL) methodSignatureForSelector)
	NSMutableArray *  ((NSString *) mutableArrayValueForKey)
	NSMutableArray *  ((NSString *) mutableArrayValueForKeyPath)
	id  ((NSZone *) mutableCopyWithZone)
	NSMutableOrderedSet *  ((NSString *) mutableOrderedSetValueForKey)
	NSMutableOrderedSet *  ((NSString *) mutableOrderedSetValueForKeyPath)
	NSMutableSet *  ((NSString *) mutableSetValueForKey)
	NSMutableSet *  ((NSString *) mutableSetValueForKeyPath)
	NSArray *  ((NSURL *) namesOfPromisedFilesDroppedAtDestination)
	void  ((id) objectDidBeginEditing)
	void  ((id) objectDidEndEditing)
	NSArray *  ((NSString *) optionDescriptionsForBinding)
	NSRange  ((NSRange) paragraphRangeForRange)
	void  ((NSPasteboard *) pasteboardChangedOwner)
	id  ((NSString *) pasteboardPropertyListForType)
	id  ((SEL) performSelector)
	CGSize  ((CALayer *) preferredSizeOfLayer)
	NSRange  ((NSCharacterSet *) rangeOfCharacterFromSet)
	NSRange  ((NSUInteger) rangeOfComposedCharacterSequenceAtIndex)
	NSRange  ((NSRange) rangeOfComposedCharacterSequencesForRange)
	NSRange  ((NSString *) rangeOfString)
	BOOL  ((NSPasteboard *) readSelectionFromPasteboard)
	id  ((NSArchiver *) replacementObjectForArchiver)
	id  ((NSCoder *) replacementObjectForCoder)
	id  ((NSKeyedArchiver *) replacementObjectForKeyedArchiver)
	id  ((NSPortCoder *) replacementObjectForPortCoder)
	BOOL  ((SEL) respondsToSelector)
	BOOL  ((id) scriptingBeginsWith)
	BOOL  ((id) scriptingContains)
	BOOL  ((id) scriptingEndsWith)
	BOOL  ((id) scriptingIsEqualTo)
	BOOL  ((id) scriptingIsGreaterThan)
	BOOL  ((id) scriptingIsGreaterThanOrEqualTo)
	BOOL  ((id) scriptingIsLessThan)
	BOOL  ((id) scriptingIsLessThanOrEqualTo)
	id  ((NSScriptObjectSpecifier *) scriptingValueForSpecifier)
	void  ((NSString *) setKey)
	void  ((NSString *) setLocalizedKey)
	void  ((NSString *) setNilValueForKey)
	void  ((void *) setObservationInfo)
	void  ((NSDictionary *) setScriptingProperties)
	void  ((id) setValue)
	void  ((NSDictionary *) setValuesForKeysWithDictionary)
	NSSize  ((NSDictionary *) sizeWithAttributes)
	id  ((NSString *) storedValueForKey)
	NSString *  ((NSStringEncoding) stringByAddingPercentEscapesUsingEncoding)
	NSString *  ((NSString *), ... stringByAppendingFormat)
	NSString *  ((NSString *) stringByAppendingPathComponent)
	NSString *  ((NSString *) stringByAppendingPathExtension)
	NSString *  ((NSString *) stringByAppendingString)
	NSString *  ((NSStringEncoding) stringByReplacingPercentEscapesUsingEncoding)
	NSString *  ((NSCharacterSet *) stringByTrimmingCharactersInSet)
	NSArray *  ((NSArray *) stringsByAppendingPaths)
	NSString *  ((NSUInteger) substringFromIndex)
	NSString *  ((NSUInteger) substringToIndex)
	NSString *  ((NSRange) substringWithRange)
	void  ((NSDictionary *) takeValuesFromDictionary)
	void  ((NSString *) unableToSetNilForKey)
	void  ((NSString *) unbind)
	NSString *  ((NSLocale *) uppercaseStringWithLocale)
	void  ((NSURL *) URLResourceDidCancelLoading)
	void  ((NSURL *) URLResourceDidFinishLoading)
	BOOL  ((NSMenuItem *) validateMenuItem)
	BOOL  ((NSToolbarItem *) validateToolbarItem)
	NSUInteger  ((NSFontPanel *) validModesForFontPanel)
	Class  ((NSString *) valueClassForBinding)
	id  ((NSString *) valueForKey)
	id  ((NSString *) valueForKeyPath)
	id  ((NSString *) valueForUndefinedKey)
	NSDictionary *  ((NSArray *) valuesForKeys)
	void  ((NSString *) willChangeValueForKey)
	NSArray *  ((NSPasteboard *) writableTypesForPasteboard)
