// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pagination

import (
	"net/http"

	"github.com/Influship/influship-go/internal/apijson"
	"github.com/Influship/influship-go/internal/requestconfig"
	"github.com/Influship/influship-go/option"
	"github.com/Influship/influship-go/packages/param"
	"github.com/Influship/influship-go/packages/respjson"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

type QueryCursor[T any] struct {
	NextCursor string `json:"next_cursor" api:"nullable"`
	HasMore    bool   `json:"has_more"`
	Data       []T    `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		NextCursor  respjson.Field
		HasMore     respjson.Field
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	cfg *requestconfig.RequestConfig
	res *http.Response
}

// Returns the unmodified JSON received from the API
func (r QueryCursor[T]) RawJSON() string { return r.JSON.raw }
func (r *QueryCursor[T]) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *QueryCursor[T]) GetNextPage() (res *QueryCursor[T], err error) {
	if len(r.Data) == 0 {
		return nil, nil
	}

	if r.JSON.HasMore.Valid() && r.HasMore == false {
		return nil, nil
	}
	next := r.NextCursor
	if len(next) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	err = cfg.Apply(option.WithQuery("cursor", next))
	if err != nil {
		return nil, err
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *QueryCursor[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &QueryCursor[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type QueryCursorAutoPager[T any] struct {
	page *QueryCursor[T]
	cur  T
	idx  int
	run  int
	err  error
	paramObj
}

func NewQueryCursorAutoPager[T any](page *QueryCursor[T], err error) *QueryCursorAutoPager[T] {
	return &QueryCursorAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *QueryCursorAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Data) == 0 {
		return false
	}
	if r.idx >= len(r.page.Data) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Data) == 0 {
			return false
		}
	}
	r.cur = r.page.Data[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *QueryCursorAutoPager[T]) Current() T {
	return r.cur
}

func (r *QueryCursorAutoPager[T]) Err() error {
	return r.err
}

func (r *QueryCursorAutoPager[T]) Index() int {
	return r.run
}

type BodyCursor[T any] struct {
	NextCursor string `json:"next_cursor" api:"nullable"`
	HasMore    bool   `json:"has_more"`
	Data       []T    `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		NextCursor  respjson.Field
		HasMore     respjson.Field
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	cfg *requestconfig.RequestConfig
	res *http.Response
}

// Returns the unmodified JSON received from the API
func (r BodyCursor[T]) RawJSON() string { return r.JSON.raw }
func (r *BodyCursor[T]) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *BodyCursor[T]) GetNextPage() (res *BodyCursor[T], err error) {
	if len(r.Data) == 0 {
		return nil, nil
	}

	if r.JSON.HasMore.Valid() && r.HasMore == false {
		return nil, nil
	}
	next := r.NextCursor
	if len(next) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	err = cfg.Apply(option.WithQuery("cursor", next))
	if err != nil {
		return nil, err
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *BodyCursor[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &BodyCursor[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type BodyCursorAutoPager[T any] struct {
	page *BodyCursor[T]
	cur  T
	idx  int
	run  int
	err  error
	paramObj
}

func NewBodyCursorAutoPager[T any](page *BodyCursor[T], err error) *BodyCursorAutoPager[T] {
	return &BodyCursorAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *BodyCursorAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Data) == 0 {
		return false
	}
	if r.idx >= len(r.page.Data) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Data) == 0 {
			return false
		}
	}
	r.cur = r.page.Data[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *BodyCursorAutoPager[T]) Current() T {
	return r.cur
}

func (r *BodyCursorAutoPager[T]) Err() error {
	return r.err
}

func (r *BodyCursorAutoPager[T]) Index() int {
	return r.run
}
