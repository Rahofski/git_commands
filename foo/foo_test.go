package foo

import "testing"

func FooTestFunc(t *testing.T) {
    expectedFooResult := "woof"
    if actualFooResult := Foo(); actualFooResult != expectedFooResult {
        t.Errorf("expected %s; got: %s", expectedFooResult, actualFooResult)
    }
}