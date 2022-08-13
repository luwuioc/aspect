package annotation

// Aspect an annotation of the aspect
//
// @Documented
//
// @Target("TYPE")
//
// @Retention("RUNTIME")
//
// usage: @Aspect or @Aspect("beanName")
//
/*
	// AuthAspect an auth aspect
	// @Aspect
	// @Aspect("authAspect")
	type AuthAspect struct {
	}
...
*/
type Aspect struct {
	Name    string `json:"name"`    // the bean name of the aspect
	Struct  string `json:"struct"`  // the target struct name of the @Aspect modification
	PkgPath string `json:"pkgPath"` // the pkg path of the target struct
}
