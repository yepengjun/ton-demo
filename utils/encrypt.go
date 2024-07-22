package utils

// xorEncryptDecrypt 对字符串进行异或加密和解密

func XorEncryptDecrypt(input string, key byte) string {
	result := make([]byte, len(input))
	for i := 0; i < len(input); i++ {
		result[i] = input[i] ^ key
	}
	return string(result)
}
