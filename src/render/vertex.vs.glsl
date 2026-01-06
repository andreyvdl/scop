#version 450 core
layout(location = 0) in vec4 aPos;
layout(location = 1) in vec2 aTex;

out vec2 texCoord;

uniform mat4 transform;

void main() {
    gl_Position = transform * aPos;
    texCoord = vec2(aTex.x, aTex.y);
}
