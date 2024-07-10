package drawConstruct

import (
	"fmt"
	svg "github.com/ajstarks/svgo"
	"github.com/user/article-construct-demo/internal/model"
	"net/http"
)

func DrawItemConstruct(item *model.Item, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "image/svg+xml")

	fmt.Printf("Item is %s\n", item)
	fmt.Print(item)

	width := 1200
	height := 1000
	canvas := svg.New(w)
	canvas.Start(width, height)
	canvas.Title("Article Construct")

	numOfVariants := totalUniqueVariantsIan(item)
	fmt.Println(numOfVariants)
	numOfCases := len(item.Cases)

	columnWidthLevel3 := width / numOfVariants
	columnWidthLevel2 := width / numOfCases
	columnWidthLevel1 := width / 2
	columnHeightLevel1 := 200
	columnHeightLevel2 := 450
	columnHeightLevel3 := 700

	canvas.Link("http://localhost:8181/item?ian="+item.Ian, item.Ian)
	canvas.Roundrect(columnWidthLevel1, columnHeightLevel1, 100, 50, 10, 10, "stroke:#68A328;stroke-width:3;fill:#73B72D")
	canvas.Text(columnWidthLevel1+50, columnHeightLevel1+25, item.Ian, "text-anchor:middle;font-size:12px;fill:black")
	canvas.LinkEnd()

	caseWidth := columnWidthLevel2 / 2
	variantWidth := columnWidthLevel3 / 2
	variantsDrawn := make(map[string]int)

	rectangleWidth := 100
	if caseWidth < 100 {
		rectangleWidth = caseWidth
	}
	rectangleWidthV := 100
	if variantWidth < 100 {
		rectangleWidthV = variantWidth
	}
	for i := 0; i < len(item.Cases); i++ {

		canvas.Line(columnWidthLevel1+50, columnHeightLevel1+50, caseWidth+rectangleWidth/2, columnHeightLevel2, "stroke:#73B72D;stroke-width:2")
		canvas.Link("http://localhost:8181/case?ian="+item.Cases[i].Ian, item.Cases[i].Ian)

		canvas.Roundrect(caseWidth, columnHeightLevel2, rectangleWidth, 50, 10, 10, "stroke:#6FA4EF;stroke-width:3;fill:#97C2FC")
		canvas.Text(caseWidth+rectangleWidth/2, columnHeightLevel2+25, item.Cases[i].Ian, "text-anchor:middle;font-size:12px;fill:black")
		canvas.LinkEnd()
		variants := item.Cases[i].Variants

		for _, variant := range variants {

			if val, found := variantsDrawn[variant.Ian]; found {
				canvas.Line(caseWidth+50, columnHeightLevel2+50, val, columnHeightLevel3, "stroke:#97C2FC;stroke-width:2")
			} else {
				canvas.Line(caseWidth+50, columnHeightLevel2+50, variantWidth+rectangleWidthV/2, columnHeightLevel3, "stroke:#97C2FC;stroke-width:2")
				canvas.Link("http://localhost:8181/variant?ian="+variant.Ian, variant.Ian)
				canvas.Roundrect(variantWidth, columnHeightLevel3, rectangleWidthV, 50, 10, 10, "stroke:#6FA4EF;stroke-width:3;fill:#97C2FC")
				canvas.Text(variantWidth+rectangleWidthV/2, columnHeightLevel3+25, variant.Ian, "text-anchor:middle;font-size:12px;fill:black")
				canvas.LinkEnd()
				variantsDrawn[variant.Ian] = variantWidth + rectangleWidthV/2
				variantWidth = variantWidth + columnWidthLevel3
			}

		}
		caseWidth = caseWidth + columnWidthLevel2
	}
	canvas.End()
}

func totalUniqueVariantsIan(item *model.Item) int {
	keys := make(map[string]bool)
	list := []string{}

	// If the key(values of the slice) is not equal
	// to the already present value in new slice (list)
	// then we append it. else we jump on another element.
	for i := 0; i < len(item.Cases); i++ {
		for _, entry := range item.Cases[i].Variants {
			if _, value := keys[entry.Ian]; !value {
				keys[entry.Ian] = true
				list = append(list, entry.Ian)
			}
		}
	}
	return len(list)
}
