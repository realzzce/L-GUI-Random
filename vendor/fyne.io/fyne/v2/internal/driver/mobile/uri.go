// +build !android

package mobile

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/storage"
)

func nativeURI(uri string) fyne.URI {
	return storage.NewURI(uri)
}
