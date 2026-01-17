#version 450 core

#define PINK vec4(1.0, 0.0, 1.0, 1.0)
#define BLACK vec4(0.0, 0.0, 0.0, 1.0)

out vec4 FragColor;

in vec2 texCoord;

uniform float transition = 0;
uniform sampler2D tex1;
uniform sampler2D tex2;

vec4 selectCoord(vec2 coord) {
    vec4 color;

    if (coord.x > 0.5 && coord.y > 0.5) {
        color = PINK;
    }
    else if (coord.x < 0.5 && coord.y < 0.5) {
        color = BLACK;
    }
    else {
        color = mix(PINK, BLACK, 0.5);
    }

    return color;
}

void main() {
    FragColor = mix(texture(tex1, texCoord), selectCoord(texCoord), transition);
}
