package drawConstruct

import (
	"fmt"
	svg "github.com/ajstarks/svgo"
	"github.com/user/article-construct-demo/internal/model"
	"net/http"
)

// plotset defines plot metadata data as float64 x,y coordinates
//type plotset struct {
//	x      int
//	y      int
//	width  int
//	height int
//}
//
//var (
//	ps = plotset{0, 0, 100, 50}
//)

const (
	width           = 1200
	height          = 1000
	rectangleWidth  = 100
	rectangleHeight = 60
	rectangleBorder = 10
	baHeightLevel   = 200
	skHeightLevel   = 450
	eaHeightLevel   = 700
	canvasTitle     = "Article Construct"
	inactiveStyle   = "stroke:#a1a1a1;stroke-width:3;fill:#b3b2b2"
	baStyle         = "stroke:#68A328;stroke-width:3;fill:#73B72D"
	style           = "stroke:#6FA4EF;stroke-width:3;fill:#97C2FC"
	baLinestyle     = "stroke:#73B72D;stroke-width:2"
	linestyle       = "stroke:#97C2FC;stroke-width:2"
	textfmt         = "text-anchor:middle;font-size:12px;fill:black"
	linkText        = "IAN %s \nNAN %s"
)

func DrawItemConstruct(item *model.Item, country string, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "image/svg+xml")

	canvas := svg.New(w)
	canvas.Start(width, height)
	canvas.Title(canvasTitle)

	numOfCases := len(item.Cases)
	numOfVariants := totalUniqueVariantsIan(item)

	baWidthLevel := (width - 100) / 2
	skWidthLevel := (width - 100) / numOfCases
	eaWidthLevel := (width - 100) / numOfVariants

	canvas.Link(fmt.Sprintf("http://localhost:8181/item?ian=%s&amp;country=%s", item.Ian, country), getLinkText(item.Ian, item.Nat))
	canvas.Roundrect(baWidthLevel, baHeightLevel, rectangleWidth, rectangleHeight, rectangleBorder, rectangleBorder, nodeBAStyle(item.ItemStatus))
	canvas.Text(baWidthLevel+rectangleWidth/2, baHeightLevel+rectangleHeight/2, item.Ian, textfmt)
	if len(item.Nat) != 0 {
		canvas.Text(baWidthLevel+rectangleWidth/2, baHeightLevel+rectangleHeight/2+10, "NAN "+item.Nat, textfmt)
	}
	canvas.LinkEnd()

	skWidthCounter := skWidthLevel / 2
	eaWidthCounter := eaWidthLevel / 2

	variantsDrawn := make(map[string]int)

	for i := 0; i < len(item.Cases); i++ {
		c := item.Cases[i]
		canvas.Line(baWidthLevel+rectangleWidth/2, baHeightLevel+rectangleHeight, skWidthCounter+rectangleWidth/2, skHeightLevel, baLinestyle)
		canvas.Link(fmt.Sprintf("http://localhost:8181/case?ian=%s&amp;country=%s", c.Ian, country), getLinkText(c.Ian, c.Nat))
		canvas.Roundrect(skWidthCounter, skHeightLevel, rectangleWidth, rectangleHeight, rectangleBorder, rectangleBorder, nodeStyle(c.ItemStatus))
		canvas.Text(skWidthCounter+rectangleWidth/2, skHeightLevel+rectangleHeight/2, c.Ian, textfmt)
		if len(c.Nat) != 0 {
			canvas.Text(skWidthCounter+rectangleWidth/2, skHeightLevel+rectangleHeight/2+10, "NAN "+c.Nat, textfmt)
		}
		canvas.LinkEnd()
		variants := c.Variants

		for _, variant := range variants {
			if val, found := variantsDrawn[variant.Ian]; found {
				canvas.Line(skWidthCounter+rectangleWidth/2, skHeightLevel+rectangleHeight, val, eaHeightLevel, linestyle)
			} else {
				canvas.Line(skWidthCounter+rectangleWidth/2, skHeightLevel+rectangleHeight, eaWidthCounter+rectangleWidth/2, eaHeightLevel, linestyle)
				canvas.Link(fmt.Sprintf("http://localhost:8181/variant?ian=%s&amp;country=%s", variant.Ian, country), getLinkText(variant.Ian, variant.Nat))
				canvas.Roundrect(eaWidthCounter, eaHeightLevel, rectangleWidth, rectangleHeight, rectangleBorder, rectangleBorder, nodeStyle(variant.ItemStatus))
				canvas.Text(eaWidthCounter+rectangleWidth/2, eaHeightLevel+rectangleHeight/2, variant.Ian, textfmt)
				if len(variant.Nat) != 0 {
					canvas.Text(eaWidthCounter+rectangleWidth/2, eaHeightLevel+rectangleHeight/2+10, "NAN "+variant.Nat, textfmt)
				}
				canvas.LinkEnd()
				variantsDrawn[variant.Ian] = eaWidthCounter + rectangleWidth/2
				eaWidthCounter = eaWidthCounter + eaWidthLevel
			}

		}
		skWidthCounter = skWidthCounter + skWidthLevel
	}
	canvas.End()
}

func totalUniqueVariantsIan(item *model.Item) int {
	keys := make(map[string]bool)
	var list []string
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

func nodeStyle(status string) string {
	if status == "inactive" {
		return inactiveStyle
	}
	return style
}

func nodeBAStyle(status string) string {
	if status == "inactive" {
		return inactiveStyle
	}
	return baStyle
}

func getLinkText(ian string, nat string) string {
	if len(nat) != 0 {
		return fmt.Sprintf(linkText, ian, nat)
	}

	return "IAN " + ian
}
