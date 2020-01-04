package importlist

import (
	"io/ioutil"
	"reflect"
	"testing"
)

func TestImportlist_importList(t *testing.T) {
	t.Run("single line", func(t *testing.T) {
		expected := []string{"github.com/sanemat/go-xgoinstall/cmd/x-go-install"}
		data, err := ioutil.ReadFile("test/tools1.go.txt")
		if err != nil {
			t.Errorf("error should be nil, but: %s", err)
		}
		list, err := importList(data)
		if err != nil {
			t.Errorf("error should be nil, but: %s", err)
		}
		if !reflect.DeepEqual(list, expected) {
			t.Errorf("\n   list: %+v\nexpected: %+v", list, expected)
		}
	})

	t.Run("multi lines", func(t *testing.T) {
		expected := []string{"github.com/sanemat/go-xgoinstall/cmd/x-go-install", "golang.org/x/lint/golint"}
		data, err := ioutil.ReadFile("test/tools2.go.txt")
		if err != nil {
			t.Errorf("error should be nil, but: %s", err)
		}
		list, err := importList(data)
		if err != nil {
			t.Errorf("error should be nil, but: %s", err)
		}
		if !reflect.DeepEqual(list, expected) {
			t.Errorf("\n   list: %+v\nexpected: %+v", list, expected)
		}
	})
}
