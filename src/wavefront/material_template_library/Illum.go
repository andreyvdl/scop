package materialtemplatelibrary

import "fmt"

/*
0. Color on and Ambient off

1. Color on and Ambient on

2. Highlight on

3. Reflection on and Ray trace on

4. Transparency: Glass on, Reflection: Ray trace on

5. Reflection: Fresnel on and Ray trace on

6. Transparency: Refraction on, Reflection: Fresnel off and Ray trace on

7. Transparency: Refraction on, Reflection: Fresnel on and Ray trace on

8. Reflection on and Ray trace off

9. Transparency: Glass on, Reflection: Ray trace off

10. Casts shadows onto invisible surfaces
*/
type Illum uint32

const (
	COLOR_ON_AND_AMBIENT_OFF Illum = iota
	COLOR_ON_AND_AMBIENT_ON
	HIGHLIGHT_ON
	REFLECTION_ON_AND_RAY_TRACE_ON
	TRANSPARENCY_GLASS_ON_REFLECTION_RAY_TRACE_ON
	REFLECTION_FRESNEL_ON_AND_RAY_TRACE_ON
	TRANSPARENCY_REFRACTION_ON_REFLECTION_FRESNEL_OFF_AND_RAY_TRACE_ON
	TRANSPARENCY_REFRACTION_ON_REFLECTION_FRESNEL_ON_AND_RAY_TRACE_ON
	REFLECTION_ON_AND_RAY_TRACE_OFF
	TRANSPARENCY_GLASS_ON_REFLECTION_RAY_TRACE_OFF
	CASTS_SHADOWS_ONTO_INVISIBLE_SURFACES
)

var illum_map = map[string]Illum{
	"0":  COLOR_ON_AND_AMBIENT_OFF,
	"1":  COLOR_ON_AND_AMBIENT_ON,
	"2":  HIGHLIGHT_ON,
	"3":  REFLECTION_ON_AND_RAY_TRACE_ON,
	"4":  TRANSPARENCY_GLASS_ON_REFLECTION_RAY_TRACE_ON,
	"5":  REFLECTION_FRESNEL_ON_AND_RAY_TRACE_ON,
	"6":  TRANSPARENCY_REFRACTION_ON_REFLECTION_FRESNEL_OFF_AND_RAY_TRACE_ON,
	"7":  TRANSPARENCY_REFRACTION_ON_REFLECTION_FRESNEL_ON_AND_RAY_TRACE_ON,
	"8":  REFLECTION_ON_AND_RAY_TRACE_OFF,
	"9":  TRANSPARENCY_GLASS_ON_REFLECTION_RAY_TRACE_OFF,
	"10": CASTS_SHADOWS_ONTO_INVISIBLE_SURFACES,
}

func NewIllum(value string) *Illum {
	i, ok := illum_map[value]
	if !ok {
		panic(fmt.Sprint("Invalid value for illum: ", value))
	}

	return &i
}
