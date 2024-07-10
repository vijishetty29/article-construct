package drawConstruct

import (
	"fmt"
	svg "github.com/ajstarks/svgo"
	"github.com/user/article-construct-demo/internal/model"
	"net/http"
)

func DrawConstruct(ct model.Construct, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "image/svg+xml")

	fmt.Printf("Item is %s\n", ct)
	fmt.Print(ct)

	width := 1200
	height := 1000
	canvas := svg.New(w)
	canvas.Start(width, height)
	canvas.Title("Article Construct")

	numOfVariants := totalUniqueVariants(ct)
	fmt.Println(numOfVariants)
	numOfCases := len(ct.Cases)

	columnWidthLevel3 := width / numOfVariants
	columnWidthLevel2 := width / numOfCases
	columnWidthLevel1 := width / 2
	columnHeightLevel1 := 200
	columnHeightLevel2 := 450
	columnHeightLevel3 := 700

	canvas.Link("https://www.google.de", ct.Ian)
	canvas.Roundrect(columnWidthLevel1, columnHeightLevel1, 100, 50, 10, 10, "stroke:#68A328;stroke-width:3;fill:#73B72D")
	canvas.Text(columnWidthLevel1+50, columnHeightLevel1+25, ct.Ian, "text-anchor:middle;font-size:12px;fill:black")
	canvas.LinkEnd()

	cwh := columnWidthLevel2 / 2
	cwv := columnWidthLevel3 / 2
	variantsDrawn := make(map[string]int)

	rectangleWidth := 100
	if cwh < 100 {
		rectangleWidth = cwh
	}
	rectangleWidthV := 100
	if cwv < 100 {
		rectangleWidthV = cwv
	}
	for i := 0; i < len(ct.Cases); i++ {

		canvas.Line(columnWidthLevel1+50, columnHeightLevel1+50, cwh+rectangleWidth/2, columnHeightLevel2, "stroke:#73B72D;stroke-width:2")
		canvas.Link("https://www.google.de", ct.Cases[i].Ian)
		canvas.Roundrect(cwh, columnHeightLevel2, rectangleWidth, 50, 10, 10, "stroke:#6FA4EF;stroke-width:3;fill:#97C2FC")
		canvas.Text(cwh+rectangleWidth/2, columnHeightLevel2+25, ct.Cases[i].Ian, "text-anchor:middle;font-size:12px;fill:black")
		canvas.LinkEnd()
		variants := ct.Cases[i].Variants

		for _, s := range variants {

			if val, found := variantsDrawn[s]; found {
				canvas.Line(cwh+50, columnHeightLevel2+50, val, columnHeightLevel3, "stroke:#97C2FC;stroke-width:2")
			} else {
				canvas.Line(cwh+50, columnHeightLevel2+50, cwv+rectangleWidthV/2, columnHeightLevel3, "stroke:#97C2FC;stroke-width:2")
				canvas.Link("https://www.google.de", ct.Cases[i].Ian)
				canvas.Roundrect(cwv, columnHeightLevel3, rectangleWidthV, 50, 10, 10, "stroke:#6FA4EF;stroke-width:3;fill:#97C2FC")
				canvas.Text(cwv+rectangleWidthV/2, columnHeightLevel3+25, s, "text-anchor:middle;font-size:12px;fill:black")
				canvas.LinkEnd()
				variantsDrawn[s] = cwv + rectangleWidthV/2
				cwv = cwv + columnWidthLevel3
			}

		}
		cwh = cwh + columnWidthLevel2
	}
	canvas.End()
}

func totalUniqueVariants(ct model.Construct) int {
	keys := make(map[string]bool)
	list := []string{}

	// If the key(values of the slice) is not equal
	// to the already present value in new slice (list)
	// then we append it. else we jump on another element.
	for i := 0; i < len(ct.Cases); i++ {
		for _, entry := range ct.Cases[i].Variants {
			if _, value := keys[entry]; !value {
				keys[entry] = true
				list = append(list, entry)
			}
		}
	}
	return len(list)
}
