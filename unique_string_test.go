package unique_string

import (
	"regexp"
	"strconv"
	"testing"
)

func TestGenerateUniqueString(t *testing.T) {
	prev := GenerateUniqueString("test")
	if prev != "rbgf3xv4ufgzg" {
		t.Fatalf("GenerateUniqueString(\"test\") should = rbgf3xv4ufgzg")
	}

	for i := 0; i < 100; i++ {
		id := GenerateUniqueString("test", strconv.Itoa(i))
		if prev == id {
			t.Fatalf("Should get a new unique string!")
		}
		prev = id

		matched, err := regexp.MatchString(
			"^([a-zA-Z0-9_-]){13}", id)
		if !matched || err != nil {
			t.Fatalf("expected match %s %v %s", id, matched, err)
		}
	}
}

func TestGenerateUniqueString_shortString(t *testing.T) {
	prev := GenerateUniqueString("t")
	if prev != "fsnwonksnm7cy" {
		t.Fatalf("GenerateUniqueString(\"t\") should = fsnwonksnm7cy")
	}
	prev = GenerateUniqueString("te")
	if prev != "o57ehhbytctgs" {
		t.Fatalf("GenerateUniqueString(\"te\") should = o57ehhbytctgs")
	}
	prev = GenerateUniqueString("tes")
	if prev != "ebse57jls4qbc" {
		t.Fatalf("GenerateUniqueString(\"tes\") should = ebse57jls4qbc")
	}
	prev = GenerateUniqueString("t", "e")
	if prev != "d5ykh7yoidm2a" {
		t.Fatalf("GenerateUniqueString(\"t\", \"e\") should = d5ykh7yoidm2a")
	}
}

func BenchmarkGenerateUniqueString(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = GenerateUniqueString("test")
	}
}
