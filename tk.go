// test project main.go
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"

	"github.com/robertkrimen/otto"
)

//获取tk
func tk() string {
	tkk := tkk()
	vm := otto.New()
	vm.Run(`var b = function (a, b) {
				for (var d = 0; d < b.length - 2; d += 3) {
					var c = b.charAt(d + 2),
						c = "a" <= c ? c.charCodeAt(0) - 87 : Number(c),
						c = "+" == b.charAt(d + 1) ? a >>> c : a << c;
					a = "+" == b.charAt(d) ? a + c & 4294967295 : a ^ c
				}
				return a
			}

			var tk =  function (a,TKK) {
				for (var e = TKK.split("."), h = Number(e[0]) || 0, g = [], d = 0, f = 0; f < a.length; f++) {
					var c = a.charCodeAt(f);
					128 > c ? g[d++] = c : (2048 > c ? g[d++] = c >> 6 | 192 : (55296 == (c & 64512) && f + 1 < a.length && 56320 == (a.charCodeAt(f + 1) & 64512) ? (c = 65536 + ((c & 1023) << 10) + (a.charCodeAt(++f) & 1023), g[d++] = c >> 18 | 240, g[d++] = c >> 12 & 63 | 128) : g[d++] = c >> 12 | 224, g[d++] = c >> 6 & 63 | 128), g[d++] = c & 63 | 128)
				}
				a = h;
				for (d = 0; d < g.length; d++) a += g[d], a = b(a, "+-a^+6");
				a = b(a, "+-3^+b+-f");
				a ^= Number(e[1]) || 0;
				0 > a && (a = (a & 2147483647) + 2147483648);
				a %= 1E6;
				return a.toString() + "." + (a ^ h)
			}
`)
	data, _ := vm.ToValue("about")
	tkks, _ := vm.ToValue(tkk)
	tk, _ := vm.Call("tk", nil, data, tkks)
	//fmt.Println(tk.String())
	return tk.String()

}
func tkk() string {
	tkjs := tkjs()
	runtime := otto.New()
	if _, err := runtime.Run(`function tkk(){` + tkjs + `  return TKK;}`); err != nil {
		panic(err)
	}
	TKK, err := runtime.Call("tkk", nil)
	if err != nil {
		panic(err)
	}
	//fmt.Println(TKK.String())
	return TKK.String()
	//httpGet()
}

func tkjs() string {
	resp, err := http.Get("https://translate.google.cn/")
	if err != nil {
		fmt.Println("服务器无法访问！")
		return ""
		// handle error
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ""
	}

	//fmt.Println(string(body))
	reg := regexp.MustCompile(`TKK=eval\('\(\(function\(\){var a\\x[\w]{10,14};var b\\x[\w]{2,4}-[\w]{9,14};return [\S]{15,25}}\)\(\)\)\'\);`)
	ret := reg.FindAllString(string(body), -1)
	//fmt.Printf("%q\n", ret)
	return ret[0]
}
