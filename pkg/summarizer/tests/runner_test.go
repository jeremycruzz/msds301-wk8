package summarizer_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestSummarizer(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Summarizer Suite")
}
