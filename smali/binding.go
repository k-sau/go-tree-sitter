package smali

// #cgo CFLAGS: -std=c11 -fPIC
// #include "parser.c"
// #if __has_include("scanner.c")
// #include "scanner.c"
// #endif
// TSLanguage *tree_sitter_smali();
import "C"

import (
	"unsafe"

	sitter "github.com/smacker/go-tree-sitter"
)

func GetLanguage() *sitter.Language {
	ptr := unsafe.Pointer(C.tree_sitter_smali())
	return sitter.NewLanguage(ptr)
}
