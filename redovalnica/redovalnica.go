package redovalnica

import "fmt"

type Student struct {
	ime     string
	priimek string
	ocene   []int
}

func DodajOceno(studenti map[string]Student, vpisnaStevilka string, ocena int) {
	if ocena < 0 || ocena > 10 {
		fmt.Println("Ocena ni v ustreznem območju")
		return
	}
	s, ok := studenti[vpisnaStevilka]
	if !ok {
		fmt.Println("Študenta ni na seznamu")
		return
	}
	s.ocene = append(s.ocene, ocena)
	studenti[vpisnaStevilka] = s
}

func povprecje(studenti map[string]Student, vpisnaStevilka string) float64 {

	s, ok := studenti[vpisnaStevilka]
	if !ok {
		return -1.0
	}
	if len(s.ocene) < 6 {
		return 0.0
	}
	var avg float64 = 0.0
	for _, v := range s.ocene {
		avg = avg + float64(v)
	}
	return avg / float64(len(s.ocene))

}

func IzpisRedovalnice(studenti map[string]Student) {
	fmt.Println("REDOVALNICA:")
	for vpisna, s := range studenti {
		fmt.Printf("%s - %s %s: ", vpisna, s.ime, s.priimek)
		fmt.Println(s.ocene)
	}
}

func IzpisiKoncniUspeh(studenti map[string]Student) {
	for vpisna, s := range studenti {
		avg := povprecje(studenti, vpisna)
		fmt.Printf("%s %s: ", s.ime, s.priimek)
		if avg >= 9 {
			fmt.Printf("povprečna ocena %f -> Odličen študent!\n", avg)
		} else if avg < 6 {
			fmt.Printf("povprečna ocena %f -> Neuspešen študent\n", avg)
		} else {
			fmt.Printf("povprečna ocena %f -> Povprečen študent\n", avg)
		}
	}
}
