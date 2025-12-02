// Package redovalnica implements a simple grade book for students.
//
// The package provides a `Student` type and functions to initialize the Student structure, add grades,
// compute averages and print the grade book and final results.
//
// The application optionally uses 3 switches:
//   - stOcen, which defines the lowest amount of grades needed for a passing grade (default: 6)
//   - minOcena, which defines the lowest possible grade (default: 5)
//   - maxOcena, which defines the highest possible grade (default: 10)
//
// Behaviour notes:
//   - `DodajOceno` validates that the grade is between 'minOcena' and 'maxOcena' and
//     prints a message when the student does not exist or the grade is
//     out of range.
//   - `povprecje` returns -1.0 when the student is not found and 0.0 when
//     the student has fewer than 6 grades.
package redovalnica

import "fmt"

// Student represents a student in the grade book.
//
// Fields are unexported because the package provides functions that
// operate on maps of `Student` values. Use the package functions to
// manipulate student grades.
type Student struct {
	ime     string
	priimek string
	ocene   []int
}

// NewStudent returns a new Student structure with the initialized parameters
// 'ime', 'priimek' and 'ocene'
func NewStudent(ime, priimek string, ocene []int) Student {
	return Student{ime, priimek, ocene}
}

// DodajOceno adds a grade `ocena` to the student identified by
// `vpisnaStevilka` in the `studenti` map.
//
// The function validates that `ocena` is in the range minOcena..maxOcena. If the
// grade is outside that range it prints "Ocena ni v ustreznem območju"
// and returns without modifying the map. If the student is not present
// in the map it prints "Študenta ni na seznamu" and returns.
func DodajOceno(studenti map[string]Student, vpisnaStevilka string, ocena, minOcena, maxOcena int) {
	if ocena < minOcena || ocena > maxOcena {
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

// povprecje returns the average grade for the student with
// `vpisnaStevilka` from the `studenti` map.
//
// If the student is not present the function returns -1.0. If the
// student has fewer than 6 grades it returns 0.0. Otherwise it returns
// the arithmetic mean of all recorded grades as a float64.
func povprecje(studenti map[string]Student, vpisnaStevilka string, stOcen int) float64 {

	s, ok := studenti[vpisnaStevilka]
	if !ok {
		return -1.0
	}
	if len(s.ocene) < stOcen {
		return 0.0
	}
	var avg float64 = 0.0
	for _, v := range s.ocene {
		avg = avg + float64(v)
	}
	return avg / float64(len(s.ocene))

}

// IzpisRedovalnice prints the contents of the grade book.
//
// Output format:
//
// REDOVALNICA:
// 63210001 - Ana Novak: [10 9 8]
// 63210002 - Boris Kralj: [6 7 5 8]
// 63210003 - Janez Novak: [4 5 3 5]
func IzpisRedovalnice(studenti map[string]Student) {
	fmt.Println("REDOVALNICA:")
	for vpisna, s := range studenti {
		fmt.Printf("%s - %s %s: ", vpisna, s.ime, s.priimek)
		fmt.Println(s.ocene)
	}
}

// IzpisiKoncniUspeh prints each student's name and their final
// performance evaluation based on the average grade computed by
// `povprecje`.
//
// For each student the function prints the student's name and a
// message according to the average grade:
//  - >= 9.0 : "Odličen študent!"
//  - < 6.0  : "Neuspešen študent"
//  - otherwise: "Povprečen študent"
func IzpisiKoncniUspeh(studenti map[string]Student, stOcen, minOcena, maxOcena int) {
	for vpisna, s := range studenti {
		avg := povprecje(studenti, vpisna, stOcen)
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
