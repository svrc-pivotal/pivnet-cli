// This file was generated by counterfeiter
package rcfakes

import (
	"sync"

	"github.com/pivotal-cf/pivnet-cli/rc"
)

type FakePivnetRCReadWriter struct {
	WriteToFileStub        func(contents []byte) error
	writeToFileMutex       sync.RWMutex
	writeToFileArgsForCall []struct {
		contents []byte
	}
	writeToFileReturns struct {
		result1 error
	}
	ReadFromFileStub        func() ([]byte, error)
	readFromFileMutex       sync.RWMutex
	readFromFileArgsForCall []struct{}
	readFromFileReturns     struct {
		result1 []byte
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakePivnetRCReadWriter) WriteToFile(contents []byte) error {
	var contentsCopy []byte
	if contents != nil {
		contentsCopy = make([]byte, len(contents))
		copy(contentsCopy, contents)
	}
	fake.writeToFileMutex.Lock()
	fake.writeToFileArgsForCall = append(fake.writeToFileArgsForCall, struct {
		contents []byte
	}{contentsCopy})
	fake.recordInvocation("WriteToFile", []interface{}{contentsCopy})
	fake.writeToFileMutex.Unlock()
	if fake.WriteToFileStub != nil {
		return fake.WriteToFileStub(contents)
	} else {
		return fake.writeToFileReturns.result1
	}
}

func (fake *FakePivnetRCReadWriter) WriteToFileCallCount() int {
	fake.writeToFileMutex.RLock()
	defer fake.writeToFileMutex.RUnlock()
	return len(fake.writeToFileArgsForCall)
}

func (fake *FakePivnetRCReadWriter) WriteToFileArgsForCall(i int) []byte {
	fake.writeToFileMutex.RLock()
	defer fake.writeToFileMutex.RUnlock()
	return fake.writeToFileArgsForCall[i].contents
}

func (fake *FakePivnetRCReadWriter) WriteToFileReturns(result1 error) {
	fake.WriteToFileStub = nil
	fake.writeToFileReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakePivnetRCReadWriter) ReadFromFile() ([]byte, error) {
	fake.readFromFileMutex.Lock()
	fake.readFromFileArgsForCall = append(fake.readFromFileArgsForCall, struct{}{})
	fake.recordInvocation("ReadFromFile", []interface{}{})
	fake.readFromFileMutex.Unlock()
	if fake.ReadFromFileStub != nil {
		return fake.ReadFromFileStub()
	} else {
		return fake.readFromFileReturns.result1, fake.readFromFileReturns.result2
	}
}

func (fake *FakePivnetRCReadWriter) ReadFromFileCallCount() int {
	fake.readFromFileMutex.RLock()
	defer fake.readFromFileMutex.RUnlock()
	return len(fake.readFromFileArgsForCall)
}

func (fake *FakePivnetRCReadWriter) ReadFromFileReturns(result1 []byte, result2 error) {
	fake.ReadFromFileStub = nil
	fake.readFromFileReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetRCReadWriter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.writeToFileMutex.RLock()
	defer fake.writeToFileMutex.RUnlock()
	fake.readFromFileMutex.RLock()
	defer fake.readFromFileMutex.RUnlock()
	return fake.invocations
}

func (fake *FakePivnetRCReadWriter) recordInvocation(key string, args []interface{}) {
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

var _ rc.PivnetRCReadWriter = new(FakePivnetRCReadWriter)
