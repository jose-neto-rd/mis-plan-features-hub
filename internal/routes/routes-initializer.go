package routes

import (
	"mis-plan-features-hub/internal/core/domain/interfaces"
	"reflect"
	"strings"
)

type RouterInitializer struct {
	Logger     interfaces.Logger
	Components []interfaces.HttpRouterInit
}

func (r *RouterInitializer) Register(component interfaces.HttpRouterInit) {
	r.Components = append(r.Components, component)
}

func (r *RouterInitializer) Init() error {
	for _, component := range r.Components {
		if component == nil {
			continue
		}
		fullType := reflect.TypeOf(component).String()
		parts := strings.Split(fullType, ".")
		structName := parts[len(parts)-1]
		r.Logger.Info("Initializing router: " + structName)
		if err := component.Init(); err != nil {
			r.Logger.Error("Error on router: " + err.Error())
		}
	}
	return nil
}
