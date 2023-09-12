package fdkaac

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestFdkAac(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "FdkAac Suite")
}
