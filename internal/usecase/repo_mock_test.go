package usecase

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

//go:generate minimock -i word-of-wisdom/internal/usecase.Repo -o ./repo_mock_test.go -n RepoMock

import (
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
)

// RepoMock implements Repo
type RepoMock struct {
	t minimock.Tester

	funcExists          func(hash string) (b1 bool)
	inspectFuncExists   func(hash string)
	afterExistsCounter  uint64
	beforeExistsCounter uint64
	ExistsMock          mRepoMockExists

	funcStore          func(hash string)
	inspectFuncStore   func(hash string)
	afterStoreCounter  uint64
	beforeStoreCounter uint64
	StoreMock          mRepoMockStore
}

// NewRepoMock returns a mock for Repo
func NewRepoMock(t minimock.Tester) *RepoMock {
	m := &RepoMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.ExistsMock = mRepoMockExists{mock: m}
	m.ExistsMock.callArgs = []*RepoMockExistsParams{}

	m.StoreMock = mRepoMockStore{mock: m}
	m.StoreMock.callArgs = []*RepoMockStoreParams{}

	return m
}

type mRepoMockExists struct {
	mock               *RepoMock
	defaultExpectation *RepoMockExistsExpectation
	expectations       []*RepoMockExistsExpectation

	callArgs []*RepoMockExistsParams
	mutex    sync.RWMutex
}

// RepoMockExistsExpectation specifies expectation struct of the Repo.Exists
type RepoMockExistsExpectation struct {
	mock    *RepoMock
	params  *RepoMockExistsParams
	results *RepoMockExistsResults
	Counter uint64
}

// RepoMockExistsParams contains parameters of the Repo.Exists
type RepoMockExistsParams struct {
	hash string
}

// RepoMockExistsResults contains results of the Repo.Exists
type RepoMockExistsResults struct {
	b1 bool
}

// Expect sets up expected params for Repo.Exists
func (mmExists *mRepoMockExists) Expect(hash string) *mRepoMockExists {
	if mmExists.mock.funcExists != nil {
		mmExists.mock.t.Fatalf("RepoMock.Exists mock is already set by Set")
	}

	if mmExists.defaultExpectation == nil {
		mmExists.defaultExpectation = &RepoMockExistsExpectation{}
	}

	mmExists.defaultExpectation.params = &RepoMockExistsParams{hash}
	for _, e := range mmExists.expectations {
		if minimock.Equal(e.params, mmExists.defaultExpectation.params) {
			mmExists.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmExists.defaultExpectation.params)
		}
	}

	return mmExists
}

// Inspect accepts an inspector function that has same arguments as the Repo.Exists
func (mmExists *mRepoMockExists) Inspect(f func(hash string)) *mRepoMockExists {
	if mmExists.mock.inspectFuncExists != nil {
		mmExists.mock.t.Fatalf("Inspect function is already set for RepoMock.Exists")
	}

	mmExists.mock.inspectFuncExists = f

	return mmExists
}

// Return sets up results that will be returned by Repo.Exists
func (mmExists *mRepoMockExists) Return(b1 bool) *RepoMock {
	if mmExists.mock.funcExists != nil {
		mmExists.mock.t.Fatalf("RepoMock.Exists mock is already set by Set")
	}

	if mmExists.defaultExpectation == nil {
		mmExists.defaultExpectation = &RepoMockExistsExpectation{mock: mmExists.mock}
	}
	mmExists.defaultExpectation.results = &RepoMockExistsResults{b1}
	return mmExists.mock
}

//Set uses given function f to mock the Repo.Exists method
func (mmExists *mRepoMockExists) Set(f func(hash string) (b1 bool)) *RepoMock {
	if mmExists.defaultExpectation != nil {
		mmExists.mock.t.Fatalf("Default expectation is already set for the Repo.Exists method")
	}

	if len(mmExists.expectations) > 0 {
		mmExists.mock.t.Fatalf("Some expectations are already set for the Repo.Exists method")
	}

	mmExists.mock.funcExists = f
	return mmExists.mock
}

// When sets expectation for the Repo.Exists which will trigger the result defined by the following
// Then helper
func (mmExists *mRepoMockExists) When(hash string) *RepoMockExistsExpectation {
	if mmExists.mock.funcExists != nil {
		mmExists.mock.t.Fatalf("RepoMock.Exists mock is already set by Set")
	}

	expectation := &RepoMockExistsExpectation{
		mock:   mmExists.mock,
		params: &RepoMockExistsParams{hash},
	}
	mmExists.expectations = append(mmExists.expectations, expectation)
	return expectation
}

// Then sets up Repo.Exists return parameters for the expectation previously defined by the When method
func (e *RepoMockExistsExpectation) Then(b1 bool) *RepoMock {
	e.results = &RepoMockExistsResults{b1}
	return e.mock
}

// Exists implements Repo
func (mmExists *RepoMock) Exists(hash string) (b1 bool) {
	mm_atomic.AddUint64(&mmExists.beforeExistsCounter, 1)
	defer mm_atomic.AddUint64(&mmExists.afterExistsCounter, 1)

	if mmExists.inspectFuncExists != nil {
		mmExists.inspectFuncExists(hash)
	}

	mm_params := &RepoMockExistsParams{hash}

	// Record call args
	mmExists.ExistsMock.mutex.Lock()
	mmExists.ExistsMock.callArgs = append(mmExists.ExistsMock.callArgs, mm_params)
	mmExists.ExistsMock.mutex.Unlock()

	for _, e := range mmExists.ExistsMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.b1
		}
	}

	if mmExists.ExistsMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmExists.ExistsMock.defaultExpectation.Counter, 1)
		mm_want := mmExists.ExistsMock.defaultExpectation.params
		mm_got := RepoMockExistsParams{hash}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmExists.t.Errorf("RepoMock.Exists got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmExists.ExistsMock.defaultExpectation.results
		if mm_results == nil {
			mmExists.t.Fatal("No results are set for the RepoMock.Exists")
		}
		return (*mm_results).b1
	}
	if mmExists.funcExists != nil {
		return mmExists.funcExists(hash)
	}
	mmExists.t.Fatalf("Unexpected call to RepoMock.Exists. %v", hash)
	return
}

// ExistsAfterCounter returns a count of finished RepoMock.Exists invocations
func (mmExists *RepoMock) ExistsAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmExists.afterExistsCounter)
}

// ExistsBeforeCounter returns a count of RepoMock.Exists invocations
func (mmExists *RepoMock) ExistsBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmExists.beforeExistsCounter)
}

// Calls returns a list of arguments used in each call to RepoMock.Exists.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmExists *mRepoMockExists) Calls() []*RepoMockExistsParams {
	mmExists.mutex.RLock()

	argCopy := make([]*RepoMockExistsParams, len(mmExists.callArgs))
	copy(argCopy, mmExists.callArgs)

	mmExists.mutex.RUnlock()

	return argCopy
}

// MinimockExistsDone returns true if the count of the Exists invocations corresponds
// the number of defined expectations
func (m *RepoMock) MinimockExistsDone() bool {
	for _, e := range m.ExistsMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.ExistsMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterExistsCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcExists != nil && mm_atomic.LoadUint64(&m.afterExistsCounter) < 1 {
		return false
	}
	return true
}

// MinimockExistsInspect logs each unmet expectation
func (m *RepoMock) MinimockExistsInspect() {
	for _, e := range m.ExistsMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to RepoMock.Exists with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.ExistsMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterExistsCounter) < 1 {
		if m.ExistsMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to RepoMock.Exists")
		} else {
			m.t.Errorf("Expected call to RepoMock.Exists with params: %#v", *m.ExistsMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcExists != nil && mm_atomic.LoadUint64(&m.afterExistsCounter) < 1 {
		m.t.Error("Expected call to RepoMock.Exists")
	}
}

type mRepoMockStore struct {
	mock               *RepoMock
	defaultExpectation *RepoMockStoreExpectation
	expectations       []*RepoMockStoreExpectation

	callArgs []*RepoMockStoreParams
	mutex    sync.RWMutex
}

// RepoMockStoreExpectation specifies expectation struct of the Repo.Store
type RepoMockStoreExpectation struct {
	mock   *RepoMock
	params *RepoMockStoreParams

	Counter uint64
}

// RepoMockStoreParams contains parameters of the Repo.Store
type RepoMockStoreParams struct {
	hash string
}

// Expect sets up expected params for Repo.Store
func (mmStore *mRepoMockStore) Expect(hash string) *mRepoMockStore {
	if mmStore.mock.funcStore != nil {
		mmStore.mock.t.Fatalf("RepoMock.Store mock is already set by Set")
	}

	if mmStore.defaultExpectation == nil {
		mmStore.defaultExpectation = &RepoMockStoreExpectation{}
	}

	mmStore.defaultExpectation.params = &RepoMockStoreParams{hash}
	for _, e := range mmStore.expectations {
		if minimock.Equal(e.params, mmStore.defaultExpectation.params) {
			mmStore.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmStore.defaultExpectation.params)
		}
	}

	return mmStore
}

// Inspect accepts an inspector function that has same arguments as the Repo.Store
func (mmStore *mRepoMockStore) Inspect(f func(hash string)) *mRepoMockStore {
	if mmStore.mock.inspectFuncStore != nil {
		mmStore.mock.t.Fatalf("Inspect function is already set for RepoMock.Store")
	}

	mmStore.mock.inspectFuncStore = f

	return mmStore
}

// Return sets up results that will be returned by Repo.Store
func (mmStore *mRepoMockStore) Return() *RepoMock {
	if mmStore.mock.funcStore != nil {
		mmStore.mock.t.Fatalf("RepoMock.Store mock is already set by Set")
	}

	if mmStore.defaultExpectation == nil {
		mmStore.defaultExpectation = &RepoMockStoreExpectation{mock: mmStore.mock}
	}

	return mmStore.mock
}

//Set uses given function f to mock the Repo.Store method
func (mmStore *mRepoMockStore) Set(f func(hash string)) *RepoMock {
	if mmStore.defaultExpectation != nil {
		mmStore.mock.t.Fatalf("Default expectation is already set for the Repo.Store method")
	}

	if len(mmStore.expectations) > 0 {
		mmStore.mock.t.Fatalf("Some expectations are already set for the Repo.Store method")
	}

	mmStore.mock.funcStore = f
	return mmStore.mock
}

// Store implements Repo
func (mmStore *RepoMock) Store(hash string) {
	mm_atomic.AddUint64(&mmStore.beforeStoreCounter, 1)
	defer mm_atomic.AddUint64(&mmStore.afterStoreCounter, 1)

	if mmStore.inspectFuncStore != nil {
		mmStore.inspectFuncStore(hash)
	}

	mm_params := &RepoMockStoreParams{hash}

	// Record call args
	mmStore.StoreMock.mutex.Lock()
	mmStore.StoreMock.callArgs = append(mmStore.StoreMock.callArgs, mm_params)
	mmStore.StoreMock.mutex.Unlock()

	for _, e := range mmStore.StoreMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return
		}
	}

	if mmStore.StoreMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmStore.StoreMock.defaultExpectation.Counter, 1)
		mm_want := mmStore.StoreMock.defaultExpectation.params
		mm_got := RepoMockStoreParams{hash}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmStore.t.Errorf("RepoMock.Store got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		return

	}
	if mmStore.funcStore != nil {
		mmStore.funcStore(hash)
		return
	}
	mmStore.t.Fatalf("Unexpected call to RepoMock.Store. %v", hash)

}

// StoreAfterCounter returns a count of finished RepoMock.Store invocations
func (mmStore *RepoMock) StoreAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmStore.afterStoreCounter)
}

// StoreBeforeCounter returns a count of RepoMock.Store invocations
func (mmStore *RepoMock) StoreBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmStore.beforeStoreCounter)
}

// Calls returns a list of arguments used in each call to RepoMock.Store.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmStore *mRepoMockStore) Calls() []*RepoMockStoreParams {
	mmStore.mutex.RLock()

	argCopy := make([]*RepoMockStoreParams, len(mmStore.callArgs))
	copy(argCopy, mmStore.callArgs)

	mmStore.mutex.RUnlock()

	return argCopy
}

// MinimockStoreDone returns true if the count of the Store invocations corresponds
// the number of defined expectations
func (m *RepoMock) MinimockStoreDone() bool {
	for _, e := range m.StoreMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.StoreMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterStoreCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcStore != nil && mm_atomic.LoadUint64(&m.afterStoreCounter) < 1 {
		return false
	}
	return true
}

// MinimockStoreInspect logs each unmet expectation
func (m *RepoMock) MinimockStoreInspect() {
	for _, e := range m.StoreMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to RepoMock.Store with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.StoreMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterStoreCounter) < 1 {
		if m.StoreMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to RepoMock.Store")
		} else {
			m.t.Errorf("Expected call to RepoMock.Store with params: %#v", *m.StoreMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcStore != nil && mm_atomic.LoadUint64(&m.afterStoreCounter) < 1 {
		m.t.Error("Expected call to RepoMock.Store")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *RepoMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockExistsInspect()

		m.MinimockStoreInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *RepoMock) MinimockWait(timeout mm_time.Duration) {
	timeoutCh := mm_time.After(timeout)
	for {
		if m.minimockDone() {
			return
		}
		select {
		case <-timeoutCh:
			m.MinimockFinish()
			return
		case <-mm_time.After(10 * mm_time.Millisecond):
		}
	}
}

func (m *RepoMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockExistsDone() &&
		m.MinimockStoreDone()
}
