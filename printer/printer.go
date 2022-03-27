package printer

import (
	"fmt"
	"os"
	"text/tabwriter"

	"linkedInLearning/tempService/models"
)

type Printer struct {
	w *tabwriter.Writer
}

// New initialises a new Printer instance
func New() *Printer {
	w := tabwriter.NewWriter(os.Stdout, 3, 0, 3, ' ', tabwriter.TabIndent)
	return &Printer{
		w: w,
	}
}

// CityHeader prints out the table header for the city table
func (p *Printer) CityHeader() {
	fmt.Fprintln(p.w, "Name\tTempC\tTempF")
}

// CityDetails prints out the given city
func (p *Printer) CityDetails(c *models.City) {
	fmt.Fprintf(p.w, "%v\t%v\t%v\n", c.Name, c.TempC, c.TempF())
}

// Cleanup closes up the printer instance
func (p *Printer) Cleanup() {
	p.w.Flush()
}
