package unit

import (
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
	"github.com/tanapon395/sa-66-example/entity"
)

func TestMemberPhoneNumber(t *testing.T) {

	g := NewGomegaWithT(t)

	t.Run(`PHONE CHECKED`, func(t *testing.T) {
		member := entity.Member{
			PhoneNumber: "0123456789",
			Password:    "1234",
			Url:         "https://www.google.com/",
		}

		ok, err := govalidator.ValidateStruct(member)

		g.Expect(ok).To(BeTrue())
		g.Expect(err).To(BeNil())
	})

	t.Run(`phone_number is required`, func(t *testing.T) {
		member := entity.Member{
			PhoneNumber: "", //ไม่ใส่ค่า
			Password:    "1234",
			Url:         "https://www.google.com/",
		}

		ok, err := govalidator.ValidateStruct(member)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("Phone is required"))
	})

	t.Run(`phone_number check 10 digit`, func(t *testing.T) {
		member := entity.Member{
			PhoneNumber: "08000000000", //11ตัว
			Password:    "1234",
			Url:         "https://www.linkedin.com/company/ilink/",
		}

		ok, err := govalidator.ValidateStruct(member)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal(("PhoneNumber length is not 10 digits.")))
	})
}

func TestUrlMember(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run(`URL CHECKED`, func(t *testing.T) {
		member := entity.Member{
			PhoneNumber: "0800000000",
			Password:    "",
			Url:         "https://www.linkedin.com/company/ilink/",
		}

		ok, err := govalidator.ValidateStruct(member)

		g.Expect(ok).To(BeTrue())
		g.Expect(err).To(BeNil())
	})

	t.Run(`Url is required`, func(t *testing.T) {
		member := entity.Member{
			PhoneNumber: "0800000000",
			Password:    "",
			Url:         "", //ไม่ใส่ URL
		}

		ok, err := govalidator.ValidateStruct(member)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("Url is required"))
	})

	t.Run(`Url is required`, func(t *testing.T) {
		member := entity.Member{
			PhoneNumber: "0800000000",
			Password:    "",
			Url:         "wwwlinkedincomcompanyilink", //URL ผิดรูปแบบ
		}

		ok, err := govalidator.ValidateStruct(member)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("Url is invalid"))
	})
}
