package utils

import (
	"math/rand"
	"server/database"
	"server/model"
	"strconv"
	"time"
)

func GenerateOTP() string {
	otp := ""
	for i := 0; i < 4; i++ {
		otp += strconv.Itoa(rand.Intn(10))
	}
	return otp
}

func InsertOrUpdateOTP(tempToken, otp string) error {
	db := database.DB
	err := db.Where(model.OtpQueue{TempToken: tempToken}).
		Assign(model.OtpQueue{TempToken: tempToken, Otp: otp, ExpiresAt: time.Now().Add(10 * time.Minute)}).
		FirstOrCreate(&model.OtpQueue{}).Error
	return err
}
