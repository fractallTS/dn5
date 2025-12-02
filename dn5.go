package main

import (
	"context"
	"log"
	"os"

	"github.com/fractallTS/dn5/redovalnica"
	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:  "dn5",
		Usage: "Dodajanje ocen študentov za posamezne domače naloge in izračun končne ocene nalog.",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:  "stOcen",
				Usage: "Definira najmanjše število ocen potrebnih za pozitivno oceno",
				Value: 6,
			},
			&cli.IntFlag{
				Name:  "minOcena",
				Usage: "Definira najmanjšo možno oceno",
				Value: 5,
			},
			&cli.IntFlag{
				Name:  "maxOcena",
				Usage: "Definira največje možno oceno",
				Value: 10,
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			stOcen := cmd.Int("stOcen")
			minOcena := cmd.Int("minOcena")
			maxOcena := cmd.Int("maxOcena")
			return Ocene(stOcen, minOcena, maxOcena)
		},
	}
	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}

func Ocene(stOcen, minOcena, maxOcena int) error {
	studenti := make(map[string]redovalnica.Student)
	studenti["63210001"] = redovalnica.NewStudent("Ana", "Novak", []int{10, 9, 8})
	studenti["63210002"] = redovalnica.NewStudent("Boris", "Kralj", []int{6, 7, 5, 8, 9, 10})
	studenti["63210003"] = redovalnica.NewStudent("Janez", "Novak", []int{4, 5, 3, 5, 6, 5})

	redovalnica.DodajOceno(studenti, "63210000", 8, minOcena, maxOcena)
	redovalnica.DodajOceno(studenti, "63210001", 11, minOcena, maxOcena)
	redovalnica.DodajOceno(studenti, "63210003", 2, minOcena, maxOcena)
	redovalnica.DodajOceno(studenti, "63210002", 7, minOcena, maxOcena)
	redovalnica.IzpisRedovalnice(studenti)
	redovalnica.IzpisiKoncniUspeh(studenti, stOcen, minOcena, maxOcena)
	return nil
}
