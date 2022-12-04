package faker

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/dop251/goja"
	"lukechampine.com/frand"
)

type Faker struct {
	*gofakeit.Faker
	rt *goja.Runtime
}

func newFaker(seed int64, rt *goja.Runtime) *Faker {
	src := frand.NewSource()

	if seed != 0 {
		src.Seed(seed)
	}

	return &Faker{Faker: gofakeit.NewCustom(src), rt: rt}
}

func (f *Faker) Ipv4Address() string {
	return f.IPv4Address()
}

func (f *Faker) Ipv6Address() string {
	return f.IPv6Address()
}

func (f *Faker) HttpStatusCodeSimple() int {
	return f.HTTPStatusCodeSimple()
}

func (f *Faker) HttpStatusCode() int {
	return f.HTTPStatusCode()
}

func (f *Faker) HttpMethod() string {
	return f.HTTPMethod()
}

func (f *Faker) Bs() string {
	return f.BS()
}

func (f *Faker) Uuid() string {
	return f.UUID()
}

func (f *Faker) RgbColor() []int {
	return f.RGBColor()
}

func (f *Faker) ImageJpeg(width int, height int) goja.ArrayBuffer {
	return f.rt.NewArrayBuffer(f.Faker.ImageJpeg(width, height))
}

func (f *Faker) ImagePng(width int, height int) goja.ArrayBuffer {
	return f.rt.NewArrayBuffer(f.Faker.ImagePng(width, height))
}
