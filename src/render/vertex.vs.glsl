#version 450 core
#pragma optimize(on)

layout(location = 0) in vec4 aPos;
layout(location = 1) in vec2 aTex;

out vec2 texCoord;

uniform float pitch;
uniform float yaw;
uniform float roll;
uniform float zoom;
uniform float width;
uniform float height;

uniform vec3 object_center;
uniform vec3 z_vec;
uniform vec3 dislocation;
uniform vec3 camera;
uniform vec3 world_up;

uniform mat4 model;
uniform mat4 view;
uniform mat4 proj;

vec4 quaternionFromEuler(float p_deg, float y_deg, float r_deg) {
    float p = radians(p_deg) * 0.5;
    float y = radians(y_deg) * 0.5;
    float r = radians(r_deg) * 0.5;

    float cy = cos(y);
    float cp = cos(p);
    float cr = cos(r);
    float sy = sin(y);
    float sp = sin(p);
    float sr = sin(r);
    return normalize(vec4(
            sp * cr * cy - sr * sy * cp,
            sp * sr * cy + sy * cp * cr,
            sp * sy * cr + sr * cp * cy,
            -sp * sr * sy + cp * cr * cy
        ));
}

mat4 quaternionToMatrix(vec4 q) {
    q = normalize(q);

    mat4 id = mat4(
            1.0, 0.0, 0.0, 0.0,
            0.0, 1.0, 0.0, 0.0,
            0.0, 0.0, 1.0, 0.0,
            0.0, 0.0, 0.0, 1.0
        );

    float xx = q.x * q.x;
    float xy = q.x * q.y;
    float xz = q.x * q.z;
    float xw = q.x * q.w;

    float yy = q.y * q.y;
    float yz = q.y * q.z;
    float yw = q.y * q.w;

    float zz = q.z * q.z;
    float zw = q.z * q.w;

    id[0].x = 1 - 2 * (yy + zz);
    id[0].y = 2 * (xy - zw);
    id[0].z = 2 * (xz + yw);

    id[1].x = 2 * (xy + zw);
    id[1].y = 1 - 2 * (xx + zz);
    id[1].z = 2 * (yz - xw);

    id[2].x = 2 * (xz - yw);
    id[2].y = 2 * (yz + xw);
    id[2].z = 1 - 2 * (xx + yy);

    return id;
}

mat4 translate(vec3 translation) {
    return mat4(
        1.0, 0.0, 0.0, 0.0,
        0.0, 1.0, 0.0, 0.0,
        0.0, 0.0, 1.0, 0.0,
        translation.x, translation.y, translation.z, 1
    );
}

mat4 lookAt(vec3 camera, vec3 target, vec3 upside) {
    vec3 forward = normalize(camera - target);
    vec3 right = normalize(cross(upside, forward));
    vec3 up = cross(forward, right);

    mat4 id = mat4(
            right.x, up.x, forward.x, 0.0,
            right.y, up.y, forward.y, 0.0,
            right.z, up.z, forward.z, 0.0,
            0.0, 0.0, 0.0, 1.0
        );

    mat4 id2 = mat4(
            1.0, 0.0, 0.0, 0.0,
            0.0, 1.0, 0.0, 0.0,
            0.0, 0.0, 1.0, 0.0,
            -camera.x, -camera.y, -camera.z, 1.0
        );

    return id * id2;
}

mat4 perspective(float fov_y, float aspect, float z_near, float z_far) {
    mat4 m = mat4(0.0);
    float f = 1.0 / tan(fov_y / 2.0);

    m[0].x = 1 / (f * aspect);
    m[1].y = 1 / f;
    m[2].z = -(z_far + z_near) / (z_far - z_near);
    m[3].z = -(2 * z_far * z_near) / (z_far - z_near);
    m[2].w = -1;

    return m;
}

void main() {
    vec4 new_position = aPos - vec4(object_center, 0.0);
    vec4 orientation = quaternionFromEuler(pitch, yaw, roll);
    mat4 m = translate(dislocation) * quaternionToMatrix(orientation);

    vec4 tgt = translate(dislocation) * vec4(0.0, 0.0, 0.0, 1.0);
    mat4 v = lookAt(camera, tgt.xyz, world_up);

    mat4 p = perspective(radians(zoom), width / height, 0.1, 1000.0);

    gl_Position = p * v * m * new_position;
    texCoord = vec2(aTex.x, aTex.y);
}
