package services

import "errors"

var (
	ErrEmptyRequest       = errors.New("пустой запрос")
	ErrPaymentNotComplete = errors.New("платеж не завершен")
	ErrRefundAmountExceed = errors.New("сумма возврата превышает доступную")
	
)
