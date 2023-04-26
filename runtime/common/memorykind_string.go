// Code generated by "stringer -type=MemoryKind -trimprefix=MemoryKind"; DO NOT EDIT.

package common

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[MemoryKindUnknown-0]
	_ = x[MemoryKindAddressValue-1]
	_ = x[MemoryKindStringValue-2]
	_ = x[MemoryKindCharacterValue-3]
	_ = x[MemoryKindNumberValue-4]
	_ = x[MemoryKindArrayValueBase-5]
	_ = x[MemoryKindDictionaryValueBase-6]
	_ = x[MemoryKindCompositeValueBase-7]
	_ = x[MemoryKindSimpleCompositeValueBase-8]
	_ = x[MemoryKindOptionalValue-9]
	_ = x[MemoryKindTypeValue-10]
	_ = x[MemoryKindPathValue-11]
	_ = x[MemoryKindIDCapabilityValue-12]
	_ = x[MemoryKindPathCapabilityValue-13]
	_ = x[MemoryKindPathLinkValue-14]
	_ = x[MemoryKindAccountLinkValue-15]
	_ = x[MemoryKindStorageReferenceValue-16]
	_ = x[MemoryKindAccountReferenceValue-17]
	_ = x[MemoryKindEphemeralReferenceValue-18]
	_ = x[MemoryKindInterpretedFunctionValue-19]
	_ = x[MemoryKindHostFunctionValue-20]
	_ = x[MemoryKindBoundFunctionValue-21]
	_ = x[MemoryKindBigInt-22]
	_ = x[MemoryKindSimpleCompositeValue-23]
	_ = x[MemoryKindPublishedValue-24]
	_ = x[MemoryKindStorageCapabilityControllerValue-25]
	_ = x[MemoryKindAccountCapabilityControllerValue-26]
	_ = x[MemoryKindAtreeArrayDataSlab-27]
	_ = x[MemoryKindAtreeArrayMetaDataSlab-28]
	_ = x[MemoryKindAtreeArrayElementOverhead-29]
	_ = x[MemoryKindAtreeMapDataSlab-30]
	_ = x[MemoryKindAtreeMapMetaDataSlab-31]
	_ = x[MemoryKindAtreeMapElementOverhead-32]
	_ = x[MemoryKindAtreeMapPreAllocatedElement-33]
	_ = x[MemoryKindAtreeEncodedSlab-34]
	_ = x[MemoryKindPrimitiveStaticType-35]
	_ = x[MemoryKindCompositeStaticType-36]
	_ = x[MemoryKindInterfaceStaticType-37]
	_ = x[MemoryKindVariableSizedStaticType-38]
	_ = x[MemoryKindConstantSizedStaticType-39]
	_ = x[MemoryKindDictionaryStaticType-40]
	_ = x[MemoryKindOptionalStaticType-41]
	_ = x[MemoryKindRestrictedStaticType-42]
	_ = x[MemoryKindReferenceStaticType-43]
	_ = x[MemoryKindCapabilityStaticType-44]
	_ = x[MemoryKindFunctionStaticType-45]
	_ = x[MemoryKindCadenceVoidValue-46]
	_ = x[MemoryKindCadenceOptionalValue-47]
	_ = x[MemoryKindCadenceBoolValue-48]
	_ = x[MemoryKindCadenceStringValue-49]
	_ = x[MemoryKindCadenceCharacterValue-50]
	_ = x[MemoryKindCadenceAddressValue-51]
	_ = x[MemoryKindCadenceIntValue-52]
	_ = x[MemoryKindCadenceNumberValue-53]
	_ = x[MemoryKindCadenceArrayValueBase-54]
	_ = x[MemoryKindCadenceArrayValueLength-55]
	_ = x[MemoryKindCadenceDictionaryValue-56]
	_ = x[MemoryKindCadenceKeyValuePair-57]
	_ = x[MemoryKindCadenceStructValueBase-58]
	_ = x[MemoryKindCadenceStructValueSize-59]
	_ = x[MemoryKindCadenceResourceValueBase-60]
	_ = x[MemoryKindCadenceAttachmentValueBase-61]
	_ = x[MemoryKindCadenceResourceValueSize-62]
	_ = x[MemoryKindCadenceAttachmentValueSize-63]
	_ = x[MemoryKindCadenceEventValueBase-64]
	_ = x[MemoryKindCadenceEventValueSize-65]
	_ = x[MemoryKindCadenceContractValueBase-66]
	_ = x[MemoryKindCadenceContractValueSize-67]
	_ = x[MemoryKindCadenceEnumValueBase-68]
	_ = x[MemoryKindCadenceEnumValueSize-69]
	_ = x[MemoryKindCadencePathLinkValue-70]
	_ = x[MemoryKindCadenceAccountLinkValue-71]
	_ = x[MemoryKindCadencePathValue-72]
	_ = x[MemoryKindCadenceTypeValue-73]
	_ = x[MemoryKindCadenceIDCapabilityValue-74]
	_ = x[MemoryKindCadencePathCapabilityValue-75]
	_ = x[MemoryKindCadenceFunctionValue-76]
	_ = x[MemoryKindCadenceOptionalType-77]
	_ = x[MemoryKindCadenceVariableSizedArrayType-78]
	_ = x[MemoryKindCadenceConstantSizedArrayType-79]
	_ = x[MemoryKindCadenceDictionaryType-80]
	_ = x[MemoryKindCadenceField-81]
	_ = x[MemoryKindCadenceParameter-82]
	_ = x[MemoryKindCadenceTypeParameter-83]
	_ = x[MemoryKindCadenceStructType-84]
	_ = x[MemoryKindCadenceResourceType-85]
	_ = x[MemoryKindCadenceAttachmentType-86]
	_ = x[MemoryKindCadenceEventType-87]
	_ = x[MemoryKindCadenceContractType-88]
	_ = x[MemoryKindCadenceStructInterfaceType-89]
	_ = x[MemoryKindCadenceResourceInterfaceType-90]
	_ = x[MemoryKindCadenceContractInterfaceType-91]
	_ = x[MemoryKindCadenceFunctionType-92]
	_ = x[MemoryKindCadenceReferenceType-93]
	_ = x[MemoryKindCadenceRestrictedType-94]
	_ = x[MemoryKindCadenceCapabilityType-95]
	_ = x[MemoryKindCadenceEnumType-96]
	_ = x[MemoryKindRawString-97]
	_ = x[MemoryKindAddressLocation-98]
	_ = x[MemoryKindBytes-99]
	_ = x[MemoryKindVariable-100]
	_ = x[MemoryKindCompositeTypeInfo-101]
	_ = x[MemoryKindCompositeField-102]
	_ = x[MemoryKindInvocation-103]
	_ = x[MemoryKindStorageMap-104]
	_ = x[MemoryKindStorageKey-105]
	_ = x[MemoryKindTypeToken-106]
	_ = x[MemoryKindErrorToken-107]
	_ = x[MemoryKindSpaceToken-108]
	_ = x[MemoryKindProgram-109]
	_ = x[MemoryKindIdentifier-110]
	_ = x[MemoryKindArgument-111]
	_ = x[MemoryKindBlock-112]
	_ = x[MemoryKindFunctionBlock-113]
	_ = x[MemoryKindParameter-114]
	_ = x[MemoryKindParameterList-115]
	_ = x[MemoryKindTypeParameter-116]
	_ = x[MemoryKindTypeParameterList-117]
	_ = x[MemoryKindTransfer-118]
	_ = x[MemoryKindMembers-119]
	_ = x[MemoryKindTypeAnnotation-120]
	_ = x[MemoryKindDictionaryEntry-121]
	_ = x[MemoryKindFunctionDeclaration-122]
	_ = x[MemoryKindCompositeDeclaration-123]
	_ = x[MemoryKindAttachmentDeclaration-124]
	_ = x[MemoryKindInterfaceDeclaration-125]
	_ = x[MemoryKindEnumCaseDeclaration-126]
	_ = x[MemoryKindFieldDeclaration-127]
	_ = x[MemoryKindTransactionDeclaration-128]
	_ = x[MemoryKindImportDeclaration-129]
	_ = x[MemoryKindVariableDeclaration-130]
	_ = x[MemoryKindSpecialFunctionDeclaration-131]
	_ = x[MemoryKindPragmaDeclaration-132]
	_ = x[MemoryKindAssignmentStatement-133]
	_ = x[MemoryKindBreakStatement-134]
	_ = x[MemoryKindContinueStatement-135]
	_ = x[MemoryKindEmitStatement-136]
	_ = x[MemoryKindExpressionStatement-137]
	_ = x[MemoryKindForStatement-138]
	_ = x[MemoryKindIfStatement-139]
	_ = x[MemoryKindReturnStatement-140]
	_ = x[MemoryKindSwapStatement-141]
	_ = x[MemoryKindSwitchStatement-142]
	_ = x[MemoryKindWhileStatement-143]
	_ = x[MemoryKindRemoveStatement-144]
	_ = x[MemoryKindBooleanExpression-145]
	_ = x[MemoryKindVoidExpression-146]
	_ = x[MemoryKindNilExpression-147]
	_ = x[MemoryKindStringExpression-148]
	_ = x[MemoryKindIntegerExpression-149]
	_ = x[MemoryKindFixedPointExpression-150]
	_ = x[MemoryKindArrayExpression-151]
	_ = x[MemoryKindDictionaryExpression-152]
	_ = x[MemoryKindIdentifierExpression-153]
	_ = x[MemoryKindInvocationExpression-154]
	_ = x[MemoryKindMemberExpression-155]
	_ = x[MemoryKindIndexExpression-156]
	_ = x[MemoryKindConditionalExpression-157]
	_ = x[MemoryKindUnaryExpression-158]
	_ = x[MemoryKindBinaryExpression-159]
	_ = x[MemoryKindFunctionExpression-160]
	_ = x[MemoryKindCastingExpression-161]
	_ = x[MemoryKindCreateExpression-162]
	_ = x[MemoryKindDestroyExpression-163]
	_ = x[MemoryKindReferenceExpression-164]
	_ = x[MemoryKindForceExpression-165]
	_ = x[MemoryKindPathExpression-166]
	_ = x[MemoryKindAttachExpression-167]
	_ = x[MemoryKindConstantSizedType-168]
	_ = x[MemoryKindDictionaryType-169]
	_ = x[MemoryKindFunctionType-170]
	_ = x[MemoryKindInstantiationType-171]
	_ = x[MemoryKindNominalType-172]
	_ = x[MemoryKindOptionalType-173]
	_ = x[MemoryKindReferenceType-174]
	_ = x[MemoryKindRestrictedType-175]
	_ = x[MemoryKindVariableSizedType-176]
	_ = x[MemoryKindPosition-177]
	_ = x[MemoryKindRange-178]
	_ = x[MemoryKindElaboration-179]
	_ = x[MemoryKindActivation-180]
	_ = x[MemoryKindActivationEntries-181]
	_ = x[MemoryKindVariableSizedSemaType-182]
	_ = x[MemoryKindConstantSizedSemaType-183]
	_ = x[MemoryKindDictionarySemaType-184]
	_ = x[MemoryKindOptionalSemaType-185]
	_ = x[MemoryKindRestrictedSemaType-186]
	_ = x[MemoryKindReferenceSemaType-187]
	_ = x[MemoryKindCapabilitySemaType-188]
	_ = x[MemoryKindOrderedMap-189]
	_ = x[MemoryKindOrderedMapEntryList-190]
	_ = x[MemoryKindOrderedMapEntry-191]
	_ = x[MemoryKindLast-192]
}

const _MemoryKind_name = "UnknownAddressValueStringValueCharacterValueNumberValueArrayValueBaseDictionaryValueBaseCompositeValueBaseSimpleCompositeValueBaseOptionalValueTypeValuePathValueIDCapabilityValuePathCapabilityValuePathLinkValueAccountLinkValueStorageReferenceValueAccountReferenceValueEphemeralReferenceValueInterpretedFunctionValueHostFunctionValueBoundFunctionValueBigIntSimpleCompositeValuePublishedValueStorageCapabilityControllerValueAccountCapabilityControllerValueAtreeArrayDataSlabAtreeArrayMetaDataSlabAtreeArrayElementOverheadAtreeMapDataSlabAtreeMapMetaDataSlabAtreeMapElementOverheadAtreeMapPreAllocatedElementAtreeEncodedSlabPrimitiveStaticTypeCompositeStaticTypeInterfaceStaticTypeVariableSizedStaticTypeConstantSizedStaticTypeDictionaryStaticTypeOptionalStaticTypeRestrictedStaticTypeReferenceStaticTypeCapabilityStaticTypeFunctionStaticTypeCadenceVoidValueCadenceOptionalValueCadenceBoolValueCadenceStringValueCadenceCharacterValueCadenceAddressValueCadenceIntValueCadenceNumberValueCadenceArrayValueBaseCadenceArrayValueLengthCadenceDictionaryValueCadenceKeyValuePairCadenceStructValueBaseCadenceStructValueSizeCadenceResourceValueBaseCadenceAttachmentValueBaseCadenceResourceValueSizeCadenceAttachmentValueSizeCadenceEventValueBaseCadenceEventValueSizeCadenceContractValueBaseCadenceContractValueSizeCadenceEnumValueBaseCadenceEnumValueSizeCadencePathLinkValueCadenceAccountLinkValueCadencePathValueCadenceTypeValueCadenceIDCapabilityValueCadencePathCapabilityValueCadenceFunctionValueCadenceOptionalTypeCadenceVariableSizedArrayTypeCadenceConstantSizedArrayTypeCadenceDictionaryTypeCadenceFieldCadenceParameterCadenceTypeParameterCadenceStructTypeCadenceResourceTypeCadenceAttachmentTypeCadenceEventTypeCadenceContractTypeCadenceStructInterfaceTypeCadenceResourceInterfaceTypeCadenceContractInterfaceTypeCadenceFunctionTypeCadenceReferenceTypeCadenceRestrictedTypeCadenceCapabilityTypeCadenceEnumTypeRawStringAddressLocationBytesVariableCompositeTypeInfoCompositeFieldInvocationStorageMapStorageKeyTypeTokenErrorTokenSpaceTokenProgramIdentifierArgumentBlockFunctionBlockParameterParameterListTypeParameterTypeParameterListTransferMembersTypeAnnotationDictionaryEntryFunctionDeclarationCompositeDeclarationAttachmentDeclarationInterfaceDeclarationEnumCaseDeclarationFieldDeclarationTransactionDeclarationImportDeclarationVariableDeclarationSpecialFunctionDeclarationPragmaDeclarationAssignmentStatementBreakStatementContinueStatementEmitStatementExpressionStatementForStatementIfStatementReturnStatementSwapStatementSwitchStatementWhileStatementRemoveStatementBooleanExpressionVoidExpressionNilExpressionStringExpressionIntegerExpressionFixedPointExpressionArrayExpressionDictionaryExpressionIdentifierExpressionInvocationExpressionMemberExpressionIndexExpressionConditionalExpressionUnaryExpressionBinaryExpressionFunctionExpressionCastingExpressionCreateExpressionDestroyExpressionReferenceExpressionForceExpressionPathExpressionAttachExpressionConstantSizedTypeDictionaryTypeFunctionTypeInstantiationTypeNominalTypeOptionalTypeReferenceTypeRestrictedTypeVariableSizedTypePositionRangeElaborationActivationActivationEntriesVariableSizedSemaTypeConstantSizedSemaTypeDictionarySemaTypeOptionalSemaTypeRestrictedSemaTypeReferenceSemaTypeCapabilitySemaTypeOrderedMapOrderedMapEntryListOrderedMapEntryLast"

var _MemoryKind_index = [...]uint16{0, 7, 19, 30, 44, 55, 69, 88, 106, 130, 143, 152, 161, 178, 197, 210, 226, 247, 268, 291, 315, 332, 350, 356, 376, 390, 422, 454, 472, 494, 519, 535, 555, 578, 605, 621, 640, 659, 678, 701, 724, 744, 762, 782, 801, 821, 839, 855, 875, 891, 909, 930, 949, 964, 982, 1003, 1026, 1048, 1067, 1089, 1111, 1135, 1161, 1185, 1211, 1232, 1253, 1277, 1301, 1321, 1341, 1361, 1384, 1400, 1416, 1440, 1466, 1486, 1505, 1534, 1563, 1584, 1596, 1612, 1632, 1649, 1668, 1689, 1705, 1724, 1750, 1778, 1806, 1825, 1845, 1866, 1887, 1902, 1911, 1926, 1931, 1939, 1956, 1970, 1980, 1990, 2000, 2009, 2019, 2029, 2036, 2046, 2054, 2059, 2072, 2081, 2094, 2107, 2124, 2132, 2139, 2153, 2168, 2187, 2207, 2228, 2248, 2267, 2283, 2305, 2322, 2341, 2367, 2384, 2403, 2417, 2434, 2447, 2466, 2478, 2489, 2504, 2517, 2532, 2546, 2561, 2578, 2592, 2605, 2621, 2638, 2658, 2673, 2693, 2713, 2733, 2749, 2764, 2785, 2800, 2816, 2834, 2851, 2867, 2884, 2903, 2918, 2932, 2948, 2965, 2979, 2991, 3008, 3019, 3031, 3044, 3058, 3075, 3083, 3088, 3099, 3109, 3126, 3147, 3168, 3186, 3202, 3220, 3237, 3255, 3265, 3284, 3299, 3303}

func (i MemoryKind) String() string {
	if i >= MemoryKind(len(_MemoryKind_index)-1) {
		return "MemoryKind(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _MemoryKind_name[_MemoryKind_index[i]:_MemoryKind_index[i+1]]
}
