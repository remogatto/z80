// z80.go: generate Go code for Z80 opcodes
//
// Based on z80.pl by Philip Kendall
// Based on modified z80.pl by Andrea Fazzi
//
// Copyright (c) 1999-2006 Philip Kendall <philip-fuse@shadowmagic.org.uk>
// Copyright (c) 2010 Andrea Fazzi <andrea.fazzi@alcacoop.it>
// Copyright (c) 2011 ⚛ <0xe2.0x9a.0x9b@gmail.com>
//
// This program is free software; you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation; either version 2 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License along
// with this program; if not, write to the Free Software Foundation, Inc.,
// 51 Franklin Street, Fifth Floor, Boston, MA 02110-1301 USA.

// +build ignore

package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
	"text/template"
)

// Returns whether 's' matches 'pattern'
func matches(s, pattern string) bool {
	return regexp.MustCompile(pattern).MatchString(s)
}

// Return the lowercased version of 's'
func lc(s string) string {
	return strings.ToLower(s)
}

// The writer to which we are currently printing
var outputStream io.Writer = nil

// Joins the strings in the 'stringList', prints them to 'outputStream'
// and sends a new line to 'outputStream'
func ln(stringList ...string) {
	for _, s := range stringList {
		fmt.Fprint(outputStream, s)
	}
	fmt.Fprintln(outputStream)
}

// Returns 'if_true' or 'if_false' depending on the value of 'cond'.
// The purpose of this function is to reduce the number of code lines.
func _if(cond bool, if_true, if_false string) string {
	if cond {
		return if_true
	}
	return if_false
}

func readWord(varString string) {
	ln("address++")
	ln("b1 := memory.ReadByte(address)")
	ln("address++")
	ln("b2 := memory.ReadByte(address)")
	ln(varString + " := joinBytes(b2, b1)")
}

func readByte(varString string) {
	ln("address++")
	ln(varString + " := memory.ReadByte(address)")
}

func result(args ...string) {
	out := ""
	for i, s := range args {
		switch s {
		case "nnnn":
			readWord("nnnn")
			out += "fmt.Sprintf(\"0x%04x\", nnnn)"
			break
		case "(nnnn)":
			readWord("_nnnn")
			out += "fmt.Sprintf(\"(0x%04x)\", _nnnn)"
			break
		case "nn":
			readByte("nn")
			out += "fmt.Sprintf(\"(0x%02x)\", nn)"
		case "(REGISTER+dd)":
			readByte("register_dd")
			out += "fmt.Sprintf(\"(REGISTER+0x%02x)\", register_dd)"
			break
		case "offset":
			readByte("offset")
			out += "fmt.Sprintf(\"0x%x\", offset)"
			break
		default:
			out += fmt.Sprintf("\"%s\"", s)
			break
		}
		if i < len(args) - 1 {
			out += "+" + "\",\"" + "+"
		}
	}
	if len(args) > 0 {
		ln(fmt.Sprintf("result += %s", out))
	}
}

// Description of each file
var description = map[string]string{
	"opcodes_cb.dat":     "z80_cb.c: Z80 CBxx opcodes",
	"opcodes_ddfd.dat":   "z80_ddfd.c Z80 {DD,FD}xx opcodes",
	"opcodes_ddfdcb.dat": "z80_ddfdcb.c Z80 {DD,FD}CBxx opcodes",
	"opcodes_ed.dat":     "z80_ed.c: Z80 CBxx opcodes",
	"opcodes_base.dat":   "opcodes_base.c: unshifted Z80 opcodes",
}

// Removes characters which cannot form a Go identifier
func turnIntoIdentifier(in string) string {
	var out bytes.Buffer
	for _, rune := range strings.ToUpper(in) {
		switch rune {
		case ' ', ',':
			out.WriteByte('_')
		case '(':
			// Indirection
			out.WriteByte('i')
		case '+':
			// Plus
			out.WriteByte('p')
		case ')', '\'':
			// Delete
		default:
			out.WriteRune(rune)
		}
	}
	return out.String()
}

func processDataFile(data_file, logical_data_file string, code *bytes.Buffer, functions *bytes.Buffer) {
	outputStream = code

	var data []byte
	var err error
	data, err = ioutil.ReadFile(data_file)
	if err != nil {
		panic(err.Error())
	}

	lines := strings.Split(string(data), "\n")

	var fallthrough_cases []string

	for _, line := range lines {
		// Remove comments
		if strings.Contains(line, "#") {
			line = line[0:strings.Index(line, "#")]
		}

		line = strings.TrimSpace(line)

		// Skip blank lines
		if len(line) == 0 {
			continue
		}

		var l []string = strings.Split(line, " ")

		var number, opcode, arguments, extra string
		number = l[0]
		if len(l) >= 2 {
			opcode = l[1]
		}
		if len(l) >= 3 {
			arguments = l[2]
		}
		if len(l) >= 4 {
			extra = l[3]
		}

		var args []string
		if arguments != "" {
			args = strings.Split(arguments, ",")
		}

		var shift_op string
		var opcodeType string
		switch logical_data_file {
		case "opcodes_cb":
			shift_op = "SHIFT_0xCB+" + number
			opcodeType = "CB"
		case "opcodes_ed":
			shift_op = "SHIFT_0xED+" + number
			opcodeType = "ED"
		case "opcodes_dd":
			shift_op = "SHIFT_0xDD+" + number
			opcodeType = "DD"
		case "opcodes_fd":
			shift_op = "SHIFT_0xFD+" + number
			opcodeType = "FD"
		case "opcodes_ddfdcb":
			shift_op = "SHIFT_0xDDCB+" + number
			opcodeType = "DDCB"
		default:
			shift_op = number
			opcodeType = ""
		}

		var comment string
		if opcode != "" {
			comment = "/* " + opcode
			if arguments != "" {
				comment += " " + arguments
			}
			if extra != "" {
				comment += " " + extra
			}
			comment += " */"
			ln(comment)
		} else {
			fallthrough_cases = append(fallthrough_cases, shift_op)
			continue
		}

		functionName := "dis_instr" + opcodeType + "__" + turnIntoIdentifier(strings.TrimSpace(opcode+" "+arguments+" "+extra))
		functionName = strings.Replace(functionName, "ixH", "IXH", -1)
		functionName = strings.Replace(functionName, "ixL", "IXL", -1)
		functionName = strings.Replace(functionName, "iyH", "IYH", -1)
		functionName = strings.Replace(functionName, "iyL", "IYL", -1)
		functionName = strings.Replace(functionName, "REGISTER", "REG", -1)
		ln("OpcodesDisMap[", shift_op, "] = ", functionName)

		outputStream = functions
		{
			ln(comment)
			ln("func ", functionName, "(memory MemoryReader, address uint16, shift int) (string, uint16, int) {")
			fmt.Println(args)
			ln("result := \"", opcode ," \"")
			if opcode == "shift" {
				ln("shift = SHIFT_0x" + strings.ToUpper(number)[2:])
			} else {
				ln("shift = 0")
			}

			result(args...)

			ln("return result, address + 1, shift")
			ln("}")

			if len(fallthrough_cases) > 0 {
				outputStream = code
				{
					ln("// Fallthrough cases")
					for _, fallthrough_case := range fallthrough_cases {
						ln("OpcodesDisMap[", fallthrough_case, "] = OpcodesDisMap[", shift_op, "]")
					}
					fallthrough_cases = fallthrough_cases[0:0]
				}
				outputStream = functions
			}
		}
		outputStream = code
	}

	outputStream = nil
}

// Main program
func main() {
	data_files := [][2]string{
		{"opcodes_base", "opcodes_base"},
		{"opcodes_cb", "opcodes_cb"},
		{"opcodes_ed", "opcodes_ed"},
		{"opcodes_ddfd", "opcodes_dd"},
		{"opcodes_ddfd", "opcodes_fd"},
		{"opcodes_ddfdcb", "opcodes_ddfdcb"},
	}

	mapping := make(map[string]string)

	// Buffer for implementations (source code in text form) of generated functions
	var functions bytes.Buffer

	for _, data_file := range data_files {
		var code bytes.Buffer
		var functions1 bytes.Buffer
		processDataFile(data_file[0]+".dat", data_file[1], &code, &functions1)

		codeStr := code.String()
		fnStr := functions1.String()

		mapping[data_file[1]] = codeStr

		switch data_file[1] {
		case "opcodes_base":
			fnStr_base := strings.Replace(fnStr, "SetSPHSPL", "SetSP", -1)
			functions.WriteString(fnStr_base)
		case "opcodes_dd":
			fnStr_dd := strings.Replace(fnStr, "REGISTER", "ix", -1)
			fnStr_dd = strings.Replace(fnStr_dd, "register", "ix", -1)
			fnStr_dd = strings.Replace(fnStr_dd, "ix()", "IX()", -1)
			fnStr_dd = strings.Replace(fnStr_dd, "SetixHixL", "SetIX", -1)
			fnStr_dd = strings.Replace(fnStr_dd, "IncixH", "IncIXH", -1)
			fnStr_dd = strings.Replace(fnStr_dd, "DecixH", "DecIXH", -1)
			fnStr_dd = strings.Replace(fnStr_dd, "IncixL", "IncIXL", -1)
			fnStr_dd = strings.Replace(fnStr_dd, "DecixL", "DecIXL", -1)
			fnStr_dd = strings.Replace(fnStr_dd, "z80.ix()", "z80.IX()", -1)
			fnStr_dd = strings.Replace(fnStr_dd, "ixH", "IXH", -1)
			fnStr_dd = strings.Replace(fnStr_dd, "ixL", "IXL", -1)
			functions.WriteString(fnStr_dd)
		case "opcodes_fd":
			fnStr_fd := strings.Replace(fnStr, "REGISTER", "iy", -1)
			fnStr_fd = strings.Replace(fnStr_fd, "register", "iy", -1)
			fnStr_fd = strings.Replace(fnStr_fd, "iy()", "IY()", -1)
			fnStr_fd = strings.Replace(fnStr_fd, "SetiyHiyL", "SetIY", -1)
			fnStr_fd = strings.Replace(fnStr_fd, "InciyH", "IncIYH", -1)
			fnStr_fd = strings.Replace(fnStr_fd, "DeciyH", "DecIYH", -1)
			fnStr_fd = strings.Replace(fnStr_fd, "InciyL", "IncIYL", -1)

			fnStr_fd = strings.Replace(fnStr_fd, "DeciyL", "DecIYL", -1)
			fnStr_fd = strings.Replace(fnStr_fd, "z80.iy()", "z80.IY()", -1)
			fnStr_fd = strings.Replace(fnStr_fd, "iyH", "IYH", -1)
			fnStr_fd = strings.Replace(fnStr_fd, "iyL", "IYL", -1)
			functions.WriteString(fnStr_fd)
		default:
			functions.WriteString(fnStr)
		}
	}

	mapping["functions"] = functions.String()

	w, err := os.Create("opcodes_dis_gen.go")
	if err != nil {
		panic(err.Error())
	}

	// Execute the template in file "opcodes_gen.go.template"
	{
		t := template.New("opcodes_dis_gen.go.template")
		t.Delims("[[", "]]")
		t, err = t.ParseFiles("opcodes_dis_gen.go.template")

		if err != nil {
			panic(err.Error())
		}

		err = t.Execute(w, mapping)
		if err != nil {
			panic(err.Error())
		}
	}

	err = w.Close()
	if err != nil {
		panic(err.Error())
	}
}
