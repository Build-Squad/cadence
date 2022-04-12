/*
 * Cadence - The resource-oriented smart contract programming language
 *
 * Copyright 2019-2022 Dapper Labs, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package sema

import (
	"github.com/onflow/cadence/runtime/ast"
	"github.com/onflow/cadence/runtime/common"
)

type ValueDeclaration interface {
	ValueDeclarationName() string
	ValueDeclarationType() Type
	ValueDeclarationDocString() string
	ValueDeclarationKind() common.DeclarationKind
	ValueDeclarationPosition() ast.Position
	ValueDeclarationIsConstant() bool
	ValueDeclarationArgumentLabels() []string
	ValueDeclarationAvailable(common.Location) bool
}

type TypeDeclaration interface {
	TypeDeclarationName() string
	TypeDeclarationType() Type
	TypeDeclarationKind() common.DeclarationKind
	TypeDeclarationPosition() ast.Position
}
