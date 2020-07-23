package utilities

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"golang.org/x/sync/errgroup"
)

type Ebook struct {
	URL   string `json:"url"`
	Title string `json:"title"`
	Image string `json:"image"`
}

type Ebooks struct {
	TotalPages  int     `json:"total_pages"`
	TotalEbooks int     `json:"total_ebooks"`
	List        []Ebook `json:"ebooks"`
}

func NewEbooks() *Ebooks {
	return &Ebooks{}
}

func (ebooks *Ebooks) getEbooksByUrl(url string) error {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return err
	}

	doc.Find(".section ul li.learn-outline-item").Each(func(i int, s *goquery.Selection) {
		docTitle, exists := s.Find("a").Attr("title")
		if !exists {
			docTitle = ""
		}
		fmt.Println("docTitle",docTitle)
		docLink, exists := s.Find("a").Attr("href")
		if !exists {
			docLink = "#"
		}
		fmt.Println("docLink",docLink)

		docImg, exists := s.Find("a").Attr("id")
		if !exists {
			docImg = ""
		}
		fmt.Println("docImg",docImg)

		Ebook := Ebook{
			URL:   docLink,
			Title: docTitle,
			Image: docImg,
		}
		ebooks.TotalEbooks++
		ebooks.List = append(ebooks.List, Ebook)
	})
	return nil
}

func (ebooks *Ebooks) GetTotalPages(url string) error {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return err
	}
	lastPageLink, _ := doc.Find(".section ul .learn-outline-item  a.learn-lesson-wr ").Attr("href")
	if lastPageLink == "javascript:void();" {
		ebooks.TotalPages = 1
		return nil
	}
	fmt.Println("lastPageLink ............",lastPageLink)

	//split := strings.Split(lastPageLink, "?p=")
	//totalPages, _ := strconv.Atoi(split[1])
	ebooks.TotalPages = 1
	return nil
}

func (ebooks *Ebooks) GetAllEbooks(currentUrl string) error {
	eg := errgroup.Group{}
	if ebooks.TotalPages > 0 {
			eg.Go(func() error {
				err := ebooks.getEbooksByUrl(currentUrl)
				if err != nil {
					return err
				}
				return nil
			})
		if err := eg.Wait(); err != nil {
			return err
		}
	}
	return nil
}
