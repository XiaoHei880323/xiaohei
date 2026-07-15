package main

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
	"time"
	"unsafe"
)

const (
	backendHealthURL = "http://127.0.0.1:8888/from/xiaohe"
	adminPageURL     = "http://127.0.0.1:8888/admin/page/index"
)

func main() {
	exePath, err := os.Executable()
	if err != nil {
		showMessage("Client Error", err.Error())
		os.Exit(1)
	}

	clientDir := filepath.Dir(exePath)
	projectDir := filepath.Dir(clientDir)
	apiDir := filepath.Join(projectDir, "api")
	apiExePath := filepath.Join(apiDir, "api.exe")
	apiConfigPath := filepath.Join(apiDir, "etc", "api-api.yaml")
	profileDir := filepath.Join(clientDir, "client-profile")

	if err := os.MkdirAll(profileDir, 0o755); err != nil {
		showMessage("Client Error", err.Error())
		os.Exit(1)
	}

	if !isBackendReady() {
		if _, err := os.Stat(apiExePath); err != nil {
			showMessage("Backend Missing", fmt.Sprintf("File not found: %s", apiExePath))
			os.Exit(1)
		}

		if _, err := os.Stat(apiConfigPath); err != nil {
			showMessage("Backend Missing", fmt.Sprintf("File not found: %s", apiConfigPath))
			os.Exit(1)
		}

		if err := startBackend(apiDir, apiExePath); err != nil {
			showMessage("Backend Start Failed", err.Error())
			os.Exit(1)
		}

		if err := waitBackendReady(45 * time.Second); err != nil {
			showMessage("Backend Not Ready", err.Error())
			os.Exit(1)
		}
	}

	browserPath := findBrowser()
	if browserPath == "" {
		showMessage("Browser Missing", "Microsoft Edge or Google Chrome was not found.")
		os.Exit(1)
	}

	cmd := exec.Command(
		browserPath,
		"--app="+adminPageURL,
		"--new-window",
		"--no-first-run",
		"--user-data-dir="+profileDir,
		"--window-size=1440,920",
	)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}

	if err := cmd.Start(); err != nil {
		showMessage("Launch Failed", err.Error())
		os.Exit(1)
	}
}

func startBackend(apiDir, apiExePath string) error {
	cmd := exec.Command(apiExePath, "-f", "etc\\api-api.yaml")
	cmd.Dir = apiDir
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	return cmd.Start()
}

func waitBackendReady(timeout time.Duration) error {
	deadline := time.Now().Add(timeout)
	for time.Now().Before(deadline) {
		if isBackendReady() {
			return nil
		}
		time.Sleep(1200 * time.Millisecond)
	}

	return fmt.Errorf("backend did not become ready in %s; check MySQL and api config", timeout.String())
}

func isBackendReady() bool {
	client := &http.Client{Timeout: 1200 * time.Millisecond}

	for _, url := range []string{adminPageURL, backendHealthURL} {
		resp, err := client.Get(url)
		if err != nil {
			continue
		}
		resp.Body.Close()
		if resp.StatusCode >= 200 && resp.StatusCode < 300 {
			return true
		}
	}

	return false
}

func findBrowser() string {
	candidates := []string{
		"C:\\Program Files (x86)\\Microsoft\\Edge\\Application\\msedge.exe",
		"C:\\Program Files\\Microsoft\\Edge\\Application\\msedge.exe",
		"C:\\Users\\Administrator\\AppData\\Local\\Microsoft\\Edge\\Application\\msedge.exe",
		"C:\\Program Files\\Google\\Chrome\\Application\\chrome.exe",
		"C:\\Program Files (x86)\\Google\\Chrome\\Application\\chrome.exe",
	}

	for _, path := range candidates {
		if _, err := os.Stat(path); err == nil {
			return path
		}
	}

	if path, err := exec.LookPath("msedge.exe"); err == nil {
		return path
	}

	if path, err := exec.LookPath("chrome.exe"); err == nil {
		return path
	}

	return ""
}

func showMessage(title, message string) {
	user32 := syscall.NewLazyDLL("user32.dll")
	messageBoxW := user32.NewProc("MessageBoxW")

	titlePtr, _ := syscall.UTF16PtrFromString(title)
	messagePtr, _ := syscall.UTF16PtrFromString(message)
	messageBoxW.Call(
		0,
		uintptr(unsafe.Pointer(messagePtr)),
		uintptr(unsafe.Pointer(titlePtr)),
		0,
	)
}
