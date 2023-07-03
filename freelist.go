package main

const initialPage = 0

type freelist struct {
	maxPage       pgnum
	releasedPages []pgnum
}

func newFreelist() *freelist {
	return &freelist{
		maxPage:       initialPage,
		releasedPages: []pgnum{},
	}
}

func (f *freelist) getNextPage() pgnum {
	if len(f.releasedPages) == 0 { // no released pages available
		f.maxPage++
		return f.maxPage
	}

	// return the first released page
	pageNum := f.releasedPages[0]
	f.releasedPages = f.releasedPages[1:]
	return pageNum
}

func (f *freelist) releasePage(pageNum pgnum) {
	f.releasedPages = append(f.releasedPages, pageNum)
}
