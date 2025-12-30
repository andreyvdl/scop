#version 450 core
layout(location = 0) in vec4 aPos;
layout(location = 1) in vec3 aColor;
layout(location = 2) in vec2 aTex;

out vec3 ourColor;
out vec2 texCoord;

void main() {
    gl_Position = aPos;
    ourColor = aColor;
    texCoord = aTex;
}
