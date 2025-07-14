package options

import (
	"github.com/Lilypad-Tech/lilypad/v2/pkg/adminService"
)

func GetDefaultAdminServiceOptions() adminService.AdminServiceClientOptions {
	return adminService.AdminServiceClientOptions{
		BaseURL: GetDefaultServeOptionString("ADMIN_BASE_URL", ""),
		ApiKey:  GetDefaultServeOptionString("ADMIN_API_KEY", ""),
	}
}
