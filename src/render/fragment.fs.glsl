#version 450 core
out vec4 FragColor;

in vec2 texCoord;

uniform float transition = 0;
uniform sampler2D tex1;
uniform sampler2D tex2;

void main() {
    FragColor = mix(texture(tex1, texCoord), vec4(0.5, 0.5, 0.5, 1), transition);
}
