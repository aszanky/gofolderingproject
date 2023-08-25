package payment

const (
	queryUpdatePayment = `
		UPDATE payment
		SET status = 1
		WHERE id = $1;
	`
)
