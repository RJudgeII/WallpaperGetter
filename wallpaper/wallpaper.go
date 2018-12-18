package wallpaper

import (
	"errors"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"syscall"
	"unsafe"
)

// Microsoft Windows system-wide parameters
const (
	// Sets the path to the desktop wallpaper file
	spiSetDeskWallpaper = 0x0014

	uiParam = 0x0000

	spiUpdateFile = 0x01
	spiSendChange = 0x02
)

var (
	user32            = syscall.NewLazyDLL("user32.dll")
	sysParametersInfo = user32.NewProc("SystemParametersInfoW")
)

// SetFromFile sets from file
func SetFromFile(filename string) error {
	filenameUTF16, err := syscall.UTF16PtrFromString(filename)
	if err != nil {
		return err
	}

	sysParametersInfo.Call(
		uintptr(spiSetDeskWallpaper),
		uintptr(uiParam),
		uintptr(unsafe.Pointer(filenameUTF16)),
		uintptr(spiUpdateFile|spiSendChange),
	)
	return nil
}

// SetFromURL sets from URL
func SetFromURL(url string) error {
	file, err := DownloadImage(url)
	if err != nil {
		return err
	}

	return SetFromFile(file)
}

// GetCacheDir gets cache dir
func GetCacheDir() (string, error) {
	return os.TempDir(), nil
}

// Desktop gets desktop
var Desktop = os.Getenv("XDG_CURRENT_DESKTOP")

// DesktopSession gets desktop environment
var DesktopSession = os.Getenv("DESKTOP_SESSION")

// ErrUnsupportedDE is thrown when Desktop is not a supported desktop environment.
var ErrUnsupportedDE = errors.New("your desktop environment is not supported")

// DownloadImage downloads image
func DownloadImage(url string) (string, error) {
	cacheDir, err := GetCacheDir()
	if err != nil {
		return "", err
	}

	filename := filepath.Join(cacheDir, filepath.Base(url))
	file, err := os.Create(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	res, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return "", errors.New("non-200 status code")
	}

	_, err = io.Copy(file, res.Body)
	if err != nil {
		return "", err
	}

	err = file.Close()
	if err != nil {
		return "", err
	}

	return filename, nil
}
