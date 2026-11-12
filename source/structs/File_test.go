package structs

import "bytes"
import "crypto/sha256"
import "encoding/hex"
import "os"
import "os/user"
import "path"
import "testing"

func fileTestDigest(value []byte) string {

	checksum := sha256.Sum256(value)

	return hex.EncodeToString(checksum[:])

}

func TestNewFile(t *testing.T) {

	folder := t.TempDir()
	file := NewFile("  hello.txt  ", folder)

	if file.Name != "hello.txt" {
		t.Errorf("NewFile() Name = %q, expected %q", file.Name, "hello.txt")
	}

	if file.Folder != folder {
		t.Errorf("NewFile() Folder = %q, expected %q", file.Folder, folder)
	}

	if len(file.Buffer) != 0 {
		t.Errorf("NewFile() Buffer = %v, expected an empty buffer", file.Buffer)
	}

	invalid := NewFile("../escape", folder)

	if invalid.Name != "" {
		t.Errorf("NewFile() accepted an unsafe name: %q", invalid.Name)
	}

}

func TestFileSetters(t *testing.T) {

	folder := t.TempDir()
	file := NewFile("hello.txt", folder)

	file.SetName("  world.txt  ")

	if file.Name != "world.txt" {
		t.Errorf("SetName() = %q, expected %q", file.Name, "world.txt")
	}

	file.SetName("../escape")

	if file.Name != "world.txt" {
		t.Errorf("SetName() accepted an unsafe name: %q", file.Name)
	}

	file.SetFolder("relative")

	if file.Folder != folder {
		t.Errorf("SetFolder() accepted a relative folder: %q", file.Folder)
	}

	file.SetExecutable(true)

	if file.IsExecutable() == false {
		t.Fatal("SetExecutable() did not update the executable flag")
	}

}

func TestFileSetBuffer(t *testing.T) {

	file := NewFile("hello.txt", t.TempDir())
	buffer := []byte("hello")

	file.SetBuffer(buffer)

	if bytes.Equal(file.Buffer, buffer) == false {
		t.Errorf("SetBuffer() Buffer = %v, expected %v", file.Buffer, buffer)
	}

	if file.Digest != fileTestDigest(buffer) {
		t.Errorf("SetBuffer() Digest = %q, expected %q", file.Digest, fileTestDigest(buffer))
	}

	buffer[0] = 'x'

	if file.Buffer[0] != 'h' {
		t.Fatal("SetBuffer() did not copy the provided buffer")
	}

	file.SetBuffer([]byte{})
	empty_digest := fileTestDigest([]byte{})

	if file.Buffer == nil {
		t.Fatal("SetBuffer() should keep an empty non-nil buffer")
	}

	if file.Digest != empty_digest {
		t.Errorf("SetBuffer() Digest = %q, expected %q", file.Digest, empty_digest)
	}

}

func TestFileWriteRead(t *testing.T) {

	folder := t.TempDir()
	content := []byte("hello file")
	file := NewFile("hello.txt", folder)

	file.SetBuffer(content)

	if err := file.Write(); err != nil {
		t.Fatalf("Write() returned an error: %s", err)
	}

	stat, err := os.Stat(path.Join(folder, "hello.txt"))

	if err != nil {
		t.Fatalf("os.Stat() returned an error: %s", err)
	}

	if stat.Mode().Perm() != 0644 {
		t.Errorf("Write() mode = %o, expected %o", stat.Mode().Perm(), 0644)
	}

	read := NewFile("hello.txt", folder)

	if err := read.Read(); err != nil {
		t.Fatalf("Read() returned an error: %s", err)
	}

	if bytes.Equal(read.Buffer, content) == false {
		t.Errorf("Read() Buffer = %v, expected %v", read.Buffer, content)
	}

	if read.Digest != file.Digest {
		t.Errorf("Read() Digest = %q, expected %q", read.Digest, file.Digest)
	}

	if read.IsExecutable() == true {
		t.Fatal("Read() marked a non-executable file as executable")
	}

	account, err := user.Current()

	if err == nil {

		if read.User.Name != account.Username {
			t.Errorf("Read() User.Name = %q, expected %q", read.User.Name, account.Username)
		}

		if read.User.Type != "user" {
			t.Errorf("Read() User.Type = %q, expected %q", read.User.Type, "user")
		}

	}

}

func TestFileWriteReadExecutable(t *testing.T) {

	folder := t.TempDir()
	file := NewFile("run.sh", folder)

	file.SetBuffer([]byte("#!/bin/sh\n"))
	file.SetExecutable(true)

	if err := file.Write(); err != nil {
		t.Fatalf("Write() returned an error: %s", err)
	}

	stat, err := os.Stat(path.Join(folder, "run.sh"))

	if err != nil {
		t.Fatalf("os.Stat() returned an error: %s", err)
	}

	if stat.Mode().Perm() != 0755 {
		t.Errorf("Write() mode = %o, expected %o", stat.Mode().Perm(), 0755)
	}

	read := NewFile("run.sh", folder)

	if err := read.Read(); err != nil {
		t.Fatalf("Read() returned an error: %s", err)
	}

	if read.IsExecutable() == false {
		t.Fatal("Read() did not mark an executable file as executable")
	}

	read.SetExecutable(false)

	if err := read.Write(); err != nil {
		t.Fatalf("Write() returned an error: %s", err)
	}

	stat, err = os.Stat(path.Join(folder, "run.sh"))

	if err != nil {
		t.Fatalf("os.Stat() returned an error: %s", err)
	}

	if stat.Mode().Perm() != 0644 {
		t.Errorf("Write() mode = %o, expected %o after clearing the executable flag", stat.Mode().Perm(), 0644)
	}

}

func TestFileWriteRecalculatesDigest(t *testing.T) {

	file := NewFile("hello.txt", t.TempDir())

	file.Buffer = []byte("direct")

	if err := file.Write(); err != nil {
		t.Fatalf("Write() returned an error: %s", err)
	}

	if file.Digest != fileTestDigest([]byte("direct")) {
		t.Errorf("Write() Digest = %q, expected %q", file.Digest, fileTestDigest([]byte("direct")))
	}

}

func TestFileReadErrors(t *testing.T) {

	var invalid File

	if invalid.Read() == nil {
		t.Fatal("Read() should fail for an invalid file")
	}

	missing := NewFile("missing.txt", t.TempDir())

	if missing.Read() == nil {
		t.Fatal("Read() should fail for a missing file")
	}

	folder := t.TempDir()
	directory := NewFile("subfolder", folder)

	if err := os.Mkdir(path.Join(folder, "subfolder"), 0755); err != nil {
		t.Fatalf("os.Mkdir() returned an error: %s", err)
	}

	if directory.Read() == nil {
		t.Fatal("Read() should fail for a directory")
	}

}

func TestFileWriteErrors(t *testing.T) {

	var invalid File

	if invalid.Write() == nil {
		t.Fatal("Write() should fail for an invalid file")
	}

	folder := path.Join(t.TempDir(), "missing")
	file := NewFile("hello.txt", folder)

	if err := file.Write(); err == nil {
		t.Fatal("Write() should fail when the folder does not exist")
	}

}

func TestFileHash(t *testing.T) {

	file := NewFile("hello.txt", t.TempDir())

	file.SetBuffer([]byte("hello"))
	hash := file.Hash()

	if hash == "" {
		t.Fatal("Hash() returned an empty string for a populated file")
	}

	if file.Hash() != hash {
		t.Fatal("Hash() is not deterministic")
	}

	file.SetBuffer([]byte("world"))

	if file.Hash() == hash {
		t.Fatal("Hash() did not change when the buffer changed")
	}

	empty := File{}

	if empty.Hash() != "" {
		t.Fatal("Hash() should be empty without a name")
	}

}
