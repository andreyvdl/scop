#version 450 core
#pragma optimize(on)

#define PINK vec4(1.0, 0.0, 1.0, 1.0)
#define BLACK vec4(0.0, 0.0, 0.0, 1.0)

out vec4 FragColor;

in vec2 texCoord;

uniform float transition = 0;
uniform sampler2D tex1;

vec4 getFaceColor() {
    switch (gl_PrimitiveID % 4) {
        case 0:
        return PINK;
        case 1:
        return BLACK;
        case 2:
        return mix(PINK, BLACK, 0.3);
        case 3:
        return mix(PINK, BLACK, 0.7);
    }
}

void main() {
    vec4 tex_color = texture(tex1, texCoord);
    vec4 face_color = getFaceColor();

    FragColor = mix(tex_color, face_color, transition);
}
