package render

import (
	"fmt"
	"os"

	"github.com/go-gl/gl/v4.5-core/gl"
)

type Shader struct {
	Id uint32
}

func NewShader(vertex_path, fragment_path string) *Shader {
	vertex_bytes, err := os.ReadFile(vertex_path)
	if err != nil {
		panic(fmt.Errorf("Failed to read vertex '%s': %s", vertex_path, err))
	}
	vertex_code, vfree := gl.Strs(string(vertex_bytes) + "\x00")
	defer vfree()

	fragment_bytes, err := os.ReadFile(fragment_path)
	if err != nil {
		panic(fmt.Errorf("Failed to read fragment '%s': %s", fragment_path, err))
	}
	fragment_code, ffree := gl.Strs(string(fragment_bytes) + "\x00")
	defer ffree()

	var status int32

	vertex := gl.CreateShader(gl.VERTEX_SHADER)
	defer gl.DeleteShader(vertex)
	gl.ShaderSource(vertex, 1, vertex_code, nil)
	gl.CompileShader(vertex)
	gl.GetShaderiv(vertex, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		log := make([]byte, 1024)
		gl.GetShaderInfoLog(vertex, 1024, nil, &log[0])
		panic(fmt.Errorf("Failed to compile vertex shader '%s': %s", vertex_path, log))
	}

	fragment := gl.CreateShader(gl.FRAGMENT_SHADER)
	defer gl.DeleteShader(fragment)
	gl.ShaderSource(fragment, 1, fragment_code, nil)
	gl.CompileShader(fragment)
	gl.GetShaderiv(fragment, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		log := make([]byte, 1024)
		gl.GetShaderInfoLog(fragment, 1024, nil, &log[0])
		panic(fmt.Errorf("Failed to compile fragment shader '%s': %s", fragment_path, log))
	}

	s := &Shader{}

	s.Id = gl.CreateProgram()
	gl.AttachShader(s.Id, vertex)
	gl.AttachShader(s.Id, fragment)
	gl.LinkProgram(s.Id)
	gl.GetProgramiv(s.Id, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		log := make([]byte, 1024)
		gl.GetProgramInfoLog(s.Id, 1024, nil, &log[0])
		panic(fmt.Errorf("Failed to link shader program '%s': %s", vertex_path, log))
	}
	return s
}

func (s *Shader) Use() {
	gl.UseProgram(s.Id)
}

func (s *Shader) SetBool(name string, value bool) {
	if value {
		s.SetInt(name, 1)
	} else {
		s.SetInt(name, 0)
	}
}

func (s *Shader) SetInt(name string, value int32) {
	gl.Uniform1i(gl.GetUniformLocation(s.Id, gl.Str(name+"\x00")), value)
}

func (s *Shader) SetFloat(name string, value float32) {
	gl.Uniform1f(gl.GetUniformLocation(s.Id, gl.Str(name+"\x00")), value)
}

func (s *Shader) Set4Float(name string, v0, v1, v2, v3 float32) {
	gl.Uniform4f(gl.GetUniformLocation(s.Id, gl.Str(name+"\x00")), v0, v1, v2, v3)
}
