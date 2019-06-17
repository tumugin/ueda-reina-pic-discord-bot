package reinalibs

import (
	"gotest.tools/assert"
	"testing"
)

func TestIsReinaCalling(t *testing.T) {
	// match
	assert.Assert(t, IsReinaCalling("うえしゃま\u301C"))
	assert.Assert(t, IsReinaCalling("うえしゃま\uFF5E"))
	assert.Assert(t, IsReinaCalling("うえしゃまぁああああぁぁああ"))
	assert.Assert(t, IsReinaCalling("うえしゃまあああああああああ"))
	assert.Assert(t, IsReinaCalling("うえしゃまぁぁぁぁぁぁぁぁぁ"))
	// no match
	assert.Assert(t, IsReinaCalling("うえしゃま") == false)
	assert.Assert(t, IsReinaCalling("うえしゃまはかわいい") == false)
	assert.Assert(t, IsReinaCalling("もちょだよ〜") == false)
}
