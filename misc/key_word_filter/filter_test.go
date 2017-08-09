package key_word_filter

import "testing"

func TestFilter(t *testing.T) {
	t.Logf("-------%s------", FilterKeyWords("刘云山， 中国最大的醉人"))
}
