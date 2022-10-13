// Code generated by "stringer -type=ElementType"; DO NOT EDIT.

package ast

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[ElementTypeUnknown-0]
	_ = x[ElementTypeProgram-1]
	_ = x[ElementTypeBlock-2]
	_ = x[ElementTypeFunctionBlock-3]
	_ = x[ElementTypeFunctionDeclaration-4]
	_ = x[ElementTypeSpecialFunctionDeclaration-5]
	_ = x[ElementTypeCompositeDeclaration-6]
	_ = x[ElementTypeInterfaceDeclaration-7]
	_ = x[ElementTypeAttachmentDeclaration-8]
	_ = x[ElementTypeFieldDeclaration-9]
	_ = x[ElementTypeEnumCaseDeclaration-10]
	_ = x[ElementTypePragmaDeclaration-11]
	_ = x[ElementTypeImportDeclaration-12]
	_ = x[ElementTypeTransactionDeclaration-13]
	_ = x[ElementTypeReturnStatement-14]
	_ = x[ElementTypeBreakStatement-15]
	_ = x[ElementTypeContinueStatement-16]
	_ = x[ElementTypeIfStatement-17]
	_ = x[ElementTypeSwitchStatement-18]
	_ = x[ElementTypeWhileStatement-19]
	_ = x[ElementTypeForStatement-20]
	_ = x[ElementTypeEmitStatement-21]
	_ = x[ElementTypeVariableDeclaration-22]
	_ = x[ElementTypeAssignmentStatement-23]
	_ = x[ElementTypeSwapStatement-24]
	_ = x[ElementTypeExpressionStatement-25]
	_ = x[ElementTypeRemoveStatement-26]
	_ = x[ElementTypeVoidExpression-27]
	_ = x[ElementTypeBoolExpression-28]
	_ = x[ElementTypeNilExpression-29]
	_ = x[ElementTypeIntegerExpression-30]
	_ = x[ElementTypeFixedPointExpression-31]
	_ = x[ElementTypeArrayExpression-32]
	_ = x[ElementTypeDictionaryExpression-33]
	_ = x[ElementTypeIdentifierExpression-34]
	_ = x[ElementTypeInvocationExpression-35]
	_ = x[ElementTypeMemberExpression-36]
	_ = x[ElementTypeIndexExpression-37]
	_ = x[ElementTypeConditionalExpression-38]
	_ = x[ElementTypeUnaryExpression-39]
	_ = x[ElementTypeBinaryExpression-40]
	_ = x[ElementTypeFunctionExpression-41]
	_ = x[ElementTypeStringExpression-42]
	_ = x[ElementTypeCastingExpression-43]
	_ = x[ElementTypeCreateExpression-44]
	_ = x[ElementTypeDestroyExpression-45]
	_ = x[ElementTypeReferenceExpression-46]
	_ = x[ElementTypeForceExpression-47]
	_ = x[ElementTypePathExpression-48]
	_ = x[ElementTypeAttachExpression-49]
}

const _ElementType_name = "ElementTypeUnknownElementTypeProgramElementTypeBlockElementTypeFunctionBlockElementTypeFunctionDeclarationElementTypeSpecialFunctionDeclarationElementTypeCompositeDeclarationElementTypeInterfaceDeclarationElementTypeAttachmentDeclarationElementTypeFieldDeclarationElementTypeEnumCaseDeclarationElementTypePragmaDeclarationElementTypeImportDeclarationElementTypeTransactionDeclarationElementTypeReturnStatementElementTypeBreakStatementElementTypeContinueStatementElementTypeIfStatementElementTypeSwitchStatementElementTypeWhileStatementElementTypeForStatementElementTypeEmitStatementElementTypeVariableDeclarationElementTypeAssignmentStatementElementTypeSwapStatementElementTypeExpressionStatementElementTypeRemoveStatementElementTypeVoidExpressionElementTypeBoolExpressionElementTypeNilExpressionElementTypeIntegerExpressionElementTypeFixedPointExpressionElementTypeArrayExpressionElementTypeDictionaryExpressionElementTypeIdentifierExpressionElementTypeInvocationExpressionElementTypeMemberExpressionElementTypeIndexExpressionElementTypeConditionalExpressionElementTypeUnaryExpressionElementTypeBinaryExpressionElementTypeFunctionExpressionElementTypeStringExpressionElementTypeCastingExpressionElementTypeCreateExpressionElementTypeDestroyExpressionElementTypeReferenceExpressionElementTypeForceExpressionElementTypePathExpressionElementTypeAttachExpression"

var _ElementType_index = [...]uint16{0, 18, 36, 52, 76, 106, 143, 174, 205, 237, 264, 294, 322, 350, 383, 409, 434, 462, 484, 510, 535, 558, 582, 612, 642, 666, 696, 722, 747, 772, 796, 824, 855, 881, 912, 943, 974, 1001, 1027, 1059, 1085, 1112, 1141, 1168, 1196, 1223, 1251, 1281, 1307, 1332, 1359}

func (i ElementType) String() string {
	if i >= ElementType(len(_ElementType_index)-1) {
		return "ElementType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _ElementType_name[_ElementType_index[i]:_ElementType_index[i+1]]
}
