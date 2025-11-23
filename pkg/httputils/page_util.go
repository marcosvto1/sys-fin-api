package httputils

import (
	"fmt"
	"strconv"
)

func GetRequestPaginator(page, size string) (int, int, []error) {
	pageNumber := 1
	pageSize := 10

	var errList []error
	var err error

	if page != "" {
		pageNumber, err = strconv.Atoi(page)
		if err != nil || pageNumber < 1 {
			errList = append(errList, fmt.Errorf("invalid page number"))

			return pageNumber, pageSize, errList
		}
	}
	if size != "" {
		pageSize, err = strconv.Atoi(size)
		if err != nil || pageSize < 1 {
			errList = append(errList, fmt.Errorf("invalid page size"))
			return pageNumber, pageSize, errList
		}
	}

	return pageNumber, pageSize, nil
}
