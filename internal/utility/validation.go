package utility

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func Validate[T any](data T) map[string]string {
	err := validator.New().Struct(data)
	res := map[string]string{}

	if err != nil {
		for _, v := range err.(validator.ValidationErrors) {
			res[v.StructField()] = TransLateTag(v)
		}
	}

	return res

}

func TransLateTag(fd validator.FieldError) string {
	switch fd.ActualTag() {
		case "required":
			return  fmt.Sprintf("field %s wajib di isi", fd.StructField())
	}

	return "Validasi Gagal"
}