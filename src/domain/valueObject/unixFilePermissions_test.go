package valueObject

import "testing"

func TestUnixFilePermissions(t *testing.T) {
	t.Run("ValidUnixFilePermissions", func(t *testing.T) {
		validUnixFilePermissions := []interface{}{
			"000", "500", "555", "644", "755", "1644", "4755",
		}

		for _, permissions := range validUnixFilePermissions {
			_, err := NewUnixFilePermissions(permissions)
			if err != nil {
				t.Errorf(
					"Expected no error for '%v', got '%s'", permissions, err.Error(),
				)
			}
		}
	})

	t.Run("InvalidUnixFilePermissions", func(t *testing.T) {
		invalidUnixFilePermissions := []interface{}{
			"", "0", "5", "55", "-1", "7778", "00000", "aaaaa", "b",
		}

		for _, permissions := range invalidUnixFilePermissions {
			_, err := NewUnixFilePermissions(permissions)
			if err == nil {
				t.Errorf("Expected error for '%v', got nil", permissions)
			}
		}
	})
}
