package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	fmt.Println("Before Unit Test")

	m.Run() // eksekusi semua unit test

	fmt.Println("After Unit Test")
}

func BenchmarkHelloWorldSatu(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("satu")
	}
}

func BenchmarkHelloWorldDua(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("dua")
	}
}

func BenchmarkSub(b *testing.B) {
	b.Run("satu", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("satu")
		}
	})
	b.Run("dua", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("dua")
		}
	})
}

func BenchmarkTable(b *testing.B) {
	benchmark := []struct {
		name    string
		request string
	}{
		{
			name:    "gandhi",
			request: "gandhi",
		},
		{
			name:    "iwan",
			request: "iwan",
		},
	}
	for _, benchmark := range benchmark {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.request)
			}
		})
	}
}

func TestHelloWorldGandhi(t *testing.T) {
	result := HelloWorld("gandhi")
	if result != "Hello gandhi" {
		//failed unit test use testing.T with function t.Fail()
		t.Fail()
	}
	fmt.Println("TestHelloWorldGandhi Done")
}

func TestHelloWorldIwan(t *testing.T) {
	result := HelloWorld("iwan")
	if result != "Hello iwan" {
		//failed unit test use testing.T with function t.FailNow()
		t.FailNow()
	}
	fmt.Println("TestHelloWorldGandhi Done")

}

func TestHelloWorldXyz(t *testing.T) {
	result := HelloWorld("xyz")
	if result != "Hello xyz" {
		//failed unit test use testing.T with function t.Error(args)
		t.Error("Result must be : Hello xyz")
	}
	fmt.Println("TestHelloWorldXyz Done")
}

func TestHelloWorldAbc(t *testing.T) {
	result := HelloWorld("abc")
	if result != "Hello abc" {
		//failed unit test use testing.T with function t.Fatal(args)
		t.Fatal("Result must be : Hello abc")
	}
	fmt.Println("TestHelloWorldAbc Done")

}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Assert")
	// assert equality
	// assert.Equal(t, 123, 123, "they should be equal")
	assert.Equal(t, "Hello Assert", result, "Result must be : Hello Assert")
	fmt.Println("TestHelloWorldAssert Done")

}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Require")
	require.Equal(t, "Hello Require", result, "Result must be : Hello Require")
	fmt.Println("TestHelloWorldRequire Done")

}

func TestSubTest(t *testing.T) {
	t.Run("nameSubTest1", func(t *testing.T) {
		result := HelloWorld("nameSubTest1")
		require.Equal(t, "Hello nameSubTest1", result, "Result must be : Hello nameSubTest1")
	})
	t.Run("nameSubTest2", func(t *testing.T) {
		result := HelloWorld("nameSubTest2")
		require.Equal(t, "Hello nameSubTest2", result, "Result must be : Hello nameSubTest2")
	})
}

func TestHelloWorldTable(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "gandhi", //name subtest
			request:  "gandhi",
			expected: "Hello gandhi",
		},
		{
			name:     "iwan", //name subtest
			request:  "iwan",
			expected: "Hello iwan",
		},
	}
	fmt.Println(tests)
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}
