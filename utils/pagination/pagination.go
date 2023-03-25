package paginateUtils

func Pagination(page int64, perPage int64) (int64, int64, int64, int64) {
	var _page int64 = 1
	var _perPage int64 = 10
	if page != 0 {
		_page = page
	}

	if perPage != 0 {
		_perPage = perPage
	}

	var skip = (_page - 1) * _perPage

	return skip, _perPage, _page, _perPage
}

func No(index, page int64, perPage int64) int64 {
	return index + 1 + ((page)-1)*(perPage)
}
