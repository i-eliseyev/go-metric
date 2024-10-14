package security

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/gofiber/fiber/v2/log"
	"github.com/i-eliseyev/go-metric/internal/common"
)

func SignMetric(metric *common.Metric, secret *string) {
	if *secret == "" {
		log.Warn("Secret key does not exist")
		return
	}
	stringToHash := makeStringToHash(metric)
	metric.Hash = calculateHash(stringToHash, secret)
}

func ValidateSignature(metric *common.Metric, secret *string) bool {
	stringToHash := makeStringToHash(metric)
	hash := calculateHash(stringToHash, secret)
	return hmac.Equal([]byte(hash), []byte(metric.Hash))
}

func makeStringToHash(metric *common.Metric) *string {
	var stringToHash string
	if metric.MType == "counter" {
		stringToHash = fmt.Sprintf("%s:counter:%d", metric.ID, *metric.Delta)
	} else {
		stringToHash = fmt.Sprintf("%s:gauge:%d", metric.ID, int(*metric.Value))
	}
	return &stringToHash
}

func calculateHash(stringToHash *string, secret *string) string {
	h := hmac.New(sha256.New, []byte(*secret))
	h.Write([]byte(*stringToHash))
	hash := h.Sum(nil)
	return hex.EncodeToString(hash)
}
