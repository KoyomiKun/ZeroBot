package zero

// New 生成空引擎
func New() *Engine {
	return &Engine{rules: []Rule{}}
}

var defaultEngine = New()

// Engine is the
type Engine struct {
	rules []Rule
}

// With 继承该 Engine 创建一个新 Engine
func (e *Engine) With(rules ...Rule) *Engine {
	engine := &Engine{
		rules: make([]Rule, len(e.rules)),
	}
	copy(engine.rules, engine.rules)              // copy the raw rules
	engine.rules = append(engine.rules, rules...) // append the rule
	return engine
}

// Use 向该 Engine 添加新 Rule
func (e *Engine) Use(rules ...Rule) {
	e.rules = append(e.rules, rules...)
}

// On 添加新的指定消息类型的匹配器(默认Engine)
func On(typ string, rules ...Rule) *Matcher { return defaultEngine.On(typ, rules...) }

// On 添加新的指定消息类型的匹配器
func (e *Engine) On(typ string, rules ...Rule) *Matcher {
	var matcher = &Matcher{
		Type:  Type(typ),
		Rules: append(e.rules, rules...),
	}
	return StoreMatcher(matcher)
}

// OnMessage 消息触发器
func OnMessage(rules ...Rule) *Matcher { return On("message", rules...) }

// OnMessage 消息触发器
func (e *Engine) OnMessage(rules ...Rule) *Matcher { return e.On("message", rules...) }

// OnNotice 系统提示触发器
func OnNotice(rules ...Rule) *Matcher { return On("notice", rules...) }

// OnNotice 系统提示触发器
func (e *Engine) OnNotice(rules ...Rule) *Matcher { return e.On("notice", rules...) }

// OnRequest 请求消息触发器
func OnRequest(rules ...Rule) *Matcher { return On("request", rules...) }

// OnRequest 请求消息触发器
func (e *Engine) OnRequest(rules ...Rule) *Matcher { return On("request", rules...) }

// OnMetaEvent 元事件触发器
func OnMetaEvent(rules ...Rule) *Matcher { return On("meta_event", rules...) }

// OnMetaEvent 元事件触发器
func (e *Engine) OnMetaEvent(rules ...Rule) *Matcher { return On("meta_event", rules...) }

// OnPrefix 前缀触发器
func OnPrefix(prefix string, rules ...Rule) *Matcher { return defaultEngine.OnPrefix(prefix, rules...) }

// OnPrefix 前缀触发器
func (e *Engine) OnPrefix(prefix string, rules ...Rule) *Matcher {
	var matcher = &Matcher{
		Type:  Type("message"),
		Rules: append(append([]Rule{PrefixRule(prefix)}, e.rules...), rules...),
	}
	return StoreMatcher(matcher)
}

// OnSuffix 后缀触发器
func OnSuffix(suffix string, rules ...Rule) *Matcher { return defaultEngine.OnSuffix(suffix, rules...) }

// OnSuffix 后缀触发器
func (e *Engine) OnSuffix(suffix string, rules ...Rule) *Matcher {
	var matcher = &Matcher{
		Type:  Type("message"),
		Rules: append(append([]Rule{SuffixRule(suffix)}, e.rules...), rules...),
	}
	return StoreMatcher(matcher)
}

// OnCommand 命令触发器
func OnCommand(commands string, rules ...Rule) *Matcher {
	return defaultEngine.OnCommand(commands, rules...)
}

// OnCommand 命令触发器
func (e *Engine) OnCommand(commands string, rules ...Rule) *Matcher {
	var matcher = &Matcher{
		Type:  Type("message"),
		Rules: append(append([]Rule{CommandRule(commands)}, e.rules...), rules...),
	}
	return StoreMatcher(matcher)
}

// OnRegex 正则触发器
func OnRegex(regexPattern string, rules ...Rule) *Matcher {
	return OnMessage(append([]Rule{RegexRule(regexPattern)}, rules...)...)
}

// OnRegex 正则触发器
func (e *Engine) OnRegex(regexPattern string, rules ...Rule) *Matcher {
	var matcher = &Matcher{
		Type:  Type("message"),
		Rules: append(append([]Rule{RegexRule(regexPattern)}, e.rules...), rules...),
	}
	return StoreMatcher(matcher)
}

// OnKeyword 关键词触发器
func OnKeyword(keyword string, rules ...Rule) *Matcher {
	return defaultEngine.OnKeyword(keyword, rules...)
}

// OnKeyword 关键词触发器
func (e *Engine) OnKeyword(keyword string, rules ...Rule) *Matcher {
	var matcher = &Matcher{
		Type:  Type("message"),
		Rules: append(append([]Rule{KeywordRule(keyword)}, e.rules...), rules...),
	}
	return StoreMatcher(matcher)
}

// OnFullMatch 完全匹配触发器
func OnFullMatch(src string, rules ...Rule) *Matcher {
	return defaultEngine.OnFullMatch(src, rules...)
}

// OnFullMatch 完全匹配触发器
func (e *Engine) OnFullMatch(src string, rules ...Rule) *Matcher {
	var matcher = &Matcher{
		Type:  Type("message"),
		Rules: append(append([]Rule{FullMatchRule(src)}, e.rules...), rules...),
	}
	return StoreMatcher(matcher)
}

// OnFullMatchGroup 完全匹配触发器组
func OnFullMatchGroup(src []string, rules ...Rule) *Matcher {
	return defaultEngine.OnFullMatchGroup(src, rules...)
}

// OnFullMatchGroup 完全匹配触发器组
func (e *Engine) OnFullMatchGroup(src []string, rules ...Rule) *Matcher {
	var matcher = &Matcher{
		Type:  Type("message"),
		Rules: append(append([]Rule{FullMatchRule(src...)}, e.rules...), rules...),
	}
	return StoreMatcher(matcher)
}

// OnKeywordGroup 关键词触发器组
func OnKeywordGroup(keywords []string, rules ...Rule) *Matcher {
	return defaultEngine.OnKeywordGroup(keywords, rules...)
}

// OnKeywordGroup 关键词触发器组
func (e *Engine) OnKeywordGroup(keywords []string, rules ...Rule) *Matcher {
	var matcher = &Matcher{
		Type:  Type("message"),
		Rules: append(append([]Rule{KeywordRule(keywords...)}, e.rules...), rules...),
	}
	return StoreMatcher(matcher)
}

// OnCommandGroup 命令触发器组
func OnCommandGroup(commands []string, rules ...Rule) *Matcher {
	return defaultEngine.OnCommandGroup(commands, rules...)
}

// OnCommandGroup 命令触发器组
func (e *Engine) OnCommandGroup(commands []string, rules ...Rule) *Matcher {
	var matcher = &Matcher{
		Type:  Type("message"),
		Rules: append(append([]Rule{CommandRule(commands...)}, e.rules...), rules...),
	}
	return StoreMatcher(matcher)
}

// OnPrefixGroup 前缀触发器组
func OnPrefixGroup(prefix []string, rules ...Rule) *Matcher {
	return defaultEngine.OnPrefixGroup(prefix, rules...)
}

// OnPrefixGroup 前缀触发器组
func (e *Engine) OnPrefixGroup(prefix []string, rules ...Rule) *Matcher {
	var matcher = &Matcher{
		Type:  Type("message"),
		Rules: append(append([]Rule{PrefixRule(prefix...)}, e.rules...), rules...),
	}
	return StoreMatcher(matcher)
}

// OnSuffixGroup 后缀触发器组
func OnSuffixGroup(suffix []string, rules ...Rule) *Matcher {
	return defaultEngine.OnSuffixGroup(suffix, rules...)
}

// OnSuffixGroup 后缀触发器组
func (e *Engine) OnSuffixGroup(suffix []string, rules ...Rule) *Matcher {
	var matcher = &Matcher{
		Type:  Type("message"),
		Rules: append(append([]Rule{SuffixRule(suffix...)}, e.rules...), rules...),
	}
	return StoreMatcher(matcher)
}