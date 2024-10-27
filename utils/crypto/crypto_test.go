package crypto

import "testing"

func TestGetHash(t *testing.T) {
	key := "damanhhao3004@gmail.com"
	expectedHash := "f6ef8cb8a83a9f43b53b22ef9c65bdc71552e697ebebe2c6e242f47775f5d0b4"

	hash := GetHash(key)
	if hash != expectedHash {
		t.Errorf("Hash không khớp. Kết quả nhận được: %s, Kết quả mong muốn: %s", hash, expectedHash)
	}
	emptyKey := ""
	expectedEmptyHash := "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855" // giá trị băm cho chuỗi rỗng
	emptyHash := GetHash(emptyKey)

	if emptyHash != expectedEmptyHash {
		t.Errorf("Hash không khớp cho chuỗi rỗng. Kết quả nhận được: %s, Kết quả mong muốn: %s", emptyHash, expectedEmptyHash)
	}
}
