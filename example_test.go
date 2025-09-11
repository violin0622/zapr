package zapr_test

// func Example() {
// 	l, sync, err := zapr.NewLogger()
// 	if err != nil {
// 		fmt.Println(`err`, err)
// 		os.Exit(1)
// 	}
// 	defer sync()
// 	pkgLogger = l.WithName(`example`)
//
// 	pkgLogger.Info(`Hello world!`)
// 	defer pkgLogger.Info(`Good bye!`)
//
// 	ctx := zapr.WithValues(context.Background(), zap.String(`fieldInCtx`, `someContent`))
// 	foo(ctx)
//
// 	// Output: ok
// }
//
// func foo(ctx context.Context) {
// 	logger := zapr.Extract(ctx, pkgLogger).WithName(`foo`)
//
// 	logger.WithValues(`func`, `foo`).Info(`Some log in foo.`)
// 	logger.V(1).Info(`Some log with higher verbosity.`)
//
// 	f := newFoo()
// 	f.tasteApple(zapr.WithV(ctx, 1))
// }
//
// type Apple struct {
// 	logr   logr.Logger
// 	others string
// }
//
// func newFoo() Apple {
// 	var f Apple
// 	f.others = `structSpecificContent`
// 	f.logr = pkgLogger.WithName(`apple`).WithValues(`suggared`, `fruit`).V(1)
// 	f.logr.Info(`New apple created.`)
// 	return f
// }
//
// func (a Apple) tasteApple(ctx context.Context) {
// 	zapr.Extract(ctx, a.logr).Info(`This message would take both fields from apple and context.`)
// }
