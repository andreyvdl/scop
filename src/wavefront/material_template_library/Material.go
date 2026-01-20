package materialtemplatelibrary

/*
Specular exponent "Ns"

Important to describe the specular's exponent for highlighting a spot.
After command "Ns" there is a single floating point number from 0 to 1000.

Syntax: Ns <exponent>

=========

Ambient color "Ka"

The first component important for phong illumination model.
Command is "ka" followed by 3 floating point numbers from 0 to 1 representing red, green and blue.

Syntax: Ka <r> <g> <b>

=========

Diffuse color "Kd"

The second component important for phong illumination model. Same as ambient color.

Syntax: Kd <r> <g> <b>

=========

Specular color "Ks"

The third component important for phong illumination model. Same as ambient and diffuse color.

Syntax: Ks <r> <g> <b>

=========

Refraction coefficient "Ni"

Important for raytracing (explained later).
If you are a physician, you know that transparent objects (i.e. glass and water) have a refraction coefficient where light rays get distracted from the surface when entering.
It is not important now how to calculate it.

Syntyx: Ni <refraction coefficient>

=========

Disolve factor or Transparency "d"

Talking about transparent objects like glass.
This command tells us how much of the color you will see from the object.

Syntax: d <disolve factor>

=========

Illumination model "illum"

Not really something you can figure out yourself.
It seems that OBJ files have around 10 different settings for illumination.

Syntax: illum <setting>
*/
type Material struct {
	Name                   string
	Specular_exponent      *Ns
	Ambient_color          *Ka
	Diffuse_color          *Kd
	Specular_color         *Ks
	Refraction_coefficient *Ni
	Dissolve_factor        *D
	Illumination_model     *Illum
}

func NewMaterial(name string) *Material {
	return &Material{Name: name, Refraction_coefficient: NewNi(1.0)}
}
