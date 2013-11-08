package govalidator_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/siesta/govalidator"
	"testing"
	"time"
)

func TestGovalidator(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Govalidator Suite")
}

var _ = Describe("Go Validation Library", func() {
	var validation *Validation

	BeforeEach(func() {
		validation = &Validation{}
	})

	Describe("Validating with Min", func() {
		var min Min
		BeforeEach(func() {
			min = Min{4}
		})

		Context("testing with less than allowed", func() {
			It("should return false with minus 0 values", func() {
				Expect(validation.Validate(min, -1)).To(BeFalse())
			})
			It("should return false with 0", func() {
				Expect(validation.Validate(min, 0)).To(BeFalse())
			})

			It("should skip this test", func() {
				//Expect(2).To(Equal(3))
			})
		})

		Context("testing with same value", func() {
			It("should return true", func() {
				Expect(validation.Validate(min, 4)).To(BeTrue())
			})
		})

		Context("testing with greater value", func() {
			It("should return true", func() {
				Expect(validation.Validate(min, 5)).To(BeTrue())
			})
			It("should return true", func() {
				Expect(validation.Validate(min, 11)).To(BeTrue())
			})
		})
		Context("testing with not integer values", func() {
			It("bool should return false", func() {
				Expect(validation.Validate(min, true)).To(BeFalse())
			})
			It("int64 should return false", func() {
				Expect(validation.Validate(min, 11.1212)).To(BeFalse())
			})
		})
	})
	Describe("Validating with Max", func() {
		var max Max
		BeforeEach(func() {
			max = Max{4}
		})
		Context("testing with less than allowed", func() {
			It("should return true with minus 0 values", func() {
				Expect(validation.Validate(max, -1)).To(BeTrue())
			})
			It("should return true with 0", func() {
				Expect(validation.Validate(max, 0)).To(BeTrue())
			})

		})

		Context("testing with same value", func() {
			It("should return true", func() {
				Expect(validation.Validate(max, 4)).To(BeTrue())
			})
		})

		Context("testing with greater value", func() {
			It("should return true", func() {
				Expect(validation.Validate(max, 5)).To(BeFalse())
			})
			It("should return true", func() {
				Expect(validation.Validate(max, 11)).To(BeFalse())
			})
		})
		Context("testing with not integer values", func() {
			It("bool should return false", func() {
				Expect(validation.Validate(max, true)).To(BeFalse())
			})
			It("int64 should return false", func() {
				Expect(validation.Validate(max, 11.1212)).To(BeFalse())
			})
		})
	})
	Describe("Validating with Len", func() {
		var length Len
		BeforeEach(func() {
			length = Len{4}
		})

		Context("testing with longer text", func() {
			It("should return false", func() {
				Expect(validation.Validate(length, "thisislongerthanforcharacter")).To(BeFalse())
			})
		})
		Context("testing with equal length text", func() {
			It("should return true", func() {
				Expect(validation.Validate(length, "four")).To(BeTrue())
			})
		})
		Context("testing with shorter text", func() {
			It("should return true", func() {
				Expect(validation.Validate(length, "two")).To(BeFalse())
			})
		})
		Context("testing with non string", func() {
			It("should return false", func() {
				Expect(validation.Validate(length, 2)).To(BeFalse())
			})
		})
	})
	Describe("Validating with Required", func() {
		Context("testing with string values", func() {
			It("should return false when empty string", func() {
				Expect(validation.Validate(Required{}, "")).To(BeFalse())
			})
			It("should return true any string", func() {
				Expect(validation.Validate(Required{}, "foo")).To(BeTrue())
			})
		})
		Context("testing with interface value", func() {
			It("should return false when empty string", func() {
				var intf interface{}

				Expect(validation.Validate(Required{}, intf)).To(BeFalse())
			})
		})
		Context("testing with int64 value", func() {
			It("should return false when empty string", func() {
				Expect(validation.Validate(Required{}, 64.9)).To(BeFalse())
			})
		})
		Context("testing with nil value", func() {
			It("should return false when data is nil", func() {
				Expect(validation.Validate(Required{}, nil)).To(BeFalse())
			})
		})
		Context("testing with bool value", func() {
			It("should return true with true", func() {
				Expect(validation.Validate(Required{}, true)).To(BeTrue())
			})
			It("should return false with false", func() {
				Expect(validation.Validate(Required{}, false)).To(BeFalse())
			})
		})
		Context("testing with integer value", func() {
			It("should return true with 1", func() {
				Expect(validation.Validate(Required{}, 1)).To(BeTrue())
			})
			It("should return true with -1", func() {
				Expect(validation.Validate(Required{}, -1)).To(BeTrue())
			})
			It("should return false with 0", func() {
				Expect(validation.Validate(Required{}, 0)).To(BeFalse())
			})
		})
		Context("testing with date time", func() {
			It("should return true with now", func() {
				Expect(validation.Validate(Required{}, time.Now())).To(BeTrue())
			})
			It("should return true with pointer to now", func() {
				now := time.Now()
				Expect(validation.Validate(Required{}, &now)).To(BeTrue())
			})
			It("should return true with yesterday", func() {
				//              today := time.Now()
				//              yesterday := time.Date(
				//                  today.Year(),
				//                  today.Month(),
				//                  today.Date()-1,
				//                  today.Hour(),
				//                  today.Minute(),
				//                  today.Second(),
				//                  today.Nanosecond(),
				//                  today.Location(),
				//              )
				//  Expect(validation.Validate(Required{}, yesterday).To(BeTrue())
			})
			It("should return false with init", func() {
				Expect(validation.Validate(Required{}, &time.Time{})).To(BeFalse())
			})

		})
		// todo write more behaviours for regex
		Context("testing with match", func() {
			It("should return true with ", func() {
				Expect(validation.Validate(NewMatch(`[a-z]`), "foobar")).To(BeTrue())
			})
			It("should return false with dash", func() {
				Expect(validation.Validate(NewMatch(`az`), "a")).To(BeFalse())
			})
			It("should return false non string regex", func() {
				Expect(validation.Validate(NewMatch(`az`), 3)).To(BeFalse())
			})
		})
		Context("testing with email", func() {
			It("should return true with cihangir@koding.com", func() {
				Expect(validation.Validate(NewEmail(), "cihangir@koding.com")).To(BeTrue())
			})
			It("should return true with cihangir+tag@koding.com", func() {
				Expect(validation.Validate(NewEmail(), "cihangir+tag@koding.com")).To(BeTrue())
			})
			It("should return false with dash", func() {
				Expect(validation.Validate(NewEmail(), "a")).To(BeFalse())
			})
			It("should return false with aa.com", func() {
				Expect(validation.Validate(NewEmail(), "aa.com")).To(BeFalse())
			})
		})
	})
})
