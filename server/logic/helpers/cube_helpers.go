package helpers

// contains checks if a slice contains a specific string.
func ContainsColor(slice []string, str string) bool {
    for _, v := range slice {
        if v == str {
            return true
        }
    }
    return false
}
