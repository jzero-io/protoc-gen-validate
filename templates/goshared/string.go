package goshared

const strTpl = `
	{{ $f := .Field }}{{ $r := .Rules }}

	{{ if $r.GetIgnoreEmpty }}
		if {{ accessor . }} != "" {
	{{ end }}

	{{ template "const" . }}
	{{ template "in" . }}

	{{ if or $r.Len (and $r.MinLen $r.MaxLen (eq $r.GetMinLen $r.GetMaxLen)) }}
		{{ if $r.Len }}
		if utf8.RuneCountInString({{ accessor . }}) != {{ $r.GetLen }} {
			{{ if $r.GetLenMessage }}
				err := {{ err . $r.GetLenMessage }}
			{{ else }}
				err := {{ err . "value length must be " $r.GetLen " runes" }}
			{{ end }}
			if !all { return err }
			errors = append(errors, err)
		{{ else }}
		if utf8.RuneCountInString({{ accessor . }}) != {{ $r.GetMinLen }} {
			{{ if $r.GetMinLenMessage }}
				err := {{ err . $r.GetMinLenMessage }}
			{{ else }}
				err := {{ err . "value length must be " $r.GetMinLen " runes" }}
			{{ end }}
			if !all { return err }
			errors = append(errors, err)
		{{ end }}
	}
	{{ else if $r.MinLen }}
		{{ if $r.MaxLen }}
			if l := utf8.RuneCountInString({{ accessor . }}); l < {{ $r.GetMinLen }} || l > {{ $r.GetMaxLen }} {
				err := {{ err . "value length must be between " $r.GetMinLen " and " $r.GetMaxLen " runes, inclusive" }}
				if !all { return err }
				errors = append(errors, err)
			}
		{{ else }}
			if utf8.RuneCountInString({{ accessor . }}) < {{ $r.GetMinLen }} {
				{{ if $r.GetMinLenMessage }}
					err := {{ err . $r.GetMinLenMessage }}
				{{ else }}
					err := {{ err . "value length must be at least " $r.GetMinLen " runes" }}
				{{ end }}
				if !all { return err }
				errors = append(errors, err)
			}
		{{ end }}
	{{ else if $r.MaxLen }}
		if utf8.RuneCountInString({{ accessor . }}) > {{ $r.GetMaxLen }} {
			{{ if $r.GetMaxLenMessage }}
				err := {{ err . $r.GetMaxLenMessage }}
			{{ else }}
				err := {{ err . "value length must be at most " $r.GetMaxLen " runes" }}
			{{ end }}
			if !all { return err }
			errors = append(errors, err)
		}
	{{ end }}

	{{ if or $r.LenBytes (and $r.MinBytes $r.MaxBytes (eq $r.GetMinBytes $r.GetMaxBytes)) }}
		{{ if $r.LenBytes }}
			if len({{ accessor . }}) != {{ $r.GetLenBytes }} {
				{{ if $r.GetLenBytesMessage }}
					err := {{ err . $r.GetLenBytesMessage }}
				{{ else }}
					err := {{ err . "value length must be " $r.GetLenBytes " bytes" }}
				{{ end }}
				if !all { return err }
				errors = append(errors, err)
			}
		{{ else }}
			if len({{ accessor . }}) != {{ $r.GetMinBytes }} {
				{{ if $r.GetMinBytesMessage }}
					err := {{ err . $r.GetMinBytesMessage }}
				{{ else }}
					err := {{ err . "value length must be " $r.GetMinBytes " bytes" }}
				{{ end }}
				if !all { return err }
				errors = append(errors, err)
			}
		{{ end }}
	{{ else if $r.MinBytes }}
		{{ if $r.MaxBytes }}
			if l := len({{ accessor . }}); l < {{ $r.GetMinBytes }} || l > {{ $r.GetMaxBytes }} {
					err := {{ err . "value length must be between " $r.GetMinBytes " and " $r.GetMaxBytes " bytes, inclusive" }}
					if !all { return err }
					errors = append(errors, err)
			}
		{{ else }}
			if len({{ accessor . }}) < {{ $r.GetMinBytes }} {
				{{ if $r.GetMinBytesMessage }}
					err := {{ err . $r.GetMinBytesMessage }}
				{{ else }}
					err := {{ err . "value length must be at least " $r.GetMinBytes " bytes" }}
				{{ end }}
				if !all { return err }
				errors = append(errors, err)
			}
		{{ end }}
	{{ else if $r.MaxBytes }}
		if len({{ accessor . }}) > {{ $r.GetMaxBytes }} {
			{{ if $r.GetMaxBytesMessage }}
				err := {{ err . $r.GetMaxBytesMessage }}
			{{ else }}
				err := {{ err . "value length must be at most " $r.GetMaxBytes " bytes" }}
			{{ end }}
			if !all { return err }
			errors = append(errors, err)
		}
	{{ end }}

	{{ if $r.Prefix }}
		if !strings.HasPrefix({{ accessor . }}, {{ lit $r.GetPrefix }}) {
			{{ if $r.GetPrefixMessage }}
				err := {{ err . $r.GetPrefixMessage }}
			{{ else }}
				err := {{ err . "value does not have prefix " (lit $r.GetPrefix) }}
			{{ end }}
			if !all { return err }
			errors = append(errors, err)
		}
	{{ end }}

	{{ if $r.Suffix }}
		if !strings.HasSuffix({{ accessor . }}, {{ lit $r.GetSuffix }}) {
			{{ if $r.GetSuffixMessage }}
				err := {{ err . $r.GetSuffixMessage }}
			{{ else }}
				err := {{ err . "value does not have suffix " (lit $r.GetSuffix) }}
			{{ end }}
			if !all { return err }
			errors = append(errors, err)
		}
	{{ end }}

	{{ if $r.Contains }}
		if !strings.Contains({{ accessor . }}, {{ lit $r.GetContains }}) {
			{{ if $r.GetContainsMessage }}
				err := {{ err . $r.GetContainsMessage }}
			{{ else }}
				err := {{ err . "value does not contain substring " (lit $r.GetContains) }}
			{{ end }}
			if !all { return err }
			errors = append(errors, err)
		}
	{{ end }}

	{{ if $r.NotContains }}
		if strings.Contains({{ accessor . }}, {{ lit $r.GetNotContains }}) {
			{{ if $r.GetNotContainsMessage }}
				err := {{ err . $r.GetNotContainsMessage }}
			{{ else }}
				err := {{ err . "value contains substring " (lit $r.GetNotContains) }}
			{{ end }}
			if !all { return err }
			errors = append(errors, err)
		}
	{{ end }}

	{{ if $r.GetIp }}
		if ip := net.ParseIP({{ accessor . }}); ip == nil {
			{{ if $r.GetIpMessage }}
				err := {{ err . $r.GetIpMessage }}
			{{ else }}
				err := {{ err . "value must be a valid IP address" }}
			{{ end }}
			if !all { return err }
			errors = append(errors, err)
		}
	{{ else if $r.GetIpv4 }}
		if ip := net.ParseIP({{ accessor . }}); ip == nil || ip.To4() == nil {
			{{ if $r.GetIpv4Message }}
				err := {{ err . $r.GetIpv4Message }}
			{{ else }}
				err := {{ err . "value must be a valid IPv4 address" }}
			{{ end }}
			if !all { return err }
			errors = append(errors, err)
		}
	{{ else if $r.GetIpv6 }}
		if ip := net.ParseIP({{ accessor . }}); ip == nil || ip.To4() != nil {
			{{ if $r.GetIpv6Message }}
				err := {{ err . $r.GetIpv6Message }}
			{{ else }}
				err := {{ err . "value must be a valid IPv6 address" }}
			{{ end }}
			if !all { return err }
			errors = append(errors, err)
		}
	{{ else if $r.GetEmail }}
		if err := m._validateEmail({{ accessor . }}); err != nil {
			{{ if $r.GetEmailMessage }}
				err = {{ errCause . "err" $r.GetEmailMessage }}
			{{ else }}
				err = {{ errCause . "err" "value must be a valid email address" }}
			{{ end }}
			if !all { return err }
			errors = append(errors, err)
		}
	{{ else if $r.GetHostname }}
		if err := m._validateHostname({{ accessor . }}); err != nil {
			{{ if $r.GetHostnameMessage }}
				err = {{ errCause . "err" $r.GetHostnameMessage }}
			{{ else }}
				err = {{ errCause . "err" "value must be a valid hostname" }}
			{{ end }}
			if !all { return err }
			errors = append(errors, err)
		}
	{{ else if $r.GetAddress }}
		if err := m._validateHostname({{ accessor . }}); err != nil {
			if ip := net.ParseIP({{ accessor . }}); ip == nil {
				{{ if $r.GetAddressMessage }}
					err = {{ err . $r.GetAddressMessage }}
				{{ else }}
					err := {{ err . "value must be a valid hostname, or ip address" }}
				{{ end }}
				if !all { return err }
				errors = append(errors, err)
			}
		}
	{{ else if $r.GetUri }}
		if uri, err := url.Parse({{ accessor . }}); err != nil {
			{{ if $r.GetUriMessage }}
				err = {{ errCause . "err" $r.GetUriMessage }}
			{{ else }}
				err = {{ errCause . "err" "value must be a valid URI" }}
			{{ end }}
			if !all { return err }
			errors = append(errors, err)
		} else if !uri.IsAbs() {
			err := {{ err . "value must be absolute" }}
			if !all { return err }
			errors = append(errors, err)
		}
	{{ else if $r.GetUriRef }}
		if _, err := url.Parse({{ accessor . }}); err != nil {
			{{ if $r.GetUriRefMessage }}
				err = {{ errCause . "err" $r.GetUriRefMessage }}
			{{ else }}
				err = {{ errCause . "err" "value must be a valid URI" }}
			{{ end }}
			if !all { return err }
			errors = append(errors, err)
		}
	{{ else if $r.GetUuid }}
		if err := m._validateUuid({{ accessor . }}); err != nil {
			{{ if $r.GetUuidMessage }}
				err = {{ errCause . "err" $r.GetUuidMessage }}
			{{ else }}
				err = {{ errCause . "err" "value must be a valid UUID" }}
			{{ end }}
			if !all { return err }
			errors = append(errors, err)
		}
	{{ end }}

	{{ if $r.Pattern }}
		if !{{ lookup $f "Pattern" }}.MatchString({{ accessor . }}) {
			{{ if $r.GetPatternMessage }}
				err := {{ err . $r.GetPatternMessage }}
			{{ else }}
				err := {{ err . "value does not match regex pattern " (lit $r.GetPattern) }}
			{{ end }}
			if !all { return err }
			errors = append(errors, err)
		}
	{{ end }}

	{{ if $r.GetIgnoreEmpty }}
		}
	{{ end }}

`
