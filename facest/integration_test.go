// +build integration

package facest

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetFaces(t *testing.T) {
	c := NewClient(os.Getenv("FACEST_INTEGRATION_API_KEY"))

	ctx := context.Background()
	res, err := c.GetFaces(ctx, nil)

	assert.NotNil(t, res, "expecting non-nil result")
	assert.Nil(t, err, "expecting nil err")

	assert.Equal(t, 1, res.Count, "expected 1 face found")
	assert.Equal(t, 1, res.PagesCount, "expected 1 page found")

	assert.Equal(t, "integration_face_id", res.Faces[0].FaceID, "expected correct FaceID")
}
