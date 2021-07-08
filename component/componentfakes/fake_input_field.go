// Code generated by counterfeiter. DO NOT EDIT.
package componentfakes

import (
	"sync"

	tcell "github.com/gdamore/tcell/v2"
	"github.com/hcjulz/damon/component"
	"github.com/rivo/tview"
)

type FakeInputField struct {
	GetTextStub        func() string
	getTextMutex       sync.RWMutex
	getTextArgsForCall []struct {
	}
	getTextReturns struct {
		result1 string
	}
	getTextReturnsOnCall map[int]struct {
		result1 string
	}
	PrimitiveStub        func() tview.Primitive
	primitiveMutex       sync.RWMutex
	primitiveArgsForCall []struct {
	}
	primitiveReturns struct {
		result1 tview.Primitive
	}
	primitiveReturnsOnCall map[int]struct {
		result1 tview.Primitive
	}
	SetAutocompleteFuncStub        func(func(currentText string) (entries []string))
	setAutocompleteFuncMutex       sync.RWMutex
	setAutocompleteFuncArgsForCall []struct {
		arg1 func(currentText string) (entries []string)
	}
	SetChangedFuncStub        func(func(text string))
	setChangedFuncMutex       sync.RWMutex
	setChangedFuncArgsForCall []struct {
		arg1 func(text string)
	}
	SetDoneFuncStub        func(func(k tcell.Key))
	setDoneFuncMutex       sync.RWMutex
	setDoneFuncArgsForCall []struct {
		arg1 func(k tcell.Key)
	}
	SetTextStub        func(string)
	setTextMutex       sync.RWMutex
	setTextArgsForCall []struct {
		arg1 string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeInputField) GetText() string {
	fake.getTextMutex.Lock()
	ret, specificReturn := fake.getTextReturnsOnCall[len(fake.getTextArgsForCall)]
	fake.getTextArgsForCall = append(fake.getTextArgsForCall, struct {
	}{})
	stub := fake.GetTextStub
	fakeReturns := fake.getTextReturns
	fake.recordInvocation("GetText", []interface{}{})
	fake.getTextMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeInputField) GetTextCallCount() int {
	fake.getTextMutex.RLock()
	defer fake.getTextMutex.RUnlock()
	return len(fake.getTextArgsForCall)
}

func (fake *FakeInputField) GetTextCalls(stub func() string) {
	fake.getTextMutex.Lock()
	defer fake.getTextMutex.Unlock()
	fake.GetTextStub = stub
}

func (fake *FakeInputField) GetTextReturns(result1 string) {
	fake.getTextMutex.Lock()
	defer fake.getTextMutex.Unlock()
	fake.GetTextStub = nil
	fake.getTextReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeInputField) GetTextReturnsOnCall(i int, result1 string) {
	fake.getTextMutex.Lock()
	defer fake.getTextMutex.Unlock()
	fake.GetTextStub = nil
	if fake.getTextReturnsOnCall == nil {
		fake.getTextReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.getTextReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeInputField) Primitive() tview.Primitive {
	fake.primitiveMutex.Lock()
	ret, specificReturn := fake.primitiveReturnsOnCall[len(fake.primitiveArgsForCall)]
	fake.primitiveArgsForCall = append(fake.primitiveArgsForCall, struct {
	}{})
	stub := fake.PrimitiveStub
	fakeReturns := fake.primitiveReturns
	fake.recordInvocation("Primitive", []interface{}{})
	fake.primitiveMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeInputField) PrimitiveCallCount() int {
	fake.primitiveMutex.RLock()
	defer fake.primitiveMutex.RUnlock()
	return len(fake.primitiveArgsForCall)
}

func (fake *FakeInputField) PrimitiveCalls(stub func() tview.Primitive) {
	fake.primitiveMutex.Lock()
	defer fake.primitiveMutex.Unlock()
	fake.PrimitiveStub = stub
}

func (fake *FakeInputField) PrimitiveReturns(result1 tview.Primitive) {
	fake.primitiveMutex.Lock()
	defer fake.primitiveMutex.Unlock()
	fake.PrimitiveStub = nil
	fake.primitiveReturns = struct {
		result1 tview.Primitive
	}{result1}
}

func (fake *FakeInputField) PrimitiveReturnsOnCall(i int, result1 tview.Primitive) {
	fake.primitiveMutex.Lock()
	defer fake.primitiveMutex.Unlock()
	fake.PrimitiveStub = nil
	if fake.primitiveReturnsOnCall == nil {
		fake.primitiveReturnsOnCall = make(map[int]struct {
			result1 tview.Primitive
		})
	}
	fake.primitiveReturnsOnCall[i] = struct {
		result1 tview.Primitive
	}{result1}
}

func (fake *FakeInputField) SetAutocompleteFunc(arg1 func(currentText string) (entries []string)) {
	fake.setAutocompleteFuncMutex.Lock()
	fake.setAutocompleteFuncArgsForCall = append(fake.setAutocompleteFuncArgsForCall, struct {
		arg1 func(currentText string) (entries []string)
	}{arg1})
	stub := fake.SetAutocompleteFuncStub
	fake.recordInvocation("SetAutocompleteFunc", []interface{}{arg1})
	fake.setAutocompleteFuncMutex.Unlock()
	if stub != nil {
		fake.SetAutocompleteFuncStub(arg1)
	}
}

func (fake *FakeInputField) SetAutocompleteFuncCallCount() int {
	fake.setAutocompleteFuncMutex.RLock()
	defer fake.setAutocompleteFuncMutex.RUnlock()
	return len(fake.setAutocompleteFuncArgsForCall)
}

func (fake *FakeInputField) SetAutocompleteFuncCalls(stub func(func(currentText string) (entries []string))) {
	fake.setAutocompleteFuncMutex.Lock()
	defer fake.setAutocompleteFuncMutex.Unlock()
	fake.SetAutocompleteFuncStub = stub
}

func (fake *FakeInputField) SetAutocompleteFuncArgsForCall(i int) func(currentText string) (entries []string) {
	fake.setAutocompleteFuncMutex.RLock()
	defer fake.setAutocompleteFuncMutex.RUnlock()
	argsForCall := fake.setAutocompleteFuncArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeInputField) SetChangedFunc(arg1 func(text string)) {
	fake.setChangedFuncMutex.Lock()
	fake.setChangedFuncArgsForCall = append(fake.setChangedFuncArgsForCall, struct {
		arg1 func(text string)
	}{arg1})
	stub := fake.SetChangedFuncStub
	fake.recordInvocation("SetChangedFunc", []interface{}{arg1})
	fake.setChangedFuncMutex.Unlock()
	if stub != nil {
		fake.SetChangedFuncStub(arg1)
	}
}

func (fake *FakeInputField) SetChangedFuncCallCount() int {
	fake.setChangedFuncMutex.RLock()
	defer fake.setChangedFuncMutex.RUnlock()
	return len(fake.setChangedFuncArgsForCall)
}

func (fake *FakeInputField) SetChangedFuncCalls(stub func(func(text string))) {
	fake.setChangedFuncMutex.Lock()
	defer fake.setChangedFuncMutex.Unlock()
	fake.SetChangedFuncStub = stub
}

func (fake *FakeInputField) SetChangedFuncArgsForCall(i int) func(text string) {
	fake.setChangedFuncMutex.RLock()
	defer fake.setChangedFuncMutex.RUnlock()
	argsForCall := fake.setChangedFuncArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeInputField) SetDoneFunc(arg1 func(k tcell.Key)) {
	fake.setDoneFuncMutex.Lock()
	fake.setDoneFuncArgsForCall = append(fake.setDoneFuncArgsForCall, struct {
		arg1 func(k tcell.Key)
	}{arg1})
	stub := fake.SetDoneFuncStub
	fake.recordInvocation("SetDoneFunc", []interface{}{arg1})
	fake.setDoneFuncMutex.Unlock()
	if stub != nil {
		fake.SetDoneFuncStub(arg1)
	}
}

func (fake *FakeInputField) SetDoneFuncCallCount() int {
	fake.setDoneFuncMutex.RLock()
	defer fake.setDoneFuncMutex.RUnlock()
	return len(fake.setDoneFuncArgsForCall)
}

func (fake *FakeInputField) SetDoneFuncCalls(stub func(func(k tcell.Key))) {
	fake.setDoneFuncMutex.Lock()
	defer fake.setDoneFuncMutex.Unlock()
	fake.SetDoneFuncStub = stub
}

func (fake *FakeInputField) SetDoneFuncArgsForCall(i int) func(k tcell.Key) {
	fake.setDoneFuncMutex.RLock()
	defer fake.setDoneFuncMutex.RUnlock()
	argsForCall := fake.setDoneFuncArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeInputField) SetText(arg1 string) {
	fake.setTextMutex.Lock()
	fake.setTextArgsForCall = append(fake.setTextArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.SetTextStub
	fake.recordInvocation("SetText", []interface{}{arg1})
	fake.setTextMutex.Unlock()
	if stub != nil {
		fake.SetTextStub(arg1)
	}
}

func (fake *FakeInputField) SetTextCallCount() int {
	fake.setTextMutex.RLock()
	defer fake.setTextMutex.RUnlock()
	return len(fake.setTextArgsForCall)
}

func (fake *FakeInputField) SetTextCalls(stub func(string)) {
	fake.setTextMutex.Lock()
	defer fake.setTextMutex.Unlock()
	fake.SetTextStub = stub
}

func (fake *FakeInputField) SetTextArgsForCall(i int) string {
	fake.setTextMutex.RLock()
	defer fake.setTextMutex.RUnlock()
	argsForCall := fake.setTextArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeInputField) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getTextMutex.RLock()
	defer fake.getTextMutex.RUnlock()
	fake.primitiveMutex.RLock()
	defer fake.primitiveMutex.RUnlock()
	fake.setAutocompleteFuncMutex.RLock()
	defer fake.setAutocompleteFuncMutex.RUnlock()
	fake.setChangedFuncMutex.RLock()
	defer fake.setChangedFuncMutex.RUnlock()
	fake.setDoneFuncMutex.RLock()
	defer fake.setDoneFuncMutex.RUnlock()
	fake.setTextMutex.RLock()
	defer fake.setTextMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeInputField) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ component.InputField = new(FakeInputField)
