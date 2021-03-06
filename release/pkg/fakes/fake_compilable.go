// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/cloudfoundry/bosh-cli/release/pkg"
)

type FakeCompilable struct {
	NameStub        func() string
	nameMutex       sync.RWMutex
	nameArgsForCall []struct{}
	nameReturns     struct {
		result1 string
	}
	FingerprintStub        func() string
	fingerprintMutex       sync.RWMutex
	fingerprintArgsForCall []struct{}
	fingerprintReturns     struct {
		result1 string
	}
	ArchivePathStub        func() string
	archivePathMutex       sync.RWMutex
	archivePathArgsForCall []struct{}
	archivePathReturns     struct {
		result1 string
	}
	ArchiveSHA1Stub        func() string
	archiveSHA1Mutex       sync.RWMutex
	archiveSHA1ArgsForCall []struct{}
	archiveSHA1Returns     struct {
		result1 string
	}
	IsCompiledStub        func() bool
	isCompiledMutex       sync.RWMutex
	isCompiledArgsForCall []struct{}
	isCompiledReturns     struct {
		result1 bool
	}
	DepsStub        func() []pkg.Compilable
	depsMutex       sync.RWMutex
	depsArgsForCall []struct{}
	depsReturns     struct {
		result1 []pkg.Compilable
	}
}

func (fake *FakeCompilable) Name() string {
	fake.nameMutex.Lock()
	fake.nameArgsForCall = append(fake.nameArgsForCall, struct{}{})
	fake.nameMutex.Unlock()
	if fake.NameStub != nil {
		return fake.NameStub()
	} else {
		return fake.nameReturns.result1
	}
}

func (fake *FakeCompilable) NameCallCount() int {
	fake.nameMutex.RLock()
	defer fake.nameMutex.RUnlock()
	return len(fake.nameArgsForCall)
}

func (fake *FakeCompilable) NameReturns(result1 string) {
	fake.NameStub = nil
	fake.nameReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeCompilable) Fingerprint() string {
	fake.fingerprintMutex.Lock()
	fake.fingerprintArgsForCall = append(fake.fingerprintArgsForCall, struct{}{})
	fake.fingerprintMutex.Unlock()
	if fake.FingerprintStub != nil {
		return fake.FingerprintStub()
	} else {
		return fake.fingerprintReturns.result1
	}
}

func (fake *FakeCompilable) FingerprintCallCount() int {
	fake.fingerprintMutex.RLock()
	defer fake.fingerprintMutex.RUnlock()
	return len(fake.fingerprintArgsForCall)
}

func (fake *FakeCompilable) FingerprintReturns(result1 string) {
	fake.FingerprintStub = nil
	fake.fingerprintReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeCompilable) ArchivePath() string {
	fake.archivePathMutex.Lock()
	fake.archivePathArgsForCall = append(fake.archivePathArgsForCall, struct{}{})
	fake.archivePathMutex.Unlock()
	if fake.ArchivePathStub != nil {
		return fake.ArchivePathStub()
	} else {
		return fake.archivePathReturns.result1
	}
}

func (fake *FakeCompilable) ArchivePathCallCount() int {
	fake.archivePathMutex.RLock()
	defer fake.archivePathMutex.RUnlock()
	return len(fake.archivePathArgsForCall)
}

func (fake *FakeCompilable) ArchivePathReturns(result1 string) {
	fake.ArchivePathStub = nil
	fake.archivePathReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeCompilable) ArchiveSHA1() string {
	fake.archiveSHA1Mutex.Lock()
	fake.archiveSHA1ArgsForCall = append(fake.archiveSHA1ArgsForCall, struct{}{})
	fake.archiveSHA1Mutex.Unlock()
	if fake.ArchiveSHA1Stub != nil {
		return fake.ArchiveSHA1Stub()
	} else {
		return fake.archiveSHA1Returns.result1
	}
}

func (fake *FakeCompilable) ArchiveSHA1CallCount() int {
	fake.archiveSHA1Mutex.RLock()
	defer fake.archiveSHA1Mutex.RUnlock()
	return len(fake.archiveSHA1ArgsForCall)
}

func (fake *FakeCompilable) ArchiveSHA1Returns(result1 string) {
	fake.ArchiveSHA1Stub = nil
	fake.archiveSHA1Returns = struct {
		result1 string
	}{result1}
}

func (fake *FakeCompilable) IsCompiled() bool {
	fake.isCompiledMutex.Lock()
	fake.isCompiledArgsForCall = append(fake.isCompiledArgsForCall, struct{}{})
	fake.isCompiledMutex.Unlock()
	if fake.IsCompiledStub != nil {
		return fake.IsCompiledStub()
	} else {
		return fake.isCompiledReturns.result1
	}
}

func (fake *FakeCompilable) IsCompiledCallCount() int {
	fake.isCompiledMutex.RLock()
	defer fake.isCompiledMutex.RUnlock()
	return len(fake.isCompiledArgsForCall)
}

func (fake *FakeCompilable) IsCompiledReturns(result1 bool) {
	fake.IsCompiledStub = nil
	fake.isCompiledReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeCompilable) Deps() []pkg.Compilable {
	fake.depsMutex.Lock()
	fake.depsArgsForCall = append(fake.depsArgsForCall, struct{}{})
	fake.depsMutex.Unlock()
	if fake.DepsStub != nil {
		return fake.DepsStub()
	} else {
		return fake.depsReturns.result1
	}
}

func (fake *FakeCompilable) DepsCallCount() int {
	fake.depsMutex.RLock()
	defer fake.depsMutex.RUnlock()
	return len(fake.depsArgsForCall)
}

func (fake *FakeCompilable) DepsReturns(result1 []pkg.Compilable) {
	fake.DepsStub = nil
	fake.depsReturns = struct {
		result1 []pkg.Compilable
	}{result1}
}

var _ pkg.Compilable = new(FakeCompilable)
