package repositories

func GetWorkflowSteps() []string {
	return []string{"to do", "doing", "waiting on information", "paid", "closed without pay", "waiting on payment"}
}
