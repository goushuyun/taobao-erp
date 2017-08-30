package key_word_filter

import "testing"

func TestFilter(t *testing.T) {
	t.Logf("-------%s------", FilterKeyWords("中国机械工程学科教程配套系列教材暨教育部高等学校机械设计制造及其自动化专业教学指）机械自动化基础"))
}
