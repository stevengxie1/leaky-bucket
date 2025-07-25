package main

import (
	"testing"
	"time"
)

// TestCounterAdd: Tests the Add function with some basic expectations
// Understanding is that for a given time, this does some math like so(?)
// We start with some constants for 120 as the kind of 'drip' rate
// that's going to in this given second, be 2 per second
// so we expect 1 remaining if we use 1
func TestCounterAdd(t *testing.T) {
	var counter = NewCounter()

	mockTime := time.Unix(0, 0)

	expectedBucketKey := "test1"
	expectedBucketSize := int64(2) // see above math
	expectedRemaining := int64(1)  // see above
	expectedAllowed := true

	gotInfo := counter.Add(expectedBucketKey, 120, 1, mockTime)

	if gotInfo.Bucket != expectedBucketKey {
		t.Errorf("Add() Bucket got %v wanted %v", gotInfo.Bucket, expectedBucketKey)
	}
	if gotInfo.BucketSize != expectedBucketSize {
		t.Errorf("Add() BucketSize got %v wanted %v", gotInfo.BucketSize, expectedBucketSize)
	}
	if gotInfo.Remaining != expectedRemaining {
		t.Errorf("Add() Remaining got %v wanted %v", gotInfo.Remaining, expectedRemaining)
	}
	if gotInfo.Allowed != expectedAllowed {
		t.Errorf("Add() expectedAllowed got %v wanted %v", gotInfo.Allowed, expectedAllowed)
	}

}

// TestCounterAdd_HitRateLimit: Tests the Add function with some 'simulated' requests in
// a controlled time window, at unix 0
// As a follow up, try to exhaust the time window
func TestCounterAdd_HitRateLimit(t *testing.T) {
	var counter = NewCounter()

	mockTime := time.Unix(0, 0)
	expectedBucketKey := "test1"
	expectedFirstRequestsAllowed := true

	gotInfo := counter.Add(expectedBucketKey, 120, 1, mockTime)
	if gotInfo.Allowed != expectedFirstRequestsAllowed {
		t.Errorf("Add() HitRateLimit got %v wanted %v", gotInfo.Allowed, expectedFirstRequestsAllowed)
	}

	gotInfo = counter.Add(expectedBucketKey, 120, 1, mockTime)
	if gotInfo.Allowed != expectedFirstRequestsAllowed {
		t.Errorf("Add() HitRateLimit got %v wanted %v", gotInfo.Allowed, expectedFirstRequestsAllowed)
	}

	// Expect a rejection
	gotInfo = counter.Add(expectedBucketKey, 120, 1, mockTime)
	if gotInfo.Allowed == true {
		t.Errorf("Add() HitRateLimit got %v wanted %v, expected to be rate limited", gotInfo.Allowed, expectedFirstRequestsAllowed)
	}
}

// TODO: Do more, edge cases, test suite
