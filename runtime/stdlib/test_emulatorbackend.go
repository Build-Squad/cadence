/*
 * Cadence - The resource-oriented smart contract programming language
 *
 * Copyright Dapper Labs, Inc.
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
 *
 */

package stdlib

import (
	"github.com/onflow/cadence/runtime/common"
	"github.com/onflow/cadence/runtime/errors"
	"github.com/onflow/cadence/runtime/interpreter"
	"github.com/onflow/cadence/runtime/sema"
)

// 'EmulatorBackend' struct.
//
// 'EmulatorBackend' is the native implementation of the 'Test.BlockchainBackend' interface.
// It provides a blockchain backed by the emulator.

const testEmulatorBackendTypeName = "EmulatorBackend"

type testEmulatorBackendType struct {
	compositeType                      *sema.CompositeType
	executeScriptFunctionType          *sema.FunctionType
	createAccountFunctionType          *sema.FunctionType
	addTransactionFunctionType         *sema.FunctionType
	executeNextTransactionFunctionType *sema.FunctionType
	commitBlockFunctionType            *sema.FunctionType
	deployContractFunctionType         *sema.FunctionType
	useConfigFunctionType              *sema.FunctionType
	logsFunctionType                   *sema.FunctionType
	serviceAccountFunctionType         *sema.FunctionType
	eventsFunctionType                 *sema.FunctionType
	resetFunctionType                  *sema.FunctionType
}

func newTestEmulatorBackendType(blockchainBackendInterfaceType *sema.InterfaceType) *testEmulatorBackendType {
	executeScriptFunctionType := interfaceFunctionType(
		blockchainBackendInterfaceType,
		testEmulatorBackendTypeExecuteScriptFunctionName,
	)

	createAccountFunctionType := interfaceFunctionType(
		blockchainBackendInterfaceType,
		testEmulatorBackendTypeCreateAccountFunctionName,
	)

	addTransactionFunctionType := interfaceFunctionType(
		blockchainBackendInterfaceType,
		testEmulatorBackendTypeAddTransactionFunctionName,
	)

	executeNextTransactionFunctionType := interfaceFunctionType(
		blockchainBackendInterfaceType,
		testEmulatorBackendTypeExecuteNextTransactionFunctionName,
	)

	commitBlockFunctionType := interfaceFunctionType(
		blockchainBackendInterfaceType,
		testEmulatorBackendTypeCommitBlockFunctionName,
	)

	deployContractFunctionType := interfaceFunctionType(
		blockchainBackendInterfaceType,
		testEmulatorBackendTypeDeployContractFunctionName,
	)

	useConfigFunctionType := interfaceFunctionType(
		blockchainBackendInterfaceType,
		testEmulatorBackendTypeUseConfigFunctionName,
	)

	logsFunctionType := interfaceFunctionType(
		blockchainBackendInterfaceType,
		testEmulatorBackendTypeLogsFunctionName,
	)

	serviceAccountFunctionType := interfaceFunctionType(
		blockchainBackendInterfaceType,
		testEmulatorBackendTypeServiceAccountFunctionName,
	)

	eventsFunctionType := interfaceFunctionType(
		blockchainBackendInterfaceType,
		testEmulatorBackendTypeEventsFunctionName,
	)

	resetFunctionType := interfaceFunctionType(
		blockchainBackendInterfaceType,
		testEmulatorBackendTypeResetFunctionName,
	)

	compositeType := &sema.CompositeType{
		Identifier: testEmulatorBackendTypeName,
		Kind:       common.CompositeKindStructure,
		Location:   TestContractLocation,
		ExplicitInterfaceConformances: []*sema.InterfaceType{
			blockchainBackendInterfaceType,
		},
	}

	var members = []*sema.Member{
		sema.NewUnmeteredPublicFunctionMember(
			compositeType,
			testEmulatorBackendTypeExecuteScriptFunctionName,
			executeScriptFunctionType,
			testEmulatorBackendTypeExecuteScriptFunctionDocString,
		),
		sema.NewUnmeteredPublicFunctionMember(
			compositeType,
			testEmulatorBackendTypeCreateAccountFunctionName,
			createAccountFunctionType,
			testEmulatorBackendTypeCreateAccountFunctionDocString,
		),
		sema.NewUnmeteredPublicFunctionMember(
			compositeType,
			testEmulatorBackendTypeAddTransactionFunctionName,
			addTransactionFunctionType,
			testEmulatorBackendTypeAddTransactionFunctionDocString,
		),
		sema.NewUnmeteredPublicFunctionMember(
			compositeType,
			testEmulatorBackendTypeExecuteNextTransactionFunctionName,
			executeNextTransactionFunctionType,
			testEmulatorBackendTypeExecuteNextTransactionFunctionDocString,
		),
		sema.NewUnmeteredPublicFunctionMember(
			compositeType,
			testEmulatorBackendTypeCommitBlockFunctionName,
			commitBlockFunctionType,
			testEmulatorBackendTypeCommitBlockFunctionDocString,
		),
		sema.NewUnmeteredPublicFunctionMember(
			compositeType,
			testEmulatorBackendTypeDeployContractFunctionName,
			deployContractFunctionType,
			testEmulatorBackendTypeDeployContractFunctionDocString,
		),
		sema.NewUnmeteredPublicFunctionMember(
			compositeType,
			testEmulatorBackendTypeUseConfigFunctionName,
			useConfigFunctionType,
			testEmulatorBackendTypeUseConfigFunctionDocString,
		),
		sema.NewUnmeteredPublicFunctionMember(
			compositeType,
			testEmulatorBackendTypeLogsFunctionName,
			logsFunctionType,
			testEmulatorBackendTypeLogsFunctionDocString,
		),
		sema.NewUnmeteredPublicFunctionMember(
			compositeType,
			testEmulatorBackendTypeServiceAccountFunctionName,
			serviceAccountFunctionType,
			testEmulatorBackendTypeServiceAccountFunctionDocString,
		),
		sema.NewUnmeteredPublicFunctionMember(
			compositeType,
			testEmulatorBackendTypeEventsFunctionName,
			eventsFunctionType,
			testEmulatorBackendTypeEventsFunctionDocString,
		),
		sema.NewUnmeteredPublicFunctionMember(
			compositeType,
			testEmulatorBackendTypeResetFunctionName,
			resetFunctionType,
			testEmulatorBackendTypeResetFunctionDocString,
		),
	}

	compositeType.Members = sema.MembersAsMap(members)
	compositeType.Fields = sema.MembersFieldNames(members)

	return &testEmulatorBackendType{
		compositeType:                      compositeType,
		executeScriptFunctionType:          executeScriptFunctionType,
		createAccountFunctionType:          createAccountFunctionType,
		addTransactionFunctionType:         addTransactionFunctionType,
		executeNextTransactionFunctionType: executeNextTransactionFunctionType,
		commitBlockFunctionType:            commitBlockFunctionType,
		deployContractFunctionType:         deployContractFunctionType,
		useConfigFunctionType:              useConfigFunctionType,
		logsFunctionType:                   logsFunctionType,
		serviceAccountFunctionType:         serviceAccountFunctionType,
		eventsFunctionType:                 eventsFunctionType,
		resetFunctionType:                  resetFunctionType,
	}
}

// 'EmulatorBackend.executeScript' function

const testEmulatorBackendTypeExecuteScriptFunctionName = "executeScript"

const testEmulatorBackendTypeExecuteScriptFunctionDocString = `
Executes a script and returns the script return value and the status.
The 'returnValue' field of the result will be nil if the script failed.
`

func (t *testEmulatorBackendType) newExecuteScriptFunction(testFramework TestFramework) *interpreter.HostFunctionValue {
	return interpreter.NewUnmeteredHostFunctionValue(
		t.executeScriptFunctionType,
		func(invocation interpreter.Invocation) interpreter.Value {
			inter := invocation.Interpreter

			script, ok := invocation.Arguments[0].(*interpreter.StringValue)
			if !ok {
				panic(errors.NewUnreachableError())
			}

			args, err := arrayValueToSlice(inter, invocation.Arguments[1])
			if err != nil {
				panic(errors.NewUnexpectedErrorFromCause(err))
			}

			result := testFramework.RunScript(inter, script.Str, args)

			return newScriptResult(inter, result.Value, result)
		},
	)
}

// 'EmulatorBackend.createAccount' function

const testEmulatorBackendTypeCreateAccountFunctionName = "createAccount"

const testEmulatorBackendTypeCreateAccountFunctionDocString = `
Creates an account by submitting an account creation transaction.
The transaction is paid by the service account.
The returned account can be used to sign and authorize transactions.
`

func (t *testEmulatorBackendType) newCreateAccountFunction(testFramework TestFramework) *interpreter.HostFunctionValue {
	return interpreter.NewUnmeteredHostFunctionValue(
		t.createAccountFunctionType,
		func(invocation interpreter.Invocation) interpreter.Value {
			account, err := testFramework.CreateAccount()
			if err != nil {
				panic(err)
			}

			inter := invocation.Interpreter
			locationRange := invocation.LocationRange

			return newTestAccountValue(
				testFramework,
				inter,
				locationRange,
				account,
			)
		},
	)
}

func newTestAccountValue(
	framework TestFramework,
	inter *interpreter.Interpreter,
	locationRange interpreter.LocationRange,
	account *Account,
) interpreter.Value {

	// Create address value
	address := interpreter.NewAddressValue(nil, account.Address)

	standardLibraryHandler := framework.StandardLibraryHandler()

	publicKey := NewPublicKeyValue(
		inter,
		locationRange,
		account.PublicKey,
		standardLibraryHandler,
		standardLibraryHandler,
	)

	// Create an 'Account' by calling its constructor.
	accountConstructor := getConstructor(inter, testAccountTypeName)
	accountValue, err := inter.InvokeExternally(
		accountConstructor,
		accountConstructor.Type,
		[]interpreter.Value{
			address,
			publicKey,
		},
	)

	if err != nil {
		panic(err)
	}

	return accountValue
}

// 'EmulatorBackend.addTransaction' function

const testEmulatorBackendTypeAddTransactionFunctionName = "addTransaction"

const testEmulatorBackendTypeAddTransactionFunctionDocString = `
Add a transaction to the current block.
`

const testTransactionTypeCodeFieldName = "code"
const testTransactionTypeAuthorizersFieldName = "authorizers"
const testTransactionTypeSignersFieldName = "signers"
const testTransactionTypeArgumentsFieldName = "arguments"

func (t *testEmulatorBackendType) newAddTransactionFunction(testFramework TestFramework) *interpreter.HostFunctionValue {
	return interpreter.NewUnmeteredHostFunctionValue(
		t.addTransactionFunctionType,
		func(invocation interpreter.Invocation) interpreter.Value {
			inter := invocation.Interpreter
			locationRange := invocation.LocationRange

			transactionValue, ok := invocation.Arguments[0].(interpreter.MemberAccessibleValue)
			if !ok {
				panic(errors.NewUnreachableError())
			}

			// Get transaction code
			codeValue := transactionValue.GetMember(
				inter,
				locationRange,
				testTransactionTypeCodeFieldName,
			)
			code, ok := codeValue.(*interpreter.StringValue)
			if !ok {
				panic(errors.NewUnreachableError())
			}

			// Get authorizers
			authorizerValue := transactionValue.GetMember(
				inter,
				locationRange,
				testTransactionTypeAuthorizersFieldName,
			)

			authorizers := addressArrayValueToSlice(inter, authorizerValue)

			// Get signers
			signersValue := transactionValue.GetMember(
				inter,
				locationRange,
				testTransactionTypeSignersFieldName,
			)

			signerAccounts := accountsArrayValueToSlice(
				inter,
				signersValue,
				locationRange,
			)

			// Get arguments
			argsValue := transactionValue.GetMember(
				inter,
				locationRange,
				testTransactionTypeArgumentsFieldName,
			)
			args, err := arrayValueToSlice(inter, argsValue)
			if err != nil {
				panic(errors.NewUnexpectedErrorFromCause(err))
			}

			err = testFramework.AddTransaction(
				inter,
				code.Str,
				authorizers,
				signerAccounts,
				args,
			)

			if err != nil {
				panic(err)
			}

			return interpreter.Void
		},
	)
}

// 'EmulatorBackend.executeNextTransaction' function

const testEmulatorBackendTypeExecuteNextTransactionFunctionName = "executeNextTransaction"

const testEmulatorBackendTypeExecuteNextTransactionFunctionDocString = `
Executes the next transaction in the block, if any.
Returns the result of the transaction, or nil if no transaction was scheduled.
`

func (t *testEmulatorBackendType) newExecuteNextTransactionFunction(testFramework TestFramework) *interpreter.HostFunctionValue {
	return interpreter.NewUnmeteredHostFunctionValue(
		t.executeNextTransactionFunctionType,
		func(invocation interpreter.Invocation) interpreter.Value {
			result := testFramework.ExecuteNextTransaction()

			// If there are no transactions to run, then return `nil`.
			if result == nil {
				return interpreter.Nil
			}

			return newTransactionResult(invocation.Interpreter, result)
		},
	)
}

// 'EmulatorBackend.commitBlock' function

const testEmulatorBackendTypeCommitBlockFunctionName = "commitBlock"

const testEmulatorBackendTypeCommitBlockFunctionDocString = `
Commit the current block. Committing will fail if there are un-executed transactions in the block.
`

func (t *testEmulatorBackendType) newCommitBlockFunction(testFramework TestFramework) *interpreter.HostFunctionValue {
	return interpreter.NewUnmeteredHostFunctionValue(
		t.commitBlockFunctionType,
		func(invocation interpreter.Invocation) interpreter.Value {
			err := testFramework.CommitBlock()
			if err != nil {
				panic(err)
			}

			return interpreter.Void
		},
	)
}

// 'EmulatorBackend.deployContract' function

const testEmulatorBackendTypeDeployContractFunctionName = "deployContract"

const testEmulatorBackendTypeDeployContractFunctionDocString = `
Deploys a given contract, and initializes it with the provided arguments.
`

func (t *testEmulatorBackendType) newDeployContractFunction(testFramework TestFramework) *interpreter.HostFunctionValue {
	return interpreter.NewUnmeteredHostFunctionValue(
		t.deployContractFunctionType,
		func(invocation interpreter.Invocation) interpreter.Value {
			inter := invocation.Interpreter

			// Contract name
			name, ok := invocation.Arguments[0].(*interpreter.StringValue)
			if !ok {
				panic(errors.NewUnreachableError())
			}

			// Contract code
			code, ok := invocation.Arguments[1].(*interpreter.StringValue)
			if !ok {
				panic(errors.NewUnreachableError())
			}

			// authorizer
			accountValue, ok := invocation.Arguments[2].(interpreter.MemberAccessibleValue)
			if !ok {
				panic(errors.NewUnreachableError())
			}

			account := accountFromValue(inter, accountValue, invocation.LocationRange)

			// Contract init arguments
			args, err := arrayValueToSlice(inter, invocation.Arguments[3])
			if err != nil {
				panic(err)
			}

			err = testFramework.DeployContract(
				inter,
				name.Str,
				code.Str,
				account,
				args,
			)

			return newErrorValue(inter, err)
		},
	)
}

// 'EmulatorBackend.useConfiguration' function

const testEmulatorBackendTypeUseConfigFunctionName = "useConfiguration"

const testEmulatorBackendTypeUseConfigFunctionDocString = `
Set the configuration to be used by the blockchain.
Overrides any existing configuration.
`

func (t *testEmulatorBackendType) newUseConfigFunction(testFramework TestFramework) *interpreter.HostFunctionValue {
	return interpreter.NewUnmeteredHostFunctionValue(
		t.useConfigFunctionType,
		func(invocation interpreter.Invocation) interpreter.Value {
			inter := invocation.Interpreter

			// configurations
			configsValue, ok := invocation.Arguments[0].(*interpreter.CompositeValue)
			if !ok {
				panic(errors.NewUnreachableError())
			}

			addresses, ok := configsValue.GetMember(
				inter,
				invocation.LocationRange,
				addressesFieldName,
			).(*interpreter.DictionaryValue)
			if !ok {
				panic(errors.NewUnreachableError())
			}

			mapping := make(map[string]common.Address, addresses.Count())

			addresses.Iterate(inter, func(locationValue, addressValue interpreter.Value) bool {
				location, ok := locationValue.(*interpreter.StringValue)
				if !ok {
					panic(errors.NewUnreachableError())
				}

				address, ok := addressValue.(interpreter.AddressValue)
				if !ok {
					panic(errors.NewUnreachableError())
				}

				mapping[location.Str] = common.Address(address)

				return true
			})

			testFramework.UseConfiguration(&Configuration{
				Addresses: mapping,
			})

			return interpreter.Void
		},
	)
}

// 'EmulatorBackend.logs' function

const testEmulatorBackendTypeLogsFunctionName = "logs"

const testEmulatorBackendTypeLogsFunctionDocString = `
Returns all the logs from the blockchain, up to the calling point.
`

func (t *testEmulatorBackendType) newLogsFunction(
	testFramework TestFramework,
) *interpreter.HostFunctionValue {
	return interpreter.NewUnmeteredHostFunctionValue(
		t.logsFunctionType,
		func(invocation interpreter.Invocation) interpreter.Value {
			logs := testFramework.Logs()
			inter := invocation.Interpreter

			arrayType := interpreter.NewVariableSizedStaticType(
				inter,
				interpreter.NewPrimitiveStaticType(
					inter,
					interpreter.PrimitiveStaticTypeString,
				),
			)

			values := make([]interpreter.Value, len(logs))
			for i, log := range logs {
				memoryUsage := common.NewStringMemoryUsage(len(log))
				values[i] = interpreter.NewStringValue(
					inter,
					memoryUsage,
					func() string {
						return log
					},
				)
			}

			return interpreter.NewArrayValue(
				inter,
				invocation.LocationRange,
				arrayType,
				common.ZeroAddress,
				values...,
			)
		},
	)
}

// 'EmulatorBackend.serviceAccount' function

const testEmulatorBackendTypeServiceAccountFunctionName = "serviceAccount"

const testEmulatorBackendTypeServiceAccountFunctionDocString = `
Returns the service account of the blockchain. Can be used to sign
transactions with this account.
`

func (t *testEmulatorBackendType) newServiceAccountFunction(
	testFramework TestFramework,
) *interpreter.HostFunctionValue {
	return interpreter.NewUnmeteredHostFunctionValue(
		t.serviceAccountFunctionType,
		func(invocation interpreter.Invocation) interpreter.Value {
			serviceAccount, err := testFramework.ServiceAccount()
			if err != nil {
				panic(err)
			}

			return newTestAccountValue(
				testFramework,
				invocation.Interpreter,
				invocation.LocationRange,
				serviceAccount,
			)
		},
	)
}

// 'EmulatorBackend.events' function

const testEmulatorBackendTypeEventsFunctionName = "events"

const testEmulatorBackendTypeEventsFunctionDocString = `
Returns all events emitted from the blockchain,
optionally filtered by event type.
`

func (t *testEmulatorBackendType) newEventsFunction(
	testFramework TestFramework,
) *interpreter.HostFunctionValue {
	return interpreter.NewUnmeteredHostFunctionValue(
		t.eventsFunctionType,
		func(invocation interpreter.Invocation) interpreter.Value {
			var eventType interpreter.StaticType = nil

			switch value := invocation.Arguments[0].(type) {
			case interpreter.NilValue:
				// Do nothing
			case *interpreter.SomeValue:
				innerValue := value.InnerValue(invocation.Interpreter, invocation.LocationRange)
				typeValue, ok := innerValue.(interpreter.TypeValue)
				if !ok {
					panic(errors.NewUnreachableError())
				}

				eventType = typeValue.Type
			default:
				panic(errors.NewUnreachableError())
			}

			return testFramework.Events(invocation.Interpreter, eventType)
		},
	)
}

// 'EmulatorBackend.reset' function

const testEmulatorBackendTypeResetFunctionName = "reset"

const testEmulatorBackendTypeResetFunctionDocString = `
Resets the state of the blockchain.
`

func (t *testEmulatorBackendType) newResetFunction(
	testFramework TestFramework,
) *interpreter.HostFunctionValue {
	return interpreter.NewUnmeteredHostFunctionValue(
		t.eventsFunctionType,
		func(invocation interpreter.Invocation) interpreter.Value {
			testFramework.Reset()
			return interpreter.Void
		},
	)
}

func (t *testEmulatorBackendType) newEmulatorBackend(
	inter *interpreter.Interpreter,
	testFramework TestFramework,
	locationRange interpreter.LocationRange,
) *interpreter.CompositeValue {
	var fields = []interpreter.CompositeField{
		{
			Name:  testEmulatorBackendTypeExecuteScriptFunctionName,
			Value: t.newExecuteScriptFunction(testFramework),
		},
		{
			Name:  testEmulatorBackendTypeCreateAccountFunctionName,
			Value: t.newCreateAccountFunction(testFramework),
		}, {
			Name:  testEmulatorBackendTypeAddTransactionFunctionName,
			Value: t.newAddTransactionFunction(testFramework),
		},
		{
			Name:  testEmulatorBackendTypeExecuteNextTransactionFunctionName,
			Value: t.newExecuteNextTransactionFunction(testFramework),
		},
		{
			Name:  testEmulatorBackendTypeCommitBlockFunctionName,
			Value: t.newCommitBlockFunction(testFramework),
		},
		{
			Name:  testEmulatorBackendTypeDeployContractFunctionName,
			Value: t.newDeployContractFunction(testFramework),
		},
		{
			Name:  testEmulatorBackendTypeUseConfigFunctionName,
			Value: t.newUseConfigFunction(testFramework),
		},
		{
			Name:  testEmulatorBackendTypeLogsFunctionName,
			Value: t.newLogsFunction(testFramework),
		},
		{
			Name:  testEmulatorBackendTypeServiceAccountFunctionName,
			Value: t.newServiceAccountFunction(testFramework),
		},
		{
			Name:  testEmulatorBackendTypeEventsFunctionName,
			Value: t.newEventsFunction(testFramework),
		},
		{
			Name:  testEmulatorBackendTypeResetFunctionName,
			Value: t.newResetFunction(testFramework),
		},
	}

	// TODO: Use SimpleCompositeValue
	return interpreter.NewCompositeValue(
		inter,
		locationRange,
		t.compositeType.Location,
		testEmulatorBackendTypeName,
		common.CompositeKindStructure,
		fields,
		common.ZeroAddress,
	)
}
