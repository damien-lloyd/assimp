package assimp

import (
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestBonesLoaded(t *testing.T) {
	scene := ImportFile("assets/models/ArmyPilot/ArmyPilot.dae", uint(
		Process_Triangulate |
			Process_FlipUVs))

	assert.NotNil(t, scene, "scene was nil")

	root := scene.RootNode()

	m := root.Transformation()
	m.Transpose()

	children := root.Children()

	assert.Greater(t, len(children[0].Meshes()), 0)

	mesh := scene.Meshes()[children[0].Meshes()[0]]

	matrix := mesh.Bones()[0].OffsetMatrix()

	result := matrix.Mat4()

	log.Println(result.String())
}