package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Pair struct {
	A, B string
}

type Relation struct {
	pairs    []Pair
	elements map[string]bool
	matrix   map[string]map[string]bool
}

func NewRelation(pairs []Pair) *Relation {
	r := &Relation{
		pairs:    pairs,
		elements: make(map[string]bool),
		matrix:   make(map[string]map[string]bool),
	}

	for _, p := range pairs {
		r.elements[p.A] = true
		r.elements[p.B] = true

		if r.matrix[p.A] == nil {
			r.matrix[p.A] = make(map[string]bool)
		}
		r.matrix[p.A][p.B] = true
	}

	return r
}

func (r *Relation) GetElements() []string {
	elements := make([]string, 0, len(r.elements))
	for e := range r.elements {
		elements = append(elements, e)
	}
	return elements
}

func (r *Relation) Contains(a, b string) bool {
	if r.matrix[a] == nil {
		return false
	}
	return r.matrix[a][b]
}

func (r *Relation) Analyze() {
	elements := r.GetElements()

	fmt.Printf("\nHimpunan relasi: %v\n", r.pairs)
	fmt.Printf("Elemen: %v\n\n", elements)

	r.checkReflexive(elements)
	r.checkSymmetric(elements)
	r.checkAsymmetric(elements)
	r.checkAntisymmetric(elements)
	r.checkTransitive(elements)
	r.checkIrreflexive(elements)
}

func (r *Relation) checkReflexive(elements []string) {
	missing := []string{}
	for _, a := range elements {
		if !r.Contains(a, a) {
			missing = append(missing, a)
		}
	}

	if len(missing) == 0 {
		fmt.Println("✓ Refleksif: Ya")
	} else {
		fmt.Println("✗ Refleksif: Tidak")
		fmt.Printf("  Karena: %s tidak berelasi dengan diri sendiri\n", missing)
	}
}

func (r *Relation) checkSymmetric(elements []string) {
	violations := []Pair{}
	for _, a := range elements {
		for _, b := range elements {
			if a != b && r.Contains(a, b) && !r.Contains(b, a) {
				violations = append(violations, Pair{a, b})
			}
		}
	}

	if len(violations) == 0 {
		fmt.Println("✓ Simetris: Ya")
	} else {
		fmt.Println("✗ Simetris: Tidak")
		for _, p := range violations {
			fmt.Printf("  Karena: ada (%s,%s) tapi tidak ada (%s,%s)\n", p.A, p.B, p.B, p.A)
		}
	}
}

func (r *Relation) checkAsymmetric(elements []string) {
	violations := []Pair{}
	hasReflexive := false

	for _, a := range elements {
		if r.Contains(a, a) {
			hasReflexive = true
		}
		for _, b := range elements {
			if r.Contains(a, b) && r.Contains(b, a) {
				violations = append(violations, Pair{a, b})
			}
		}
	}

	if len(violations) == 0 && !hasReflexive {
		fmt.Println("✓ Asimetris: Ya")
	} else {
		fmt.Println("✗ Asimetris: Tidak")
		if hasReflexive {
			fmt.Println("  Karena: ada elemen yang berelasi dengan diri sendiri")
		}
		for _, p := range violations {
			if p.A != p.B {
				fmt.Printf("  Karena: (%s,%s) dan (%s,%s) sama-sama ada\n", p.A, p.B, p.B, p.A)
			}
		}
	}
}

func (r *Relation) checkAntisymmetric(elements []string) {
	violations := []Pair{}
	for _, a := range elements {
		for _, b := range elements {
			if a != b && r.Contains(a, b) && r.Contains(b, a) {
				violations = append(violations, Pair{a, b})
			}
		}
	}

	if len(violations) == 0 {
		fmt.Println("✓ Anti-simetris: Ya")
	} else {
		fmt.Println("✗ Anti-simetris: Tidak")
		for _, p := range violations {
			fmt.Printf("  Karena: (%s,%s) dan (%s,%s) sama-sama ada\n", p.A, p.B, p.B, p.A)
		}
	}
}

func (r *Relation) checkTransitive(elements []string) {
	violations := []struct{ a, b, c string }{}

	for _, a := range elements {
		for _, b := range elements {
			if r.Contains(a, b) {
				for _, c := range elements {
					if r.Contains(b, c) && !r.Contains(a, c) {
						violations = append(violations, struct{ a, b, c string }{a, b, c})
					}
				}
			}
		}
	}

	if len(violations) == 0 {
		fmt.Println("✓ Transitif: Ya")
	} else {
		fmt.Println("✗ Transitif: Tidak")
		for _, v := range violations {
			fmt.Printf("  Karena: ada (%s,%s) dan (%s,%s) tapi tidak ada (%s,%s)\n",
				v.a, v.b, v.b, v.c, v.a, v.c)
		}
	}
}

func (r *Relation) checkIrreflexive(elements []string) {
	reflexiveElements := []string{}
	for _, a := range elements {
		if r.Contains(a, a) {
			reflexiveElements = append(reflexiveElements, a)
		}
	}

	if len(reflexiveElements) == 0 {
		fmt.Println("✓ Irrefleksif: Ya")
	} else {
		fmt.Println("✗ Irrefleksif: Tidak")
		fmt.Printf("  Karena: %s berelasi dengan diri sendiri\n", reflexiveElements)
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("=== ANALISIS SIFAT RELASI ===")
	fmt.Println("Format input: (a,b) (c,d) (e,f)")
	fmt.Println("Contoh: (1,1) (1,2) (2,1) (2,2)")
	fmt.Print("\nMasukkan himpunan relasi: ")

	scanner.Scan()
	input := scanner.Text()

	pairs := []Pair{}
	parts := strings.Fields(input)

	for _, part := range parts {
		part = strings.TrimSpace(part)
		if len(part) < 5 || part[0] != '(' || part[len(part)-1] != ')' {
			continue
		}

		content := part[1 : len(part)-1]
		elements := strings.Split(content, ",")

		if len(elements) == 2 {
			a := strings.TrimSpace(elements[0])
			b := strings.TrimSpace(elements[1])
			pairs = append(pairs, Pair{A: a, B: b})
		}
	}

	if len(pairs) == 0 {
		fmt.Println("Error: Tidak ada pasangan yang valid ditemukan")
		fmt.Println("Pastikan format input sesuai contoh: (a,b) (c,d)")
		return
	}

	r := NewRelation(pairs)
	r.Analyze()
}
