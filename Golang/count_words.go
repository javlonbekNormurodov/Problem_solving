package main

import "fmt"

func countWords(words1 []string, words2 []string) int {
	sum := 0
	m := make(map[string]int)
	s := make(map[string]int)
	for _, v := range words1 {
		if m[v] == 0 {
			m[v] = 1
		} else if m[v] == 1 {
			m[v] = 3
		} else {
			m[v] = 0
		}
	}
	for _, v := range words2 {
		if s[v] == 0 {
			s[v] = 1
		} else if s[v] == 1 {
			s[v] = 3
		} else {
			s[v] = 0
		}
	}
	fmt.Println("m => ", m)
	fmt.Println("s => ", s)
	if len(m) > len(s) {
		for i := range m {
			if m[i] == 1 && s[i] == 1 {
				sum++
			}
		}
	} else {
		for i := range s {
			if s[i] == 1 && m[i] == 1 {
				sum++
			}
		}
	}

	return sum
}

func main() {
	words1 := []string{"tnzaifelzbtrprqawvtx", "tnzaifelzbtrprqawvtx", "tnzaifelzbtrprqawvtx", "tnzaifelzbtrprqawvtx",
		"digprmnrigylafrklfrfdwbfc", "digprmnrigylafrklfrfdwbfc", "digprmnrigylafrklfrfdwbfc", "digprmnrigylafrklfrfdwbfc",
		"digprmnrigylafrklfrfdwbfc", "mmhcnybfnftqubvxk", "mmhcnybfnftqubvxk", "mmhcnybfnftqubvxk", "hvnqssvjocczpzqbqfl",
		"agpmawruhuszulr", "agpmawruhuszulr", "agpmawruhuszulr", "dxvawtnrnifqinrugec", "dxvawtnrnifqinrugec", "dxvawtnrnifqinrugec",
		"dxvawtnrnifqinrugec", "dxvawtnrnifqinrugec",
		"jhlgpaptxysljmzfvzulhjzmi", "kclpodfwwhwhlmujnvieyi", "kclpodfwwhwhlmujnvieyi", "kclpodfwwhwhlmujnvieyi",
		"kclpodfwwhwhlmujnvieyi", "kclpodfwwhwhlmujnvieyi", "wkysonbcmexhpmbtcghzwxzkk", "wkysonbcmexhpmbtcghzwxzkk",
		"wkysonbcmexhpmbtcghzwxzkk", "wkysonbcmexhpmbtcghzwxzkk", "wkysonbcmexhpmbtcghzwxzkk", "ddhxlqhitxjyzxrsie",
		"ddhxlqhitxjyzxrsie", "ddhxlqhitxjyzxrsie", "ddhxlqhitxjyzxrsie", "kdenfbwmpwdmmdqcjad", "kdenfbwmpwdmmdqcjad",
		"xmwxnntkfnonezwozvnajhcztihqe", "xmwxnntkfnonezwozvnajhcztihqe", "xmwxnntkfnonezwozvnajhcztihqe"}
	words2 := []string{"vzkz", "w", "dlayelrwwh", "lycjdinyndbkzraxujyerecrrubmk", "fdc", "w",
		"vtzkwxfiufcuqlrqvatlvc", "dlayelrwwh", "drwrlrchhjnyihznkbqoihwh", "dlayelrwwh", "lsznkshrmisgipyubywnq",
		"fcaqmwvuyivkwoh", "qhboqjsgjuvqjtzqxhiwkx", "gqhhm", "xmwxnntkfnonezwozvnajhcztihqe", "tnzaifelzbtrprqawvtx",
		"fevvqophwbpeuvhymymiuhxrl", "aysjqneyeerodnizdfestwars", "lsznkshrmisgipyubywnq", "wkysonbcmexhpmbtcghzwxzkk",
		"leuhtskblydvsgiosmf", "agpmawruhuszulr", "digprmnrigylafrklfrfdwbfc", "hutltx", "dxvawtnrnifqinrugec", "uqwqtrvlsa",
		"mmhcnybfnftqubvxk", "kclpodfwwhwhlmujnvieyi", "w", "kboxeyogqc", "extlnruqgdr", "dereqou", "ffsrsbmchi",
		"raygvlgudpdldjotvfmd", "pkdnctly", "rjgt", "iertztrjcalbuiprftltimg", "cqbiddmptsemkehllsy", "irzvpvbpqfyidzmkymypukpqmbsw",
		"pejzio"}
	fmt.Println(countWords(words1, words2))
}
