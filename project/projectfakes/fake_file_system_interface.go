// This file was generated by counterfeiter
package projectfakes

import (
	"sync"

	"github.com/odlp/inflight/project"
)

type FakeFileSystemInterface struct {
	WriteToFileStub        func(filePath string, text string)
	writeToFileMutex       sync.RWMutex
	writeToFileArgsForCall []struct {
		filePath string
		text     string
	}
	ReadFromFileStub        func(filePath string) ([]byte, error)
	readFromFileMutex       sync.RWMutex
	readFromFileArgsForCall []struct {
		filePath string
	}
	readFromFileReturns struct {
		result1 []byte
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeFileSystemInterface) WriteToFile(filePath string, text string) {
	fake.writeToFileMutex.Lock()
	fake.writeToFileArgsForCall = append(fake.writeToFileArgsForCall, struct {
		filePath string
		text     string
	}{filePath, text})
	fake.recordInvocation("WriteToFile", []interface{}{filePath, text})
	fake.writeToFileMutex.Unlock()
	if fake.WriteToFileStub != nil {
		fake.WriteToFileStub(filePath, text)
	}
}

func (fake *FakeFileSystemInterface) WriteToFileCallCount() int {
	fake.writeToFileMutex.RLock()
	defer fake.writeToFileMutex.RUnlock()
	return len(fake.writeToFileArgsForCall)
}

func (fake *FakeFileSystemInterface) WriteToFileArgsForCall(i int) (string, string) {
	fake.writeToFileMutex.RLock()
	defer fake.writeToFileMutex.RUnlock()
	return fake.writeToFileArgsForCall[i].filePath, fake.writeToFileArgsForCall[i].text
}

func (fake *FakeFileSystemInterface) ReadFromFile(filePath string) ([]byte, error) {
	fake.readFromFileMutex.Lock()
	fake.readFromFileArgsForCall = append(fake.readFromFileArgsForCall, struct {
		filePath string
	}{filePath})
	fake.recordInvocation("ReadFromFile", []interface{}{filePath})
	fake.readFromFileMutex.Unlock()
	if fake.ReadFromFileStub != nil {
		return fake.ReadFromFileStub(filePath)
	} else {
		return fake.readFromFileReturns.result1, fake.readFromFileReturns.result2
	}
}

func (fake *FakeFileSystemInterface) ReadFromFileCallCount() int {
	fake.readFromFileMutex.RLock()
	defer fake.readFromFileMutex.RUnlock()
	return len(fake.readFromFileArgsForCall)
}

func (fake *FakeFileSystemInterface) ReadFromFileArgsForCall(i int) string {
	fake.readFromFileMutex.RLock()
	defer fake.readFromFileMutex.RUnlock()
	return fake.readFromFileArgsForCall[i].filePath
}

func (fake *FakeFileSystemInterface) ReadFromFileReturns(result1 []byte, result2 error) {
	fake.ReadFromFileStub = nil
	fake.readFromFileReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeFileSystemInterface) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.writeToFileMutex.RLock()
	defer fake.writeToFileMutex.RUnlock()
	fake.readFromFileMutex.RLock()
	defer fake.readFromFileMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeFileSystemInterface) recordInvocation(key string, args []interface{}) {
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

var _ project.FileSystemInterface = new(FakeFileSystemInterface)