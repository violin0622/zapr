package ctx_test

import (
	"context"
	"fmt"
	"testing"

	"go.uber.org/zap"

	zcx "github.com/violin0622/zapr/ctx"
)

func TestValues(t *testing.T) {
	testcases := []struct {
		desc   string
		expect []any
	}{{
		desc: `empty`,
	}, {
		desc:   `one zap field`,
		expect: []any{zap.String(`elem`, `one`)},
	}, {
		desc:   `one any field`,
		expect: []any{`elem`, `one`},
	}}

	for i, tc := range testcases {
		t.Run(fmt.Sprintf(`%d_%s`, i, tc.desc), func(t *testing.T) {
			ctx := context.Background()
			ctx = zcx.WithValues(ctx, tc.expect...)

			actual := zcx.Values(ctx)
			if len(actual) != len(tc.expect) {
				t.Errorf(`expect != actual: %d vs %d`, len(tc.expect), len(actual))
			}
			for j := range tc.expect {
				if tc.expect[j] != actual[j] {
					t.Errorf(`ellm %d expect != actual: %+v vs %+v`, j, tc.expect[j], actual[j])
				}
			}
		})
	}
}

func TestV(t *testing.T) {
	testcases := []struct {
		desc              string
		stored, v, expect int
	}{{
		desc: `empty`,
	}, {
		desc:   `absolute`,
		v:      3,
		expect: 3,
	}, {
		desc:   `relative`,
		stored: 1,
		v:      2,
		expect: 3,
	}}

	for i, tc := range testcases {
		t.Run(fmt.Sprintf(`%d_%s`, i, tc.desc), func(t *testing.T) {
			ctx := context.Background()
			ctx = zcx.WithV(ctx, tc.stored)
			ctx = zcx.WithV(ctx, tc.v)

			actual := zcx.V(ctx)
			if actual != tc.expect {
				t.Errorf(`expect != actual: %d vs %d`, tc.expect, actual)
			}
		})
	}
}
