package jsdec_tiny

import (
	_ "embed"

	"github.com/robertkrimen/otto"
)

type Jsdec struct {
	js_text string
	js_vm   *otto.Otto
}

//go:embed dec.js
var fileByte []byte

func (m *Jsdec) openFile() error {
	m.js_text = string(fileByte)
	return nil
}

func (m *Jsdec) ModInit() error {
	return m.openFile()
}

func (m *Jsdec) LoadJS() error {
	m.js_vm = otto.New()
	_, err := m.js_vm.Run(m.js_text)
	return err
}

func (m *Jsdec) Decrypt(jsf string) (s string, e error) {
	ret, e := m.js_vm.Call("autoscan", nil, jsf)
	s = ret.String()
	return
}
