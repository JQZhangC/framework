package pass

import (
	"fmt"
	"net/url"
	"regexp"
)

type RuleConfig struct {
	MinRune  int
	HasAlaph bool
	HasNum   bool
}

var _defRuleConfig = RuleConfig{8, true, true}

var (
	_pwdExpAlaph = regexp.MustCompile(`[a-zA-Z]+`)
	_pwdExpNum   = regexp.MustCompile(`[0-9]+`)
)

// CheckPass 检查密码强度
func CheckPass(pass string, cfgs ...RuleConfig) error {
	var cfg = _defRuleConfig
	if len(cfgs) > 0 {
		cfg = cfgs[0]
	}
	if len(pass) < cfg.MinRune {
		return fmt.Errorf("密码小于%d位", cfg.MinRune)
	}
	if cfg.HasAlaph && !_pwdExpAlaph.MatchString(pass) {
		return fmt.Errorf("密码未包含字符")
	}
	if cfg.HasNum && !_pwdExpNum.MatchString(pass) {
		return fmt.Errorf("密码未包含数字")
	}
	return nil
}

var _expSchema = regexp.MustCompile(`^.+://`)

// ParseDsnURL 提取url规范的dsn上的账号信息
func ParseDsnURL(dsn string) (user, pass string) {
	if !_expSchema.MatchString(dsn) {
		dsn = fmt.Sprintf("mock://%s", dsn)
	}
	u, err := url.Parse(dsn)
	if err != nil || u.User == nil {
		return
	}
	user = u.User.Username()
	pass, _ = u.User.Password()
	return
}
