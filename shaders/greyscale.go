package shaders

var GreyscaleShader = `
#version 330 core

in vec2  vTexCoords;

out vec4 fragColor;

uniform vec4 uTexBounds;
uniform sampler2D uTexture;

void main() {
	// Get our current screen coordinate
	vec2 t = (vTexCoords - uTexBounds.xy) / uTexBounds.zw;

	// Sum our 3 color channels
	float sum  = texture(uTexture, t).r;
	      sum += texture(uTexture, t).g;
	      sum += texture(uTexture, t).b;

	// Divide by 3, and set the output to the result
	vec4 color = vec4( sum/3, sum/3, sum/3, 1.0);
	fragColor = color;
}`
