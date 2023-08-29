package echoroutesview

import (
	"bytes"
	"fmt"
	"net/http"
	"sort"
	"text/template"

	"github.com/Masterminds/sprig/v3"
	"github.com/labstack/echo/v4"
)

type (
	Registrable interface {
		GET(string, echo.HandlerFunc, ...echo.MiddlewareFunc) *echo.Route
		Routes() []*echo.Route
	}
)

func RegisterRoutesViewer(echoInstance Registrable) error {

	tpl, err := parseTemplate(echoInstance)
	if err != nil {
		return err
	}

	echoInstance.GET("/registered-routes", func(ctx echo.Context) error {
		return ctx.HTML(http.StatusOK, tpl)
	})

	return nil
}

func getOrderedRoutes(routes []*echo.Route) []*echo.Route {
	routesCopy := make([]*echo.Route, len(routes))
	copy(routesCopy, routes)
	sort.SliceStable(routesCopy, func(i, j int) bool {
		return fmt.Sprintf("%s-%s", routesCopy[i].Path, routesCopy[i].Method) < fmt.Sprintf("%s-%s", routesCopy[j].Path, routesCopy[j].Method)
	})
	return routesCopy
}

func parseTemplate(echoInstance Registrable) (string, error) {
	var outputBuffer bytes.Buffer
	if err := template.Must(
		template.New("viewer").Funcs(sprig.FuncMap()).Parse(routerViewer),
	).Execute(&outputBuffer, getOrderedRoutes(echoInstance.Routes())); err != nil {
		return "", fmt.Errorf("can not parse template: %w", err)
	}
	return outputBuffer.String(), nil
}
